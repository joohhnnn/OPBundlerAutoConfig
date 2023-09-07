package ethtransaction

import (
	"log"
	"math/big"
	"strings" 
	"github.com/ethereum/go-ethereum/rpc"
)

func SendTransaction() {
	// Connect to Ethereum node
	client, err := rpc.Dial("http://localhost:9545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	var accounts []string
	err = client.Call(&accounts, "eth_accounts")
	if err != nil || len(accounts) == 0 {
		log.Fatalf("Failed to get accounts: %v", err)
	}

	// Use the first account as the signer
	signerAddress := accounts[0]
	log.Printf("Signer address: %s", signerAddress)

	toAddress := "0x3fab184622dc19b6109349b94811493bf2a45362"
	
	// Get initial balance as a string and remove extra quotes
	var initialBalanceStr string
	err = client.Call(&initialBalanceStr, "eth_getBalance", toAddress, "latest")
	if err != nil {
		log.Fatalf("Failed to get initial balance: %v", err)
	}
	initialBalanceStr = strings.ReplaceAll(initialBalanceStr, "\"", "")

	// Convert hex string to big.Int
	initialBalance, success := new(big.Int).SetString(initialBalanceStr[2:], 16)
	if !success {
		log.Fatalf("Failed to parse initial balance: %v", err)
	}
	log.Printf("Initial balance of toAddress: %s", initialBalance.String())

	// Send Transaction
	txData := map[string]string{
		"from":     signerAddress,
		"to":       toAddress,
		"value":    "0xDE0B6B3A7640000",  // 1 ETH in Wei
		"gas":      "0x186A0",           // 100k Gas limit
		"gasPrice": "0x4A817C800",       // 20 Gwei
	}
	var txHash string
	err = client.Call(&txHash, "personal_sendTransaction", txData, "your_passphrase")
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	// Wait for the transaction to be mined and get the final balance
	// (Simplified by just querying the balance)
	
	// Get final balance as a string and remove extra quotes
	var finalBalanceStr string
	err = client.Call(&finalBalanceStr, "eth_getBalance", toAddress, "latest")
	if err != nil {
		log.Fatalf("Failed to get final balance: %v", err)
	}
	finalBalanceStr = strings.ReplaceAll(finalBalanceStr, "\"", "")

	// Convert hex string to big.Int
	finalBalance, success := new(big.Int).SetString(finalBalanceStr[2:], 16)
	if !success {
		log.Fatalf("Failed to parse final balance: %v", err)
	}
	log.Printf("Final balance of toAddress: %s", finalBalance.String())
}
