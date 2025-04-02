package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/zheli/flare-claim-rewards/pkg/cmd/rewardmanager"
)

func main() {
	log.Println("Reward manager CLI. Only claim() is supported currently")

	// pass argument as json file path
	rewardEpochFolder := os.Args[1]

	jsonFullPath := filepath.Join(rewardEpochFolder, "reward-distribution-data.json")

	err := rewardmanager.Claim(jsonFullPath)
	if err != nil {
		log.Fatal(err)
	}

}
