package parser

import (
	"encoding/json"
	"fmt"
	"math/big"

	eth "github.com/ethereum/go-ethereum/common"

	flareContracts "github.com/zheli/flare-ftso-v2-reward-claim/pkg/contracts"
)

// define input data struct
type RewardDistributionData struct {
	RewardEpochId *big.Int
	Network string
	AppliedMinConditions bool
	RewardClaims []RewardClaim
}

// define reward claim struct
type RewardClaim struct {
	MerkleProof []string
	Body RewardClaimBody
}

// define reward claim body struct
type RewardClaimBody struct {
	Beneficiary eth.Address
	ClaimType uint8
	Amount string
	RewardEpochId *big.Int
}

// ParseRewardDistributionData parses a JSON string into RewardsV2InterfaceRewardClaimWithProof array
func ParseRewardDistributionData(data []byte, ftsoV1Wallet string, ftsoV2Identity string) ([]flareContracts.RewardsV2InterfaceRewardClaimWithProof, error) {
	// Parse JSON data into RewardDistributionData struct
	var rewardData RewardDistributionData
	if err := json.Unmarshal(data, &rewardData); err != nil {
		return nil, fmt.Errorf("failed to parse reward distribution data: %v", err)
	}

	// Filter claims for our addresses
	var filteredClaims []flareContracts.RewardsV2InterfaceRewardClaimWithProof
	for _, claim := range rewardData.RewardClaims {
		if claim.Body.Beneficiary == eth.HexToAddress(ftsoV1Wallet) || claim.Body.Beneficiary == eth.HexToAddress(ftsoV2Identity) {
			merkleProof := make([][32]byte, 8)
			for i, proof := range claim.MerkleProof {
				merkleProof[i] = eth.HexToHash(proof)
			}

			amount, ok := new(big.Int).SetString(claim.Body.Amount, 10)
			if !ok {
				return nil, fmt.Errorf("failed to parse amount: %v", claim.Body.Amount)
			}

			filteredClaims = append(filteredClaims, flareContracts.RewardsV2InterfaceRewardClaimWithProof{
				MerkleProof: merkleProof,
				Body: flareContracts.RewardsV2InterfaceRewardClaim{
					Beneficiary: [20]byte(claim.Body.Beneficiary),
					ClaimType: claim.Body.ClaimType,
					Amount: amount,
					RewardEpochId: claim.Body.RewardEpochId,
				},
			})
		}
	}

	return filteredClaims, nil
}