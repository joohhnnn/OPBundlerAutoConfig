package ethtransaction

import (
	"encoding/json"
	"log"
	"math/big"
	"strings"
	"time"

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
	
	// Get initial balance as a JSON RawMessage, then convert it to a string
	var initialBalanceRaw json.RawMessage
	err = client.Call(&initialBalanceRaw, "eth_getBalance", toAddress, "latest")
	if err != nil {
		log.Fatalf("Failed to get initial balance: %v", err)
	}
	initialBalanceStr := string(initialBalanceRaw)
	initialBalanceStr = strings.Trim(initialBalanceStr, "\"")

	// Convert hex string to big.Int
	initialBalance, success := new(big.Int).SetString(initialBalanceStr[2:], 16)
	if !success {
		log.Fatalf("Failed to parse initial balance: %v", err)
	}
	log.Printf("Initial balance of toAddress: %s", initialBalance.String())

	// Send Transaction
	txData := map[string]interface{}{
		"from":     signerAddress,
		"to":       toAddress,
		"value":    "0xDE0B6B3A7640000",  // 1 ETH in Wei
		"gas":      "0x186A0",           // 100k Gas limit
		"gasPrice": "0x4A817C800",       // 20 Gwei
	}
	var txHash string
	err = client.Call(&txHash, "eth_sendTransaction", txData)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	// Get the transaction receipt to confirm it was successful
	var receipt map[string]interface{}
	for {
		err = client.Call(&receipt, "eth_getTransactionReceipt", txHash)
		if err != nil {
			log.Fatalf("Failed to get transaction receipt: %v", err)
		}

		// Check if the receipt is nil, which means the transaction is not yet confirmed
		if receipt != nil {
			break
		}

		// Wait for a short period before trying again
		time.Sleep(1 * time.Second)
	}

	// Check the status of the transaction
	status := receipt["status"].(string)
	if status == "0x1" {
		log.Println("Transaction was successful")
	} else {
		log.Println("Transaction failed")
	}

	// Get final balance as a JSON RawMessage, then convert it to a string
	var finalBalanceRaw json.RawMessage
	err = client.Call(&finalBalanceRaw, "eth_getBalance", toAddress, "latest")
	if err != nil {
		log.Fatalf("Failed to get final balance: %v", err)
	}
	finalBalanceStr := string(finalBalanceRaw)
	finalBalanceStr = strings.Trim(finalBalanceStr, "\"")

	// Convert hex string to big.Int
	finalBalance, success := new(big.Int).SetString(finalBalanceStr[2:], 16)
	if !success {
		log.Fatalf("Failed to parse final balance: %v", err)
	}
	log.Printf("Final balance of toAddress: %s", finalBalance.String())
}
