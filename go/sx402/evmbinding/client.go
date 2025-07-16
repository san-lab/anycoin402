package evmbinding

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const Base_sepolia = "base-sepolia"
const Amoy = "amoy"
const Sepolia = "sepolia"
const Holesky = "holesky"
const ZkSync_sepolia = "zksync-sepolia"
const Arbitrum_sepolia = "arbitrum-sepolia"
const OP_Sepolia = "op-sepolia"

var rpcEndpoints = map[string]string{
	Base_sepolia:     "https://sepolia.base.org",
	Sepolia:          "https://ethereum-sepolia-rpc.publicnode.com",
	Amoy:             "https://rpc-amoy.polygon.technology/",
	Holesky:          "https://ethereum-holesky.publicnode.com",
	ZkSync_sepolia:   "https://sepolia.era.zksync.dev",
	Arbitrum_sepolia: "https://sepolia-rollup.arbitrum.io/rpc", // "https://arbitrum-sepolia.gateway.tenderly.co",
	OP_Sepolia:       "https://optimism-sepolia.gateway.tenderly.co",
}

var ExplorerURLs = map[string]string{
	Base_sepolia:     "https://sepolia.basescan.org", // Basescan (Etherscan-style)
	Sepolia:          "https://sepolia.etherscan.io", // Official Sepolia Etherscan
	Amoy:             "https://amoy.polygonscan.com", // Polygonscan for Amoy
	Holesky:          "https://holesky.etherscan.io", // Etherscan for Holesky
	ZkSync_sepolia:   "https://sepolia-era.zksync.network/",
	Arbitrum_sepolia: "https://sepolia.arbiscan.io",
	OP_Sepolia:       "https://sepolia-optimism.etherscan.io/",
}

func init() {
	LoadConfigs()
}

func LoadConfigs() {
	log.Println(LoadOverrides("config/rpcs.json", rpcEndpoints))
	log.Println(LoadOverrides("config/explorers.json", rpcEndpoints))
}

func LoadOverrides(relativePath string, targetMap map[string]string) error {
	// Absolute path from the working directory (project root)
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		return fmt.Errorf("could not resolve path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("could not read file %s: %w", absPath, err)
	}

	var overrides map[string]string
	if err := json.Unmarshal(data, &overrides); err != nil {
		return fmt.Errorf("invalid JSON in %s: %w", absPath, err)
	}

	maps.Copy(targetMap, overrides)

	return nil
}

var ChainIDs = map[string]*big.Int{
	Base_sepolia:     big.NewInt(84532),
	Sepolia:          big.NewInt(11155111),
	Amoy:             big.NewInt(80002),
	Holesky:          big.NewInt(17000),
	ZkSync_sepolia:   big.NewInt(300),
	Arbitrum_sepolia: big.NewInt(421614),
	OP_Sepolia:       big.NewInt(11155420),
}

func GetClientByNetwork(network string) (client *ethclient.Client, err error) {
	url, ok := GetRPCEndpoint(network)
	if !ok {
		err = fmt.Errorf("Unknown network: %s", network)
		return
	}
	return ethclient.Dial(url)
}

func GetlientByChainID(chainID *big.Int) (client *ethclient.Client, err error) {
	network := ""
	for k, v := range ChainIDs {
		if v.Cmp(chainID) == 0 {
			network = k
			break
		}
	}
	if len(network) == 0 {
		err = fmt.Errorf("Unsupported ChainID: %v", chainID)
		return
	}
	return GetClientByNetwork(network)
}

func InitClients() map[string]*ethclient.Client {
	clients := make(map[string]*ethclient.Client)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for network, url := range rpcEndpoints {
		wg.Add(1)
		go func(network, url string) {
			defer wg.Done()
			client, err := ethclient.Dial(url)
			if err != nil {
				log.Printf("❌ Failed to connect to %s: %v", network, err)
				return
			}
			log.Printf("✅ Connected to %s", network)
			mu.Lock()
			clients[network] = client
			mu.Unlock()
		}(network, url)
	}

	wg.Wait()
	return clients
}

func GetRPCEndpoint(network string) (string, bool) {
	rpc, ok := rpcEndpoints[network]
	return rpc, ok
}

