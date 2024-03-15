// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2oo

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

// L2ooMetaData contains all meta data concerning the L2oo contract.
var L2ooMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"name\":\"OutputFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"OutputProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"name\":\"OutputVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevNextOutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newNextOutputIndex\",\"type\":\"uint256\"}],\"name\":\"OutputsDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"name\":\"emitOutputFailed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"emitOutputProposed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"name\":\"emitOutputVerified\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prevNextOutputIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newNextOutputIndex\",\"type\":\"uint256\"}],\"name\":\"emitOutputsDeleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextVerifyIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setNextVerifyIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b50610017565b61065f806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100675760003560e01c806345c2150c1461006d5780636b405db01461008957806384d67709146100a757806389a5ad39146100c3578063931da5d7146100df578063a511b012146100fb57610067565b60006000fd5b61008760048036038101906100829190610417565b610117565b005b61009161014a565b60405161009e91906104ab565b60405180910390f35b6100c160048036038101906100bc919061039a565b610153565b005b6100dd60048036038101906100d89190610333565b610163565b005b6100f960048036038101906100f491906103c5565b6101a8565b005b610115600480360381019061011091906103c5565b610258565b005b80827f4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b660405160405180910390a35b5050565b60006000505481565b8060006000508190909055505b50565b818385600019167fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e28460405161019991906104ab565b60405180910390a45b50505050565b600060005054831415156101f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e89061048a565b60405180910390fd5b6000600081815054809291906102069061050c565b91905090905550806fffffffffffffffffffffffffffffffff168260001916847f98beebb56f3d9f23ca4c2e85cac607b80e030b3ca74fcc7b38f98bdd5146ef6b60405160405180910390a45b505050565b60006000505483101515156102a2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102999061048a565b60405180910390fd5b806fffffffffffffffffffffffffffffffff168260001916847f39dab377091f907ce5cdb2ef4eeea3961f83ca13a3b6daff133ec7f99a3ec34360405160405180910390a45b50505056610628565b600081359050610300816105d7565b5b92915050565b600081359050610316816105f2565b5b92915050565b60008135905061032c8161060d565b5b92915050565b60006000600060006080858703121561034c5760006000fd5b600061035a878288016102f1565b945050602061036b8782880161031d565b935050604061037c8782880161031d565b925050606061038d8782880161031d565b9150505b92959194509250565b6000602082840312156103ad5760006000fd5b60006103bb8482850161031d565b9150505b92915050565b600060006000606084860312156103dc5760006000fd5b60006103ea8682870161031d565b93505060206103fb868287016102f1565b925050604061040c86828701610307565b9150505b9250925092565b600060006040838503121561042c5760006000fd5b600061043a8582860161031d565b925050602061044b8582860161031d565b9150505b9250929050565b60006104636027836104c7565b915061046e82610587565b6040820190505b919050565b61048381610501565b82525b5050565b600060208201905081810360008301526104a381610456565b90505b919050565b60006020820190506104c0600083018461047a565b5b92915050565b60008282526020820190505b92915050565b60008190505b919050565b60006fffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600061051782610501565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561054a57610549610556565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4c324f75747075744f7261636c653a20496e76616c6964204c32206f7574707560008201527f7420696e6465780000000000000000000000000000000000000000000000000060208201525b50565b6105e0816104d9565b811415156105ee5760006000fd5b5b50565b6105fb816104e4565b811415156106095760006000fd5b5b50565b61061681610501565b811415156106245760006000fd5b5b50565bfea26469706673582212202fff7f36bd94bb71c937b0eae21d621429cd59b1275a080511b05aa1719bfad564736f6c63430008020033",
}

// L2ooABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ooMetaData.ABI instead.
var L2ooABI = L2ooMetaData.ABI

// L2ooBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ooMetaData.Bin instead.
var L2ooBin = L2ooMetaData.Bin

// DeployL2oo deploys a new Ethereum contract, binding an instance of L2oo to it.
func DeployL2oo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2oo, error) {
	parsed, err := L2ooMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ooBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2oo{L2ooCaller: L2ooCaller{contract: contract}, L2ooTransactor: L2ooTransactor{contract: contract}, L2ooFilterer: L2ooFilterer{contract: contract}}, nil
}

