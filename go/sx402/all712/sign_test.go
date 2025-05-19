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

const somePayReqs = `{"scheme":"exact","network":"base-sepolia","maxAmountRequired":"15000","resource":"http://localhost:4021/weather","description":"","mimeType":"","payTo":"0x64A8303112D05027F1f1d4ed7e54482799861db0","maxTimeoutSeconds":60,"asset":"0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9","extra":{"name":"EURO_S","version":"2"}}`
