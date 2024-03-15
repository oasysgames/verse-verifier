// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2ooverifier

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

// OasysL2OutputOracleVerifierOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type OasysL2OutputOracleVerifierOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// OasysL2OutputOracleVerifierSetting is an auto generated low-level Go binding around an user-defined struct.
type OasysL2OutputOracleVerifierSetting struct {
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
}

// L2ooverifierMetaData contains all meta data concerning the L2ooverifier contract.
var L2ooverifierMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysL2OutputOracleVerifier.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assertLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysL2OutputOracleVerifier.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"approve\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ooAssertLogsLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysL2OutputOracleVerifier.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structOasysL2OutputOracleVerifier.Setting\",\"name\":\"_l2ooSetting\",\"type\":\"tuple\"}],\"name\":\"setL2ooSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setting\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b50610017565b6113ed806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100675760003560e01c8063358b3cbb1461006d57806374d5ecb71461008b5780637d2d9b63146100a75780639ac45572146100c3578063c0c9ce30146100df578063c4c5b2ff146100fe57610067565b60006000fd5b610075610132565b6040516100829190610e65565b60405180910390f35b6100a560048036038101906100a09190610c1e565b610147565b005b6100c160048036038101906100bc9190610b6e565b610160565b005b6100dd60048036038101906100d89190610b6e565b610421565b005b6100e76106e2565b6040516100f5929190610e3b565b60405180910390f35b61011860048036038101906101139190610c49565b610716565b604051610129959493929190610de0565b60405180910390f35b60006000600050805490509050610144565b90565b80600160005081816101599190611317565b9050505b50565b60006000506040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff168152602001868152602001858036038101906101a39190610bf3565b81526020016101b885856108d463ffffffff16565b815260200160011515815260200150908060018154018082558091505060019003906000526020600020906006020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001016000509090556040820151816002016000506000820151816000016000509060001916905560208201518160010160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060408201518160010160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050506060820151816004016000509080519060200190610305929190610977565b5060808201518160050160006101000a81548160ff02191690831515021790555050508473ffffffffffffffffffffffffffffffffffffffff1663931da5d785600160005060000160005054600160005060010160009054906101000a90046fffffffffffffffffffffffffffffffff166040518463ffffffff1660e01b815260040161039493929190610e81565b600060405180830381600087803b1580156103af5760006000fd5b505af11580156103c4573d600060003e3d6000fd5b5050505082600001356000191660001916848673ffffffffffffffffffffffffffffffffffffffff167f3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f60405160405180910390a45b5050505050565b60006000506040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff168152602001868152602001858036038101906104649190610bf3565b815260200161047985856108d463ffffffff16565b815260200160001515815260200150908060018154018082558091505060019003906000526020600020906006020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001016000509090556040820151816002016000506000820151816000016000509060001916905560208201518160010160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060408201518160010160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550505060608201518160040160005090805190602001906105c6929190610977565b5060808201518160050160006101000a81548160ff02191690831515021790555050508473ffffffffffffffffffffffffffffffffffffffff1663a511b01285600160005060000160005054600160005060010160009054906101000a90046fffffffffffffffffffffffffffffffff166040518463ffffffff1660e01b815260040161065593929190610e81565b600060405180830381600087803b1580156106705760006000fd5b505af1158015610685573d600060003e3d6000fd5b5050505082600001356000191660001916848673ffffffffffffffffffffffffffffffffffffffff167ff7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca360405160405180910390a45b5050505050565b60016000508060000160005054908060010160009054906101000a90046fffffffffffffffffffffffffffffffff16905082565b6000600050818154811061072957600080fd5b906000526020600020906006020160005b915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000505490806002016000506040518060600160405290816000820160005054600019166000191681526020016001820160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020016001820160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815260200150509080600401600050805461083e9061108e565b80601f016020809104026020016040519081016040528092919081815260200182805461086a9061108e565b80156108b75780601f1061088c576101008083540402835291602001916108b7565b820191906000526020600020905b81548152906001019060200180831161089a57829003601f168201915b5050505050908060050160009054906101000a900460ff16905085565b60606000600090505b8383905081101561096f57818484838181101515610924577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020028101906109369190610eb9565b60405160200161094893929190610db9565b604051602081830303815290604052915081505b8080610967906110f5565b9150506108dd565b505b92915050565b8280546109839061108e565b90600052602060002090601f0160209004810192826109a557600085556109f1565b82601f106109be57805160ff19168380011785556109f1565b828001600101855582156109f1579182015b828111156109f057825182600050909055916020019190600101906109d0565b5b5090506109fe9190610a02565b5090565b610a07565b80821115610a215760008181506000905550600101610a07565b5090566113b6565b600081359050610a388161134a565b5b92915050565b6000600083601f8401121515610a555760006000fd5b8235905067ffffffffffffffff811115610a6f5760006000fd5b602083019150836020820283011115610a885760006000fd5b5b9250929050565b600081359050610a9f81611365565b5b92915050565b600060608284031215610ab95760006000fd5b8190505b92915050565b600060608284031215610ad65760006000fd5b610ae06060610f17565b90506000610af084828501610a90565b6000830152506020610b0484828501610b42565b6020830152506040610b1884828501610b42565b6040830152505b92915050565b600060408284031215610b385760006000fd5b8190505b92915050565b600081359050610b5181611380565b5b92915050565b600081359050610b678161139b565b5b92915050565b6000600060006000600060c08688031215610b895760006000fd5b6000610b9788828901610a29565b9550506020610ba888828901610b58565b9450506040610bb988828901610aa6565b93505060a086013567ffffffffffffffff811115610bd75760006000fd5b610be388828901610a3f565b92509250505b9295509295909350565b600060608284031215610c065760006000fd5b6000610c1484828501610ac3565b9150505b92915050565b600060408284031215610c315760006000fd5b6000610c3f84828501610b25565b9150505b92915050565b600060208284031215610c5c5760006000fd5b6000610c6a84828501610b58565b9150505b92915050565b610c7d81610f68565b82525b5050565b610c8d81610f7b565b82525b5050565b610c9d81610f88565b82525b5050565b610cad81610f88565b82525b5050565b6000610cc08385610f5c565b9350610ccd838584611002565b82840190505b9392505050565b6000610ce582610f3e565b610cef8185610f4a565b9350610cff818560208601611012565b610d088161124f565b84019150505b92915050565b6000610d1f82610f3e565b610d298185610f5c565b9350610d39818560208601611012565b8084019150505b92915050565b606082016000820151610d5c6000850182610c94565b506020820151610d6f6020850182610d89565b506040820151610d826040850182610d89565b50505b5050565b610d9281610f93565b82525b5050565b610da281610f93565b82525b5050565b610db281610fd1565b82525b5050565b6000610dc58286610d14565b9150610dd2828486610cb4565b91508190505b949350505050565b600060e082019050610df56000830188610c74565b610e026020830187610da9565b610e0f6040830186610d46565b81810360a0830152610e218185610cda565b9050610e3060c0830184610c84565b5b9695505050505050565b6000604082019050610e506000830185610ca4565b610e5d6020830184610d99565b5b9392505050565b6000602082019050610e7a6000830184610da9565b5b92915050565b6000606082019050610e966000830186610da9565b610ea36020830185610ca4565b610eb06040830184610d99565b5b949350505050565b60006000833560016020038436030381121515610ed65760006000fd5b80840192508235915067ffffffffffffffff821115610ef55760006000fd5b602083019250600182023603831315610f0e5760006000fd5b505b9250929050565b6000610f21610f33565b9050610f2d82826110c3565b5b919050565b600060405190505b90565b6000815190505b919050565b60008282526020820190505b92915050565b60008190505b92915050565b6000610f7382610fb0565b90505b919050565b600081151590505b919050565b60008190505b919050565b60006fffffffffffffffffffffffffffffffff821690505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b6000610fe782610f88565b90505b919050565b6000610ffa82610f93565b90505b919050565b828183376000838301525b505050565b60005b838110156110315780820151818401525b602081019050611015565b83811115611040576000848401525b505b505050565b60008101600083018061105981611221565b905061106581846112f3565b50505060018101602083018061107a81611238565b90506110868184611326565b5050505b5050565b6000600282049050600182168015156110a857607f821691505b602082108114156110bc576110bb6111a1565b5b505b919050565b6110cc8261124f565b810181811067ffffffffffffffff821117156110eb576110ea6111d2565b5b80604052505b5050565b600061110082610fd1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561113357611132611170565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b600061120e8261126f565b90505b919050565b60008190505b919050565b6000813561122e81611365565b809150505b919050565b6000813561124581611380565b809150505b919050565b6000601f19601f83011690505b919050565b60008160001b90505b919050565b60008160001c90505b919050565b60006fffffffffffffffffffffffffffffffff61129984611261565b935080198316925080841683179150505b92915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6112dc84611261565b935080198316925080841683179150505b92915050565b6112fc82610fdc565b61130f61130882611203565b83546112b0565b8255505b5050565b6113218282611047565b5b5050565b61132f82610fef565b61134261133b82611216565b835461127d565b8255505b5050565b61135381610f68565b811415156113615760006000fd5b5b50565b61136e81610f88565b8114151561137c5760006000fd5b5b50565b61138981610f93565b811415156113975760006000fd5b5b50565b6113a481610fd1565b811415156113b25760006000fd5b5b50565bfea26469706673582212207889bd33f13309b3d1e7665a1a425c2b15761ea66daffbf3c82f5c995b1045da64736f6c63430008020033",
}

