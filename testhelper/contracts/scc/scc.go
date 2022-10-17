// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scc

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
)

// SccMetaData contains all meta data concerning the Scc contract.
var SccMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"OtherEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_batchSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prevTotalElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"StateBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchVerified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"emitOtherEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"emitStateBatchAppended\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"emitStateBatchDeleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"emitStateBatchFailed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"emitStateBatchVerified\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b50610017565b610828806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100675760003560e01c806320c590211461006d578063875dcafd14610089578063982bc5b0146100a5578063afeab107146100c1578063d1f93ca6146100dd578063fc7e9c6f146100f957610067565b60006000fd5b610087600480360381019061008291906103c6565b610117565b005b6100a3600480360381019061009e9190610405565b610154565b005b6100bf60048036038101906100ba91906103c6565b61019a565b005b6100db60048036038101906100d6919061039b565b610221565b005b6100f760048036038101906100f291906103c6565b610252565b005b6101016102f4565b60405161010e91906105a1565b60405180910390f35b817f8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64826040516101479190610517565b60405180910390a25b5050565b847f16be4c5129a4e03cf3350262e181dc02ddfb4a6008d925368c0899fcd97ca9c58585858560405161018a9493929190610533565b60405180910390a25b5050505050565b60006000505482101515156101e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101db90610580565b60405180910390fd5b817f1723478079cff2362bd8896c78c4c8bee974428fc01131b52d79078349af3e10826040516102149190610517565b60405180910390a25b5050565b807f43b925d43c14ebe0fced53177afd61fc35ed62026274d516917655247e98f99360405160405180910390a25b50565b6000600050548214151561029b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029290610580565b60405180910390fd5b6000600081815054809291906102b0906106d3565b91905090905550817fc032f530357a4853a125c129880b7801f1f30fb46fdd0e2f3bbc6e053199dca9826040516102e79190610517565b60405180910390a25b5050565b60006000505481566107f1565b600061031461030f846105e4565b6105bd565b90508281526020810184848401111561032d5760006000fd5b61033884828561065c565b505b9392505050565b600081359050610350816107bb565b5b92915050565b600082601f830112151561036b5760006000fd5b813561037b848260208601610301565b9150505b92915050565b600081359050610394816107d6565b5b92915050565b6000602082840312156103ae5760006000fd5b60006103bc84828501610385565b9150505b92915050565b60006000604083850312156103db5760006000fd5b60006103e985828601610385565b92505060206103fa85828601610341565b9150505b9250929050565b6000600060006000600060a086880312156104205760006000fd5b600061042e88828901610385565b955050602061043f88828901610341565b945050604061045088828901610385565b935050606061046188828901610385565b925050608086013567ffffffffffffffff81111561047f5760006000fd5b61048b88828901610357565b9150505b9295509295909350565b6104a281610646565b82525b5050565b60006104b482610616565b6104be8185610622565b93506104ce81856020860161066c565b6104d78161077f565b84019150505b92915050565b60006104f0601483610634565b91506104fb82610791565b6020820190505b919050565b61051081610651565b82525b5050565b600060208201905061052c6000830184610499565b5b92915050565b60006080820190506105486000830187610499565b6105556020830186610507565b6105626040830185610507565b818103606083015261057481846104a9565b90505b95945050505050565b60006020820190508181036000830152610599816104e3565b90505b919050565b60006020820190506105b66000830184610507565b5b92915050565b60006105c76105d9565b90506105d382826106a1565b5b919050565b600060405190505b90565b600067ffffffffffffffff8211156105ff576105fe61074e565b5b6106088261077f565b90506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b919050565b60008190505b919050565b828183376000838301525b505050565b60005b8381101561068b5780820151818401525b60208101905061066f565b8381111561069a576000848401525b505b505050565b6106aa8261077f565b810181811067ffffffffffffffff821117156106c9576106c861074e565b5b80604052505b5050565b60006106de82610651565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156107115761071061071d565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b7f496e76616c696420626174636820696e6465782e00000000000000000000000060008201525b50565b6107c481610646565b811415156107d25760006000fd5b5b50565b6107df81610651565b811415156107ed5760006000fd5b5b50565bfea2646970667358221220b5533d426207d5d4557160bbb6c7051644233493544584fde6f18e3c3b6b2c0e64736f6c63430008020033",
}

