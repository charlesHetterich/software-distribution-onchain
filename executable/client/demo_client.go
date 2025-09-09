// File: client/test_client.go
package client

import (
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

const PASEO_ASSET_HUB_RPC = "wss://asset-hub-paseo-rpc.dwellir.com"

// TestClient handles basic Polkadot API tests
type TestClient struct {
	api *gsrpc.SubstrateAPI
}

// NewTestClient creates a new test client
func NewTestClient(rpcEndpoint string) (*TestClient, error) {
	fmt.Printf("🔗 Connecting to %s...\n", rpcEndpoint)

	api, err := gsrpc.NewSubstrateAPI(rpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %v", err)
	}

	fmt.Println("✅ Connected successfully!")
	return &TestClient{api: api}, nil
}

// GetGenesisHash gets the block hash at height 0 (genesis block)
func (c *TestClient) GetGenesisHash() (*types.Hash, error) {
	fmt.Println("📋 Getting genesis block hash (block 0)...")

	// Use GetBlockHash with 0 to get genesis block
	hash, err := c.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, fmt.Errorf("failed to get genesis hash: %v", err)
	}

	fmt.Printf("✅ Genesis hash: %s\n", hash.Hex())
	return &hash, nil
}

// GetLatestBlockHash gets the latest block hash
func (c *TestClient) GetLatestBlockHash() (*types.Hash, error) {
	fmt.Println("📋 Getting latest block hash...")

	hash, err := c.api.RPC.Chain.GetBlockHashLatest()
	if err != nil {
		return nil, fmt.Errorf("failed to get latest hash: %v", err)
	}

	fmt.Printf("✅ Latest hash: %s\n", hash.Hex())
	return &hash, nil
}

// GetChainInfo gets basic chain information
func (c *TestClient) GetChainInfo() error {
	fmt.Println("📋 Getting chain information...")

	// Get chain name
	chain, err := c.api.RPC.System.Chain()
	if err != nil {
		return fmt.Errorf("failed to get chain: %v", err)
	}
	fmt.Printf("🔗 Chain: %s\n", chain)

	// Get node name
	name, err := c.api.RPC.System.Name()
	if err != nil {
		return fmt.Errorf("failed to get node name: %v", err)
	}
	fmt.Printf("🖥️  Node: %s\n", name)

	// Get node version
	version, err := c.api.RPC.System.Version()
	if err != nil {
		return fmt.Errorf("failed to get node version: %v", err)
	}
	fmt.Printf("📦 Node Version: %s\n", version)

	// Get runtime version
	runtimeVersion, err := c.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return fmt.Errorf("failed to get runtime version: %v", err)
	}
	fmt.Printf("🔢 Runtime: %s v%d\n", runtimeVersion.SpecName, runtimeVersion.SpecVersion)

	return nil
}

// GetBlockByNumber gets block information by number
func (c *TestClient) GetBlockByNumber(blockNumber uint64) error {
	fmt.Printf("📋 Getting block #%d...\n", blockNumber)

	// Get block hash for the number - pass uint64 directly
	hash, err := c.api.RPC.Chain.GetBlockHash(blockNumber)
	if err != nil {
		return fmt.Errorf("failed to get block hash: %v", err)
	}

	// Get the full block
	block, err := c.api.RPC.Chain.GetBlock(hash)
	if err != nil {
		return fmt.Errorf("failed to get block: %v", err)
	}

	fmt.Printf("✅ Block #%d:\n", blockNumber)
	fmt.Printf("   Hash: %s\n", hash.Hex())
	fmt.Printf("   Parent: %s\n", block.Block.Header.ParentHash.Hex())
	fmt.Printf("   Extrinsics: %d\n", len(block.Block.Extrinsics))

	return nil
}

// TestBasicOperations runs a series of basic tests
func (c *TestClient) TestBasicOperations() error {
	fmt.Println("\n🧪 Running basic API tests...")
	fmt.Println("=====================================")

	// Test 1: Get genesis hash (block 0)
	if _, err := c.GetGenesisHash(); err != nil {
		return err
	}

	// Test 2: Get latest block hash
	if _, err := c.GetLatestBlockHash(); err != nil {
		return err
	}

	// Test 3: Get chain information
	if err := c.GetChainInfo(); err != nil {
		return err
	}

	// Test 4: Get specific blocks
	fmt.Println("\n📦 Testing block retrieval...")
	if err := c.GetBlockByNumber(0); err != nil {
		return err
	}

	if err := c.GetBlockByNumber(1); err != nil {
		return err
	}

	fmt.Println("\n🎉 All tests passed!")
	return nil
}

// Close closes the API connection
func (c *TestClient) Close() {
	if c.api != nil {
		// The gsrpc client doesn't have an explicit close method
		// The connection will be closed when the client is garbage collected
		fmt.Println("🔌 Connection will be closed")
	}
}
