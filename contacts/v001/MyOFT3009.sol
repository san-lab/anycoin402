// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.22;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IOFT, SendParam, OFTLimit, OFTReceipt, OFTFeeDetail, MessagingReceipt, MessagingFee } from "@layerzerolabs/oft-evm/contracts/interfaces/IOFT.sol";
import "./OFT3009.sol";


contract MyOFT3009 is OFT3009 {
    constructor(
        string memory _name,
        string memory _version,
        address _lzEndpoint,
        address _delegate
    ) OFT3009(_name, _version, _lzEndpoint, _delegate)  {}


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
        ) public virtual payable returns (bool) {
        address spender = msg.sender;
        _spendAllowance(_from, spender, _sendParam.amountLD );
        _transfer(_from, spender, _sendParam.amountLD);
        _send( _sendParam, _fee, _refundAddress);
        return true;
    }

    function sendWithAuthorization(
        SendParam calldata _sendParam,
        MessagingFee calldata _fee,
        address _from, 
        uint256 validAfter,
        uint256 validBefore,
        bytes32 nonce,
        bytes memory signature,
        address _refundAddress
        ) public virtual payable returns (bool) {
        
        _requireValidAuthorization(_from, nonce, validAfter, validBefore);
        _requireValidSignature(
            _from,
            keccak256(
                abi.encode(
                    TRANSFER_WITH_AUTHORIZATION_TYPEHASH,
                    _from,
                    _sendParam.to,
                    _sendParam.amountLD,
                    validAfter,
                    validBefore,
                    nonce
                )
            ),
            signature
        );

        _markAuthorizationAsUsed(_from, nonce);
        _transfer(_from, msg.sender, _sendParam.amountLD);
        _send(_sendParam, _fee, _refundAddress);
        return true;
    }


    





    

}
