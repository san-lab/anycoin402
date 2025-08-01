// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.22;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IOFT, SendParam, OFTLimit, OFTReceipt, OFTFeeDetail, MessagingReceipt, MessagingFee } from "@layerzerolabs/oft-evm/contracts/interfaces/IOFT.sol";
import "./OFT3009CC.sol";



contract MyOFT3009CC is OFT3009CC {
    constructor(
        string memory _name,
        string memory _version,
        address _lzEndpoint,
        address _delegate
    ) OFT3009CC(_name, _version, _lzEndpoint, _delegate)  {}

    bytes32 public constant CROSS_CHAIN_TRANSFER_TYPEHASH = keccak256(
        "CrossChainTransferWithAuthorization(address from,address to,uint256 amount,uint256 minimalAmount,uint256 destinationChain,uint256 validAfter,uint256 validBefore,bytes32 nonce)"
    );

    // facilitator=>(dstEid=>markup))
    mapping (address=>mapping(uint32=>uint256)) public markups;

    function setMarkup(uint256 _markup, uint32 dstEid) external {
        markups[msg.sender][dstEid]=_markup;
    }

    function sharedDecimals() public pure override returns (uint8) {
        return 6;
    }

    function decimals() public pure override returns(uint8) {
        return 6;
    }

    function mint(address to, uint amount)  external  onlyOwner  {
        _mint(to, amount);
    }


    function transferWithAuthorization(
        address from,
        address to,
        uint256 value,
        uint256 validAfter,
        uint256 validBefore,
        bytes32 nonce,
        bytes memory signature
    ) external  {
        _transferWithAuthorization(
            from,
            to,
            value,
            validAfter,
            validBefore,
            nonce,
            signature
        );
    }


   
      function transferWithAuthorization(
        address from,
        address to,
        uint256 value,
        uint256 validAfter,
        uint256 validBefore,
        bytes32 nonce,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external   {
        _transferWithAuthorization(
            from,
            to,
            value,
            validAfter,
            validBefore,
            nonce,
            v,
            r,
            s
        );
    }



      function sendFrom(
        SendParam calldata _sendParam,
        MessagingFee calldata _fee,
        address _from, 
        address _refundAddress
        ) public payable returns (bool) {
        address spender = msg.sender;
        _spendAllowance(_from, spender, _sendParam.amountLD );
        _transfer(_from, spender, _sendParam.amountLD);
        _send( _sendParam, _fee, _refundAddress);
        return true;
    }

    function sendWithCCAuthorization(
        SendParam calldata _sendParam,
        MessagingFee calldata _fee,
        address _from, 
        uint256 validAfter,
        uint256 validBefore,
        bytes32 nonce,
        bytes memory signature,
        address _refundAddress
        ) public payable returns (bool) {
        
        _requireValidAuthorization(_from, nonce, validAfter, validBefore);
        _requireValidSignature(
            _from,
            keccak256(
                abi.encode(
                    CROSS_CHAIN_TRANSFER_TYPEHASH,
                    _from,
                    _sendParam.to,
                    _sendParam.amountLD,
                    _sendParam.minAmountLD,
                    _sendParam.dstEid,
                    validAfter,
                    validBefore,
                    nonce
                )
            ),
            signature
        );

        _markAuthorizationAsUsed(_from, nonce);
        _transfer(_from, msg.sender, _sendParam.amountLD);
        SendParam memory newSendPAram = _sendParam;
        newSendPAram.amountLD = _sendParam.amountLD - markups[msg.sender][_dstEid];
        _send(newSendParam, _fee, _refundAddress);
        return true;
    }

    function HashCCData(address from, address to, 
            uint256 amountLD, uint256 minAmountLD, uint256 dstEid, 
            uint256 validAfter, uint256 validBefore, 
            bytes32 nonce) public pure returns (bytes32) {
       return keccak256(
                abi.encode(
                    CROSS_CHAIN_TRANSFER_TYPEHASH,
                    from,
                    to,
                    amountLD,
                    minAmountLD,
                    dstEid,
                    validAfter,
                    validBefore,
                    nonce
                )
            );
    }


    
    

}
