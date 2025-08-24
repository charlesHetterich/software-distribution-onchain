package main

import (
	"fmt"

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
}
