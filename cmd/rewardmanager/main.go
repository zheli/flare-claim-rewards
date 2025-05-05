
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zheli/flare-ftso-v2-reward-claim/pkg/contracts"
)

func main() {
	// Connect to Flare network
	client, err := ethclient.Dial("https://flare.public-rpc.com")
	if err != nil {
		log.Fatal(err)
	}

	v1Wallet := os.Getenv("FTSO_V1_WALLET")
	v2Identity := os.Getenv("FTSO_V2_IDENTITY")

	// Contract address for mainnet
	contractAddress := common.HexToAddress("0xC8f55c5aA2C752eE285Bd872855C749f4ee6239B")

	// Get private key from environment variable
	privateKeyStr := os.Getenv("PRIVATE_KEY")
	if privateKeyStr == "" {
		log.Fatal("PRIVATE_KEY environment variable not set")
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create auth transaction
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
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
	contract, err := contracts.NewRewardManager(contractAddress, client)
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
