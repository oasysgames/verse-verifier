// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakemanager

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

// StakemanagerMetaData contains all meta data concerning the Stakemanager contract.
var StakemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInUse\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyJoined\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountMismatched\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBLS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBLSLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Locked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ObsoletedMethod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBlockProducer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotLastBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PastEpoch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAsOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AddedRewardBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldBLSPublicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newBLSPublicKey\",\"type\":\"bytes\"}],\"name\":\"BLSPublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedCommissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedUnstake\",\"type\":\"uint256\"}],\"name\":\"ClaimedLockedUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"OperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedUnstake\",\"type\":\"uint256\"}],\"name\":\"UnstakedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"ValidatorActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"ValidatorDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"until\",\"type\":\"uint256\"}],\"name\":\"ValidatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorJoined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorSlashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"activateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addRewardBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowlist\",\"outputs\":[{\"internalType\":\"contractIAllowlist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"blsPublicKeyToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidateManager\",\"outputs\":[{\"internalType\":\"contractICandidateValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"claimCommissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedUnstake\",\"type\":\"uint256\"}],\"name\":\"claimLockedUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimUnstakes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"environment\",\"outputs\":[{\"internalType\":\"contractIEnvironment\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getBlockAndSlashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getCommissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commissions\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockedUnstake\",\"type\":\"uint256\"}],\"name\":\"getLockedUnstake\",\"outputs\":[{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getLockedUnstakeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getLockedUnstakes\",\"outputs\":[{\"internalType\":\"enumToken.Type[]\",\"name\":\"tokens\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"unlockTimes\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"claimable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getOperatorStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getStakerStakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"oasStakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"woasStakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"soasStakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getTotalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amounts\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getUnstakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"oasUnstakes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"woasUnstakes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"soasUnstakes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"candidate\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakes\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getValidatorOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getValidatorStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getValidatorStakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"blsPublicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bool[]\",\"name\":\"candidates\",\"type\":\"bool[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEnvironment\",\"name\":\"_environment\",\"type\":\"address\"},{\"internalType\":\"contractIAllowlist\",\"name\":\"_allowlist\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"joinValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"restakeCommissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"restakeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeUpdates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumToken.Type\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstakeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"}],\"name\":\"updateBLSPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimCommission\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615ff280620000216000396000f3fe60806040526004361061027c5760003560e01c806374e2b63c1161014f578063cf5c13db116100c1578063e4b2477b1161007a578063e4b2477b14610880578063f3621e43146108b6578063f65a5ed2146108d6578063f8d6b1ab146108f6578063fa52c7d814610916578063ff3d3f601461094657600080fd5b8063cf5c13db14610784578063d0051adf146107a4578063d1f18ee1146107d5578063dbd61d8714610807578063df93c84214610827578063e1aca3411461086057600080fd5b80639168ae72116101135780639168ae72146106ae5780639c508219146106e4578063a6a41f4414610704578063ac7475ed14610724578063ad71bd3614610744578063cbc0fac61461076457600080fd5b806374e2b63c146105fd5780637b520aa8146106225780637befa74f14610658578063883252341461066b5780639043150b146106a657600080fd5b80632b47da52116101f3578063485cc955116101ac578063485cc9551461051b5780635c4fc4c51461053b5780635d94ccf61461056b5780635efc766e1461058b5780636b2b3369146105ab57806372431991146105cb57600080fd5b80632b47da52146104345780632ee462b31461046c57806333f32d781461048c578063428e8562146104ac57806345367f23146104cc57806346dfce7b146104ec57600080fd5b8063190b925711610245578063190b925714610332578063195afea1146103605780631c1b4f3a146103805780632168e8b4146103a057806322226367146103ce5780632b42ed8c1461040357600080fd5b8062c8ae891461028157806302fb4d85146102a35780630ddda63c146102c3578063158ef93e146102e35780631903cf1614610312575b600080fd5b34801561028d57600080fd5b506102a161029c3660046150d9565b610966565b005b3480156102af57600080fd5b506102a16102be36600461515f565b610b1d565b3480156102cf57600080fd5b506102a16102de36600461518b565b610d23565b3480156102ef57600080fd5b506000546102fd9060ff1681565b60405190151581526020015b60405180910390f35b34801561031e57600080fd5b506102a161032d366004615236565b610f2e565b34801561033e57600080fd5b5061035261034d36600461518b565b611127565b604051908152602001610309565b34801561036c57600080fd5b5061035261037b36600461515f565b611148565b34801561038c57600080fd5b5061035261039b36600461518b565b61117f565b3480156103ac57600080fd5b506103c06103bb3660046152e0565b61118f565b604051610309929190615346565b3480156103da57600080fd5b506103ee6103e936600461515f565b611275565b60408051928352602083019190915201610309565b34801561040f57600080fd5b5061042361041e366004615368565b611343565b6040516103099594939291906153d3565b34801561044057600080fd5b50600154610454906001600160a01b031681565b6040516001600160a01b039091168152602001610309565b34801561047857600080fd5b5061035261048736600461515f565b611690565b34801561049857600080fd5b506103526104a7366004615433565b611743565b3480156104b857600080fd5b506102a16104c7366004615236565b611909565b3480156104d857600080fd5b506103526104e736600461518b565b611b02565b3480156104f857600080fd5b5061050c610507366004615368565b611b98565b604051610309939291906154d2565b34801561052757600080fd5b506102a1610536366004615508565b611df5565b34801561054757600080fd5b5061055b61055636600461515f565b611e74565b6040516103099493929190615579565b34801561057757600080fd5b506102a161058636600461518b565b611f40565b34801561059757600080fd5b506104546105a636600461518b565b61204a565b3480156105b757600080fd5b506102a16105c63660046155a4565b612074565b3480156105d757600080fd5b506105eb6105e63660046155c1565b61216d565b60405161030996959493929190615677565b34801561060957600080fd5b506000546104549061010090046001600160a01b031681565b34801561062e57600080fd5b5061045461063d3660046155a4565b6006602052600090815260409020546001600160a01b031681565b6102a161066636600461572e565b612229565b34801561067757600080fd5b5061068b6106863660046155a4565b6123f4565b60408051938452602084019290925290820152606001610309565b6102a1612474565b3480156106ba57600080fd5b506104546106c93660046155a4565b6007602052600090815260409020546001600160a01b031681565b3480156106f057600080fd5b506103526106ff36600461515f565b6124a9565b34801561071057600080fd5b50600954610454906001600160a01b031681565b34801561073057600080fd5b506102a161073f3660046155a4565b612585565b34801561075057600080fd5b506103c061075f3660046152e0565b61268b565b34801561077057600080fd5b506102a161077f36600461515f565b612769565b34801561079057600080fd5b506102a161079f36600461515f565b612811565b3480156107b057600080fd5b506107c46107bf366004615773565b612a10565b6040516103099594939291906157a8565b3480156107e157600080fd5b506107f56107f036600461515f565b612cc8565b60405161030996959493929190615825565b34801561081357600080fd5b50610352610822366004615870565b612ed3565b34801561083357600080fd5b506103526108423660046155a4565b6001600160a01b031660009081526007602052604090206006015490565b34801561086c57600080fd5b506102a161087b36600461572e565b612f19565b34801561088c57600080fd5b5061045461089b36600461518b565b600a602052600090815260409020546001600160a01b031681565b3480156108c257600080fd5b506102a16108d1366004615870565b612f32565b3480156108e257600080fd5b506104546108f136600461518b565b61303e565b34801561090257600080fd5b506102a16109113660046155a4565b61304e565b34801561092257600080fd5b506109366109313660046155a4565b6130b0565b60405161030994939291906158a0565b34801561095257600080fd5b506102a161096136600461572e565b61316f565b336000818152600460205260409020546001600160a01b031661099c576040516372898ae960e11b815260040160405180910390fd5b3360006109ac60208286886158dd565b6109b591615907565b6000818152600a60205260409020549091506001600160a01b0316156109ee5760405163055ee1f160e31b815260040160405180910390fd5b6001600160a01b0382166000908152600460205260408120600b81018054919291610a1890615925565b80601f0160208091040260200160405190810160405280929190818152602001828054610a4490615925565b8015610a915780601f10610a6657610100808354040283529160200191610a91565b820191906000526020600020905b815481529060010190602001808311610a7457829003601f168201915b50505050509050610aad8787846134af9092919063ffffffff16565b6000838152600a60205260409081902080546001600160a01b0319166001600160a01b03871690811790915590517f7ea7e12060119574f657de08c5ef0970a24d7734612fb00c418ad40c7d4a84fd90610b0c9084908b908b90615960565b60405180910390a250505050505050565b6001600160a01b038083166000908152600660209081526040808320548416808452600490925290912054909116610b68576040516372898ae960e11b815260040160405180910390fd5b334114610b8857604051631cf4735960e01b815260040160405180910390fd5b6001600160a01b038084166000908152600660209081526040808320548416835260049182905280832083548251633fa4f24560e01b815292519195610c9b946101009092041692633fa4f2459281830192610120928290030181865afa158015610bf7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1b91906159a6565b600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c929190615a21565b84919087613537565b82546040519192506001600160a01b0316907f1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda90600090a28015610d1c5781546040518281526001600160a01b03909116907feb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802906020015b60405180910390a25b5050505050565b336000818152600460205260409020546001600160a01b0316610d59576040516372898ae960e11b815260040160405180910390fd5b336000818152600760205260409020546001600160a01b0316610d8f5760405163cf83d93d60e01b815260040160405180910390fd5b600060019054906101000a90046001600160a01b03166001600160a01b031663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa158015610de2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e069190615a4f565b15610e2457604051631e59ccd960e01b815260040160405180910390fd5b60008054338252600460205260408220610e4c9161010090046001600160a01b031686613616565b905080610e6c57604051637bc90c0560e11b815260040160405180910390fd5b610ef03333600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ec4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee89190615a21565b600085613636565b604051818152339081907fddd8e3ffe5c76cd1a7cca4b98662cb0a4e3da53ee24a873da632d28ba1043836906020015b60405180910390a350505050565b6001600160a01b03828116600090815260046020526040902080548492163314801590610f68575060018101546001600160a01b03163314155b15610f8657604051630101292160e31b815260040160405180910390fd5b6001600160a01b03808516600090815260046020526040902054859116610fc0576040516372898ae960e11b815260040160405180910390fd5b600060019054906101000a90046001600160a01b03166001600160a01b031663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa158015611013573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110379190615a4f565b1561105557604051631e59ccd960e01b815260040160405180910390fd5b6110ee600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110cf9190615a21565b6001600160a01b0387166000908152600460205260409020908661374b565b846001600160a01b03167fc11dfc9c24621433bb10587dc4bbae26a33a4aff53914e0d4c9fddf224a8cb7d85604051610d139190615a6a565b6002818154811061113757600080fd5b600091825260209091200154905081565b600080546001600160a01b038481168352600460205260408320611175929091610100909104168461375d565b5090505b92915050565b6003818154811061113757600080fd5b606060006111a384846005805490506138f1565b9093509050826001600160401b038111156111c0576111c06151a4565b6040519080825280602002602001820160405280156111e9578160200160208202803683370190505b50915060005b8381101561126d5760056112038287615a93565b8154811061121357611213615aab565b9060005260206000200160009054906101000a90046001600160a01b031683828151811061124357611243615aab565b6001600160a01b03909216602092830291909101909101528061126581615ac1565b9150506111ef565b509250929050565b600080611338600084116112ff57600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112fa9190615a21565b611301565b835b6001600160a01b038616600090815260046020908152604080832093835260098401825280832054600a9094019091529020549091565b909590945092505050565b6001600160a01b0384166000908152600760205260408120606091829182918291886113e557600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e09190615a21565b6113e7565b885b98506113f988886005805490506138f1565b9097509150866001600160401b03811115611416576114166151a4565b60405190808252806020026020018201604052801561143f578160200160208202803683370190505b509550866001600160401b0381111561145a5761145a6151a4565b604051908082528060200260200182016040528015611483578160200160208202803683370190505b509450866001600160401b0381111561149e5761149e6151a4565b6040519080825280602002602001820160405280156114c7578160200160208202803683370190505b509350866001600160401b038111156114e2576114e26151a4565b60405190808252806020026020018201604052801561150b578160200160208202803683370190505b50925060005b87811015611683576005611525828b615a93565b8154811061153557611535615aab565b9060005260206000200160009054906101000a90046001600160a01b031687828151811061156557611565615aab565b60200260200101906001600160a01b031690816001600160a01b0316815250506115b687828151811061159a5761159a615aab565b602002602001015160008c85613929909392919063ffffffff16565b8682815181106115c8576115c8615aab565b6020026020010181815250506116058782815181106115e9576115e9615aab565b602002602001015160018c85613929909392919063ffffffff16565b85828151811061161757611617615aab565b60200260200101818152505061165487828151811061163857611638615aab565b602002602001015160028c85613929909392919063ffffffff16565b84828151811061166657611666615aab565b60209081029190910101528061167b81615ac1565b915050611511565b5050945094509450945094565b600080821161171557600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117109190615a21565b611717565b815b6001600160a01b038416600090815260046020526040902090925061173c90836139e8565b9392505050565b600080600183600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561179c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117c09190615a21565b6117ca9190615adc565b6117d49190615adc565b845190915060005b84811015611900576117ef600184615a93565b6000805460405163fcbb371b60e01b81526004810184905292955090916101009091046001600160a01b03169063fcbb371b9060240161012060405180830381865afa158015611843573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061186791906159a6565b905060005b838110156118eb5760006118c88387600460008d878151811061189157611891615aab565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206139fb9092919063ffffffff16565b5090506118d58188615a93565b96505080806118e390615ac1565b91505061186c565b505080806118f890615ac1565b9150506117dc565b50505092915050565b6001600160a01b03828116600090815260046020526040902080548492163314801590611943575060018101546001600160a01b03163314155b1561196157604051630101292160e31b815260040160405180910390fd5b6001600160a01b0380851660009081526004602052604090205485911661199b576040516372898ae960e11b815260040160405180910390fd5b600060019054906101000a90046001600160a01b03166001600160a01b031663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a129190615a4f565b15611a3057604051631e59ccd960e01b815260040160405180910390fd5b611ac9600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aaa9190615a21565b6001600160a01b03871660009081526004602052604090209086613b6a565b846001600160a01b03167f0ad9bf1b8c026a174a2f30954417002a6ea00c9b08c1b8846c7951c687be809585604051610d139190615a6a565b6000808211611b8757600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b829190615a21565b611b89565b815b91506111796002600384613b77565b6001600160a01b0384166000908152600460205260408120606091829186611c3657600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c319190615a21565b611c38565b865b9650611c4c868683600701805490506138f1565b9095509150846001600160401b03811115611c6957611c696151a4565b604051908082528060200260200182016040528015611c92578160200160208202803683370190505b509350846001600160401b03811115611cad57611cad6151a4565b604051908082528060200260200182016040528015611cd6578160200160208202803683370190505b50925060005b85811015611de9576000600781848201611cf6858c615a93565b81548110611d0657611d06615aab565b60009182526020808320909101546001600160a01b03908116845290830193909352604090910190208054885191935090911690879084908110611d4c57611d4c615aab565b6001600160a01b0392831660209182029290920101528354611d739183911660028c613929565b8354611d8c9083906001600160a01b031660018d613929565b8454611da59084906001600160a01b031660008e613929565b611daf9190615a93565b611db99190615a93565b858381518110611dcb57611dcb615aab565b60209081029190910101525080611de181615ac1565b915050611cdc565b50509450945094915050565b334114611e1557604051631cf4735960e01b815260040160405180910390fd5b60005460ff1615611e385760405162dc149f60e41b815260040160405180910390fd5b6000805460016001600160a81b03199091166101006001600160a01b039586160217811790915580546001600160a01b03191691909216179055565b6001600160a01b03821660009081526007602052604081206006018054829182918291829187908110611ea957611ea9615aab565b6000918252602090912060408051606081019091526003909202018054829060ff166002811115611edc57611edc615541565b6002811115611eed57611eed615541565b81526020016001820154815260200160028201548152505090508060000151816020015182604001518360400151600014158015611f2f575083604001514210155b929a91995097509095509350505050565b336000818152600760205260409020546001600160a01b0316611f765760405163cf83d93d60e01b815260040160405180910390fd5b336000908152600760205260408120600601805484908110611f9a57611f9a615aab565b9060005260206000209060030201905060008160020154905080421015611fd4576040516303cb96db60e21b815260040160405180910390fd5b80611ff257604051630c8d9eab60e31b815260040160405180910390fd5b6000600283015560405184815233907fbf5f92dc2b945251eadf78c2ca629ae64053d979bfbc43a7b17a463707906bf99060200160405180910390a2815460018301546120449160ff16903390613c90565b50505050565b6005818154811061205a57600080fd5b6000918252602090912001546001600160a01b0316905081565b6001600160a01b03818116600090815260066020526040902054339116156120af5760405163055ee1f160e31b815260040160405180910390fd5b6001600160a01b03811660009081526004602052604090206120d19083613daf565b60058054600181019091557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b038381166001600160a01b03199283168117909355841660009081526006602090815260409182902080549093168417909255519182527fd5828184f48f65962d10eac907318df85953d4e3542a0f09b5932ee3fe398bdd910160405180910390a15050565b600954604051632d73a02f60e01b815260048101859052602481018490526044810183905260609182918291829182916000916001600160a01b0390911690632d73a02f90606401600060405180830381865afa1580156121d2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526121fa9190810190615cf6565b9091929394509091929350809650819750829850839950849a50859b5050505050505093975093979195509350565b6001600160a01b03808416600090815260046020526040902054849116612263576040516372898ae960e11b815260040160405180910390fd5b600060019054906101000a90046001600160a01b03166001600160a01b031663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156122b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122da9190615a4f565b156122f857604051631e59ccd960e01b815260040160405180910390fd5b8161231657604051637bc90c0560e11b815260040160405180910390fd5b612321833384613df1565b6123af3385600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612379573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061239d9190615a21565b6123a8906001615a93565b8686613636565b836001600160a01b0316336001600160a01b03167f8fc656e319452025372383dc27d933046d412b8253de50a10739eeaa59862be68585604051610f20929190615e1a565b6001600160a01b0380821660009081526007602052604081208154919283928392916124299183916101009091041684613e92565b60005490945061244a90829061010090046001600160a01b03166001613e92565b60005490935061246b90829061010090046001600160a01b03166002613e92565b93959294505050565b6040513481527f1de49774d094a85fc1bbbd16e8d09a865fb848218f41e2da4369f528c42ee42e9060200160405180910390a1565b6001600160a01b0380831660008181526006602090815260408083205485168352600490915281206001810154919390929116146124eb576000915050611179565b6000831161256f57600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612546573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061256a9190615a21565b612571565b825b925061257d81846139e8565b949350505050565b336000818152600460205260409020546001600160a01b03166125bb576040516372898ae960e11b815260040160405180910390fd5b6001600160a01b03828116600090815260066020526040902054339116156125f65760405163055ee1f160e31b815260040160405180910390fd5b6001600160a01b038082166000908152600460205260409020600181015490911661262182866140b8565b6001600160a01b0385811660008181526006602090815260409182902080546001600160a01b031916888616908117909155825194861685529084019290925290917f758820d0b14a01c1fa60b8d2bbef25ed1b6a5af4802e5dec3f08679255ba8bf39101610d13565b6060600061269f84846008805490506138f1565b9093509050826001600160401b038111156126bc576126bc6151a4565b6040519080825280602002602001820160405280156126e5578160200160208202803683370190505b50915060005b8381101561126d5760086126ff8287615a93565b8154811061270f5761270f615aab565b9060005260206000200160009054906101000a90046001600160a01b031683828151811061273f5761273f615aab565b6001600160a01b03909216602092830291909101909101528061276181615ac1565b9150506126eb565b336000818152600460205260409020546001600160a01b031661279f576040516372898ae960e11b815260040160405180910390fd5b600080543382526004602052604082206127c79161010090046001600160a01b031685613616565b60405181815290915033907f882d5671e5b36af50883197c33d48ba56ce337589958e871ba82fb0a54adf3e89060200160405180910390a280156120445761204460003383613c90565b6001600160a01b0380831660009081526004602052604090205483911661284b576040516372898ae960e11b815260040160405180910390fd5b336000818152600760205260409020546001600160a01b03166128815760405163cf83d93d60e01b815260040160405180910390fd5b600060019054906101000a90046001600160a01b03166001600160a01b031663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128f89190615a4f565b1561291657604051631e59ccd960e01b815260040160405180910390fd5b600080546001600160a01b03868116835260046020908152604080852033865260079092528420612951939092610100909104169087614132565b90508061297157604051637bc90c0560e11b815260040160405180910390fd5b6129c93386600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ec4573d6000803e3d6000fd5b6040518181526001600160a01b0386169033907fddd8e3ffe5c76cd1a7cca4b98662cb0a4e3da53ee24a873da632d28ba10438369060200160405180910390a35050505050565b6001600160a01b038316600090815260076020526040812060068101546060928392839283929190612a4590899089906138f1565b9097509150866001600160401b03811115612a6257612a626151a4565b604051908082528060200260200182016040528015612a8b578160200160208202803683370190505b509550866001600160401b03811115612aa657612aa66151a4565b604051908082528060200260200182016040528015612acf578160200160208202803683370190505b509450866001600160401b03811115612aea57612aea6151a4565b604051908082528060200260200182016040528015612b13578160200160208202803683370190505b509350866001600160401b03811115612b2e57612b2e6151a4565b604051908082528060200260200182016040528015612b57578160200160208202803683370190505b50925060005b87811015612cbb57600060068301612b75838c615a93565b81548110612b8557612b85615aab565b6000918252602090912060408051606081019091526003909202018054829060ff166002811115612bb857612bb8615541565b6002811115612bc957612bc9615541565b81526020016001820154815260200160028201548152505090508060000151888381518110612bfa57612bfa615aab565b60200260200101906002811115612c1357612c13615541565b90816002811115612c2657612c26615541565b815250508060200151878381518110612c4157612c41615aab565b6020026020010181815250508060400151868381518110612c6457612c64615aab565b6020908102919091010152604081015115801590612c86575080604001514210155b858381518110612c9857612c98615aab565b911515602092830291909101909101525080612cb381615ac1565b915050612b5d565b5050939792965093509350565b600080808080606086612d4f57600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612d28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d4c9190615a21565b96505b6001600160a01b03888116600090815260046020908152604080832060018101548c855260028201909352922054921697509060ff1660008981526003830160205260409020549015965060ff169450612da981896139e8565b925080600b018054612dba90615925565b80601f0160208091040260200160405190810160405280929190818152602001828054612de690615925565b8015612e335780601f10612e0857610100808354040283529160200191612e33565b820191906000526020600020905b815481529060010190602001808311612e1657829003601f168201915b50505050509150858015612e45575084155b8015612ec6575060005460405163fcbb371b60e01b8152600481018a90526101009091046001600160a01b03169063fcbb371b9060240161012060405180830381865afa158015612e9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ebe91906159a6565b60c001518310155b9350509295509295509295565b600080546001600160a01b038481168352600460209081526040808520888416865260079092528420612f1093909261010090910416908561416b565b50949350505050565b604051634ee5a1b960e01b815260040160405180910390fd5b6001600160a01b03808316600090815260046020526040902054839116612f6c576040516372898ae960e11b815260040160405180910390fd5b336000818152600760205260409020546001600160a01b0316612fa25760405163cf83d93d60e01b815260040160405180910390fd5b600080546001600160a01b03868116835260046020908152604080852033865260079092528420612fdd939092610100909104169087614132565b604080516001600160a01b03881681526020810183905291925033917f2ef606d064225d24c1514dc94907c134faee1237445c2f63f410cce0852b2054910160405180910390a280156130365761303660003383613c90565b505050505050565b6008818154811061205a57600080fd5b336000818152600760205260409020546001600160a01b03166130845760405163cf83d93d60e01b815260040160405180910390fd5b60008054338252600760205260409091206130ac9161010090046001600160a01b0316614379565b5050565b6004602052600090815260409020805460018201546006830154600b840180546001600160a01b0394851695949093169391926130ec90615925565b80601f016020809104026020016040519081016040528092919081815260200182805461311890615925565b80156131655780601f1061313a57610100808354040283529160200191613165565b820191906000526020600020905b81548152906001019060200180831161314857829003601f168201915b5050505050905084565b6001600160a01b038084166000908152600460205260409020548491166131a9576040516372898ae960e11b815260040160405180910390fd5b336000818152600760205260409020546001600160a01b03166131df5760405163cf83d93d60e01b815260040160405180910390fd5b600060019054906101000a90046001600160a01b03166001600160a01b031663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa158015613232573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132569190615a4f565b1561327457604051631e59ccd960e01b815260040160405180910390fd5b33600090815260076020908152604080832083546001600160a01b038a811686526004909452919093206132b292849261010090041690888861439d565b9350836132d257604051637bc90c0560e11b815260040160405180910390fd5b8060060160405180606001604052808760028111156132f3576132f3615541565b81526020810187905260400161330c42620d2f00615a93565b9052815460018181018455600093845260209093208251600390920201805492939092839160ff199091169083600281111561334a5761334a615541565b0217905550602082015181600101556040820151816002015550506133f66003600060019054906101000a90046001600160a01b03166001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156133bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133e19190615a21565b6133ec906001615a93565b60029190876144d3565b50600954604051635692619d60e11b81526001600160a01b0388811660048301529091169063ad24c33a90602401600060405180830381600087803b15801561343e57600080fd5b505af1158015613452573d6000803e3d6000fd5b50505060068201546001600160a01b038816915033907fb649014faa7a0e23357e091fb8a67a128c33dc9480f846f7e41cb3a6c9d806109061349690600190615adc565b60405190815260200160405180910390a3505050505050565b603081146134d057604051637477579960e11b815260040160405180910390fd5b6134de6020600083856158dd565b6134e791615907565b15801561350b57506134fd6030602083856158dd565b61350691615e35565b60801c155b1561352957604051634ee9493360e11b815260040160405180910390fd5b612044600b84018383614fe4565b600082815260098501602052604081205461356057600083815260098601602052604090208290555b6000838152600a8601602052604081205461357c906001615a93565b6000858152600a88016020526040902081905560e086015190915081108015906135c857506003860160006135b2866001615a93565b815260208101919091526040016000205460ff16155b15612f10576101008501516135dd9085615a93565b91505b81841015612f1057836135f281615ac1565b60008181526003890160205260409020805460ff1916600117905594506135e09050565b600080600061362686868661375d565b6006880155925050509392505050565b6001600160a01b03808616600090815260076020526040902080549091166136af5780546001600160a01b0387166001600160a01b031991821681178355600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180549092161790555b6001600160a01b03851660009081526004602052604090206136d6908290869086866145ae565b6136e4600260038685614676565b600954604051635692619d60e11b81526001600160a01b0387811660048301529091169063ad24c33a90602401600060405180830381600087803b15801561372b57600080fd5b505af115801561373f573d6000803e3d6000fd5b50505050505050505050565b613758838383600061475c565b505050565b6000808460060154905060006001856001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156137a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137cd9190615a21565b6137d79190615adc565b90508315806137ee5750806137ec8386615a93565b115b15613800576137fd8282615adc565b93505b60005b848110156138e757613816600184615a93565b60405163fcbb371b60e01b8152600481018290529093506000906001600160a01b0388169063fcbb371b9060240161012060405180830381865afa158015613862573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061388691906159a6565b905060006138958983876139fb565b509050806138a45750506138d5565b60a08201516138b45750506138d5565b6138c6818360a00151606460196147e2565b6138d09087615a93565b955050505b806138df81615ac1565b915050613803565b5050935093915050565b600080826138ff8587615a93565b106139115761390e8584615adc565b93505b8361391c8187615a93565b915091505b935093915050565b60006139df85600201600085600281111561394657613946615541565b600281111561395757613957615541565b81526020019081526020016000206000866001600160a01b03166001600160a01b03168152602001908152602001600020838760010160008760028111156139a1576139a1615541565b60028111156139b2576139b2615541565b8152602080820192909252604090810160009081206001600160a01b038b16825290925290209190613b77565b95945050505050565b600061173c600484016005850184613b77565b6000818152600284016020526040812054819060ff1680613a2c5750600083815260038601602052604090205460ff165b15613a3c57506000905080613921565b613a4685846139e8565b905060498310613a86578360c00151811015613a655760009150613921565b62989680613a7582610a34615e6e565b613a7f9190615e8d565b9150613b1f565b80613a9657506000905080613921565b613aa26019600a615f8b565b613ab385608001516064601961480e565b613abd9083615e6e565b613ac79190615e8d565b915081613ad75760009150613921565b613afa84606001518560400151613aee9190615e6e565b6301e13380601961480e565b613b049083615e6e565b9150613b126019600a615f8b565b613b1c9083615e8d565b91505b6000838152600a860160205260409020548015613b61576000848152600987016020526040902054613b5d84613b558484615adc565b8360196147e2565b9350505b50935093915050565b613758838383600161475c565b8254600090801580613ba557508285600081548110613b9857613b98615aab565b9060005260206000200154115b15613bb457600091505061173c565b8285613bc1600184615adc565b81548110613bd157613bd1615aab565b906000526020600020015411613c105783613bed600183615adc565b81548110613bfd57613bfd615aab565b906000526020600020015491505061173c565b600181118015613c4657508285613c28600284615adc565b81548110613c3857613c38615aab565b906000526020600020015411155b15613c575783613bed600283615adc565b6000613c66868560008561484a565b9050848181548110613c7a57613c7a615aab565b9060005260206000200154925050509392505050565b600080846002811115613ca557613ca5615541565b1415613d04576040516001600160a01b038416908390600081818185875af1925050503d8060008114613cf4576040519150601f19603f3d011682016040523d82523d6000602084013e613cf9565b606091505b505080915050613d86565b613d0d846148f2565b60405163a9059cbb60e01b81526001600160a01b03858116600483015260248201859052919091169063a9059cbb906044015b6020604051808303816000875af1158015613d5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d839190615a4f565b90505b806120445783604051630db5347560e11b8152600401613da69190615f97565b60405180910390fd5b81546001600160a01b031615613dd757604051621d934160e11b815260040160405180910390fd5b81546001600160a01b031916331782556130ac82826140b8565b6000836002811115613e0557613e05615541565b1415613e2b5780341461375857604051630fe5b06560e11b815260040160405180910390fd5b3415613e4a5760405163a745ac8560e01b815260040160405180910390fd5b6000613e55846148f2565b6040516323b872dd60e01b81526001600160a01b0385811660048301523060248301526044820185905291909116906323b872dd90606401613d40565b600080846003016000846002811115613ead57613ead615541565b6002811115613ebe57613ebe615541565b8152602081019190915260400160002054905080613ee057600091505061173c565b6000846001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613f20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f449190615a21565b90506000613f53600184615adc565b9050600081118015613fb6575081876003016000876002811115613f7957613f79615541565b6002811115613f8a57613f8a615541565b81526020019081526020016000208281548110613fa957613fa9615aab565b9060005260206000200154115b15613fc95780613fc581615fa5565b9150505b81876003016000876002811115613fe257613fe2615541565b6002811115613ff357613ff3615541565b8152602001908152602001600020828154811061401257614012615aab565b9060005260206000200154111561402f576000935050505061173c565b6000805b8281116140ac5788600401600088600281111561405257614052615541565b600281111561406357614063615541565b8152602001908152602001600020818154811061408257614082615aab565b9060005260206000200154826140989190615a93565b9150806140a481615ac1565b915050614033565b50979650505050505050565b6001600160a01b0381166140df57604051637138356f60e01b815260040160405180910390fd5b81546001600160a01b038281169116141561410d5760405163e037058f60e01b815260040160405180910390fd5b60019190910180546001600160a01b0319166001600160a01b03909216919091179055565b60008060006141438787878761416b565b86546001600160a01b0316600090815260058a01602052604090205592505050949350505050565b81546001600160a01b039081166000908152600586016020908152604080832054815163900cf0cf60e01b81529151939490938593600193928a169263900cf0cf92600480830193928290030181865afa1580156141cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141f19190615a21565b6141fb9190615adc565b90508315806142125750806142108386615a93565b115b15614224576142218282615adc565b93505b60005b8481101561436e5761423a600184615a93565b8654909350600090614259908a906001600160a01b0316600287613929565b8754614272908b906001600160a01b0316600188613929565b885461428b908c906001600160a01b0316600089613929565b6142959190615a93565b61429f9190615a93565b9050806142ac575061435c565b60008061432a8a6001600160a01b031663fcbb371b886040518263ffffffff1660e01b81526004016142e091815260200190565b61012060405180830381865afa1580156142fe573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061432291906159a6565b8a908861495d565b91509150816000141561433f5750505061435c565b61434c82848360196147e2565b6143569088615a93565b96505050505b8061436681615ac1565b915050614227565b505094509492505050565b614385828260016149ae565b614391828260026149ae565b6130ac828260006149ae565b60006144b98660020160008560028111156143ba576143ba615541565b60028111156143cb576143cb615541565b81526020808201929092526040908101600090812088546001600160a01b03908116835290845290829020825163900cf0cf60e01b815292519093918a169263900cf0cf9260048083019391928290030181865afa158015614431573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906144559190615a21565b614460906001615a93565b8489600101600088600281111561447957614479615541565b600281111561448a5761448a615541565b8152602080820192909252604090810160009081208b546001600160a01b0316825290925290209291906144d3565b9150816144c8575060006139df565b612f10848684614cd4565b835460009080158061450a5750856144ec600183615adc565b815481106144fc576144fc615aab565b906000526020600020015484105b1561452857604051630eae4c9760e01b815260040160405180910390fd5b6000614535878787614d54565b9050600086828154811061454b5761454b615aab565b90600052602060002001549050808511156145665780614568565b845b945084156145a2578487838154811061458357614583615aab565b90600052602060002001600082825461459c9190615adc565b90915550505b50929695505050505050565b61465d8560020160008460028111156145c9576145c9615541565b60028111156145da576145da615541565b81526020808201929092526040908101600090812087546001600160a01b031682529092528120908690849060018a019087600281111561461d5761461d615541565b600281111561462e5761462e615541565b8152602080820192909252604090810160009081208a546001600160a01b031682529092529020929190614676565b8454610d1c90849086906001600160a01b031684614f64565b83546001811180156146ad57508461468f600183615adc565b8154811061469f5761469f615aab565b906000526020600020015483105b80156146de5750846146c0600283615adc565b815481106146d0576146d0615aab565b906000526020600020015483105b156146fc57604051630eae4c9760e01b815260040160405180910390fd5b6000614709868686614d54565b8654925090505b81811015613036578285828154811061472b5761472b615aab565b9060005260206000200160008282546147449190615a93565b9091555081905061475481615ac1565b915050614710565b815160005b8181101561303657600084828151811061477d5761477d615aab565b6020026020010151905085811180156147ad5750600081815260028801602052604090205460ff16151584151514155b156147cf5760008181526002880160205260409020805460ff19168515151790555b50806147da81615ac1565b915050614761565b60006147ef82600a615f8b565b6147fa85858561480e565b6148049087615e6e565b6139df9190615e8d565b60008061481c836001615a93565b61482790600a615f8b565b6148319086615e6e565b9050600a61483f8583615e8d565b614804906005615a93565b6000818314156148665761485f600183615adc565b905061257d565b600060026148748486615a93565b61487e9190615e8d565b90508486828154811061489357614893615aab565b906000526020600020015411156148b8576148b08686868461484a565b91505061257d565b848682815481106148cb576148cb615aab565b906000526020600020015410156139df576148b086866148ec846001615a93565b8661484a565b6000600182600281111561490857614908615541565b141561491c57506001602960991b01919050565b600282600281111561493057614930615541565b141561494457506002602960991b01919050565b604051638698bf3760e01b815260040160405180910390fd5b60008061496b8585856139fb565b90925090508161497e5760009150613921565b60a0840151156139215761499a828560a00151606460196147e2565b6149a49083615adc565b9150935093915050565b60006149bb848484613e92565b9050806149c85750505050565b60008460030160008460028111156149e2576149e2615541565b60028111156149f3576149f3615541565b8152602001908152602001600020805490509050836001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614a45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a699190615a21565b856003016000856002811115614a8157614a81615541565b6002811115614a9257614a92615541565b8152602001908152602001600020600183614aad9190615adc565b81548110614abd57614abd615aab565b906000526020600020015411614b5857846003016000846002811115614ae557614ae5615541565b6002811115614af657614af6615541565b81526020019081526020016000206000614b109190615068565b846004016000846002811115614b2857614b28615541565b6002811115614b3957614b39615541565b81526020019081526020016000206000614b539190615068565b614cbd565b6040518060200160405280866003016000866002811115614b7b57614b7b615541565b6002811115614b8c57614b8c615541565b8152602001908152602001600020600184614ba79190615adc565b81548110614bb757614bb7615aab565b9060005260206000200154815250856003016000856002811115614bdd57614bdd615541565b6002811115614bee57614bee615541565b81526020810191909152604001600020614c09916001615089565b506040518060200160405280866004016000866002811115614c2d57614c2d615541565b6002811115614c3e57614c3e615541565b8152602001908152602001600020600184614c599190615adc565b81548110614c6957614c69615aab565b9060005260206000200154815250856004016000856002811115614c8f57614c8f615541565b6002811115614ca057614ca0615541565b81526020810191909152604001600020614cbb916001615089565b505b8454610d1c9084906001600160a01b031684613c90565b61204483600501836001600160a01b031663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614d19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614d3d9190615a21565b614d48906001615a93565b600486019190846144d3565b825460009080614d8f57505082546001818101855560008581526020808220909301849055845491820185558481529182200181905561173c565b6000614d9c600183615adc565b90506000868281548110614db257614db2615aab565b9060005260206000200154905080851415614dd15750915061173c9050565b80851115614e395786546001810188556000888152602090200185905585548690819084908110614e0457614e04615aab565b6000918252602080832090910154835460018181018655948452919092200155614e2f908390615a93565b935050505061173c565b600082118015614e6e575086614e50600184615adc565b81548110614e6057614e60615aab565b906000526020600020015485145b15614e7e57614e2f600183615adc565b86878381548110614e9157614e91615aab565b6000918252602080832090910154835460018101855593835291209091015585548690819084908110614ec657614ec6615aab565b6000918252602080832090910154835460018101855593835291209091015586548590889084908110614efb57614efb615aab565b6000918252602090912001558115614f395785614f19600184615adc565b81548110614f2957614f29615aab565b9060005260206000200154614f3c565b60005b868381548110614f4e57614f4e615aab565b60009182526020909120015550915061173c9050565b6001600160a01b038216600090815260088501602052604090205460ff16614fd2576001600160a01b038216600081815260088601602090815260408220805460ff191660019081179091556007880180549182018155835291200180546001600160a01b03191690911790555b61204460048501600586018584614676565b828054614ff090615925565b90600052602060002090601f0160209004810192826150125760008555615058565b82601f1061502b5782800160ff19823516178555615058565b82800160010185558215615058579182015b8281111561505857823582559160200191906001019061503d565b506150649291506150c4565b5090565b508054600082559060005260206000209081019061508691906150c4565b50565b828054828255906000526020600020908101928215615058579160200282015b828111156150585782518255916020019190600101906150a9565b5b8082111561506457600081556001016150c5565b600080602083850312156150ec57600080fd5b82356001600160401b038082111561510357600080fd5b818501915085601f83011261511757600080fd5b81358181111561512657600080fd5b86602082850101111561513857600080fd5b60209290920196919550909350505050565b6001600160a01b038116811461508657600080fd5b6000806040838503121561517257600080fd5b823561517d8161514a565b946020939093013593505050565b60006020828403121561519d57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b03811182821017156151dd576151dd6151a4565b60405290565b604051601f8201601f191681016001600160401b038111828210171561520b5761520b6151a4565b604052919050565b60006001600160401b0382111561522c5761522c6151a4565b5060051b60200190565b6000806040838503121561524957600080fd5b82356152548161514a565b91506020838101356001600160401b0381111561527057600080fd5b8401601f8101861361528157600080fd5b803561529461528f82615213565b6151e3565b81815260059190911b820183019083810190888311156152b357600080fd5b928401925b828410156152d1578335825292840192908401906152b8565b80955050505050509250929050565b600080604083850312156152f357600080fd5b50508035926020909101359150565b600081518084526020808501945080840160005b8381101561533b5781516001600160a01b031687529582019590820190600101615316565b509495945050505050565b6040815260006153596040830185615302565b90508260208301529392505050565b6000806000806080858703121561537e57600080fd5b84356153898161514a565b966020860135965060408601359560600135945092505050565b600081518084526020808501945080840160005b8381101561533b578151875295820195908201906001016153b7565b60a0815260006153e660a0830188615302565b82810360208401526153f881886153a3565b9050828103604084015261540c81876153a3565b9050828103606084015261542081866153a3565b9150508260808301529695505050505050565b6000806040838503121561544657600080fd5b82356001600160401b0381111561545c57600080fd5b8301601f8101851361546d57600080fd5b8035602061547d61528f83615213565b82815260059290921b8301810191818101908884111561549c57600080fd5b938201935b838510156154c35784356154b48161514a565b825293820193908201906154a1565b98969091013596505050505050565b6060815260006154e56060830186615302565b82810360208401526154f781866153a3565b915050826040830152949350505050565b6000806040838503121561551b57600080fd5b82356155268161514a565b915060208301356155368161514a565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b6003811061557557634e487b7160e01b600052602160045260246000fd5b9052565b608081016155878287615557565b846020830152836040830152821515606083015295945050505050565b6000602082840312156155b657600080fd5b813561173c8161514a565b6000806000606084860312156155d657600080fd5b505081359360208301359350604090920135919050565b60005b838110156156085781810151838201526020016155f0565b838111156120445750506000910152565b600081518084526156318160208601602086016155ed565b601f01601f19169290920160200192915050565b600081518084526020808501945080840160005b8381101561533b578151151587529582019590820190600101615659565b60c08152600061568a60c0830189615302565b60208382038185015261569d828a615302565b915083820360408501526156b182896153a3565b915083820360608501528187518084528284019150828160051b850101838a0160005b8381101561570257601f198784030185526156f0838351615619565b948601949250908501906001016156d4565b50508681036080880152615716818a615645565b955050505050508260a0830152979650505050505050565b60008060006060848603121561574357600080fd5b833561574e8161514a565b925060208401356003811061576257600080fd5b929592945050506040919091013590565b60008060006060848603121561578857600080fd5b83356157938161514a565b95602085013595506040909401359392505050565b60a0808252865190820181905260009060209060c0840190828a01845b828110156157e8576157d8848351615557565b92840192908401906001016157c5565b505050838103828501526157fc81896153a3565b915050828103604084015261581181876153a3565b905082810360608401526154208186615645565b60018060a01b038716815285151560208201528415156040820152831515606082015282608082015260c060a0820152600061586460c0830184615619565b98975050505050505050565b60008060006060848603121561588557600080fd5b83356158908161514a565b925060208401356157628161514a565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906158d390830184615619565b9695505050505050565b600080858511156158ed57600080fd5b838611156158fa57600080fd5b5050820193919092039150565b8035602083101561117957600019602084900360031b1b1692915050565b600181811c9082168061593957607f821691505b6020821081141561595a57634e487b7160e01b600052602260045260246000fd5b50919050565b6040815260006159736040830186615619565b8281036020840152838152838560208301376000602085830101526020601f19601f860116820101915050949350505050565b600061012082840312156159b957600080fd5b6159c16151ba565b825181526020830151602082015260408301516040820152606083015160608201526080830151608082015260a083015160a082015260c083015160c082015260e083015160e08201526101008084015181830152508091505092915050565b600060208284031215615a3357600080fd5b5051919050565b80518015158114615a4a57600080fd5b919050565b600060208284031215615a6157600080fd5b61173c82615a3a565b60208152600061173c60208301846153a3565b634e487b7160e01b600052601160045260246000fd5b60008219821115615aa657615aa6615a7d565b500190565b634e487b7160e01b600052603260045260246000fd5b6000600019821415615ad557615ad5615a7d565b5060010190565b600082821015615aee57615aee615a7d565b500390565b600082601f830112615b0457600080fd5b81516020615b1461528f83615213565b82815260059290921b84018101918181019086841115615b3357600080fd5b8286015b84811015615b57578051615b4a8161514a565b8352918301918301615b37565b509695505050505050565b600082601f830112615b7357600080fd5b81516020615b8361528f83615213565b82815260059290921b84018101918181019086841115615ba257600080fd5b8286015b84811015615b5757615bb781615a3a565b8352918301918301615ba6565b600082601f830112615bd557600080fd5b81516020615be561528f83615213565b82815260059290921b84018101918181019086841115615c0457600080fd5b8286015b84811015615b575780518352918301918301615c08565b6000601f8381840112615c3157600080fd5b82516020615c4161528f83615213565b82815260059290921b85018101918181019087841115615c6057600080fd5b8287015b848110156140ac5780516001600160401b0380821115615c845760008081fd5b818a0191508a603f830112615c995760008081fd5b85820151604082821115615caf57615caf6151a4565b615cc0828b01601f191689016151e3565b92508183528c81838601011115615cd75760008081fd5b615ce6828985018387016155ed565b5050845250918301918301615c64565b600080600080600080600080610100898b031215615d1357600080fd5b88516001600160401b0380821115615d2a57600080fd5b615d368c838d01615af3565b995060208b0151915080821115615d4c57600080fd5b615d588c838d01615af3565b985060408b0151915080821115615d6e57600080fd5b615d7a8c838d01615b62565b975060608b0151915080821115615d9057600080fd5b615d9c8c838d01615b62565b965060808b0151915080821115615db257600080fd5b615dbe8c838d01615bc4565b955060a08b0151915080821115615dd457600080fd5b615de08c838d01615c1f565b945060c08b0151915080821115615df657600080fd5b50615e038b828c01615b62565b92505060e089015190509295985092959890939650565b60408101615e288285615557565b8260208301529392505050565b6fffffffffffffffffffffffffffffffff198135818116916010851015615e665780818660100360031b1b83161692505b505092915050565b6000816000190483118215151615615e8857615e88615a7d565b500290565b600082615eaa57634e487b7160e01b600052601260045260246000fd5b500490565b600181815b8085111561126d578160001904821115615ed057615ed0615a7d565b80851615615edd57918102915b93841c9390800290615eb4565b600082615ef957506001611179565b81615f0657506000611179565b8160018114615f1c5760028114615f2657615f42565b6001915050611179565b60ff841115615f3757615f37615a7d565b50506001821b611179565b5060208310610133831016604e8410600b8410161715615f65575081810a611179565b615f6f8383615eaf565b8060001904821115615f8357615f83615a7d565b029392505050565b600061173c8383615eea565b602081016111798284615557565b600081615fb457615fb4615a7d565b50600019019056fea2646970667358221220b3c1bc7342f8d8fdc7ee0ce718b653ec3c77c2df6b236ed9e233f78255c3dd6a64736f6c634300080c0033",
}

// StakemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakemanagerMetaData.ABI instead.
var StakemanagerABI = StakemanagerMetaData.ABI

// StakemanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakemanagerMetaData.Bin instead.
var StakemanagerBin = StakemanagerMetaData.Bin

// DeployStakemanager deploys a new Ethereum contract, binding an instance of Stakemanager to it.
func DeployStakemanager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stakemanager, error) {
	parsed, err := StakemanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakemanagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stakemanager{StakemanagerCaller: StakemanagerCaller{contract: contract}, StakemanagerTransactor: StakemanagerTransactor{contract: contract}, StakemanagerFilterer: StakemanagerFilterer{contract: contract}}, nil
}

// Stakemanager is an auto generated Go binding around an Ethereum contract.
type Stakemanager struct {
	StakemanagerCaller     // Read-only binding to the contract
	StakemanagerTransactor // Write-only binding to the contract
	StakemanagerFilterer   // Log filterer for contract events
}

// StakemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakemanagerSession struct {
	Contract     *Stakemanager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakemanagerCallerSession struct {
	Contract *StakemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakemanagerTransactorSession struct {
	Contract     *StakemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakemanagerRaw struct {
	Contract *Stakemanager // Generic contract binding to access the raw methods on
}

// StakemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakemanagerCallerRaw struct {
	Contract *StakemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakemanagerTransactorRaw struct {
	Contract *StakemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakemanager creates a new instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanager(address common.Address, backend bind.ContractBackend) (*Stakemanager, error) {
	contract, err := bindStakemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakemanager{StakemanagerCaller: StakemanagerCaller{contract: contract}, StakemanagerTransactor: StakemanagerTransactor{contract: contract}, StakemanagerFilterer: StakemanagerFilterer{contract: contract}}, nil
}

// NewStakemanagerCaller creates a new read-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerCaller(address common.Address, caller bind.ContractCaller) (*StakemanagerCaller, error) {
	contract, err := bindStakemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerCaller{contract: contract}, nil
}

// NewStakemanagerTransactor creates a new write-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakemanagerTransactor, error) {
	contract, err := bindStakemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerTransactor{contract: contract}, nil
}

// NewStakemanagerFilterer creates a new log filterer instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakemanagerFilterer, error) {
	contract, err := bindStakemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakemanagerFilterer{contract: contract}, nil
}

