package ethservice

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zheli/flare-claim-rewards/pkg/contracts"
)

const (
	rewardManagerAddress = "0xC8f55c5aA2C752eE285Bd872855C749f4ee6239B"
)

// SendClaimTx sends a claim transaction to the reward manager contract
func SendClaimTx(ethClient *ethclient.Client, walletAddress common.Address, walletPK *ecdsa.PrivateKey, rewardRecipient common.Address, rewardClaim contracts.RewardsV2InterfaceRewardClaimWithProof) (error) {
	// Get chain ID
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create auth transaction
	auth, err := bind.NewKeyedTransactorWithChainID(walletPK, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// Create contract instance
	contract, err := contracts.NewRewardManager(common.HexToAddress(rewardManagerAddress), ethClient)
	if err != nil {
		log.Fatal(err)
	}
	
	// Get current nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// Set gas price
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice

	sendWrappedFlare := true

	tx, err := contract.Claim(
		auth, 
		walletAddress, 
		rewardRecipient, 
		rewardClaim.Body.RewardEpochId, 
		sendWrappedFlare, 
		[]contracts.RewardsV2InterfaceRewardClaimWithProof{rewardClaim},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Transaction sent: %s, see https://flare-explorer.flare.network/tx/%s\n", tx.Hash().Hex(), tx.Hash().Hex())

	return nil
}