func SendTransaction(client *ethclient.Client, signedTx *types.Transaction) (*common.Hash, error) {

	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("could not send tx: %v", err)
	}
	h := signedTx.Hash()
	return &h, nil
}

func CheckTokenBalance(client *ethclient.Client, tokenAddress, ownerAddress common.Address) (*big.Int, error) {

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(tokenABI))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse ABI: %v", err)
	}

	// Pack the input (balanceOf(address))
	data, err := parsedABI.Pack("balanceOf", ownerAddress)
	if err != nil {
		log.Fatalf("Failed to pack data: %v", err)
	}

	// Prepare the call message
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	// Call the contract
	ctx := context.Background()
	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		log.Printf("Failed to call contract: %v\n", err)
		return nil, err
	}

	// Unpack the result
	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", result)

	return balance, err
}

// Returns if the nonce is "known"
func CheckAuthorizationState(client *ethclient.Client, tokenAddress, payer common.Address, nonce [32]byte) (bool, error) {
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(tokenABI))
	if err != nil {
		return false, fmt.Errorf("Failed to parse ABI: %v", err)
	}

	// Pack the input (balanceOf(address))
	data, err := parsedABI.Pack("authorizationState", payer, nonce)
	if err != nil {
		log.Fatalf("Failed to pack data: %v", err)
	}

	// Prepare the call message
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	// Call the contract
	ctx := context.Background()
	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		log.Printf("Failed to call contract: %v", err)
		return false, err
	}

	// Unpack the result

	unpacked, err := parsedABI.Unpack("authorizationState", result)
	known, ok := unpacked[0].(bool)
	if !ok {
		log.Fatalf("Unpacked result is not a bool")
	}

	return known, err
}

func TransferWithAuthorization(
	client *ethclient.Client,
	signer *ecdsa.PrivateKey,
	token, from, to common.Address,
	value, validAfter, validBefore *big.Int,
	nonce, r, s [32]byte,
	v byte,
) (*common.Hash, error) {
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(trWithAuthABI))
	if err != nil {
		return nil, err
	}

	// Pack the input (balanceOf(address))
	data, err := parsedABI.Pack("transferWithAuthorization", from, to, value, validAfter, validBefore, nonce, v, r, s)
	if err != nil {
		return nil, fmt.Errorf("Failed to pack data: %w", err)
	}

	facilAddress := crypto.PubkeyToAddress(signer.PublicKey)

	fromNonce, err := getNonce(context.Background(), client, facilAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Estimate gas
	gasLimit := uint64(500000) // or estimate with client.EstimateGas()

	tx := types.NewTransaction(
		fromNonce,
		token,
		big.NewInt(0), // No ETH being sent
		gasLimit,
		gasPrice,
		data,
	)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error recoovering ChainID: %w", err)
	}
	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), signer)
	if err != nil {
		return nil, fmt.Errorf("error signing transaction: %w", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("could not send tx: %v", err)
	}
	h := signedTx.Hash()
	return &h, nil

}

const tokenABI = `[
  {
    "constant": true,
    "inputs": [{"name": "_owner", "type": "address"}],
    "name": "balanceOf",
    "outputs": [{"name": "balance", "type": "uint256"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [
      {"name": "authorizer", "type": "address"},
      {"name": "nonce", "type": "bytes32"}
    ],
    "name": "authorizationState",
    "outputs": [{"name": "", "type": "bool"}],
    "stateMutability": "view",
    "type": "function"
  }
]`

const trWithAuthABI = `[{
	"constant": false,
	"inputs": [
	  { "name": "from", "type": "address" },
	  { "name": "to", "type": "address" },
	  { "name": "value", "type": "uint256" },
	  { "name": "validAfter", "type": "uint256" },
	  { "name": "validBefore", "type": "uint256" },
	  { "name": "nonce", "type": "bytes32" },
	  { "name": "v", "type": "uint8" },
	  { "name": "r", "type": "bytes32" },
	  { "name": "s", "type": "bytes32" }
	],
	"name": "transferWithAuthorization",
	"outputs": [],
	"stateMutability": "nonpayable",
	"type": "function"
  }]`