// L2ooverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ooverifierMetaData.ABI instead.
var L2ooverifierABI = L2ooverifierMetaData.ABI

// L2ooverifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ooverifierMetaData.Bin instead.
var L2ooverifierBin = L2ooverifierMetaData.Bin

// DeployL2ooverifier deploys a new Ethereum contract, binding an instance of L2ooverifier to it.
func DeployL2ooverifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2ooverifier, error) {
	parsed, err := L2ooverifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ooverifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ooverifier{L2ooverifierCaller: L2ooverifierCaller{contract: contract}, L2ooverifierTransactor: L2ooverifierTransactor{contract: contract}, L2ooverifierFilterer: L2ooverifierFilterer{contract: contract}}, nil
}

// L2ooverifier is an auto generated Go binding around an Ethereum contract.
type L2ooverifier struct {
	L2ooverifierCaller     // Read-only binding to the contract
	L2ooverifierTransactor // Write-only binding to the contract
	L2ooverifierFilterer   // Log filterer for contract events
}

// L2ooverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ooverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ooverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ooverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ooverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ooverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ooverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ooverifierSession struct {
	Contract     *L2ooverifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ooverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ooverifierCallerSession struct {
	Contract *L2ooverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// L2ooverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ooverifierTransactorSession struct {
	Contract     *L2ooverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2ooverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ooverifierRaw struct {
	Contract *L2ooverifier // Generic contract binding to access the raw methods on
}

// L2ooverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ooverifierCallerRaw struct {
	Contract *L2ooverifierCaller // Generic read-only contract binding to access the raw methods on
}

