// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sccverifier

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

// OasysStateCommitmentChainVerifierChainBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type OasysStateCommitmentChainVerifierChainBatchHeader struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	ExtraData         []byte
}

// SccverifierMetaData contains all meta data concerning the Sccverifier contract.
var SccverifierMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assertLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"approve\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sccAssertLogsLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b50610017565b6110fb806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100515760003560e01c8063b32213bf14610057578063c4c5b2ff14610075578063d17f1fb1146100a8578063dea2619d146100c457610051565b60006000fd5b61005f6100e0565b60405161006c9190610d48565b60405180910390f35b61008f600480360381019061008a9190610ae3565b6100f5565b60405161009f9493929190610cd8565b60405180910390f35b6100c260048036038101906100bd9190610a59565b6102d2565b005b6100de60048036038101906100d99190610a59565b610519565b005b600060006000508054905090506100f2565b90565b6000600050818154811061010857600080fd5b906000526020600020906008020160005b915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000506040518060a00160405290816000820160005054815260200160018201600050546000191660001916815260200160028201600050548152602001600382016000505481526020016004820160005080546101a490610f1d565b80601f01602080910402602001604051908101604052809291908181526020018280546101d090610f1d565b801561021d5780601f106101f25761010080835404028352916020019161021d565b820191906000526020600020905b81548152906001019060200180831161020057829003601f168201915b5050505050815260200150509080600601600050805461023c90610f1d565b80601f016020809104026020016040519081016040528092919081815260200182805461026890610f1d565b80156102b55780601f1061028a576101008083540402835291602001916102b5565b820191906000526020600020905b81548152906001019060200180831161029857829003601f168201915b5050505050908060070160009054906101000a900460ff16905084565b600060005060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001610314858561076063ffffffff16565b815260200160001515815260200150908060018154018082558091505060019003906000526020600020906008020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001016000506000820151816000016000509090556020820151816001016000509060001916905560408201518160020160005090905560608201518160030160005090905560808201518160040160005090805190602001906103ff929190610803565b5050506040820151816006016000509080519060200190610421929190610803565b5060608201518160070160006101000a81548160ff02191690831515021790555050508373ffffffffffffffffffffffffffffffffffffffff1663982bc5b0846000015185602001516040518363ffffffff1660e01b8152600401610487929190610d64565b600060405180830381600087803b1580156104a25760006000fd5b505af11580156104b7573d600060003e3d6000fd5b5050505082600001518473ffffffffffffffffffffffffffffffffffffffff167f2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd856020015160405161050a9190610d2c565b60405180910390a35b50505050565b600060005060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200161055b858561076063ffffffff16565b815260200160011515815260200150908060018154018082558091505060019003906000526020600020906008020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101600050600082015181600001600050909055602082015181600101600050906000191690556040820151816002016000509090556060820151816003016000509090556080820151816004016000509080519060200190610646929190610803565b5050506040820151816006016000509080519060200190610668929190610803565b5060608201518160070160006101000a81548160ff02191690831515021790555050508373ffffffffffffffffffffffffffffffffffffffff1663d1f93ca6846000015185602001516040518363ffffffff1660e01b81526004016106ce929190610d64565b600060405180830381600087803b1580156106e95760006000fd5b505af11580156106fe573d600060003e3d6000fd5b5050505082600001518473ffffffffffffffffffffffffffffffffffffffff167f4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f585602001516040516107519190610d2c565b60405180910390a35b50505050565b60606000600090505b838390508110156107fb578184848381811015156107b0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020028101906107c29190610d8e565b6040516020016107d493929190610cb1565b604051602081830303815290604052915081505b80806107f390610f84565b915050610769565b505b92915050565b82805461080f90610f1d565b90600052602060002090601f016020900481019282610831576000855561087d565b82601f1061084a57805160ff191683800117855561087d565b8280016001018555821561087d579182015b8281111561087c578251826000509090559160200191906001019061085c565b5b50905061088a919061088e565b5090565b610893565b808211156108ad5760008181506000905550600101610893565b5090566110c4565b60006108c86108c384610e13565b610dec565b9050828152602081018484840111156108e15760006000fd5b6108ec848285610ed8565b505b9392505050565b60008135905061090481611073565b5b92915050565b6000600083601f84011215156109215760006000fd5b8235905067ffffffffffffffff81111561093b5760006000fd5b6020830191508360208202830111156109545760006000fd5b5b9250929050565b60008135905061096b8161108e565b5b92915050565b600082601f83011215156109865760006000fd5b81356109968482602086016108b5565b9150505b92915050565b600060a082840312156109b35760006000fd5b6109bd60a0610dec565b905060006109cd84828501610a43565b60008301525060206109e18482850161095c565b60208301525060406109f584828501610a43565b6040830152506060610a0984828501610a43565b606083015250608082013567ffffffffffffffff811115610a2a5760006000fd5b610a3684828501610972565b6080830152505b92915050565b600081359050610a52816110a9565b5b92915050565b600060006000600060608587031215610a725760006000fd5b6000610a80878288016108f5565b945050602085013567ffffffffffffffff811115610a9e5760006000fd5b610aaa878288016109a0565b935050604085013567ffffffffffffffff811115610ac85760006000fd5b610ad48782880161090b565b92509250505b92959194509250565b600060208284031215610af65760006000fd5b6000610b0484828501610a43565b9150505b92915050565b610b1781610e81565b82525b5050565b610b2781610e94565b82525b5050565b610b3781610ea1565b82525b5050565b610b4781610ea1565b82525b5050565b6000610b5a8385610e75565b9350610b67838584610ed8565b82840190505b9392505050565b6000610b7f82610e45565b610b898185610e51565b9350610b99818560208601610ee8565b610ba281611061565b84019150505b92915050565b6000610bb982610e45565b610bc38185610e63565b9350610bd3818560208601610ee8565b610bdc81611061565b84019150505b92915050565b6000610bf382610e45565b610bfd8185610e75565b9350610c0d818560208601610ee8565b8084019150505b92915050565b600060a083016000830151610c326000860182610c91565b506020830151610c456020860182610b2e565b506040830151610c586040860182610c91565b506060830151610c6b6060860182610c91565b5060808301518482036080860152610c838282610b74565b915050809150505b92915050565b610c9a81610ecd565b82525b5050565b610caa81610ecd565b82525b5050565b6000610cbd8286610be8565b9150610cca828486610b4e565b91508190505b949350505050565b6000608082019050610ced6000830187610b0e565b8181036020830152610cff8186610c1a565b90508181036040830152610d138185610bae565b9050610d226060830184610b1e565b5b95945050505050565b6000602082019050610d416000830184610b3e565b5b92915050565b6000602082019050610d5d6000830184610ca1565b5b92915050565b6000604082019050610d796000830185610ca1565b610d866020830184610b3e565b5b9392505050565b60006000833560016020038436030381121515610dab5760006000fd5b80840192508235915067ffffffffffffffff821115610dca5760006000fd5b602083019250600182023603831315610de35760006000fd5b505b9250929050565b6000610df6610e08565b9050610e028282610f52565b5b919050565b600060405190505b90565b600067ffffffffffffffff821115610e2e57610e2d611030565b5b610e3782611061565b90506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b6000610e8c82610eac565b90505b919050565b600081151590505b919050565b60008190505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b828183376000838301525b505050565b60005b83811015610f075780820151818401525b602081019050610eeb565b83811115610f16576000848401525b505b505050565b600060028204905060018216801515610f3757607f821691505b60208210811415610f4b57610f4a610fff565b5b505b919050565b610f5b82611061565b810181811067ffffffffffffffff82111715610f7a57610f79611030565b5b80604052505b5050565b6000610f8f82610ecd565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610fc257610fc1610fce565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b61107c81610e81565b8114151561108a5760006000fd5b5b50565b61109781610ea1565b811415156110a55760006000fd5b5b50565b6110b281610ecd565b811415156110c05760006000fd5b5b50565bfea26469706673582212204d44894fd5041eefdb372bfc04685a2f32bf916c3508e4e7b972828adafef7e964736f6c63430008020033",
}

// SccverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use SccverifierMetaData.ABI instead.
var SccverifierABI = SccverifierMetaData.ABI

// SccverifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SccverifierMetaData.Bin instead.
var SccverifierBin = SccverifierMetaData.Bin

// DeploySccverifier deploys a new Ethereum contract, binding an instance of Sccverifier to it.
func DeploySccverifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sccverifier, error) {
	parsed, err := SccverifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SccverifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sccverifier{SccverifierCaller: SccverifierCaller{contract: contract}, SccverifierTransactor: SccverifierTransactor{contract: contract}, SccverifierFilterer: SccverifierFilterer{contract: contract}}, nil
}

// Sccverifier is an auto generated Go binding around an Ethereum contract.
type Sccverifier struct {
	SccverifierCaller     // Read-only binding to the contract
	SccverifierTransactor // Write-only binding to the contract
	SccverifierFilterer   // Log filterer for contract events
}

// SccverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type SccverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SccverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SccverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SccverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SccverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SccverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SccverifierSession struct {
	Contract     *Sccverifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SccverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SccverifierCallerSession struct {
	Contract *SccverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SccverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SccverifierTransactorSession struct {
	Contract     *SccverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SccverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type SccverifierRaw struct {
	Contract *Sccverifier // Generic contract binding to access the raw methods on
}

// SccverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SccverifierCallerRaw struct {
	Contract *SccverifierCaller // Generic read-only contract binding to access the raw methods on
}

// SccverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SccverifierTransactorRaw struct {
	Contract *SccverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSccverifier creates a new instance of Sccverifier, bound to a specific deployed contract.
func NewSccverifier(address common.Address, backend bind.ContractBackend) (*Sccverifier, error) {
	contract, err := bindSccverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sccverifier{SccverifierCaller: SccverifierCaller{contract: contract}, SccverifierTransactor: SccverifierTransactor{contract: contract}, SccverifierFilterer: SccverifierFilterer{contract: contract}}, nil
}

// NewSccverifierCaller creates a new read-only instance of Sccverifier, bound to a specific deployed contract.
func NewSccverifierCaller(address common.Address, caller bind.ContractCaller) (*SccverifierCaller, error) {
	contract, err := bindSccverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SccverifierCaller{contract: contract}, nil
}

// NewSccverifierTransactor creates a new write-only instance of Sccverifier, bound to a specific deployed contract.
func NewSccverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*SccverifierTransactor, error) {
	contract, err := bindSccverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SccverifierTransactor{contract: contract}, nil
}