// bindStakemanager binds a generic wrapper to an already deployed contract.
func bindStakemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakemanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.StakemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transact(opts, method, params...)
}

// Allowlist is a free data retrieval call binding the contract method 0x2b47da52.
//
// Solidity: function allowlist() view returns(address)
func (_Stakemanager *StakemanagerCaller) Allowlist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "allowlist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allowlist is a free data retrieval call binding the contract method 0x2b47da52.
//
// Solidity: function allowlist() view returns(address)
func (_Stakemanager *StakemanagerSession) Allowlist() (common.Address, error) {
	return _Stakemanager.Contract.Allowlist(&_Stakemanager.CallOpts)
}

// Allowlist is a free data retrieval call binding the contract method 0x2b47da52.
//
// Solidity: function allowlist() view returns(address)
func (_Stakemanager *StakemanagerCallerSession) Allowlist() (common.Address, error) {
	return _Stakemanager.Contract.Allowlist(&_Stakemanager.CallOpts)
}

// BlsPublicKeyToOwner is a free data retrieval call binding the contract method 0xe4b2477b.
//
// Solidity: function blsPublicKeyToOwner(bytes32 ) view returns(address)
func (_Stakemanager *StakemanagerCaller) BlsPublicKeyToOwner(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "blsPublicKeyToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsPublicKeyToOwner is a free data retrieval call binding the contract method 0xe4b2477b.
//
// Solidity: function blsPublicKeyToOwner(bytes32 ) view returns(address)
func (_Stakemanager *StakemanagerSession) BlsPublicKeyToOwner(arg0 [32]byte) (common.Address, error) {
	return _Stakemanager.Contract.BlsPublicKeyToOwner(&_Stakemanager.CallOpts, arg0)
}

// BlsPublicKeyToOwner is a free data retrieval call binding the contract method 0xe4b2477b.
//
// Solidity: function blsPublicKeyToOwner(bytes32 ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) BlsPublicKeyToOwner(arg0 [32]byte) (common.Address, error) {
	return _Stakemanager.Contract.BlsPublicKeyToOwner(&_Stakemanager.CallOpts, arg0)
}

