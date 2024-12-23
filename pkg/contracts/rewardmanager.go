// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IRewardManagerUnclaimedRewardState is an auto generated low-level Go binding around an user-defined struct.
type IRewardManagerUnclaimedRewardState struct {
	Initialised bool
	Amount      *big.Int
	Weight      *big.Int
}

// RewardsV2InterfaceRewardClaim is an auto generated low-level Go binding around an user-defined struct.
type RewardsV2InterfaceRewardClaim struct {
	RewardEpochId *big.Int
	Beneficiary   [20]byte
	Amount        *big.Int
	ClaimType     uint8
}

// RewardsV2InterfaceRewardClaimWithProof is an auto generated low-level Go binding around an user-defined struct.
type RewardsV2InterfaceRewardClaimWithProof struct {
	MerkleProof [][32]byte
	Body        RewardsV2InterfaceRewardClaim
}

// RewardsV2InterfaceRewardState is an auto generated low-level Go binding around an user-defined struct.
type RewardsV2InterfaceRewardState struct {
	RewardEpochId *big.Int
	Beneficiary   [20]byte
	Amount        *big.Int
	ClaimType     uint8
	Initialised   bool
}

// RewardManagerMetaData contains all meta data concerning the RewardManager contract.
var RewardManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oldRewardManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardManagerId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"enumRewardsV2Interface.ClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint120\",\"name\":\"amount\",\"type\":\"uint120\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"RewardClaimsEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"RewardClaimsExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_rewardOwners\",\"type\":\"address[]\"},{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20\",\"name\":\"beneficiary\",\"type\":\"bytes20\"},{\"internalType\":\"uint120\",\"name\":\"amount\",\"type\":\"uint120\"},{\"internalType\":\"enumRewardsV2Interface.ClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"}],\"internalType\":\"structRewardsV2Interface.RewardClaim\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structRewardsV2Interface.RewardClaimWithProof[]\",\"name\":\"_proofs\",\"type\":\"tuple[]\"}],\"name\":\"autoClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cChainStake\",\"outputs\":[{\"internalType\":\"contractICChainStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cChainStakeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardOwner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"_wrap\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20\",\"name\":\"beneficiary\",\"type\":\"bytes20\"},{\"internalType\":\"uint120\",\"name\":\"amount\",\"type\":\"uint120\"},{\"internalType\":\"enumRewardsV2Interface.ClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"}],\"internalType\":\"structRewardsV2Interface.RewardClaim\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structRewardsV2Interface.RewardClaimWithProof[]\",\"name\":\"_proofs\",\"type\":\"tuple[]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardAmountWei\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardOwner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"_wrap\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20\",\"name\":\"beneficiary\",\"type\":\"bytes20\"},{\"internalType\":\"uint120\",\"name\":\"amount\",\"type\":\"uint120\"},{\"internalType\":\"enumRewardsV2Interface.ClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"}],\"internalType\":\"structRewardsV2Interface.RewardClaim\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structRewardsV2Interface.RewardClaimWithProof[]\",\"name\":\"_proofs\",\"type\":\"tuple[]\"}],\"name\":\"claimProxy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardAmountWei\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimSetupManager\",\"outputs\":[{\"internalType\":\"contractIIClaimSetupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanupBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"closeExpiredRewardEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableCChainStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enablePChainStakeMirror\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstClaimableRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsCalculator\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsCalculator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftsoRewardManagerProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialRewardEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardOwner\",\"type\":\"address\"}],\"name\":\"getNextClaimableRewardEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardEpochIdToExpireNext\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardEpochIdsWithClaimableRewards\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_startEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"_endEpochId\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRewardEpochTotals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalRewardsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalInflationRewardsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialisedRewardsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_claimedRewardsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_burnedRewardsWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardOffersManagerList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardOwner\",\"type\":\"address\"}],\"name\":\"getStateOfRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20\",\"name\":\"beneficiary\",\"type\":\"bytes20\"},{\"internalType\":\"uint120\",\"name\":\"amount\",\"type\":\"uint120\"},{\"internalType\":\"enumRewardsV2Interface.ClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"initialised\",\"type\":\"bool\"}],\"internalType\":\"structRewardsV2Interface.RewardState[][]\",\"name\":\"_rewardStates\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardOwner\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getStateOfRewardsAt\",\"outputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20\",\"name\":\"beneficiary\",\"type\":\"bytes20\"},{\"internalType\":\"uint120\",\"name\":\"amount\",\"type\":\"uint120\"},{\"internalType\":\"enumRewardsV2Interface.ClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"initialised\",\"type\":\"bool\"}],\"internalType\":\"structRewardsV2Interface.RewardState[]\",\"name\":\"_rewardStates\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPoolSupplyData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockedFundsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalInflationAuthorizedWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalClaimedWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalRewardsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalInflationRewardsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalClaimedWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBurnedWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"enumRewardsV2Interface.ClaimType\",\"name\":\"_claimType\",\"type\":\"uint8\"}],\"name\":\"getUnclaimedRewardState\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"initialised\",\"type\":\"bool\"},{\"internalType\":\"uint120\",\"name\":\"amount\",\"type\":\"uint120\"},{\"internalType\":\"uint128\",\"name\":\"weight\",\"type\":\"uint128\"}],\"internalType\":\"structIRewardManager.UnclaimedRewardState\",\"name\":\"_state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20\",\"name\":\"beneficiary\",\"type\":\"bytes20\"},{\"internalType\":\"uint120\",\"name\":\"amount\",\"type\":\"uint120\"},{\"internalType\":\"enumRewardsV2Interface.ClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"}],\"internalType\":\"structRewardsV2Interface.RewardClaim\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structRewardsV2Interface.RewardClaimWithProof[]\",\"name\":\"_proofs\",\"type\":\"tuple[]\"}],\"name\":\"initialiseWeightBasedClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newRewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"noOfInitialisedWeightBasedClaims\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldRewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainStakeMirror\",\"outputs\":[{\"internalType\":\"contractIPChainStakeMirror\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainStakeMirrorEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"_inflation\",\"type\":\"bool\"}],\"name\":\"receiveRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManagerId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setInitialRewardData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRewardManager\",\"type\":\"address\"}],\"name\":\"setNewRewardManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_rewardOffersManagerList\",\"type\":\"address[]\"}],\"name\":\"setRewardOffersManagerList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wNat\",\"outputs\":[{\"internalType\":\"contractIWNat\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RewardManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardManagerMetaData.ABI instead.
var RewardManagerABI = RewardManagerMetaData.ABI

// RewardManager is an auto generated Go binding around an Ethereum contract.
type RewardManager struct {
	RewardManagerCaller     // Read-only binding to the contract
	RewardManagerTransactor // Write-only binding to the contract
	RewardManagerFilterer   // Log filterer for contract events
}

// RewardManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardManagerSession struct {
	Contract     *RewardManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardManagerCallerSession struct {
	Contract *RewardManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RewardManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardManagerTransactorSession struct {
	Contract     *RewardManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RewardManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardManagerRaw struct {
	Contract *RewardManager // Generic contract binding to access the raw methods on
}

// RewardManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardManagerCallerRaw struct {
	Contract *RewardManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RewardManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardManagerTransactorRaw struct {
	Contract *RewardManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardManager creates a new instance of RewardManager, bound to a specific deployed contract.
func NewRewardManager(address common.Address, backend bind.ContractBackend) (*RewardManager, error) {
	contract, err := bindRewardManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardManager{RewardManagerCaller: RewardManagerCaller{contract: contract}, RewardManagerTransactor: RewardManagerTransactor{contract: contract}, RewardManagerFilterer: RewardManagerFilterer{contract: contract}}, nil
}

// NewRewardManagerCaller creates a new read-only instance of RewardManager, bound to a specific deployed contract.
func NewRewardManagerCaller(address common.Address, caller bind.ContractCaller) (*RewardManagerCaller, error) {
	contract, err := bindRewardManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardManagerCaller{contract: contract}, nil
}

// NewRewardManagerTransactor creates a new write-only instance of RewardManager, bound to a specific deployed contract.
func NewRewardManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardManagerTransactor, error) {
	contract, err := bindRewardManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardManagerTransactor{contract: contract}, nil
}

// NewRewardManagerFilterer creates a new log filterer instance of RewardManager, bound to a specific deployed contract.
func NewRewardManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardManagerFilterer, error) {
	contract, err := bindRewardManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardManagerFilterer{contract: contract}, nil
}

// bindRewardManager binds a generic wrapper to an already deployed contract.
func bindRewardManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardManager *RewardManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardManager.Contract.RewardManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardManager *RewardManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.Contract.RewardManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardManager *RewardManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardManager.Contract.RewardManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardManager *RewardManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardManager *RewardManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardManager *RewardManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardManager.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_RewardManager *RewardManagerCaller) Active(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "active")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_RewardManager *RewardManagerSession) Active() (bool, error) {
	return _RewardManager.Contract.Active(&_RewardManager.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_RewardManager *RewardManagerCallerSession) Active() (bool, error) {
	return _RewardManager.Contract.Active(&_RewardManager.CallOpts)
}

// CChainStake is a free data retrieval call binding the contract method 0xe7dea8e6.
//
// Solidity: function cChainStake() view returns(address)
func (_RewardManager *RewardManagerCaller) CChainStake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "cChainStake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CChainStake is a free data retrieval call binding the contract method 0xe7dea8e6.
//
// Solidity: function cChainStake() view returns(address)
func (_RewardManager *RewardManagerSession) CChainStake() (common.Address, error) {
	return _RewardManager.Contract.CChainStake(&_RewardManager.CallOpts)
}

// CChainStake is a free data retrieval call binding the contract method 0xe7dea8e6.
//
// Solidity: function cChainStake() view returns(address)
func (_RewardManager *RewardManagerCallerSession) CChainStake() (common.Address, error) {
	return _RewardManager.Contract.CChainStake(&_RewardManager.CallOpts)
}

// CChainStakeEnabled is a free data retrieval call binding the contract method 0x064be532.
//
// Solidity: function cChainStakeEnabled() view returns(bool)
func (_RewardManager *RewardManagerCaller) CChainStakeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "cChainStakeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CChainStakeEnabled is a free data retrieval call binding the contract method 0x064be532.
//
// Solidity: function cChainStakeEnabled() view returns(bool)
func (_RewardManager *RewardManagerSession) CChainStakeEnabled() (bool, error) {
	return _RewardManager.Contract.CChainStakeEnabled(&_RewardManager.CallOpts)
}

// CChainStakeEnabled is a free data retrieval call binding the contract method 0x064be532.
//
// Solidity: function cChainStakeEnabled() view returns(bool)
func (_RewardManager *RewardManagerCallerSession) CChainStakeEnabled() (bool, error) {
	return _RewardManager.Contract.CChainStakeEnabled(&_RewardManager.CallOpts)
}

// ClaimSetupManager is a free data retrieval call binding the contract method 0xc4db9619.
//
// Solidity: function claimSetupManager() view returns(address)
func (_RewardManager *RewardManagerCaller) ClaimSetupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "claimSetupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimSetupManager is a free data retrieval call binding the contract method 0xc4db9619.
//
// Solidity: function claimSetupManager() view returns(address)
func (_RewardManager *RewardManagerSession) ClaimSetupManager() (common.Address, error) {
	return _RewardManager.Contract.ClaimSetupManager(&_RewardManager.CallOpts)
}

// ClaimSetupManager is a free data retrieval call binding the contract method 0xc4db9619.
//
// Solidity: function claimSetupManager() view returns(address)
func (_RewardManager *RewardManagerCallerSession) ClaimSetupManager() (common.Address, error) {
	return _RewardManager.Contract.ClaimSetupManager(&_RewardManager.CallOpts)
}

// CleanupBlockNumber is a free data retrieval call binding the contract method 0xdeea13e7.
//
// Solidity: function cleanupBlockNumber() view returns(uint256)
func (_RewardManager *RewardManagerCaller) CleanupBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "cleanupBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CleanupBlockNumber is a free data retrieval call binding the contract method 0xdeea13e7.
//
// Solidity: function cleanupBlockNumber() view returns(uint256)
func (_RewardManager *RewardManagerSession) CleanupBlockNumber() (*big.Int, error) {
	return _RewardManager.Contract.CleanupBlockNumber(&_RewardManager.CallOpts)
}

// CleanupBlockNumber is a free data retrieval call binding the contract method 0xdeea13e7.
//
// Solidity: function cleanupBlockNumber() view returns(uint256)
func (_RewardManager *RewardManagerCallerSession) CleanupBlockNumber() (*big.Int, error) {
	return _RewardManager.Contract.CleanupBlockNumber(&_RewardManager.CallOpts)
}

// FirstClaimableRewardEpochId is a free data retrieval call binding the contract method 0x9a3410bc.
//
// Solidity: function firstClaimableRewardEpochId() view returns(uint24)
func (_RewardManager *RewardManagerCaller) FirstClaimableRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "firstClaimableRewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstClaimableRewardEpochId is a free data retrieval call binding the contract method 0x9a3410bc.
//
// Solidity: function firstClaimableRewardEpochId() view returns(uint24)
func (_RewardManager *RewardManagerSession) FirstClaimableRewardEpochId() (*big.Int, error) {
	return _RewardManager.Contract.FirstClaimableRewardEpochId(&_RewardManager.CallOpts)
}

// FirstClaimableRewardEpochId is a free data retrieval call binding the contract method 0x9a3410bc.
//
// Solidity: function firstClaimableRewardEpochId() view returns(uint24)
func (_RewardManager *RewardManagerCallerSession) FirstClaimableRewardEpochId() (*big.Int, error) {
	return _RewardManager.Contract.FirstClaimableRewardEpochId(&_RewardManager.CallOpts)
}

// FlareSystemsCalculator is a free data retrieval call binding the contract method 0x8e467784.
//
// Solidity: function flareSystemsCalculator() view returns(address)
func (_RewardManager *RewardManagerCaller) FlareSystemsCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "flareSystemsCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsCalculator is a free data retrieval call binding the contract method 0x8e467784.
//
// Solidity: function flareSystemsCalculator() view returns(address)
func (_RewardManager *RewardManagerSession) FlareSystemsCalculator() (common.Address, error) {
	return _RewardManager.Contract.FlareSystemsCalculator(&_RewardManager.CallOpts)
}

// FlareSystemsCalculator is a free data retrieval call binding the contract method 0x8e467784.
//
// Solidity: function flareSystemsCalculator() view returns(address)
func (_RewardManager *RewardManagerCallerSession) FlareSystemsCalculator() (common.Address, error) {
	return _RewardManager.Contract.FlareSystemsCalculator(&_RewardManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_RewardManager *RewardManagerCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_RewardManager *RewardManagerSession) FlareSystemsManager() (common.Address, error) {
	return _RewardManager.Contract.FlareSystemsManager(&_RewardManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_RewardManager *RewardManagerCallerSession) FlareSystemsManager() (common.Address, error) {
	return _RewardManager.Contract.FlareSystemsManager(&_RewardManager.CallOpts)
}

// FtsoRewardManagerProxy is a free data retrieval call binding the contract method 0x80b24173.
//
// Solidity: function ftsoRewardManagerProxy() view returns(address)
func (_RewardManager *RewardManagerCaller) FtsoRewardManagerProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "ftsoRewardManagerProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FtsoRewardManagerProxy is a free data retrieval call binding the contract method 0x80b24173.
//
// Solidity: function ftsoRewardManagerProxy() view returns(address)
func (_RewardManager *RewardManagerSession) FtsoRewardManagerProxy() (common.Address, error) {
	return _RewardManager.Contract.FtsoRewardManagerProxy(&_RewardManager.CallOpts)
}

// FtsoRewardManagerProxy is a free data retrieval call binding the contract method 0x80b24173.
//
// Solidity: function ftsoRewardManagerProxy() view returns(address)
func (_RewardManager *RewardManagerCallerSession) FtsoRewardManagerProxy() (common.Address, error) {
	return _RewardManager.Contract.FtsoRewardManagerProxy(&_RewardManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_RewardManager *RewardManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_RewardManager *RewardManagerSession) GetAddressUpdater() (common.Address, error) {
	return _RewardManager.Contract.GetAddressUpdater(&_RewardManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_RewardManager *RewardManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _RewardManager.Contract.GetAddressUpdater(&_RewardManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_RewardManager *RewardManagerCaller) GetCurrentRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getCurrentRewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_RewardManager *RewardManagerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _RewardManager.Contract.GetCurrentRewardEpochId(&_RewardManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_RewardManager *RewardManagerCallerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _RewardManager.Contract.GetCurrentRewardEpochId(&_RewardManager.CallOpts)
}

// GetExpectedBalance is a free data retrieval call binding the contract method 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (_RewardManager *RewardManagerCaller) GetExpectedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getExpectedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedBalance is a free data retrieval call binding the contract method 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (_RewardManager *RewardManagerSession) GetExpectedBalance() (*big.Int, error) {
	return _RewardManager.Contract.GetExpectedBalance(&_RewardManager.CallOpts)
}

// GetExpectedBalance is a free data retrieval call binding the contract method 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (_RewardManager *RewardManagerCallerSession) GetExpectedBalance() (*big.Int, error) {
	return _RewardManager.Contract.GetExpectedBalance(&_RewardManager.CallOpts)
}

// GetInitialRewardEpochId is a free data retrieval call binding the contract method 0x30c40bc9.
//
// Solidity: function getInitialRewardEpochId() view returns(uint256)
func (_RewardManager *RewardManagerCaller) GetInitialRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getInitialRewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialRewardEpochId is a free data retrieval call binding the contract method 0x30c40bc9.
//
// Solidity: function getInitialRewardEpochId() view returns(uint256)
func (_RewardManager *RewardManagerSession) GetInitialRewardEpochId() (*big.Int, error) {
	return _RewardManager.Contract.GetInitialRewardEpochId(&_RewardManager.CallOpts)
}

// GetInitialRewardEpochId is a free data retrieval call binding the contract method 0x30c40bc9.
//
// Solidity: function getInitialRewardEpochId() view returns(uint256)
func (_RewardManager *RewardManagerCallerSession) GetInitialRewardEpochId() (*big.Int, error) {
	return _RewardManager.Contract.GetInitialRewardEpochId(&_RewardManager.CallOpts)
}

// GetNextClaimableRewardEpochId is a free data retrieval call binding the contract method 0xd6ac4f72.
//
// Solidity: function getNextClaimableRewardEpochId(address _rewardOwner) view returns(uint256)
func (_RewardManager *RewardManagerCaller) GetNextClaimableRewardEpochId(opts *bind.CallOpts, _rewardOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getNextClaimableRewardEpochId", _rewardOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextClaimableRewardEpochId is a free data retrieval call binding the contract method 0xd6ac4f72.
//
// Solidity: function getNextClaimableRewardEpochId(address _rewardOwner) view returns(uint256)
func (_RewardManager *RewardManagerSession) GetNextClaimableRewardEpochId(_rewardOwner common.Address) (*big.Int, error) {
	return _RewardManager.Contract.GetNextClaimableRewardEpochId(&_RewardManager.CallOpts, _rewardOwner)
}

// GetNextClaimableRewardEpochId is a free data retrieval call binding the contract method 0xd6ac4f72.
//
// Solidity: function getNextClaimableRewardEpochId(address _rewardOwner) view returns(uint256)
func (_RewardManager *RewardManagerCallerSession) GetNextClaimableRewardEpochId(_rewardOwner common.Address) (*big.Int, error) {
	return _RewardManager.Contract.GetNextClaimableRewardEpochId(&_RewardManager.CallOpts, _rewardOwner)
}

// GetRewardEpochIdToExpireNext is a free data retrieval call binding the contract method 0xa71de676.
//
// Solidity: function getRewardEpochIdToExpireNext() view returns(uint256)
func (_RewardManager *RewardManagerCaller) GetRewardEpochIdToExpireNext(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getRewardEpochIdToExpireNext")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardEpochIdToExpireNext is a free data retrieval call binding the contract method 0xa71de676.
//
// Solidity: function getRewardEpochIdToExpireNext() view returns(uint256)
func (_RewardManager *RewardManagerSession) GetRewardEpochIdToExpireNext() (*big.Int, error) {
	return _RewardManager.Contract.GetRewardEpochIdToExpireNext(&_RewardManager.CallOpts)
}

// GetRewardEpochIdToExpireNext is a free data retrieval call binding the contract method 0xa71de676.
//
// Solidity: function getRewardEpochIdToExpireNext() view returns(uint256)
func (_RewardManager *RewardManagerCallerSession) GetRewardEpochIdToExpireNext() (*big.Int, error) {
	return _RewardManager.Contract.GetRewardEpochIdToExpireNext(&_RewardManager.CallOpts)
}

// GetRewardEpochIdsWithClaimableRewards is a free data retrieval call binding the contract method 0xd8def818.
//
// Solidity: function getRewardEpochIdsWithClaimableRewards() view returns(uint24 _startEpochId, uint24 _endEpochId)
func (_RewardManager *RewardManagerCaller) GetRewardEpochIdsWithClaimableRewards(opts *bind.CallOpts) (struct {
	StartEpochId *big.Int
	EndEpochId   *big.Int
}, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getRewardEpochIdsWithClaimableRewards")

	outstruct := new(struct {
		StartEpochId *big.Int
		EndEpochId   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartEpochId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndEpochId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRewardEpochIdsWithClaimableRewards is a free data retrieval call binding the contract method 0xd8def818.
//
// Solidity: function getRewardEpochIdsWithClaimableRewards() view returns(uint24 _startEpochId, uint24 _endEpochId)
func (_RewardManager *RewardManagerSession) GetRewardEpochIdsWithClaimableRewards() (struct {
	StartEpochId *big.Int
	EndEpochId   *big.Int
}, error) {
	return _RewardManager.Contract.GetRewardEpochIdsWithClaimableRewards(&_RewardManager.CallOpts)
}

// GetRewardEpochIdsWithClaimableRewards is a free data retrieval call binding the contract method 0xd8def818.
//
// Solidity: function getRewardEpochIdsWithClaimableRewards() view returns(uint24 _startEpochId, uint24 _endEpochId)
func (_RewardManager *RewardManagerCallerSession) GetRewardEpochIdsWithClaimableRewards() (struct {
	StartEpochId *big.Int
	EndEpochId   *big.Int
}, error) {
	return _RewardManager.Contract.GetRewardEpochIdsWithClaimableRewards(&_RewardManager.CallOpts)
}

// GetRewardEpochTotals is a free data retrieval call binding the contract method 0xdf339638.
//
// Solidity: function getRewardEpochTotals(uint24 _rewardEpochId) view returns(uint256 _totalRewardsWei, uint256 _totalInflationRewardsWei, uint256 _initialisedRewardsWei, uint256 _claimedRewardsWei, uint256 _burnedRewardsWei)
func (_RewardManager *RewardManagerCaller) GetRewardEpochTotals(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	TotalRewardsWei          *big.Int
	TotalInflationRewardsWei *big.Int
	InitialisedRewardsWei    *big.Int
	ClaimedRewardsWei        *big.Int
	BurnedRewardsWei         *big.Int
}, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getRewardEpochTotals", _rewardEpochId)

	outstruct := new(struct {
		TotalRewardsWei          *big.Int
		TotalInflationRewardsWei *big.Int
		InitialisedRewardsWei    *big.Int
		ClaimedRewardsWei        *big.Int
		BurnedRewardsWei         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalRewardsWei = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalInflationRewardsWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitialisedRewardsWei = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ClaimedRewardsWei = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BurnedRewardsWei = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRewardEpochTotals is a free data retrieval call binding the contract method 0xdf339638.
//
// Solidity: function getRewardEpochTotals(uint24 _rewardEpochId) view returns(uint256 _totalRewardsWei, uint256 _totalInflationRewardsWei, uint256 _initialisedRewardsWei, uint256 _claimedRewardsWei, uint256 _burnedRewardsWei)
func (_RewardManager *RewardManagerSession) GetRewardEpochTotals(_rewardEpochId *big.Int) (struct {
	TotalRewardsWei          *big.Int
	TotalInflationRewardsWei *big.Int
	InitialisedRewardsWei    *big.Int
	ClaimedRewardsWei        *big.Int
	BurnedRewardsWei         *big.Int
}, error) {
	return _RewardManager.Contract.GetRewardEpochTotals(&_RewardManager.CallOpts, _rewardEpochId)
}

// GetRewardEpochTotals is a free data retrieval call binding the contract method 0xdf339638.
//
// Solidity: function getRewardEpochTotals(uint24 _rewardEpochId) view returns(uint256 _totalRewardsWei, uint256 _totalInflationRewardsWei, uint256 _initialisedRewardsWei, uint256 _claimedRewardsWei, uint256 _burnedRewardsWei)
func (_RewardManager *RewardManagerCallerSession) GetRewardEpochTotals(_rewardEpochId *big.Int) (struct {
	TotalRewardsWei          *big.Int
	TotalInflationRewardsWei *big.Int
	InitialisedRewardsWei    *big.Int
	ClaimedRewardsWei        *big.Int
	BurnedRewardsWei         *big.Int
}, error) {
	return _RewardManager.Contract.GetRewardEpochTotals(&_RewardManager.CallOpts, _rewardEpochId)
}

// GetRewardOffersManagerList is a free data retrieval call binding the contract method 0x7a17a34d.
//
// Solidity: function getRewardOffersManagerList() view returns(address[])
func (_RewardManager *RewardManagerCaller) GetRewardOffersManagerList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getRewardOffersManagerList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardOffersManagerList is a free data retrieval call binding the contract method 0x7a17a34d.
//
// Solidity: function getRewardOffersManagerList() view returns(address[])
func (_RewardManager *RewardManagerSession) GetRewardOffersManagerList() ([]common.Address, error) {
	return _RewardManager.Contract.GetRewardOffersManagerList(&_RewardManager.CallOpts)
}

// GetRewardOffersManagerList is a free data retrieval call binding the contract method 0x7a17a34d.
//
// Solidity: function getRewardOffersManagerList() view returns(address[])
func (_RewardManager *RewardManagerCallerSession) GetRewardOffersManagerList() ([]common.Address, error) {
	return _RewardManager.Contract.GetRewardOffersManagerList(&_RewardManager.CallOpts)
}

// GetStateOfRewards is a free data retrieval call binding the contract method 0xf1367b7f.
//
// Solidity: function getStateOfRewards(address _rewardOwner) view returns((uint24,bytes20,uint120,uint8,bool)[][] _rewardStates)
func (_RewardManager *RewardManagerCaller) GetStateOfRewards(opts *bind.CallOpts, _rewardOwner common.Address) ([][]RewardsV2InterfaceRewardState, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getStateOfRewards", _rewardOwner)

	if err != nil {
		return *new([][]RewardsV2InterfaceRewardState), err
	}

	out0 := *abi.ConvertType(out[0], new([][]RewardsV2InterfaceRewardState)).(*[][]RewardsV2InterfaceRewardState)

	return out0, err

}

// GetStateOfRewards is a free data retrieval call binding the contract method 0xf1367b7f.
//
// Solidity: function getStateOfRewards(address _rewardOwner) view returns((uint24,bytes20,uint120,uint8,bool)[][] _rewardStates)
func (_RewardManager *RewardManagerSession) GetStateOfRewards(_rewardOwner common.Address) ([][]RewardsV2InterfaceRewardState, error) {
	return _RewardManager.Contract.GetStateOfRewards(&_RewardManager.CallOpts, _rewardOwner)
}

// GetStateOfRewards is a free data retrieval call binding the contract method 0xf1367b7f.
//
// Solidity: function getStateOfRewards(address _rewardOwner) view returns((uint24,bytes20,uint120,uint8,bool)[][] _rewardStates)
func (_RewardManager *RewardManagerCallerSession) GetStateOfRewards(_rewardOwner common.Address) ([][]RewardsV2InterfaceRewardState, error) {
	return _RewardManager.Contract.GetStateOfRewards(&_RewardManager.CallOpts, _rewardOwner)
}

// GetStateOfRewardsAt is a free data retrieval call binding the contract method 0x06c7e243.
//
// Solidity: function getStateOfRewardsAt(address _rewardOwner, uint24 _rewardEpochId) view returns((uint24,bytes20,uint120,uint8,bool)[] _rewardStates)
func (_RewardManager *RewardManagerCaller) GetStateOfRewardsAt(opts *bind.CallOpts, _rewardOwner common.Address, _rewardEpochId *big.Int) ([]RewardsV2InterfaceRewardState, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getStateOfRewardsAt", _rewardOwner, _rewardEpochId)

	if err != nil {
		return *new([]RewardsV2InterfaceRewardState), err
	}

	out0 := *abi.ConvertType(out[0], new([]RewardsV2InterfaceRewardState)).(*[]RewardsV2InterfaceRewardState)

	return out0, err

}

// GetStateOfRewardsAt is a free data retrieval call binding the contract method 0x06c7e243.
//
// Solidity: function getStateOfRewardsAt(address _rewardOwner, uint24 _rewardEpochId) view returns((uint24,bytes20,uint120,uint8,bool)[] _rewardStates)
func (_RewardManager *RewardManagerSession) GetStateOfRewardsAt(_rewardOwner common.Address, _rewardEpochId *big.Int) ([]RewardsV2InterfaceRewardState, error) {
	return _RewardManager.Contract.GetStateOfRewardsAt(&_RewardManager.CallOpts, _rewardOwner, _rewardEpochId)
}

// GetStateOfRewardsAt is a free data retrieval call binding the contract method 0x06c7e243.
//
// Solidity: function getStateOfRewardsAt(address _rewardOwner, uint24 _rewardEpochId) view returns((uint24,bytes20,uint120,uint8,bool)[] _rewardStates)
func (_RewardManager *RewardManagerCallerSession) GetStateOfRewardsAt(_rewardOwner common.Address, _rewardEpochId *big.Int) ([]RewardsV2InterfaceRewardState, error) {
	return _RewardManager.Contract.GetStateOfRewardsAt(&_RewardManager.CallOpts, _rewardOwner, _rewardEpochId)
}

// GetTokenPoolSupplyData is a free data retrieval call binding the contract method 0x2dafdbbf.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (_RewardManager *RewardManagerCaller) GetTokenPoolSupplyData(opts *bind.CallOpts) (struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getTokenPoolSupplyData")

	outstruct := new(struct {
		LockedFundsWei              *big.Int
		TotalInflationAuthorizedWei *big.Int
		TotalClaimedWei             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockedFundsWei = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalInflationAuthorizedWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalClaimedWei = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenPoolSupplyData is a free data retrieval call binding the contract method 0x2dafdbbf.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (_RewardManager *RewardManagerSession) GetTokenPoolSupplyData() (struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}, error) {
	return _RewardManager.Contract.GetTokenPoolSupplyData(&_RewardManager.CallOpts)
}

// GetTokenPoolSupplyData is a free data retrieval call binding the contract method 0x2dafdbbf.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (_RewardManager *RewardManagerCallerSession) GetTokenPoolSupplyData() (struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}, error) {
	return _RewardManager.Contract.GetTokenPoolSupplyData(&_RewardManager.CallOpts)
}

// GetTotals is a free data retrieval call binding the contract method 0x84e10a90.
//
// Solidity: function getTotals() view returns(uint256 _totalRewardsWei, uint256 _totalInflationRewardsWei, uint256 _totalClaimedWei, uint256 _totalBurnedWei)
func (_RewardManager *RewardManagerCaller) GetTotals(opts *bind.CallOpts) (struct {
	TotalRewardsWei          *big.Int
	TotalInflationRewardsWei *big.Int
	TotalClaimedWei          *big.Int
	TotalBurnedWei           *big.Int
}, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getTotals")

	outstruct := new(struct {
		TotalRewardsWei          *big.Int
		TotalInflationRewardsWei *big.Int
		TotalClaimedWei          *big.Int
		TotalBurnedWei           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalRewardsWei = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalInflationRewardsWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalClaimedWei = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalBurnedWei = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTotals is a free data retrieval call binding the contract method 0x84e10a90.
//
// Solidity: function getTotals() view returns(uint256 _totalRewardsWei, uint256 _totalInflationRewardsWei, uint256 _totalClaimedWei, uint256 _totalBurnedWei)
func (_RewardManager *RewardManagerSession) GetTotals() (struct {
	TotalRewardsWei          *big.Int
	TotalInflationRewardsWei *big.Int
	TotalClaimedWei          *big.Int
	TotalBurnedWei           *big.Int
}, error) {
	return _RewardManager.Contract.GetTotals(&_RewardManager.CallOpts)
}

// GetTotals is a free data retrieval call binding the contract method 0x84e10a90.
//
// Solidity: function getTotals() view returns(uint256 _totalRewardsWei, uint256 _totalInflationRewardsWei, uint256 _totalClaimedWei, uint256 _totalBurnedWei)
func (_RewardManager *RewardManagerCallerSession) GetTotals() (struct {
	TotalRewardsWei          *big.Int
	TotalInflationRewardsWei *big.Int
	TotalClaimedWei          *big.Int
	TotalBurnedWei           *big.Int
}, error) {
	return _RewardManager.Contract.GetTotals(&_RewardManager.CallOpts)
}

// GetUnclaimedRewardState is a free data retrieval call binding the contract method 0x9ee5de33.
//
// Solidity: function getUnclaimedRewardState(address _beneficiary, uint24 _rewardEpochId, uint8 _claimType) view returns((bool,uint120,uint128) _state)
func (_RewardManager *RewardManagerCaller) GetUnclaimedRewardState(opts *bind.CallOpts, _beneficiary common.Address, _rewardEpochId *big.Int, _claimType uint8) (IRewardManagerUnclaimedRewardState, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "getUnclaimedRewardState", _beneficiary, _rewardEpochId, _claimType)

	if err != nil {
		return *new(IRewardManagerUnclaimedRewardState), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardManagerUnclaimedRewardState)).(*IRewardManagerUnclaimedRewardState)

	return out0, err

}

// GetUnclaimedRewardState is a free data retrieval call binding the contract method 0x9ee5de33.
//
// Solidity: function getUnclaimedRewardState(address _beneficiary, uint24 _rewardEpochId, uint8 _claimType) view returns((bool,uint120,uint128) _state)
func (_RewardManager *RewardManagerSession) GetUnclaimedRewardState(_beneficiary common.Address, _rewardEpochId *big.Int, _claimType uint8) (IRewardManagerUnclaimedRewardState, error) {
	return _RewardManager.Contract.GetUnclaimedRewardState(&_RewardManager.CallOpts, _beneficiary, _rewardEpochId, _claimType)
}

// GetUnclaimedRewardState is a free data retrieval call binding the contract method 0x9ee5de33.
//
// Solidity: function getUnclaimedRewardState(address _beneficiary, uint24 _rewardEpochId, uint8 _claimType) view returns((bool,uint120,uint128) _state)
func (_RewardManager *RewardManagerCallerSession) GetUnclaimedRewardState(_beneficiary common.Address, _rewardEpochId *big.Int, _claimType uint8) (IRewardManagerUnclaimedRewardState, error) {
	return _RewardManager.Contract.GetUnclaimedRewardState(&_RewardManager.CallOpts, _beneficiary, _rewardEpochId, _claimType)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_RewardManager *RewardManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_RewardManager *RewardManagerSession) Governance() (common.Address, error) {
	return _RewardManager.Contract.Governance(&_RewardManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_RewardManager *RewardManagerCallerSession) Governance() (common.Address, error) {
	return _RewardManager.Contract.Governance(&_RewardManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_RewardManager *RewardManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_RewardManager *RewardManagerSession) GovernanceSettings() (common.Address, error) {
	return _RewardManager.Contract.GovernanceSettings(&_RewardManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_RewardManager *RewardManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _RewardManager.Contract.GovernanceSettings(&_RewardManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_RewardManager *RewardManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_RewardManager *RewardManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _RewardManager.Contract.IsExecutor(&_RewardManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_RewardManager *RewardManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _RewardManager.Contract.IsExecutor(&_RewardManager.CallOpts, _address)
}

// NewRewardManager is a free data retrieval call binding the contract method 0xfbde354d.
//
// Solidity: function newRewardManager() view returns(address)
func (_RewardManager *RewardManagerCaller) NewRewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "newRewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewRewardManager is a free data retrieval call binding the contract method 0xfbde354d.
//
// Solidity: function newRewardManager() view returns(address)
func (_RewardManager *RewardManagerSession) NewRewardManager() (common.Address, error) {
	return _RewardManager.Contract.NewRewardManager(&_RewardManager.CallOpts)
}

// NewRewardManager is a free data retrieval call binding the contract method 0xfbde354d.
//
// Solidity: function newRewardManager() view returns(address)
func (_RewardManager *RewardManagerCallerSession) NewRewardManager() (common.Address, error) {
	return _RewardManager.Contract.NewRewardManager(&_RewardManager.CallOpts)
}

// NoOfInitialisedWeightBasedClaims is a free data retrieval call binding the contract method 0x4b6e018d.
//
// Solidity: function noOfInitialisedWeightBasedClaims(uint256 rewardEpochId) view returns(uint256)
func (_RewardManager *RewardManagerCaller) NoOfInitialisedWeightBasedClaims(opts *bind.CallOpts, rewardEpochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "noOfInitialisedWeightBasedClaims", rewardEpochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoOfInitialisedWeightBasedClaims is a free data retrieval call binding the contract method 0x4b6e018d.
//
// Solidity: function noOfInitialisedWeightBasedClaims(uint256 rewardEpochId) view returns(uint256)
func (_RewardManager *RewardManagerSession) NoOfInitialisedWeightBasedClaims(rewardEpochId *big.Int) (*big.Int, error) {
	return _RewardManager.Contract.NoOfInitialisedWeightBasedClaims(&_RewardManager.CallOpts, rewardEpochId)
}

// NoOfInitialisedWeightBasedClaims is a free data retrieval call binding the contract method 0x4b6e018d.
//
// Solidity: function noOfInitialisedWeightBasedClaims(uint256 rewardEpochId) view returns(uint256)
func (_RewardManager *RewardManagerCallerSession) NoOfInitialisedWeightBasedClaims(rewardEpochId *big.Int) (*big.Int, error) {
	return _RewardManager.Contract.NoOfInitialisedWeightBasedClaims(&_RewardManager.CallOpts, rewardEpochId)
}

// OldRewardManager is a free data retrieval call binding the contract method 0x36300be2.
//
// Solidity: function oldRewardManager() view returns(address)
func (_RewardManager *RewardManagerCaller) OldRewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "oldRewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldRewardManager is a free data retrieval call binding the contract method 0x36300be2.
//
// Solidity: function oldRewardManager() view returns(address)
func (_RewardManager *RewardManagerSession) OldRewardManager() (common.Address, error) {
	return _RewardManager.Contract.OldRewardManager(&_RewardManager.CallOpts)
}

// OldRewardManager is a free data retrieval call binding the contract method 0x36300be2.
//
// Solidity: function oldRewardManager() view returns(address)
func (_RewardManager *RewardManagerCallerSession) OldRewardManager() (common.Address, error) {
	return _RewardManager.Contract.OldRewardManager(&_RewardManager.CallOpts)
}

// PChainStakeMirror is a free data retrieval call binding the contract method 0x62d9c89a.
//
// Solidity: function pChainStakeMirror() view returns(address)
func (_RewardManager *RewardManagerCaller) PChainStakeMirror(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "pChainStakeMirror")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PChainStakeMirror is a free data retrieval call binding the contract method 0x62d9c89a.
//
// Solidity: function pChainStakeMirror() view returns(address)
func (_RewardManager *RewardManagerSession) PChainStakeMirror() (common.Address, error) {
	return _RewardManager.Contract.PChainStakeMirror(&_RewardManager.CallOpts)
}

// PChainStakeMirror is a free data retrieval call binding the contract method 0x62d9c89a.
//
// Solidity: function pChainStakeMirror() view returns(address)
func (_RewardManager *RewardManagerCallerSession) PChainStakeMirror() (common.Address, error) {
	return _RewardManager.Contract.PChainStakeMirror(&_RewardManager.CallOpts)
}

// PChainStakeMirrorEnabled is a free data retrieval call binding the contract method 0x7bf756c9.
//
// Solidity: function pChainStakeMirrorEnabled() view returns(bool)
func (_RewardManager *RewardManagerCaller) PChainStakeMirrorEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "pChainStakeMirrorEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PChainStakeMirrorEnabled is a free data retrieval call binding the contract method 0x7bf756c9.
//
// Solidity: function pChainStakeMirrorEnabled() view returns(bool)
func (_RewardManager *RewardManagerSession) PChainStakeMirrorEnabled() (bool, error) {
	return _RewardManager.Contract.PChainStakeMirrorEnabled(&_RewardManager.CallOpts)
}

// PChainStakeMirrorEnabled is a free data retrieval call binding the contract method 0x7bf756c9.
//
// Solidity: function pChainStakeMirrorEnabled() view returns(bool)
func (_RewardManager *RewardManagerCallerSession) PChainStakeMirrorEnabled() (bool, error) {
	return _RewardManager.Contract.PChainStakeMirrorEnabled(&_RewardManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_RewardManager *RewardManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_RewardManager *RewardManagerSession) ProductionMode() (bool, error) {
	return _RewardManager.Contract.ProductionMode(&_RewardManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_RewardManager *RewardManagerCallerSession) ProductionMode() (bool, error) {
	return _RewardManager.Contract.ProductionMode(&_RewardManager.CallOpts)
}

// RewardManagerId is a free data retrieval call binding the contract method 0x2ae07e9a.
//
// Solidity: function rewardManagerId() view returns(uint256)
func (_RewardManager *RewardManagerCaller) RewardManagerId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "rewardManagerId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardManagerId is a free data retrieval call binding the contract method 0x2ae07e9a.
//
// Solidity: function rewardManagerId() view returns(uint256)
func (_RewardManager *RewardManagerSession) RewardManagerId() (*big.Int, error) {
	return _RewardManager.Contract.RewardManagerId(&_RewardManager.CallOpts)
}

// RewardManagerId is a free data retrieval call binding the contract method 0x2ae07e9a.
//
// Solidity: function rewardManagerId() view returns(uint256)
func (_RewardManager *RewardManagerCallerSession) RewardManagerId() (*big.Int, error) {
	return _RewardManager.Contract.RewardManagerId(&_RewardManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_RewardManager *RewardManagerCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "timelockedCalls", selector)

	outstruct := new(struct {
		AllowedAfterTimestamp *big.Int
		EncodedCall           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowedAfterTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_RewardManager *RewardManagerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _RewardManager.Contract.TimelockedCalls(&_RewardManager.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_RewardManager *RewardManagerCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _RewardManager.Contract.TimelockedCalls(&_RewardManager.CallOpts, selector)
}

// WNat is a free data retrieval call binding the contract method 0x9edbf007.
//
// Solidity: function wNat() view returns(address)
func (_RewardManager *RewardManagerCaller) WNat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "wNat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WNat is a free data retrieval call binding the contract method 0x9edbf007.
//
// Solidity: function wNat() view returns(address)
func (_RewardManager *RewardManagerSession) WNat() (common.Address, error) {
	return _RewardManager.Contract.WNat(&_RewardManager.CallOpts)
}

// WNat is a free data retrieval call binding the contract method 0x9edbf007.
//
// Solidity: function wNat() view returns(address)
func (_RewardManager *RewardManagerCallerSession) WNat() (common.Address, error) {
	return _RewardManager.Contract.WNat(&_RewardManager.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_RewardManager *RewardManagerTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_RewardManager *RewardManagerSession) Activate() (*types.Transaction, error) {
	return _RewardManager.Contract.Activate(&_RewardManager.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_RewardManager *RewardManagerTransactorSession) Activate() (*types.Transaction, error) {
	return _RewardManager.Contract.Activate(&_RewardManager.TransactOpts)
}

// AutoClaim is a paid mutator transaction binding the contract method 0x15f253fb.
//
// Solidity: function autoClaim(address[] _rewardOwners, uint24 _rewardEpochId, (bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns()
func (_RewardManager *RewardManagerTransactor) AutoClaim(opts *bind.TransactOpts, _rewardOwners []common.Address, _rewardEpochId *big.Int, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "autoClaim", _rewardOwners, _rewardEpochId, _proofs)
}

// AutoClaim is a paid mutator transaction binding the contract method 0x15f253fb.
//
// Solidity: function autoClaim(address[] _rewardOwners, uint24 _rewardEpochId, (bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns()
func (_RewardManager *RewardManagerSession) AutoClaim(_rewardOwners []common.Address, _rewardEpochId *big.Int, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.Contract.AutoClaim(&_RewardManager.TransactOpts, _rewardOwners, _rewardEpochId, _proofs)
}

// AutoClaim is a paid mutator transaction binding the contract method 0x15f253fb.
//
// Solidity: function autoClaim(address[] _rewardOwners, uint24 _rewardEpochId, (bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns()
func (_RewardManager *RewardManagerTransactorSession) AutoClaim(_rewardOwners []common.Address, _rewardEpochId *big.Int, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.Contract.AutoClaim(&_RewardManager.TransactOpts, _rewardOwners, _rewardEpochId, _proofs)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_RewardManager *RewardManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_RewardManager *RewardManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _RewardManager.Contract.CancelGovernanceCall(&_RewardManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_RewardManager *RewardManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _RewardManager.Contract.CancelGovernanceCall(&_RewardManager.TransactOpts, _selector)
}

// Claim is a paid mutator transaction binding the contract method 0x8e33aba5.
//
// Solidity: function claim(address _rewardOwner, address _recipient, uint24 _rewardEpochId, bool _wrap, (bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns(uint256 _rewardAmountWei)
func (_RewardManager *RewardManagerTransactor) Claim(opts *bind.TransactOpts, _rewardOwner common.Address, _recipient common.Address, _rewardEpochId *big.Int, _wrap bool, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "claim", _rewardOwner, _recipient, _rewardEpochId, _wrap, _proofs)
}

// Claim is a paid mutator transaction binding the contract method 0x8e33aba5.
//
// Solidity: function claim(address _rewardOwner, address _recipient, uint24 _rewardEpochId, bool _wrap, (bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns(uint256 _rewardAmountWei)
func (_RewardManager *RewardManagerSession) Claim(_rewardOwner common.Address, _recipient common.Address, _rewardEpochId *big.Int, _wrap bool, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.Contract.Claim(&_RewardManager.TransactOpts, _rewardOwner, _recipient, _rewardEpochId, _wrap, _proofs)
}

// Claim is a paid mutator transaction binding the contract method 0x8e33aba5.
//
// Solidity: function claim(address _rewardOwner, address _recipient, uint24 _rewardEpochId, bool _wrap, (bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns(uint256 _rewardAmountWei)
func (_RewardManager *RewardManagerTransactorSession) Claim(_rewardOwner common.Address, _recipient common.Address, _rewardEpochId *big.Int, _wrap bool, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.Contract.Claim(&_RewardManager.TransactOpts, _rewardOwner, _recipient, _rewardEpochId, _wrap, _proofs)
}

// ClaimProxy is a paid mutator transaction binding the contract method 0xaa145443.
//
// Solidity: function claimProxy(address _msgSender, address _rewardOwner, address _recipient, uint24 _rewardEpochId, bool _wrap, (bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns(uint256 _rewardAmountWei)
func (_RewardManager *RewardManagerTransactor) ClaimProxy(opts *bind.TransactOpts, _msgSender common.Address, _rewardOwner common.Address, _recipient common.Address, _rewardEpochId *big.Int, _wrap bool, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "claimProxy", _msgSender, _rewardOwner, _recipient, _rewardEpochId, _wrap, _proofs)
}

// ClaimProxy is a paid mutator transaction binding the contract method 0xaa145443.
//
// Solidity: function claimProxy(address _msgSender, address _rewardOwner, address _recipient, uint24 _rewardEpochId, bool _wrap, (bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns(uint256 _rewardAmountWei)
func (_RewardManager *RewardManagerSession) ClaimProxy(_msgSender common.Address, _rewardOwner common.Address, _recipient common.Address, _rewardEpochId *big.Int, _wrap bool, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.Contract.ClaimProxy(&_RewardManager.TransactOpts, _msgSender, _rewardOwner, _recipient, _rewardEpochId, _wrap, _proofs)
}

// ClaimProxy is a paid mutator transaction binding the contract method 0xaa145443.
//
// Solidity: function claimProxy(address _msgSender, address _rewardOwner, address _recipient, uint24 _rewardEpochId, bool _wrap, (bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns(uint256 _rewardAmountWei)
func (_RewardManager *RewardManagerTransactorSession) ClaimProxy(_msgSender common.Address, _rewardOwner common.Address, _recipient common.Address, _rewardEpochId *big.Int, _wrap bool, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.Contract.ClaimProxy(&_RewardManager.TransactOpts, _msgSender, _rewardOwner, _recipient, _rewardEpochId, _wrap, _proofs)
}

// CloseExpiredRewardEpoch is a paid mutator transaction binding the contract method 0xd6c1dbee.
//
// Solidity: function closeExpiredRewardEpoch(uint256 _rewardEpochId) returns()
func (_RewardManager *RewardManagerTransactor) CloseExpiredRewardEpoch(opts *bind.TransactOpts, _rewardEpochId *big.Int) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "closeExpiredRewardEpoch", _rewardEpochId)
}

// CloseExpiredRewardEpoch is a paid mutator transaction binding the contract method 0xd6c1dbee.
//
// Solidity: function closeExpiredRewardEpoch(uint256 _rewardEpochId) returns()
func (_RewardManager *RewardManagerSession) CloseExpiredRewardEpoch(_rewardEpochId *big.Int) (*types.Transaction, error) {
	return _RewardManager.Contract.CloseExpiredRewardEpoch(&_RewardManager.TransactOpts, _rewardEpochId)
}

// CloseExpiredRewardEpoch is a paid mutator transaction binding the contract method 0xd6c1dbee.
//
// Solidity: function closeExpiredRewardEpoch(uint256 _rewardEpochId) returns()
func (_RewardManager *RewardManagerTransactorSession) CloseExpiredRewardEpoch(_rewardEpochId *big.Int) (*types.Transaction, error) {
	return _RewardManager.Contract.CloseExpiredRewardEpoch(&_RewardManager.TransactOpts, _rewardEpochId)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_RewardManager *RewardManagerTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_RewardManager *RewardManagerSession) Deactivate() (*types.Transaction, error) {
	return _RewardManager.Contract.Deactivate(&_RewardManager.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_RewardManager *RewardManagerTransactorSession) Deactivate() (*types.Transaction, error) {
	return _RewardManager.Contract.Deactivate(&_RewardManager.TransactOpts)
}

// EnableCChainStake is a paid mutator transaction binding the contract method 0xfd95b2e0.
//
// Solidity: function enableCChainStake() returns()
func (_RewardManager *RewardManagerTransactor) EnableCChainStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "enableCChainStake")
}

// EnableCChainStake is a paid mutator transaction binding the contract method 0xfd95b2e0.
//
// Solidity: function enableCChainStake() returns()
func (_RewardManager *RewardManagerSession) EnableCChainStake() (*types.Transaction, error) {
	return _RewardManager.Contract.EnableCChainStake(&_RewardManager.TransactOpts)
}

// EnableCChainStake is a paid mutator transaction binding the contract method 0xfd95b2e0.
//
// Solidity: function enableCChainStake() returns()
func (_RewardManager *RewardManagerTransactorSession) EnableCChainStake() (*types.Transaction, error) {
	return _RewardManager.Contract.EnableCChainStake(&_RewardManager.TransactOpts)
}

// EnableClaims is a paid mutator transaction binding the contract method 0xea28edad.
//
// Solidity: function enableClaims() returns()
func (_RewardManager *RewardManagerTransactor) EnableClaims(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "enableClaims")
}

// EnableClaims is a paid mutator transaction binding the contract method 0xea28edad.
//
// Solidity: function enableClaims() returns()
func (_RewardManager *RewardManagerSession) EnableClaims() (*types.Transaction, error) {
	return _RewardManager.Contract.EnableClaims(&_RewardManager.TransactOpts)
}

// EnableClaims is a paid mutator transaction binding the contract method 0xea28edad.
//
// Solidity: function enableClaims() returns()
func (_RewardManager *RewardManagerTransactorSession) EnableClaims() (*types.Transaction, error) {
	return _RewardManager.Contract.EnableClaims(&_RewardManager.TransactOpts)
}

// EnablePChainStakeMirror is a paid mutator transaction binding the contract method 0xb006b4e3.
//
// Solidity: function enablePChainStakeMirror() returns()
func (_RewardManager *RewardManagerTransactor) EnablePChainStakeMirror(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "enablePChainStakeMirror")
}

// EnablePChainStakeMirror is a paid mutator transaction binding the contract method 0xb006b4e3.
//
// Solidity: function enablePChainStakeMirror() returns()
func (_RewardManager *RewardManagerSession) EnablePChainStakeMirror() (*types.Transaction, error) {
	return _RewardManager.Contract.EnablePChainStakeMirror(&_RewardManager.TransactOpts)
}

// EnablePChainStakeMirror is a paid mutator transaction binding the contract method 0xb006b4e3.
//
// Solidity: function enablePChainStakeMirror() returns()
func (_RewardManager *RewardManagerTransactorSession) EnablePChainStakeMirror() (*types.Transaction, error) {
	return _RewardManager.Contract.EnablePChainStakeMirror(&_RewardManager.TransactOpts)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_RewardManager *RewardManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_RewardManager *RewardManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _RewardManager.Contract.ExecuteGovernanceCall(&_RewardManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_RewardManager *RewardManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _RewardManager.Contract.ExecuteGovernanceCall(&_RewardManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_RewardManager *RewardManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_RewardManager *RewardManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.Initialise(&_RewardManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_RewardManager *RewardManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.Initialise(&_RewardManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// InitialiseWeightBasedClaims is a paid mutator transaction binding the contract method 0x3ce7522a.
//
// Solidity: function initialiseWeightBasedClaims((bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns()
func (_RewardManager *RewardManagerTransactor) InitialiseWeightBasedClaims(opts *bind.TransactOpts, _proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "initialiseWeightBasedClaims", _proofs)
}

// InitialiseWeightBasedClaims is a paid mutator transaction binding the contract method 0x3ce7522a.
//
// Solidity: function initialiseWeightBasedClaims((bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns()
func (_RewardManager *RewardManagerSession) InitialiseWeightBasedClaims(_proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.Contract.InitialiseWeightBasedClaims(&_RewardManager.TransactOpts, _proofs)
}

// InitialiseWeightBasedClaims is a paid mutator transaction binding the contract method 0x3ce7522a.
//
// Solidity: function initialiseWeightBasedClaims((bytes32[],(uint24,bytes20,uint120,uint8))[] _proofs) returns()
func (_RewardManager *RewardManagerTransactorSession) InitialiseWeightBasedClaims(_proofs []RewardsV2InterfaceRewardClaimWithProof) (*types.Transaction, error) {
	return _RewardManager.Contract.InitialiseWeightBasedClaims(&_RewardManager.TransactOpts, _proofs)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xa02e86e5.
//
// Solidity: function receiveRewards(uint24 _rewardEpochId, bool _inflation) payable returns()
func (_RewardManager *RewardManagerTransactor) ReceiveRewards(opts *bind.TransactOpts, _rewardEpochId *big.Int, _inflation bool) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "receiveRewards", _rewardEpochId, _inflation)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xa02e86e5.
//
// Solidity: function receiveRewards(uint24 _rewardEpochId, bool _inflation) payable returns()
func (_RewardManager *RewardManagerSession) ReceiveRewards(_rewardEpochId *big.Int, _inflation bool) (*types.Transaction, error) {
	return _RewardManager.Contract.ReceiveRewards(&_RewardManager.TransactOpts, _rewardEpochId, _inflation)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xa02e86e5.
//
// Solidity: function receiveRewards(uint24 _rewardEpochId, bool _inflation) payable returns()
func (_RewardManager *RewardManagerTransactorSession) ReceiveRewards(_rewardEpochId *big.Int, _inflation bool) (*types.Transaction, error) {
	return _RewardManager.Contract.ReceiveRewards(&_RewardManager.TransactOpts, _rewardEpochId, _inflation)
}

// SetInitialRewardData is a paid mutator transaction binding the contract method 0x1de56098.
//
// Solidity: function setInitialRewardData() returns()
func (_RewardManager *RewardManagerTransactor) SetInitialRewardData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "setInitialRewardData")
}

// SetInitialRewardData is a paid mutator transaction binding the contract method 0x1de56098.
//
// Solidity: function setInitialRewardData() returns()
func (_RewardManager *RewardManagerSession) SetInitialRewardData() (*types.Transaction, error) {
	return _RewardManager.Contract.SetInitialRewardData(&_RewardManager.TransactOpts)
}

// SetInitialRewardData is a paid mutator transaction binding the contract method 0x1de56098.
//
// Solidity: function setInitialRewardData() returns()
func (_RewardManager *RewardManagerTransactorSession) SetInitialRewardData() (*types.Transaction, error) {
	return _RewardManager.Contract.SetInitialRewardData(&_RewardManager.TransactOpts)
}

// SetNewRewardManager is a paid mutator transaction binding the contract method 0x470ce087.
//
// Solidity: function setNewRewardManager(address _newRewardManager) returns()
func (_RewardManager *RewardManagerTransactor) SetNewRewardManager(opts *bind.TransactOpts, _newRewardManager common.Address) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "setNewRewardManager", _newRewardManager)
}

// SetNewRewardManager is a paid mutator transaction binding the contract method 0x470ce087.
//
// Solidity: function setNewRewardManager(address _newRewardManager) returns()
func (_RewardManager *RewardManagerSession) SetNewRewardManager(_newRewardManager common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.SetNewRewardManager(&_RewardManager.TransactOpts, _newRewardManager)
}

// SetNewRewardManager is a paid mutator transaction binding the contract method 0x470ce087.
//
// Solidity: function setNewRewardManager(address _newRewardManager) returns()
func (_RewardManager *RewardManagerTransactorSession) SetNewRewardManager(_newRewardManager common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.SetNewRewardManager(&_RewardManager.TransactOpts, _newRewardManager)
}

// SetRewardOffersManagerList is a paid mutator transaction binding the contract method 0xb753801d.
//
// Solidity: function setRewardOffersManagerList(address[] _rewardOffersManagerList) returns()
func (_RewardManager *RewardManagerTransactor) SetRewardOffersManagerList(opts *bind.TransactOpts, _rewardOffersManagerList []common.Address) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "setRewardOffersManagerList", _rewardOffersManagerList)
}

// SetRewardOffersManagerList is a paid mutator transaction binding the contract method 0xb753801d.
//
// Solidity: function setRewardOffersManagerList(address[] _rewardOffersManagerList) returns()
func (_RewardManager *RewardManagerSession) SetRewardOffersManagerList(_rewardOffersManagerList []common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.SetRewardOffersManagerList(&_RewardManager.TransactOpts, _rewardOffersManagerList)
}

// SetRewardOffersManagerList is a paid mutator transaction binding the contract method 0xb753801d.
//
// Solidity: function setRewardOffersManagerList(address[] _rewardOffersManagerList) returns()
func (_RewardManager *RewardManagerTransactorSession) SetRewardOffersManagerList(_rewardOffersManagerList []common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.SetRewardOffersManagerList(&_RewardManager.TransactOpts, _rewardOffersManagerList)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_RewardManager *RewardManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_RewardManager *RewardManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _RewardManager.Contract.SwitchToProductionMode(&_RewardManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_RewardManager *RewardManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _RewardManager.Contract.SwitchToProductionMode(&_RewardManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_RewardManager *RewardManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_RewardManager *RewardManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.UpdateContractAddresses(&_RewardManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_RewardManager *RewardManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.UpdateContractAddresses(&_RewardManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// RewardManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the RewardManager contract.
type RewardManagerGovernanceCallTimelockedIterator struct {
	Event *RewardManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerGovernanceCallTimelocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardManagerGovernanceCallTimelocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the RewardManager contract.
type RewardManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_RewardManager *RewardManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*RewardManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &RewardManagerGovernanceCallTimelockedIterator{contract: _RewardManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_RewardManager *RewardManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *RewardManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerGovernanceCallTimelocked)
				if err := _RewardManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_RewardManager *RewardManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*RewardManagerGovernanceCallTimelocked, error) {
	event := new(RewardManagerGovernanceCallTimelocked)
	if err := _RewardManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the RewardManager contract.
type RewardManagerGovernanceInitialisedIterator struct {
	Event *RewardManagerGovernanceInitialised // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerGovernanceInitialised)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardManagerGovernanceInitialised)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerGovernanceInitialised represents a GovernanceInitialised event raised by the RewardManager contract.
type RewardManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_RewardManager *RewardManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*RewardManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &RewardManagerGovernanceInitialisedIterator{contract: _RewardManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_RewardManager *RewardManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *RewardManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerGovernanceInitialised)
				if err := _RewardManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_RewardManager *RewardManagerFilterer) ParseGovernanceInitialised(log types.Log) (*RewardManagerGovernanceInitialised, error) {
	event := new(RewardManagerGovernanceInitialised)
	if err := _RewardManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the RewardManager contract.
type RewardManagerGovernedProductionModeEnteredIterator struct {
	Event *RewardManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerGovernedProductionModeEntered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardManagerGovernedProductionModeEntered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the RewardManager contract.
type RewardManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_RewardManager *RewardManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*RewardManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &RewardManagerGovernedProductionModeEnteredIterator{contract: _RewardManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_RewardManager *RewardManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *RewardManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerGovernedProductionModeEntered)
				if err := _RewardManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernedProductionModeEntered is a log parse operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_RewardManager *RewardManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*RewardManagerGovernedProductionModeEntered, error) {
	event := new(RewardManagerGovernedProductionModeEntered)
	if err := _RewardManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the RewardManager contract.
type RewardManagerRewardClaimedIterator struct {
	Event *RewardManagerRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardManagerRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerRewardClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardManagerRewardClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardManagerRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerRewardClaimed represents a RewardClaimed event raised by the RewardManager contract.
type RewardManagerRewardClaimed struct {
	Beneficiary   common.Address
	RewardOwner   common.Address
	Recipient     common.Address
	RewardEpochId *big.Int
	ClaimType     uint8
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x06f77960d1401cc7d724b5c2b5ad672b9dbf08d8b11516a38c21697c23fbb0d2.
//
// Solidity: event RewardClaimed(address indexed beneficiary, address indexed rewardOwner, address indexed recipient, uint24 rewardEpochId, uint8 claimType, uint120 amount)
func (_RewardManager *RewardManagerFilterer) FilterRewardClaimed(opts *bind.FilterOpts, beneficiary []common.Address, rewardOwner []common.Address, recipient []common.Address) (*RewardManagerRewardClaimedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var rewardOwnerRule []interface{}
	for _, rewardOwnerItem := range rewardOwner {
		rewardOwnerRule = append(rewardOwnerRule, rewardOwnerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "RewardClaimed", beneficiaryRule, rewardOwnerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RewardManagerRewardClaimedIterator{contract: _RewardManager.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x06f77960d1401cc7d724b5c2b5ad672b9dbf08d8b11516a38c21697c23fbb0d2.
//
// Solidity: event RewardClaimed(address indexed beneficiary, address indexed rewardOwner, address indexed recipient, uint24 rewardEpochId, uint8 claimType, uint120 amount)
func (_RewardManager *RewardManagerFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *RewardManagerRewardClaimed, beneficiary []common.Address, rewardOwner []common.Address, recipient []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var rewardOwnerRule []interface{}
	for _, rewardOwnerItem := range rewardOwner {
		rewardOwnerRule = append(rewardOwnerRule, rewardOwnerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "RewardClaimed", beneficiaryRule, rewardOwnerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerRewardClaimed)
				if err := _RewardManager.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimed is a log parse operation binding the contract event 0x06f77960d1401cc7d724b5c2b5ad672b9dbf08d8b11516a38c21697c23fbb0d2.
//
// Solidity: event RewardClaimed(address indexed beneficiary, address indexed rewardOwner, address indexed recipient, uint24 rewardEpochId, uint8 claimType, uint120 amount)
func (_RewardManager *RewardManagerFilterer) ParseRewardClaimed(log types.Log) (*RewardManagerRewardClaimed, error) {
	event := new(RewardManagerRewardClaimed)
	if err := _RewardManager.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerRewardClaimsEnabledIterator is returned from FilterRewardClaimsEnabled and is used to iterate over the raw logs and unpacked data for RewardClaimsEnabled events raised by the RewardManager contract.
type RewardManagerRewardClaimsEnabledIterator struct {
	Event *RewardManagerRewardClaimsEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardManagerRewardClaimsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerRewardClaimsEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardManagerRewardClaimsEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardManagerRewardClaimsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerRewardClaimsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerRewardClaimsEnabled represents a RewardClaimsEnabled event raised by the RewardManager contract.
type RewardManagerRewardClaimsEnabled struct {
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimsEnabled is a free log retrieval operation binding the contract event 0x1cfb844c44f9325fc9ad6cc6191a4a24b0415137fe300b6c9071523a253f7a08.
//
// Solidity: event RewardClaimsEnabled(uint256 indexed rewardEpochId)
func (_RewardManager *RewardManagerFilterer) FilterRewardClaimsEnabled(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*RewardManagerRewardClaimsEnabledIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "RewardClaimsEnabled", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &RewardManagerRewardClaimsEnabledIterator{contract: _RewardManager.contract, event: "RewardClaimsEnabled", logs: logs, sub: sub}, nil
}

// WatchRewardClaimsEnabled is a free log subscription operation binding the contract event 0x1cfb844c44f9325fc9ad6cc6191a4a24b0415137fe300b6c9071523a253f7a08.
//
// Solidity: event RewardClaimsEnabled(uint256 indexed rewardEpochId)
func (_RewardManager *RewardManagerFilterer) WatchRewardClaimsEnabled(opts *bind.WatchOpts, sink chan<- *RewardManagerRewardClaimsEnabled, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "RewardClaimsEnabled", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerRewardClaimsEnabled)
				if err := _RewardManager.contract.UnpackLog(event, "RewardClaimsEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimsEnabled is a log parse operation binding the contract event 0x1cfb844c44f9325fc9ad6cc6191a4a24b0415137fe300b6c9071523a253f7a08.
//
// Solidity: event RewardClaimsEnabled(uint256 indexed rewardEpochId)
func (_RewardManager *RewardManagerFilterer) ParseRewardClaimsEnabled(log types.Log) (*RewardManagerRewardClaimsEnabled, error) {
	event := new(RewardManagerRewardClaimsEnabled)
	if err := _RewardManager.contract.UnpackLog(event, "RewardClaimsEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerRewardClaimsExpiredIterator is returned from FilterRewardClaimsExpired and is used to iterate over the raw logs and unpacked data for RewardClaimsExpired events raised by the RewardManager contract.
type RewardManagerRewardClaimsExpiredIterator struct {
	Event *RewardManagerRewardClaimsExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardManagerRewardClaimsExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerRewardClaimsExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardManagerRewardClaimsExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardManagerRewardClaimsExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerRewardClaimsExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerRewardClaimsExpired represents a RewardClaimsExpired event raised by the RewardManager contract.
type RewardManagerRewardClaimsExpired struct {
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimsExpired is a free log retrieval operation binding the contract event 0x5d05c64f281304391697cf987812e1a736413a062a9bdf39af4102209eb6fa58.
//
// Solidity: event RewardClaimsExpired(uint256 indexed rewardEpochId)
func (_RewardManager *RewardManagerFilterer) FilterRewardClaimsExpired(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*RewardManagerRewardClaimsExpiredIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "RewardClaimsExpired", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &RewardManagerRewardClaimsExpiredIterator{contract: _RewardManager.contract, event: "RewardClaimsExpired", logs: logs, sub: sub}, nil
}

// WatchRewardClaimsExpired is a free log subscription operation binding the contract event 0x5d05c64f281304391697cf987812e1a736413a062a9bdf39af4102209eb6fa58.
//
// Solidity: event RewardClaimsExpired(uint256 indexed rewardEpochId)
func (_RewardManager *RewardManagerFilterer) WatchRewardClaimsExpired(opts *bind.WatchOpts, sink chan<- *RewardManagerRewardClaimsExpired, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "RewardClaimsExpired", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerRewardClaimsExpired)
				if err := _RewardManager.contract.UnpackLog(event, "RewardClaimsExpired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimsExpired is a log parse operation binding the contract event 0x5d05c64f281304391697cf987812e1a736413a062a9bdf39af4102209eb6fa58.
//
// Solidity: event RewardClaimsExpired(uint256 indexed rewardEpochId)
func (_RewardManager *RewardManagerFilterer) ParseRewardClaimsExpired(log types.Log) (*RewardManagerRewardClaimsExpired, error) {
	event := new(RewardManagerRewardClaimsExpired)
	if err := _RewardManager.contract.UnpackLog(event, "RewardClaimsExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the RewardManager contract.
type RewardManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *RewardManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerTimelockedGovernanceCallCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardManagerTimelockedGovernanceCallCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the RewardManager contract.
type RewardManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_RewardManager *RewardManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*RewardManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &RewardManagerTimelockedGovernanceCallCanceledIterator{contract: _RewardManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_RewardManager *RewardManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *RewardManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerTimelockedGovernanceCallCanceled)
				if err := _RewardManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_RewardManager *RewardManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*RewardManagerTimelockedGovernanceCallCanceled, error) {
	event := new(RewardManagerTimelockedGovernanceCallCanceled)
	if err := _RewardManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the RewardManager contract.
type RewardManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *RewardManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerTimelockedGovernanceCallExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardManagerTimelockedGovernanceCallExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the RewardManager contract.
type RewardManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_RewardManager *RewardManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*RewardManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &RewardManagerTimelockedGovernanceCallExecutedIterator{contract: _RewardManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_RewardManager *RewardManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *RewardManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerTimelockedGovernanceCallExecuted)
				if err := _RewardManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_RewardManager *RewardManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*RewardManagerTimelockedGovernanceCallExecuted, error) {
	event := new(RewardManagerTimelockedGovernanceCallExecuted)
	if err := _RewardManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
