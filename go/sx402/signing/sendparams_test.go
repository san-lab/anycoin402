package signing

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"

	x402types "github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/oft"
)

func TestQuoteSendParams(t *testing.T) {

	hexDat, _ := hex.DecodeString(sendFromData[2:])
	//sendParam := new(
	mabi, err := abi.JSON(strings.NewReader(ABI))
	//err = rlp.DecodeBytes(tx.Data(), &v)
	if err != nil {
		t.Error(err)
	}
	send := mabi.Methods["sendFrom"]

	// 5. Decode
	args, err := send.Inputs.Unpack(hexDat[4:])
	if err != nil {
		log.Fatalf("failed to decode args: %v", err)
	}

	fmt.Println(len(args))

	sendParams := new(oft.SendParam)
	CopyFields(sendParams, args[0])

	mfee := new(oft.MessagingFee)
	CopyFields(mfee, args[1])
	bt, _ := json.MarshalIndent(sendParams, " ", " ")
	fmt.Println("extra:", sendParams.ExtraOptions)
	fmt.Println(string(bt))
	fmt.Printf(`[%v, 0x%x, %v, %v, "0x%x", "0x", "0x"]`, sendParams.DstEid, sendParams.To, sendParams.AmountLD, sendParams.MinAmountLD, sendParams.ExtraOptions)
	fmt.Println()
}

const sendFromData = `0xa6e329f500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000916200a4b8e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000aab05558448c8a9597287db9f61e2d751645b12a000000000000000000000000aab05558448c8a9597287db9f61e2d751645b12a0000000000000000000000000000000000000000000000000000000000009d4b000000000000000000000000aab05558448c8a9597287db9f61e2d751645b12a0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000002000300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000`

var fromAddress = common.HexToAddress("0xAA5A032fA7a69f884072C6BA59FE22A90dFbc780")

