package main

import (
	"fmt"
	"os"

	"software-distribution-onchain/client"
	"software-distribution-onchain/contract"
	"software-distribution-onchain/fingerprint"
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
	testClient, err := client.NewTestClient(client.PASEO_ASSET_HUB_RPC)
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

	contract.RunContractTest()
}
