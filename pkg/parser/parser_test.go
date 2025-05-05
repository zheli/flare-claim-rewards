package parser

import (
	"math/big"
	"testing"

	eth "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

const (
	testFTOSV1Wallet   = "0x0000000000000000000000000000000000000001"
	testFTOSV2Identity = "0x0000000000000000000000000000000000000002"
)

func TestParseRewardDistributionData(t *testing.T) {
	// Test input JSON
	testJSON := []byte(`{
		"rewardEpochId": 251,
		"network": "flare",
		"appliedMinConditions": false,
		"rewardClaims": [
			{
				"merkleProof": [
					"0x2e7ea06977a109200b35d190631c212b7676faf81bbfe5a9a4f16498630fc89e",
					"0x6f87e0eaf8ba5a76120f37a80bec978ad70176f853e198112ab77fa6194a1f92",
					"0x541eac1ad2b6e616ef4782873625e8e688389b9d627364e2866357a7647c41d0",
					"0x5ff2da4880e74d772d1057e2fb9494673727c9eae63b912b6236d008713c06f4",
					"0xa9d331a8f0deae8cdfaff4f97cf85197b924a400696a2ff4b0e57995e05ae393",
					"0x3f1d2b61fca9ca33743bddea3f21e6e384e353fad136e4b4d60d510062e98be4",
					"0x031e07638f27407831411a84b45466fd936d09315c71decddc9e92b1ba14bc69",
					"0xa05b5f8ef7ef0f9d6e0f0b6d47ba851fbd2de6193cde441c8ea757349312e16f"
				],
				"body": {
					"beneficiary": "0x0000000000000000000000000000000000000001",
					"claimType": 0,
					"amount": "403643555586096443338936",
					"rewardEpochId": 251
				}
			},
			{
				"merkleProof": [
					"0x2222222222222222222222222222222222222222222222222222222222222222",
					"0x3333333333333333333333333333333333333333333333333333333333332222",
					"0x4444444444444444444444444444444444444444444444444444444444442222",
					"0x5555555555555555555555555555555555555555555555555555555555552222",
					"0x6666666666666666666666666666666666666666666666666666666666662222",
					"0x7777777777777777777777777777777777777777777777777777777777772222",
					"0x8888888888888888888888888888888888888888888888888888888888882222",
					"0x9999999999999999999999999999999999999999999999999999999999992222"
				],
				"body": {
					"beneficiary": "0x0000000000000000000000000000000000000002",
					"claimType": 0,
					"amount": "100000000000000000000",
					"rewardEpochId": 251
				}
			},
			{
				"merkleProof": [
					"0x1111111111111111111111111111111111111111111111111111111111113333",
					"0x2222222222222222222222222222222222222222222222222222222222223333",
					"0x3333333333333333333333333333333333333333333333333333333333333333",
					"0x4444444444444444444444444444444444444444444444444444444444443333",
					"0x5555555555555555555555555555555555555555555555555555555555553333",
					"0x6666666666666666666666666666666666666666666666666666666666663333",
					"0x7777777777777777777777777777777777777777777777777777777777773333",
					"0x8888888888888888888888888888888888888888888888888888888888883333"
				],
				"body": {
					"beneficiary": "0x000000000000000000000000000000000000dead",
					"claimType": 0,
					"amount": "200000000000000000000",
					"rewardEpochId": 251
				}
			}
		]
	}`)

	// Parse the test data
	claims, err := ParseRewardDistributionData(testJSON, testFTOSV1Wallet, testFTOSV2Identity)

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert we got exactly one claim (for ftsoV1Wallet)
	assert.Equal(t, 2, len(claims))

	// Check the parsed claim details
	expectedAmount := new(big.Int)
	expectedAmount.SetString("403643555586096443338936", 10)
	expectedEpochID := big.NewInt(251)

	claim := claims[0]
	assert.Equal(t, [20]byte(eth.HexToAddress(testFTOSV1Wallet)), claim.Body.Beneficiary)
	assert.Equal(t, uint8(0), claim.Body.ClaimType)
	assert.Equal(t, expectedAmount, claim.Body.Amount)
	assert.Equal(t, expectedEpochID, claim.Body.RewardEpochId)

	// Check first and last merkle proof entries
	assert.Equal(t, [32]byte(eth.HexToHash("0x2e7ea06977a109200b35d190631c212b7676faf81bbfe5a9a4f16498630fc89e")), claim.MerkleProof[0])
	assert.Equal(t, [32]byte(eth.HexToHash("0xa05b5f8ef7ef0f9d6e0f0b6d47ba851fbd2de6193cde441c8ea757349312e16f")), claim.MerkleProof[7])
}

func TestParseRewardDistributionData_InvalidJSON(t *testing.T) {
	// Test with invalid JSON
	invalidJSON := []byte(`{invalid json}`)
	claims, err := ParseRewardDistributionData(invalidJSON, testFTOSV1Wallet, testFTOSV2Identity)

	// Assert error occurred and claims is nil
	assert.Error(t, err)
	assert.Nil(t, claims)
} 