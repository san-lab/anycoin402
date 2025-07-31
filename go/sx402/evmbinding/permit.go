package evmbinding

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/san-lab/sx402/all712"
)

const permitABI = `[
  {
    "inputs": [
      { "internalType": "address", "name": "owner", "type": "address" },
      { "internalType": "address", "name": "spender", "type": "address" },
      { "internalType": "uint256", "name": "value", "type": "uint256" },
      { "internalType": "uint256", "name": "deadline", "type": "uint256" },
      { "internalType": "uint8", "name": "v", "type": "uint8" },
      { "internalType": "bytes32", "name": "r", "type": "bytes32" },
      { "internalType": "bytes32", "name": "s", "type": "bytes32" }
    ],
    "name": "permit",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "sender", "type": "address" },
      { "internalType": "address", "name": "recipient", "type": "address" },
      { "internalType": "uint256", "name": "amount", "type": "uint256" }
    ],
    "name": "transferFrom",
    "outputs": [
      { "internalType": "bool", "name": "", "type": "bool" }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]`

func EnactPermit(permit *all712.PermitMessage, facilKey *ecdsa.PrivateKey) (*common.Hash, error) {
	client, err := GetlientByChainID(permit.Domain.ChainID)

	parsedABI, err := abi.JSON(strings.NewReader(permitABI))
	if err != nil {
		return nil, err
	}

	var r, s [32]byte
	var v byte
	// Convert r, s (hex strings to []byte)
	sig, err := hex.DecodeString(strings.TrimPrefix(permit.Signature, "0x"))
	if err != nil || len(r) != 32 {
		return nil, err
	}
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	v = sig[64]
	if v < 27 {
		v += 27
	}

	input, err := parsedABI.Pack("permit", permit.Message.Owner, permit.Message.Spender,
		permit.Message.Value, permit.Message.Deadline, v, r, s)
	if err != nil {
		return nil, err
	}

	// 3. Get nonce for your account
	facilAddress := crypto.PubkeyToAddress(facilKey.PublicKey)

	fromNonce, err := getNonce(context.Background(), client, facilAddress)
	if err != nil {
		return nil, err
	}

	// 4. Suggest gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	// 5. Estimate gas limit
	msg := ethereum.CallMsg{
		From:     facilAddress,
		To:       &permit.Domain.VerifyingContract,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     input,
	}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, err
	}

	// 6. Create the transaction
	tx := types.NewTransaction(fromNonce, permit.Domain.VerifyingContract, big.NewInt(0), gasLimit+10000, gasPrice, input)

	// 7. Sign the transaction
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), facilKey)
	if err != nil {
		return nil, err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	// Return transaction hash
	h := signedTx.Hash()
	return &h, nil

}

func TransferFrom(from, to, asset common.Address, amount, chainID *big.Int, facilKey *ecdsa.PrivateKey) (*common.Hash, error) {
	client, err := GetlientByChainID(chainID)

	parsedABI, err := abi.JSON(strings.NewReader(permitABI))
	if err != nil {
		return nil, err
	}

	// 3. Get nonce for your account
	facilAddress := crypto.PubkeyToAddress(facilKey.PublicKey)

	input, err := parsedABI.Pack("transferFrom", from, to, amount)
	if err != nil {
		return nil, err
	}

	fromNonce, err := getNonce(context.Background(), client, facilAddress)
	if err != nil {
		return nil, err
	}

	// 4. Suggest gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	// 5. Estimate gas limit
	/* Racy, racy
	msg := ethereum.CallMsg{
		From:     from,
		To:       &asset,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     input,
	}

	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	*/
	gasLimit := uint64(100000)

	// 6. Create the transaction
	tx := types.NewTransaction(fromNonce, asset, big.NewInt(0), gasLimit, gasPrice, input)

	// 7. Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), facilKey)
	if err != nil {
		return nil, err
	}

	// 8. Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	// Return transaction hash
	h := signedTx.Hash()
	return &h, nil

}
