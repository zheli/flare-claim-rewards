package main

import (
	// "context"
	// "fmt"
	"log"
	// "math/big"
	"os"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/zheli/flare-ftso-v2-reward-claim/pkg/cmd"
	"github.com/zheli/flare-claim-rewards/pkg/cmd/rewardmanager"
	// "github.com/zheli/flare-ftso-v2-reward-claim/pkg/contracts"
)

func main() {
	log.Println("Reward manager CLI. Only claim() is supported currently")

	// pass argument as json file path
	jsonFilePath := os.Args[1]

	err := rewardmanager.Claim(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

}