// NewSccverifierFilterer creates a new log filterer instance of Sccverifier, bound to a specific deployed contract.
func NewSccverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*SccverifierFilterer, error) {
	contract, err := bindSccverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SccverifierFilterer{contract: contract}, nil
}

// bindSccverifier binds a generic wrapper to an already deployed contract.
func bindSccverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SccverifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sccverifier *SccverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sccverifier.Contract.SccverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sccverifier *SccverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sccverifier.Contract.SccverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sccverifier *SccverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sccverifier.Contract.SccverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sccverifier *SccverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sccverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sccverifier *SccverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sccverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sccverifier *SccverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sccverifier.Contract.contract.Transact(opts, method, params...)
}

// AssertLogs is a free data retrieval call binding the contract method 0xc4c5b2ff.
//
// Solidity: function assertLogs(uint256 ) view returns(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes signatures, bool approve)
func (_Sccverifier *SccverifierCaller) AssertLogs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StateCommitmentChain common.Address
	BatchHeader          OasysStateCommitmentChainVerifierChainBatchHeader
	Signatures           []byte
	Approve              bool
}, error) {
	var out []interface{}
	err := _Sccverifier.contract.Call(opts, &out, "assertLogs", arg0)

	outstruct := new(struct {
		StateCommitmentChain common.Address
		BatchHeader          OasysStateCommitmentChainVerifierChainBatchHeader
		Signatures           []byte
		Approve              bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StateCommitmentChain = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BatchHeader = *abi.ConvertType(out[1], new(OasysStateCommitmentChainVerifierChainBatchHeader)).(*OasysStateCommitmentChainVerifierChainBatchHeader)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Approve = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// AssertLogs is a free data retrieval call binding the contract method 0xc4c5b2ff.
//
// Solidity: function assertLogs(uint256 ) view returns(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes signatures, bool approve)
func (_Sccverifier *SccverifierSession) AssertLogs(arg0 *big.Int) (struct {
	StateCommitmentChain common.Address
	BatchHeader          OasysStateCommitmentChainVerifierChainBatchHeader
	Signatures           []byte
	Approve              bool
}, error) {
	return _Sccverifier.Contract.AssertLogs(&_Sccverifier.CallOpts, arg0)
}

// AssertLogs is a free data retrieval call binding the contract method 0xc4c5b2ff.
//
// Solidity: function assertLogs(uint256 ) view returns(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes signatures, bool approve)
func (_Sccverifier *SccverifierCallerSession) AssertLogs(arg0 *big.Int) (struct {
	StateCommitmentChain common.Address
	BatchHeader          OasysStateCommitmentChainVerifierChainBatchHeader
	Signatures           []byte
	Approve              bool
}, error) {
	return _Sccverifier.Contract.AssertLogs(&_Sccverifier.CallOpts, arg0)
}

// SccAssertLogsLen is a free data retrieval call binding the contract method 0xb32213bf.
//
// Solidity: function sccAssertLogsLen() view returns(uint256)
func (_Sccverifier *SccverifierCaller) SccAssertLogsLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sccverifier.contract.Call(opts, &out, "sccAssertLogsLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SccAssertLogsLen is a free data retrieval call binding the contract method 0xb32213bf.
//
// Solidity: function sccAssertLogsLen() view returns(uint256)
func (_Sccverifier *SccverifierSession) SccAssertLogsLen() (*big.Int, error) {
	return _Sccverifier.Contract.SccAssertLogsLen(&_Sccverifier.CallOpts)
}

// SccAssertLogsLen is a free data retrieval call binding the contract method 0xb32213bf.
//
// Solidity: function sccAssertLogsLen() view returns(uint256)
func (_Sccverifier *SccverifierCallerSession) SccAssertLogsLen() (*big.Int, error) {
	return _Sccverifier.Contract.SccAssertLogsLen(&_Sccverifier.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactor) Approve(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.contract.Transact(opts, "approve", stateCommitmentChain, batchHeader, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierSession) Approve(stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Approve(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactorSession) Approve(stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Approve(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactor) Reject(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.contract.Transact(opts, "reject", stateCommitmentChain, batchHeader, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierSession) Reject(stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Reject(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactorSession) Reject(stateCommitmentChain common.Address, batchHeader OasysStateCommitmentChainVerifierChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Reject(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// SccverifierStateBatchApprovedIterator is returned from FilterStateBatchApproved and is used to iterate over the raw logs and unpacked data for StateBatchApproved events raised by the Sccverifier contract.
type SccverifierStateBatchApprovedIterator struct {
	Event *SccverifierStateBatchApproved // Event containing the contract specifics and raw log

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
func (it *SccverifierStateBatchApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccverifierStateBatchApproved)
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
		it.Event = new(SccverifierStateBatchApproved)
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
func (it *SccverifierStateBatchApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccverifierStateBatchApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccverifierStateBatchApproved represents a StateBatchApproved event raised by the Sccverifier contract.
type SccverifierStateBatchApproved struct {
	StateCommitmentChain common.Address
	BatchIndex           *big.Int
	BatchRoot            [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStateBatchApproved is a free log retrieval operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) FilterStateBatchApproved(opts *bind.FilterOpts, stateCommitmentChain []common.Address, batchIndex []*big.Int) (*SccverifierStateBatchApprovedIterator, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Sccverifier.contract.FilterLogs(opts, "StateBatchApproved", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &SccverifierStateBatchApprovedIterator{contract: _Sccverifier.contract, event: "StateBatchApproved", logs: logs, sub: sub}, nil
}

// WatchStateBatchApproved is a free log subscription operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) WatchStateBatchApproved(opts *bind.WatchOpts, sink chan<- *SccverifierStateBatchApproved, stateCommitmentChain []common.Address, batchIndex []*big.Int) (event.Subscription, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Sccverifier.contract.WatchLogs(opts, "StateBatchApproved", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccverifierStateBatchApproved)
				if err := _Sccverifier.contract.UnpackLog(event, "StateBatchApproved", log); err != nil {
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

// ParseStateBatchApproved is a log parse operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) ParseStateBatchApproved(log types.Log) (*SccverifierStateBatchApproved, error) {
	event := new(SccverifierStateBatchApproved)
	if err := _Sccverifier.contract.UnpackLog(event, "StateBatchApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SccverifierStateBatchRejectedIterator is returned from FilterStateBatchRejected and is used to iterate over the raw logs and unpacked data for StateBatchRejected events raised by the Sccverifier contract.
type SccverifierStateBatchRejectedIterator struct {
	Event *SccverifierStateBatchRejected // Event containing the contract specifics and raw log

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
func (it *SccverifierStateBatchRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SccverifierStateBatchRejected)
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
		it.Event = new(SccverifierStateBatchRejected)
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
func (it *SccverifierStateBatchRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SccverifierStateBatchRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SccverifierStateBatchRejected represents a StateBatchRejected event raised by the Sccverifier contract.
type SccverifierStateBatchRejected struct {
	StateCommitmentChain common.Address
	BatchIndex           *big.Int
	BatchRoot            [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStateBatchRejected is a free log retrieval operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) FilterStateBatchRejected(opts *bind.FilterOpts, stateCommitmentChain []common.Address, batchIndex []*big.Int) (*SccverifierStateBatchRejectedIterator, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Sccverifier.contract.FilterLogs(opts, "StateBatchRejected", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &SccverifierStateBatchRejectedIterator{contract: _Sccverifier.contract, event: "StateBatchRejected", logs: logs, sub: sub}, nil
}

// WatchStateBatchRejected is a free log subscription operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) WatchStateBatchRejected(opts *bind.WatchOpts, sink chan<- *SccverifierStateBatchRejected, stateCommitmentChain []common.Address, batchIndex []*big.Int) (event.Subscription, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Sccverifier.contract.WatchLogs(opts, "StateBatchRejected", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SccverifierStateBatchRejected)
				if err := _Sccverifier.contract.UnpackLog(event, "StateBatchRejected", log); err != nil {
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

// ParseStateBatchRejected is a log parse operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_Sccverifier *SccverifierFilterer) ParseStateBatchRejected(log types.Log) (*SccverifierStateBatchRejected, error) {
	event := new(SccverifierStateBatchRejected)
	if err := _Sccverifier.contract.UnpackLog(event, "StateBatchRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
