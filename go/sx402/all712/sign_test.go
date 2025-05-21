package all712

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var boss = common.HexToAddress("0xaab05558448C8a9597287Db9F61e2d751645B12a")

func TestSignature(t *testing.T) {

	privkeyhex := "56c11c2fee673894e85151857339066cd244d4932f23e660ce8502c867d0927e"
	privkey, err := crypto.HexToECDSA(privkeyhex)
	if err != nil {
		t.Error(err)

	}

	nonce := crypto.Keccak256Hash([]byte("SixthNonce"))
	log.Printf("nonce: 0x%x\n", nonce)
	tokenName := "EURO_S"

	from := boss
	//to := common.HexToAddress("0x64A8303112D05027F1f1d4ed7e54482799861db0")
	to := common.HexToAddress("0x8cE22df7DE8FC669D0727eeC2C99b0063d7D2ECC")
	tokenVersion := "2"

	auth := Authorization{
		From:        from.Hex(),
		To:          to.Hex(),
		Value:       "1500",
		ValidAfter:  "0",
		ValidBefore: "1748735999",
		Nonce:       nonce.Hex(),
	}

	sig, err := SignERC3009Authorization(&auth, privkey, sepoliaChainId, tokenName, tokenVersion, BaseSepoliaEURSAddress)
	if err != nil {
		t.Error(err)
	}

	log.Printf("0x%x\n", sig)
	log.Printf("sig: r: 0x%x, s: 0x%x v: 0x%x\n", sig[:32], sig[32:64], sig[64]+27)

}

func TestEncodePayment(t *testing.T) {
	jsonpayment := `{"x402Version":1,"scheme":"exact","network":"base-sepolia","payload":
{"signature":"0x01bb20f75aa19952b568cb0ae674daf4e85d27266985810219b03c913729ae0647479cddc541611fa8fd7f9d84457433183ba1c723d9e31202208d2caa674c981b",
"authorization":
{"from":"0xaab05558448C8a9597287Db9F61e2d751645B12a","to":"0xCEF702Bd69926B13ab7150624daA7aFEE0300786",
"value":"1230","validAfter":"1747130688","validBefore":"1748735999",
"nonce":"0x4e91383183648c3e72a1f317acc930227bf36e12cf376bcf6ab1e8f5e9e2a6dc"}}}`

	ep := base64.StdEncoding.EncodeToString([]byte(jsonpayment))
	log.Println(ep)
}

func TestAddSignature(t *testing.T) {
	preq := new(types.PaymentRequirements)
	err := json.Unmarshal([]byte(somePayReqs), preq)
	if err != nil {
		t.Error(err)
	}
	preq.MaxAmountRequired = "1234"
	privkeyhex := "56c11c2fee673894e85151857339066cd244d4932f23e660ce8502c867d0927e"
	privkey, err := crypto.HexToECDSA(privkeyhex)

	ppld, err := AddAuthorizationSignature(preq, privkey)
	if err != nil {
		t.Error(err)
	}

	envlp := new(Envelope)
	envlp.PaymentPayload = ppld
	envlp.PaymentRequirements = preq
	envlp.X402Version = 1

	curlstr := "curl $X402_URL -d '%s'\n"
	ebts, _ := json.Marshal(envlp)
	fmt.Printf(curlstr, string(ebts))

	extra := map[string]string{}
	err = json.Unmarshal(*preq.Extra, &extra)
	if err != nil {
		t.Error(err)
	}
	ok, recovered, err := VerifyTransferWithAuthorizationSignature(ppld.Payload.Signature, *ppld.Payload.Authorization, extra["name"], extra["version"], big.NewInt(84532), common.HexToAddress(preq.Asset))
	if err != nil {
		t.Error(err)
	}
	log.Println(ok, "recovered payer: ", recovered)
}

func TestJMignature(t *testing.T) {
	//signature := "0xa5f8b738aedbb7096d6c9eff2a00506855e4bda75454d4b61389730e7bb4acbc11470566bdde97245f8fc873140e203efa8c79a233af9c80c3f1219aa5edd0991b"
	signature := "0x98823dd1606f5fa70c60af0a85b40973a319e5fc6177fc535bba9f8a9fdb9e78556b918d851c60c8fb118ff21386dfd7065c0e40b458f4ccc39c3e2829380f2c1b"
	tokenName := "USDC"
	tokenVersion := "1"
	tokenAddress := "0x036CbD53842c5426634e7929541eC2318f3dCF7"

	auth := types.ExactEvmPayloadAuthorization{
		From:        "0xcd3c55547cda4da34c745e8f3b57698166c8aeeb",
		To:          "0x5F18bD40CF6cBbf034ff3d2003576B95E73D32e3",
		Value:       "1000000000000000000",
		ValidAfter:  "0",
		ValidBefore: "1747731903",
		Nonce:       "0xb55ba27ac38b7e4f4b9aa4289bae8813e42c39024b4c253afb2b5d3df0c6065e",
	}

	ok, recovered, err := VerifyTransferWithAuthorizationSignature(signature, auth, tokenName, tokenVersion, big.NewInt(84532), common.HexToAddress(tokenAddress))
	if err != nil {
		t.Error(err)
	}
	log.Println(ok, "recovered payer: ", recovered)
}