// CandidateManager is a free data retrieval call binding the contract method 0xa6a41f44.
//
// Solidity: function candidateManager() view returns(address)
func (_Stakemanager *StakemanagerCaller) CandidateManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "candidateManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CandidateManager is a free data retrieval call binding the contract method 0xa6a41f44.
//
// Solidity: function candidateManager() view returns(address)
func (_Stakemanager *StakemanagerSession) CandidateManager() (common.Address, error) {
	return _Stakemanager.Contract.CandidateManager(&_Stakemanager.CallOpts)
}

// CandidateManager is a free data retrieval call binding the contract method 0xa6a41f44.
//
// Solidity: function candidateManager() view returns(address)
func (_Stakemanager *StakemanagerCallerSession) CandidateManager() (common.Address, error) {
	return _Stakemanager.Contract.CandidateManager(&_Stakemanager.CallOpts)
}

// Environment is a free data retrieval call binding the contract method 0x74e2b63c.
//
// Solidity: function environment() view returns(address)
func (_Stakemanager *StakemanagerCaller) Environment(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "environment")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Environment is a free data retrieval call binding the contract method 0x74e2b63c.
//
// Solidity: function environment() view returns(address)
func (_Stakemanager *StakemanagerSession) Environment() (common.Address, error) {
	return _Stakemanager.Contract.Environment(&_Stakemanager.CallOpts)
}

// Environment is a free data retrieval call binding the contract method 0x74e2b63c.
//
// Solidity: function environment() view returns(address)
func (_Stakemanager *StakemanagerCallerSession) Environment() (common.Address, error) {
	return _Stakemanager.Contract.Environment(&_Stakemanager.CallOpts)
}

