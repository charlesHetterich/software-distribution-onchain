package fingerprint

import (
	"encoding/hex"
	"fmt"
	"net"
	"runtime"
	"sort"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"golang.org/x/crypto/sha3"
)

// DeviceFingerprint represents device hardware information
type DeviceFingerprint struct {
	CPUModel     string   `json:"cpu_model"`
	CPUCores     int32    `json:"cpu_cores"`
	HostID       string   `json:"host_id"`
	Hostname     string   `json:"hostname"`
	MACAddresses []string `json:"mac_addresses"`
	OS           string   `json:"os"`
	Architecture string   `json:"architecture"`
	Platform     string   `json:"platform"`
	Uptime       uint64   `json:"uptime"`
}

// CollectDeviceInfo gathers comprehensive device information
func CollectDeviceInfo() (*DeviceFingerprint, error) {
	fingerprint := &DeviceFingerprint{}

	// Get CPU information
	if cpuInfo, err := cpu.Info(); err == nil && len(cpuInfo) > 0 {
		fingerprint.CPUModel = strings.TrimSpace(cpuInfo[0].ModelName)
		fingerprint.CPUCores = cpuInfo[0].Cores
	} else {
		fmt.Printf("Warning: Could not get CPU info: %v\n", err)
	}

	// Get host information
	if hostInfo, err := host.Info(); err == nil {
		fingerprint.HostID = hostInfo.HostID
		fingerprint.Hostname = hostInfo.Hostname
		fingerprint.OS = hostInfo.OS
		fingerprint.Platform = hostInfo.Platform
		fingerprint.Uptime = hostInfo.Uptime
	} else {
		fmt.Printf("Warning: Could not get host info: %v\n", err)
	}

	// Get runtime architecture
	fingerprint.Architecture = runtime.GOARCH

	// Get network interfaces (MAC addresses)
	if interfaces, err := net.Interfaces(); err == nil {
		var macs []string
		for _, iface := range interfaces {
			// Skip loopback and virtual interfaces
			if len(iface.HardwareAddr) > 0 &&
				!strings.Contains(strings.ToLower(iface.Name), "lo") &&
				!strings.Contains(strings.ToLower(iface.Name), "docker") &&
				!strings.Contains(strings.ToLower(iface.Name), "veth") {
				macs = append(macs, iface.HardwareAddr.String())
			}
		}
		sort.Strings(macs) // Ensure consistent ordering
		fingerprint.MACAddresses = macs
	} else {
		fmt.Printf("Warning: Could not get network interfaces: %v\n", err)
	}

	return fingerprint, nil
}

// GenerateFingerprint creates a hash from device information
func (df *DeviceFingerprint) GenerateFingerprint() string {
	var identifiers []string

	// Add non-empty fields to identifiers
	if df.CPUModel != "" {
		identifiers = append(identifiers, fmt.Sprintf("cpu:%s", df.CPUModel))
	}
	if df.CPUCores > 0 {
		identifiers = append(identifiers, fmt.Sprintf("cores:%d", df.CPUCores))
	}
	if df.HostID != "" {
		identifiers = append(identifiers, fmt.Sprintf("host:%s", df.HostID))
	}
	if df.Hostname != "" {
		identifiers = append(identifiers, fmt.Sprintf("hostname:%s", df.Hostname))
	}
	if df.OS != "" {
		identifiers = append(identifiers, fmt.Sprintf("os:%s", df.OS))
	}
	if df.Platform != "" {
		identifiers = append(identifiers, fmt.Sprintf("platform:%s", df.Platform))
	}
	if df.Architecture != "" {
		identifiers = append(identifiers, fmt.Sprintf("arch:%s", df.Architecture))
	}
	if len(df.MACAddresses) > 0 {
		identifiers = append(identifiers, fmt.Sprintf("mac:%s", strings.Join(df.MACAddresses, ",")))
	}

	// Sort for consistency
	sort.Strings(identifiers)

	// Create combined string
	combined := strings.Join(identifiers, "|")

	// Generate Keccak-256 hash (same as Ethereum uses)
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(combined))
	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash)
}

// PrintDetailedInfo displays comprehensive device information
func (df *DeviceFingerprint) PrintDetailedInfo() {
	fmt.Println("=== DEVICE FINGERPRINT DETAILS ===")
	fmt.Printf("CPU Model:     %s\n", df.CPUModel)
	fmt.Printf("CPU Cores:     %d\n", df.CPUCores)
	fmt.Printf("Host ID:       %s\n", df.HostID)
	fmt.Printf("Hostname:      %s\n", df.Hostname)
	fmt.Printf("OS:            %s\n", df.OS)
	fmt.Printf("Platform:      %s\n", df.Platform)
	fmt.Printf("Architecture:  %s\n", df.Architecture)
	fmt.Printf("Uptime:        %d seconds\n", df.Uptime)
	fmt.Printf("MAC Addresses: %v\n", df.MACAddresses)
	fmt.Println("=====================================")
}

// GenerateHashStep shows step-by-step hash generation
func (df *DeviceFingerprint) GenerateHashWithSteps() string {
	var identifiers []string

	fmt.Println("\n=== HASH GENERATION STEPS ===")

	// Build identifiers array with logging
	if df.CPUModel != "" {
		id := fmt.Sprintf("cpu:%s", df.CPUModel)
		identifiers = append(identifiers, id)
		fmt.Printf("✓ Added: %s\n", id)
	}
	if df.CPUCores > 0 {
		id := fmt.Sprintf("cores:%d", df.CPUCores)
		identifiers = append(identifiers, id)
		fmt.Printf("✓ Added: %s\n", id)
	}
	if df.HostID != "" {
		id := fmt.Sprintf("host:%s", df.HostID)
		identifiers = append(identifiers, id)
		fmt.Printf("✓ Added: %s\n", id)
	}
	if df.Hostname != "" {
		id := fmt.Sprintf("hostname:%s", df.Hostname)
		identifiers = append(identifiers, id)
		fmt.Printf("✓ Added: %s\n", id)
	}
	if df.OS != "" {
		id := fmt.Sprintf("os:%s", df.OS)
		identifiers = append(identifiers, id)
		fmt.Printf("✓ Added: %s\n", id)
	}
	if df.Platform != "" {
		id := fmt.Sprintf("platform:%s", df.Platform)
		identifiers = append(identifiers, id)
		fmt.Printf("✓ Added: %s\n", id)
	}
	if df.Architecture != "" {
		id := fmt.Sprintf("arch:%s", df.Architecture)
		identifiers = append(identifiers, id)
		fmt.Printf("✓ Added: %s\n", id)
	}
	if len(df.MACAddresses) > 0 {
		id := fmt.Sprintf("mac:%s", strings.Join(df.MACAddresses, ","))
		identifiers = append(identifiers, id)
		fmt.Printf("✓ Added: %s\n", id)
	}

	fmt.Printf("\nIdentifiers before sorting: %v\n", identifiers)

	// Sort for consistency
	sort.Strings(identifiers)
	fmt.Printf("Identifiers after sorting:  %v\n", identifiers)

	// Create combined string
	combined := strings.Join(identifiers, "|")
	fmt.Printf("Combined string: %s\n", combined)
	fmt.Printf("Combined length: %d characters\n", len(combined))

	// Generate Keccak-256 hash (same as Ethereum uses)
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(combined))
	hash := hasher.Sum(nil)
	hashStr := hex.EncodeToString(hash)

	fmt.Printf("Keccak-256 hash: %s\n", hashStr)
	fmt.Println("=============================")

	return hashStr
}