const message1 = `{"from" : "0xcef702bd69926b13ab7150624daa7afee0300786",
"nonce": "0x0b38c4346596ec21682d78991d9337bafe26e7b772b3c99457b1d4b00e1862d4",
"to": "0xCEF702Bd69926B13ab7150624daA7aFEE0300786",
"validAfter":"1747751973",
"validBefore": "1747752573",
"value": "2000"}`

const sig1 = `0xc7aac39eb838c7bc03d2f6a829257a0f184dd6f5752483115251014180fdd4fb710c00bae8aed1921b13cdcb84600c7206bd6d034d1cdbb92fc6e4e493be9d9c1c`

const message2 = `{"from":"0xcef702bd69926b13ab7150624daa7afee0300786","to":"0xCEF702Bd69926B13ab7150624daA7aFEE0300786","value":"2000","validAfter":"1747814778","validBefore":"1747815378","nonce":"0x75d69d62a999e935e2ad4741616c56ae36a751ebfbdcef0bb44c0203f0db650b"}`
const sig2 = `0xe6f5066174d74b16f49e28e00d83b35e14e1bb3325006213ffecafc12e2b146e538f9addef06512951693c2e7c35b29fceea9f1c0f58abb92aebca22bd9b3c7a1b`

func TestMessage(t *testing.T) {
	auth := new(types.ExactEvmPayloadAuthorization)
	err := json.Unmarshal([]byte(message2), auth)
	if err != nil {
		t.Error(err)
	}
	_, payer, err := VerifyTransferWithAuthorizationSignature(sig2, *auth, "EURO_S", "2", big.NewInt(84532), BaseSepoliaEURSAddress)
	fmt.Println(payer, err)

}

const somePayReqs = `{"scheme":"exact","network":"base-sepolia","maxAmountRequired":"15000","resource":"http://localhost:4021/weather","description":"","mimeType":"","payTo":"0x64A8303112D05027F1f1d4ed7e54482799861db0","maxTimeoutSeconds":60,"asset":"0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9","extra":{"name":"EURO_S","version":"2"}}`

const jmheader = `eyJ4NDAyVmVyc2lvbiI6MSwic2NoZW1lIjoiZXhhY3QiLCJuZXR3b3JrIjoiYmFzZS1zZXBvbGlhIiwicGF5bG9hZCI6eyJzaWduYXR1cmUiOiIweDE1YzQ1OTdhNjMxNjBjNmVlODU3MzZjODM0YzI1NDJmMmQ4OGY5NjFjNTZlYTllYTM1ZWQ2MzQzNjNiNTVhZWQxMzExZDRhNzE5ZTRhZWRmNjBhZjY5YmY5ZmQ0NjUzM2ZkNGVjODRhZGRmYThmODVlNzNkN2RmMzM5MDRkZTg3MWMiLCJhdXRob3JpemF0aW9uIjp7ImZyb20iOiIweGNkM2M1NTU0N2NkYTRkYTM0Yzc0NWU4ZjNiNTc2OTgxNjZjOGFlZWIiLCJ0byI6IjB4Q0VGNzAyQmQ2OTkyNkIxM2FiNzE1MDYyNGRhQTdhRkVFMDMwMDc4NiIsInZhbHVlIjoiMCIsInZhbGlkQWZ0ZXIiOiIwIiwidmFsaWRCZWZvcmUiOiIxNzQ3ODIyMzcyIiwibm9uY2UiOiIweGI1NWJhMjdhYzM4YjdlNGY0YjlhYTQyODliYWU4ODEzZTQyYzM5MDI0YjRjMjUzYWZiMmI1ZDNkZjBjNjA2NWUifX19`

func TestJMH(t *testing.T) {
	bts, err := base64.StdEncoding.DecodeString(jmheader)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(bts))
	ph := new(types.PaymentPayload)
	err = json.Unmarshal(bts, ph)
	if err != nil {
		t.Error(err)
	}
	auth := ph.Payload.Authorization
	_, payer, err := VerifyTransferWithAuthorizationSignature(ph.Payload.Signature, *auth, "USDC", "2", big.NewInt(84532), common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e"))
	fmt.Println(payer, err)

}