// SccABI is the input ABI used to generate the binding from.
// Deprecated: Use SccMetaData.ABI instead.
var SccABI = SccMetaData.ABI

// SccBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SccMetaData.Bin instead.
var SccBin = SccMetaData.Bin

// DeployScc deploys a new Ethereum contract, binding an instance of Scc to it.
func DeployScc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Scc, error) {
	parsed, err := SccMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SccBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Scc{SccCaller: SccCaller{contract: contract}, SccTransactor: SccTransactor{contract: contract}, SccFilterer: SccFilterer{contract: contract}}, nil
}

// Scc is an auto generated Go binding around an Ethereum contract.
type Scc struct {
	SccCaller     // Read-only binding to the contract
	SccTransactor // Write-only binding to the contract
	SccFilterer   // Log filterer for contract events
}

// SccCaller is an auto generated read-only Go binding around an Ethereum contract.
type SccCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SccTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SccTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SccFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SccFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SccSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SccSession struct {
	Contract     *Scc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SccCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SccCallerSession struct {
	Contract *SccCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SccTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SccTransactorSession struct {
	Contract     *SccTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SccRaw is an auto generated low-level Go binding around an Ethereum contract.
type SccRaw struct {
	Contract *Scc // Generic contract binding to access the raw methods on
}

// SccCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SccCallerRaw struct {
	Contract *SccCaller // Generic read-only contract binding to access the raw methods on
}

// SccTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SccTransactorRaw struct {
	Contract *SccTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScc creates a new instance of Scc, bound to a specific deployed contract.
func NewScc(address common.Address, backend bind.ContractBackend) (*Scc, error) {
	contract, err := bindScc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Scc{SccCaller: SccCaller{contract: contract}, SccTransactor: SccTransactor{contract: contract}, SccFilterer: SccFilterer{contract: contract}}, nil
}

// NewSccCaller creates a new read-only instance of Scc, bound to a specific deployed contract.
func NewSccCaller(address common.Address, caller bind.ContractCaller) (*SccCaller, error) {
	contract, err := bindScc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SccCaller{contract: contract}, nil
}

// NewSccTransactor creates a new write-only instance of Scc, bound to a specific deployed contract.
func NewSccTransactor(address common.Address, transactor bind.ContractTransactor) (*SccTransactor, error) {
	contract, err := bindScc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SccTransactor{contract: contract}, nil
}

// NewSccFilterer creates a new log filterer instance of Scc, bound to a specific deployed contract.
func NewSccFilterer(address common.Address, filterer bind.ContractFilterer) (*SccFilterer, error) {
	contract, err := bindScc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SccFilterer{contract: contract}, nil
}

// bindScc binds a generic wrapper to an already deployed contract.
func bindScc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SccABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scc *SccRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scc.Contract.SccCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scc *SccRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scc.Contract.SccTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scc *SccRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scc.Contract.SccTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scc *SccCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scc *SccTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scc *SccTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scc.Contract.contract.Transact(opts, method, params...)
}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint256)
func (_Scc *SccCaller) NextIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Scc.contract.Call(opts, &out, "nextIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint256)
func (_Scc *SccSession) NextIndex() (*big.Int, error) {
	return _Scc.Contract.NextIndex(&_Scc.CallOpts)
}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint256)
func (_Scc *SccCallerSession) NextIndex() (*big.Int, error) {
	return _Scc.Contract.NextIndex(&_Scc.CallOpts)
}

// EmitOtherEvent is a paid mutator transaction binding the contract method 0xafeab107.
//
// Solidity: function emitOtherEvent(uint256 batchIndex) returns()
func (_Scc *SccTransactor) EmitOtherEvent(opts *bind.TransactOpts, batchIndex *big.Int) (*types.Transaction, error) {
	return _Scc.contract.Transact(opts, "emitOtherEvent", batchIndex)
}

// EmitOtherEvent is a paid mutator transaction binding the contract method 0xafeab107.
//
// Solidity: function emitOtherEvent(uint256 batchIndex) returns()
func (_Scc *SccSession) EmitOtherEvent(batchIndex *big.Int) (*types.Transaction, error) {
	return _Scc.Contract.EmitOtherEvent(&_Scc.TransactOpts, batchIndex)
}

