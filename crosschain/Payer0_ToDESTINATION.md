# x402 Cross-Chain Scheme: `payer0_ToDESTINATION`

## Overview

The `payer0_ToDESTINATION` scheme extends the x402 exact payment scheme to support **cross-chain token transfers**. It enables clients to authorize a token transfer on the source chain with a guarantee that a **minimal amount** arrives on a specified **destination chain**.

---

## PaymentRequirements

When a resource requires payment, the server responds with HTTP 402 and a JSON body:

```json
{
  "x402Version": 1,
  "accepts": [
    {
      "scheme": "payer0_ToDESTINATION",
      "asset": "<tokenAddress>",
      "network": "<sourceChainId>",
      "maxAmountRequired": "<stringified integer>",
      "payTo": "<recipient address>",
      "description": "<optional description>",
      "extra": {
        "name": "<name>",
        "version": "<version>",
        "dstEid": 40265
      }
    }
  ],
  "error": "<optional message>"
}
```


scheme: MUST be "payer0_ToDESTINATION" for this cross-chain scheme.

asset, network, maxAmountRequired, payTo: as per exact scheme, defining what, where, and how much.

extra.dstEid: mandatory numeric destination chain ID where tokens must be received.

PaymentPayload
Clients respond with a JSON payload in the X-Payment header that encodes a signed CrossChainTransferAuthorization message:

```json
{
  "domain": {
    "name": "<Token Name>",
    "version": "<Token Version>",
    "chainId": <sourceChainId>,
    "verifyingContract": "<tokenAddress>"
  },
  "message": {
    "from": "<payer address>",
    "to": "<payTo>",
    "amount": "<exact amount to send>",
    "minimalAmount": "<minimum acceptable amount on destination>",
    "destinationChain": <destinationChainId>,
    "validAfter": "<timestamp>",
    "validBefore": "<timestamp>",
    "nonce": "<32-byte nonce hex>"
  },
  "signature": "<EIP-712 signature over the above>"
}
```

message.amount: token amount authorized to transfer on the source chain, must be ≤ maxAmountRequired.

message.minimalAmount: minimal amount the recipient should receive on the destination chain (slippage allowed).

message.destinationChain: MUST match the dstEid from PaymentRequirements.extra.

Signature scheme follows EIP-712 to ensure integrity and replay protection.
