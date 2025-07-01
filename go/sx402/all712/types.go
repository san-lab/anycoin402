package all712

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type EVMPayload types.ExactEvmPayload

/*
type Payload struct {
	Signature     string         `json:"signature"`
	Authorization *Authorization `json:"authorization"`
}

type Authorization struct {
	From        string   `json:"from"`
	To          string   `json:"to"`
	Value       string   `json:"value"`
	ValidAfter  *big.Int `json:"validAfter"`
	ValidBefore *big.Int `json:"validBefore"`
	Nonce       string   `json:"nonce"`
}
*/

func mustNewType(t string) abi.Type {
	typ, err := abi.NewType(t, "", nil)
	if err != nil {
		log.Fatalf("failed to create ABI type %s: %v", t, err)
	}
	return typ
}

type Envelope struct {
	X402Version         int                        `json:"x402Version"`
	PaymentPayload      *PaymentPayload            `json:"paymentPayload"`
	PaymentRequirements *types.PaymentRequirements `json:"paymentRequirements"`
}

type PaymentPayload struct {
	X402Version int             `json:"x402Version"`
	Scheme      string          `json:"scheme"`
	Network     string          `json:"network"`
	Payload     json.RawMessage `json:"payload"`
}

type Permit struct {
	Domain    Domain        `json:"domain"`
	Message   PermitMessage `json:"message"`
	Nonce     *big.Int      `json:"nonce"`
	Signature []byte        `json:"signature,omitempty"`
}

type PermitMessage struct {
	Owner    common.Address `json:"owner"`
	Spender  common.Address `json:"spender"`
	Value    *big.Int       `json:"value"`
	Deadline *big.Int       `json:"deadline"`
}

type Domain struct {
	Name              string         `json:"name,omitempty"`
	Version           string         `json:"version,omitempty"`
	ChainID           *big.Int       `json:"chainId,omitempty"`
	VerifyingContract common.Address `json:"verifyingContract,omitempty"`
}

func (permit *Permit) Digest() ([]byte, error) {
	return EIP712PermitHash(
		permit.Message.Owner,
		permit.Message.Spender,
		permit.Domain.VerifyingContract,
		permit.Message.Value,
		permit.Message.Deadline,
		permit.Domain.ChainID,
		permit.Nonce,
		permit.Domain.Name,
		permit.Domain.Version,
	)
}
