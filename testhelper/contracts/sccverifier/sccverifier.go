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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assertLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"approve\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structOasysStateCommitmentChainVerifier.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b50610017565b6110db806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100465760003560e01c8063c4c5b2ff1461004c578063d17f1fb11461007f578063dea2619d1461009b57610046565b60006000fd5b61006660048036038101906100619190610b38565b6100b7565b6040516100769493929190610d05565b60405180910390f35b61009960048036038101906100949190610ab4565b610294565b005b6100b560048036038101906100b09190610ab4565b61055b565b005b600060005081815481106100ca57600080fd5b906000526020600020906008020160005b915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000506040518060a001604052908160008201600050548152602001600182016000505460001916600019168152602001600282016000505481526020016003820160005054815260200160048201600050805461016690610efd565b80601f016020809104026020016040519081016040528092919081815260200182805461019290610efd565b80156101df5780601f106101b4576101008083540402835291602001916101df565b820191906000526020600020905b8154815290600101906020018083116101c257829003601f168201915b505050505081526020015050908060060160005080546101fe90610efd565b80601f016020809104026020016040519081016040528092919081815260200182805461022a90610efd565b80156102775780601f1061024c57610100808354040283529160200191610277565b820191906000526020600020905b81548152906001019060200180831161025a57829003601f168201915b5050505050908060070160009054906101000a900460ff16905084565b60606000600090505b8251811015610322578183828151811015156102e2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040516020016102fb929190610ce0565b604051602081830303815290604052915081505b808061031a90610f64565b91505061029d565b50600060005060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200183815260200160001515815260200150908060018154018082558091505060019003906000526020600020906008020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101600050600082015181600001600050909055602082015181600101600050906000191690556040820151816002016000509090556060820151816003016000509090556080820151816004016000509080519060200190610441929190610822565b5050506040820151816006016000509080519060200190610463929190610822565b5060608201518160070160006101000a81548160ff02191690831515021790555050508373ffffffffffffffffffffffffffffffffffffffff1663982bc5b0846000015185602001516040518363ffffffff1660e01b81526004016104c9929190610d75565b600060405180830381600087803b1580156104e45760006000fd5b505af11580156104f9573d600060003e3d6000fd5b5050505082600001518473ffffffffffffffffffffffffffffffffffffffff167f2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd856020015160405161054c9190610d59565b60405180910390a3505b505050565b60606000600090505b82518110156105e9578183828151811015156105a9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040516020016105c2929190610ce0565b604051602081830303815290604052915081505b80806105e190610f64565b915050610564565b50600060005060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200183815260200160011515815260200150908060018154018082558091505060019003906000526020600020906008020160005b9091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101600050600082015181600001600050909055602082015181600101600050906000191690556040820151816002016000509090556060820151816003016000509090556080820151816004016000509080519060200190610708929190610822565b505050604082015181600601600050908051906020019061072a929190610822565b5060608201518160070160006101000a81548160ff02191690831515021790555050508373ffffffffffffffffffffffffffffffffffffffff1663d1f93ca6846000015185602001516040518363ffffffff1660e01b8152600401610790929190610d75565b600060405180830381600087803b1580156107ab5760006000fd5b505af11580156107c0573d600060003e3d6000fd5b5050505082600001518473ffffffffffffffffffffffffffffffffffffffff167f4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f585602001516040516108139190610d59565b60405180910390a3505b505050565b82805461082e90610efd565b90600052602060002090601f016020900481019282610850576000855561089c565b82601f1061086957805160ff191683800117855561089c565b8280016001018555821561089c579182015b8281111561089b578251826000509090559160200191906001019061087b565b5b5090506108a991906108ad565b5090565b6108b2565b808211156108cc57600081815060009055506001016108b2565b5090566110a4565b60006108e76108e284610dc6565b610d9f565b9050808382526020820190508260005b85811015610928578135850161090d88826109cd565b8452602084019350602083019250505b6001810190506108f7565b5050505b9392505050565b600061094661094184610df3565b610d9f565b90508281526020810184848401111561095f5760006000fd5b61096a848285610eb8565b505b9392505050565b60008135905061098281611053565b5b92915050565b600082601f830112151561099d5760006000fd5b81356109ad8482602086016108d4565b9150505b92915050565b6000813590506109c68161106e565b5b92915050565b600082601f83011215156109e15760006000fd5b81356109f1848260208601610933565b9150505b92915050565b600060a08284031215610a0e5760006000fd5b610a1860a0610d9f565b90506000610a2884828501610a9e565b6000830152506020610a3c848285016109b7565b6020830152506040610a5084828501610a9e565b6040830152506060610a6484828501610a9e565b606083015250608082013567ffffffffffffffff811115610a855760006000fd5b610a91848285016109cd565b6080830152505b92915050565b600081359050610aad81611089565b5b92915050565b60006000600060608486031215610acb5760006000fd5b6000610ad986828701610973565b935050602084013567ffffffffffffffff811115610af75760006000fd5b610b03868287016109fb565b925050604084013567ffffffffffffffff811115610b215760006000fd5b610b2d86828701610989565b9150505b9250925092565b600060208284031215610b4b5760006000fd5b6000610b5984828501610a9e565b9150505b92915050565b610b6c81610e61565b82525b5050565b610b7c81610e74565b82525b5050565b610b8c81610e81565b82525b5050565b610b9c81610e81565b82525b5050565b6000610bae82610e25565b610bb88185610e31565b9350610bc8818560208601610ec8565b610bd181611041565b84019150505b92915050565b6000610be882610e25565b610bf28185610e43565b9350610c02818560208601610ec8565b610c0b81611041565b84019150505b92915050565b6000610c2282610e25565b610c2c8185610e55565b9350610c3c818560208601610ec8565b8084019150505b92915050565b600060a083016000830151610c616000860182610cc0565b506020830151610c746020860182610b83565b506040830151610c876040860182610cc0565b506060830151610c9a6060860182610cc0565b5060808301518482036080860152610cb28282610ba3565b915050809150505b92915050565b610cc981610ead565b82525b5050565b610cd981610ead565b82525b5050565b6000610cec8285610c17565b9150610cf88284610c17565b91508190505b9392505050565b6000608082019050610d1a6000830187610b63565b8181036020830152610d2c8186610c49565b90508181036040830152610d408185610bdd565b9050610d4f6060830184610b73565b5b95945050505050565b6000602082019050610d6e6000830184610b93565b5b92915050565b6000604082019050610d8a6000830185610cd0565b610d976020830184610b93565b5b9392505050565b6000610da9610dbb565b9050610db58282610f32565b5b919050565b600060405190505b90565b600067ffffffffffffffff821115610de157610de0611010565b5b6020820290506020810190505b919050565b600067ffffffffffffffff821115610e0e57610e0d611010565b5b610e1782611041565b90506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b6000610e6c82610e8c565b90505b919050565b600081151590505b919050565b60008190505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b828183376000838301525b505050565b60005b83811015610ee75780820151818401525b602081019050610ecb565b83811115610ef6576000848401525b505b505050565b600060028204905060018216801515610f1757607f821691505b60208210811415610f2b57610f2a610fdf565b5b505b919050565b610f3b82611041565b810181811067ffffffffffffffff82111715610f5a57610f59611010565b5b80604052505b5050565b6000610f6f82610ead565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610fa257610fa1610fae565b5b6001820190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b565b6000601f19601f83011690505b919050565b61105c81610e61565b8114151561106a5760006000fd5b5b50565b61107781610e81565b811415156110855760006000fd5b5b50565b61109281610ead565b811415156110a05760006000fd5b5b50565bfea26469706673582212206e58f4579a6a3a1f92e90c42d31c4b71f563212fa6678f66f19b6df57a6c43fd64736f6c63430008020033",
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
	parsed, err := abi.JSON(strings.NewReader(SccverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
