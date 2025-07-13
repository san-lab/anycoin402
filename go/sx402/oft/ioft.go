// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oft

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

// OftMetaData contains all meta data concerning the Oft contract.
var OftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLocalDecimals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"InvalidOptions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"SimulationResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"}],\"name\":\"SlippageExceeded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"AuthorizationCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"AuthorizationUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"EnforcedOptionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inspector\",\"type\":\"address\"}],\"name\":\"MsgInspectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"preCrimeAddress\",\"type\":\"address\"}],\"name\":\"PreCrimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANCEL_AUTHORIZATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEND\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEND_AND_CALL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_WITH_AUTHORIZATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvalRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"authorizationState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancelAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"cancelAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_extraOptions\",\"type\":\"bytes\"}],\"name\":\"combineOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalConversionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"}],\"name\":\"enforcedOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"enforcedOption\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"isPeer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structInboundPacket[]\",\"name\":\"_packets\",\"type\":\"tuple[]\"}],\"name\":\"lzReceiveAndRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceiveSimulate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgInspector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oftVersion\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preCrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"}],\"name\":\"quoteOFT\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTLimit\",\"name\":\"oftLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"feeAmountLD\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structOFTFeeDetail[]\",\"name\":\"oftFeeDetails\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_payInLzToken\",\"type\":\"bool\"}],\"name\":\"quoteSend\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"msgFee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"msgReceipt\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendWithAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"setEnforcedOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgInspector\",\"type\":\"address\"}],\"name\":\"setMsgInspector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_preCrime\",\"type\":\"address\"}],\"name\":\"setPreCrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"transferWithAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"transferWithAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OftABI is the input ABI used to generate the binding from.
// Deprecated: Use OftMetaData.ABI instead.
var OftABI = OftMetaData.ABI

// Oft is an auto generated Go binding around an Ethereum contract.
type Oft struct {
	OftCaller     // Read-only binding to the contract
	OftTransactor // Write-only binding to the contract
	OftFilterer   // Log filterer for contract events
}

// OftCaller is an auto generated read-only Go binding around an Ethereum contract.
type OftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OftSession struct {
	Contract     *Oft              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OftCallerSession struct {
	Contract *OftCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OftTransactorSession struct {
	Contract     *OftTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OftRaw is an auto generated low-level Go binding around an Ethereum contract.
type OftRaw struct {
	Contract *Oft // Generic contract binding to access the raw methods on
}

// OftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OftCallerRaw struct {
	Contract *OftCaller // Generic read-only contract binding to access the raw methods on
}

// OftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OftTransactorRaw struct {
	Contract *OftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOft creates a new instance of Oft, bound to a specific deployed contract.
func NewOft(address common.Address, backend bind.ContractBackend) (*Oft, error) {
	contract, err := bindOft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oft{OftCaller: OftCaller{contract: contract}, OftTransactor: OftTransactor{contract: contract}, OftFilterer: OftFilterer{contract: contract}}, nil
}

// NewOftCaller creates a new read-only instance of Oft, bound to a specific deployed contract.
func NewOftCaller(address common.Address, caller bind.ContractCaller) (*OftCaller, error) {
	contract, err := bindOft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OftCaller{contract: contract}, nil
}

// NewOftTransactor creates a new write-only instance of Oft, bound to a specific deployed contract.
func NewOftTransactor(address common.Address, transactor bind.ContractTransactor) (*OftTransactor, error) {
	contract, err := bindOft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OftTransactor{contract: contract}, nil
}

// NewOftFilterer creates a new log filterer instance of Oft, bound to a specific deployed contract.
func NewOftFilterer(address common.Address, filterer bind.ContractFilterer) (*OftFilterer, error) {
	contract, err := bindOft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OftFilterer{contract: contract}, nil
}

// bindOft binds a generic wrapper to an already deployed contract.
func bindOft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oft *OftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oft.Contract.OftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oft *OftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oft.Contract.OftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oft *OftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oft.Contract.OftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oft *OftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oft *OftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oft *OftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oft.Contract.contract.Transact(opts, method, params...)
}

// CANCELAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xd9169487.
//
// Solidity: function CANCEL_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oft *OftCaller) CANCELAUTHORIZATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "CANCEL_AUTHORIZATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xd9169487.
//
// Solidity: function CANCEL_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oft *OftSession) CANCELAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Oft.Contract.CANCELAUTHORIZATIONTYPEHASH(&_Oft.CallOpts)
}

// CANCELAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xd9169487.
//
// Solidity: function CANCEL_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oft *OftCallerSession) CANCELAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Oft.Contract.CANCELAUTHORIZATIONTYPEHASH(&_Oft.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Oft *OftCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Oft *OftSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Oft.Contract.DOMAINSEPARATOR(&_Oft.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Oft *OftCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Oft.Contract.DOMAINSEPARATOR(&_Oft.CallOpts)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Oft *OftCaller) SEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Oft *OftSession) SEND() (uint16, error) {
	return _Oft.Contract.SEND(&_Oft.CallOpts)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Oft *OftCallerSession) SEND() (uint16, error) {
	return _Oft.Contract.SEND(&_Oft.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Oft *OftCaller) SENDANDCALL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "SEND_AND_CALL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Oft *OftSession) SENDANDCALL() (uint16, error) {
	return _Oft.Contract.SENDANDCALL(&_Oft.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Oft *OftCallerSession) SENDANDCALL() (uint16, error) {
	return _Oft.Contract.SENDANDCALL(&_Oft.CallOpts)
}

// TRANSFERWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xa0cc6a68.
//
// Solidity: function TRANSFER_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oft *OftCaller) TRANSFERWITHAUTHORIZATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "TRANSFER_WITH_AUTHORIZATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xa0cc6a68.
//
// Solidity: function TRANSFER_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oft *OftSession) TRANSFERWITHAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Oft.Contract.TRANSFERWITHAUTHORIZATIONTYPEHASH(&_Oft.CallOpts)
}

