package rewardmanager

import (
	"log"
	"os"

	eth "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"github.com/zheli/flare-claim-rewards/pkg/ethservice"
	"github.com/zheli/flare-claim-rewards/pkg/parser"
)

// claim command
func Claim(jsonFilePath string) (error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// read from .env file
	v1Wallet := os.Getenv("FTSO_V1_WALLET")
	v2Identity := os.Getenv("FTSO_V2_IDENTITY")
	if v1Wallet == "" || v2Identity == "" {
		log.Fatal("FTSO_V1_WALLET or FTSO_V2_IDENTITY environment variable not set")
	}
	log.Println("V1 wallet:", v1Wallet)
	log.Println("V2 identity:", v2Identity)

	// Get private key from environment variable
	privateKeyStr := os.Getenv("FTSO_V2_PRIVATE_KEY")
	if privateKeyStr == "" {
		log.Fatal("FTSO_V2_PRIVATE_KEY environment variable not set")
	}

	rewardRecipient := os.Getenv("REWARD_RECIPIENT")
	if rewardRecipient == "" {
		log.Fatal("REWARD_RECIPIENT environment variable not set")
	}

	log.Println("Claiming rewards from", jsonFilePath)
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()


	rewardClaims, err := parser.ParseRewardDistributionData(jsonFile, v1Wallet, v2Identity)
	if err != nil {
		return err
	}

	// print all reward claims
	for i, claim := range rewardClaims {
		// there is 18 decimals in the amount, convert it to human readable format by dividing by 10^18 and keep all numbers after the decimal point
		amountStr := decimal.NewFromBigInt(claim.Body.Amount, -18).String()
		log.Printf("Claim %d: Beneficiary: 0x%x, Reward Amount: %s WFLRs", i, claim.Body.Beneficiary[:], amountStr)
	}

	// Connect to Flare network
	client, err := ethclient.Dial("https://flare.public-rpc.com")
	if err != nil {
		log.Fatal(err)
	}

	err = ethservice.SendClaimTx(client, privateKeyStr, eth.HexToAddress(rewardRecipient), rewardClaims)
	if err != nil {
		return err
	}

	return nil
}
