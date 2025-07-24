# 🌊 Go-x402 Facilitator

This project is a **Go-based implementation of an x402 payment facilitator**. Its primary purpose is to provide a flexible platform for experimenting with and extending the x402 protocol beyond its initial scope—focusing on diverse token support, multi-chain environments, and advanced cross-chain payment scenarios.

---

## 💡 What is x402?

**x402** is an open standard that reimagines internet payments by leveraging the rarely-used HTTP `402 Payment Required` status code. It enables **instant, on-chain payments** (primarily with stablecoins) directly within standard HTTP interactions—making value transfer as seamless as data transfer.

The core idea is to **facilitate internet-native payments** for resources like API calls, data access, or digital content.

### 🧠 How It Works

1. A client (human or AI agent) requests a resource requiring payment.
2. The server responds with an HTTP 402, specifying payment details.
3. The client crafts and signs a payment authorization **off-chain**.
4. The client retries the request, attaching the signed payment.
5. The server (or facilitator) verifies and **submits it on-chain** to grant access.

### 🚀 Key Benefits of x402

- **Micropayments**: Enables economically viable payments for tiny amounts—ideal for pay-per-API or per-article models.
- **AI Agent Autonomy**: Allows AI agents to independently discover and pay for digital resources.
- **Frictionless Access**: No accounts, passwords, or sign-ups—**payment = access**.
- **Instant Settlement**: On-chain finality provides faster and more reliable confirmation than traditional rails.

---

## 🎯 Purpose of This Project

While the x402 protocol provides a powerful new paradigm, early implementations are often narrow in scope (e.g., limited to USDC or Ethereum). The **Go-x402 Facilitator** project aims to expand and test the boundaries of x402 through:

### ✅ Objectives

- **Test Multi-Token Compatibility**  
  Verify x402’s adaptability with various **ERC-20 tokens**, not just USDC.

- **Explore Multi-Chain Deployments**  
  Validate facilitator performance across **multiple EVM-compatible chains**.

- **Experiment with Cross-Chain Payments**  
  Integrate and test cross-chain x402 payments using **LayerZero's Omnichain Fungible Token (OFT)** standard—enabling seamless value transfer between blockchains.

By building a **flexible, Go-based** implementation, this project helps deepen understanding of x402 and its potential for internet-native, blockchain-powered commerce.

> 🔒 This facilitator is **stateless** and **never holds user funds**. Its sole role is to **verify and submit signed payment payloads** to the appropriate blockchains.

---

## 🛠️ Key Technologies & Standards

- **x402 Protocol**  
  The core standard for on-chain payments over HTTP.

- **[EIP-3009](https://eips.ethereum.org/EIPS/eip-3009)** – `transferWithAuthorization`  
  Enables off-chain token approvals for on-chain transfers.

- **[LayerZero OFT](https://layerzero.network)** – *Omnichain Fungible Token*  
  Allows fungible tokens to move across blockchains while maintaining a unified supply.

---

