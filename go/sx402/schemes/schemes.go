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

// Scheme Types - an ugly artifact of some x402 design choices
const ExactType = "exac"
const PermitType = "permit"
const Payer0Legacy = "payer0Legacy"
const Payer0Type = "payer0"

var Exact_EURS_SchemeExtraBytes []byte
var Exact_USDC_SchemeExtraBytes []byte
var Exact_EURC_SchemeExtraBytes []byte
var Payer0_SchemeExtraBytes []byte

var Permit_USDC_SchemeExtraBytes []byte

type Scheme struct {
	SchemeName string
	Type       string
	Network    string
	Asset      string
	Extra      *ExtraInfo
}

// ---ASSETS-------
const ARBITRUM_SEPOLIA_EURSM = "0xd7A4537267741d00F9654856b81F0AEe409B7aD9"
const AMOY_EURSM = "0xb1e5EC11F3F453BB5a1E2E4D0C1ea8ECE059401d"
const BASE_SEPOLIA_EURSM = "0x0190C8a558ad75d7929bE7d06b07D4cdCdAC18c4"
const OP_SEPOLIA_EURSM = "0xE6280F5EED02505D7e5cE762E35C878A99b54ac2"

const OP_SEPOLIA_DRAFT = "0xE6280F5EED02505D7e5cE762E35C878A99b54ac2"

const BASE_SEPOLIA_EURS = "0x89D5F29be7753E4c0ad43D08A5067Afc99231CC9" //"0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9"
const AMOY_EURS = "0x73a4F05628fE6976a5d45Fd321b4eD588D8c9Eb3"
const OP_SEPOLIA_EURS = "0x34E2c5d3ac5D07a280f49f0c0B7c69E29BC68F09"
const ARBITRUM_SEPOLIA_EURS = "0x8069a68DdaAFE2227f1AF283D23fD6FC2C59b6EC" //"0x8069a68DdaAFE2227f1AF283D23fD6FC2C59b6EC"

const BASE_SEPOLIA_USDC = "0x036CbD53842c5426634e7929541eC2318f3dCF7e"
const AMOY_USDC = "0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582"
const SEPOLIA_EURC = "0x08210F9170F89Ab7658F0B5E3fF39b0E03C594D4"
const SEPOLIA_USDC = "0x93dB8F200E46FD10dbA87E7563148C3cf6985352"
const ZKSYNC_SEPOLIA_USDC = "0xAe045DE5638162fa134807Cb558E15A3F5A7F853"

// ---SCHEME NAMES-------
const Scheme_Exact_USDC = "exact"
const Scheme_Exact_EURS = "exact_EURS"
const Scheme_Exact_EURC = "exact_EURC"
const Scheme_Permit_USDC = "permit_USDC"
const Scheme_Payer0_toArbitrum = "payer0_toArbitrum"
const Scheme_Payer0_toBase = "payer0_toBase"

const Scheme_Payer0M_toBase = "payer0_toBase_withMarkup"

const Scheme_Payer0Plus_toBase = "PZ_toBase"
const Scheme_Payer0Plus_toArbitrum = "PZ_toArbitrum"
const Scheme_Payer0Plus_toOP = "PZ_toOP"
const Scheme_Payer0Plus_toAmoy = "PZ_toAmoy"

const Scheme_Exact_Draft = "exact_EURM_draft"

// ---EXTRA INFO MAPS----
type ExtraInfo map[string]string

func NewExtraInfo(name, version string) *ExtraInfo {
	nei := ExtraInfo{}
	nei["name"] = name
	nei["version"] = version
	return &nei
}

// Setting new parameter on a clone of the origin
func (ei *ExtraInfo) Set(name, value string) *ExtraInfo {
	nei := ExtraInfo{}
	for k, v := range *ei {
		nei[k] = v
	}
	nei[name] = value
	return &nei
}

// Sugar
func (ei *ExtraInfo) SetDstEid(dstEid string) *ExtraInfo {
	return ei.Set("dstEid", dstEid)
}

var ExtraUSDC = NewExtraInfo("USDC", "2")
var ExtraEURC = NewExtraInfo("EURC", "2")
var ExtraEURS = NewExtraInfo("EURS", "1")
var ExtraEURSM = NewExtraInfo("EURSM", "1")
var ExtraPermitUSDC = ExtraUSDC.Set("facilitator", "0xfAc178B1C359D41e9162A1A6385380de96809048")

// ----------------- SCHEMES -------------
type SchemeKey struct {
	Name    string
	Network string
}

var SchemeMap = map[SchemeKey]Scheme{}

func NewScheme(name, sType, network, asset string, extra *ExtraInfo) Scheme {
	key := SchemeKey{name, network}
	s := Scheme{name, sType, network, asset, extra}
	SchemeMap[key] = s
	return s
}

