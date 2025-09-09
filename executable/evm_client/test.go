package evm_client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
)

// ERC20 ABI for basic token functions
const ERC20_ABI = `[
	{
		"inputs": [],
		"name": "name",
		"outputs": [{"internalType": "string", "name": "", "type": "string"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "symbol",
		"outputs": [{"internalType": "string", "name": "", "type": "string"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "totalSupply",
		"outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [{"internalType": "address", "name": "account", "type": "address"}],
		"name": "balanceOf",
		"outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
		"stateMutability": "view",
		"type": "function"
	}
]`

type TokenClient struct {
	client          *ethclient.Client
	contractAddress common.Address
	contractABI     abi.ABI
	privateKey      *ecdsa.PrivateKey
	fromAddress     common.Address
}

func NewTokenClient(rpcURL, contractAddress, privateKeyHex string) (*TokenClient, error) {
	// Connect to the Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	// Parse contract ABI
	contractABI, err := abi.JSON(strings.NewReader(ERC20_ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// Get public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &TokenClient{
		client:          client,
		contractAddress: common.HexToAddress(contractAddress),
		contractABI:     contractABI,
		privateKey:      privateKey,
		fromAddress:     fromAddress,
	}, nil
}

func (tc *TokenClient) GetName() (string, error) {
	data, err := tc.contractABI.Pack("name")
	if err != nil {
		return "", err
	}

	msg := ethereum.CallMsg{
		To:   &tc.contractAddress,
		Data: data,
	}

	result, err := tc.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", err
	}

	var name string
	err = tc.contractABI.UnpackIntoInterface(&name, "name", result)
	return name, err
}

func (tc *TokenClient) GetSymbol() (string, error) {
	data, err := tc.contractABI.Pack("symbol")
	if err != nil {
		return "", err
	}

	msg := ethereum.CallMsg{
		To:   &tc.contractAddress,
		Data: data,
	}

	result, err := tc.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", err
	}

	var symbol string
	err = tc.contractABI.UnpackIntoInterface(&symbol, "symbol", result)
	return symbol, err
}

func (tc *TokenClient) GetTotalSupply() (*big.Int, error) {
	data, err := tc.contractABI.Pack("totalSupply")
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &tc.contractAddress,
		Data: data,
	}

	result, err := tc.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}

	var totalSupply *big.Int
	err = tc.contractABI.UnpackIntoInterface(&totalSupply, "totalSupply", result)
	return totalSupply, err
}

func (tc *TokenClient) GetBalance(address common.Address) (*big.Int, error) {
	data, err := tc.contractABI.Pack("balanceOf", address)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &tc.contractAddress,
		Data: data,
	}

	result, err := tc.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}

	var balance *big.Int
	err = tc.contractABI.UnpackIntoInterface(&balance, "balanceOf", result)
	return balance, err
}

func (tc *TokenClient) TestContract() error {
	fmt.Println("🚀 Testing MyToken Contract with Go")
	fmt.Println("===================================")

	// Test contract name
	name, err := tc.GetName()
	if err != nil {
		return fmt.Errorf("failed to get name: %v", err)
	}
	fmt.Printf("Token Name: %s\n", name)

	// Test contract symbol
	symbol, err := tc.GetSymbol()
	if err != nil {
		return fmt.Errorf("failed to get symbol: %v", err)
	}
	fmt.Printf("Token Symbol: %s\n", symbol)

	// Test total supply
	totalSupply, err := tc.GetTotalSupply()
	if err != nil {
		return fmt.Errorf("failed to get total supply: %v", err)
	}

	// Convert from wei to tokens (divide by 10^18)
	totalSupplyTokens := new(big.Float).Quo(new(big.Float).SetInt(totalSupply), big.NewFloat(1e18))
	fmt.Printf("Total Supply: %s tokens\n", totalSupplyTokens.Text('f', 2))

	// Test balance of deployer
	balance, err := tc.GetBalance(tc.fromAddress)
	if err != nil {
		return fmt.Errorf("failed to get balance: %v", err)
	}

	balanceTokens := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("Deployer Balance: %s tokens\n", balanceTokens.Text('f', 2))
	fmt.Printf("Deployer Address: %s\n", tc.fromAddress.Hex())

	return nil
}

func (tc *TokenClient) Close() {
	tc.client.Close()
}

func RunContractTest() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	PRIVATE_KEY := os.Getenv("PRIVATE_KEY")
	RPC_URL := os.Getenv("RPC_URL")
	CONTRACT_ADDRESS := os.Getenv("CONTRACT_ADDRESS")

	tokenClient, err := NewTokenClient(RPC_URL, CONTRACT_ADDRESS, PRIVATE_KEY)
	if err != nil {
		return fmt.Errorf("failed to create token client: %v", err)
	}
	defer tokenClient.Close()

	return tokenClient.TestContract()
}