// L2ooverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ooverifierTransactorRaw struct {
	Contract *L2ooverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ooverifier creates a new instance of L2ooverifier, bound to a specific deployed contract.
func NewL2ooverifier(address common.Address, backend bind.ContractBackend) (*L2ooverifier, error) {
	contract, err := bindL2ooverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ooverifier{L2ooverifierCaller: L2ooverifierCaller{contract: contract}, L2ooverifierTransactor: L2ooverifierTransactor{contract: contract}, L2ooverifierFilterer: L2ooverifierFilterer{contract: contract}}, nil
}

// NewL2ooverifierCaller creates a new read-only instance of L2ooverifier, bound to a specific deployed contract.
func NewL2ooverifierCaller(address common.Address, caller bind.ContractCaller) (*L2ooverifierCaller, error) {
	contract, err := bindL2ooverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ooverifierCaller{contract: contract}, nil
}

// NewL2ooverifierTransactor creates a new write-only instance of L2ooverifier, bound to a specific deployed contract.
func NewL2ooverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ooverifierTransactor, error) {
	contract, err := bindL2ooverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ooverifierTransactor{contract: contract}, nil
}

// NewL2ooverifierFilterer creates a new log filterer instance of L2ooverifier, bound to a specific deployed contract.
func NewL2ooverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ooverifierFilterer, error) {
	contract, err := bindL2ooverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ooverifierFilterer{contract: contract}, nil
}

