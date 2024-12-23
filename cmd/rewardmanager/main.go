package main

import (
	"log"
	"os"
	"github.com/zheli/flare-claim-rewards/pkg/cmd/rewardmanager"
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