func TestQuoteSend(t *testing.T) {
	privKey, _ := crypto.HexToECDSA(privkeyhex)
	contractAddress := common.HexToAddress("0x89D5F29be7753E4c0ad43D08A5067Afc99231CC9")
	toAddress := common.HexToAddress("0xCEF702Bd69926B13ab7150624daA7aFEE0300786")
	benefAddress := common.HexToAddress("0xaab05558448C8a9597287Db9F61e2d751645B12a")

	// Connect to an Ethereum node (e.g., Infura, local node)
	client, err := evmbinding.GetlientByChainID(big.NewInt(84532))
	if err != nil {

		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}

	defer client.Close()

	// Define your contract's deployed address.

	// Instantiate the contract binding.
	contract, err := oft.NewOft(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate the contract: %v", err)
	}

	// Create a new CallOpts (read-only call)
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	to := [32]byte{}
	copy(to[12:], toAddress[:])
	fmt.Printf("0x%x\n", to)

	// Prepare the SendParam structure.
	// You must set these fields with appropriate values.
	sendParam := oft.SendParam{
		DstEid:       40231,              // example destination chain ID
		To:           to,                 // populate with a properly formatted bytes32 value (e.g., via common.LeftPadBytes(addr.Bytes(), 32))
		AmountLD:     big.NewInt(100000), // e.g., 1 * 10^18 for 18 decimals
		MinAmountLD:  big.NewInt(100000), // minimum acceptable amount received on dest chain
		ExtraOptions: []byte{0, 3},       // if you have extra options, encode them here
		ComposeMsg:   []byte{},           // provide if necessary
		OftCmd:       []byte{},           // provide if necessary
	}

	// Set the _payInLzToken parameter (true or false as required).
	payInLzToken := false

	// Call quoteSend on the contract.
	messagingFee, err := contract.QuoteSend(callOpts, sendParam, payInLzToken)
	if err != nil {
		log.Fatalf("QuoteSend call failed: %v", err)
	}

	// The return value is expected to be a struct, for example:
	// type MessagingFee struct {
	//     NativeFee  *big.Int
	//     LzTokenFee *big.Int
	// }
	fmt.Printf("Native Fee: %s\n", messagingFee.NativeFee.String())
	fmt.Printf("LzToken Fee: %s\n", messagingFee.LzTokenFee.String())

	messagingFee.NativeFee.Add(messagingFee.NativeFee, big.NewInt(10096650793526))
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(84532))
	if err != nil {
		log.Fatalf("Failed to create TransactOpts: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	// Optionally customize gas
	auth.GasLimit = uint64(1000000)                           // set if needed
	auth.GasPrice = gasPrice.Add(gasPrice, big.NewInt(30000)) //
	auth.Value = messagingFee.NativeFee                       //.Mul(messagingFee.NativeFee, big.NewInt(2))
	/*
		//txh, err := contract.Send(auth, sendParam, messagingFee, benefAddress)
		txh, err := contract.SendFrom(auth, sendParam, messagingFee, fromAddress, benefAddress)

		if err != nil {
			t.Error(err)

		}
		fmt.Printf("transaction hash: %s", txh.Hash().Hex())
	*/
	//Prepare authorization
	nonce := crypto.Keccak256Hash([]byte(time.ANSIC))

	validAfter := big.NewInt(time.Now().Unix() - 300)
	validBefore := big.NewInt(time.Now().Add(time.Hour).Unix())

	authorization := all712.EVMPayload{}
	authorization.Authorization = new(x402types.ExactEvmPayloadAuthorization)
	sauth := authorization.Authorization
	sauth.ValidAfter = validAfter.String()
	sauth.ValidBefore = validBefore.String()
	sauth.From = fromAddress.Hex()
	sauth.To = toAddress.Hex()
	sauth.Nonce = fmt.Sprintf("0x%x", nonce)
	sauth.Value = sendParam.AmountLD.String()

	aa5priv, err := crypto.HexToECDSA(aa5privhx)
	if err != nil {
		t.Error()
		return
	}

	sig, err := SignERC3009Authorization(authorization.Authorization, aa5priv, big.NewInt(84532), "EURS", "1", contractAddress)
	if err != nil {
		t.Error(err)
		return
	}

	//VerifyTransferWithAuthorizationSignature(hex.EncodeToString(sig), *sauth, "EURS", "1", big.NewInt(84532), contractAddress)
	//authorization.Signature = fmt.Sprintf("0x%x", sig)
	txh, err := contract.SendWithAuthorization(auth, sendParam, messagingFee, fromAddress, validAfter, validBefore, nonce, sig, benefAddress)
	if err != nil {
		t.Error(err)

	}
	fmt.Printf("transaction hash: %s", txh.Hash().Hex())

}

var rawTxHex = `0x02f9025683066eee220184078b30c483037d9f9479ef5e7e18a5719d0b8af4e7ea0a54818441aa06860bb2affbd7b1b901e4c7c7f5b3000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000bb2affbd7b10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000aab05558448c8a9597287db9f61e2d751645b12a0000000000000000000000000000000000000000000000000000000000009d35000000000000000000000000cef702bd69926b13ab7150624daa7afee03007860000000000000000000000000000000000000000000000000186cc6acd4b00000000000000000000000000000000000000000000000000000186cc6acd4b000000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000002000300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c080a00fa5af8963594c86910df4576db3b69dcd74a2c504bb99a2e787fc4de1f0c59fa052ab1f3e6551077cc6f2472467e820a273d73d335621cef0e1e631d522847d89`

func TestSendParams(t *testing.T) {
	if len(rawTxHex) >= 2 && rawTxHex[:2] == "0x" {
		rawTxHex = rawTxHex[2:]
	}

	rawTxBytes, err := hex.DecodeString(rawTxHex)
	if err != nil {
		log.Fatalf("Failed to decode hex: %v", err)
	}

	// Check type byte (EIP-2718 typed transaction envelope)
	txType := rawTxBytes[0]
	if txType != 0x02 {
		log.Fatalf("Unexpected transaction type: got 0x%x, expected 0x02 (EIP-1559)", txType)
	}

	// Strip the type byte (EIP-1559 payload follows)
	txPayload := rawTxBytes[:]

	// Decode using go-ethereum's EIP-1559 transaction structure
	var tx = new(types.Transaction)
	err = tx.UnmarshalBinary(txPayload)
	if err != nil {
		log.Fatalf("Failed to decode EIP-1559 tx: %v", err)
	}

	// Output transaction details
	fmt.Println("Transaction decoded successfully:")
	fmt.Printf("Nonce: %d\n", tx.Nonce())
	fmt.Printf("To: %v\n", tx.To())
	fmt.Printf("Value: %s\n", tx.Value())
	fmt.Printf("Gas Tip Cap (maxPriorityFeePerGas): %s\n", tx.GasTipCap()) // dynamic fee
	fmt.Printf("Gas Fee Cap (maxFeePerGas): %s\n", tx.GasFeeCap())         // dynamic fee
	fmt.Printf("Gas: %d\n", tx.Gas())
	fmt.Printf("Data: 0x%x\n", tx.Data())

	signer := types.NewLondonSigner(tx.ChainId())
	from, err := types.Sender(signer, tx)
	if err != nil {
		log.Fatalf("Failed to derive sender: %v", err)
	}
	fmt.Printf("From: %s\n", from.Hex())
	fmt.Printf("ChainID: %s\n", tx.ChainId().String())

	fmt.Printf("fee: 0x%x\n", fee)
	fmt.Println("cmp  0xbb2affbd7b1")

	//sendParam := new(
	mabi, err := abi.JSON(strings.NewReader(ABI))
	//err = rlp.DecodeBytes(tx.Data(), &v)
	if err != nil {
		t.Error(err)
	}
	send := mabi.Methods["send"]

	// 5. Decode
	args, err := send.Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		log.Fatalf("failed to decode args: %v", err)
	}

	fmt.Println("Beneficiary:", args[2])
	// Type assertions
	sendParam := oft.SendParam{}
	CopyFields(&sendParam, args[0])
	fmt.Println(sendParam.ExtraOptions)

}