// EmitOtherEvent is a paid mutator transaction binding the contract method 0xafeab107.
//
// Solidity: function emitOtherEvent(uint256 batchIndex) returns()
func (_Scc *SccTransactorSession) EmitOtherEvent(batchIndex *big.Int) (*types.Transaction, error) {
	return _Scc.Contract.EmitOtherEvent(&_Scc.TransactOpts, batchIndex)
}

// EmitStateBatchAppended is a paid mutator transaction binding the contract method 0x875dcafd.
//
// Solidity: function emitStateBatchAppended(uint256 batchIndex, bytes32 batchRoot, uint256 batchSize, uint256 prevTotalElements, bytes extraData) returns()
func (_Scc *SccTransactor) EmitStateBatchAppended(opts *bind.TransactOpts, batchIndex *big.Int, batchRoot [32]byte, batchSize *big.Int, prevTotalElements *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Scc.contract.Transact(opts, "emitStateBatchAppended", batchIndex, batchRoot, batchSize, prevTotalElements, extraData)
}

// EmitStateBatchAppended is a paid mutator transaction binding the contract method 0x875dcafd.
//
// Solidity: function emitStateBatchAppended(uint256 batchIndex, bytes32 batchRoot, uint256 batchSize, uint256 prevTotalElements, bytes extraData) returns()
func (_Scc *SccSession) EmitStateBatchAppended(batchIndex *big.Int, batchRoot [32]byte, batchSize *big.Int, prevTotalElements *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Scc.Contract.EmitStateBatchAppended(&_Scc.TransactOpts, batchIndex, batchRoot, batchSize, prevTotalElements, extraData)
}

// EmitStateBatchAppended is a paid mutator transaction binding the contract method 0x875dcafd.
//
// Solidity: function emitStateBatchAppended(uint256 batchIndex, bytes32 batchRoot, uint256 batchSize, uint256 prevTotalElements, bytes extraData) returns()
func (_Scc *SccTransactorSession) EmitStateBatchAppended(batchIndex *big.Int, batchRoot [32]byte, batchSize *big.Int, prevTotalElements *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Scc.Contract.EmitStateBatchAppended(&_Scc.TransactOpts, batchIndex, batchRoot, batchSize, prevTotalElements, extraData)
}

// EmitStateBatchDeleted is a paid mutator transaction binding the contract method 0x20c59021.
//
// Solidity: function emitStateBatchDeleted(uint256 batchIndex, bytes32 batchRoot) returns()
func (_Scc *SccTransactor) EmitStateBatchDeleted(opts *bind.TransactOpts, batchIndex *big.Int, batchRoot [32]byte) (*types.Transaction, error) {
	return _Scc.contract.Transact(opts, "emitStateBatchDeleted", batchIndex, batchRoot)
}

// EmitStateBatchDeleted is a paid mutator transaction binding the contract method 0x20c59021.
//
// Solidity: function emitStateBatchDeleted(uint256 batchIndex, bytes32 batchRoot) returns()
func (_Scc *SccSession) EmitStateBatchDeleted(batchIndex *big.Int, batchRoot [32]byte) (*types.Transaction, error) {
	return _Scc.Contract.EmitStateBatchDeleted(&_Scc.TransactOpts, batchIndex, batchRoot)
}

// EmitStateBatchDeleted is a paid mutator transaction binding the contract method 0x20c59021.
//
// Solidity: function emitStateBatchDeleted(uint256 batchIndex, bytes32 batchRoot) returns()
func (_Scc *SccTransactorSession) EmitStateBatchDeleted(batchIndex *big.Int, batchRoot [32]byte) (*types.Transaction, error) {
	return _Scc.Contract.EmitStateBatchDeleted(&_Scc.TransactOpts, batchIndex, batchRoot)
}

// EmitStateBatchFailed is a paid mutator transaction binding the contract method 0x982bc5b0.
//
// Solidity: function emitStateBatchFailed(uint256 batchIndex, bytes32 batchRoot) returns()
func (_Scc *SccTransactor) EmitStateBatchFailed(opts *bind.TransactOpts, batchIndex *big.Int, batchRoot [32]byte) (*types.Transaction, error) {
	return _Scc.contract.Transact(opts, "emitStateBatchFailed", batchIndex, batchRoot)
}

