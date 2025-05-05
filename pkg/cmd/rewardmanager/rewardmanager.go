package rewardmanager

import (
	"crypto/ecdsa"
	"log"
	"os"

	eth "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	// Get private key and reward recipient from environment variable
	v1PrivateKeyStr := os.Getenv("FTSO_V1_PRIVATE_KEY")
	v2PrivateKeyStr := os.Getenv("FTSO_V2_PRIVATE_KEY")
	rewardRecipient := os.Getenv("REWARD_RECIPIENT")
	if v1PrivateKeyStr == "" || v2PrivateKeyStr == "" || rewardRecipient == "" {
		log.Fatal("FTSO_V1_PRIVATE_KEY, FTSO_V2_PRIVATE_KEY, or REWARD_RECIPIENT environment variable not set")
	}

	v1PrivateKey, v1WalletAddress, err := getPrivateKeyAndWalletAddress(v1PrivateKeyStr)
	if err != nil {
		log.Fatal("Invalid private key for v1 wallet:", err)
	}
	v2PrivateKey, v2WalletAddress, err := getPrivateKeyAndWalletAddress(v2PrivateKeyStr)
	if err != nil {
		log.Fatal("Invalid private key for v2 wallet:", err)
	}

	log.Println("Claiming rewards from", jsonFilePath)
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()


	rewardClaims, err := parser.ParseRewardDistributionData(jsonFile, v1WalletAddress, v2WalletAddress, true)
	if err != nil {
		return err
	}

	// Connect to Flare network
	client, err := ethclient.Dial("https://flare.public-rpc.com")
	if err != nil {
		log.Fatal(err)
	}

	for i, claim := range rewardClaims {
		// there is 18 decimals in the amount, convert it to human readable format by dividing by 10^18 and keep all numbers after the decimal point
		amountStr := decimal.NewFromBigInt(claim.Body.Amount, -18).String()
		log.Printf("Claim %d: Beneficiary: 0x%x, Reward Amount: %s WFLRs, Reward Epoch ID: %d, Reward Type: %d", i, claim.Body.Beneficiary[:], amountStr, claim.Body.RewardEpochId.Int64(), claim.Body.ClaimType)

		if claim.Body.Beneficiary == v1WalletAddress {
			log.Printf("Claiming for v1 wallet: %s", v1WalletAddress.Hex())
			err = ethservice.SendClaimTx(client, v1WalletAddress, v1PrivateKey, eth.HexToAddress(rewardRecipient), claim)
		} else if claim.Body.Beneficiary == v2WalletAddress {
			log.Printf("Claiming for v2 wallet: %s", v2WalletAddress.Hex())
			err = ethservice.SendClaimTx(client, v2WalletAddress, v2PrivateKey, eth.HexToAddress(rewardRecipient), claim)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func getPrivateKeyAndWalletAddress(privateKeyStr string) (*ecdsa.PrivateKey, eth.Address, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, eth.Address{}, err
	}
	return privateKey, crypto.PubkeyToAddress(privateKey.PublicKey), nil
}