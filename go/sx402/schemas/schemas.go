package schemas

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/san-lab/sx402/evmbinding"
)

/* exact schema
struct {
	Scheme            string            `json:"scheme"`
	Network           string            `json:"network"`
	MaxAmountRequired string            `json:"maxAmountRequired"`
	Resource          string            `json:"resource"`
	Description       string            `json:"description"`
	MimeType          string            `json:"mimeType"`
	PayTo             string            `json:"payTo"`
	MaxTimeoutSeconds int               `json:"maxTimeoutSeconds"`
	Asset             string            `json:"asset"`
	Extra             map[string]string `json:"extra"`
}
*/

var EURO_SSchemaExtraBytes []byte
var USDC_SchemaExtraBytes []byte
var EURC_SchemaExtraBytes []byte

type Scheme struct {
	SchemeName string
	Network    string
}

const BASE_SEPOLIA_EUROS = "0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9"
const BASE_SEPOLIA_USDC = "0x036CbD53842c5426634e7929541eC2318f3dCF7e"
const AMOY_USDC = "0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582"
const SEPOLIA_EURC = "0x08210F9170F89Ab7658F0B5E3fF39b0E03C594D4"
const SEPOLIA_USDC = "0x93dB8F200E46FD10dbA87E7563148C3cf6985352"

var UsdcOnBaseSepolia = Scheme{"exact", evmbinding.Base_sepolia}
var EurosOnBaseSepolia = Scheme{"EUROS", evmbinding.Base_sepolia}
var UsdcOnAmoy = Scheme{"exact", evmbinding.Amoy}
var EurcOnSepolia = Scheme{"EURC", evmbinding.Sepolia}
var UsdcOnSepolia = Scheme{"USDC", evmbinding.Sepolia}

var Assets = map[Scheme]string{
	UsdcOnBaseSepolia:  BASE_SEPOLIA_USDC,
	EurosOnBaseSepolia: BASE_SEPOLIA_EUROS,
	UsdcOnAmoy:         AMOY_USDC,
	EurcOnSepolia:      SEPOLIA_EURC,
	UsdcOnSepolia:      SEPOLIA_USDC,
}

var extras = map[Scheme]*json.RawMessage{
	UsdcOnBaseSepolia:  (*json.RawMessage)(&USDC_SchemaExtraBytes),
	EurosOnBaseSepolia: (*json.RawMessage)(&EURO_SSchemaExtraBytes),
	UsdcOnAmoy:         (*json.RawMessage)(&USDC_SchemaExtraBytes),
	UsdcOnSepolia:      (*json.RawMessage)(&USDC_SchemaExtraBytes),
	EurcOnSepolia:      (*json.RawMessage)(&EURC_SchemaExtraBytes),
}

func init() {
	var err error
	EURO_SSchemaExtraBytes, err = json.Marshal(map[string]string{"name": "EURO_S", "version": "2"})
	if err != nil {
		log.Fatal(err)
	}
	USDC_SchemaExtraBytes, err = json.Marshal(map[string]string{"name": "USDC", "version": "2"})
	if err != nil {
		log.Fatal(err)
	}

	EURC_SchemaExtraBytes, err = json.Marshal(map[string]string{"name": "EURC", "version": "2"})
	if err != nil {
		log.Fatal(err)
	}

}

func GetSchema(name, network string) (*Scheme, error) {
	s := new(Scheme)
	s.Network = network
	s.SchemeName = name
	_, ok := Assets[*s]
	_, ok2 := extras[*s]
	if ok && ok2 {
		return s, nil
	}
	return nil, fmt.Errorf("unsupported schema: %s, %s, %v, %v", s.Network, s.SchemeName, ok, ok2)
}

func (s *Scheme) Requirement(resource, price, payto string) *types.PaymentRequirements {
	return &types.PaymentRequirements{
		PayTo:             payto,
		MaxTimeoutSeconds: 120,
		Asset:             Assets[*s],
		MaxAmountRequired: price,
		Resource:          resource,
		Extra:             extras[*s],
		Network:           s.Network,
		Scheme:            s.SchemeName,
	}
}