// CopyFields copies values from src to dst where fields have the same name and type.
func CopyFields(dst interface{}, src interface{}) error {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)

	// dst must be a pointer to a struct
	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to a struct")
	}
	dstVal = dstVal.Elem()

	// src must be a struct
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	if srcVal.Kind() != reflect.Struct {
		return fmt.Errorf("source must be a struct or pointer to struct")
	}

	dstType := dstVal.Type()
	//srcType := srcVal.Type()

	for i := 0; i < dstVal.NumField(); i++ {
		dstField := dstVal.Field(i)
		dstFieldType := dstType.Field(i)

		// Find matching field by name
		srcFieldVal := srcVal.FieldByName(dstFieldType.Name)
		if !srcFieldVal.IsValid() {
			continue
		}

		// Only settable and type-compatible fields
		if dstField.CanSet() && srcFieldVal.Type() == dstField.Type() {
			dstField.Set(srcFieldVal)
		}
	}

	return nil
}

const quoteSendData = `0x3b6f743b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009d35000000000000000000000000cef702bd69926b13ab7150624daa7afee03007860000000000000000000000000000000000000000000000000186cc6acd4b00000000000000000000000000000000000000000000000000000186cc6acd4b000000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000002000300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000`

const fee = 12862084601777 //12862084601777, 114576763194326

