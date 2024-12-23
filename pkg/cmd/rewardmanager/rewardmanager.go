package rewardmanager

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zheli/flare-claim-rewards/pkg/parser"
)

func Claim(jsonFilePath string) (error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	log.Println("Claiming rewards from", jsonFilePath)
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// read from .env file
	v1Wallet := os.Getenv("FTSO_V1_WALLET")
	v2Identity := os.Getenv("FTSO_V2_IDENTITY")
	log.Println("V1 wallet:", v1Wallet)
	log.Println("V2 identity:", v2Identity)

	rewardClaims, err := parser.ParseRewardDistributionData(jsonFile, v1Wallet, v2Identity)
	if err != nil {
		return err
	}

	// print all reward claims
	for i, claim := range rewardClaims {
		log.Printf("Claim %d: Beneficiary: 0x%x", i, claim.Body.Beneficiary[:])
	}

	return nil
}

