package store

import (
	"log"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/san-lab/sx402/schemes"
)

type Accepts []*types.PaymentRequirements

func (ac *Accepts) addRequirement(scheme_name, network, resourceURI, price string) bool {
	scheme, err := schemes.GetScheme(scheme_name, network)
	if err != nil {
		log.Println(err)
		return false
	} else {

		*ac = append(*ac, scheme.Requirement(resourceURI, price, store_wallet))
	}
	return true
}

func (ac *Accepts) addSchemeInstance(scheme schemes.Scheme, resourceURI, price string) {
	*ac = append(*ac, scheme.Requirement(resourceURI, price, store_wallet))
}
