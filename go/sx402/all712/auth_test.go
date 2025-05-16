package all712

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

const test_auth1 = `{
	"signature": "0x2d6a7588d6acca505cbf0d9a4a227e0c52c6c34008c8e8986a1283259764173608a2ce6496642e377d6da8dbbf5836e9bd15092f9ecab05ded3d6293af148b571c",
	"authorization": {
	  "from": "0x857b06519E91e3A54538791bDbb0E22373e36b66",
	  "to": "0x209693Bc6afc0C5328bA36FaF03C514EF312287C",
	  "value": "10000",
	  "validAfter": "1740672089",
	  "validBefore": "1740672154",
	  "nonce": "0xf3746613c2d920b5fdabc0856f2aeb2d4f88ee6037b8cc5d04a71a4462f13480"
	}
  }`

const test_auth2 = `{"signature":"0x9e4259bd90c4776404f3b6261d91e23af523dfcf53edb354c68f13b156e820de159c8a4e8e1b19b0b9748520f8e761030c58806614b48a2fd30325bb15df03651c",
"authorization":{"from":"0xaab05558448C8a9597287Db9F61e2d751645B12a","to":"0xCEF702Bd69926B13ab7150624daA7aFEE0300786",
"value":"1000","validAfter":"1747126904","validBefore":"1747127024",
"nonce":"0x4c8fd27a64f1cf5c955ebd5813d772cc56f03f45e0c8196ec180cea98449bb3b"
}
}`

var sepoliaChainId = big.NewInt(84532)

var baseChainId = big.NewInt(8453)

const USDCName = "USDC"
const USDCOnBaseSepolia = "0x036CbD53842c5426634e7929541eC2318f3dCF7e"

func TestAuth(t *testing.T) {
	ea := EVMPayload{}
	err := json.Unmarshal([]byte(test_auth2), &ea)
	if err != nil {
		t.Error(err)
	}

	ok, addr, err := VerifyTransferWithAuthorizationSignature(ea.Signature, *ea.Authorization, USDCName, "2", sepoliaChainId, common.HexToAddress(USDCOnBaseSepolia))
	log.Println(err)
	log.Println(addr)
	log.Println(ok)
}

func TestDecodePayload(t *testing.T) {
	bt, err := base64.StdEncoding.DecodeString(base64Payment2)
	log.Println(err)
	log.Println(string(bt))
}

const paymentRequiredMsg = `{"x402Version":1,"error":"X-PAYMENT header is required","accepts":[{"scheme":"exact","network":"base-sepolia","maxAmountRequired":"1000","resource":"http://localhost:4021/weather","description":"","mimeType":"","payTo":"0xCEF702Bd69926B13ab7150624daA7aFEE0300786","maxTimeoutSeconds":60,"asset":"0x036CbD53842c5426634e7929541eC2318f3dCF7e","extra":{"name":"USDC","version":"2"}}]}`

const base64Payment = `eyJ4NDAyVmVyc2lvbiI6MSwic2NoZW1lIjoiZXhhY3QiLCJuZXR3b3JrIjoiYmFzZS1zZXBvbGlhIiwicGF5bG9hZCI6eyJzaWduYXR1cmUiOiIweDllNDI1OWJkOTBjNDc3NjQwNGYzYjYyNjFkOTFlMjNhZjUyM2RmY2Y1M2VkYjM1NGM2OGYxM2IxNTZlODIwZGUxNTljOGE0ZThlMWIxOWIwYjk3NDg1MjBmOGU3NjEwMzBjNTg4MDY2MTRiNDhhMmZkMzAzMjViYjE1ZGYwMzY1MWMiLCJhdXRob3JpemF0aW9uIjp7ImZyb20iOiIweGFhYjA1NTU4NDQ4QzhhOTU5NzI4N0RiOUY2MWUyZDc1MTY0NUIxMmEiLCJ0byI6IjB4Q0VGNzAyQmQ2OTkyNkIxM2FiNzE1MDYyNGRhQTdhRkVFMDMwMDc4NiIsInZhbHVlIjoiMTAwMCIsInZhbGlkQWZ0ZXIiOiIxNzQ3MTI2OTA0IiwidmFsaWRCZWZvcmUiOiIxNzQ3MTI3MDI0Iiwibm9uY2UiOiIweDRjOGZkMjdhNjRmMWNmNWM5NTVlYmQ1ODEzZDc3MmNjNTZmMDNmNDVlMGM4MTk2ZWMxODBjZWE5ODQ0OWJiM2IifX19`
const base64Payment2 = `eyJ4NDAyVmVyc2lvbiI6MSwic2NoZW1lIjoiZXhhY3QiLCJuZXR3b3JrIjoiYmFzZS1zZXBvbGlhIiwicGF5bG9hZCI6eyJzaWduYXR1cmUiOiIweDg4NTc0NzYwYjFkNGMwYWUwNmI5MzUzZmQ4MjdmOGFkYWQ5OTU0NDEzODQ4NDhmOTMxOTJiMmYyNWJlYWY1MTc2YmYyZWY1YzA4NTI4ZmEwYzlhZTc1NDNlMTU2ODRmMzBkNWQ2NThkMzQ4Zjg0N2FhYzMwMDY2ZGE4NjI4OWI0MWIiLCJhdXRob3JpemF0aW9uIjp7ImZyb20iOiIweGFhYjA1NTU4NDQ4QzhhOTU5NzI4N0RiOUY2MWUyZDc1MTY0NUIxMmEiLCJ0byI6IjB4Q0VGNzAyQmQ2OTkyNkIxM2FiNzE1MDYyNGRhQTdhRkVFMDMwMDc4NiIsInZhbHVlIjoiMTAwMCIsInZhbGlkQWZ0ZXIiOiIxNzQ3MTQyMTA3IiwidmFsaWRCZWZvcmUiOiIxNzQ3MTQyMjI3Iiwibm9uY2UiOiIweGM3NWYyYWUzYTUzMzAyYTczNTI4YWMxZTIxZWUwM2NlNGM2NTYyNmZjZjAzMjVjOTIwMDk4YTdkMGMyYzllOGEifX19`
const base64PaymentResponse = `eyJzdWNjZXNzIjp0cnVlLCJ0cmFuc2FjdGlvbiI6IjB4ZmJjMjRhMGE3NmQ4NzExMjM0ODMzNGQ1YzZjNzMyYjIxMzk3OGRmNWI3ZmQ2NGVkYWY3NTU2YTllOWIwYTI5NyIsIm5ldHdvcmsiOiJiYXNlLXNlcG9saWEiLCJwYXllciI6IjB4YWFiMDU1NTg0NDhDOGE5NTk3Mjg3RGI5RjYxZTJkNzUxNjQ1QjEyYSJ9`