// EmitStateBatchFailed is a paid mutator transaction binding the contract method 0x982bc5b0.
//
// Solidity: function emitStateBatchFailed(uint256 batchIndex, bytes32 batchRoot) returns()
func (_Scc *SccSession) EmitStateBatchFailed(batchIndex *big.Int, batchRoot [32]byte) (*types.Transaction, error) {
	return _Scc.Contract.EmitStateBatchFailed(&_Scc.TransactOpts, batchIndex, batchRoot)
}

// EmitStateBatchFailed is a paid mutator transaction binding the contract method 0x982bc5b0.
//
// Solidity: function emitStateBatchFailed(uint256 batchIndex, bytes32 batchRoot) returns()
func (_Scc *SccTransactorSession) EmitStateBatchFailed(batchIndex *big.Int, batchRoot [32]byte) (*types.Transaction, error) {
	return _Scc.Contract.EmitStateBatchFailed(&_Scc.TransactOpts, batchIndex, batchRoot)
}

// EmitStateBatchVerified is a paid mutator transaction binding the contract method 0xd1f93ca6.
//
// Solidity: function emitStateBatchVerified(uint256 batchIndex, bytes32 batchRoot) returns()
func (_Scc *SccTransactor) EmitStateBatchVerified(opts *bind.TransactOpts, batchIndex *big.Int, batchRoot [32]byte) (*types.Transaction, error) {
	return _Scc.contract.Transact(opts, "emitStateBatchVerified", batchIndex, batchRoot)
}

// EmitStateBatchVerified is a paid mutator transaction binding the contract method 0xd1f93ca6.
//
// Solidity: function emitStateBatchVerified(uint256 batchIndex, bytes32 batchRoot) returns()
func (_Scc *SccSession) EmitStateBatchVerified(batchIndex *big.Int, batchRoot [32]byte) (*types.Transaction, error) {
	return _Scc.Contract.EmitStateBatchVerified(&_Scc.TransactOpts, batchIndex, batchRoot)
}

// EmitStateBatchVerified is a paid mutator transaction binding the contract method 0xd1f93ca6.
//
// Solidity: function emitStateBatchVerified(uint256 batchIndex, bytes32 batchRoot) returns()
func (_Scc *SccTransactorSession) EmitStateBatchVerified(batchIndex *big.Int, batchRoot [32]byte) (*types.Transaction, error) {
	return _Scc.Contract.EmitStateBatchVerified(&_Scc.TransactOpts, batchIndex, batchRoot)
}

