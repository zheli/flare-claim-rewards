package ethservice

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zheli/flare-claim-rewards/pkg/contracts"
)

const (
	rewardManagerAddress = "0xC8f55c5aA2C752eE285Bd872855C749f4ee6239B"
)

// SendClaimTx sends a claim transaction to the reward manager contract
func SendClaimTx(ethClient *ethclient.Client, v2IdentityPKStr string, rewardRecipient common.Address, rewardClaims []contracts.RewardsV2InterfaceRewardClaimWithProof) (error) {
	// Parse private key
	v2IdentityPK, err := crypto.HexToECDSA(v2IdentityPKStr)
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// Get chain ID
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create auth transaction
	auth, err := bind.NewKeyedTransactorWithChainID(v2IdentityPK, chainID)
	if err != nil {
		log.Fatal(err)
	}


	// Create contract instance
	contract, err := contracts.NewRewardManager(common.HexToAddress(rewardManagerAddress), ethClient)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, claim := range rewardClaims {
		// Create CallOpts for simulation
		transactOpts := &bind.TransactOpts{
			From: auth.From,
			Context: context.Background(),
			NoSend: true,
		}

		// // Get current nonce
		// nonce, err := ethClient.PendingNonceAt(context.Background(), auth.From)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// auth.Nonce = big.NewInt(int64(nonce))

		// // Set gas price
		// gasPrice, err := ethClient.SuggestGasPrice(context.Background())
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// auth.GasPrice = gasPrice

		wrapToken := true

		tx, err := contract.Claim(
			transactOpts, 
			claim.Body.Beneficiary, 
			rewardRecipient, 
			claim.Body.RewardEpochId, 
			wrapToken, 
			rewardClaims,
		)
		// tx, err := contract.Claim(auth, claim.Body.Beneficiary, claim.Body.Beneficiary, claim.Body.RewardEpochId, true, rewardClaims)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	}

	return nil
}