// L2oo is an auto generated Go binding around an Ethereum contract.
type L2oo struct {
	L2ooCaller     // Read-only binding to the contract
	L2ooTransactor // Write-only binding to the contract
	L2ooFilterer   // Log filterer for contract events
}

// L2ooCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ooCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ooTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ooTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ooFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ooFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ooSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ooSession struct {
	Contract     *L2oo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ooCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ooCallerSession struct {
	Contract *L2ooCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// L2ooTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ooTransactorSession struct {
	Contract     *L2ooTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ooRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ooRaw struct {
	Contract *L2oo // Generic contract binding to access the raw methods on
}

// L2ooCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ooCallerRaw struct {
	Contract *L2ooCaller // Generic read-only contract binding to access the raw methods on
}

// L2ooTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ooTransactorRaw struct {
	Contract *L2ooTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2oo creates a new instance of L2oo, bound to a specific deployed contract.
func NewL2oo(address common.Address, backend bind.ContractBackend) (*L2oo, error) {
	contract, err := bindL2oo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2oo{L2ooCaller: L2ooCaller{contract: contract}, L2ooTransactor: L2ooTransactor{contract: contract}, L2ooFilterer: L2ooFilterer{contract: contract}}, nil
}

// NewL2ooCaller creates a new read-only instance of L2oo, bound to a specific deployed contract.
func NewL2ooCaller(address common.Address, caller bind.ContractCaller) (*L2ooCaller, error) {
	contract, err := bindL2oo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ooCaller{contract: contract}, nil
}

// NewL2ooTransactor creates a new write-only instance of L2oo, bound to a specific deployed contract.
func NewL2ooTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ooTransactor, error) {
	contract, err := bindL2oo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ooTransactor{contract: contract}, nil
}

// NewL2ooFilterer creates a new log filterer instance of L2oo, bound to a specific deployed contract.
func NewL2ooFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ooFilterer, error) {
	contract, err := bindL2oo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ooFilterer{contract: contract}, nil
}

// bindL2oo binds a generic wrapper to an already deployed contract.
func bindL2oo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2ooMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2oo *L2ooRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2oo.Contract.L2ooCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2oo *L2ooRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2oo.Contract.L2ooTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2oo *L2ooRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2oo.Contract.L2ooTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2oo *L2ooCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2oo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2oo *L2ooTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2oo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2oo *L2ooTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2oo.Contract.contract.Transact(opts, method, params...)
}

