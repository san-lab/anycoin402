package schemes

import (
	"fmt"
	"log"
	"testing"

	"github.com/san-lab/sx402/evmbinding"
)

func TestExists(t *testing.T) {
	network := evmbinding.Base_sepolia
	usdc := "exact"
	s, err := GetScheme(usdc, network)
	if err != nil {
		t.Error(err)
	}
	log.Println(Assets[*s])

	euros := "EUROS"
	s, err = GetScheme(euros, network)
	if err != nil {
		t.Error(err)
	}
	log.Println(Assets[*s])
	r := s.Requirement("this", "42", "x")
	fmt.Println(string(*r.Extra))

}