// GetBlockAndSlashes is a free data retrieval call binding the contract method 0x22226367.
//
// Solidity: function getBlockAndSlashes(address validator, uint256 epoch) view returns(uint256 blocks, uint256 slashes)
func (_Stakemanager *StakemanagerCaller) GetBlockAndSlashes(opts *bind.CallOpts, validator common.Address, epoch *big.Int) (struct {
	Blocks  *big.Int
	Slashes *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getBlockAndSlashes", validator, epoch)

	outstruct := new(struct {
		Blocks  *big.Int
		Slashes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Blocks = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slashes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBlockAndSlashes is a free data retrieval call binding the contract method 0x22226367.
//
// Solidity: function getBlockAndSlashes(address validator, uint256 epoch) view returns(uint256 blocks, uint256 slashes)
func (_Stakemanager *StakemanagerSession) GetBlockAndSlashes(validator common.Address, epoch *big.Int) (struct {
	Blocks  *big.Int
	Slashes *big.Int
}, error) {
	return _Stakemanager.Contract.GetBlockAndSlashes(&_Stakemanager.CallOpts, validator, epoch)
}

// GetBlockAndSlashes is a free data retrieval call binding the contract method 0x22226367.
//
// Solidity: function getBlockAndSlashes(address validator, uint256 epoch) view returns(uint256 blocks, uint256 slashes)
func (_Stakemanager *StakemanagerCallerSession) GetBlockAndSlashes(validator common.Address, epoch *big.Int) (struct {
	Blocks  *big.Int
	Slashes *big.Int
}, error) {
	return _Stakemanager.Contract.GetBlockAndSlashes(&_Stakemanager.CallOpts, validator, epoch)
}

// GetCommissions is a free data retrieval call binding the contract method 0x195afea1.
//
// Solidity: function getCommissions(address validator, uint256 epochs) view returns(uint256 commissions)
func (_Stakemanager *StakemanagerCaller) GetCommissions(opts *bind.CallOpts, validator common.Address, epochs *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getCommissions", validator, epochs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommissions is a free data retrieval call binding the contract method 0x195afea1.
//
// Solidity: function getCommissions(address validator, uint256 epochs) view returns(uint256 commissions)
func (_Stakemanager *StakemanagerSession) GetCommissions(validator common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetCommissions(&_Stakemanager.CallOpts, validator, epochs)
}

// GetCommissions is a free data retrieval call binding the contract method 0x195afea1.
//
// Solidity: function getCommissions(address validator, uint256 epochs) view returns(uint256 commissions)
func (_Stakemanager *StakemanagerCallerSession) GetCommissions(validator common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetCommissions(&_Stakemanager.CallOpts, validator, epochs)
}

// GetLockedUnstake is a free data retrieval call binding the contract method 0x5c4fc4c5.
//
// Solidity: function getLockedUnstake(address staker, uint256 lockedUnstake) view returns(uint8 token, uint256 amount, uint256 unlockTime, bool claimable)
func (_Stakemanager *StakemanagerCaller) GetLockedUnstake(opts *bind.CallOpts, staker common.Address, lockedUnstake *big.Int) (struct {
	Token      uint8
	Amount     *big.Int
	UnlockTime *big.Int
	Claimable  bool
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getLockedUnstake", staker, lockedUnstake)

	outstruct := new(struct {
		Token      uint8
		Amount     *big.Int
		UnlockTime *big.Int
		Claimable  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Claimable = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetLockedUnstake is a free data retrieval call binding the contract method 0x5c4fc4c5.
//
// Solidity: function getLockedUnstake(address staker, uint256 lockedUnstake) view returns(uint8 token, uint256 amount, uint256 unlockTime, bool claimable)
func (_Stakemanager *StakemanagerSession) GetLockedUnstake(staker common.Address, lockedUnstake *big.Int) (struct {
	Token      uint8
	Amount     *big.Int
	UnlockTime *big.Int
	Claimable  bool
}, error) {
	return _Stakemanager.Contract.GetLockedUnstake(&_Stakemanager.CallOpts, staker, lockedUnstake)
}

// GetLockedUnstake is a free data retrieval call binding the contract method 0x5c4fc4c5.
//
// Solidity: function getLockedUnstake(address staker, uint256 lockedUnstake) view returns(uint8 token, uint256 amount, uint256 unlockTime, bool claimable)
func (_Stakemanager *StakemanagerCallerSession) GetLockedUnstake(staker common.Address, lockedUnstake *big.Int) (struct {
	Token      uint8
	Amount     *big.Int
	UnlockTime *big.Int
	Claimable  bool
}, error) {
	return _Stakemanager.Contract.GetLockedUnstake(&_Stakemanager.CallOpts, staker, lockedUnstake)
}

// GetLockedUnstakeCount is a free data retrieval call binding the contract method 0xdf93c842.
//
// Solidity: function getLockedUnstakeCount(address staker) view returns(uint256 count)
func (_Stakemanager *StakemanagerCaller) GetLockedUnstakeCount(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getLockedUnstakeCount", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedUnstakeCount is a free data retrieval call binding the contract method 0xdf93c842.
//
// Solidity: function getLockedUnstakeCount(address staker) view returns(uint256 count)
func (_Stakemanager *StakemanagerSession) GetLockedUnstakeCount(staker common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.GetLockedUnstakeCount(&_Stakemanager.CallOpts, staker)
}

// GetLockedUnstakeCount is a free data retrieval call binding the contract method 0xdf93c842.
//
// Solidity: function getLockedUnstakeCount(address staker) view returns(uint256 count)
func (_Stakemanager *StakemanagerCallerSession) GetLockedUnstakeCount(staker common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.GetLockedUnstakeCount(&_Stakemanager.CallOpts, staker)
}

// GetLockedUnstakes is a free data retrieval call binding the contract method 0xd0051adf.
//
// Solidity: function getLockedUnstakes(address staker, uint256 cursor, uint256 howMany) view returns(uint8[] tokens, uint256[] amounts, uint256[] unlockTimes, bool[] claimable, uint256 newCursor)
func (_Stakemanager *StakemanagerCaller) GetLockedUnstakes(opts *bind.CallOpts, staker common.Address, cursor *big.Int, howMany *big.Int) (struct {
	Tokens      []uint8
	Amounts     []*big.Int
	UnlockTimes []*big.Int
	Claimable   []bool
	NewCursor   *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getLockedUnstakes", staker, cursor, howMany)

	outstruct := new(struct {
		Tokens      []uint8
		Amounts     []*big.Int
		UnlockTimes []*big.Int
		Claimable   []bool
		NewCursor   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.UnlockTimes = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Claimable = *abi.ConvertType(out[3], new([]bool)).(*[]bool)
	outstruct.NewCursor = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLockedUnstakes is a free data retrieval call binding the contract method 0xd0051adf.
//
// Solidity: function getLockedUnstakes(address staker, uint256 cursor, uint256 howMany) view returns(uint8[] tokens, uint256[] amounts, uint256[] unlockTimes, bool[] claimable, uint256 newCursor)
func (_Stakemanager *StakemanagerSession) GetLockedUnstakes(staker common.Address, cursor *big.Int, howMany *big.Int) (struct {
	Tokens      []uint8
	Amounts     []*big.Int
	UnlockTimes []*big.Int
	Claimable   []bool
	NewCursor   *big.Int
}, error) {
	return _Stakemanager.Contract.GetLockedUnstakes(&_Stakemanager.CallOpts, staker, cursor, howMany)
}

// GetLockedUnstakes is a free data retrieval call binding the contract method 0xd0051adf.
//
// Solidity: function getLockedUnstakes(address staker, uint256 cursor, uint256 howMany) view returns(uint8[] tokens, uint256[] amounts, uint256[] unlockTimes, bool[] claimable, uint256 newCursor)
func (_Stakemanager *StakemanagerCallerSession) GetLockedUnstakes(staker common.Address, cursor *big.Int, howMany *big.Int) (struct {
	Tokens      []uint8
	Amounts     []*big.Int
	UnlockTimes []*big.Int
	Claimable   []bool
	NewCursor   *big.Int
}, error) {
	return _Stakemanager.Contract.GetLockedUnstakes(&_Stakemanager.CallOpts, staker, cursor, howMany)
}

// GetOperatorStakes is a free data retrieval call binding the contract method 0x9c508219.
//
// Solidity: function getOperatorStakes(address operator, uint256 epoch) view returns(uint256 stakes)
func (_Stakemanager *StakemanagerCaller) GetOperatorStakes(opts *bind.CallOpts, operator common.Address, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getOperatorStakes", operator, epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorStakes is a free data retrieval call binding the contract method 0x9c508219.
//
// Solidity: function getOperatorStakes(address operator, uint256 epoch) view returns(uint256 stakes)
func (_Stakemanager *StakemanagerSession) GetOperatorStakes(operator common.Address, epoch *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetOperatorStakes(&_Stakemanager.CallOpts, operator, epoch)
}

// GetOperatorStakes is a free data retrieval call binding the contract method 0x9c508219.
//
// Solidity: function getOperatorStakes(address operator, uint256 epoch) view returns(uint256 stakes)
func (_Stakemanager *StakemanagerCallerSession) GetOperatorStakes(operator common.Address, epoch *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetOperatorStakes(&_Stakemanager.CallOpts, operator, epoch)
}

// GetRewards is a free data retrieval call binding the contract method 0xdbd61d87.
//
// Solidity: function getRewards(address staker, address validator, uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerCaller) GetRewards(opts *bind.CallOpts, staker common.Address, validator common.Address, epochs *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getRewards", staker, validator, epochs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewards is a free data retrieval call binding the contract method 0xdbd61d87.
//
// Solidity: function getRewards(address staker, address validator, uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerSession) GetRewards(staker common.Address, validator common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetRewards(&_Stakemanager.CallOpts, staker, validator, epochs)
}

// GetRewards is a free data retrieval call binding the contract method 0xdbd61d87.
//
// Solidity: function getRewards(address staker, address validator, uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerCallerSession) GetRewards(staker common.Address, validator common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetRewards(&_Stakemanager.CallOpts, staker, validator, epochs)
}

// GetStakerStakes is a free data retrieval call binding the contract method 0x2b42ed8c.
//
// Solidity: function getStakerStakes(address staker, uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] _validators, uint256[] oasStakes, uint256[] woasStakes, uint256[] soasStakes, uint256 newCursor)
func (_Stakemanager *StakemanagerCaller) GetStakerStakes(opts *bind.CallOpts, staker common.Address, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Validators []common.Address
	OasStakes  []*big.Int
	WoasStakes []*big.Int
	SoasStakes []*big.Int
	NewCursor  *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getStakerStakes", staker, epoch, cursor, howMany)

	outstruct := new(struct {
		Validators []common.Address
		OasStakes  []*big.Int
		WoasStakes []*big.Int
		SoasStakes []*big.Int
		NewCursor  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OasStakes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.WoasStakes = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.SoasStakes = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.NewCursor = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakerStakes is a free data retrieval call binding the contract method 0x2b42ed8c.
//
// Solidity: function getStakerStakes(address staker, uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] _validators, uint256[] oasStakes, uint256[] woasStakes, uint256[] soasStakes, uint256 newCursor)
func (_Stakemanager *StakemanagerSession) GetStakerStakes(staker common.Address, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Validators []common.Address
	OasStakes  []*big.Int
	WoasStakes []*big.Int
	SoasStakes []*big.Int
	NewCursor  *big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerStakes(&_Stakemanager.CallOpts, staker, epoch, cursor, howMany)
}

// GetStakerStakes is a free data retrieval call binding the contract method 0x2b42ed8c.
//
// Solidity: function getStakerStakes(address staker, uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] _validators, uint256[] oasStakes, uint256[] woasStakes, uint256[] soasStakes, uint256 newCursor)
func (_Stakemanager *StakemanagerCallerSession) GetStakerStakes(staker common.Address, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Validators []common.Address
	OasStakes  []*big.Int
	WoasStakes []*big.Int
	SoasStakes []*big.Int
	NewCursor  *big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerStakes(&_Stakemanager.CallOpts, staker, epoch, cursor, howMany)
}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 cursor, uint256 howMany) view returns(address[] _stakers, uint256 newCursor)
func (_Stakemanager *StakemanagerCaller) GetStakers(opts *bind.CallOpts, cursor *big.Int, howMany *big.Int) (struct {
	Stakers   []common.Address
	NewCursor *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getStakers", cursor, howMany)

	outstruct := new(struct {
		Stakers   []common.Address
		NewCursor *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stakers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.NewCursor = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 cursor, uint256 howMany) view returns(address[] _stakers, uint256 newCursor)
func (_Stakemanager *StakemanagerSession) GetStakers(cursor *big.Int, howMany *big.Int) (struct {
	Stakers   []common.Address
	NewCursor *big.Int
}, error) {
	return _Stakemanager.Contract.GetStakers(&_Stakemanager.CallOpts, cursor, howMany)
}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 cursor, uint256 howMany) view returns(address[] _stakers, uint256 newCursor)
func (_Stakemanager *StakemanagerCallerSession) GetStakers(cursor *big.Int, howMany *big.Int) (struct {
	Stakers   []common.Address
	NewCursor *big.Int
}, error) {
	return _Stakemanager.Contract.GetStakers(&_Stakemanager.CallOpts, cursor, howMany)
}

// GetTotalRewards is a free data retrieval call binding the contract method 0x33f32d78.
//
// Solidity: function getTotalRewards(address[] _validators, uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerCaller) GetTotalRewards(opts *bind.CallOpts, _validators []common.Address, epochs *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getTotalRewards", _validators, epochs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRewards is a free data retrieval call binding the contract method 0x33f32d78.
//
// Solidity: function getTotalRewards(address[] _validators, uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerSession) GetTotalRewards(_validators []common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalRewards(&_Stakemanager.CallOpts, _validators, epochs)
}

// GetTotalRewards is a free data retrieval call binding the contract method 0x33f32d78.
//
// Solidity: function getTotalRewards(address[] _validators, uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerCallerSession) GetTotalRewards(_validators []common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalRewards(&_Stakemanager.CallOpts, _validators, epochs)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x45367f23.
//
// Solidity: function getTotalStake(uint256 epoch) view returns(uint256 amounts)
func (_Stakemanager *StakemanagerCaller) GetTotalStake(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getTotalStake", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x45367f23.
//
// Solidity: function getTotalStake(uint256 epoch) view returns(uint256 amounts)
func (_Stakemanager *StakemanagerSession) GetTotalStake(epoch *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalStake(&_Stakemanager.CallOpts, epoch)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x45367f23.
//
// Solidity: function getTotalStake(uint256 epoch) view returns(uint256 amounts)
func (_Stakemanager *StakemanagerCallerSession) GetTotalStake(epoch *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalStake(&_Stakemanager.CallOpts, epoch)
}

// GetUnstakes is a free data retrieval call binding the contract method 0x88325234.
//
// Solidity: function getUnstakes(address staker) view returns(uint256 oasUnstakes, uint256 woasUnstakes, uint256 soasUnstakes)
func (_Stakemanager *StakemanagerCaller) GetUnstakes(opts *bind.CallOpts, staker common.Address) (struct {
	OasUnstakes  *big.Int
	WoasUnstakes *big.Int
	SoasUnstakes *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getUnstakes", staker)

	outstruct := new(struct {
		OasUnstakes  *big.Int
		WoasUnstakes *big.Int
		SoasUnstakes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OasUnstakes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WoasUnstakes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SoasUnstakes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUnstakes is a free data retrieval call binding the contract method 0x88325234.
//
// Solidity: function getUnstakes(address staker) view returns(uint256 oasUnstakes, uint256 woasUnstakes, uint256 soasUnstakes)
func (_Stakemanager *StakemanagerSession) GetUnstakes(staker common.Address) (struct {
	OasUnstakes  *big.Int
	WoasUnstakes *big.Int
	SoasUnstakes *big.Int
}, error) {
	return _Stakemanager.Contract.GetUnstakes(&_Stakemanager.CallOpts, staker)
}

// GetUnstakes is a free data retrieval call binding the contract method 0x88325234.
//
// Solidity: function getUnstakes(address staker) view returns(uint256 oasUnstakes, uint256 woasUnstakes, uint256 soasUnstakes)
func (_Stakemanager *StakemanagerCallerSession) GetUnstakes(staker common.Address) (struct {
	OasUnstakes  *big.Int
	WoasUnstakes *big.Int
	SoasUnstakes *big.Int
}, error) {
	return _Stakemanager.Contract.GetUnstakes(&_Stakemanager.CallOpts, staker)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xd1f18ee1.
//
// Solidity: function getValidatorInfo(address validator, uint256 epoch) view returns(address operator, bool active, bool jailed, bool candidate, uint256 stakes, bytes blsPublicKey)
func (_Stakemanager *StakemanagerCaller) GetValidatorInfo(opts *bind.CallOpts, validator common.Address, epoch *big.Int) (struct {
	Operator     common.Address
	Active       bool
	Jailed       bool
	Candidate    bool
	Stakes       *big.Int
	BlsPublicKey []byte
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorInfo", validator, epoch)

	outstruct := new(struct {
		Operator     common.Address
		Active       bool
		Jailed       bool
		Candidate    bool
		Stakes       *big.Int
		BlsPublicKey []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Active = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Jailed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Candidate = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Stakes = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BlsPublicKey = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xd1f18ee1.
//
// Solidity: function getValidatorInfo(address validator, uint256 epoch) view returns(address operator, bool active, bool jailed, bool candidate, uint256 stakes, bytes blsPublicKey)
func (_Stakemanager *StakemanagerSession) GetValidatorInfo(validator common.Address, epoch *big.Int) (struct {
	Operator     common.Address
	Active       bool
	Jailed       bool
	Candidate    bool
	Stakes       *big.Int
	BlsPublicKey []byte
}, error) {
	return _Stakemanager.Contract.GetValidatorInfo(&_Stakemanager.CallOpts, validator, epoch)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xd1f18ee1.
//
// Solidity: function getValidatorInfo(address validator, uint256 epoch) view returns(address operator, bool active, bool jailed, bool candidate, uint256 stakes, bytes blsPublicKey)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorInfo(validator common.Address, epoch *big.Int) (struct {
	Operator     common.Address
	Active       bool
	Jailed       bool
	Candidate    bool
	Stakes       *big.Int
	BlsPublicKey []byte
}, error) {
	return _Stakemanager.Contract.GetValidatorInfo(&_Stakemanager.CallOpts, validator, epoch)
}

// GetValidatorOwners is a free data retrieval call binding the contract method 0x2168e8b4.
//
// Solidity: function getValidatorOwners(uint256 cursor, uint256 howMany) view returns(address[] owners, uint256 newCursor)
func (_Stakemanager *StakemanagerCaller) GetValidatorOwners(opts *bind.CallOpts, cursor *big.Int, howMany *big.Int) (struct {
	Owners    []common.Address
	NewCursor *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorOwners", cursor, howMany)

	outstruct := new(struct {
		Owners    []common.Address
		NewCursor *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.NewCursor = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorOwners is a free data retrieval call binding the contract method 0x2168e8b4.
//
// Solidity: function getValidatorOwners(uint256 cursor, uint256 howMany) view returns(address[] owners, uint256 newCursor)
func (_Stakemanager *StakemanagerSession) GetValidatorOwners(cursor *big.Int, howMany *big.Int) (struct {
	Owners    []common.Address
	NewCursor *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorOwners(&_Stakemanager.CallOpts, cursor, howMany)
}

// GetValidatorOwners is a free data retrieval call binding the contract method 0x2168e8b4.
//
// Solidity: function getValidatorOwners(uint256 cursor, uint256 howMany) view returns(address[] owners, uint256 newCursor)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorOwners(cursor *big.Int, howMany *big.Int) (struct {
	Owners    []common.Address
	NewCursor *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorOwners(&_Stakemanager.CallOpts, cursor, howMany)
}

// GetValidatorStakes is a free data retrieval call binding the contract method 0x2ee462b3.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch) view returns(uint256 stakes)
func (_Stakemanager *StakemanagerCaller) GetValidatorStakes(opts *bind.CallOpts, validator common.Address, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorStakes", validator, epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorStakes is a free data retrieval call binding the contract method 0x2ee462b3.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch) view returns(uint256 stakes)
func (_Stakemanager *StakemanagerSession) GetValidatorStakes(validator common.Address, epoch *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetValidatorStakes(&_Stakemanager.CallOpts, validator, epoch)
}

// GetValidatorStakes is a free data retrieval call binding the contract method 0x2ee462b3.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch) view returns(uint256 stakes)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorStakes(validator common.Address, epoch *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetValidatorStakes(&_Stakemanager.CallOpts, validator, epoch)
}

// GetValidatorStakes0 is a free data retrieval call binding the contract method 0x46dfce7b.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] _stakers, uint256[] stakes, uint256 newCursor)
func (_Stakemanager *StakemanagerCaller) GetValidatorStakes0(opts *bind.CallOpts, validator common.Address, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Stakers   []common.Address
	Stakes    []*big.Int
	NewCursor *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorStakes0", validator, epoch, cursor, howMany)

	outstruct := new(struct {
		Stakers   []common.Address
		Stakes    []*big.Int
		NewCursor *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stakers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.NewCursor = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorStakes0 is a free data retrieval call binding the contract method 0x46dfce7b.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] _stakers, uint256[] stakes, uint256 newCursor)
func (_Stakemanager *StakemanagerSession) GetValidatorStakes0(validator common.Address, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Stakers   []common.Address
	Stakes    []*big.Int
	NewCursor *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorStakes0(&_Stakemanager.CallOpts, validator, epoch, cursor, howMany)
}

// GetValidatorStakes0 is a free data retrieval call binding the contract method 0x46dfce7b.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] _stakers, uint256[] stakes, uint256 newCursor)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorStakes0(validator common.Address, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Stakers   []common.Address
	Stakes    []*big.Int
	NewCursor *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorStakes0(&_Stakemanager.CallOpts, validator, epoch, cursor, howMany)
}

// GetValidators is a free data retrieval call binding the contract method 0x72431991.
//
// Solidity: function getValidators(uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] owners, address[] operators, uint256[] stakes, bytes[] blsPublicKeys, bool[] candidates, uint256 newCursor)
func (_Stakemanager *StakemanagerCaller) GetValidators(opts *bind.CallOpts, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Owners        []common.Address
	Operators     []common.Address
	Stakes        []*big.Int
	BlsPublicKeys [][]byte
	Candidates    []bool
	NewCursor     *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidators", epoch, cursor, howMany)

	outstruct := new(struct {
		Owners        []common.Address
		Operators     []common.Address
		Stakes        []*big.Int
		BlsPublicKeys [][]byte
		Candidates    []bool
		NewCursor     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Operators = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.BlsPublicKeys = *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)
	outstruct.Candidates = *abi.ConvertType(out[4], new([]bool)).(*[]bool)
	outstruct.NewCursor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidators is a free data retrieval call binding the contract method 0x72431991.
//
// Solidity: function getValidators(uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] owners, address[] operators, uint256[] stakes, bytes[] blsPublicKeys, bool[] candidates, uint256 newCursor)
func (_Stakemanager *StakemanagerSession) GetValidators(epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Owners        []common.Address
	Operators     []common.Address
	Stakes        []*big.Int
	BlsPublicKeys [][]byte
	Candidates    []bool
	NewCursor     *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidators(&_Stakemanager.CallOpts, epoch, cursor, howMany)
}

// GetValidators is a free data retrieval call binding the contract method 0x72431991.
//
// Solidity: function getValidators(uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] owners, address[] operators, uint256[] stakes, bytes[] blsPublicKeys, bool[] candidates, uint256 newCursor)
func (_Stakemanager *StakemanagerCallerSession) GetValidators(epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Owners        []common.Address
	Operators     []common.Address
	Stakes        []*big.Int
	BlsPublicKeys [][]byte
	Candidates    []bool
	NewCursor     *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidators(&_Stakemanager.CallOpts, epoch, cursor, howMany)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Stakemanager *StakemanagerCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Stakemanager *StakemanagerSession) Initialized() (bool, error) {
	return _Stakemanager.Contract.Initialized(&_Stakemanager.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Stakemanager *StakemanagerCallerSession) Initialized() (bool, error) {
	return _Stakemanager.Contract.Initialized(&_Stakemanager.CallOpts)
}

// OperatorToOwner is a free data retrieval call binding the contract method 0x7b520aa8.
//
// Solidity: function operatorToOwner(address ) view returns(address)
func (_Stakemanager *StakemanagerCaller) OperatorToOwner(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "operatorToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorToOwner is a free data retrieval call binding the contract method 0x7b520aa8.
//
// Solidity: function operatorToOwner(address ) view returns(address)
func (_Stakemanager *StakemanagerSession) OperatorToOwner(arg0 common.Address) (common.Address, error) {
	return _Stakemanager.Contract.OperatorToOwner(&_Stakemanager.CallOpts, arg0)
}

// OperatorToOwner is a free data retrieval call binding the contract method 0x7b520aa8.
//
// Solidity: function operatorToOwner(address ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) OperatorToOwner(arg0 common.Address) (common.Address, error) {
	return _Stakemanager.Contract.OperatorToOwner(&_Stakemanager.CallOpts, arg0)
}

// StakeAmounts is a free data retrieval call binding the contract method 0x1c1b4f3a.
//
// Solidity: function stakeAmounts(uint256 ) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) StakeAmounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "stakeAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAmounts is a free data retrieval call binding the contract method 0x1c1b4f3a.
//
// Solidity: function stakeAmounts(uint256 ) view returns(uint256)
func (_Stakemanager *StakemanagerSession) StakeAmounts(arg0 *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.StakeAmounts(&_Stakemanager.CallOpts, arg0)
}

// StakeAmounts is a free data retrieval call binding the contract method 0x1c1b4f3a.
//
// Solidity: function stakeAmounts(uint256 ) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) StakeAmounts(arg0 *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.StakeAmounts(&_Stakemanager.CallOpts, arg0)
}

// StakeUpdates is a free data retrieval call binding the contract method 0x190b9257.
//
// Solidity: function stakeUpdates(uint256 ) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) StakeUpdates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "stakeUpdates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeUpdates is a free data retrieval call binding the contract method 0x190b9257.
//
// Solidity: function stakeUpdates(uint256 ) view returns(uint256)
func (_Stakemanager *StakemanagerSession) StakeUpdates(arg0 *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.StakeUpdates(&_Stakemanager.CallOpts, arg0)
}

// StakeUpdates is a free data retrieval call binding the contract method 0x190b9257.
//
// Solidity: function stakeUpdates(uint256 ) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) StakeUpdates(arg0 *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.StakeUpdates(&_Stakemanager.CallOpts, arg0)
}

// StakerSigners is a free data retrieval call binding the contract method 0xf65a5ed2.
//
// Solidity: function stakerSigners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCaller) StakerSigners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "stakerSigners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerSigners is a free data retrieval call binding the contract method 0xf65a5ed2.
//
// Solidity: function stakerSigners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerSession) StakerSigners(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.StakerSigners(&_Stakemanager.CallOpts, arg0)
}

// StakerSigners is a free data retrieval call binding the contract method 0xf65a5ed2.
//
// Solidity: function stakerSigners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) StakerSigners(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.StakerSigners(&_Stakemanager.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address signer)
func (_Stakemanager *StakemanagerCaller) Stakers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "stakers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address signer)
func (_Stakemanager *StakemanagerSession) Stakers(arg0 common.Address) (common.Address, error) {
	return _Stakemanager.Contract.Stakers(&_Stakemanager.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address signer)
func (_Stakemanager *StakemanagerCallerSession) Stakers(arg0 common.Address) (common.Address, error) {
	return _Stakemanager.Contract.Stakers(&_Stakemanager.CallOpts, arg0)
}

// ValidatorOwners is a free data retrieval call binding the contract method 0x5efc766e.
//
// Solidity: function validatorOwners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCaller) ValidatorOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "validatorOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorOwners is a free data retrieval call binding the contract method 0x5efc766e.
//
// Solidity: function validatorOwners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerSession) ValidatorOwners(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.ValidatorOwners(&_Stakemanager.CallOpts, arg0)
}

// ValidatorOwners is a free data retrieval call binding the contract method 0x5efc766e.
//
// Solidity: function validatorOwners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) ValidatorOwners(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.ValidatorOwners(&_Stakemanager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, uint256 lastClaimCommission, bytes blsPublicKey)
func (_Stakemanager *StakemanagerCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	LastClaimCommission *big.Int
	BlsPublicKey        []byte
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Owner               common.Address
		Operator            common.Address
		LastClaimCommission *big.Int
		BlsPublicKey        []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.LastClaimCommission = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BlsPublicKey = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, uint256 lastClaimCommission, bytes blsPublicKey)
func (_Stakemanager *StakemanagerSession) Validators(arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	LastClaimCommission *big.Int
	BlsPublicKey        []byte
}, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, uint256 lastClaimCommission, bytes blsPublicKey)
func (_Stakemanager *StakemanagerCallerSession) Validators(arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	LastClaimCommission *big.Int
	BlsPublicKey        []byte
}, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0x1903cf16.
//
// Solidity: function activateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerTransactor) ActivateValidator(opts *bind.TransactOpts, validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "activateValidator", validator, epochs)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0x1903cf16.
//
// Solidity: function activateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerSession) ActivateValidator(validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ActivateValidator(&_Stakemanager.TransactOpts, validator, epochs)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0x1903cf16.
//
// Solidity: function activateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) ActivateValidator(validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ActivateValidator(&_Stakemanager.TransactOpts, validator, epochs)
}

// AddRewardBalance is a paid mutator transaction binding the contract method 0x9043150b.
//
// Solidity: function addRewardBalance() payable returns()
func (_Stakemanager *StakemanagerTransactor) AddRewardBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "addRewardBalance")
}

// AddRewardBalance is a paid mutator transaction binding the contract method 0x9043150b.
//
// Solidity: function addRewardBalance() payable returns()
func (_Stakemanager *StakemanagerSession) AddRewardBalance() (*types.Transaction, error) {
	return _Stakemanager.Contract.AddRewardBalance(&_Stakemanager.TransactOpts)
}

// AddRewardBalance is a paid mutator transaction binding the contract method 0x9043150b.
//
// Solidity: function addRewardBalance() payable returns()
func (_Stakemanager *StakemanagerTransactorSession) AddRewardBalance() (*types.Transaction, error) {
	return _Stakemanager.Contract.AddRewardBalance(&_Stakemanager.TransactOpts)
}

// ClaimCommissions is a paid mutator transaction binding the contract method 0xcbc0fac6.
//
// Solidity: function claimCommissions(address , uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimCommissions(opts *bind.TransactOpts, arg0 common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimCommissions", arg0, epochs)
}

// ClaimCommissions is a paid mutator transaction binding the contract method 0xcbc0fac6.
//
// Solidity: function claimCommissions(address , uint256 epochs) returns()
func (_Stakemanager *StakemanagerSession) ClaimCommissions(arg0 common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimCommissions(&_Stakemanager.TransactOpts, arg0, epochs)
}

// ClaimCommissions is a paid mutator transaction binding the contract method 0xcbc0fac6.
//
// Solidity: function claimCommissions(address , uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimCommissions(arg0 common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimCommissions(&_Stakemanager.TransactOpts, arg0, epochs)
}

// ClaimLockedUnstake is a paid mutator transaction binding the contract method 0x5d94ccf6.
//
// Solidity: function claimLockedUnstake(uint256 lockedUnstake) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimLockedUnstake(opts *bind.TransactOpts, lockedUnstake *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimLockedUnstake", lockedUnstake)
}

// ClaimLockedUnstake is a paid mutator transaction binding the contract method 0x5d94ccf6.
//
// Solidity: function claimLockedUnstake(uint256 lockedUnstake) returns()
func (_Stakemanager *StakemanagerSession) ClaimLockedUnstake(lockedUnstake *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimLockedUnstake(&_Stakemanager.TransactOpts, lockedUnstake)
}

// ClaimLockedUnstake is a paid mutator transaction binding the contract method 0x5d94ccf6.
//
// Solidity: function claimLockedUnstake(uint256 lockedUnstake) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimLockedUnstake(lockedUnstake *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimLockedUnstake(&_Stakemanager.TransactOpts, lockedUnstake)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf3621e43.
//
// Solidity: function claimRewards(address , address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimRewards(opts *bind.TransactOpts, arg0 common.Address, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimRewards", arg0, validator, epochs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf3621e43.
//
// Solidity: function claimRewards(address , address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerSession) ClaimRewards(arg0 common.Address, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimRewards(&_Stakemanager.TransactOpts, arg0, validator, epochs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf3621e43.
//
// Solidity: function claimRewards(address , address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimRewards(arg0 common.Address, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimRewards(&_Stakemanager.TransactOpts, arg0, validator, epochs)
}

// ClaimUnstakes is a paid mutator transaction binding the contract method 0xf8d6b1ab.
//
// Solidity: function claimUnstakes(address ) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimUnstakes(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimUnstakes", arg0)
}

// ClaimUnstakes is a paid mutator transaction binding the contract method 0xf8d6b1ab.
//
// Solidity: function claimUnstakes(address ) returns()
func (_Stakemanager *StakemanagerSession) ClaimUnstakes(arg0 common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimUnstakes(&_Stakemanager.TransactOpts, arg0)
}

// ClaimUnstakes is a paid mutator transaction binding the contract method 0xf8d6b1ab.
//
// Solidity: function claimUnstakes(address ) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimUnstakes(arg0 common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimUnstakes(&_Stakemanager.TransactOpts, arg0)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x428e8562.
//
// Solidity: function deactivateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerTransactor) DeactivateValidator(opts *bind.TransactOpts, validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "deactivateValidator", validator, epochs)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x428e8562.
//
// Solidity: function deactivateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerSession) DeactivateValidator(validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.DeactivateValidator(&_Stakemanager.TransactOpts, validator, epochs)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x428e8562.
//
// Solidity: function deactivateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) DeactivateValidator(validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.DeactivateValidator(&_Stakemanager.TransactOpts, validator, epochs)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _environment, address _allowlist) returns()
func (_Stakemanager *StakemanagerTransactor) Initialize(opts *bind.TransactOpts, _environment common.Address, _allowlist common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "initialize", _environment, _allowlist)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _environment, address _allowlist) returns()
func (_Stakemanager *StakemanagerSession) Initialize(_environment common.Address, _allowlist common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Initialize(&_Stakemanager.TransactOpts, _environment, _allowlist)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _environment, address _allowlist) returns()
func (_Stakemanager *StakemanagerTransactorSession) Initialize(_environment common.Address, _allowlist common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Initialize(&_Stakemanager.TransactOpts, _environment, _allowlist)
}

// JoinValidator is a paid mutator transaction binding the contract method 0x6b2b3369.
//
// Solidity: function joinValidator(address operator) returns()
func (_Stakemanager *StakemanagerTransactor) JoinValidator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "joinValidator", operator)
}

// JoinValidator is a paid mutator transaction binding the contract method 0x6b2b3369.
//
// Solidity: function joinValidator(address operator) returns()
func (_Stakemanager *StakemanagerSession) JoinValidator(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.JoinValidator(&_Stakemanager.TransactOpts, operator)
}

// JoinValidator is a paid mutator transaction binding the contract method 0x6b2b3369.
//
// Solidity: function joinValidator(address operator) returns()
func (_Stakemanager *StakemanagerTransactorSession) JoinValidator(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.JoinValidator(&_Stakemanager.TransactOpts, operator)
}

// RestakeCommissions is a paid mutator transaction binding the contract method 0x0ddda63c.
//
// Solidity: function restakeCommissions(uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactor) RestakeCommissions(opts *bind.TransactOpts, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "restakeCommissions", epochs)
}

// RestakeCommissions is a paid mutator transaction binding the contract method 0x0ddda63c.
//
// Solidity: function restakeCommissions(uint256 epochs) returns()
func (_Stakemanager *StakemanagerSession) RestakeCommissions(epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.RestakeCommissions(&_Stakemanager.TransactOpts, epochs)
}

// RestakeCommissions is a paid mutator transaction binding the contract method 0x0ddda63c.
//
// Solidity: function restakeCommissions(uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) RestakeCommissions(epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.RestakeCommissions(&_Stakemanager.TransactOpts, epochs)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0xcf5c13db.
//
// Solidity: function restakeRewards(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactor) RestakeRewards(opts *bind.TransactOpts, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "restakeRewards", validator, epochs)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0xcf5c13db.
//
// Solidity: function restakeRewards(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerSession) RestakeRewards(validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.RestakeRewards(&_Stakemanager.TransactOpts, validator, epochs)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0xcf5c13db.
//
// Solidity: function restakeRewards(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) RestakeRewards(validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.RestakeRewards(&_Stakemanager.TransactOpts, validator, epochs)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address operator, uint256 blocks) returns()
func (_Stakemanager *StakemanagerTransactor) Slash(opts *bind.TransactOpts, operator common.Address, blocks *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "slash", operator, blocks)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address operator, uint256 blocks) returns()
func (_Stakemanager *StakemanagerSession) Slash(operator common.Address, blocks *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Slash(&_Stakemanager.TransactOpts, operator, blocks)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address operator, uint256 blocks) returns()
func (_Stakemanager *StakemanagerTransactorSession) Slash(operator common.Address, blocks *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Slash(&_Stakemanager.TransactOpts, operator, blocks)
}

// Stake is a paid mutator transaction binding the contract method 0x7befa74f.
//
// Solidity: function stake(address validator, uint8 token, uint256 amount) payable returns()
func (_Stakemanager *StakemanagerTransactor) Stake(opts *bind.TransactOpts, validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "stake", validator, token, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7befa74f.
//
// Solidity: function stake(address validator, uint8 token, uint256 amount) payable returns()
func (_Stakemanager *StakemanagerSession) Stake(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Stake(&_Stakemanager.TransactOpts, validator, token, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7befa74f.
//
// Solidity: function stake(address validator, uint8 token, uint256 amount) payable returns()
func (_Stakemanager *StakemanagerTransactorSession) Stake(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Stake(&_Stakemanager.TransactOpts, validator, token, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xe1aca341.
//
// Solidity: function unstake(address , uint8 , uint256 ) returns()
func (_Stakemanager *StakemanagerTransactor) Unstake(opts *bind.TransactOpts, arg0 common.Address, arg1 uint8, arg2 *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "unstake", arg0, arg1, arg2)
}

// Unstake is a paid mutator transaction binding the contract method 0xe1aca341.
//
// Solidity: function unstake(address , uint8 , uint256 ) returns()
func (_Stakemanager *StakemanagerSession) Unstake(arg0 common.Address, arg1 uint8, arg2 *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, arg0, arg1, arg2)
}

// Unstake is a paid mutator transaction binding the contract method 0xe1aca341.
//
// Solidity: function unstake(address , uint8 , uint256 ) returns()
func (_Stakemanager *StakemanagerTransactorSession) Unstake(arg0 common.Address, arg1 uint8, arg2 *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, arg0, arg1, arg2)
}

// UnstakeV2 is a paid mutator transaction binding the contract method 0xff3d3f60.
//
// Solidity: function unstakeV2(address validator, uint8 token, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) UnstakeV2(opts *bind.TransactOpts, validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "unstakeV2", validator, token, amount)
}

// UnstakeV2 is a paid mutator transaction binding the contract method 0xff3d3f60.
//
// Solidity: function unstakeV2(address validator, uint8 token, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) UnstakeV2(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UnstakeV2(&_Stakemanager.TransactOpts, validator, token, amount)
}

// UnstakeV2 is a paid mutator transaction binding the contract method 0xff3d3f60.
//
// Solidity: function unstakeV2(address validator, uint8 token, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) UnstakeV2(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UnstakeV2(&_Stakemanager.TransactOpts, validator, token, amount)
}

// UpdateBLSPublicKey is a paid mutator transaction binding the contract method 0x00c8ae89.
//
// Solidity: function updateBLSPublicKey(bytes blsPublicKey) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateBLSPublicKey(opts *bind.TransactOpts, blsPublicKey []byte) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateBLSPublicKey", blsPublicKey)
}

// UpdateBLSPublicKey is a paid mutator transaction binding the contract method 0x00c8ae89.
//
// Solidity: function updateBLSPublicKey(bytes blsPublicKey) returns()
func (_Stakemanager *StakemanagerSession) UpdateBLSPublicKey(blsPublicKey []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateBLSPublicKey(&_Stakemanager.TransactOpts, blsPublicKey)
}

// UpdateBLSPublicKey is a paid mutator transaction binding the contract method 0x00c8ae89.
//
// Solidity: function updateBLSPublicKey(bytes blsPublicKey) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateBLSPublicKey(blsPublicKey []byte) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateBLSPublicKey(&_Stakemanager.TransactOpts, blsPublicKey)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xac7475ed.
//
// Solidity: function updateOperator(address operator) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateOperator", operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xac7475ed.
//
// Solidity: function updateOperator(address operator) returns()
func (_Stakemanager *StakemanagerSession) UpdateOperator(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateOperator(&_Stakemanager.TransactOpts, operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xac7475ed.
//
// Solidity: function updateOperator(address operator) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateOperator(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateOperator(&_Stakemanager.TransactOpts, operator)
}

// StakemanagerAddedRewardBalanceIterator is returned from FilterAddedRewardBalance and is used to iterate over the raw logs and unpacked data for AddedRewardBalance events raised by the Stakemanager contract.
type StakemanagerAddedRewardBalanceIterator struct {
	Event *StakemanagerAddedRewardBalance // Event containing the contract specifics and raw log

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
func (it *StakemanagerAddedRewardBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerAddedRewardBalance)
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
		it.Event = new(StakemanagerAddedRewardBalance)
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
func (it *StakemanagerAddedRewardBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerAddedRewardBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerAddedRewardBalance represents a AddedRewardBalance event raised by the Stakemanager contract.
type StakemanagerAddedRewardBalance struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddedRewardBalance is a free log retrieval operation binding the contract event 0x1de49774d094a85fc1bbbd16e8d09a865fb848218f41e2da4369f528c42ee42e.
//
// Solidity: event AddedRewardBalance(uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterAddedRewardBalance(opts *bind.FilterOpts) (*StakemanagerAddedRewardBalanceIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "AddedRewardBalance")
	if err != nil {
		return nil, err
	}
	return &StakemanagerAddedRewardBalanceIterator{contract: _Stakemanager.contract, event: "AddedRewardBalance", logs: logs, sub: sub}, nil
}

// WatchAddedRewardBalance is a free log subscription operation binding the contract event 0x1de49774d094a85fc1bbbd16e8d09a865fb848218f41e2da4369f528c42ee42e.
//
// Solidity: event AddedRewardBalance(uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchAddedRewardBalance(opts *bind.WatchOpts, sink chan<- *StakemanagerAddedRewardBalance) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "AddedRewardBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerAddedRewardBalance)
				if err := _Stakemanager.contract.UnpackLog(event, "AddedRewardBalance", log); err != nil {
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

// ParseAddedRewardBalance is a log parse operation binding the contract event 0x1de49774d094a85fc1bbbd16e8d09a865fb848218f41e2da4369f528c42ee42e.
//
// Solidity: event AddedRewardBalance(uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseAddedRewardBalance(log types.Log) (*StakemanagerAddedRewardBalance, error) {
	event := new(StakemanagerAddedRewardBalance)
	if err := _Stakemanager.contract.UnpackLog(event, "AddedRewardBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerBLSPublicKeyUpdatedIterator is returned from FilterBLSPublicKeyUpdated and is used to iterate over the raw logs and unpacked data for BLSPublicKeyUpdated events raised by the Stakemanager contract.
type StakemanagerBLSPublicKeyUpdatedIterator struct {
	Event *StakemanagerBLSPublicKeyUpdated // Event containing the contract specifics and raw log

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
func (it *StakemanagerBLSPublicKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerBLSPublicKeyUpdated)
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
		it.Event = new(StakemanagerBLSPublicKeyUpdated)
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
func (it *StakemanagerBLSPublicKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerBLSPublicKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerBLSPublicKeyUpdated represents a BLSPublicKeyUpdated event raised by the Stakemanager contract.
type StakemanagerBLSPublicKeyUpdated struct {
	Validator       common.Address
	OldBLSPublicKey []byte
	NewBLSPublicKey []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBLSPublicKeyUpdated is a free log retrieval operation binding the contract event 0x7ea7e12060119574f657de08c5ef0970a24d7734612fb00c418ad40c7d4a84fd.
//
// Solidity: event BLSPublicKeyUpdated(address indexed validator, bytes oldBLSPublicKey, bytes newBLSPublicKey)
func (_Stakemanager *StakemanagerFilterer) FilterBLSPublicKeyUpdated(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerBLSPublicKeyUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "BLSPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerBLSPublicKeyUpdatedIterator{contract: _Stakemanager.contract, event: "BLSPublicKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchBLSPublicKeyUpdated is a free log subscription operation binding the contract event 0x7ea7e12060119574f657de08c5ef0970a24d7734612fb00c418ad40c7d4a84fd.
//
// Solidity: event BLSPublicKeyUpdated(address indexed validator, bytes oldBLSPublicKey, bytes newBLSPublicKey)
func (_Stakemanager *StakemanagerFilterer) WatchBLSPublicKeyUpdated(opts *bind.WatchOpts, sink chan<- *StakemanagerBLSPublicKeyUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "BLSPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerBLSPublicKeyUpdated)
				if err := _Stakemanager.contract.UnpackLog(event, "BLSPublicKeyUpdated", log); err != nil {
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

// ParseBLSPublicKeyUpdated is a log parse operation binding the contract event 0x7ea7e12060119574f657de08c5ef0970a24d7734612fb00c418ad40c7d4a84fd.
//
// Solidity: event BLSPublicKeyUpdated(address indexed validator, bytes oldBLSPublicKey, bytes newBLSPublicKey)
func (_Stakemanager *StakemanagerFilterer) ParseBLSPublicKeyUpdated(log types.Log) (*StakemanagerBLSPublicKeyUpdated, error) {
	event := new(StakemanagerBLSPublicKeyUpdated)
	if err := _Stakemanager.contract.UnpackLog(event, "BLSPublicKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerClaimedCommissionsIterator is returned from FilterClaimedCommissions and is used to iterate over the raw logs and unpacked data for ClaimedCommissions events raised by the Stakemanager contract.
type StakemanagerClaimedCommissionsIterator struct {
	Event *StakemanagerClaimedCommissions // Event containing the contract specifics and raw log

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
func (it *StakemanagerClaimedCommissionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerClaimedCommissions)
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
		it.Event = new(StakemanagerClaimedCommissions)
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
func (it *StakemanagerClaimedCommissionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerClaimedCommissionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerClaimedCommissions represents a ClaimedCommissions event raised by the Stakemanager contract.
type StakemanagerClaimedCommissions struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimedCommissions is a free log retrieval operation binding the contract event 0x882d5671e5b36af50883197c33d48ba56ce337589958e871ba82fb0a54adf3e8.
//
// Solidity: event ClaimedCommissions(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterClaimedCommissions(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerClaimedCommissionsIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ClaimedCommissions", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerClaimedCommissionsIterator{contract: _Stakemanager.contract, event: "ClaimedCommissions", logs: logs, sub: sub}, nil
}

// WatchClaimedCommissions is a free log subscription operation binding the contract event 0x882d5671e5b36af50883197c33d48ba56ce337589958e871ba82fb0a54adf3e8.
//
// Solidity: event ClaimedCommissions(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchClaimedCommissions(opts *bind.WatchOpts, sink chan<- *StakemanagerClaimedCommissions, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ClaimedCommissions", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerClaimedCommissions)
				if err := _Stakemanager.contract.UnpackLog(event, "ClaimedCommissions", log); err != nil {
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

// ParseClaimedCommissions is a log parse operation binding the contract event 0x882d5671e5b36af50883197c33d48ba56ce337589958e871ba82fb0a54adf3e8.
//
// Solidity: event ClaimedCommissions(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseClaimedCommissions(log types.Log) (*StakemanagerClaimedCommissions, error) {
	event := new(StakemanagerClaimedCommissions)
	if err := _Stakemanager.contract.UnpackLog(event, "ClaimedCommissions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerClaimedLockedUnstakeIterator is returned from FilterClaimedLockedUnstake and is used to iterate over the raw logs and unpacked data for ClaimedLockedUnstake events raised by the Stakemanager contract.
type StakemanagerClaimedLockedUnstakeIterator struct {
	Event *StakemanagerClaimedLockedUnstake // Event containing the contract specifics and raw log

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
func (it *StakemanagerClaimedLockedUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerClaimedLockedUnstake)
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
		it.Event = new(StakemanagerClaimedLockedUnstake)
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
func (it *StakemanagerClaimedLockedUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerClaimedLockedUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerClaimedLockedUnstake represents a ClaimedLockedUnstake event raised by the Stakemanager contract.
type StakemanagerClaimedLockedUnstake struct {
	Staker        common.Address
	LockedUnstake *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClaimedLockedUnstake is a free log retrieval operation binding the contract event 0xbf5f92dc2b945251eadf78c2ca629ae64053d979bfbc43a7b17a463707906bf9.
//
// Solidity: event ClaimedLockedUnstake(address indexed staker, uint256 lockedUnstake)
func (_Stakemanager *StakemanagerFilterer) FilterClaimedLockedUnstake(opts *bind.FilterOpts, staker []common.Address) (*StakemanagerClaimedLockedUnstakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ClaimedLockedUnstake", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerClaimedLockedUnstakeIterator{contract: _Stakemanager.contract, event: "ClaimedLockedUnstake", logs: logs, sub: sub}, nil
}

// WatchClaimedLockedUnstake is a free log subscription operation binding the contract event 0xbf5f92dc2b945251eadf78c2ca629ae64053d979bfbc43a7b17a463707906bf9.
//
// Solidity: event ClaimedLockedUnstake(address indexed staker, uint256 lockedUnstake)
func (_Stakemanager *StakemanagerFilterer) WatchClaimedLockedUnstake(opts *bind.WatchOpts, sink chan<- *StakemanagerClaimedLockedUnstake, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ClaimedLockedUnstake", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerClaimedLockedUnstake)
				if err := _Stakemanager.contract.UnpackLog(event, "ClaimedLockedUnstake", log); err != nil {
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

// ParseClaimedLockedUnstake is a log parse operation binding the contract event 0xbf5f92dc2b945251eadf78c2ca629ae64053d979bfbc43a7b17a463707906bf9.
//
// Solidity: event ClaimedLockedUnstake(address indexed staker, uint256 lockedUnstake)
func (_Stakemanager *StakemanagerFilterer) ParseClaimedLockedUnstake(log types.Log) (*StakemanagerClaimedLockedUnstake, error) {
	event := new(StakemanagerClaimedLockedUnstake)
	if err := _Stakemanager.contract.UnpackLog(event, "ClaimedLockedUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerClaimedRewardsIterator is returned from FilterClaimedRewards and is used to iterate over the raw logs and unpacked data for ClaimedRewards events raised by the Stakemanager contract.
type StakemanagerClaimedRewardsIterator struct {
	Event *StakemanagerClaimedRewards // Event containing the contract specifics and raw log

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
func (it *StakemanagerClaimedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerClaimedRewards)
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
		it.Event = new(StakemanagerClaimedRewards)
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
func (it *StakemanagerClaimedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerClaimedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerClaimedRewards represents a ClaimedRewards event raised by the Stakemanager contract.
type StakemanagerClaimedRewards struct {
	Staker    common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimedRewards is a free log retrieval operation binding the contract event 0x2ef606d064225d24c1514dc94907c134faee1237445c2f63f410cce0852b2054.
//
// Solidity: event ClaimedRewards(address indexed staker, address validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterClaimedRewards(opts *bind.FilterOpts, staker []common.Address) (*StakemanagerClaimedRewardsIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ClaimedRewards", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerClaimedRewardsIterator{contract: _Stakemanager.contract, event: "ClaimedRewards", logs: logs, sub: sub}, nil
}

// WatchClaimedRewards is a free log subscription operation binding the contract event 0x2ef606d064225d24c1514dc94907c134faee1237445c2f63f410cce0852b2054.
//
// Solidity: event ClaimedRewards(address indexed staker, address validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchClaimedRewards(opts *bind.WatchOpts, sink chan<- *StakemanagerClaimedRewards, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ClaimedRewards", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerClaimedRewards)
				if err := _Stakemanager.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
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

// ParseClaimedRewards is a log parse operation binding the contract event 0x2ef606d064225d24c1514dc94907c134faee1237445c2f63f410cce0852b2054.
//
// Solidity: event ClaimedRewards(address indexed staker, address validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseClaimedRewards(log types.Log) (*StakemanagerClaimedRewards, error) {
	event := new(StakemanagerClaimedRewards)
	if err := _Stakemanager.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerOperatorUpdatedIterator is returned from FilterOperatorUpdated and is used to iterate over the raw logs and unpacked data for OperatorUpdated events raised by the Stakemanager contract.
type StakemanagerOperatorUpdatedIterator struct {
	Event *StakemanagerOperatorUpdated // Event containing the contract specifics and raw log

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
func (it *StakemanagerOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerOperatorUpdated)
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
		it.Event = new(StakemanagerOperatorUpdated)
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
func (it *StakemanagerOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerOperatorUpdated represents a OperatorUpdated event raised by the Stakemanager contract.
type StakemanagerOperatorUpdated struct {
	Validator   common.Address
	OldOperator common.Address
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorUpdated is a free log retrieval operation binding the contract event 0x758820d0b14a01c1fa60b8d2bbef25ed1b6a5af4802e5dec3f08679255ba8bf3.
//
// Solidity: event OperatorUpdated(address indexed validator, address oldOperator, address newOperator)
func (_Stakemanager *StakemanagerFilterer) FilterOperatorUpdated(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerOperatorUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "OperatorUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerOperatorUpdatedIterator{contract: _Stakemanager.contract, event: "OperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorUpdated is a free log subscription operation binding the contract event 0x758820d0b14a01c1fa60b8d2bbef25ed1b6a5af4802e5dec3f08679255ba8bf3.
//
// Solidity: event OperatorUpdated(address indexed validator, address oldOperator, address newOperator)
func (_Stakemanager *StakemanagerFilterer) WatchOperatorUpdated(opts *bind.WatchOpts, sink chan<- *StakemanagerOperatorUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "OperatorUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerOperatorUpdated)
				if err := _Stakemanager.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
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

// ParseOperatorUpdated is a log parse operation binding the contract event 0x758820d0b14a01c1fa60b8d2bbef25ed1b6a5af4802e5dec3f08679255ba8bf3.
//
// Solidity: event OperatorUpdated(address indexed validator, address oldOperator, address newOperator)
func (_Stakemanager *StakemanagerFilterer) ParseOperatorUpdated(log types.Log) (*StakemanagerOperatorUpdated, error) {
	event := new(StakemanagerOperatorUpdated)
	if err := _Stakemanager.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerReStakedIterator is returned from FilterReStaked and is used to iterate over the raw logs and unpacked data for ReStaked events raised by the Stakemanager contract.
type StakemanagerReStakedIterator struct {
	Event *StakemanagerReStaked // Event containing the contract specifics and raw log

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
func (it *StakemanagerReStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerReStaked)
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
		it.Event = new(StakemanagerReStaked)
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
func (it *StakemanagerReStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerReStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerReStaked represents a ReStaked event raised by the Stakemanager contract.
type StakemanagerReStaked struct {
	Staker    common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReStaked is a free log retrieval operation binding the contract event 0xddd8e3ffe5c76cd1a7cca4b98662cb0a4e3da53ee24a873da632d28ba1043836.
//
// Solidity: event ReStaked(address indexed staker, address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterReStaked(opts *bind.FilterOpts, staker []common.Address, validator []common.Address) (*StakemanagerReStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ReStaked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerReStakedIterator{contract: _Stakemanager.contract, event: "ReStaked", logs: logs, sub: sub}, nil
}

// WatchReStaked is a free log subscription operation binding the contract event 0xddd8e3ffe5c76cd1a7cca4b98662cb0a4e3da53ee24a873da632d28ba1043836.
//
// Solidity: event ReStaked(address indexed staker, address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchReStaked(opts *bind.WatchOpts, sink chan<- *StakemanagerReStaked, staker []common.Address, validator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ReStaked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerReStaked)
				if err := _Stakemanager.contract.UnpackLog(event, "ReStaked", log); err != nil {
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

// ParseReStaked is a log parse operation binding the contract event 0xddd8e3ffe5c76cd1a7cca4b98662cb0a4e3da53ee24a873da632d28ba1043836.
//
// Solidity: event ReStaked(address indexed staker, address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseReStaked(log types.Log) (*StakemanagerReStaked, error) {
	event := new(StakemanagerReStaked)
	if err := _Stakemanager.contract.UnpackLog(event, "ReStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Stakemanager contract.
type StakemanagerStakedIterator struct {
	Event *StakemanagerStaked // Event containing the contract specifics and raw log

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
func (it *StakemanagerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerStaked)
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
		it.Event = new(StakemanagerStaked)
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
func (it *StakemanagerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerStaked represents a Staked event raised by the Stakemanager contract.
type StakemanagerStaked struct {
	Staker    common.Address
	Validator common.Address
	Token     uint8
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x8fc656e319452025372383dc27d933046d412b8253de50a10739eeaa59862be6.
//
// Solidity: event Staked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address, validator []common.Address) (*StakemanagerStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Staked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerStakedIterator{contract: _Stakemanager.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x8fc656e319452025372383dc27d933046d412b8253de50a10739eeaa59862be6.
//
// Solidity: event Staked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakemanagerStaked, staker []common.Address, validator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Staked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerStaked)
				if err := _Stakemanager.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x8fc656e319452025372383dc27d933046d412b8253de50a10739eeaa59862be6.
//
// Solidity: event Staked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseStaked(log types.Log) (*StakemanagerStaked, error) {
	event := new(StakemanagerStaked)
	if err := _Stakemanager.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Stakemanager contract.
type StakemanagerUnstakedIterator struct {
	Event *StakemanagerUnstaked // Event containing the contract specifics and raw log

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
func (it *StakemanagerUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerUnstaked)
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
		it.Event = new(StakemanagerUnstaked)
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
func (it *StakemanagerUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerUnstaked represents a Unstaked event raised by the Stakemanager contract.
type StakemanagerUnstaked struct {
	Staker    common.Address
	Validator common.Address
	Token     uint8
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xf2812c3df2511a467cbe777b1ee98b1ddb9918bb0f09568a269d2fb58233cb52.
//
// Solidity: event Unstaked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterUnstaked(opts *bind.FilterOpts, staker []common.Address, validator []common.Address) (*StakemanagerUnstakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Unstaked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerUnstakedIterator{contract: _Stakemanager.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xf2812c3df2511a467cbe777b1ee98b1ddb9918bb0f09568a269d2fb58233cb52.
//
// Solidity: event Unstaked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakemanagerUnstaked, staker []common.Address, validator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Unstaked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerUnstaked)
				if err := _Stakemanager.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0xf2812c3df2511a467cbe777b1ee98b1ddb9918bb0f09568a269d2fb58233cb52.
//
// Solidity: event Unstaked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseUnstaked(log types.Log) (*StakemanagerUnstaked, error) {
	event := new(StakemanagerUnstaked)
	if err := _Stakemanager.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerUnstakedV2Iterator is returned from FilterUnstakedV2 and is used to iterate over the raw logs and unpacked data for UnstakedV2 events raised by the Stakemanager contract.
type StakemanagerUnstakedV2Iterator struct {
	Event *StakemanagerUnstakedV2 // Event containing the contract specifics and raw log

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
func (it *StakemanagerUnstakedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerUnstakedV2)
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
		it.Event = new(StakemanagerUnstakedV2)
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
func (it *StakemanagerUnstakedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerUnstakedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerUnstakedV2 represents a UnstakedV2 event raised by the Stakemanager contract.
type StakemanagerUnstakedV2 struct {
	Staker        common.Address
	Validator     common.Address
	LockedUnstake *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakedV2 is a free log retrieval operation binding the contract event 0xb649014faa7a0e23357e091fb8a67a128c33dc9480f846f7e41cb3a6c9d80610.
//
// Solidity: event UnstakedV2(address indexed staker, address indexed validator, uint256 lockedUnstake)
func (_Stakemanager *StakemanagerFilterer) FilterUnstakedV2(opts *bind.FilterOpts, staker []common.Address, validator []common.Address) (*StakemanagerUnstakedV2Iterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "UnstakedV2", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerUnstakedV2Iterator{contract: _Stakemanager.contract, event: "UnstakedV2", logs: logs, sub: sub}, nil
}

// WatchUnstakedV2 is a free log subscription operation binding the contract event 0xb649014faa7a0e23357e091fb8a67a128c33dc9480f846f7e41cb3a6c9d80610.
//
// Solidity: event UnstakedV2(address indexed staker, address indexed validator, uint256 lockedUnstake)
func (_Stakemanager *StakemanagerFilterer) WatchUnstakedV2(opts *bind.WatchOpts, sink chan<- *StakemanagerUnstakedV2, staker []common.Address, validator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "UnstakedV2", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerUnstakedV2)
				if err := _Stakemanager.contract.UnpackLog(event, "UnstakedV2", log); err != nil {
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

// ParseUnstakedV2 is a log parse operation binding the contract event 0xb649014faa7a0e23357e091fb8a67a128c33dc9480f846f7e41cb3a6c9d80610.
//
// Solidity: event UnstakedV2(address indexed staker, address indexed validator, uint256 lockedUnstake)
func (_Stakemanager *StakemanagerFilterer) ParseUnstakedV2(log types.Log) (*StakemanagerUnstakedV2, error) {
	event := new(StakemanagerUnstakedV2)
	if err := _Stakemanager.contract.UnpackLog(event, "UnstakedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorActivatedIterator is returned from FilterValidatorActivated and is used to iterate over the raw logs and unpacked data for ValidatorActivated events raised by the Stakemanager contract.
type StakemanagerValidatorActivatedIterator struct {
	Event *StakemanagerValidatorActivated // Event containing the contract specifics and raw log

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
func (it *StakemanagerValidatorActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorActivated)
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
		it.Event = new(StakemanagerValidatorActivated)
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
func (it *StakemanagerValidatorActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorActivated represents a ValidatorActivated event raised by the Stakemanager contract.
type StakemanagerValidatorActivated struct {
	Validator common.Address
	Epochs    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorActivated is a free log retrieval operation binding the contract event 0xc11dfc9c24621433bb10587dc4bbae26a33a4aff53914e0d4c9fddf224a8cb7d.
//
// Solidity: event ValidatorActivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorActivated(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorActivatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorActivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorActivatedIterator{contract: _Stakemanager.contract, event: "ValidatorActivated", logs: logs, sub: sub}, nil
}

// WatchValidatorActivated is a free log subscription operation binding the contract event 0xc11dfc9c24621433bb10587dc4bbae26a33a4aff53914e0d4c9fddf224a8cb7d.
//
// Solidity: event ValidatorActivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorActivated(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorActivated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorActivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorActivated)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
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

// ParseValidatorActivated is a log parse operation binding the contract event 0xc11dfc9c24621433bb10587dc4bbae26a33a4aff53914e0d4c9fddf224a8cb7d.
//
// Solidity: event ValidatorActivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorActivated(log types.Log) (*StakemanagerValidatorActivated, error) {
	event := new(StakemanagerValidatorActivated)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorDeactivatedIterator is returned from FilterValidatorDeactivated and is used to iterate over the raw logs and unpacked data for ValidatorDeactivated events raised by the Stakemanager contract.
type StakemanagerValidatorDeactivatedIterator struct {
	Event *StakemanagerValidatorDeactivated // Event containing the contract specifics and raw log

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
func (it *StakemanagerValidatorDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorDeactivated)
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
		it.Event = new(StakemanagerValidatorDeactivated)
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
func (it *StakemanagerValidatorDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorDeactivated represents a ValidatorDeactivated event raised by the Stakemanager contract.
type StakemanagerValidatorDeactivated struct {
	Validator common.Address
	Epochs    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeactivated is a free log retrieval operation binding the contract event 0x0ad9bf1b8c026a174a2f30954417002a6ea00c9b08c1b8846c7951c687be8095.
//
// Solidity: event ValidatorDeactivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorDeactivated(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorDeactivatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorDeactivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorDeactivatedIterator{contract: _Stakemanager.contract, event: "ValidatorDeactivated", logs: logs, sub: sub}, nil
}

// WatchValidatorDeactivated is a free log subscription operation binding the contract event 0x0ad9bf1b8c026a174a2f30954417002a6ea00c9b08c1b8846c7951c687be8095.
//
// Solidity: event ValidatorDeactivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorDeactivated(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorDeactivated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorDeactivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorDeactivated)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorDeactivated", log); err != nil {
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

// ParseValidatorDeactivated is a log parse operation binding the contract event 0x0ad9bf1b8c026a174a2f30954417002a6ea00c9b08c1b8846c7951c687be8095.
//
// Solidity: event ValidatorDeactivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorDeactivated(log types.Log) (*StakemanagerValidatorDeactivated, error) {
	event := new(StakemanagerValidatorDeactivated)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the Stakemanager contract.
type StakemanagerValidatorJailedIterator struct {
	Event *StakemanagerValidatorJailed // Event containing the contract specifics and raw log

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
func (it *StakemanagerValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorJailed)
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
		it.Event = new(StakemanagerValidatorJailed)
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
func (it *StakemanagerValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorJailed represents a ValidatorJailed event raised by the Stakemanager contract.
type StakemanagerValidatorJailed struct {
	Validator common.Address
	Until     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0xeb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802.
//
// Solidity: event ValidatorJailed(address indexed validator, uint256 until)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorJailed(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorJailedIterator{contract: _Stakemanager.contract, event: "ValidatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0xeb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802.
//
// Solidity: event ValidatorJailed(address indexed validator, uint256 until)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorJailed)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
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

// ParseValidatorJailed is a log parse operation binding the contract event 0xeb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802.
//
// Solidity: event ValidatorJailed(address indexed validator, uint256 until)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorJailed(log types.Log) (*StakemanagerValidatorJailed, error) {
	event := new(StakemanagerValidatorJailed)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorJoinedIterator is returned from FilterValidatorJoined and is used to iterate over the raw logs and unpacked data for ValidatorJoined events raised by the Stakemanager contract.
type StakemanagerValidatorJoinedIterator struct {
	Event *StakemanagerValidatorJoined // Event containing the contract specifics and raw log

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
func (it *StakemanagerValidatorJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorJoined)
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
		it.Event = new(StakemanagerValidatorJoined)
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
func (it *StakemanagerValidatorJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorJoined represents a ValidatorJoined event raised by the Stakemanager contract.
type StakemanagerValidatorJoined struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJoined is a free log retrieval operation binding the contract event 0xd5828184f48f65962d10eac907318df85953d4e3542a0f09b5932ee3fe398bdd.
//
// Solidity: event ValidatorJoined(address validator)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorJoined(opts *bind.FilterOpts) (*StakemanagerValidatorJoinedIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorJoined")
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorJoinedIterator{contract: _Stakemanager.contract, event: "ValidatorJoined", logs: logs, sub: sub}, nil
}

// WatchValidatorJoined is a free log subscription operation binding the contract event 0xd5828184f48f65962d10eac907318df85953d4e3542a0f09b5932ee3fe398bdd.
//
// Solidity: event ValidatorJoined(address validator)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorJoined(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorJoined) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorJoined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorJoined)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorJoined", log); err != nil {
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

// ParseValidatorJoined is a log parse operation binding the contract event 0xd5828184f48f65962d10eac907318df85953d4e3542a0f09b5932ee3fe398bdd.
//
// Solidity: event ValidatorJoined(address validator)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorJoined(log types.Log) (*StakemanagerValidatorJoined, error) {
	event := new(StakemanagerValidatorJoined)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorJoined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the Stakemanager contract.
type StakemanagerValidatorSlashedIterator struct {
	Event *StakemanagerValidatorSlashed // Event containing the contract specifics and raw log

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
func (it *StakemanagerValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorSlashed)
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
		it.Event = new(StakemanagerValidatorSlashed)
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
func (it *StakemanagerValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorSlashed represents a ValidatorSlashed event raised by the Stakemanager contract.
type StakemanagerValidatorSlashed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0x1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda.
//
// Solidity: event ValidatorSlashed(address indexed validator)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorSlashedIterator{contract: _Stakemanager.contract, event: "ValidatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0x1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda.
//
// Solidity: event ValidatorSlashed(address indexed validator)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorSlashed)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
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

// ParseValidatorSlashed is a log parse operation binding the contract event 0x1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda.
//
// Solidity: event ValidatorSlashed(address indexed validator)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorSlashed(log types.Log) (*StakemanagerValidatorSlashed, error) {
	event := new(StakemanagerValidatorSlashed)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