// NextVerifyIndex is a free data retrieval call binding the contract method 0x6b405db0.
//
// Solidity: function nextVerifyIndex() view returns(uint256)
func (_L2oo *L2ooCaller) NextVerifyIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2oo.contract.Call(opts, &out, "nextVerifyIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextVerifyIndex is a free data retrieval call binding the contract method 0x6b405db0.
//
// Solidity: function nextVerifyIndex() view returns(uint256)
func (_L2oo *L2ooSession) NextVerifyIndex() (*big.Int, error) {
	return _L2oo.Contract.NextVerifyIndex(&_L2oo.CallOpts)
}

// NextVerifyIndex is a free data retrieval call binding the contract method 0x6b405db0.
//
// Solidity: function nextVerifyIndex() view returns(uint256)
func (_L2oo *L2ooCallerSession) NextVerifyIndex() (*big.Int, error) {
	return _L2oo.Contract.NextVerifyIndex(&_L2oo.CallOpts)
}

// EmitOutputFailed is a paid mutator transaction binding the contract method 0xa511b012.
//
// Solidity: function emitOutputFailed(uint256 l2OutputIndex, bytes32 outputRoot, uint128 l2BlockNumber) returns()
func (_L2oo *L2ooTransactor) EmitOutputFailed(opts *bind.TransactOpts, l2OutputIndex *big.Int, outputRoot [32]byte, l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _L2oo.contract.Transact(opts, "emitOutputFailed", l2OutputIndex, outputRoot, l2BlockNumber)
}

// EmitOutputFailed is a paid mutator transaction binding the contract method 0xa511b012.
//
// Solidity: function emitOutputFailed(uint256 l2OutputIndex, bytes32 outputRoot, uint128 l2BlockNumber) returns()
func (_L2oo *L2ooSession) EmitOutputFailed(l2OutputIndex *big.Int, outputRoot [32]byte, l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.EmitOutputFailed(&_L2oo.TransactOpts, l2OutputIndex, outputRoot, l2BlockNumber)
}

// EmitOutputFailed is a paid mutator transaction binding the contract method 0xa511b012.
//
// Solidity: function emitOutputFailed(uint256 l2OutputIndex, bytes32 outputRoot, uint128 l2BlockNumber) returns()
func (_L2oo *L2ooTransactorSession) EmitOutputFailed(l2OutputIndex *big.Int, outputRoot [32]byte, l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.EmitOutputFailed(&_L2oo.TransactOpts, l2OutputIndex, outputRoot, l2BlockNumber)
}

// EmitOutputProposed is a paid mutator transaction binding the contract method 0x89a5ad39.
//
// Solidity: function emitOutputProposed(bytes32 outputRoot, uint256 l2OutputIndex, uint256 l2BlockNumber, uint256 l1Timestamp) returns()
func (_L2oo *L2ooTransactor) EmitOutputProposed(opts *bind.TransactOpts, outputRoot [32]byte, l2OutputIndex *big.Int, l2BlockNumber *big.Int, l1Timestamp *big.Int) (*types.Transaction, error) {
	return _L2oo.contract.Transact(opts, "emitOutputProposed", outputRoot, l2OutputIndex, l2BlockNumber, l1Timestamp)
}

// EmitOutputProposed is a paid mutator transaction binding the contract method 0x89a5ad39.
//
// Solidity: function emitOutputProposed(bytes32 outputRoot, uint256 l2OutputIndex, uint256 l2BlockNumber, uint256 l1Timestamp) returns()
func (_L2oo *L2ooSession) EmitOutputProposed(outputRoot [32]byte, l2OutputIndex *big.Int, l2BlockNumber *big.Int, l1Timestamp *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.EmitOutputProposed(&_L2oo.TransactOpts, outputRoot, l2OutputIndex, l2BlockNumber, l1Timestamp)
}

// EmitOutputProposed is a paid mutator transaction binding the contract method 0x89a5ad39.
//
// Solidity: function emitOutputProposed(bytes32 outputRoot, uint256 l2OutputIndex, uint256 l2BlockNumber, uint256 l1Timestamp) returns()
func (_L2oo *L2ooTransactorSession) EmitOutputProposed(outputRoot [32]byte, l2OutputIndex *big.Int, l2BlockNumber *big.Int, l1Timestamp *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.EmitOutputProposed(&_L2oo.TransactOpts, outputRoot, l2OutputIndex, l2BlockNumber, l1Timestamp)
}

// EmitOutputVerified is a paid mutator transaction binding the contract method 0x931da5d7.
//
// Solidity: function emitOutputVerified(uint256 l2OutputIndex, bytes32 outputRoot, uint128 l2BlockNumber) returns()
func (_L2oo *L2ooTransactor) EmitOutputVerified(opts *bind.TransactOpts, l2OutputIndex *big.Int, outputRoot [32]byte, l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _L2oo.contract.Transact(opts, "emitOutputVerified", l2OutputIndex, outputRoot, l2BlockNumber)
}

// EmitOutputVerified is a paid mutator transaction binding the contract method 0x931da5d7.
//
// Solidity: function emitOutputVerified(uint256 l2OutputIndex, bytes32 outputRoot, uint128 l2BlockNumber) returns()
func (_L2oo *L2ooSession) EmitOutputVerified(l2OutputIndex *big.Int, outputRoot [32]byte, l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.EmitOutputVerified(&_L2oo.TransactOpts, l2OutputIndex, outputRoot, l2BlockNumber)
}

// EmitOutputVerified is a paid mutator transaction binding the contract method 0x931da5d7.
//
// Solidity: function emitOutputVerified(uint256 l2OutputIndex, bytes32 outputRoot, uint128 l2BlockNumber) returns()
func (_L2oo *L2ooTransactorSession) EmitOutputVerified(l2OutputIndex *big.Int, outputRoot [32]byte, l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.EmitOutputVerified(&_L2oo.TransactOpts, l2OutputIndex, outputRoot, l2BlockNumber)
}

// EmitOutputsDeleted is a paid mutator transaction binding the contract method 0x45c2150c.
//
// Solidity: function emitOutputsDeleted(uint256 prevNextOutputIndex, uint256 newNextOutputIndex) returns()
func (_L2oo *L2ooTransactor) EmitOutputsDeleted(opts *bind.TransactOpts, prevNextOutputIndex *big.Int, newNextOutputIndex *big.Int) (*types.Transaction, error) {
	return _L2oo.contract.Transact(opts, "emitOutputsDeleted", prevNextOutputIndex, newNextOutputIndex)
}

// EmitOutputsDeleted is a paid mutator transaction binding the contract method 0x45c2150c.
//
// Solidity: function emitOutputsDeleted(uint256 prevNextOutputIndex, uint256 newNextOutputIndex) returns()
func (_L2oo *L2ooSession) EmitOutputsDeleted(prevNextOutputIndex *big.Int, newNextOutputIndex *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.EmitOutputsDeleted(&_L2oo.TransactOpts, prevNextOutputIndex, newNextOutputIndex)
}

// EmitOutputsDeleted is a paid mutator transaction binding the contract method 0x45c2150c.
//
// Solidity: function emitOutputsDeleted(uint256 prevNextOutputIndex, uint256 newNextOutputIndex) returns()
func (_L2oo *L2ooTransactorSession) EmitOutputsDeleted(prevNextOutputIndex *big.Int, newNextOutputIndex *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.EmitOutputsDeleted(&_L2oo.TransactOpts, prevNextOutputIndex, newNextOutputIndex)
}

// SetNextVerifyIndex is a paid mutator transaction binding the contract method 0x84d67709.
//
// Solidity: function setNextVerifyIndex(uint256 val) returns()
func (_L2oo *L2ooTransactor) SetNextVerifyIndex(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _L2oo.contract.Transact(opts, "setNextVerifyIndex", val)
}

// SetNextVerifyIndex is a paid mutator transaction binding the contract method 0x84d67709.
//
// Solidity: function setNextVerifyIndex(uint256 val) returns()
func (_L2oo *L2ooSession) SetNextVerifyIndex(val *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.SetNextVerifyIndex(&_L2oo.TransactOpts, val)
}

// SetNextVerifyIndex is a paid mutator transaction binding the contract method 0x84d67709.
//
// Solidity: function setNextVerifyIndex(uint256 val) returns()
func (_L2oo *L2ooTransactorSession) SetNextVerifyIndex(val *big.Int) (*types.Transaction, error) {
	return _L2oo.Contract.SetNextVerifyIndex(&_L2oo.TransactOpts, val)
}

// L2ooOutputFailedIterator is returned from FilterOutputFailed and is used to iterate over the raw logs and unpacked data for OutputFailed events raised by the L2oo contract.
type L2ooOutputFailedIterator struct {
	Event *L2ooOutputFailed // Event containing the contract specifics and raw log

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
func (it *L2ooOutputFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ooOutputFailed)
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
		it.Event = new(L2ooOutputFailed)
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
func (it *L2ooOutputFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ooOutputFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ooOutputFailed represents a OutputFailed event raised by the L2oo contract.
type L2ooOutputFailed struct {
	L2OutputIndex *big.Int
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputFailed is a free log retrieval operation binding the contract event 0x39dab377091f907ce5cdb2ef4eeea3961f83ca13a3b6daff133ec7f99a3ec343.
//
// Solidity: event OutputFailed(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_L2oo *L2ooFilterer) FilterOutputFailed(opts *bind.FilterOpts, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (*L2ooOutputFailedIterator, error) {

	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _L2oo.contract.FilterLogs(opts, "OutputFailed", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &L2ooOutputFailedIterator{contract: _L2oo.contract, event: "OutputFailed", logs: logs, sub: sub}, nil
}

// WatchOutputFailed is a free log subscription operation binding the contract event 0x39dab377091f907ce5cdb2ef4eeea3961f83ca13a3b6daff133ec7f99a3ec343.
//
// Solidity: event OutputFailed(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_L2oo *L2ooFilterer) WatchOutputFailed(opts *bind.WatchOpts, sink chan<- *L2ooOutputFailed, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _L2oo.contract.WatchLogs(opts, "OutputFailed", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ooOutputFailed)
				if err := _L2oo.contract.UnpackLog(event, "OutputFailed", log); err != nil {
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

// ParseOutputFailed is a log parse operation binding the contract event 0x39dab377091f907ce5cdb2ef4eeea3961f83ca13a3b6daff133ec7f99a3ec343.
//
// Solidity: event OutputFailed(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_L2oo *L2ooFilterer) ParseOutputFailed(log types.Log) (*L2ooOutputFailed, error) {
	event := new(L2ooOutputFailed)
	if err := _L2oo.contract.UnpackLog(event, "OutputFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ooOutputProposedIterator is returned from FilterOutputProposed and is used to iterate over the raw logs and unpacked data for OutputProposed events raised by the L2oo contract.
type L2ooOutputProposedIterator struct {
	Event *L2ooOutputProposed // Event containing the contract specifics and raw log

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
func (it *L2ooOutputProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ooOutputProposed)
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
		it.Event = new(L2ooOutputProposed)
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
func (it *L2ooOutputProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ooOutputProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ooOutputProposed represents a OutputProposed event raised by the L2oo contract.
type L2ooOutputProposed struct {
	OutputRoot    [32]byte
	L2OutputIndex *big.Int
	L2BlockNumber *big.Int
	L1Timestamp   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputProposed is a free log retrieval operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_L2oo *L2ooFilterer) FilterOutputProposed(opts *bind.FilterOpts, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (*L2ooOutputProposedIterator, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _L2oo.contract.FilterLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &L2ooOutputProposedIterator{contract: _L2oo.contract, event: "OutputProposed", logs: logs, sub: sub}, nil
}

// WatchOutputProposed is a free log subscription operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_L2oo *L2ooFilterer) WatchOutputProposed(opts *bind.WatchOpts, sink chan<- *L2ooOutputProposed, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _L2oo.contract.WatchLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ooOutputProposed)
				if err := _L2oo.contract.UnpackLog(event, "OutputProposed", log); err != nil {
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

// ParseOutputProposed is a log parse operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_L2oo *L2ooFilterer) ParseOutputProposed(log types.Log) (*L2ooOutputProposed, error) {
	event := new(L2ooOutputProposed)
	if err := _L2oo.contract.UnpackLog(event, "OutputProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ooOutputVerifiedIterator is returned from FilterOutputVerified and is used to iterate over the raw logs and unpacked data for OutputVerified events raised by the L2oo contract.
type L2ooOutputVerifiedIterator struct {
	Event *L2ooOutputVerified // Event containing the contract specifics and raw log

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
func (it *L2ooOutputVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ooOutputVerified)
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
		it.Event = new(L2ooOutputVerified)
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
func (it *L2ooOutputVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ooOutputVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ooOutputVerified represents a OutputVerified event raised by the L2oo contract.
type L2ooOutputVerified struct {
	L2OutputIndex *big.Int
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputVerified is a free log retrieval operation binding the contract event 0x98beebb56f3d9f23ca4c2e85cac607b80e030b3ca74fcc7b38f98bdd5146ef6b.
//
// Solidity: event OutputVerified(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_L2oo *L2ooFilterer) FilterOutputVerified(opts *bind.FilterOpts, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (*L2ooOutputVerifiedIterator, error) {

	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _L2oo.contract.FilterLogs(opts, "OutputVerified", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &L2ooOutputVerifiedIterator{contract: _L2oo.contract, event: "OutputVerified", logs: logs, sub: sub}, nil
}

// WatchOutputVerified is a free log subscription operation binding the contract event 0x98beebb56f3d9f23ca4c2e85cac607b80e030b3ca74fcc7b38f98bdd5146ef6b.
//
// Solidity: event OutputVerified(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_L2oo *L2ooFilterer) WatchOutputVerified(opts *bind.WatchOpts, sink chan<- *L2ooOutputVerified, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _L2oo.contract.WatchLogs(opts, "OutputVerified", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ooOutputVerified)
				if err := _L2oo.contract.UnpackLog(event, "OutputVerified", log); err != nil {
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

// ParseOutputVerified is a log parse operation binding the contract event 0x98beebb56f3d9f23ca4c2e85cac607b80e030b3ca74fcc7b38f98bdd5146ef6b.
//
// Solidity: event OutputVerified(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_L2oo *L2ooFilterer) ParseOutputVerified(log types.Log) (*L2ooOutputVerified, error) {
	event := new(L2ooOutputVerified)
	if err := _L2oo.contract.UnpackLog(event, "OutputVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ooOutputsDeletedIterator is returned from FilterOutputsDeleted and is used to iterate over the raw logs and unpacked data for OutputsDeleted events raised by the L2oo contract.
type L2ooOutputsDeletedIterator struct {
	Event *L2ooOutputsDeleted // Event containing the contract specifics and raw log

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
func (it *L2ooOutputsDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ooOutputsDeleted)
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
		it.Event = new(L2ooOutputsDeleted)
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
func (it *L2ooOutputsDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ooOutputsDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ooOutputsDeleted represents a OutputsDeleted event raised by the L2oo contract.
type L2ooOutputsDeleted struct {
	PrevNextOutputIndex *big.Int
	NewNextOutputIndex  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterOutputsDeleted is a free log retrieval operation binding the contract event 0x4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b6.
//
// Solidity: event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex)
func (_L2oo *L2ooFilterer) FilterOutputsDeleted(opts *bind.FilterOpts, prevNextOutputIndex []*big.Int, newNextOutputIndex []*big.Int) (*L2ooOutputsDeletedIterator, error) {

	var prevNextOutputIndexRule []interface{}
	for _, prevNextOutputIndexItem := range prevNextOutputIndex {
		prevNextOutputIndexRule = append(prevNextOutputIndexRule, prevNextOutputIndexItem)
	}
	var newNextOutputIndexRule []interface{}
	for _, newNextOutputIndexItem := range newNextOutputIndex {
		newNextOutputIndexRule = append(newNextOutputIndexRule, newNextOutputIndexItem)
	}

	logs, sub, err := _L2oo.contract.FilterLogs(opts, "OutputsDeleted", prevNextOutputIndexRule, newNextOutputIndexRule)
	if err != nil {
		return nil, err
	}
	return &L2ooOutputsDeletedIterator{contract: _L2oo.contract, event: "OutputsDeleted", logs: logs, sub: sub}, nil
}

// WatchOutputsDeleted is a free log subscription operation binding the contract event 0x4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b6.
//
// Solidity: event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex)
func (_L2oo *L2ooFilterer) WatchOutputsDeleted(opts *bind.WatchOpts, sink chan<- *L2ooOutputsDeleted, prevNextOutputIndex []*big.Int, newNextOutputIndex []*big.Int) (event.Subscription, error) {

	var prevNextOutputIndexRule []interface{}
	for _, prevNextOutputIndexItem := range prevNextOutputIndex {
		prevNextOutputIndexRule = append(prevNextOutputIndexRule, prevNextOutputIndexItem)
	}
	var newNextOutputIndexRule []interface{}
	for _, newNextOutputIndexItem := range newNextOutputIndex {
		newNextOutputIndexRule = append(newNextOutputIndexRule, newNextOutputIndexItem)
	}

	logs, sub, err := _L2oo.contract.WatchLogs(opts, "OutputsDeleted", prevNextOutputIndexRule, newNextOutputIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ooOutputsDeleted)
				if err := _L2oo.contract.UnpackLog(event, "OutputsDeleted", log); err != nil {
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

// ParseOutputsDeleted is a log parse operation binding the contract event 0x4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b6.
//
// Solidity: event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex)
func (_L2oo *L2ooFilterer) ParseOutputsDeleted(log types.Log) (*L2ooOutputsDeleted, error) {
	event := new(L2ooOutputsDeleted)
	if err := _L2oo.contract.UnpackLog(event, "OutputsDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
