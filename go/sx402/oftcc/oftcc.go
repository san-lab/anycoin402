// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oftcc

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EnforcedOptionParam is an auto generated low-level Go binding around an user-defined struct.
type EnforcedOptionParam struct {
	Eid     uint32
	MsgType uint16
	Options []byte
}

// InboundPacket is an auto generated low-level Go binding around an user-defined struct.
type InboundPacket struct {
	Origin    Origin
	DstEid    uint32
	Receiver  common.Address
	Guid      [32]byte
	Value     *big.Int
	Executor  common.Address
	Message   []byte
	ExtraData []byte
}

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// OFTFeeDetail is an auto generated low-level Go binding around an user-defined struct.
type OFTFeeDetail struct {
	FeeAmountLD *big.Int
	Description string
}

// OFTLimit is an auto generated low-level Go binding around an user-defined struct.
type OFTLimit struct {
	MinAmountLD *big.Int
	MaxAmountLD *big.Int
}

// OFTReceipt is an auto generated low-level Go binding around an user-defined struct.
type OFTReceipt struct {
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// SendParam is an auto generated low-level Go binding around an user-defined struct.
type SendParam struct {
	DstEid       uint32
	To           [32]byte
	AmountLD     *big.Int
	MinAmountLD  *big.Int
	ExtraOptions []byte
	ComposeMsg   []byte
	OftCmd       []byte
}

// OftccMetaData contains all meta data concerning the Oftcc contract.
var OftccMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"AmountSDOverflowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLocalDecimals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"InvalidOptions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"SimulationResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"}],\"name\":\"SlippageExceeded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"AuthorizationCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"AuthorizationUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"EnforcedOptionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inspector\",\"type\":\"address\"}],\"name\":\"MsgInspectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"preCrimeAddress\",\"type\":\"address\"}],\"name\":\"PreCrimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANCEL_AUTHORIZATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_TRANSFER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstEid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"HashCCData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEND\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEND_AND_CALL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_WITH_AUTHORIZATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvalRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"authorizationState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancelAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"cancelAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_extraOptions\",\"type\":\"bytes\"}],\"name\":\"combineOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalConversionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"}],\"name\":\"enforcedOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"enforcedOption\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"isPeer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"localMarkups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structInboundPacket[]\",\"name\":\"_packets\",\"type\":\"tuple[]\"}],\"name\":\"lzReceiveAndRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceiveSimulate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"markups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgInspector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oftVersion\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preCrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"}],\"name\":\"quoteOFT\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTLimit\",\"name\":\"oftLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"feeAmountLD\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structOFTFeeDetail[]\",\"name\":\"oftFeeDetails\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_payInLzToken\",\"type\":\"bool\"}],\"name\":\"quoteSend\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"msgFee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"msgReceipt\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendWithCCAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_markup\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"}],\"name\":\"setCrosschainMarkup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"setEnforcedOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_markup\",\"type\":\"uint256\"}],\"name\":\"setLocalMarkup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgInspector\",\"type\":\"address\"}],\"name\":\"setMsgInspector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_preCrime\",\"type\":\"address\"}],\"name\":\"setPreCrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"transferWithAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"transferWithAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OftccABI is the input ABI used to generate the binding from.
// Deprecated: Use OftccMetaData.ABI instead.
var OftccABI = OftccMetaData.ABI

// Oftcc is an auto generated Go binding around an Ethereum contract.
type Oftcc struct {
	OftccCaller     // Read-only binding to the contract
	OftccTransactor // Write-only binding to the contract
	OftccFilterer   // Log filterer for contract events
}

// OftccCaller is an auto generated read-only Go binding around an Ethereum contract.
type OftccCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OftccTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OftccTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OftccFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OftccFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OftccSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OftccSession struct {
	Contract     *Oftcc            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OftccCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OftccCallerSession struct {
	Contract *OftccCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OftccTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OftccTransactorSession struct {
	Contract     *OftccTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OftccRaw is an auto generated low-level Go binding around an Ethereum contract.
type OftccRaw struct {
	Contract *Oftcc // Generic contract binding to access the raw methods on
}

// OftccCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OftccCallerRaw struct {
	Contract *OftccCaller // Generic read-only contract binding to access the raw methods on
}

// OftccTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OftccTransactorRaw struct {
	Contract *OftccTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOftcc creates a new instance of Oftcc, bound to a specific deployed contract.
func NewOftcc(address common.Address, backend bind.ContractBackend) (*Oftcc, error) {
	contract, err := bindOftcc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oftcc{OftccCaller: OftccCaller{contract: contract}, OftccTransactor: OftccTransactor{contract: contract}, OftccFilterer: OftccFilterer{contract: contract}}, nil
}

// NewOftccCaller creates a new read-only instance of Oftcc, bound to a specific deployed contract.
func NewOftccCaller(address common.Address, caller bind.ContractCaller) (*OftccCaller, error) {
	contract, err := bindOftcc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OftccCaller{contract: contract}, nil
}

// NewOftccTransactor creates a new write-only instance of Oftcc, bound to a specific deployed contract.
func NewOftccTransactor(address common.Address, transactor bind.ContractTransactor) (*OftccTransactor, error) {
	contract, err := bindOftcc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OftccTransactor{contract: contract}, nil
}

// NewOftccFilterer creates a new log filterer instance of Oftcc, bound to a specific deployed contract.
func NewOftccFilterer(address common.Address, filterer bind.ContractFilterer) (*OftccFilterer, error) {
	contract, err := bindOftcc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OftccFilterer{contract: contract}, nil
}

// bindOftcc binds a generic wrapper to an already deployed contract.
func bindOftcc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OftccMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oftcc *OftccRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oftcc.Contract.OftccCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oftcc *OftccRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oftcc.Contract.OftccTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oftcc *OftccRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oftcc.Contract.OftccTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oftcc *OftccCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oftcc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oftcc *OftccTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oftcc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oftcc *OftccTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oftcc.Contract.contract.Transact(opts, method, params...)
}

// CANCELAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xd9169487.
//
// Solidity: function CANCEL_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oftcc *OftccCaller) CANCELAUTHORIZATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "CANCEL_AUTHORIZATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xd9169487.
//
// Solidity: function CANCEL_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oftcc *OftccSession) CANCELAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Oftcc.Contract.CANCELAUTHORIZATIONTYPEHASH(&_Oftcc.CallOpts)
}

// CANCELAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xd9169487.
//
// Solidity: function CANCEL_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oftcc *OftccCallerSession) CANCELAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Oftcc.Contract.CANCELAUTHORIZATIONTYPEHASH(&_Oftcc.CallOpts)
}

// CROSSCHAINTRANSFERTYPEHASH is a free data retrieval call binding the contract method 0xde3f8fbe.
//
// Solidity: function CROSS_CHAIN_TRANSFER_TYPEHASH() view returns(bytes32)
func (_Oftcc *OftccCaller) CROSSCHAINTRANSFERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "CROSS_CHAIN_TRANSFER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CROSSCHAINTRANSFERTYPEHASH is a free data retrieval call binding the contract method 0xde3f8fbe.
//
// Solidity: function CROSS_CHAIN_TRANSFER_TYPEHASH() view returns(bytes32)
func (_Oftcc *OftccSession) CROSSCHAINTRANSFERTYPEHASH() ([32]byte, error) {
	return _Oftcc.Contract.CROSSCHAINTRANSFERTYPEHASH(&_Oftcc.CallOpts)
}

// CROSSCHAINTRANSFERTYPEHASH is a free data retrieval call binding the contract method 0xde3f8fbe.
//
// Solidity: function CROSS_CHAIN_TRANSFER_TYPEHASH() view returns(bytes32)
func (_Oftcc *OftccCallerSession) CROSSCHAINTRANSFERTYPEHASH() ([32]byte, error) {
	return _Oftcc.Contract.CROSSCHAINTRANSFERTYPEHASH(&_Oftcc.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Oftcc *OftccCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Oftcc *OftccSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Oftcc.Contract.DOMAINSEPARATOR(&_Oftcc.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Oftcc *OftccCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Oftcc.Contract.DOMAINSEPARATOR(&_Oftcc.CallOpts)
}

// HashCCData is a free data retrieval call binding the contract method 0xf43434a4.
//
// Solidity: function HashCCData(address from, address to, uint256 amountLD, uint256 minAmountLD, uint256 dstEid, uint256 validAfter, uint256 validBefore, bytes32 nonce) pure returns(bytes32)
func (_Oftcc *OftccCaller) HashCCData(opts *bind.CallOpts, from common.Address, to common.Address, amountLD *big.Int, minAmountLD *big.Int, dstEid *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "HashCCData", from, to, amountLD, minAmountLD, dstEid, validAfter, validBefore, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashCCData is a free data retrieval call binding the contract method 0xf43434a4.
//
// Solidity: function HashCCData(address from, address to, uint256 amountLD, uint256 minAmountLD, uint256 dstEid, uint256 validAfter, uint256 validBefore, bytes32 nonce) pure returns(bytes32)
func (_Oftcc *OftccSession) HashCCData(from common.Address, to common.Address, amountLD *big.Int, minAmountLD *big.Int, dstEid *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte) ([32]byte, error) {
	return _Oftcc.Contract.HashCCData(&_Oftcc.CallOpts, from, to, amountLD, minAmountLD, dstEid, validAfter, validBefore, nonce)
}

// HashCCData is a free data retrieval call binding the contract method 0xf43434a4.
//
// Solidity: function HashCCData(address from, address to, uint256 amountLD, uint256 minAmountLD, uint256 dstEid, uint256 validAfter, uint256 validBefore, bytes32 nonce) pure returns(bytes32)
func (_Oftcc *OftccCallerSession) HashCCData(from common.Address, to common.Address, amountLD *big.Int, minAmountLD *big.Int, dstEid *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte) ([32]byte, error) {
	return _Oftcc.Contract.HashCCData(&_Oftcc.CallOpts, from, to, amountLD, minAmountLD, dstEid, validAfter, validBefore, nonce)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Oftcc *OftccCaller) SEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Oftcc *OftccSession) SEND() (uint16, error) {
	return _Oftcc.Contract.SEND(&_Oftcc.CallOpts)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Oftcc *OftccCallerSession) SEND() (uint16, error) {
	return _Oftcc.Contract.SEND(&_Oftcc.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Oftcc *OftccCaller) SENDANDCALL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "SEND_AND_CALL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Oftcc *OftccSession) SENDANDCALL() (uint16, error) {
	return _Oftcc.Contract.SENDANDCALL(&_Oftcc.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Oftcc *OftccCallerSession) SENDANDCALL() (uint16, error) {
	return _Oftcc.Contract.SENDANDCALL(&_Oftcc.CallOpts)
}

// TRANSFERWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xa0cc6a68.
//
// Solidity: function TRANSFER_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oftcc *OftccCaller) TRANSFERWITHAUTHORIZATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "TRANSFER_WITH_AUTHORIZATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xa0cc6a68.
//
// Solidity: function TRANSFER_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oftcc *OftccSession) TRANSFERWITHAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Oftcc.Contract.TRANSFERWITHAUTHORIZATIONTYPEHASH(&_Oftcc.CallOpts)
}

// TRANSFERWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xa0cc6a68.
//
// Solidity: function TRANSFER_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oftcc *OftccCallerSession) TRANSFERWITHAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Oftcc.Contract.TRANSFERWITHAUTHORIZATIONTYPEHASH(&_Oftcc.CallOpts)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Oftcc *OftccCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Oftcc *OftccSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Oftcc.Contract.AllowInitializePath(&_Oftcc.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Oftcc *OftccCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Oftcc.Contract.AllowInitializePath(&_Oftcc.CallOpts, origin)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Oftcc *OftccCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Oftcc *OftccSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Oftcc.Contract.Allowance(&_Oftcc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Oftcc *OftccCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Oftcc.Contract.Allowance(&_Oftcc.CallOpts, owner, spender)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Oftcc *OftccCaller) ApprovalRequired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "approvalRequired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Oftcc *OftccSession) ApprovalRequired() (bool, error) {
	return _Oftcc.Contract.ApprovalRequired(&_Oftcc.CallOpts)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Oftcc *OftccCallerSession) ApprovalRequired() (bool, error) {
	return _Oftcc.Contract.ApprovalRequired(&_Oftcc.CallOpts)
}

// AuthorizationState is a free data retrieval call binding the contract method 0xe94a0102.
//
// Solidity: function authorizationState(address authorizer, bytes32 nonce) view returns(bool)
func (_Oftcc *OftccCaller) AuthorizationState(opts *bind.CallOpts, authorizer common.Address, nonce [32]byte) (bool, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "authorizationState", authorizer, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizationState is a free data retrieval call binding the contract method 0xe94a0102.
//
// Solidity: function authorizationState(address authorizer, bytes32 nonce) view returns(bool)
func (_Oftcc *OftccSession) AuthorizationState(authorizer common.Address, nonce [32]byte) (bool, error) {
	return _Oftcc.Contract.AuthorizationState(&_Oftcc.CallOpts, authorizer, nonce)
}

// AuthorizationState is a free data retrieval call binding the contract method 0xe94a0102.
//
// Solidity: function authorizationState(address authorizer, bytes32 nonce) view returns(bool)
func (_Oftcc *OftccCallerSession) AuthorizationState(authorizer common.Address, nonce [32]byte) (bool, error) {
	return _Oftcc.Contract.AuthorizationState(&_Oftcc.CallOpts, authorizer, nonce)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Oftcc *OftccCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Oftcc *OftccSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Oftcc.Contract.BalanceOf(&_Oftcc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Oftcc *OftccCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Oftcc.Contract.BalanceOf(&_Oftcc.CallOpts, account)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Oftcc *OftccCaller) CombineOptions(opts *bind.CallOpts, _eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "combineOptions", _eid, _msgType, _extraOptions)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Oftcc *OftccSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _Oftcc.Contract.CombineOptions(&_Oftcc.CallOpts, _eid, _msgType, _extraOptions)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Oftcc *OftccCallerSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _Oftcc.Contract.CombineOptions(&_Oftcc.CallOpts, _eid, _msgType, _extraOptions)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Oftcc *OftccCaller) DecimalConversionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "decimalConversionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Oftcc *OftccSession) DecimalConversionRate() (*big.Int, error) {
	return _Oftcc.Contract.DecimalConversionRate(&_Oftcc.CallOpts)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Oftcc *OftccCallerSession) DecimalConversionRate() (*big.Int, error) {
	return _Oftcc.Contract.DecimalConversionRate(&_Oftcc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Oftcc *OftccCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Oftcc *OftccSession) Decimals() (uint8, error) {
	return _Oftcc.Contract.Decimals(&_Oftcc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Oftcc *OftccCallerSession) Decimals() (uint8, error) {
	return _Oftcc.Contract.Decimals(&_Oftcc.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Oftcc *OftccCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Oftcc *OftccSession) Endpoint() (common.Address, error) {
	return _Oftcc.Contract.Endpoint(&_Oftcc.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Oftcc *OftccCallerSession) Endpoint() (common.Address, error) {
	return _Oftcc.Contract.Endpoint(&_Oftcc.CallOpts)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_Oftcc *OftccCaller) EnforcedOptions(opts *bind.CallOpts, eid uint32, msgType uint16) ([]byte, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "enforcedOptions", eid, msgType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_Oftcc *OftccSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _Oftcc.Contract.EnforcedOptions(&_Oftcc.CallOpts, eid, msgType)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_Oftcc *OftccCallerSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _Oftcc.Contract.EnforcedOptions(&_Oftcc.CallOpts, eid, msgType)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Oftcc *OftccCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Oftcc *OftccSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Oftcc.Contract.IsComposeMsgSender(&_Oftcc.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Oftcc *OftccCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Oftcc.Contract.IsComposeMsgSender(&_Oftcc.CallOpts, arg0, arg1, _sender)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Oftcc *OftccCaller) IsPeer(opts *bind.CallOpts, _eid uint32, _peer [32]byte) (bool, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "isPeer", _eid, _peer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Oftcc *OftccSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _Oftcc.Contract.IsPeer(&_Oftcc.CallOpts, _eid, _peer)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Oftcc *OftccCallerSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _Oftcc.Contract.IsPeer(&_Oftcc.CallOpts, _eid, _peer)
}

// LocalMarkups is a free data retrieval call binding the contract method 0x83002918.
//
// Solidity: function localMarkups(address ) view returns(uint256)
func (_Oftcc *OftccCaller) LocalMarkups(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "localMarkups", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalMarkups is a free data retrieval call binding the contract method 0x83002918.
//
// Solidity: function localMarkups(address ) view returns(uint256)
func (_Oftcc *OftccSession) LocalMarkups(arg0 common.Address) (*big.Int, error) {
	return _Oftcc.Contract.LocalMarkups(&_Oftcc.CallOpts, arg0)
}

// LocalMarkups is a free data retrieval call binding the contract method 0x83002918.
//
// Solidity: function localMarkups(address ) view returns(uint256)
func (_Oftcc *OftccCallerSession) LocalMarkups(arg0 common.Address) (*big.Int, error) {
	return _Oftcc.Contract.LocalMarkups(&_Oftcc.CallOpts, arg0)
}

// Markups is a free data retrieval call binding the contract method 0x2148ce13.
//
// Solidity: function markups(address , uint32 ) view returns(uint256)
func (_Oftcc *OftccCaller) Markups(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "markups", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Markups is a free data retrieval call binding the contract method 0x2148ce13.
//
// Solidity: function markups(address , uint32 ) view returns(uint256)
func (_Oftcc *OftccSession) Markups(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Oftcc.Contract.Markups(&_Oftcc.CallOpts, arg0, arg1)
}

// Markups is a free data retrieval call binding the contract method 0x2148ce13.
//
// Solidity: function markups(address , uint32 ) view returns(uint256)
func (_Oftcc *OftccCallerSession) Markups(arg0 common.Address, arg1 uint32) (*big.Int, error) {
	return _Oftcc.Contract.Markups(&_Oftcc.CallOpts, arg0, arg1)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Oftcc *OftccCaller) MsgInspector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "msgInspector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Oftcc *OftccSession) MsgInspector() (common.Address, error) {
	return _Oftcc.Contract.MsgInspector(&_Oftcc.CallOpts)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Oftcc *OftccCallerSession) MsgInspector() (common.Address, error) {
	return _Oftcc.Contract.MsgInspector(&_Oftcc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Oftcc *OftccCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Oftcc *OftccSession) Name() (string, error) {
	return _Oftcc.Contract.Name(&_Oftcc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Oftcc *OftccCallerSession) Name() (string, error) {
	return _Oftcc.Contract.Name(&_Oftcc.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Oftcc *OftccCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Oftcc *OftccSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Oftcc.Contract.NextNonce(&_Oftcc.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Oftcc *OftccCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Oftcc.Contract.NextNonce(&_Oftcc.CallOpts, arg0, arg1)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Oftcc *OftccCaller) OApp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "oApp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Oftcc *OftccSession) OApp() (common.Address, error) {
	return _Oftcc.Contract.OApp(&_Oftcc.CallOpts)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Oftcc *OftccCallerSession) OApp() (common.Address, error) {
	return _Oftcc.Contract.OApp(&_Oftcc.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Oftcc *OftccCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Oftcc *OftccSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Oftcc.Contract.OAppVersion(&_Oftcc.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Oftcc *OftccCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Oftcc.Contract.OAppVersion(&_Oftcc.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Oftcc *OftccCaller) OftVersion(opts *bind.CallOpts) (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "oftVersion")

	outstruct := new(struct {
		InterfaceId [4]byte
		Version     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InterfaceId = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.Version = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Oftcc *OftccSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _Oftcc.Contract.OftVersion(&_Oftcc.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Oftcc *OftccCallerSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _Oftcc.Contract.OftVersion(&_Oftcc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oftcc *OftccCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oftcc *OftccSession) Owner() (common.Address, error) {
	return _Oftcc.Contract.Owner(&_Oftcc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oftcc *OftccCallerSession) Owner() (common.Address, error) {
	return _Oftcc.Contract.Owner(&_Oftcc.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Oftcc *OftccCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Oftcc *OftccSession) Peers(eid uint32) ([32]byte, error) {
	return _Oftcc.Contract.Peers(&_Oftcc.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Oftcc *OftccCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _Oftcc.Contract.Peers(&_Oftcc.CallOpts, eid)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Oftcc *OftccCaller) PreCrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "preCrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Oftcc *OftccSession) PreCrime() (common.Address, error) {
	return _Oftcc.Contract.PreCrime(&_Oftcc.CallOpts)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Oftcc *OftccCallerSession) PreCrime() (common.Address, error) {
	return _Oftcc.Contract.PreCrime(&_Oftcc.CallOpts)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Oftcc *OftccCaller) QuoteOFT(opts *bind.CallOpts, _sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "quoteOFT", _sendParam)

	outstruct := new(struct {
		OftLimit      OFTLimit
		OftFeeDetails []OFTFeeDetail
		OftReceipt    OFTReceipt
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OftLimit = *abi.ConvertType(out[0], new(OFTLimit)).(*OFTLimit)
	outstruct.OftFeeDetails = *abi.ConvertType(out[1], new([]OFTFeeDetail)).(*[]OFTFeeDetail)
	outstruct.OftReceipt = *abi.ConvertType(out[2], new(OFTReceipt)).(*OFTReceipt)

	return *outstruct, err

}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Oftcc *OftccSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _Oftcc.Contract.QuoteOFT(&_Oftcc.CallOpts, _sendParam)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Oftcc *OftccCallerSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _Oftcc.Contract.QuoteOFT(&_Oftcc.CallOpts, _sendParam)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Oftcc *OftccCaller) QuoteSend(opts *bind.CallOpts, _sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "quoteSend", _sendParam, _payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Oftcc *OftccSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Oftcc.Contract.QuoteSend(&_Oftcc.CallOpts, _sendParam, _payInLzToken)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Oftcc *OftccCallerSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Oftcc.Contract.QuoteSend(&_Oftcc.CallOpts, _sendParam, _payInLzToken)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() pure returns(uint8)
func (_Oftcc *OftccCaller) SharedDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "sharedDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() pure returns(uint8)
func (_Oftcc *OftccSession) SharedDecimals() (uint8, error) {
	return _Oftcc.Contract.SharedDecimals(&_Oftcc.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() pure returns(uint8)
func (_Oftcc *OftccCallerSession) SharedDecimals() (uint8, error) {
	return _Oftcc.Contract.SharedDecimals(&_Oftcc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Oftcc *OftccCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Oftcc *OftccSession) Symbol() (string, error) {
	return _Oftcc.Contract.Symbol(&_Oftcc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Oftcc *OftccCallerSession) Symbol() (string, error) {
	return _Oftcc.Contract.Symbol(&_Oftcc.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Oftcc *OftccCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Oftcc *OftccSession) Token() (common.Address, error) {
	return _Oftcc.Contract.Token(&_Oftcc.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Oftcc *OftccCallerSession) Token() (common.Address, error) {
	return _Oftcc.Contract.Token(&_Oftcc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Oftcc *OftccCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Oftcc *OftccSession) TotalSupply() (*big.Int, error) {
	return _Oftcc.Contract.TotalSupply(&_Oftcc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Oftcc *OftccCallerSession) TotalSupply() (*big.Int, error) {
	return _Oftcc.Contract.TotalSupply(&_Oftcc.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Oftcc *OftccCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oftcc.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Oftcc *OftccSession) Version() (string, error) {
	return _Oftcc.Contract.Version(&_Oftcc.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Oftcc *OftccCallerSession) Version() (string, error) {
	return _Oftcc.Contract.Version(&_Oftcc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Oftcc *OftccTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Oftcc *OftccSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.Approve(&_Oftcc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Oftcc *OftccTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.Approve(&_Oftcc.TransactOpts, spender, value)
}

// CancelAuthorization is a paid mutator transaction binding the contract method 0x5a049a70.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oftcc *OftccTransactor) CancelAuthorization(opts *bind.TransactOpts, authorizer common.Address, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "cancelAuthorization", authorizer, nonce, v, r, s)
}

// CancelAuthorization is a paid mutator transaction binding the contract method 0x5a049a70.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oftcc *OftccSession) CancelAuthorization(authorizer common.Address, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oftcc.Contract.CancelAuthorization(&_Oftcc.TransactOpts, authorizer, nonce, v, r, s)
}

// CancelAuthorization is a paid mutator transaction binding the contract method 0x5a049a70.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oftcc *OftccTransactorSession) CancelAuthorization(authorizer common.Address, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oftcc.Contract.CancelAuthorization(&_Oftcc.TransactOpts, authorizer, nonce, v, r, s)
}

// CancelAuthorization0 is a paid mutator transaction binding the contract method 0xb7b72899.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, bytes signature) returns()
func (_Oftcc *OftccTransactor) CancelAuthorization0(opts *bind.TransactOpts, authorizer common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "cancelAuthorization0", authorizer, nonce, signature)
}

// CancelAuthorization0 is a paid mutator transaction binding the contract method 0xb7b72899.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, bytes signature) returns()
func (_Oftcc *OftccSession) CancelAuthorization0(authorizer common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oftcc.Contract.CancelAuthorization0(&_Oftcc.TransactOpts, authorizer, nonce, signature)
}

// CancelAuthorization0 is a paid mutator transaction binding the contract method 0xb7b72899.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, bytes signature) returns()
func (_Oftcc *OftccTransactorSession) CancelAuthorization0(authorizer common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oftcc.Contract.CancelAuthorization0(&_Oftcc.TransactOpts, authorizer, nonce, signature)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oftcc *OftccTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oftcc *OftccSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oftcc.Contract.LzReceive(&_Oftcc.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oftcc *OftccTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oftcc.Contract.LzReceive(&_Oftcc.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Oftcc *OftccTransactor) LzReceiveAndRevert(opts *bind.TransactOpts, _packets []InboundPacket) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "lzReceiveAndRevert", _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Oftcc *OftccSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _Oftcc.Contract.LzReceiveAndRevert(&_Oftcc.TransactOpts, _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Oftcc *OftccTransactorSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _Oftcc.Contract.LzReceiveAndRevert(&_Oftcc.TransactOpts, _packets)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oftcc *OftccTransactor) LzReceiveSimulate(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "lzReceiveSimulate", _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oftcc *OftccSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oftcc.Contract.LzReceiveSimulate(&_Oftcc.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oftcc *OftccTransactorSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oftcc.Contract.LzReceiveSimulate(&_Oftcc.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Oftcc *OftccTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Oftcc *OftccSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.Mint(&_Oftcc.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Oftcc *OftccTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.Mint(&_Oftcc.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oftcc *OftccTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oftcc *OftccSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oftcc.Contract.RenounceOwnership(&_Oftcc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oftcc *OftccTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oftcc.Contract.RenounceOwnership(&_Oftcc.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Oftcc *OftccTransactor) Send(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "send", _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Oftcc *OftccSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.Send(&_Oftcc.TransactOpts, _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Oftcc *OftccTransactorSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.Send(&_Oftcc.TransactOpts, _sendParam, _fee, _refundAddress)
}

// SendFrom is a paid mutator transaction binding the contract method 0xa6e329f5.
//
// Solidity: function sendFrom((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, address _refundAddress) payable returns(bool)
func (_Oftcc *OftccTransactor) SendFrom(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _from common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "sendFrom", _sendParam, _fee, _from, _refundAddress)
}

// SendFrom is a paid mutator transaction binding the contract method 0xa6e329f5.
//
// Solidity: function sendFrom((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, address _refundAddress) payable returns(bool)
func (_Oftcc *OftccSession) SendFrom(_sendParam SendParam, _fee MessagingFee, _from common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SendFrom(&_Oftcc.TransactOpts, _sendParam, _fee, _from, _refundAddress)
}

// SendFrom is a paid mutator transaction binding the contract method 0xa6e329f5.
//
// Solidity: function sendFrom((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, address _refundAddress) payable returns(bool)
func (_Oftcc *OftccTransactorSession) SendFrom(_sendParam SendParam, _fee MessagingFee, _from common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SendFrom(&_Oftcc.TransactOpts, _sendParam, _fee, _from, _refundAddress)
}

// SendWithCCAuthorization is a paid mutator transaction binding the contract method 0x85ddcc42.
//
// Solidity: function sendWithCCAuthorization((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature, address _refundAddress) payable returns(bool)
func (_Oftcc *OftccTransactor) SendWithCCAuthorization(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _from common.Address, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "sendWithCCAuthorization", _sendParam, _fee, _from, validAfter, validBefore, nonce, signature, _refundAddress)
}

// SendWithCCAuthorization is a paid mutator transaction binding the contract method 0x85ddcc42.
//
// Solidity: function sendWithCCAuthorization((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature, address _refundAddress) payable returns(bool)
func (_Oftcc *OftccSession) SendWithCCAuthorization(_sendParam SendParam, _fee MessagingFee, _from common.Address, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SendWithCCAuthorization(&_Oftcc.TransactOpts, _sendParam, _fee, _from, validAfter, validBefore, nonce, signature, _refundAddress)
}

// SendWithCCAuthorization is a paid mutator transaction binding the contract method 0x85ddcc42.
//
// Solidity: function sendWithCCAuthorization((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature, address _refundAddress) payable returns(bool)
func (_Oftcc *OftccTransactorSession) SendWithCCAuthorization(_sendParam SendParam, _fee MessagingFee, _from common.Address, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SendWithCCAuthorization(&_Oftcc.TransactOpts, _sendParam, _fee, _from, validAfter, validBefore, nonce, signature, _refundAddress)
}

// SetCrosschainMarkup is a paid mutator transaction binding the contract method 0xa158c790.
//
// Solidity: function setCrosschainMarkup(uint256 _markup, uint32 dstEid) returns()
func (_Oftcc *OftccTransactor) SetCrosschainMarkup(opts *bind.TransactOpts, _markup *big.Int, dstEid uint32) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "setCrosschainMarkup", _markup, dstEid)
}

// SetCrosschainMarkup is a paid mutator transaction binding the contract method 0xa158c790.
//
// Solidity: function setCrosschainMarkup(uint256 _markup, uint32 dstEid) returns()
func (_Oftcc *OftccSession) SetCrosschainMarkup(_markup *big.Int, dstEid uint32) (*types.Transaction, error) {
	return _Oftcc.Contract.SetCrosschainMarkup(&_Oftcc.TransactOpts, _markup, dstEid)
}

// SetCrosschainMarkup is a paid mutator transaction binding the contract method 0xa158c790.
//
// Solidity: function setCrosschainMarkup(uint256 _markup, uint32 dstEid) returns()
func (_Oftcc *OftccTransactorSession) SetCrosschainMarkup(_markup *big.Int, dstEid uint32) (*types.Transaction, error) {
	return _Oftcc.Contract.SetCrosschainMarkup(&_Oftcc.TransactOpts, _markup, dstEid)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Oftcc *OftccTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Oftcc *OftccSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SetDelegate(&_Oftcc.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Oftcc *OftccTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SetDelegate(&_Oftcc.TransactOpts, _delegate)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Oftcc *OftccTransactor) SetEnforcedOptions(opts *bind.TransactOpts, _enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "setEnforcedOptions", _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Oftcc *OftccSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Oftcc.Contract.SetEnforcedOptions(&_Oftcc.TransactOpts, _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Oftcc *OftccTransactorSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Oftcc.Contract.SetEnforcedOptions(&_Oftcc.TransactOpts, _enforcedOptions)
}

// SetLocalMarkup is a paid mutator transaction binding the contract method 0x8f91bcc2.
//
// Solidity: function setLocalMarkup(uint256 _markup) returns()
func (_Oftcc *OftccTransactor) SetLocalMarkup(opts *bind.TransactOpts, _markup *big.Int) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "setLocalMarkup", _markup)
}

// SetLocalMarkup is a paid mutator transaction binding the contract method 0x8f91bcc2.
//
// Solidity: function setLocalMarkup(uint256 _markup) returns()
func (_Oftcc *OftccSession) SetLocalMarkup(_markup *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.SetLocalMarkup(&_Oftcc.TransactOpts, _markup)
}

// SetLocalMarkup is a paid mutator transaction binding the contract method 0x8f91bcc2.
//
// Solidity: function setLocalMarkup(uint256 _markup) returns()
func (_Oftcc *OftccTransactorSession) SetLocalMarkup(_markup *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.SetLocalMarkup(&_Oftcc.TransactOpts, _markup)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Oftcc *OftccTransactor) SetMsgInspector(opts *bind.TransactOpts, _msgInspector common.Address) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "setMsgInspector", _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Oftcc *OftccSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SetMsgInspector(&_Oftcc.TransactOpts, _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Oftcc *OftccTransactorSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SetMsgInspector(&_Oftcc.TransactOpts, _msgInspector)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Oftcc *OftccTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Oftcc *OftccSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Oftcc.Contract.SetPeer(&_Oftcc.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Oftcc *OftccTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Oftcc.Contract.SetPeer(&_Oftcc.TransactOpts, _eid, _peer)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Oftcc *OftccTransactor) SetPreCrime(opts *bind.TransactOpts, _preCrime common.Address) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "setPreCrime", _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Oftcc *OftccSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SetPreCrime(&_Oftcc.TransactOpts, _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Oftcc *OftccTransactorSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.SetPreCrime(&_Oftcc.TransactOpts, _preCrime)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Oftcc *OftccTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Oftcc *OftccSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.Transfer(&_Oftcc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Oftcc *OftccTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.Transfer(&_Oftcc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Oftcc *OftccTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Oftcc *OftccSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.TransferFrom(&_Oftcc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Oftcc *OftccTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oftcc.Contract.TransferFrom(&_Oftcc.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oftcc *OftccTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oftcc *OftccSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.TransferOwnership(&_Oftcc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oftcc *OftccTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oftcc.Contract.TransferOwnership(&_Oftcc.TransactOpts, newOwner)
}

// TransferWithAuthorization is a paid mutator transaction binding the contract method 0xcf092995.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature) returns()
func (_Oftcc *OftccTransactor) TransferWithAuthorization(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "transferWithAuthorization", from, to, value, validAfter, validBefore, nonce, signature)
}

// TransferWithAuthorization is a paid mutator transaction binding the contract method 0xcf092995.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature) returns()
func (_Oftcc *OftccSession) TransferWithAuthorization(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oftcc.Contract.TransferWithAuthorization(&_Oftcc.TransactOpts, from, to, value, validAfter, validBefore, nonce, signature)
}

// TransferWithAuthorization is a paid mutator transaction binding the contract method 0xcf092995.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature) returns()
func (_Oftcc *OftccTransactorSession) TransferWithAuthorization(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oftcc.Contract.TransferWithAuthorization(&_Oftcc.TransactOpts, from, to, value, validAfter, validBefore, nonce, signature)
}

// TransferWithAuthorization0 is a paid mutator transaction binding the contract method 0xe3ee160e.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oftcc *OftccTransactor) TransferWithAuthorization0(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oftcc.contract.Transact(opts, "transferWithAuthorization0", from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// TransferWithAuthorization0 is a paid mutator transaction binding the contract method 0xe3ee160e.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oftcc *OftccSession) TransferWithAuthorization0(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oftcc.Contract.TransferWithAuthorization0(&_Oftcc.TransactOpts, from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// TransferWithAuthorization0 is a paid mutator transaction binding the contract method 0xe3ee160e.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oftcc *OftccTransactorSession) TransferWithAuthorization0(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oftcc.Contract.TransferWithAuthorization0(&_Oftcc.TransactOpts, from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// OftccApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Oftcc contract.
type OftccApprovalIterator struct {
	Event *OftccApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccApproval represents a Approval event raised by the Oftcc contract.
type OftccApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Oftcc *OftccFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OftccApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OftccApprovalIterator{contract: _Oftcc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Oftcc *OftccFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OftccApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccApproval)
				if err := _Oftcc.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Oftcc *OftccFilterer) ParseApproval(log types.Log) (*OftccApproval, error) {
	event := new(OftccApproval)
	if err := _Oftcc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccAuthorizationCanceledIterator is returned from FilterAuthorizationCanceled and is used to iterate over the raw logs and unpacked data for AuthorizationCanceled events raised by the Oftcc contract.
type OftccAuthorizationCanceledIterator struct {
	Event *OftccAuthorizationCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccAuthorizationCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccAuthorizationCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccAuthorizationCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccAuthorizationCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccAuthorizationCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccAuthorizationCanceled represents a AuthorizationCanceled event raised by the Oftcc contract.
type OftccAuthorizationCanceled struct {
	Authorizer common.Address
	Nonce      [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthorizationCanceled is a free log retrieval operation binding the contract event 0x1cdd46ff242716cdaa72d159d339a485b3438398348d68f09d7c8c0a59353d81.
//
// Solidity: event AuthorizationCanceled(address indexed authorizer, bytes32 indexed nonce)
func (_Oftcc *OftccFilterer) FilterAuthorizationCanceled(opts *bind.FilterOpts, authorizer []common.Address, nonce [][32]byte) (*OftccAuthorizationCanceledIterator, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "AuthorizationCanceled", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &OftccAuthorizationCanceledIterator{contract: _Oftcc.contract, event: "AuthorizationCanceled", logs: logs, sub: sub}, nil
}

// WatchAuthorizationCanceled is a free log subscription operation binding the contract event 0x1cdd46ff242716cdaa72d159d339a485b3438398348d68f09d7c8c0a59353d81.
//
// Solidity: event AuthorizationCanceled(address indexed authorizer, bytes32 indexed nonce)
func (_Oftcc *OftccFilterer) WatchAuthorizationCanceled(opts *bind.WatchOpts, sink chan<- *OftccAuthorizationCanceled, authorizer []common.Address, nonce [][32]byte) (event.Subscription, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "AuthorizationCanceled", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccAuthorizationCanceled)
				if err := _Oftcc.contract.UnpackLog(event, "AuthorizationCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizationCanceled is a log parse operation binding the contract event 0x1cdd46ff242716cdaa72d159d339a485b3438398348d68f09d7c8c0a59353d81.
//
// Solidity: event AuthorizationCanceled(address indexed authorizer, bytes32 indexed nonce)
func (_Oftcc *OftccFilterer) ParseAuthorizationCanceled(log types.Log) (*OftccAuthorizationCanceled, error) {
	event := new(OftccAuthorizationCanceled)
	if err := _Oftcc.contract.UnpackLog(event, "AuthorizationCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccAuthorizationUsedIterator is returned from FilterAuthorizationUsed and is used to iterate over the raw logs and unpacked data for AuthorizationUsed events raised by the Oftcc contract.
type OftccAuthorizationUsedIterator struct {
	Event *OftccAuthorizationUsed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccAuthorizationUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccAuthorizationUsed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccAuthorizationUsed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccAuthorizationUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccAuthorizationUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccAuthorizationUsed represents a AuthorizationUsed event raised by the Oftcc contract.
type OftccAuthorizationUsed struct {
	Authorizer common.Address
	Nonce      [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthorizationUsed is a free log retrieval operation binding the contract event 0x98de503528ee59b575ef0c0a2576a82497bfc029a5685b209e9ec333479b10a5.
//
// Solidity: event AuthorizationUsed(address indexed authorizer, bytes32 indexed nonce)
func (_Oftcc *OftccFilterer) FilterAuthorizationUsed(opts *bind.FilterOpts, authorizer []common.Address, nonce [][32]byte) (*OftccAuthorizationUsedIterator, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "AuthorizationUsed", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &OftccAuthorizationUsedIterator{contract: _Oftcc.contract, event: "AuthorizationUsed", logs: logs, sub: sub}, nil
}

// WatchAuthorizationUsed is a free log subscription operation binding the contract event 0x98de503528ee59b575ef0c0a2576a82497bfc029a5685b209e9ec333479b10a5.
//
// Solidity: event AuthorizationUsed(address indexed authorizer, bytes32 indexed nonce)
func (_Oftcc *OftccFilterer) WatchAuthorizationUsed(opts *bind.WatchOpts, sink chan<- *OftccAuthorizationUsed, authorizer []common.Address, nonce [][32]byte) (event.Subscription, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "AuthorizationUsed", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccAuthorizationUsed)
				if err := _Oftcc.contract.UnpackLog(event, "AuthorizationUsed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizationUsed is a log parse operation binding the contract event 0x98de503528ee59b575ef0c0a2576a82497bfc029a5685b209e9ec333479b10a5.
//
// Solidity: event AuthorizationUsed(address indexed authorizer, bytes32 indexed nonce)
func (_Oftcc *OftccFilterer) ParseAuthorizationUsed(log types.Log) (*OftccAuthorizationUsed, error) {
	event := new(OftccAuthorizationUsed)
	if err := _Oftcc.contract.UnpackLog(event, "AuthorizationUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccEnforcedOptionSetIterator is returned from FilterEnforcedOptionSet and is used to iterate over the raw logs and unpacked data for EnforcedOptionSet events raised by the Oftcc contract.
type OftccEnforcedOptionSetIterator struct {
	Event *OftccEnforcedOptionSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccEnforcedOptionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccEnforcedOptionSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccEnforcedOptionSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccEnforcedOptionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccEnforcedOptionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccEnforcedOptionSet represents a EnforcedOptionSet event raised by the Oftcc contract.
type OftccEnforcedOptionSet struct {
	EnforcedOptions []EnforcedOptionParam
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEnforcedOptionSet is a free log retrieval operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Oftcc *OftccFilterer) FilterEnforcedOptionSet(opts *bind.FilterOpts) (*OftccEnforcedOptionSetIterator, error) {

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return &OftccEnforcedOptionSetIterator{contract: _Oftcc.contract, event: "EnforcedOptionSet", logs: logs, sub: sub}, nil
}

// WatchEnforcedOptionSet is a free log subscription operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Oftcc *OftccFilterer) WatchEnforcedOptionSet(opts *bind.WatchOpts, sink chan<- *OftccEnforcedOptionSet) (event.Subscription, error) {

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccEnforcedOptionSet)
				if err := _Oftcc.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnforcedOptionSet is a log parse operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Oftcc *OftccFilterer) ParseEnforcedOptionSet(log types.Log) (*OftccEnforcedOptionSet, error) {
	event := new(OftccEnforcedOptionSet)
	if err := _Oftcc.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccMsgInspectorSetIterator is returned from FilterMsgInspectorSet and is used to iterate over the raw logs and unpacked data for MsgInspectorSet events raised by the Oftcc contract.
type OftccMsgInspectorSetIterator struct {
	Event *OftccMsgInspectorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccMsgInspectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccMsgInspectorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccMsgInspectorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccMsgInspectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccMsgInspectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccMsgInspectorSet represents a MsgInspectorSet event raised by the Oftcc contract.
type OftccMsgInspectorSet struct {
	Inspector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMsgInspectorSet is a free log retrieval operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Oftcc *OftccFilterer) FilterMsgInspectorSet(opts *bind.FilterOpts) (*OftccMsgInspectorSetIterator, error) {

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return &OftccMsgInspectorSetIterator{contract: _Oftcc.contract, event: "MsgInspectorSet", logs: logs, sub: sub}, nil
}

// WatchMsgInspectorSet is a free log subscription operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Oftcc *OftccFilterer) WatchMsgInspectorSet(opts *bind.WatchOpts, sink chan<- *OftccMsgInspectorSet) (event.Subscription, error) {

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccMsgInspectorSet)
				if err := _Oftcc.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMsgInspectorSet is a log parse operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Oftcc *OftccFilterer) ParseMsgInspectorSet(log types.Log) (*OftccMsgInspectorSet, error) {
	event := new(OftccMsgInspectorSet)
	if err := _Oftcc.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccOFTReceivedIterator is returned from FilterOFTReceived and is used to iterate over the raw logs and unpacked data for OFTReceived events raised by the Oftcc contract.
type OftccOFTReceivedIterator struct {
	Event *OftccOFTReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccOFTReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccOFTReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccOFTReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccOFTReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccOFTReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccOFTReceived represents a OFTReceived event raised by the Oftcc contract.
type OftccOFTReceived struct {
	Guid             [32]byte
	SrcEid           uint32
	ToAddress        common.Address
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTReceived is a free log retrieval operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Oftcc *OftccFilterer) FilterOFTReceived(opts *bind.FilterOpts, guid [][32]byte, toAddress []common.Address) (*OftccOFTReceivedIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &OftccOFTReceivedIterator{contract: _Oftcc.contract, event: "OFTReceived", logs: logs, sub: sub}, nil
}

// WatchOFTReceived is a free log subscription operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Oftcc *OftccFilterer) WatchOFTReceived(opts *bind.WatchOpts, sink chan<- *OftccOFTReceived, guid [][32]byte, toAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccOFTReceived)
				if err := _Oftcc.contract.UnpackLog(event, "OFTReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOFTReceived is a log parse operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Oftcc *OftccFilterer) ParseOFTReceived(log types.Log) (*OftccOFTReceived, error) {
	event := new(OftccOFTReceived)
	if err := _Oftcc.contract.UnpackLog(event, "OFTReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccOFTSentIterator is returned from FilterOFTSent and is used to iterate over the raw logs and unpacked data for OFTSent events raised by the Oftcc contract.
type OftccOFTSentIterator struct {
	Event *OftccOFTSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccOFTSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccOFTSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccOFTSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccOFTSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccOFTSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccOFTSent represents a OFTSent event raised by the Oftcc contract.
type OftccOFTSent struct {
	Guid             [32]byte
	DstEid           uint32
	FromAddress      common.Address
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTSent is a free log retrieval operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Oftcc *OftccFilterer) FilterOFTSent(opts *bind.FilterOpts, guid [][32]byte, fromAddress []common.Address) (*OftccOFTSentIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &OftccOFTSentIterator{contract: _Oftcc.contract, event: "OFTSent", logs: logs, sub: sub}, nil
}

// WatchOFTSent is a free log subscription operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Oftcc *OftccFilterer) WatchOFTSent(opts *bind.WatchOpts, sink chan<- *OftccOFTSent, guid [][32]byte, fromAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccOFTSent)
				if err := _Oftcc.contract.UnpackLog(event, "OFTSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOFTSent is a log parse operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Oftcc *OftccFilterer) ParseOFTSent(log types.Log) (*OftccOFTSent, error) {
	event := new(OftccOFTSent)
	if err := _Oftcc.contract.UnpackLog(event, "OFTSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oftcc contract.
type OftccOwnershipTransferredIterator struct {
	Event *OftccOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccOwnershipTransferred represents a OwnershipTransferred event raised by the Oftcc contract.
type OftccOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oftcc *OftccFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OftccOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OftccOwnershipTransferredIterator{contract: _Oftcc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oftcc *OftccFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OftccOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccOwnershipTransferred)
				if err := _Oftcc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oftcc *OftccFilterer) ParseOwnershipTransferred(log types.Log) (*OftccOwnershipTransferred, error) {
	event := new(OftccOwnershipTransferred)
	if err := _Oftcc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the Oftcc contract.
type OftccPeerSetIterator struct {
	Event *OftccPeerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccPeerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccPeerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccPeerSet represents a PeerSet event raised by the Oftcc contract.
type OftccPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Oftcc *OftccFilterer) FilterPeerSet(opts *bind.FilterOpts) (*OftccPeerSetIterator, error) {

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &OftccPeerSetIterator{contract: _Oftcc.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Oftcc *OftccFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *OftccPeerSet) (event.Subscription, error) {

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccPeerSet)
				if err := _Oftcc.contract.UnpackLog(event, "PeerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Oftcc *OftccFilterer) ParsePeerSet(log types.Log) (*OftccPeerSet, error) {
	event := new(OftccPeerSet)
	if err := _Oftcc.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccPreCrimeSetIterator is returned from FilterPreCrimeSet and is used to iterate over the raw logs and unpacked data for PreCrimeSet events raised by the Oftcc contract.
type OftccPreCrimeSetIterator struct {
	Event *OftccPreCrimeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccPreCrimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccPreCrimeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccPreCrimeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccPreCrimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccPreCrimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccPreCrimeSet represents a PreCrimeSet event raised by the Oftcc contract.
type OftccPreCrimeSet struct {
	PreCrimeAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPreCrimeSet is a free log retrieval operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Oftcc *OftccFilterer) FilterPreCrimeSet(opts *bind.FilterOpts) (*OftccPreCrimeSetIterator, error) {

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return &OftccPreCrimeSetIterator{contract: _Oftcc.contract, event: "PreCrimeSet", logs: logs, sub: sub}, nil
}

// WatchPreCrimeSet is a free log subscription operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Oftcc *OftccFilterer) WatchPreCrimeSet(opts *bind.WatchOpts, sink chan<- *OftccPreCrimeSet) (event.Subscription, error) {

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccPreCrimeSet)
				if err := _Oftcc.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreCrimeSet is a log parse operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Oftcc *OftccFilterer) ParsePreCrimeSet(log types.Log) (*OftccPreCrimeSet, error) {
	event := new(OftccPreCrimeSet)
	if err := _Oftcc.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftccTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Oftcc contract.
type OftccTransferIterator struct {
	Event *OftccTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftccTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftccTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftccTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftccTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftccTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftccTransfer represents a Transfer event raised by the Oftcc contract.
type OftccTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Oftcc *OftccFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OftccTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Oftcc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OftccTransferIterator{contract: _Oftcc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Oftcc *OftccFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OftccTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Oftcc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftccTransfer)
				if err := _Oftcc.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Oftcc *OftccFilterer) ParseTransfer(log types.Log) (*OftccTransfer, error) {
	event := new(OftccTransfer)
	if err := _Oftcc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
