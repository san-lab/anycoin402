package evmbinding

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var BaseSepoliaUSDCAddress = common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e")
var BaseSepoliaEURSAddress = common.HexToAddress("0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9")
var BaseSepoliaEURSMAddress = common.HexToAddress("0x7AE6D004cd29570974357f30c016f8ea3e1A628A")
var boss = common.HexToAddress("0xaab05558448C8a9597287Db9F61e2d751645B12a")

func TestGetBalance(t *testing.T) {
	client, _ := ethclient.Dial(rpcEndpoints["base-sepolia"])
	b, e := CheckTokenBalance(client, BaseSepoliaEURSMAddress, boss)
	if e != nil {
		t.Error(e)
	}
	fmt.Printf("The balance of %s at token %s is %v", boss, BaseSepoliaUSDCAddress, b)
}

func TestParseTrAuth(t *testing.T) {
	parsedABI, err := abi.JSON(strings.NewReader(trWithAuthABI))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(parsedABI)

}

func TestMarkup(t *testing.T) {
	tokenA := "0x7AE6D004cd29570974357f30c016f8ea3e1A628A"
	tokenB := "0xd4A90f96d8626e7Be636F1CFe843005867cA4eAe"

	m, err := GetDetailedMarkup(Base_sepolia, tokenA, 40231, "0xfAc178B1C359D41e9162A1A6385380de96809048")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(m)
	m, err = GetMarkup(Base_sepolia, tokenB, "0xfAc178B1C359D41e9162A1A6385380de96809048")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(m)
}
