// SPDX-License-Identifier: MIT

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

    
    function setCrosschainMarkup(uint256 _markup, uint32 dstEid) external {
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
        _send(_from, _sendParam, _fee, _refundAddress);
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
        _send(_from, _sendParam, _fee, _refundAddress);
        return true;
    }



    // Forking this method because we really need to serve the usecase where _from!=msg.sender 
    // Another option would be to modify the SendParam, but it is marked calldata/immutabel throughout the stack...
    function _send(
        address _from,
        SendParam calldata _sendParam,
        MessagingFee calldata _fee,
        address _refundAddress
    ) internal virtual returns (MessagingReceipt memory msgReceipt, OFTReceipt memory oftReceipt) {
        // @dev Applies the token transfers regarding this send() operation.
        // - amountSentLD is the amount in local decimals that was ACTUALLY sent/debited from the sender.
        // - amountReceivedLD is the amount in local decimals that will be received/credited to the recipient on the remote OFT instance.
        (uint256 amountSentLD, uint256 amountReceivedLD) = _debit(
            _from,
            _sendParam.amountLD,
            _sendParam.minAmountLD,
            _sendParam.dstEid
        );
        // @dev Builds the options and OFT message to quote in the endpoint.
        (bytes memory message, bytes memory options) = _buildMsgAndOptions(_sendParam, amountReceivedLD);

        // @dev Sends the message to the LayerZero endpoint and returns the LayerZero msg receipt.
        msgReceipt = _lzSend(_sendParam.dstEid, message, options, _fee, _refundAddress);
        // @dev Formulate the OFT receipt.
        oftReceipt = OFTReceipt(amountSentLD, amountReceivedLD);

        emit OFTSent(msgReceipt.guid, _sendParam.dstEid, msg.sender, amountSentLD, amountReceivedLD);
    }

    function _debit(
        address _from,
        uint256 _amountLD,
        uint256 _minAmountLD,
        uint32 _dstEid
    ) internal virtual override returns (uint256 amountSentLD, uint256 amountReceivedLD) {

        if (msg.sender != _from) { //  3rd Party/markup usecase
            uint256 markup = markups[msg.sender][_dstEid];
            if (markup > 0) {
                    _amountLD -= markup ;  // We assume the _debitView will check > _minAmountLD
                                            // And we assume Solidity >= 0.8 (underflow)
                    _transfer(_from, msg.sender, markup);
            }
            
        }

        (amountSentLD, amountReceivedLD) = _debitView(_amountLD, _minAmountLD, _dstEid);

        // @dev In NON-default OFT, amountSentLD could be 100, with a 10% fee, the amountReceivedLD amount is 90,
        // therefore amountSentLD CAN differ from amountReceivedLD.

        // @dev Default OFT burns on src.
        _burn(_from, amountSentLD);
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