var ExactUsdcOnBaseSepolia = NewScheme(Scheme_Exact_USDC, ExactType, evmbinding.Base_sepolia, BASE_SEPOLIA_USDC, ExtraUSDC)
var ExactUsdcOnAmoy = NewScheme(Scheme_Exact_USDC, ExactType, evmbinding.Amoy, AMOY_USDC, ExtraUSDC)
var ExactEurcOnSepolia = NewScheme(Scheme_Exact_EURC, ExactType, evmbinding.Sepolia, SEPOLIA_EURC, ExtraEURC)
var ExactUsdcOnSepolia = NewScheme(Scheme_Exact_USDC, ExactType, evmbinding.Sepolia, SEPOLIA_USDC, ExtraUSDC)
var ExactUsdcOnZksyncSepolia = NewScheme(Scheme_Exact_USDC, ExactType, evmbinding.ZkSync_sepolia, ZKSYNC_SEPOLIA_USDC, ExtraUSDC)
var PermitUsdcOnBaseSepolia = NewScheme(Scheme_Permit_USDC, PermitType, evmbinding.Base_sepolia, BASE_SEPOLIA_USDC, ExtraPermitUSDC)

var ExactEursOnBaseSepolia = NewScheme(Scheme_Exact_EURS, ExactType, evmbinding.Base_sepolia, BASE_SEPOLIA_EURS, ExtraEURS)
var ExactEursOnOpSepolia = NewScheme(Scheme_Exact_EURS, ExactType, evmbinding.OP_Sepolia, OP_SEPOLIA_EURS, ExtraEURS)
var ExactEursOnArbitrumSepolia = NewScheme(Scheme_Exact_EURS, ExactType, evmbinding.Arbitrum_sepolia, ARBITRUM_SEPOLIA_EURS, ExtraEURS)
var ExactEursOnAmoy = NewScheme(Scheme_Exact_EURS, ExactType, evmbinding.Amoy, AMOY_EURS, ExtraEURS)

var Payer0EURSBaseToArbitrum = NewScheme(Scheme_Payer0_toArbitrum, Payer0Legacy, evmbinding.Base_sepolia, BASE_SEPOLIA_EURS, ExtraEURS.SetDstEid("40231"))
var Payer0EURSArbitrumToBase = NewScheme(Scheme_Payer0_toBase, Payer0Legacy, evmbinding.Arbitrum_sepolia, ARBITRUM_SEPOLIA_EURS, ExtraEURS.SetDstEid("40245"))

var Payer0MarkupArbitrumToBase = NewScheme(Scheme_Payer0M_toBase, Payer0Legacy, evmbinding.Arbitrum_sepolia, ARBITRUM_SEPOLIA_EURSM, ExtraEURSM.SetDstEid("40245").Set("maxMarkup", "42"))

var P0_Arbitrum_toBase = NewScheme(Scheme_Payer0Plus_toBase, Payer0Type, evmbinding.Arbitrum_sepolia, ARBITRUM_SEPOLIA_EURSM, ExtraEURSM.SetDstEid("40245"))
var P0_Base_toArbitrum = NewScheme(Scheme_Payer0Plus_toArbitrum, Payer0Type, evmbinding.Base_sepolia, BASE_SEPOLIA_EURSM,
	ExtraEURSM.SetDstEid("40231"))
var P0_Amoy_toArbitrum = NewScheme(Scheme_Payer0Plus_toArbitrum, Payer0Type, evmbinding.Amoy, AMOY_EURSM,
	ExtraEURSM.SetDstEid("40231"))
var P0_OP_toBase = NewScheme(Scheme_Payer0Plus_toBase, Payer0Type, evmbinding.OP_Sepolia, OP_SEPOLIA_EURSM,
	ExtraEURSM.SetDstEid("40245"))

var ExactEURMOnOp = NewScheme(Scheme_Exact_Draft, ExactType, evmbinding.OP_Sepolia, OP_SEPOLIA_DRAFT, ExtraEURSM)

//---------SCHEMES END---------------

func GetScheme(name, network string) (*Scheme, error) {
	s, ok := SchemeMap[SchemeKey{name, network}]
	if ok {
		return &s, nil
	}
	return nil, fmt.Errorf("unsupported scheme: %s, %s", s.Network, s.SchemeName)
}

func (s *Scheme) Requirement(resource, price, payto string) *types.PaymentRequirements {
	var extra json.RawMessage
	extra, err := json.Marshal(s.Extra)
	if err != nil {
		log.Println("Error marshalling ExtraInfo. This cannot happen.", err)
	}
	return &types.PaymentRequirements{
		PayTo:             payto,
		MaxTimeoutSeconds: 120,
		Asset:             s.Asset,
		MaxAmountRequired: price,
		Resource:          resource,
		Extra:             &extra,
		Network:           s.Network,
		Scheme:            s.SchemeName,
	}
}
