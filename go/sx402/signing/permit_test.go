package signing

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/san-lab/sx402/all712"
)

func TestPermit(t *testing.T) {
	permit := new(all712.Permit)
	permit.Domain.ChainID = big.NewInt(84532)
	permit.Domain.Name = "USDC"
	permit.Domain.Version = "2"
	permit.Domain.VerifyingContract = common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e")

	permit.Message.Owner = common.HexToAddress("0xaab05558448c8a9597287db9f61e2d751645b12a")
	permit.Message.Spender = common.HexToAddress("0xCEF702Bd69926B13ab7150624daA7aFEE0300786")
	permit.Message.Deadline = big.NewInt(time.Now().Add(time.Hour * 3).Unix())
	permit.Message.Value = big.NewInt(1100)

	permit.Nonce = big.NewInt(0)

	privkey, err := crypto.HexToECDSA(privkeyhex)
	if err != nil {
		t.Error(err)
		return
	}
	signature, err := SignEIP2612Permit(permit, privkey)
	if err != nil {
		t.Error(err)
		return
	}
	permit.Signature = fmt.Sprintf("0x%x", signature)

	fmt.Printf("owner: %s\nspender: %s\n value: %v\n deadline: %v\n  0x%x\n",
		permit.Message.Owner,
		permit.Message.Spender,
		permit.Message.Value,
		permit.Message.Deadline,
		permit.Signature)

	bt, _ := json.MarshalIndent(permit, " ", " ")
	fmt.Println(string(bt))

	rec, err := VerifyPermitSignature(permit)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Recovered: %s", rec.Hex())
}

func TestPaymentHeader(t *testing.T) {
	ppld := new(all712.PaymentPayload)
	err := json.Unmarshal([]byte(header), ppld)
	if err != nil {
		t.Error()
		return
	}
	permit := new(all712.Permit)
	err = json.Unmarshal(ppld.Payload, permit)
	if err != nil {
		t.Error()
		return
	}
	rec, err := VerifyPermitSignature(permit)
	if err != nil {
		t.Error()
		return
	}
	fmt.Println("Recovered address:", rec)

}

const header = `{"x402Version":1,"scheme":"permit_USDC","network":"base-sepolia","payload":{"domain":{"name":"USDC","version":"2","chainId":84532,"verifyingContract":"0x036CbD53842c5426634e7929541eC2318f3dCF7e"},"message":{"owner":"0xaab05558448c8a9597287db9f61e2d751645b12a","spender":"0xfAc178B1C359D41e9162A1A6385380de96809048","value":2200,"deadline":1751542836},"nonce":12,"signature":"0x395670754ee052f0298716f7f2f585a6d9e30d7a2dc63b9d7653e96b8e9d3d6408e3ff9d5484ba83592904b3e8972d8fef30ea0ca62accf842b769864740e9751b"}}`