// SccOtherEventIterator is returned from FilterOtherEvent and is used to iterate over the raw logs and unpacked data for OtherEvent events raised by the Scc contract.
type SccOtherEventIterator struct {
	Event *SccOtherEvent // Event containing the contract specifics and raw log

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
func (it *SccOtherEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccOtherEvent)
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
		it.Event = new(SccOtherEvent)
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
func (it *SccOtherEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccOtherEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccOtherEvent represents a OtherEvent event raised by the Scc contract.
type SccOtherEvent struct {
	BatchIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOtherEvent is a free log retrieval operation binding the contract event 0x43b925d43c14ebe0fced53177afd61fc35ed62026274d516917655247e98f993.
//
// Solidity: event OtherEvent(uint256 indexed _batchIndex)
func (_Scc *SccFilterer) FilterOtherEvent(opts *bind.FilterOpts, _batchIndex []*big.Int) (*SccOtherEventIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.FilterLogs(opts, "OtherEvent", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &SccOtherEventIterator{contract: _Scc.contract, event: "OtherEvent", logs: logs, sub: sub}, nil
}

// WatchOtherEvent is a free log subscription operation binding the contract event 0x43b925d43c14ebe0fced53177afd61fc35ed62026274d516917655247e98f993.
//
// Solidity: event OtherEvent(uint256 indexed _batchIndex)
func (_Scc *SccFilterer) WatchOtherEvent(opts *bind.WatchOpts, sink chan<- *SccOtherEvent, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.WatchLogs(opts, "OtherEvent", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccOtherEvent)
				if err := _Scc.contract.UnpackLog(event, "OtherEvent", log); err != nil {
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

// ParseOtherEvent is a log parse operation binding the contract event 0x43b925d43c14ebe0fced53177afd61fc35ed62026274d516917655247e98f993.
//
// Solidity: event OtherEvent(uint256 indexed _batchIndex)
func (_Scc *SccFilterer) ParseOtherEvent(log types.Log) (*SccOtherEvent, error) {
	event := new(SccOtherEvent)
	if err := _Scc.contract.UnpackLog(event, "OtherEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SccStateBatchAppendedIterator is returned from FilterStateBatchAppended and is used to iterate over the raw logs and unpacked data for StateBatchAppended events raised by the Scc contract.
type SccStateBatchAppendedIterator struct {
	Event *SccStateBatchAppended // Event containing the contract specifics and raw log

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
func (it *SccStateBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccStateBatchAppended)
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
		it.Event = new(SccStateBatchAppended)
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
func (it *SccStateBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccStateBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccStateBatchAppended represents a StateBatchAppended event raised by the Scc contract.
type SccStateBatchAppended struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	ExtraData         []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStateBatchAppended is a free log retrieval operation binding the contract event 0x16be4c5129a4e03cf3350262e181dc02ddfb4a6008d925368c0899fcd97ca9c5.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _extraData)
func (_Scc *SccFilterer) FilterStateBatchAppended(opts *bind.FilterOpts, _batchIndex []*big.Int) (*SccStateBatchAppendedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.FilterLogs(opts, "StateBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &SccStateBatchAppendedIterator{contract: _Scc.contract, event: "StateBatchAppended", logs: logs, sub: sub}, nil
}

// WatchStateBatchAppended is a free log subscription operation binding the contract event 0x16be4c5129a4e03cf3350262e181dc02ddfb4a6008d925368c0899fcd97ca9c5.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _extraData)
func (_Scc *SccFilterer) WatchStateBatchAppended(opts *bind.WatchOpts, sink chan<- *SccStateBatchAppended, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.WatchLogs(opts, "StateBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccStateBatchAppended)
				if err := _Scc.contract.UnpackLog(event, "StateBatchAppended", log); err != nil {
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

// ParseStateBatchAppended is a log parse operation binding the contract event 0x16be4c5129a4e03cf3350262e181dc02ddfb4a6008d925368c0899fcd97ca9c5.
//
// Solidity: event StateBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _extraData)
func (_Scc *SccFilterer) ParseStateBatchAppended(log types.Log) (*SccStateBatchAppended, error) {
	event := new(SccStateBatchAppended)
	if err := _Scc.contract.UnpackLog(event, "StateBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SccStateBatchDeletedIterator is returned from FilterStateBatchDeleted and is used to iterate over the raw logs and unpacked data for StateBatchDeleted events raised by the Scc contract.
type SccStateBatchDeletedIterator struct {
	Event *SccStateBatchDeleted // Event containing the contract specifics and raw log

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
func (it *SccStateBatchDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccStateBatchDeleted)
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
		it.Event = new(SccStateBatchDeleted)
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
func (it *SccStateBatchDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccStateBatchDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccStateBatchDeleted represents a StateBatchDeleted event raised by the Scc contract.
type SccStateBatchDeleted struct {
	BatchIndex *big.Int
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStateBatchDeleted is a free log retrieval operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_Scc *SccFilterer) FilterStateBatchDeleted(opts *bind.FilterOpts, _batchIndex []*big.Int) (*SccStateBatchDeletedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.FilterLogs(opts, "StateBatchDeleted", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &SccStateBatchDeletedIterator{contract: _Scc.contract, event: "StateBatchDeleted", logs: logs, sub: sub}, nil
}

// WatchStateBatchDeleted is a free log subscription operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_Scc *SccFilterer) WatchStateBatchDeleted(opts *bind.WatchOpts, sink chan<- *SccStateBatchDeleted, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.WatchLogs(opts, "StateBatchDeleted", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccStateBatchDeleted)
				if err := _Scc.contract.UnpackLog(event, "StateBatchDeleted", log); err != nil {
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

// ParseStateBatchDeleted is a log parse operation binding the contract event 0x8747b69ce8fdb31c3b9b0a67bd8049ad8c1a69ea417b69b12174068abd9cbd64.
//
// Solidity: event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_Scc *SccFilterer) ParseStateBatchDeleted(log types.Log) (*SccStateBatchDeleted, error) {
	event := new(SccStateBatchDeleted)
	if err := _Scc.contract.UnpackLog(event, "StateBatchDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SccStateBatchFailedIterator is returned from FilterStateBatchFailed and is used to iterate over the raw logs and unpacked data for StateBatchFailed events raised by the Scc contract.
type SccStateBatchFailedIterator struct {
	Event *SccStateBatchFailed // Event containing the contract specifics and raw log

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
func (it *SccStateBatchFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccStateBatchFailed)
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
		it.Event = new(SccStateBatchFailed)
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
func (it *SccStateBatchFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccStateBatchFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccStateBatchFailed represents a StateBatchFailed event raised by the Scc contract.
type SccStateBatchFailed struct {
	BatchIndex *big.Int
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStateBatchFailed is a free log retrieval operation binding the contract event 0x1723478079cff2362bd8896c78c4c8bee974428fc01131b52d79078349af3e10.
//
// Solidity: event StateBatchFailed(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_Scc *SccFilterer) FilterStateBatchFailed(opts *bind.FilterOpts, _batchIndex []*big.Int) (*SccStateBatchFailedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.FilterLogs(opts, "StateBatchFailed", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &SccStateBatchFailedIterator{contract: _Scc.contract, event: "StateBatchFailed", logs: logs, sub: sub}, nil
}

// WatchStateBatchFailed is a free log subscription operation binding the contract event 0x1723478079cff2362bd8896c78c4c8bee974428fc01131b52d79078349af3e10.
//
// Solidity: event StateBatchFailed(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_Scc *SccFilterer) WatchStateBatchFailed(opts *bind.WatchOpts, sink chan<- *SccStateBatchFailed, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.WatchLogs(opts, "StateBatchFailed", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccStateBatchFailed)
				if err := _Scc.contract.UnpackLog(event, "StateBatchFailed", log); err != nil {
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

// ParseStateBatchFailed is a log parse operation binding the contract event 0x1723478079cff2362bd8896c78c4c8bee974428fc01131b52d79078349af3e10.
//
// Solidity: event StateBatchFailed(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_Scc *SccFilterer) ParseStateBatchFailed(log types.Log) (*SccStateBatchFailed, error) {
	event := new(SccStateBatchFailed)
	if err := _Scc.contract.UnpackLog(event, "StateBatchFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SccStateBatchVerifiedIterator is returned from FilterStateBatchVerified and is used to iterate over the raw logs and unpacked data for StateBatchVerified events raised by the Scc contract.
type SccStateBatchVerifiedIterator struct {
	Event *SccStateBatchVerified // Event containing the contract specifics and raw log

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
func (it *SccStateBatchVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccStateBatchVerified)
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
		it.Event = new(SccStateBatchVerified)
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
func (it *SccStateBatchVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccStateBatchVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccStateBatchVerified represents a StateBatchVerified event raised by the Scc contract.
type SccStateBatchVerified struct {
	BatchIndex *big.Int
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStateBatchVerified is a free log retrieval operation binding the contract event 0xc032f530357a4853a125c129880b7801f1f30fb46fdd0e2f3bbc6e053199dca9.
//
// Solidity: event StateBatchVerified(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_Scc *SccFilterer) FilterStateBatchVerified(opts *bind.FilterOpts, _batchIndex []*big.Int) (*SccStateBatchVerifiedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.FilterLogs(opts, "StateBatchVerified", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &SccStateBatchVerifiedIterator{contract: _Scc.contract, event: "StateBatchVerified", logs: logs, sub: sub}, nil
}

// WatchStateBatchVerified is a free log subscription operation binding the contract event 0xc032f530357a4853a125c129880b7801f1f30fb46fdd0e2f3bbc6e053199dca9.
//
// Solidity: event StateBatchVerified(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_Scc *SccFilterer) WatchStateBatchVerified(opts *bind.WatchOpts, sink chan<- *SccStateBatchVerified, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Scc.contract.WatchLogs(opts, "StateBatchVerified", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccStateBatchVerified)
				if err := _Scc.contract.UnpackLog(event, "StateBatchVerified", log); err != nil {
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

// ParseStateBatchVerified is a log parse operation binding the contract event 0xc032f530357a4853a125c129880b7801f1f30fb46fdd0e2f3bbc6e053199dca9.
//
// Solidity: event StateBatchVerified(uint256 indexed _batchIndex, bytes32 _batchRoot)
func (_Scc *SccFilterer) ParseStateBatchVerified(log types.Log) (*SccStateBatchVerified, error) {
	event := new(SccStateBatchVerified)
	if err := _Scc.contract.UnpackLog(event, "StateBatchVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