// bindL2ooverifier binds a generic wrapper to an already deployed contract.
func bindL2ooverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2ooverifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ooverifier *L2ooverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ooverifier.Contract.L2ooverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ooverifier *L2ooverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ooverifier.Contract.L2ooverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ooverifier *L2ooverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ooverifier.Contract.L2ooverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ooverifier *L2ooverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ooverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ooverifier *L2ooverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ooverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ooverifier *L2ooverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ooverifier.Contract.contract.Transact(opts, method, params...)
}

// AssertLogs is a free data retrieval call binding the contract method 0xc4c5b2ff.
//
// Solidity: function assertLogs(uint256 ) view returns(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes signatures, bool approve)
func (_L2ooverifier *L2ooverifierCaller) AssertLogs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	L2Output       OasysL2OutputOracleVerifierOutputProposal
	Signatures     []byte
	Approve        bool
}, error) {
	var out []interface{}
	err := _L2ooverifier.contract.Call(opts, &out, "assertLogs", arg0)

	outstruct := new(struct {
		L2OutputOracle common.Address
		L2OutputIndex  *big.Int
		L2Output       OasysL2OutputOracleVerifierOutputProposal
		Signatures     []byte
		Approve        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.L2OutputOracle = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.L2OutputIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.L2Output = *abi.ConvertType(out[2], new(OasysL2OutputOracleVerifierOutputProposal)).(*OasysL2OutputOracleVerifierOutputProposal)
	outstruct.Signatures = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.Approve = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// AssertLogs is a free data retrieval call binding the contract method 0xc4c5b2ff.
//
// Solidity: function assertLogs(uint256 ) view returns(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes signatures, bool approve)
func (_L2ooverifier *L2ooverifierSession) AssertLogs(arg0 *big.Int) (struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	L2Output       OasysL2OutputOracleVerifierOutputProposal
	Signatures     []byte
	Approve        bool
}, error) {
	return _L2ooverifier.Contract.AssertLogs(&_L2ooverifier.CallOpts, arg0)
}

// AssertLogs is a free data retrieval call binding the contract method 0xc4c5b2ff.
//
// Solidity: function assertLogs(uint256 ) view returns(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes signatures, bool approve)
func (_L2ooverifier *L2ooverifierCallerSession) AssertLogs(arg0 *big.Int) (struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	L2Output       OasysL2OutputOracleVerifierOutputProposal
	Signatures     []byte
	Approve        bool
}, error) {
	return _L2ooverifier.Contract.AssertLogs(&_L2ooverifier.CallOpts, arg0)
}

// L2ooAssertLogsLen is a free data retrieval call binding the contract method 0x358b3cbb.
//
// Solidity: function l2ooAssertLogsLen() view returns(uint256)
func (_L2ooverifier *L2ooverifierCaller) L2ooAssertLogsLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ooverifier.contract.Call(opts, &out, "l2ooAssertLogsLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ooAssertLogsLen is a free data retrieval call binding the contract method 0x358b3cbb.
//
// Solidity: function l2ooAssertLogsLen() view returns(uint256)
func (_L2ooverifier *L2ooverifierSession) L2ooAssertLogsLen() (*big.Int, error) {
	return _L2ooverifier.Contract.L2ooAssertLogsLen(&_L2ooverifier.CallOpts)
}

// L2ooAssertLogsLen is a free data retrieval call binding the contract method 0x358b3cbb.
//
// Solidity: function l2ooAssertLogsLen() view returns(uint256)
func (_L2ooverifier *L2ooverifierCallerSession) L2ooAssertLogsLen() (*big.Int, error) {
	return _L2ooverifier.Contract.L2ooAssertLogsLen(&_L2ooverifier.CallOpts)
}

// Setting is a free data retrieval call binding the contract method 0xc0c9ce30.
//
// Solidity: function setting() view returns(bytes32 outputRoot, uint128 l2BlockNumber)
func (_L2ooverifier *L2ooverifierCaller) Setting(opts *bind.CallOpts) (struct {
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _L2ooverifier.contract.Call(opts, &out, "setting")

	outstruct := new(struct {
		OutputRoot    [32]byte
		L2BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutputRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.L2BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Setting is a free data retrieval call binding the contract method 0xc0c9ce30.
//
// Solidity: function setting() view returns(bytes32 outputRoot, uint128 l2BlockNumber)
func (_L2ooverifier *L2ooverifierSession) Setting() (struct {
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _L2ooverifier.Contract.Setting(&_L2ooverifier.CallOpts)
}

// Setting is a free data retrieval call binding the contract method 0xc0c9ce30.
//
// Solidity: function setting() view returns(bytes32 outputRoot, uint128 l2BlockNumber)
func (_L2ooverifier *L2ooverifierCallerSession) Setting() (struct {
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _L2ooverifier.Contract.Setting(&_L2ooverifier.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_L2ooverifier *L2ooverifierTransactor) Approve(opts *bind.TransactOpts, l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysL2OutputOracleVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _L2ooverifier.contract.Transact(opts, "approve", l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_L2ooverifier *L2ooverifierSession) Approve(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysL2OutputOracleVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _L2ooverifier.Contract.Approve(&_L2ooverifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_L2ooverifier *L2ooverifierTransactorSession) Approve(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysL2OutputOracleVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _L2ooverifier.Contract.Approve(&_L2ooverifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_L2ooverifier *L2ooverifierTransactor) Reject(opts *bind.TransactOpts, l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysL2OutputOracleVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _L2ooverifier.contract.Transact(opts, "reject", l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_L2ooverifier *L2ooverifierSession) Reject(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysL2OutputOracleVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _L2ooverifier.Contract.Reject(&_L2ooverifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_L2ooverifier *L2ooverifierTransactorSession) Reject(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output OasysL2OutputOracleVerifierOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _L2ooverifier.Contract.Reject(&_L2ooverifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// SetL2ooSetting is a paid mutator transaction binding the contract method 0x74d5ecb7.
//
// Solidity: function setL2ooSetting((bytes32,uint128) _l2ooSetting) returns()
func (_L2ooverifier *L2ooverifierTransactor) SetL2ooSetting(opts *bind.TransactOpts, _l2ooSetting OasysL2OutputOracleVerifierSetting) (*types.Transaction, error) {
	return _L2ooverifier.contract.Transact(opts, "setL2ooSetting", _l2ooSetting)
}

// SetL2ooSetting is a paid mutator transaction binding the contract method 0x74d5ecb7.
//
// Solidity: function setL2ooSetting((bytes32,uint128) _l2ooSetting) returns()
func (_L2ooverifier *L2ooverifierSession) SetL2ooSetting(_l2ooSetting OasysL2OutputOracleVerifierSetting) (*types.Transaction, error) {
	return _L2ooverifier.Contract.SetL2ooSetting(&_L2ooverifier.TransactOpts, _l2ooSetting)
}

// SetL2ooSetting is a paid mutator transaction binding the contract method 0x74d5ecb7.
//
// Solidity: function setL2ooSetting((bytes32,uint128) _l2ooSetting) returns()
func (_L2ooverifier *L2ooverifierTransactorSession) SetL2ooSetting(_l2ooSetting OasysL2OutputOracleVerifierSetting) (*types.Transaction, error) {
	return _L2ooverifier.Contract.SetL2ooSetting(&_L2ooverifier.TransactOpts, _l2ooSetting)
}

// L2ooverifierL2OutputApprovedIterator is returned from FilterL2OutputApproved and is used to iterate over the raw logs and unpacked data for L2OutputApproved events raised by the L2ooverifier contract.
type L2ooverifierL2OutputApprovedIterator struct {
	Event *L2ooverifierL2OutputApproved // Event containing the contract specifics and raw log

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
func (it *L2ooverifierL2OutputApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ooverifierL2OutputApproved)
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
		it.Event = new(L2ooverifierL2OutputApproved)
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
func (it *L2ooverifierL2OutputApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ooverifierL2OutputApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ooverifierL2OutputApproved represents a L2OutputApproved event raised by the L2ooverifier contract.
type L2ooverifierL2OutputApproved struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	OutputRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterL2OutputApproved is a free log retrieval operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_L2ooverifier *L2ooverifierFilterer) FilterL2OutputApproved(opts *bind.FilterOpts, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (*L2ooverifierL2OutputApprovedIterator, error) {

	var l2OutputOracleRule []interface{}
	for _, l2OutputOracleItem := range l2OutputOracle {
		l2OutputOracleRule = append(l2OutputOracleRule, l2OutputOracleItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _L2ooverifier.contract.FilterLogs(opts, "L2OutputApproved", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &L2ooverifierL2OutputApprovedIterator{contract: _L2ooverifier.contract, event: "L2OutputApproved", logs: logs, sub: sub}, nil
}

// WatchL2OutputApproved is a free log subscription operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_L2ooverifier *L2ooverifierFilterer) WatchL2OutputApproved(opts *bind.WatchOpts, sink chan<- *L2ooverifierL2OutputApproved, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (event.Subscription, error) {

	var l2OutputOracleRule []interface{}
	for _, l2OutputOracleItem := range l2OutputOracle {
		l2OutputOracleRule = append(l2OutputOracleRule, l2OutputOracleItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _L2ooverifier.contract.WatchLogs(opts, "L2OutputApproved", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ooverifierL2OutputApproved)
				if err := _L2ooverifier.contract.UnpackLog(event, "L2OutputApproved", log); err != nil {
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

// ParseL2OutputApproved is a log parse operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_L2ooverifier *L2ooverifierFilterer) ParseL2OutputApproved(log types.Log) (*L2ooverifierL2OutputApproved, error) {
	event := new(L2ooverifierL2OutputApproved)
	if err := _L2ooverifier.contract.UnpackLog(event, "L2OutputApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ooverifierL2OutputRejectedIterator is returned from FilterL2OutputRejected and is used to iterate over the raw logs and unpacked data for L2OutputRejected events raised by the L2ooverifier contract.
type L2ooverifierL2OutputRejectedIterator struct {
	Event *L2ooverifierL2OutputRejected // Event containing the contract specifics and raw log

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
func (it *L2ooverifierL2OutputRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ooverifierL2OutputRejected)
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
		it.Event = new(L2ooverifierL2OutputRejected)
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
func (it *L2ooverifierL2OutputRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ooverifierL2OutputRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ooverifierL2OutputRejected represents a L2OutputRejected event raised by the L2ooverifier contract.
type L2ooverifierL2OutputRejected struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	OutputRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterL2OutputRejected is a free log retrieval operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_L2ooverifier *L2ooverifierFilterer) FilterL2OutputRejected(opts *bind.FilterOpts, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (*L2ooverifierL2OutputRejectedIterator, error) {

	var l2OutputOracleRule []interface{}
	for _, l2OutputOracleItem := range l2OutputOracle {
		l2OutputOracleRule = append(l2OutputOracleRule, l2OutputOracleItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _L2ooverifier.contract.FilterLogs(opts, "L2OutputRejected", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &L2ooverifierL2OutputRejectedIterator{contract: _L2ooverifier.contract, event: "L2OutputRejected", logs: logs, sub: sub}, nil
}

// WatchL2OutputRejected is a free log subscription operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_L2ooverifier *L2ooverifierFilterer) WatchL2OutputRejected(opts *bind.WatchOpts, sink chan<- *L2ooverifierL2OutputRejected, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (event.Subscription, error) {

	var l2OutputOracleRule []interface{}
	for _, l2OutputOracleItem := range l2OutputOracle {
		l2OutputOracleRule = append(l2OutputOracleRule, l2OutputOracleItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _L2ooverifier.contract.WatchLogs(opts, "L2OutputRejected", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ooverifierL2OutputRejected)
				if err := _L2ooverifier.contract.UnpackLog(event, "L2OutputRejected", log); err != nil {
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

// ParseL2OutputRejected is a log parse operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_L2ooverifier *L2ooverifierFilterer) ParseL2OutputRejected(log types.Log) (*L2ooverifierL2OutputRejected, error) {
	event := new(L2ooverifierL2OutputRejected)
	if err := _L2ooverifier.contract.UnpackLog(event, "L2OutputRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
