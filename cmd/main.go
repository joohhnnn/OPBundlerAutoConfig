package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"OPBundlerAutoConfig/ethereum"
)

// Function to modify Bundler configuration
func modifyBundlerConfig(bundlerPath string) error {
	cmd := exec.Command("./scripts/modify_bundler.shh", bundlerPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute script: %w", err)
	}
	return nil
}

// Function to modify Optimism configuration
func modifyOptimismConfig(optimismPath string) error {
	cmd := exec.Command("./scripts/modify_op_devnet.sh", optimismPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute script: %w", err)
	}
	return nil
}

func main() {
	// Define command-line arguments
	var bundlerPath string
	var optimismPath string

	flag.StringVar(&bundlerPath, "bundlerPath", "", "Path to the Bundler configuration file")
	flag.StringVar(&optimismPath, "optimismPath", "", "Path to the Optimism configuration file")

	// Parse command-line arguments
	flag.Parse()

	// Validate arguments
	if bundlerPath == "" || optimismPath == "" {
		fmt.Println("Both --bundler and --optimism flags are required.")
		os.Exit(1)
	}

	// Perform configuration modifications
	if err := modifyBundlerConfig(bundlerPath); err != nil {
		fmt.Printf("Failed to modify Bundler config: %s\n", err)
		os.Exit(1)
	}

	if err := modifyOptimismConfig(optimismPath); err != nil {
		fmt.Printf("Failed to modify Optimism config: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Configuration successfully modified.")

	ethtransaction.SendTransaction()

	// Execute the start_services.sh script
	cmd := exec.Command("./scripts/start_services.sh", optimismPath, bundlerPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to start services: %s\n", err)
		os.Exit(1)
	}

}
