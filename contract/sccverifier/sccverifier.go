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

// Lib_OVMCodecChainBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type Lib_OVMCodecChainBatchHeader struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	ExtraData         []byte
}

// SccverifierMetaData contains all meta data concerning the Sccverifier contract.
var SccverifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"InvalidAddressSort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"OutdatedValidatorAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verified\",\"type\":\"uint256\"}],\"name\":\"StakeAmountShortage\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_OVMCodec.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_OVMCodec.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"recoverSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_OVMCodec.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611af5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806345c3edd814610046578063d17f1fb114610076578063dea2619d14610092575b600080fd5b610060600480360381019061005b9190610ee0565b6100ae565b60405161006d919061103d565b60405180910390f35b610090600480360381019061008b919061105f565b6100fa565b005b6100ac60048036038101906100a7919061105f565b6101ce565b005b60606100f0468686600001518760200151876040516020016100d49594939291906111bc565b60405160208183030381529060405280519060200120836102a2565b9050949350505050565b61010783836000846105d5565b8273ffffffffffffffffffffffffffffffffffffffff16639594dd64836040518263ffffffff1660e01b81526004016101409190611337565b600060405180830381600087803b15801561015a57600080fd5b505af115801561016e573d6000803e3d6000fd5b5050505081600001518373ffffffffffffffffffffffffffffffffffffffff167f2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd84602001516040516101c19190611368565b60405180910390a3505050565b6101db83836001846105d5565b8273ffffffffffffffffffffffffffffffffffffffff1663e0c2b418836040518263ffffffff1660e01b81526004016102149190611337565b600060405180830381600087803b15801561022e57600080fd5b505af1158015610242573d6000803e3d6000fd5b5050505081600001518373ffffffffffffffffffffffffffffffffffffffff167f4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f584602001516040516102959190611368565b60405180910390a3505050565b6060815167ffffffffffffffff8111156102bf576102be610b74565b5b6040519080825280602002602001820160405280156102ed5781602001602082028036833780820191505090505b50905060008360405160200161030391906113da565b6040516020818303038152906040528051906020012090506000805b84518110156105cc57600085828151811061033d5761033c611400565b5b602002602001015190506000806103548684610629565b915091506001600481111561036c5761036b61142f565b5b81600481111561037f5761037e61142f565b5b14156103c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b7906114bb565b60405180910390fd5b600260048111156103d4576103d361142f565b5b8160048111156103e7576103e661142f565b5b1415610428576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041f90611527565b60405180910390fd5b6003600481111561043c5761043b61142f565b5b81600481111561044f5761044e61142f565b5b1415610490576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610487906115b9565b60405180910390fd5b6004808111156104a3576104a261142f565b5b8160048111156104b6576104b561142f565b5b14156104f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ee9061164b565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1611610565576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055c906116b7565b60405180910390fd5b8187858151811061057957610578611400565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505081945050505080806105c490611706565b91505061031f565b50505092915050565b6000610617468686600001518760200151876040516020016105fb9594939291906111bc565b60405160208183030381529060405280519060200120836102a2565b9050610622816106ac565b5050505050565b60008060418351141561066b5760008060006020860151925060408601519150606086015160001a905061065f87828585610991565b945094505050506106a5565b60408351141561069c576000806020850151915060408501519050610691868383610a9e565b9350935050506106a5565b60006002915091505b9250929050565b600061100190506000825190506000805b8281101561089c5760008582815181106106da576106d9611400565b5b6020026020010151905060008573ffffffffffffffffffffffffffffffffffffffff16637b520aa8836040518263ffffffff1660e01b815260040161071f919061175e565b60206040518083038186803b15801561073757600080fd5b505afa15801561074b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076f919061178e565b90506000808773ffffffffffffffffffffffffffffffffffffffff1663d1f18ee18460006040518363ffffffff1660e01b81526004016107b0929190611800565b60a06040518083038186803b1580156107c857600080fd5b505afa1580156107dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108009190611853565b945050505091508173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161461087757826040517f5b73c50400000000000000000000000000000000000000000000000000000000815260040161086e919061175e565b60405180910390fd5b808661088391906118ce565b955050505050808061089490611706565b9150506106bd565b506000603360648573ffffffffffffffffffffffffffffffffffffffff166345367f2360006040518263ffffffff1660e01b81526004016108dd9190611924565b60206040518083038186803b1580156108f557600080fd5b505afa158015610909573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092d919061193f565b610937919061199b565b61094191906119cc565b90508082101561098a5780826040517f09b3dd4a000000000000000000000000000000000000000000000000000000008152600401610981929190611a35565b60405180910390fd5b5050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c11156109cc576000600391509150610a95565b601b8560ff16141580156109e45750601c8560ff1614155b156109f6576000600491509150610a95565b600060018787878760405160008152602001604052604051610a1b9493929190611a7a565b6020604051602081039080840390855afa158015610a3d573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a8c57600060019250925050610a95565b80600092509250505b94509492505050565b6000806000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85169150601b8560ff1c019050610ade87828885610991565b935093505050935093915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b2b82610b00565b9050919050565b610b3b81610b20565b8114610b4657600080fd5b50565b600081359050610b5881610b32565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610bac82610b63565b810181811067ffffffffffffffff82111715610bcb57610bca610b74565b5b80604052505050565b6000610bde610aec565b9050610bea8282610ba3565b919050565b600080fd5b6000819050919050565b610c0781610bf4565b8114610c1257600080fd5b50565b600081359050610c2481610bfe565b92915050565b6000819050919050565b610c3d81610c2a565b8114610c4857600080fd5b50565b600081359050610c5a81610c34565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff821115610c8557610c84610b74565b5b610c8e82610b63565b9050602081019050919050565b82818337600083830152505050565b6000610cbd610cb884610c6a565b610bd4565b905082815260208101848484011115610cd957610cd8610c65565b5b610ce4848285610c9b565b509392505050565b600082601f830112610d0157610d00610c60565b5b8135610d11848260208601610caa565b91505092915050565b600060a08284031215610d3057610d2f610b5e565b5b610d3a60a0610bd4565b90506000610d4a84828501610c15565b6000830152506020610d5e84828501610c4b565b6020830152506040610d7284828501610c15565b6040830152506060610d8684828501610c15565b606083015250608082013567ffffffffffffffff811115610daa57610da9610bef565b5b610db684828501610cec565b60808301525092915050565b60008115159050919050565b610dd781610dc2565b8114610de257600080fd5b50565b600081359050610df481610dce565b92915050565b600067ffffffffffffffff821115610e1557610e14610b74565b5b602082029050602081019050919050565b600080fd5b6000610e3e610e3984610dfa565b610bd4565b90508083825260208201905060208402830185811115610e6157610e60610e26565b5b835b81811015610ea857803567ffffffffffffffff811115610e8657610e85610c60565b5b808601610e938982610cec565b85526020850194505050602081019050610e63565b5050509392505050565b600082601f830112610ec757610ec6610c60565b5b8135610ed7848260208601610e2b565b91505092915050565b60008060008060808587031215610efa57610ef9610af6565b5b6000610f0887828801610b49565b945050602085013567ffffffffffffffff811115610f2957610f28610afb565b5b610f3587828801610d1a565b9350506040610f4687828801610de5565b925050606085013567ffffffffffffffff811115610f6757610f66610afb565b5b610f7387828801610eb2565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610fb481610b20565b82525050565b6000610fc68383610fab565b60208301905092915050565b6000602082019050919050565b6000610fea82610f7f565b610ff48185610f8a565b9350610fff83610f9b565b8060005b838110156110305781516110178882610fba565b975061102283610fd2565b925050600181019050611003565b5085935050505092915050565b600060208201905081810360008301526110578184610fdf565b905092915050565b60008060006060848603121561107857611077610af6565b5b600061108686828701610b49565b935050602084013567ffffffffffffffff8111156110a7576110a6610afb565b5b6110b386828701610d1a565b925050604084013567ffffffffffffffff8111156110d4576110d3610afb565b5b6110e086828701610eb2565b9150509250925092565b6000819050919050565b61110561110082610bf4565b6110ea565b82525050565b60008160601b9050919050565b60006111238261110b565b9050919050565b600061113582611118565b9050919050565b61114d61114882610b20565b61112a565b82525050565b6000819050919050565b61116e61116982610c2a565b611153565b82525050565b60008160f81b9050919050565b600061118c82611174565b9050919050565b600061119e82611181565b9050919050565b6111b66111b182610dc2565b611193565b82525050565b60006111c882886110f4565b6020820191506111d8828761113c565b6014820191506111e882866110f4565b6020820191506111f8828561115d565b60208201915061120882846111a5565b6001820191508190509695505050505050565b61122481610bf4565b82525050565b61123381610c2a565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611273578082015181840152602081019050611258565b83811115611282576000848401525b50505050565b600061129382611239565b61129d8185611244565b93506112ad818560208601611255565b6112b681610b63565b840191505092915050565b600060a0830160008301516112d9600086018261121b565b5060208301516112ec602086018261122a565b5060408301516112ff604086018261121b565b506060830151611312606086018261121b565b506080830151848203608086015261132a8282611288565b9150508091505092915050565b6000602082019050818103600083015261135181846112c1565b905092915050565b61136281610c2a565b82525050565b600060208201905061137d6000830184611359565b92915050565b600081905092915050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b60006113c4601c83611383565b91506113cf8261138e565b601c82019050919050565b60006113e5826113b7565b91506113f1828461115d565b60208201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600082825260208201905092915050565b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b60006114a560188361145e565b91506114b08261146f565b602082019050919050565b600060208201905081810360008301526114d481611498565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b6000611511601f8361145e565b915061151c826114db565b602082019050919050565b6000602082019050818103600083015261154081611504565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b60006115a360228361145e565b91506115ae82611547565b604082019050919050565b600060208201905081810360008301526115d281611596565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b600061163560228361145e565b9150611640826115d9565b604082019050919050565b6000602082019050818103600083015261166481611628565b9050919050565b7f696e76616c6964206164647265737320736f7274000000000000000000000000600082015250565b60006116a160148361145e565b91506116ac8261166b565b602082019050919050565b600060208201905081810360008301526116d081611694565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061171182610bf4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611744576117436116d7565b5b600182019050919050565b61175881610b20565b82525050565b6000602082019050611773600083018461174f565b92915050565b60008151905061178881610b32565b92915050565b6000602082840312156117a4576117a3610af6565b5b60006117b284828501611779565b91505092915050565b6000819050919050565b6000819050919050565b60006117ea6117e56117e0846117bb565b6117c5565b610bf4565b9050919050565b6117fa816117cf565b82525050565b6000604082019050611815600083018561174f565b61182260208301846117f1565b9392505050565b60008151905061183881610dce565b92915050565b60008151905061184d81610bfe565b92915050565b600080600080600060a0868803121561186f5761186e610af6565b5b600061187d88828901611779565b955050602061188e88828901611829565b945050604061189f88828901611829565b93505060606118b088828901611829565b92505060806118c18882890161183e565b9150509295509295909350565b60006118d982610bf4565b91506118e483610bf4565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611919576119186116d7565b5b828201905092915050565b600060208201905061193960008301846117f1565b92915050565b60006020828403121561195557611954610af6565b5b60006119638482850161183e565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006119a682610bf4565b91506119b183610bf4565b9250826119c1576119c061196c565b5b828204905092915050565b60006119d782610bf4565b91506119e283610bf4565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611a1b57611a1a6116d7565b5b828202905092915050565b611a2f81610bf4565b82525050565b6000604082019050611a4a6000830185611a26565b611a576020830184611a26565b9392505050565b600060ff82169050919050565b611a7481611a5e565b82525050565b6000608082019050611a8f6000830187611359565b611a9c6020830186611a6b565b611aa96040830185611359565b611ab66060830184611359565b9594505050505056fea2646970667358221220671939d61781827f6d1bac6cde08875551739d056ea616c0b6756696dad468a964736f6c63430008090033",
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

