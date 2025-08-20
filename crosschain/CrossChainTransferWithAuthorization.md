# CrossChainTransferWithAuthorization (EIP-712)

##  Rationale

The `CrossChainTransferWithAuthorization` structure extends [EIP-3009](https://eips.ethereum.org/EIPS/eip-3009) to support **cross-chain token transfers**. It introduces:

- **`minimalAmount`**: ensures slippage protection on the destination chain.
- **`destinationChain`**: indicates the target chain ID for the transfer.

This enables secure, gasless cross-chain transfers via meta-transactions and relayers, while ensuring that the user-signed authorization enforces the minimum acceptable received amount.

---

##  EIP-712 Struct Definition

```solidity
struct CrossChainTransferWithAuthorization {
  address from;
  address to;
  uint256 amount;
  uint256 minimalAmount;
  uint256 destinationChain;
  uint256 validAfter;
  uint256 validBefore;
  bytes32 nonce;
}
```

---

##  EIP-712 Type Definition

```
CrossChainTransferWithAuthorization(
  address from,
  address to,
  uint256 amount,
  uint256 minimalAmount,
  uint32 destinationChain,
  uint256 validAfter,
  uint256 validBefore,
  bytes32 nonce
)
```

---

##  Type Hash (Solidity)

```solidity
bytes32 constant CROSS_CHAIN_TRANSFER_WITH_AUTHORIZATION_TYPEHASH = keccak256(
  "CrossChainTransferWithAuthorization(address from,address to,uint256 amount,uint256 minimalAmount,uint256 destinationChain,uint256 validAfter,uint256 validBefore,bytes32 nonce)"
);
```

---

##  EIP-712 Hash Calculation

```solidity
bytes32 structHash = keccak256(abi.encode(
  CROSS_CHAIN_TRANSFER_WITH_AUTHORIZATION_TYPEHASH,
  from,
  to,
  amount,
  minimalAmount,
  destinationChain,
  validAfter,
  validBefore,
  nonce
));
```

```solidity
bytes32 digest = keccak256(abi.encodePacked(
  "\x19\x01",
  domainSeparator,
  structHash
));
```

---

##  Use Cases

- Secure **cross-chain transfers** with guaranteed minimum output.
- Gasless UX for users authorizing transfers via wallets off-chain.