package schemes

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/san-lab/sx402/evmbinding"
)

/* exact scheme
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

var Exact_EURO_S_SchemeExtraBytes []byte
var Exact_USDC_SchemeExtraBytes []byte
var Exact_EURC_SchemeExtraBytes []byte

var Permit_USDC_SchemeExtraBytes []byte

type Scheme struct {
	SchemeName string
	Network    string
}

const BASE_SEPOLIA_EUROS = "0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9"
const BASE_SEPOLIA_USDC = "0x036CbD53842c5426634e7929541eC2318f3dCF7e"
const AMOY_USDC = "0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582"
const SEPOLIA_EURC = "0x08210F9170F89Ab7658F0B5E3fF39b0E03C594D4"
const SEPOLIA_USDC = "0x93dB8F200E46FD10dbA87E7563148C3cf6985352"
const ZKSYNC_SEPOLIA_USDC = "0xAe045DE5638162fa134807Cb558E15A3F5A7F853"

const Scheme_Exact_USDC = "exact"
const Scheme_Exact_EUROS = "exact_EUROS"
const Scheme_Exact_EURC = "exact_EURC"
const Scheme_Permit_USDC = "permit_USDC"

var ExactUsdcOnBaseSepolia = Scheme{Scheme_Exact_USDC, evmbinding.Base_sepolia}
var ExactEurosOnBaseSepolia = Scheme{Scheme_Exact_EUROS, evmbinding.Base_sepolia}
var ExactUsdcOnAmoy = Scheme{Scheme_Exact_USDC, evmbinding.Amoy}
var ExactEurcOnSepolia = Scheme{Scheme_Exact_EURC, evmbinding.Sepolia}
var ExactUsdcOnSepolia = Scheme{Scheme_Exact_USDC, evmbinding.Sepolia}
var ExactUsdcOnZksyncSepolia = Scheme{Scheme_Exact_USDC, evmbinding.ZkSync_sepolia}
var PermitUsdcOnBaseSepolia = Scheme{Scheme_Permit_USDC, evmbinding.Base_sepolia}

var Assets = map[Scheme]string{
	ExactUsdcOnBaseSepolia:   BASE_SEPOLIA_USDC,
	ExactEurosOnBaseSepolia:  BASE_SEPOLIA_EUROS,
	ExactUsdcOnAmoy:          AMOY_USDC,
	ExactEurcOnSepolia:       SEPOLIA_EURC,
	ExactUsdcOnSepolia:       SEPOLIA_USDC,
	ExactUsdcOnZksyncSepolia: ZKSYNC_SEPOLIA_USDC,
	PermitUsdcOnBaseSepolia:  BASE_SEPOLIA_USDC,
}

// it works because pointers, otherwise init() timing goes wrrrr...
var extras = map[Scheme]*json.RawMessage{
	ExactUsdcOnBaseSepolia:   (*json.RawMessage)(&Exact_USDC_SchemeExtraBytes),
	ExactEurosOnBaseSepolia:  (*json.RawMessage)(&Exact_EURO_S_SchemeExtraBytes),
	ExactUsdcOnAmoy:          (*json.RawMessage)(&Exact_USDC_SchemeExtraBytes),
	ExactUsdcOnSepolia:       (*json.RawMessage)(&Exact_USDC_SchemeExtraBytes),
	ExactEurcOnSepolia:       (*json.RawMessage)(&Exact_EURC_SchemeExtraBytes),
	ExactUsdcOnZksyncSepolia: (*json.RawMessage)(&Exact_USDC_SchemeExtraBytes),
	PermitUsdcOnBaseSepolia:  (*json.RawMessage)(&Permit_USDC_SchemeExtraBytes),
}

func init() {
	var err error
	Exact_EURO_S_SchemeExtraBytes, err = json.Marshal(map[string]string{"name": "EURO_S", "version": "2"})
	if err != nil {
		log.Fatal(err)
	}
	Exact_USDC_SchemeExtraBytes, err = json.Marshal(map[string]string{"name": "USDC", "version": "2"})
	if err != nil {
		log.Fatal(err)
	}

	Exact_EURC_SchemeExtraBytes, err = json.Marshal(map[string]string{"name": "EURC", "version": "2"})
	if err != nil {
		log.Fatal(err)
	}
	Permit_USDC_SchemeExtraBytes, err = json.Marshal(map[string]string{"name": "USDC", "version": "2", "facilitator": "0xfAc178B1C359D41e9162A1A6385380de96809048"})
	if err != nil {
		log.Fatal(err)
	}
}

func GetScheme(name, network string) (*Scheme, error) {
	s := new(Scheme)
	s.Network = network
	s.SchemeName = name
	_, ok := Assets[*s]
	_, ok2 := extras[*s]
	if ok && ok2 {
		return s, nil
	}
	return nil, fmt.Errorf("unsupported scheme: %s, %s, %v, %v", s.Network, s.SchemeName, ok, ok2)
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