const ABI = `[
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_name",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_version",
				"type": "string"
			},
			{
				"internalType": "address",
				"name": "_lzEndpoint",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_delegate",
				"type": "address"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "allowance",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "needed",
				"type": "uint256"
			}
		],
		"name": "ERC20InsufficientAllowance",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "balance",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "needed",
				"type": "uint256"
			}
		],
		"name": "ERC20InsufficientBalance",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "approver",
				"type": "address"
			}
		],
		"name": "ERC20InvalidApprover",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			}
		],
		"name": "ERC20InvalidReceiver",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			}
		],
		"name": "ERC20InvalidSender",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			}
		],
		"name": "ERC20InvalidSpender",
		"type": "error"
	},
	{
		"inputs": [],
		"name": "InvalidDelegate",
		"type": "error"
	},
	{
		"inputs": [],
		"name": "InvalidEndpointCall",
		"type": "error"
	},
	{
		"inputs": [],
		"name": "InvalidLocalDecimals",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "bytes",
				"name": "options",
				"type": "bytes"
			}
		],
		"name": "InvalidOptions",
		"type": "error"
	},
	{
		"inputs": [],
		"name": "LzTokenUnavailable",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "eid",
				"type": "uint32"
			}
		],
		"name": "NoPeer",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "msgValue",
				"type": "uint256"
			}
		],
		"name": "NotEnoughNative",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "addr",
				"type": "address"
			}
		],
		"name": "OnlyEndpoint",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "eid",
				"type": "uint32"
			},
			{
				"internalType": "bytes32",
				"name": "sender",
				"type": "bytes32"
			}
		],
		"name": "OnlyPeer",
		"type": "error"
	},
	{
		"inputs": [],
		"name": "OnlySelf",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "OwnableInvalidOwner",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			}
		],
		"name": "OwnableUnauthorizedAccount",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			}
		],
		"name": "SafeERC20FailedOperation",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "bytes",
				"name": "result",
				"type": "bytes"
			}
		],
		"name": "SimulationResult",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountLD",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "minAmountLD",
				"type": "uint256"
			}
		],
		"name": "SlippageExceeded",
		"type": "error"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Approval",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "authorizer",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "nonce",
				"type": "bytes32"
			}
		],
		"name": "AuthorizationCanceled",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "authorizer",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "nonce",
				"type": "bytes32"
			}
		],
		"name": "AuthorizationUsed",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "eid",
						"type": "uint32"
					},
					{
						"internalType": "uint16",
						"name": "msgType",
						"type": "uint16"
					},
					{
						"internalType": "bytes",
						"name": "options",
						"type": "bytes"
					}
				],
				"indexed": false,
				"internalType": "struct EnforcedOptionParam[]",
				"name": "_enforcedOptions",
				"type": "tuple[]"
			}
		],
		"name": "EnforcedOptionSet",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "inspector",
				"type": "address"
			}
		],
		"name": "MsgInspectorSet",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "guid",
				"type": "bytes32"
			},
			{
				"indexed": false,
				"internalType": "uint32",
				"name": "srcEid",
				"type": "uint32"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "toAddress",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amountReceivedLD",
				"type": "uint256"
			}
		],
		"name": "OFTReceived",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "guid",
				"type": "bytes32"
			},
			{
				"indexed": false,
				"internalType": "uint32",
				"name": "dstEid",
				"type": "uint32"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "fromAddress",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amountSentLD",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amountReceivedLD",
				"type": "uint256"
			}
		],
		"name": "OFTSent",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint32",
				"name": "eid",
				"type": "uint32"
			},
			{
				"indexed": false,
				"internalType": "bytes32",
				"name": "peer",
				"type": "bytes32"
			}
		],
		"name": "PeerSet",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "preCrimeAddress",
				"type": "address"
			}
		],
		"name": "PreCrimeSet",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	},
	{
		"inputs": [],
		"name": "CANCEL_AUTHORIZATION_TYPEHASH",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "DOMAIN_SEPARATOR",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "SEND",
		"outputs": [
			{
				"internalType": "uint16",
				"name": "",
				"type": "uint16"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "SEND_AND_CALL",
		"outputs": [
			{
				"internalType": "uint16",
				"name": "",
				"type": "uint16"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "TRANSFER_WITH_AUTHORIZATION_TYPEHASH",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "srcEid",
						"type": "uint32"
					},
					{
						"internalType": "bytes32",
						"name": "sender",
						"type": "bytes32"
					},
					{
						"internalType": "uint64",
						"name": "nonce",
						"type": "uint64"
					}
				],
				"internalType": "struct Origin",
				"name": "origin",
				"type": "tuple"
			}
		],
		"name": "allowInitializePath",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			}
		],
		"name": "allowance",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "approvalRequired",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "approve",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "authorizer",
				"type": "address"
			},
			{
				"internalType": "bytes32",
				"name": "nonce",
				"type": "bytes32"
			}
		],
		"name": "authorizationState",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "authorizer",
				"type": "address"
			},
			{
				"internalType": "bytes32",
				"name": "nonce",
				"type": "bytes32"
			},
			{
				"internalType": "uint8",
				"name": "v",
				"type": "uint8"
			},
			{
				"internalType": "bytes32",
				"name": "r",
				"type": "bytes32"
			},
			{
				"internalType": "bytes32",
				"name": "s",
				"type": "bytes32"
			}
		],
		"name": "cancelAuthorization",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "authorizer",
				"type": "address"
			},
			{
				"internalType": "bytes32",
				"name": "nonce",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "cancelAuthorization",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "_eid",
				"type": "uint32"
			},
			{
				"internalType": "uint16",
				"name": "_msgType",
				"type": "uint16"
			},
			{
				"internalType": "bytes",
				"name": "_extraOptions",
				"type": "bytes"
			}
		],
		"name": "combineOptions",
		"outputs": [
			{
				"internalType": "bytes",
				"name": "",
				"type": "bytes"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "decimalConversionRate",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "decimals",
		"outputs": [
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "endpoint",
		"outputs": [
			{
				"internalType": "contract ILayerZeroEndpointV2",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "eid",
				"type": "uint32"
			},
			{
				"internalType": "uint16",
				"name": "msgType",
				"type": "uint16"
			}
		],
		"name": "enforcedOptions",
		"outputs": [
			{
				"internalType": "bytes",
				"name": "enforcedOption",
				"type": "bytes"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "srcEid",
						"type": "uint32"
					},
					{
						"internalType": "bytes32",
						"name": "sender",
						"type": "bytes32"
					},
					{
						"internalType": "uint64",
						"name": "nonce",
						"type": "uint64"
					}
				],
				"internalType": "struct Origin",
				"name": "",
				"type": "tuple"
			},
			{
				"internalType": "bytes",
				"name": "",
				"type": "bytes"
			},
			{
				"internalType": "address",
				"name": "_sender",
				"type": "address"
			}
		],
		"name": "isComposeMsgSender",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "_eid",
				"type": "uint32"
			},
			{
				"internalType": "bytes32",
				"name": "_peer",
				"type": "bytes32"
			}
		],
		"name": "isPeer",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "srcEid",
						"type": "uint32"
					},
					{
						"internalType": "bytes32",
						"name": "sender",
						"type": "bytes32"
					},
					{
						"internalType": "uint64",
						"name": "nonce",
						"type": "uint64"
					}
				],
				"internalType": "struct Origin",
				"name": "_origin",
				"type": "tuple"
			},
			{
				"internalType": "bytes32",
				"name": "_guid",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "_message",
				"type": "bytes"
			},
			{
				"internalType": "address",
				"name": "_executor",
				"type": "address"
			},
			{
				"internalType": "bytes",
				"name": "_extraData",
				"type": "bytes"
			}
		],
		"name": "lzReceive",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"components": [
							{
								"internalType": "uint32",
								"name": "srcEid",
								"type": "uint32"
							},
							{
								"internalType": "bytes32",
								"name": "sender",
								"type": "bytes32"
							},
							{
								"internalType": "uint64",
								"name": "nonce",
								"type": "uint64"
							}
						],
						"internalType": "struct Origin",
						"name": "origin",
						"type": "tuple"
					},
					{
						"internalType": "uint32",
						"name": "dstEid",
						"type": "uint32"
					},
					{
						"internalType": "address",
						"name": "receiver",
						"type": "address"
					},
					{
						"internalType": "bytes32",
						"name": "guid",
						"type": "bytes32"
					},
					{
						"internalType": "uint256",
						"name": "value",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "executor",
						"type": "address"
					},
					{
						"internalType": "bytes",
						"name": "message",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "extraData",
						"type": "bytes"
					}
				],
				"internalType": "struct InboundPacket[]",
				"name": "_packets",
				"type": "tuple[]"
			}
		],
		"name": "lzReceiveAndRevert",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "srcEid",
						"type": "uint32"
					},
					{
						"internalType": "bytes32",
						"name": "sender",
						"type": "bytes32"
					},
					{
						"internalType": "uint64",
						"name": "nonce",
						"type": "uint64"
					}
				],
				"internalType": "struct Origin",
				"name": "_origin",
				"type": "tuple"
			},
			{
				"internalType": "bytes32",
				"name": "_guid",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "_message",
				"type": "bytes"
			},
			{
				"internalType": "address",
				"name": "_executor",
				"type": "address"
			},
			{
				"internalType": "bytes",
				"name": "_extraData",
				"type": "bytes"
			}
		],
		"name": "lzReceiveSimulate",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "mint",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "msgInspector",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "name",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "",
				"type": "uint32"
			},
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"name": "nextNonce",
		"outputs": [
			{
				"internalType": "uint64",
				"name": "nonce",
				"type": "uint64"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "oApp",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "oAppVersion",
		"outputs": [
			{
				"internalType": "uint64",
				"name": "senderVersion",
				"type": "uint64"
			},
			{
				"internalType": "uint64",
				"name": "receiverVersion",
				"type": "uint64"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "oftVersion",
		"outputs": [
			{
				"internalType": "bytes4",
				"name": "interfaceId",
				"type": "bytes4"
			},
			{
				"internalType": "uint64",
				"name": "version",
				"type": "uint64"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "eid",
				"type": "uint32"
			}
		],
		"name": "peers",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "peer",
				"type": "bytes32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "preCrime",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "dstEid",
						"type": "uint32"
					},
					{
						"internalType": "bytes32",
						"name": "to",
						"type": "bytes32"
					},
					{
						"internalType": "uint256",
						"name": "amountLD",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "minAmountLD",
						"type": "uint256"
					},
					{
						"internalType": "bytes",
						"name": "extraOptions",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "composeMsg",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "oftCmd",
						"type": "bytes"
					}
				],
				"internalType": "struct SendParam",
				"name": "_sendParam",
				"type": "tuple"
			}
		],
		"name": "quoteOFT",
		"outputs": [
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "minAmountLD",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "maxAmountLD",
						"type": "uint256"
					}
				],
				"internalType": "struct OFTLimit",
				"name": "oftLimit",
				"type": "tuple"
			},
			{
				"components": [
					{
						"internalType": "int256",
						"name": "feeAmountLD",
						"type": "int256"
					},
					{
						"internalType": "string",
						"name": "description",
						"type": "string"
					}
				],
				"internalType": "struct OFTFeeDetail[]",
				"name": "oftFeeDetails",
				"type": "tuple[]"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "amountSentLD",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "amountReceivedLD",
						"type": "uint256"
					}
				],
				"internalType": "struct OFTReceipt",
				"name": "oftReceipt",
				"type": "tuple"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "dstEid",
						"type": "uint32"
					},
					{
						"internalType": "bytes32",
						"name": "to",
						"type": "bytes32"
					},
					{
						"internalType": "uint256",
						"name": "amountLD",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "minAmountLD",
						"type": "uint256"
					},
					{
						"internalType": "bytes",
						"name": "extraOptions",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "composeMsg",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "oftCmd",
						"type": "bytes"
					}
				],
				"internalType": "struct SendParam",
				"name": "_sendParam",
				"type": "tuple"
			},
			{
				"internalType": "bool",
				"name": "_payInLzToken",
				"type": "bool"
			}
		],
		"name": "quoteSend",
		"outputs": [
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "nativeFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "lzTokenFee",
						"type": "uint256"
					}
				],
				"internalType": "struct MessagingFee",
				"name": "msgFee",
				"type": "tuple"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "dstEid",
						"type": "uint32"
					},
					{
						"internalType": "bytes32",
						"name": "to",
						"type": "bytes32"
					},
					{
						"internalType": "uint256",
						"name": "amountLD",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "minAmountLD",
						"type": "uint256"
					},
					{
						"internalType": "bytes",
						"name": "extraOptions",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "composeMsg",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "oftCmd",
						"type": "bytes"
					}
				],
				"internalType": "struct SendParam",
				"name": "_sendParam",
				"type": "tuple"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "nativeFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "lzTokenFee",
						"type": "uint256"
					}
				],
				"internalType": "struct MessagingFee",
				"name": "_fee",
				"type": "tuple"
			},
			{
				"internalType": "address",
				"name": "_refundAddress",
				"type": "address"
			}
		],
		"name": "send",
		"outputs": [
			{
				"components": [
					{
						"internalType": "bytes32",
						"name": "guid",
						"type": "bytes32"
					},
					{
						"internalType": "uint64",
						"name": "nonce",
						"type": "uint64"
					},
					{
						"components": [
							{
								"internalType": "uint256",
								"name": "nativeFee",
								"type": "uint256"
							},
							{
								"internalType": "uint256",
								"name": "lzTokenFee",
								"type": "uint256"
							}
						],
						"internalType": "struct MessagingFee",
						"name": "fee",
						"type": "tuple"
					}
				],
				"internalType": "struct MessagingReceipt",
				"name": "msgReceipt",
				"type": "tuple"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "amountSentLD",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "amountReceivedLD",
						"type": "uint256"
					}
				],
				"internalType": "struct OFTReceipt",
				"name": "oftReceipt",
				"type": "tuple"
			}
		],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "dstEid",
						"type": "uint32"
					},
					{
						"internalType": "bytes32",
						"name": "to",
						"type": "bytes32"
					},
					{
						"internalType": "uint256",
						"name": "amountLD",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "minAmountLD",
						"type": "uint256"
					},
					{
						"internalType": "bytes",
						"name": "extraOptions",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "composeMsg",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "oftCmd",
						"type": "bytes"
					}
				],
				"internalType": "struct SendParam",
				"name": "_sendParam",
				"type": "tuple"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "nativeFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "lzTokenFee",
						"type": "uint256"
					}
				],
				"internalType": "struct MessagingFee",
				"name": "_fee",
				"type": "tuple"
			},
			{
				"internalType": "address",
				"name": "_from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_refundAddress",
				"type": "address"
			}
		],
		"name": "sendFrom",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "dstEid",
						"type": "uint32"
					},
					{
						"internalType": "bytes32",
						"name": "to",
						"type": "bytes32"
					},
					{
						"internalType": "uint256",
						"name": "amountLD",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "minAmountLD",
						"type": "uint256"
					},
					{
						"internalType": "bytes",
						"name": "extraOptions",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "composeMsg",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "oftCmd",
						"type": "bytes"
					}
				],
				"internalType": "struct SendParam",
				"name": "_sendParam",
				"type": "tuple"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "nativeFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "lzTokenFee",
						"type": "uint256"
					}
				],
				"internalType": "struct MessagingFee",
				"name": "_fee",
				"type": "tuple"
			},
			{
				"internalType": "address",
				"name": "_from",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "validAfter",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "validBefore",
				"type": "uint256"
			},
			{
				"internalType": "bytes32",
				"name": "nonce",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			},
			{
				"internalType": "address",
				"name": "_refundAddress",
				"type": "address"
			}
		],
		"name": "sendWithAuthorization",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_delegate",
				"type": "address"
			}
		],
		"name": "setDelegate",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "uint32",
						"name": "eid",
						"type": "uint32"
					},
					{
						"internalType": "uint16",
						"name": "msgType",
						"type": "uint16"
					},
					{
						"internalType": "bytes",
						"name": "options",
						"type": "bytes"
					}
				],
				"internalType": "struct EnforcedOptionParam[]",
				"name": "_enforcedOptions",
				"type": "tuple[]"
			}
		],
		"name": "setEnforcedOptions",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_msgInspector",
				"type": "address"
			}
		],
		"name": "setMsgInspector",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "_eid",
				"type": "uint32"
			},
			{
				"internalType": "bytes32",
				"name": "_peer",
				"type": "bytes32"
			}
		],
		"name": "setPeer",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_preCrime",
				"type": "address"
			}
		],
		"name": "setPreCrime",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "sharedDecimals",
		"outputs": [
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "symbol",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "token",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "totalSupply",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "transfer",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "transferFrom",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "validAfter",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "validBefore",
				"type": "uint256"
			},
			{
				"internalType": "bytes32",
				"name": "nonce",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "transferWithAuthorization",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "validAfter",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "validBefore",
				"type": "uint256"
			},
			{
				"internalType": "bytes32",
				"name": "nonce",
				"type": "bytes32"
			},
			{
				"internalType": "uint8",
				"name": "v",
				"type": "uint8"
			},
			{
				"internalType": "bytes32",
				"name": "r",
				"type": "bytes32"
			},
			{
				"internalType": "bytes32",
				"name": "s",
				"type": "bytes32"
			}
		],
		"name": "transferWithAuthorization",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "version",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`
