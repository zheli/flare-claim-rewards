
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to Flare network
	client, err := ethclient.Dial("https://flare.public-rpc.com")
	if err != nil {
		log.Fatal(err)
	}

	// Contract address - replace with your contract address
	contractAddress := common.HexToAddress("YOUR_CONTRACT_ADDRESS")

	// Your private key - replace with your private key
	privateKey := "YOUR_PRIVATE_KEY"
	
	// Create auth transaction
	auth, err := bind.NewTransactorWithChainID(
		strings.NewReader(privateKey),
		password,
		chainID,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Get current nonce
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// Set gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice

	// Create contract instance
	contract, err := NewRewardManager(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Call contract function (example: activate)
	tx, err := contract.Activate(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
}
