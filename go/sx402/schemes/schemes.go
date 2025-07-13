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

var Exact_EURS_SchemeExtraBytes []byte
var Exact_USDC_SchemeExtraBytes []byte
var Exact_EURC_SchemeExtraBytes []byte
var Payer0_SchemeExtraBytes []byte

var Permit_USDC_SchemeExtraBytes []byte

type Scheme struct {
	SchemeName string
	Network    string
}

const BASE_SEPOLIA_EURS = "0x89D5F29be7753E4c0ad43D08A5067Afc99231CC9" //"0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9"
const AMOY_EURS = "0x73a4F05628fE6976a5d45Fd321b4eD588D8c9Eb3"
const OP_SEPOLIA_EURS = "0x34E2c5d3ac5D07a280f49f0c0B7c69E29BC68F09"
const ARITRUM_SEPOLIA_EURS = "0x8069a68DdaAFE2227f1AF283D23fD6FC2C59b6EC"

const BASE_SEPOLIA_USDC = "0x036CbD53842c5426634e7929541eC2318f3dCF7e"
const AMOY_USDC = "0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582"
const SEPOLIA_EURC = "0x08210F9170F89Ab7658F0B5E3fF39b0E03C594D4"
const SEPOLIA_USDC = "0x93dB8F200E46FD10dbA87E7563148C3cf6985352"
const ZKSYNC_SEPOLIA_USDC = "0xAe045DE5638162fa134807Cb558E15A3F5A7F853"

const Scheme_Exact_USDC = "exact"
const Scheme_Exact_EURS = "exact_EURS"
const Scheme_Exact_EURC = "exact_EURC"
const Scheme_Permit_USDC = "permit_USDC"
const Scheme_Payer0 = "payer0"

// ----------------- SCHEMES -------------
var ExactUsdcOnBaseSepolia = Scheme{Scheme_Exact_USDC, evmbinding.Base_sepolia}
var ExactUsdcOnAmoy = Scheme{Scheme_Exact_USDC, evmbinding.Amoy}
var ExactEurcOnSepolia = Scheme{Scheme_Exact_EURC, evmbinding.Sepolia}
var ExactUsdcOnSepolia = Scheme{Scheme_Exact_USDC, evmbinding.Sepolia}
var ExactUsdcOnZksyncSepolia = Scheme{Scheme_Exact_USDC, evmbinding.ZkSync_sepolia}
var PermitUsdcOnBaseSepolia = Scheme{Scheme_Permit_USDC, evmbinding.Base_sepolia}

var ExactEursOnBaseSepolia = Scheme{Scheme_Exact_EURS, evmbinding.Base_sepolia}
var ExactEursOnOpSepolia = Scheme{Scheme_Exact_EURS, evmbinding.OP_Sepolia}
var ExactEursOnArbitrumSepolia = Scheme{Scheme_Exact_EURS, evmbinding.Arbitrum_sepolia}
var ExactEursOnAmoy = Scheme{Scheme_Exact_EURS, evmbinding.Amoy}

var Payer0OnBaseSepolia = Scheme{Scheme_Payer0, evmbinding.Base_sepolia}

//---------SCHEMES END---------------

var Assets = map[Scheme]string{
	ExactUsdcOnBaseSepolia: BASE_SEPOLIA_USDC,

	ExactEursOnBaseSepolia:     BASE_SEPOLIA_EURS,
	ExactEursOnArbitrumSepolia: ARITRUM_SEPOLIA_EURS,
	ExactEursOnAmoy:            AMOY_EURS,
	ExactEursOnOpSepolia:       OP_SEPOLIA_EURS,

	ExactUsdcOnAmoy:          AMOY_USDC,
	ExactEurcOnSepolia:       SEPOLIA_EURC,
	ExactUsdcOnSepolia:       SEPOLIA_USDC,
	ExactUsdcOnZksyncSepolia: ZKSYNC_SEPOLIA_USDC,
	PermitUsdcOnBaseSepolia:  BASE_SEPOLIA_USDC,
	Payer0OnBaseSepolia:      BASE_SEPOLIA_EURS,
}

// it works because pointers, otherwise init() timing goes wrrrr...
var extras = map[string]*json.RawMessage{ //by Scheme name
	Scheme_Exact_USDC:  (*json.RawMessage)(&Exact_USDC_SchemeExtraBytes),
	Scheme_Exact_EURS:  (*json.RawMessage)(&Exact_EURS_SchemeExtraBytes),
	Scheme_Exact_EURC:  (*json.RawMessage)(&Exact_EURC_SchemeExtraBytes),
	Scheme_Permit_USDC: (*json.RawMessage)(&Permit_USDC_SchemeExtraBytes),
	Scheme_Payer0:      (*json.RawMessage)(&Payer0_SchemeExtraBytes),

	//Payer0O
}

func init() {
	var err error
	Exact_EURS_SchemeExtraBytes, err = json.Marshal(map[string]string{"name": "EURS", "version": "1"})
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

	Payer0_SchemeExtraBytes, err = json.Marshal(map[string]string{"name": "EURS", "version": "1", "dstEid": "40231"})
	if err != nil {
		log.Fatal(err)
	}
}

func GetScheme(name, network string) (*Scheme, error) {
	s := new(Scheme)
	s.Network = network
	s.SchemeName = name
	_, ok := Assets[*s]
	_, ok2 := extras[*&s.SchemeName]
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
		Extra:             extras[s.SchemeName],
		Network:           s.Network,
		Scheme:            s.SchemeName,
	}
}
