package all712

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestPermit(t *testing.T) {
	permit := new(Permit)
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

	permit.Signature, err = SignEIP2612Permit(permit, privkey)
	if err != nil {
		t.Error(err)
		return
	}
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
