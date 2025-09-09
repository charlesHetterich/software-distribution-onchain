package main

import (
	"fmt"
	"os"

	"software-distribution-onchain/evm_client"
	"software-distribution-onchain/fingerprint"
	"software-distribution-onchain/substrate_client"
)

func main() {
	fmt.Println("🔍 Device Fingerprint Test")
	fmt.Println("==========================")

	fmt.Println("📋 Collecting device information...")
	deviceInfo, err := fingerprint.CollectDeviceInfo()
	if err != nil {
		fmt.Printf("❌ Error collecting device info: %v\n", err)
		return
	}

	deviceInfo.PrintDetailedInfo()

	hash := deviceInfo.GenerateHashWithSteps()

	fmt.Println("\n🎯 FINAL RESULT")
	fmt.Printf("Device Hash: %s\n", hash)
	fmt.Printf("Hash Length: %d characters\n", len(hash))

	fmt.Println("🚀 Polkadot API Test")
	fmt.Println("===================")

	// Create test client
	testClient, err := substrate_client.NewTestClient(substrate_client.PASEO_ASSET_HUB_RPC)
	if err != nil {
		fmt.Printf("❌ Failed to create client: %v\n", err)
		os.Exit(1)
	}
	defer testClient.Close()

	// Run tests
	if err := testClient.TestBasicOperations(); err != nil {
		fmt.Printf("❌ Test failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Polkadot API connection test successful!")

	// Go version of https://docs.polkadot.com/develop/smart-contracts/dev-environments/hardhat/#interacting-with-your-contract
	evm_client.RunContractTest()
}
