package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/zheli/flare-claim-rewards/pkg/cmd/rewardmanager"
)

func main() {
	log.Println("Reward manager CLI. Only claim() is supported currently")

	// pass argument as json file path
	rewardFolder := os.Args[1]
	startEpoch := os.Args[2]
	endEpoch := os.Args[3]

	if rewardFolder == "" || startEpoch == "" || endEpoch == "" {
		fmt.Println("Usage: ./rewardmanager <rewardEpochFolder> <startEpoch> <endEpoch>")
		os.Exit(1)
	}

	startEpochInt, err := strconv.Atoi(startEpoch)
	if err != nil {
		log.Fatal(startEpoch, "is not a valid integer")
	}

	endEpochInt, err := strconv.Atoi(endEpoch)
	if err != nil {
		log.Fatal(endEpoch, "is not a valid integer")
	}

	for epoch := startEpochInt; epoch <= endEpochInt; epoch++ {
		jsonFullPath := filepath.Join(rewardFolder, strconv.Itoa(epoch), "reward-distribution-data.json")
		log.Println("Claiming rewards for epoch", epoch, "from", jsonFullPath)

		if _, err := os.Stat(jsonFullPath); os.IsNotExist(err) {
			log.Fatal(fmt.Sprintf("%s File does not exist for epoch %d. Skip the rest of the epochs.", jsonFullPath, epoch))
			os.Exit(1)
		}

		err = rewardmanager.Claim(jsonFullPath)
		if err != nil {
			log.Println("Error claiming rewards for epoch", epoch, "from", jsonFullPath, "with error:", err, "continue to the next epoch")
		}

		time.Sleep(10 * time.Second)
	}

}
