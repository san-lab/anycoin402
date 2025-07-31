// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.22;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { OFT } from "@layerzerolabs/oft-evm/contracts/OFT.sol";

/**
 * @title EIP-3009
 * @notice Provide internal implementation for gas-abstracted transfers
 * @dev Contracts that inherit from this must wrap these with publicly
 * accessible functions, optionally adding modifiers where necessary
 */
abstract contract OFT3009 is OFT  {
    // keccak256("TransferWithAuthorization(address from,address to,uint256 value,uint256 validAfter,uint256 validBefore,bytes32 nonce)")
    bytes32
        public constant TRANSFER_WITH_AUTHORIZATION_TYPEHASH = 0x7c7c6cdb67a18743f49ec6fa9b35f50d52ed05cbed4cc592e13b44501c1a2267;

    // keccak256("CancelAuthorization(address authorizer,bytes32 nonce)")
    bytes32
        public constant CANCEL_AUTHORIZATION_TYPEHASH = 0x158b0a9edf7a828aad02f63cd515c68ef2f50ba807396f6d12842833a1597429;

    /**
     * @dev authorizer address => nonce => bool (true if nonce is used)
     */
    mapping(address => mapping(bytes32 => bool)) private _authorizationStates;

    event AuthorizationUsed(address indexed authorizer, bytes32 indexed nonce);
    event AuthorizationCanceled(
        address indexed authorizer,
        bytes32 indexed nonce
    );

    bytes32 public DOMAIN_SEPARATOR;
    string public version;


     constructor(
        string memory _name,
        string memory _version,
        address _lzEndpoint,
        address _delegate
    ) OFT(_name, _name, _lzEndpoint, _delegate) Ownable(_delegate) {
        version = _version;
        DOMAIN_SEPARATOR = makeDomainSeparator(_name, _version);
    }

   
   
    function authorizationState(address authorizer, bytes32 nonce)
        external
        view
        returns (bool)
    {
        return _authorizationStates[authorizer][nonce];
    }

    
    function _transferWithAuthorization(
        address from,
        address to,
        uint256 value,
        uint256 validAfter,
        uint256 validBefore,
        bytes32 nonce,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) internal {
        _transferWithAuthorization(
            from,
            to,
            value,
            validAfter,
            validBefore,
            nonce,
            abi.encodePacked(r, s, v)
        );
    }

   
    function _transferWithAuthorization(
        address from,
        address to,
        uint256 value,
        uint256 validAfter,
        uint256 validBefore,
        bytes32 nonce,
        bytes memory signature
    ) internal {
        _requireValidAuthorization(from, nonce, validAfter, validBefore);
        _requireValidSignature(
            from,
            keccak256(
                abi.encode(
                    TRANSFER_WITH_AUTHORIZATION_TYPEHASH,
                    from,
                    to,
                    value,
                    validAfter,
                    validBefore,
                    nonce
                )
            ),
            signature
        );

        _markAuthorizationAsUsed(from, nonce);
        _transfer(from, to, value);
    }

    
    

    function _cancelAuthorization(
        address authorizer,
        bytes32 nonce,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) internal {
        _cancelAuthorization(authorizer, nonce, abi.encodePacked(r, s, v));
    }

    
    function _cancelAuthorization(
        address authorizer,
        bytes32 nonce,
        bytes memory signature
    ) internal {
        _requireUnusedAuthorization(authorizer, nonce);
        _requireValidSignature(
            authorizer,
            keccak256(
                abi.encode(CANCEL_AUTHORIZATION_TYPEHASH, authorizer, nonce)
            ),
            signature
        );

        _authorizationStates[authorizer][nonce] = true;
        emit AuthorizationCanceled(authorizer, nonce);
    }

    function cancelAuthorization(
        address authorizer,
        bytes32 nonce,
        bytes memory signature
    ) external {
        _cancelAuthorization(authorizer, nonce, signature);
    }
   

   function cancelAuthorization(
        address authorizer,
        bytes32 nonce,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external {
        _cancelAuthorization(authorizer, nonce, v, r, s);
    }

    
    function _requireValidSignature(
        address signer,
        bytes32 dataHash,
        bytes memory signature
    ) internal view {
        require(
            isValidSignatureNow(
                signer,
                toTypedDataHash(DOMAIN_SEPARATOR, dataHash),
                signature
            ),
            "FiatTokenV2: invalid signature"
        );
    }

    function toTypedDataHash(bytes32 domainSeparator, bytes32 structHash)
        internal
        pure
        returns (bytes32 digest)
    {
        assembly {
            let ptr := mload(0x40)
            mstore(ptr, "\x19\x01")
            mstore(add(ptr, 0x02), domainSeparator)
            mstore(add(ptr, 0x22), structHash)
            digest := keccak256(ptr, 0x42)
        }
    }

    
    function _requireUnusedAuthorization(address authorizer, bytes32 nonce)
        internal
        view
    {
        require(
            !_authorizationStates[authorizer][nonce],
            "FiatTokenV2: authorization is used or canceled"
        );
    }

    
    function _requireValidAuthorization(
        address authorizer,
        bytes32 nonce,
        uint256 validAfter,
        uint256 validBefore
    ) internal view {
        require(
            block.timestamp > validAfter,
            "FiatTokenV2: authorization is not yet valid"
        );
        require(block.timestamp < validBefore, "FiatTokenV2: authorization is expired");
        _requireUnusedAuthorization(authorizer, nonce);
    }

   
    function _markAuthorizationAsUsed(address authorizer, bytes32 nonce)
        internal
    {
        _authorizationStates[authorizer][nonce] = true;
        emit AuthorizationUsed(authorizer, nonce);
    }

    function isValidSignatureNow(
        address signer,
        bytes32 digest,
        bytes memory signature
    )  internal  view returns (bool) {
        if (!isContract(signer)) {
            return recover(digest, signature) == signer;
        }
        return isValidERC1271SignatureNow(signer, digest, signature);
    }

    
    function isValidERC1271SignatureNow(
        address signer,
        bytes32 digest,
        bytes memory signature
    ) internal view returns (bool) {
        (bool success, bytes memory result) = signer.staticcall(
            abi.encodeWithSelector(
                IERC1271.isValidSignature.selector,
                digest,
                signature
            )
        );
        return (success &&
            result.length >= 32 &&
            abi.decode(result, (bytes32)) ==
            bytes32(IERC1271.isValidSignature.selector));
    }

   
    function isContract(address addr) internal view returns (bool) {
        uint256 size;
        assembly {
            size := extcodesize(addr)
        }
        return size > 0;
    }

    
    function makeDomainSeparator(
        string memory _name,
        string memory _version
    ) internal view returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    // keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)")
                    0x8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f,
                    keccak256(bytes(_name)),
                    keccak256(bytes(_version)),
                    block.chainid,
                    address(this)
                )
            );
    
    }



   function recover(
        bytes32 digest,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) internal pure returns (address) {
        // EIP-2 still allows signature malleability for ecrecover(). Remove this possibility and make the signature
        // unique. Appendix F in the Ethereum Yellow paper (https://ethereum.github.io/yellowpaper/paper.pdf), defines
        // the valid range for s in (281): 0 < s < secp256k1n ÷ 2 + 1, and for v in (282): v ∈ {27, 28}. Most
        // signatures from current libraries generate a unique signature with an s-value in the lower half order.
        //
        // If your library generates malleable signatures, such as s-values in the upper range, calculate a new s-value
        // with 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141 - s1 and flip v from 27 to 28 or
        // vice versa. If your library also generates signatures with 0/1 for v instead 27/28, add 27 to v to accept
        // these malleable signatures as well.
        if (
            uint256(s) >
            0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0
        ) {
            revert("ECRecover: invalid signature 's' value");
        }

        if (v != 27 && v != 28) {
            revert("ECRecover: invalid signature 'v' value");
        }

        // If the signature is valid (and not malleable), return the signer address
        address signer = ecrecover(digest, v, r, s);
        require(signer != address(0), "ECRecover: invalid signature");

        return signer;
    }

    function recover(bytes32 digest, bytes memory signature)
        internal
        pure
        returns (address)
    {
        require(signature.length == 65, "ECRecover: invalid signature length");

        bytes32 r;
        bytes32 s;
        uint8 v;

        // ecrecover takes the signature parameters, and the only way to get them
        // currently is to use assembly.
        /// @solidity memory-safe-assembly
        assembly {
            r := mload(add(signature, 0x20))
            s := mload(add(signature, 0x40))
            v := byte(0, mload(add(signature, 0x60)))
        }
        return recover(digest, v, r, s);
    }
 
}

interface IERC1271 {
    /**
     *  @return magicValue   bytes4 magic value 0x1626ba7e when function passes
     */
    function isValidSignature(bytes32 hash, bytes memory signature)
        external
        view
        returns (bytes4 magicValue);
}