// RecoverSigners is a free data retrieval call binding the contract method 0x45c3edd8.
//
// Solidity: function recoverSigners(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bool approved, bytes[] signatures) view returns(address[] signers)
func (_Sccverifier *SccverifierCaller) RecoverSigners(opts *bind.CallOpts, stateCommitmentChain common.Address, batchHeader Lib_OVMCodecChainBatchHeader, approved bool, signatures [][]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Sccverifier.contract.Call(opts, &out, "recoverSigners", stateCommitmentChain, batchHeader, approved, signatures)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RecoverSigners is a free data retrieval call binding the contract method 0x45c3edd8.
//
// Solidity: function recoverSigners(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bool approved, bytes[] signatures) view returns(address[] signers)
func (_Sccverifier *SccverifierSession) RecoverSigners(stateCommitmentChain common.Address, batchHeader Lib_OVMCodecChainBatchHeader, approved bool, signatures [][]byte) ([]common.Address, error) {
	return _Sccverifier.Contract.RecoverSigners(&_Sccverifier.CallOpts, stateCommitmentChain, batchHeader, approved, signatures)
}

// RecoverSigners is a free data retrieval call binding the contract method 0x45c3edd8.
//
// Solidity: function recoverSigners(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bool approved, bytes[] signatures) view returns(address[] signers)
func (_Sccverifier *SccverifierCallerSession) RecoverSigners(stateCommitmentChain common.Address, batchHeader Lib_OVMCodecChainBatchHeader, approved bool, signatures [][]byte) ([]common.Address, error) {
	return _Sccverifier.Contract.RecoverSigners(&_Sccverifier.CallOpts, stateCommitmentChain, batchHeader, approved, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactor) Approve(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader Lib_OVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.contract.Transact(opts, "approve", stateCommitmentChain, batchHeader, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierSession) Approve(stateCommitmentChain common.Address, batchHeader Lib_OVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Approve(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactorSession) Approve(stateCommitmentChain common.Address, batchHeader Lib_OVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Approve(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactor) Reject(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader Lib_OVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.contract.Transact(opts, "reject", stateCommitmentChain, batchHeader, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierSession) Reject(stateCommitmentChain common.Address, batchHeader Lib_OVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _Sccverifier.Contract.Reject(&_Sccverifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_Sccverifier *SccverifierTransactorSession) Reject(stateCommitmentChain common.Address, batchHeader Lib_OVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
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
