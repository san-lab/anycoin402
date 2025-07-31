package signing

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
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/schemes"
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
	tokenName := "USDC"

	from := boss
	//to := common.HexToAddress("0x64A8303112D05027F1f1d4ed7e54482799861db0")
	to := boss
	tokenVersion := "2"

	auth := types.ExactEvmPayloadAuthorization{
		From:        from.Hex(),
		To:          to.Hex(),
		Value:       "1500",
		ValidAfter:  "0",
		ValidBefore: "1748735999",
		Nonce:       nonce.Hex(),
	}

	sig, err := SignERC3009Authorization(&auth, privkey, big.NewInt(80002), tokenName, tokenVersion, common.HexToAddress(schemes.AMOY_USDC))
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
	preq.Asset = schemes.AMOY_USDC

	preq.MaxAmountRequired = "1234"

	privkey, err := crypto.HexToECDSA(privkeyhex)
	if err != nil {
		t.Error(err)
	}

	ppld, err := CreateAuthorizationWithSignature(preq, privkey)
	if err != nil {
		t.Error(err)
	}

	envlp := new(all712.Envelope)
	ppldbytes, _ := json.Marshal(ppld.Payload)
	envlp.PaymentPayload = &all712.PaymentPayload{ppld.X402Version, ppld.Scheme, ppld.Network, ppldbytes}
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
	recovered, _, _, err := VerifyTransferWithAuthorizationSignature(ppld.Payload.Signature, *ppld.Payload.Authorization, extra["name"], extra["version"], big.NewInt(84532), common.HexToAddress(preq.Asset))
	if err != nil {
		t.Error(err)
	}
	log.Println("recovered payer: ", recovered)
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

	recovered, _, _, err := VerifyTransferWithAuthorizationSignature(signature, auth, tokenName, tokenVersion, big.NewInt(84532), common.HexToAddress(tokenAddress))
	if err != nil {
		t.Error(err)
	}
	log.Println("recovered payer: ", recovered)
}

const message1 = `{"from" : "0xaab05558448c8a9597287db9f61e2d751645b12a",
"nonce": "0x6ee9e29abfc331d7f4552fe55bd5ff45ebe67c5b0423172533dd1c882ac92a98",
"to": "0xCEF702Bd69926B13ab7150624daA7aFEE0300786",
"validAfter":"1748356918",
"validBefore": "1748357518",
"value": "2200"}`

const sig1 = `0xde7c0389cfd4fa942b67f6d243a3658019491f2db1e4a5414ac81d8cc69bb8be54331f0a77ecea4d44aa9efa2261d45e2fe0f0237ecd43af623a957104c93d5d1b`

const message2 = `{"from":"0xcef702bd69926b13ab7150624daa7afee0300786","to":"0xCEF702Bd69926B13ab7150624daA7aFEE0300786","value":"2000","validAfter":"1747814778","validBefore":"1747815378","nonce":"0x75d69d62a999e935e2ad4741616c56ae36a751ebfbdcef0bb44c0203f0db650b"}`
const sig2 = `0xe6f5066174d74b16f49e28e00d83b35e14e1bb3325006213ffecafc12e2b146e538f9addef06512951693c2e7c35b29fceea9f1c0f58abb92aebca22bd9b3c7a1b`

func TestMessage(t *testing.T) {
	auth := new(types.ExactEvmPayloadAuthorization)
	err := json.Unmarshal([]byte(message1), auth)
	if err != nil {
		t.Error(err)
	}
	payer, _, _, err := VerifyTransferWithAuthorizationSignature(sig1, *auth, "USDC", "2", big.NewInt(80002), common.HexToAddress(schemes.AMOY_USDC))
	fmt.Println(payer, err)

	privkeyhex := "56c11c2fee673894e85151857339066cd244d4932f23e660ce8502c867d0927e"
	privkey, err := crypto.HexToECDSA(privkeyhex)
	if err != nil {
		t.Error(err)

	}

	nonce := crypto.Keccak256Hash([]byte("SixthNonce"))
	auth.Nonce = nonce.Hex()

	sig2, err := SignERC3009Authorization(auth, privkey, big.NewInt(80002), "USDC", "2", common.HexToAddress(schemes.AMOY_USDC))
	payer, _, _, err = VerifyTransferWithAuthorizationSignature(fmt.Sprintf("0x%x", sig2), *auth, "USDC", "2", big.NewInt(80002), common.HexToAddress(schemes.AMOY_USDC))
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
	payer, _, _, err := VerifyTransferWithAuthorizationSignature(ph.Payload.Signature, *auth, "USDC", "2", big.NewInt(84532), common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e"))
	fmt.Println(payer, err)

}

func TestDomain(t *testing.T) {
	dsh := all712.MakeDomainSeparator("USDC", "2", evmbinding.ChainIDs["amoy"], common.HexToAddress(schemes.AMOY_USDC))

	fmt.Println(dsh.Hex())
}