// TRANSFERWITHAUTHORIZATIONTYPEHASH is a free data retrieval call binding the contract method 0xa0cc6a68.
//
// Solidity: function TRANSFER_WITH_AUTHORIZATION_TYPEHASH() view returns(bytes32)
func (_Oft *OftCallerSession) TRANSFERWITHAUTHORIZATIONTYPEHASH() ([32]byte, error) {
	return _Oft.Contract.TRANSFERWITHAUTHORIZATIONTYPEHASH(&_Oft.CallOpts)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Oft *OftCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Oft *OftSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Oft.Contract.AllowInitializePath(&_Oft.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Oft *OftCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Oft.Contract.AllowInitializePath(&_Oft.CallOpts, origin)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Oft *OftCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Oft *OftSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Oft.Contract.Allowance(&_Oft.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Oft *OftCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Oft.Contract.Allowance(&_Oft.CallOpts, owner, spender)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Oft *OftCaller) ApprovalRequired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "approvalRequired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Oft *OftSession) ApprovalRequired() (bool, error) {
	return _Oft.Contract.ApprovalRequired(&_Oft.CallOpts)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Oft *OftCallerSession) ApprovalRequired() (bool, error) {
	return _Oft.Contract.ApprovalRequired(&_Oft.CallOpts)
}

// AuthorizationState is a free data retrieval call binding the contract method 0xe94a0102.
//
// Solidity: function authorizationState(address authorizer, bytes32 nonce) view returns(bool)
func (_Oft *OftCaller) AuthorizationState(opts *bind.CallOpts, authorizer common.Address, nonce [32]byte) (bool, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "authorizationState", authorizer, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizationState is a free data retrieval call binding the contract method 0xe94a0102.
//
// Solidity: function authorizationState(address authorizer, bytes32 nonce) view returns(bool)
func (_Oft *OftSession) AuthorizationState(authorizer common.Address, nonce [32]byte) (bool, error) {
	return _Oft.Contract.AuthorizationState(&_Oft.CallOpts, authorizer, nonce)
}

// AuthorizationState is a free data retrieval call binding the contract method 0xe94a0102.
//
// Solidity: function authorizationState(address authorizer, bytes32 nonce) view returns(bool)
func (_Oft *OftCallerSession) AuthorizationState(authorizer common.Address, nonce [32]byte) (bool, error) {
	return _Oft.Contract.AuthorizationState(&_Oft.CallOpts, authorizer, nonce)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Oft *OftCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Oft *OftSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Oft.Contract.BalanceOf(&_Oft.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Oft *OftCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Oft.Contract.BalanceOf(&_Oft.CallOpts, account)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Oft *OftCaller) CombineOptions(opts *bind.CallOpts, _eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "combineOptions", _eid, _msgType, _extraOptions)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Oft *OftSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _Oft.Contract.CombineOptions(&_Oft.CallOpts, _eid, _msgType, _extraOptions)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Oft *OftCallerSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _Oft.Contract.CombineOptions(&_Oft.CallOpts, _eid, _msgType, _extraOptions)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Oft *OftCaller) DecimalConversionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "decimalConversionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Oft *OftSession) DecimalConversionRate() (*big.Int, error) {
	return _Oft.Contract.DecimalConversionRate(&_Oft.CallOpts)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Oft *OftCallerSession) DecimalConversionRate() (*big.Int, error) {
	return _Oft.Contract.DecimalConversionRate(&_Oft.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Oft *OftCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Oft *OftSession) Decimals() (uint8, error) {
	return _Oft.Contract.Decimals(&_Oft.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Oft *OftCallerSession) Decimals() (uint8, error) {
	return _Oft.Contract.Decimals(&_Oft.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Oft *OftCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Oft *OftSession) Endpoint() (common.Address, error) {
	return _Oft.Contract.Endpoint(&_Oft.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Oft *OftCallerSession) Endpoint() (common.Address, error) {
	return _Oft.Contract.Endpoint(&_Oft.CallOpts)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_Oft *OftCaller) EnforcedOptions(opts *bind.CallOpts, eid uint32, msgType uint16) ([]byte, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "enforcedOptions", eid, msgType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_Oft *OftSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _Oft.Contract.EnforcedOptions(&_Oft.CallOpts, eid, msgType)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_Oft *OftCallerSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _Oft.Contract.EnforcedOptions(&_Oft.CallOpts, eid, msgType)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Oft *OftCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Oft *OftSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Oft.Contract.IsComposeMsgSender(&_Oft.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Oft *OftCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Oft.Contract.IsComposeMsgSender(&_Oft.CallOpts, arg0, arg1, _sender)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Oft *OftCaller) IsPeer(opts *bind.CallOpts, _eid uint32, _peer [32]byte) (bool, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "isPeer", _eid, _peer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Oft *OftSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _Oft.Contract.IsPeer(&_Oft.CallOpts, _eid, _peer)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Oft *OftCallerSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _Oft.Contract.IsPeer(&_Oft.CallOpts, _eid, _peer)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Oft *OftCaller) MsgInspector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "msgInspector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Oft *OftSession) MsgInspector() (common.Address, error) {
	return _Oft.Contract.MsgInspector(&_Oft.CallOpts)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Oft *OftCallerSession) MsgInspector() (common.Address, error) {
	return _Oft.Contract.MsgInspector(&_Oft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Oft *OftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Oft *OftSession) Name() (string, error) {
	return _Oft.Contract.Name(&_Oft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Oft *OftCallerSession) Name() (string, error) {
	return _Oft.Contract.Name(&_Oft.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Oft *OftCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Oft *OftSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Oft.Contract.NextNonce(&_Oft.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Oft *OftCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Oft.Contract.NextNonce(&_Oft.CallOpts, arg0, arg1)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Oft *OftCaller) OApp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "oApp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Oft *OftSession) OApp() (common.Address, error) {
	return _Oft.Contract.OApp(&_Oft.CallOpts)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Oft *OftCallerSession) OApp() (common.Address, error) {
	return _Oft.Contract.OApp(&_Oft.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Oft *OftCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "oAppVersion")

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
func (_Oft *OftSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Oft.Contract.OAppVersion(&_Oft.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Oft *OftCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Oft.Contract.OAppVersion(&_Oft.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Oft *OftCaller) OftVersion(opts *bind.CallOpts) (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "oftVersion")

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
func (_Oft *OftSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _Oft.Contract.OftVersion(&_Oft.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Oft *OftCallerSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _Oft.Contract.OftVersion(&_Oft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oft *OftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oft *OftSession) Owner() (common.Address, error) {
	return _Oft.Contract.Owner(&_Oft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oft *OftCallerSession) Owner() (common.Address, error) {
	return _Oft.Contract.Owner(&_Oft.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Oft *OftCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Oft *OftSession) Peers(eid uint32) ([32]byte, error) {
	return _Oft.Contract.Peers(&_Oft.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Oft *OftCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _Oft.Contract.Peers(&_Oft.CallOpts, eid)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Oft *OftCaller) PreCrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "preCrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Oft *OftSession) PreCrime() (common.Address, error) {
	return _Oft.Contract.PreCrime(&_Oft.CallOpts)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Oft *OftCallerSession) PreCrime() (common.Address, error) {
	return _Oft.Contract.PreCrime(&_Oft.CallOpts)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Oft *OftCaller) QuoteOFT(opts *bind.CallOpts, _sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "quoteOFT", _sendParam)

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
func (_Oft *OftSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _Oft.Contract.QuoteOFT(&_Oft.CallOpts, _sendParam)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Oft *OftCallerSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _Oft.Contract.QuoteOFT(&_Oft.CallOpts, _sendParam)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Oft *OftCaller) QuoteSend(opts *bind.CallOpts, _sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "quoteSend", _sendParam, _payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Oft *OftSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Oft.Contract.QuoteSend(&_Oft.CallOpts, _sendParam, _payInLzToken)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Oft *OftCallerSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Oft.Contract.QuoteSend(&_Oft.CallOpts, _sendParam, _payInLzToken)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_Oft *OftCaller) SharedDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "sharedDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_Oft *OftSession) SharedDecimals() (uint8, error) {
	return _Oft.Contract.SharedDecimals(&_Oft.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_Oft *OftCallerSession) SharedDecimals() (uint8, error) {
	return _Oft.Contract.SharedDecimals(&_Oft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Oft *OftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Oft *OftSession) Symbol() (string, error) {
	return _Oft.Contract.Symbol(&_Oft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Oft *OftCallerSession) Symbol() (string, error) {
	return _Oft.Contract.Symbol(&_Oft.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Oft *OftCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Oft *OftSession) Token() (common.Address, error) {
	return _Oft.Contract.Token(&_Oft.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Oft *OftCallerSession) Token() (common.Address, error) {
	return _Oft.Contract.Token(&_Oft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Oft *OftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Oft *OftSession) TotalSupply() (*big.Int, error) {
	return _Oft.Contract.TotalSupply(&_Oft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Oft *OftCallerSession) TotalSupply() (*big.Int, error) {
	return _Oft.Contract.TotalSupply(&_Oft.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Oft *OftCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oft.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Oft *OftSession) Version() (string, error) {
	return _Oft.Contract.Version(&_Oft.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Oft *OftCallerSession) Version() (string, error) {
	return _Oft.Contract.Version(&_Oft.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Oft *OftTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Oft *OftSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oft.Contract.Approve(&_Oft.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Oft *OftTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oft.Contract.Approve(&_Oft.TransactOpts, spender, value)
}

// CancelAuthorization is a paid mutator transaction binding the contract method 0x5a049a70.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oft *OftTransactor) CancelAuthorization(opts *bind.TransactOpts, authorizer common.Address, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "cancelAuthorization", authorizer, nonce, v, r, s)
}

// CancelAuthorization is a paid mutator transaction binding the contract method 0x5a049a70.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oft *OftSession) CancelAuthorization(authorizer common.Address, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oft.Contract.CancelAuthorization(&_Oft.TransactOpts, authorizer, nonce, v, r, s)
}

// CancelAuthorization is a paid mutator transaction binding the contract method 0x5a049a70.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oft *OftTransactorSession) CancelAuthorization(authorizer common.Address, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oft.Contract.CancelAuthorization(&_Oft.TransactOpts, authorizer, nonce, v, r, s)
}

// CancelAuthorization0 is a paid mutator transaction binding the contract method 0xb7b72899.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, bytes signature) returns()
func (_Oft *OftTransactor) CancelAuthorization0(opts *bind.TransactOpts, authorizer common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "cancelAuthorization0", authorizer, nonce, signature)
}

// CancelAuthorization0 is a paid mutator transaction binding the contract method 0xb7b72899.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, bytes signature) returns()
func (_Oft *OftSession) CancelAuthorization0(authorizer common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oft.Contract.CancelAuthorization0(&_Oft.TransactOpts, authorizer, nonce, signature)
}

// CancelAuthorization0 is a paid mutator transaction binding the contract method 0xb7b72899.
//
// Solidity: function cancelAuthorization(address authorizer, bytes32 nonce, bytes signature) returns()
func (_Oft *OftTransactorSession) CancelAuthorization0(authorizer common.Address, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oft.Contract.CancelAuthorization0(&_Oft.TransactOpts, authorizer, nonce, signature)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oft *OftTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oft *OftSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oft.Contract.LzReceive(&_Oft.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oft *OftTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oft.Contract.LzReceive(&_Oft.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Oft *OftTransactor) LzReceiveAndRevert(opts *bind.TransactOpts, _packets []InboundPacket) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "lzReceiveAndRevert", _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Oft *OftSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _Oft.Contract.LzReceiveAndRevert(&_Oft.TransactOpts, _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Oft *OftTransactorSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _Oft.Contract.LzReceiveAndRevert(&_Oft.TransactOpts, _packets)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oft *OftTransactor) LzReceiveSimulate(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "lzReceiveSimulate", _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oft *OftSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oft.Contract.LzReceiveSimulate(&_Oft.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Oft *OftTransactorSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Oft.Contract.LzReceiveSimulate(&_Oft.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Oft *OftTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Oft *OftSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oft.Contract.Mint(&_Oft.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Oft *OftTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oft.Contract.Mint(&_Oft.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oft *OftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oft *OftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oft.Contract.RenounceOwnership(&_Oft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oft *OftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oft.Contract.RenounceOwnership(&_Oft.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Oft *OftTransactor) Send(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "send", _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Oft *OftSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oft.Contract.Send(&_Oft.TransactOpts, _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Oft *OftTransactorSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oft.Contract.Send(&_Oft.TransactOpts, _sendParam, _fee, _refundAddress)
}

// SendFrom is a paid mutator transaction binding the contract method 0xa6e329f5.
//
// Solidity: function sendFrom((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, address _refundAddress) returns(bool)
func (_Oft *OftTransactor) SendFrom(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _from common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "sendFrom", _sendParam, _fee, _from, _refundAddress)
}

// SendFrom is a paid mutator transaction binding the contract method 0xa6e329f5.
//
// Solidity: function sendFrom((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, address _refundAddress) returns(bool)
func (_Oft *OftSession) SendFrom(_sendParam SendParam, _fee MessagingFee, _from common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SendFrom(&_Oft.TransactOpts, _sendParam, _fee, _from, _refundAddress)
}

// SendFrom is a paid mutator transaction binding the contract method 0xa6e329f5.
//
// Solidity: function sendFrom((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, address _refundAddress) returns(bool)
func (_Oft *OftTransactorSession) SendFrom(_sendParam SendParam, _fee MessagingFee, _from common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SendFrom(&_Oft.TransactOpts, _sendParam, _fee, _from, _refundAddress)
}

// SendWithAuthorization is a paid mutator transaction binding the contract method 0xeb3ebb18.
//
// Solidity: function sendWithAuthorization((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature, address _refundAddress) returns(bool)
func (_Oft *OftTransactor) SendWithAuthorization(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _from common.Address, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "sendWithAuthorization", _sendParam, _fee, _from, validAfter, validBefore, nonce, signature, _refundAddress)
}

// SendWithAuthorization is a paid mutator transaction binding the contract method 0xeb3ebb18.
//
// Solidity: function sendWithAuthorization((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature, address _refundAddress) returns(bool)
func (_Oft *OftSession) SendWithAuthorization(_sendParam SendParam, _fee MessagingFee, _from common.Address, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SendWithAuthorization(&_Oft.TransactOpts, _sendParam, _fee, _from, validAfter, validBefore, nonce, signature, _refundAddress)
}

// SendWithAuthorization is a paid mutator transaction binding the contract method 0xeb3ebb18.
//
// Solidity: function sendWithAuthorization((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _from, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature, address _refundAddress) returns(bool)
func (_Oft *OftTransactorSession) SendWithAuthorization(_sendParam SendParam, _fee MessagingFee, _from common.Address, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SendWithAuthorization(&_Oft.TransactOpts, _sendParam, _fee, _from, validAfter, validBefore, nonce, signature, _refundAddress)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Oft *OftTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Oft *OftSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SetDelegate(&_Oft.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Oft *OftTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SetDelegate(&_Oft.TransactOpts, _delegate)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Oft *OftTransactor) SetEnforcedOptions(opts *bind.TransactOpts, _enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "setEnforcedOptions", _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Oft *OftSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Oft.Contract.SetEnforcedOptions(&_Oft.TransactOpts, _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Oft *OftTransactorSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Oft.Contract.SetEnforcedOptions(&_Oft.TransactOpts, _enforcedOptions)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Oft *OftTransactor) SetMsgInspector(opts *bind.TransactOpts, _msgInspector common.Address) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "setMsgInspector", _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Oft *OftSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SetMsgInspector(&_Oft.TransactOpts, _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Oft *OftTransactorSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SetMsgInspector(&_Oft.TransactOpts, _msgInspector)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Oft *OftTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Oft *OftSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Oft.Contract.SetPeer(&_Oft.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Oft *OftTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Oft.Contract.SetPeer(&_Oft.TransactOpts, _eid, _peer)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Oft *OftTransactor) SetPreCrime(opts *bind.TransactOpts, _preCrime common.Address) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "setPreCrime", _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Oft *OftSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SetPreCrime(&_Oft.TransactOpts, _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Oft *OftTransactorSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _Oft.Contract.SetPreCrime(&_Oft.TransactOpts, _preCrime)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Oft *OftTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Oft *OftSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oft.Contract.Transfer(&_Oft.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Oft *OftTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oft.Contract.Transfer(&_Oft.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Oft *OftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Oft *OftSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oft.Contract.TransferFrom(&_Oft.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Oft *OftTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Oft.Contract.TransferFrom(&_Oft.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oft *OftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oft *OftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oft.Contract.TransferOwnership(&_Oft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oft *OftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oft.Contract.TransferOwnership(&_Oft.TransactOpts, newOwner)
}

// TransferWithAuthorization is a paid mutator transaction binding the contract method 0xcf092995.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature) returns()
func (_Oft *OftTransactor) TransferWithAuthorization(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "transferWithAuthorization", from, to, value, validAfter, validBefore, nonce, signature)
}

// TransferWithAuthorization is a paid mutator transaction binding the contract method 0xcf092995.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature) returns()
func (_Oft *OftSession) TransferWithAuthorization(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oft.Contract.TransferWithAuthorization(&_Oft.TransactOpts, from, to, value, validAfter, validBefore, nonce, signature)
}

// TransferWithAuthorization is a paid mutator transaction binding the contract method 0xcf092995.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature) returns()
func (_Oft *OftTransactorSession) TransferWithAuthorization(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _Oft.Contract.TransferWithAuthorization(&_Oft.TransactOpts, from, to, value, validAfter, validBefore, nonce, signature)
}

// TransferWithAuthorization0 is a paid mutator transaction binding the contract method 0xe3ee160e.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oft *OftTransactor) TransferWithAuthorization0(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oft.contract.Transact(opts, "transferWithAuthorization0", from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// TransferWithAuthorization0 is a paid mutator transaction binding the contract method 0xe3ee160e.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oft *OftSession) TransferWithAuthorization0(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oft.Contract.TransferWithAuthorization0(&_Oft.TransactOpts, from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// TransferWithAuthorization0 is a paid mutator transaction binding the contract method 0xe3ee160e.
//
// Solidity: function transferWithAuthorization(address from, address to, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Oft *OftTransactorSession) TransferWithAuthorization0(from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Oft.Contract.TransferWithAuthorization0(&_Oft.TransactOpts, from, to, value, validAfter, validBefore, nonce, v, r, s)
}

// OftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Oft contract.
type OftApprovalIterator struct {
	Event *OftApproval // Event containing the contract specifics and raw log

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
func (it *OftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftApproval)
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
		it.Event = new(OftApproval)
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
func (it *OftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftApproval represents a Approval event raised by the Oft contract.
type OftApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Oft *OftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Oft.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OftApprovalIterator{contract: _Oft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Oft *OftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OftApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Oft.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftApproval)
				if err := _Oft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Oft *OftFilterer) ParseApproval(log types.Log) (*OftApproval, error) {
	event := new(OftApproval)
	if err := _Oft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftAuthorizationCanceledIterator is returned from FilterAuthorizationCanceled and is used to iterate over the raw logs and unpacked data for AuthorizationCanceled events raised by the Oft contract.
type OftAuthorizationCanceledIterator struct {
	Event *OftAuthorizationCanceled // Event containing the contract specifics and raw log

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
func (it *OftAuthorizationCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftAuthorizationCanceled)
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
		it.Event = new(OftAuthorizationCanceled)
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
func (it *OftAuthorizationCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftAuthorizationCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftAuthorizationCanceled represents a AuthorizationCanceled event raised by the Oft contract.
type OftAuthorizationCanceled struct {
	Authorizer common.Address
	Nonce      [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthorizationCanceled is a free log retrieval operation binding the contract event 0x1cdd46ff242716cdaa72d159d339a485b3438398348d68f09d7c8c0a59353d81.
//
// Solidity: event AuthorizationCanceled(address indexed authorizer, bytes32 indexed nonce)
func (_Oft *OftFilterer) FilterAuthorizationCanceled(opts *bind.FilterOpts, authorizer []common.Address, nonce [][32]byte) (*OftAuthorizationCanceledIterator, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Oft.contract.FilterLogs(opts, "AuthorizationCanceled", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &OftAuthorizationCanceledIterator{contract: _Oft.contract, event: "AuthorizationCanceled", logs: logs, sub: sub}, nil
}

// WatchAuthorizationCanceled is a free log subscription operation binding the contract event 0x1cdd46ff242716cdaa72d159d339a485b3438398348d68f09d7c8c0a59353d81.
//
// Solidity: event AuthorizationCanceled(address indexed authorizer, bytes32 indexed nonce)
func (_Oft *OftFilterer) WatchAuthorizationCanceled(opts *bind.WatchOpts, sink chan<- *OftAuthorizationCanceled, authorizer []common.Address, nonce [][32]byte) (event.Subscription, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Oft.contract.WatchLogs(opts, "AuthorizationCanceled", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftAuthorizationCanceled)
				if err := _Oft.contract.UnpackLog(event, "AuthorizationCanceled", log); err != nil {
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
func (_Oft *OftFilterer) ParseAuthorizationCanceled(log types.Log) (*OftAuthorizationCanceled, error) {
	event := new(OftAuthorizationCanceled)
	if err := _Oft.contract.UnpackLog(event, "AuthorizationCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftAuthorizationUsedIterator is returned from FilterAuthorizationUsed and is used to iterate over the raw logs and unpacked data for AuthorizationUsed events raised by the Oft contract.
type OftAuthorizationUsedIterator struct {
	Event *OftAuthorizationUsed // Event containing the contract specifics and raw log

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
func (it *OftAuthorizationUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftAuthorizationUsed)
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
		it.Event = new(OftAuthorizationUsed)
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
func (it *OftAuthorizationUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftAuthorizationUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftAuthorizationUsed represents a AuthorizationUsed event raised by the Oft contract.
type OftAuthorizationUsed struct {
	Authorizer common.Address
	Nonce      [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthorizationUsed is a free log retrieval operation binding the contract event 0x98de503528ee59b575ef0c0a2576a82497bfc029a5685b209e9ec333479b10a5.
//
// Solidity: event AuthorizationUsed(address indexed authorizer, bytes32 indexed nonce)
func (_Oft *OftFilterer) FilterAuthorizationUsed(opts *bind.FilterOpts, authorizer []common.Address, nonce [][32]byte) (*OftAuthorizationUsedIterator, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Oft.contract.FilterLogs(opts, "AuthorizationUsed", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &OftAuthorizationUsedIterator{contract: _Oft.contract, event: "AuthorizationUsed", logs: logs, sub: sub}, nil
}

// WatchAuthorizationUsed is a free log subscription operation binding the contract event 0x98de503528ee59b575ef0c0a2576a82497bfc029a5685b209e9ec333479b10a5.
//
// Solidity: event AuthorizationUsed(address indexed authorizer, bytes32 indexed nonce)
func (_Oft *OftFilterer) WatchAuthorizationUsed(opts *bind.WatchOpts, sink chan<- *OftAuthorizationUsed, authorizer []common.Address, nonce [][32]byte) (event.Subscription, error) {

	var authorizerRule []interface{}
	for _, authorizerItem := range authorizer {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Oft.contract.WatchLogs(opts, "AuthorizationUsed", authorizerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftAuthorizationUsed)
				if err := _Oft.contract.UnpackLog(event, "AuthorizationUsed", log); err != nil {
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
func (_Oft *OftFilterer) ParseAuthorizationUsed(log types.Log) (*OftAuthorizationUsed, error) {
	event := new(OftAuthorizationUsed)
	if err := _Oft.contract.UnpackLog(event, "AuthorizationUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftEnforcedOptionSetIterator is returned from FilterEnforcedOptionSet and is used to iterate over the raw logs and unpacked data for EnforcedOptionSet events raised by the Oft contract.
type OftEnforcedOptionSetIterator struct {
	Event *OftEnforcedOptionSet // Event containing the contract specifics and raw log

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
func (it *OftEnforcedOptionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftEnforcedOptionSet)
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
		it.Event = new(OftEnforcedOptionSet)
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
func (it *OftEnforcedOptionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftEnforcedOptionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftEnforcedOptionSet represents a EnforcedOptionSet event raised by the Oft contract.
type OftEnforcedOptionSet struct {
	EnforcedOptions []EnforcedOptionParam
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEnforcedOptionSet is a free log retrieval operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Oft *OftFilterer) FilterEnforcedOptionSet(opts *bind.FilterOpts) (*OftEnforcedOptionSetIterator, error) {

	logs, sub, err := _Oft.contract.FilterLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return &OftEnforcedOptionSetIterator{contract: _Oft.contract, event: "EnforcedOptionSet", logs: logs, sub: sub}, nil
}

// WatchEnforcedOptionSet is a free log subscription operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Oft *OftFilterer) WatchEnforcedOptionSet(opts *bind.WatchOpts, sink chan<- *OftEnforcedOptionSet) (event.Subscription, error) {

	logs, sub, err := _Oft.contract.WatchLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftEnforcedOptionSet)
				if err := _Oft.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
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
func (_Oft *OftFilterer) ParseEnforcedOptionSet(log types.Log) (*OftEnforcedOptionSet, error) {
	event := new(OftEnforcedOptionSet)
	if err := _Oft.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftMsgInspectorSetIterator is returned from FilterMsgInspectorSet and is used to iterate over the raw logs and unpacked data for MsgInspectorSet events raised by the Oft contract.
type OftMsgInspectorSetIterator struct {
	Event *OftMsgInspectorSet // Event containing the contract specifics and raw log

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
func (it *OftMsgInspectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftMsgInspectorSet)
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
		it.Event = new(OftMsgInspectorSet)
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
func (it *OftMsgInspectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftMsgInspectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftMsgInspectorSet represents a MsgInspectorSet event raised by the Oft contract.
type OftMsgInspectorSet struct {
	Inspector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMsgInspectorSet is a free log retrieval operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Oft *OftFilterer) FilterMsgInspectorSet(opts *bind.FilterOpts) (*OftMsgInspectorSetIterator, error) {

	logs, sub, err := _Oft.contract.FilterLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return &OftMsgInspectorSetIterator{contract: _Oft.contract, event: "MsgInspectorSet", logs: logs, sub: sub}, nil
}

// WatchMsgInspectorSet is a free log subscription operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Oft *OftFilterer) WatchMsgInspectorSet(opts *bind.WatchOpts, sink chan<- *OftMsgInspectorSet) (event.Subscription, error) {

	logs, sub, err := _Oft.contract.WatchLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftMsgInspectorSet)
				if err := _Oft.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
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
func (_Oft *OftFilterer) ParseMsgInspectorSet(log types.Log) (*OftMsgInspectorSet, error) {
	event := new(OftMsgInspectorSet)
	if err := _Oft.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftOFTReceivedIterator is returned from FilterOFTReceived and is used to iterate over the raw logs and unpacked data for OFTReceived events raised by the Oft contract.
type OftOFTReceivedIterator struct {
	Event *OftOFTReceived // Event containing the contract specifics and raw log

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
func (it *OftOFTReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftOFTReceived)
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
		it.Event = new(OftOFTReceived)
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
func (it *OftOFTReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftOFTReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftOFTReceived represents a OFTReceived event raised by the Oft contract.
type OftOFTReceived struct {
	Guid             [32]byte
	SrcEid           uint32
	ToAddress        common.Address
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTReceived is a free log retrieval operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Oft *OftFilterer) FilterOFTReceived(opts *bind.FilterOpts, guid [][32]byte, toAddress []common.Address) (*OftOFTReceivedIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Oft.contract.FilterLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &OftOFTReceivedIterator{contract: _Oft.contract, event: "OFTReceived", logs: logs, sub: sub}, nil
}

// WatchOFTReceived is a free log subscription operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Oft *OftFilterer) WatchOFTReceived(opts *bind.WatchOpts, sink chan<- *OftOFTReceived, guid [][32]byte, toAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Oft.contract.WatchLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftOFTReceived)
				if err := _Oft.contract.UnpackLog(event, "OFTReceived", log); err != nil {
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
func (_Oft *OftFilterer) ParseOFTReceived(log types.Log) (*OftOFTReceived, error) {
	event := new(OftOFTReceived)
	if err := _Oft.contract.UnpackLog(event, "OFTReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftOFTSentIterator is returned from FilterOFTSent and is used to iterate over the raw logs and unpacked data for OFTSent events raised by the Oft contract.
type OftOFTSentIterator struct {
	Event *OftOFTSent // Event containing the contract specifics and raw log

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
func (it *OftOFTSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftOFTSent)
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
		it.Event = new(OftOFTSent)
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
func (it *OftOFTSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftOFTSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftOFTSent represents a OFTSent event raised by the Oft contract.
type OftOFTSent struct {
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
func (_Oft *OftFilterer) FilterOFTSent(opts *bind.FilterOpts, guid [][32]byte, fromAddress []common.Address) (*OftOFTSentIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Oft.contract.FilterLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &OftOFTSentIterator{contract: _Oft.contract, event: "OFTSent", logs: logs, sub: sub}, nil
}

// WatchOFTSent is a free log subscription operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Oft *OftFilterer) WatchOFTSent(opts *bind.WatchOpts, sink chan<- *OftOFTSent, guid [][32]byte, fromAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Oft.contract.WatchLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftOFTSent)
				if err := _Oft.contract.UnpackLog(event, "OFTSent", log); err != nil {
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
func (_Oft *OftFilterer) ParseOFTSent(log types.Log) (*OftOFTSent, error) {
	event := new(OftOFTSent)
	if err := _Oft.contract.UnpackLog(event, "OFTSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oft contract.
type OftOwnershipTransferredIterator struct {
	Event *OftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftOwnershipTransferred)
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
		it.Event = new(OftOwnershipTransferred)
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
func (it *OftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftOwnershipTransferred represents a OwnershipTransferred event raised by the Oft contract.
type OftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oft *OftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OftOwnershipTransferredIterator{contract: _Oft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oft *OftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftOwnershipTransferred)
				if err := _Oft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oft *OftFilterer) ParseOwnershipTransferred(log types.Log) (*OftOwnershipTransferred, error) {
	event := new(OftOwnershipTransferred)
	if err := _Oft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the Oft contract.
type OftPeerSetIterator struct {
	Event *OftPeerSet // Event containing the contract specifics and raw log

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
func (it *OftPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftPeerSet)
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
		it.Event = new(OftPeerSet)
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
func (it *OftPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftPeerSet represents a PeerSet event raised by the Oft contract.
type OftPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Oft *OftFilterer) FilterPeerSet(opts *bind.FilterOpts) (*OftPeerSetIterator, error) {

	logs, sub, err := _Oft.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &OftPeerSetIterator{contract: _Oft.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Oft *OftFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *OftPeerSet) (event.Subscription, error) {

	logs, sub, err := _Oft.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftPeerSet)
				if err := _Oft.contract.UnpackLog(event, "PeerSet", log); err != nil {
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
func (_Oft *OftFilterer) ParsePeerSet(log types.Log) (*OftPeerSet, error) {
	event := new(OftPeerSet)
	if err := _Oft.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftPreCrimeSetIterator is returned from FilterPreCrimeSet and is used to iterate over the raw logs and unpacked data for PreCrimeSet events raised by the Oft contract.
type OftPreCrimeSetIterator struct {
	Event *OftPreCrimeSet // Event containing the contract specifics and raw log

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
func (it *OftPreCrimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftPreCrimeSet)
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
		it.Event = new(OftPreCrimeSet)
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
func (it *OftPreCrimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftPreCrimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftPreCrimeSet represents a PreCrimeSet event raised by the Oft contract.
type OftPreCrimeSet struct {
	PreCrimeAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPreCrimeSet is a free log retrieval operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Oft *OftFilterer) FilterPreCrimeSet(opts *bind.FilterOpts) (*OftPreCrimeSetIterator, error) {

	logs, sub, err := _Oft.contract.FilterLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return &OftPreCrimeSetIterator{contract: _Oft.contract, event: "PreCrimeSet", logs: logs, sub: sub}, nil
}

// WatchPreCrimeSet is a free log subscription operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Oft *OftFilterer) WatchPreCrimeSet(opts *bind.WatchOpts, sink chan<- *OftPreCrimeSet) (event.Subscription, error) {

	logs, sub, err := _Oft.contract.WatchLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftPreCrimeSet)
				if err := _Oft.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
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
func (_Oft *OftFilterer) ParsePreCrimeSet(log types.Log) (*OftPreCrimeSet, error) {
	event := new(OftPreCrimeSet)
	if err := _Oft.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Oft contract.
type OftTransferIterator struct {
	Event *OftTransfer // Event containing the contract specifics and raw log

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
func (it *OftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftTransfer)
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
		it.Event = new(OftTransfer)
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
func (it *OftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftTransfer represents a Transfer event raised by the Oft contract.
type OftTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Oft *OftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Oft.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OftTransferIterator{contract: _Oft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Oft *OftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OftTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Oft.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftTransfer)
				if err := _Oft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Oft *OftFilterer) ParseTransfer(log types.Log) (*OftTransfer, error) {
	event := new(OftTransfer)
	if err := _Oft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
