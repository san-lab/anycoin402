package schemas

import (
	"fmt"
	"log"
	"testing"

	"github.com/san-lab/sx402/evmbinding"
)

func TestExists(t *testing.T) {
	network := evmbinding.Base_sepolia
	usdc := "exact"
	s, err := GetSchema(usdc, network)
	if err != nil {
		t.Error(err)
	}
	log.Println(assets[*s])

	euros := "EUROS"
	s, err = GetSchema(euros, network)
	if err != nil {
		t.Error(err)
	}
	log.Println(assets[*s])
	r := s.Requirement("this", "42", "x")
	fmt.Println(string(*r.Extra))

}
