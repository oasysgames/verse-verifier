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

// LibOVMCodecChainBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type LibOVMCodecChainBatchHeader struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	ExtraData         []byte
}

// TypesOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// OasysL2OutputOracleVerifierMetaData contains all meta data concerning the OasysL2OutputOracleVerifier contract.
var OasysL2OutputOracleVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"InvalidAddressSort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verified\",\"type\":\"uint256\"}],\"name\":\"StakeAmountShortage\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"L2OutputRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"StateBatchRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_OVMCodec.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateCommitmentChain\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"batchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevTotalElements\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structLib_OVMCodec.ChainBatchHeader\",\"name\":\"batchHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611659806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80639ac45572116100505780639ac45572146100d3578063d17f1fb1146100e6578063dea2619d146100f957600080fd5b806354fd4d501461006c5780637d2d9b63146100be575b600080fd5b6100a86040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100b59190610eb5565b60405180910390f35b6100d16100cc366004610ef8565b61010c565b005b6100d16100e1366004610ef8565b6101fe565b6100d16100f436600461119a565b6102e1565b6100d161010736600461119a565b61047c565b61012b61011c8686866001610627565b6101268385611264565b610784565b6040517fd882343600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86169063d88234369061017f9087908790600401611291565b600060405180830381600087803b15801561019957600080fd5b505af11580156101ad573d6000803e3d6000fd5b50506040518535925086915073ffffffffffffffffffffffffffffffffffffffff8816907f3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f90600090a45050505050565b61020e61011c8686866000610627565b6040517f859c029d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86169063859c029d906102629087908790600401611291565b600060405180830381600087803b15801561027c57600080fd5b505af1158015610290573d6000803e3d6000fd5b50506040518535925086915073ffffffffffffffffffffffffffffffffffffffff8816907ff7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca390600090a45050505050565b8151602080840151604080514681850152606088901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016818301526054810194909452607484019190915260006094840152805160758185030181526095840182528051908301207f19457468657265756d205369676e6564204d6573736167653a0a33320000000060b585015260d1808501919091528151808503909101815260f19093019052815191012061039a905b82610784565b6040517f9594dd6400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841690639594dd64906103ec9085906004016112e4565b600060405180830381600087803b15801561040657600080fd5b505af115801561041a573d6000803e3d6000fd5b5050505081600001518373ffffffffffffffffffffffffffffffffffffffff167f2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd846020015160405161046f91815260200190565b60405180910390a3505050565b8151602080840151604080514681850152606088901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001681830152605481019490945260748401919091527f01000000000000000000000000000000000000000000000000000000000000006094840152805160758185030181526095840182528051908301207f19457468657265756d205369676e6564204d6573736167653a0a33320000000060b585015260d1808501919091528151808503909101815260f19093019052815191012061055290610394565b6040517fe0c2b41800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063e0c2b418906105a49085906004016112e4565b600060405180830381600087803b1580156105be57600080fd5b505af11580156105d2573d6000803e3d6000fd5b5050505081600001518373ffffffffffffffffffffffffffffffffffffffff167f4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5846020015160405161046f91815260200190565b600080468686863561063f6040890160208a0161132f565b61064f60608a0160408b0161132f565b60405160200161069b93929190928352608091821b7fffffffffffffffffffffffffffffffff000000000000000000000000000000009081166020850152911b16603082015260400190565b604051602081830303815290604052805190602001208660405160200161070e95949392919094855260609390931b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602085015260348401919091526054830152151560f81b607482015260750190565b6040516020818303038152906040529050808051906020012060405160200161076391907f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0190565b60405160208183030381529060405280519060200120915050949350505050565b610796610791838361079a565b6108ec565b5050565b6060815167ffffffffffffffff8111156107b6576107b6610fc0565b6040519080825280602002602001820160405280156107df578160200160208202803683370190505b5090506000805b83518110156108e4576000610814868684815181106108075761080761134a565b6020026020010151610b4a565b90508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1611610898576040517f3d1fd1b000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b808483815181106108ab576108ab61134a565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101529150806108dc816113a8565b9150506107e6565b505092915050565b600061100090506000611001905060008273ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610947573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096b91906113e0565b84519091506000805b82811015610a50578473ffffffffffffffffffffffffffffffffffffffff16639c5082198883815181106109aa576109aa61134a565b6020026020010151866040518363ffffffff1660e01b81526004016109f192919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b602060405180830381865afa158015610a0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3291906113e0565b610a3c90836113f9565b915080610a48816113a8565b915050610974565b506040517f45367f230000000000000000000000000000000000000000000000000000000081526004810184905260009060649073ffffffffffffffffffffffffffffffffffffffff8716906345367f2390602401602060405180830381865afa158015610ac2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae691906113e0565b610af1906033611411565b610afb919061144e565b905080821015610b41576040517f09b3dd4a000000000000000000000000000000000000000000000000000000008152600481018290526024810183905260440161088f565b50505050505050565b6000806000610b598585610ced565b90925090506000816004811115610b7257610b72611489565b03610b7f57509050610ce7565b6001816004811115610b9357610b93611489565b03610bcc57836040517ff31b6ee500000000000000000000000000000000000000000000000000000000815260040161088f91906114b8565b6002816004811115610be057610be0611489565b03610c1957836040517ff31b6ee500000000000000000000000000000000000000000000000000000000815260040161088f919061150a565b6003816004811115610c2d57610c2d611489565b03610c6657836040517ff31b6ee500000000000000000000000000000000000000000000000000000000815260040161088f919061155c565b6004816004811115610c7a57610c7a611489565b03610cb357836040517ff31b6ee500000000000000000000000000000000000000000000000000000000815260040161088f91906115d4565b836040517ff31b6ee500000000000000000000000000000000000000000000000000000000815260040161088f91906114b8565b92915050565b6000808251604103610d235760208301516040840151606085015160001a610d1787828585610d32565b94509450505050610d2b565b506000905060025b9250929050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610d695750600090506003610e41565b8460ff16601b14158015610d8157508460ff16601c14155b15610d925750600090506004610e41565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610de6573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116610e3a57600060019250925050610e41565b9150600090505b94509492505050565b6000815180845260005b81811015610e7057602081850181015186830182015201610e54565b81811115610e82576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610ec86020830184610e4a565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610ef357600080fd5b919050565b600080600080600085870360c0811215610f1157600080fd5b610f1a87610ecf565b95506020870135945060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082011215610f5357600080fd5b5060408601925060a086013567ffffffffffffffff80821115610f7557600080fd5b818801915088601f830112610f8957600080fd5b813581811115610f9857600080fd5b8960208260051b8501011115610fad57600080fd5b9699959850939650602001949392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff8111828210171561101257611012610fc0565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561105f5761105f610fc0565b604052919050565b600082601f83011261107857600080fd5b813567ffffffffffffffff81111561109257611092610fc0565b6110c360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611018565b8181528460208386010111156110d857600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8084111561111057611110610fc0565b8360051b6020611121818301611018565b8681529350908401908084018783111561113a57600080fd5b855b8381101561116e578035858111156111545760008081fd5b6111608a828a01611067565b83525090820190820161113c565b50505050509392505050565b600082601f83011261118b57600080fd5b610ec8838335602085016110f5565b6000806000606084860312156111af57600080fd5b6111b884610ecf565b9250602084013567ffffffffffffffff808211156111d557600080fd5b9085019060a082880312156111e957600080fd5b6111f1610fef565b8235815260208301356020820152604083013560408201526060830135606082015260808301358281111561122557600080fd5b61123189828601611067565b6080830152509350604086013591508082111561124d57600080fd5b5061125a8682870161117a565b9150509250925092565b6000610ec83684846110f5565b80356fffffffffffffffffffffffffffffffff81168114610ef357600080fd5b6000608082019050838252823560208301526112af60208401611271565b6fffffffffffffffffffffffffffffffff8082166040850152806112d560408701611271565b16606085015250509392505050565b60208152815160208201526020820151604082015260408201516060820152606082015160808201526000608083015160a08084015261132760c0840182610e4a565b949350505050565b60006020828403121561134157600080fd5b610ec882611271565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036113d9576113d9611379565b5060010190565b6000602082840312156113f257600080fd5b5051919050565b6000821982111561140c5761140c611379565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561144957611449611379565b500290565b600082611484577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6040815260006114cb6040830184610e4a565b8281036020840152601881527f45434453413a20696e76616c6964207369676e6174757265000000000000000060208201526040810191505092915050565b60408152600061151d6040830184610e4a565b8281036020840152601f81527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060208201526040810191505092915050565b60408152600061156f6040830184610e4a565b8281036020840152602281527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60208201527f756500000000000000000000000000000000000000000000000000000000000060408201526060810191505092915050565b6040815260006115e76040830184610e4a565b8281036020840152602281527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60208201527f75650000000000000000000000000000000000000000000000000000000000006040820152606081019150509291505056fea164736f6c634300080f000a",
}

// OasysL2OutputOracleVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use OasysL2OutputOracleVerifierMetaData.ABI instead.
var OasysL2OutputOracleVerifierABI = OasysL2OutputOracleVerifierMetaData.ABI

// OasysL2OutputOracleVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OasysL2OutputOracleVerifierMetaData.Bin instead.
var OasysL2OutputOracleVerifierBin = OasysL2OutputOracleVerifierMetaData.Bin

// DeployOasysL2OutputOracleVerifier deploys a new Ethereum contract, binding an instance of OasysL2OutputOracleVerifier to it.
func DeployOasysL2OutputOracleVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OasysL2OutputOracleVerifier, error) {
	parsed, err := OasysL2OutputOracleVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OasysL2OutputOracleVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OasysL2OutputOracleVerifier{OasysL2OutputOracleVerifierCaller: OasysL2OutputOracleVerifierCaller{contract: contract}, OasysL2OutputOracleVerifierTransactor: OasysL2OutputOracleVerifierTransactor{contract: contract}, OasysL2OutputOracleVerifierFilterer: OasysL2OutputOracleVerifierFilterer{contract: contract}}, nil
}

// OasysL2OutputOracleVerifier is an auto generated Go binding around an Ethereum contract.
type OasysL2OutputOracleVerifier struct {
	OasysL2OutputOracleVerifierCaller     // Read-only binding to the contract
	OasysL2OutputOracleVerifierTransactor // Write-only binding to the contract
	OasysL2OutputOracleVerifierFilterer   // Log filterer for contract events
}

// OasysL2OutputOracleVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type OasysL2OutputOracleVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysL2OutputOracleVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OasysL2OutputOracleVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysL2OutputOracleVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OasysL2OutputOracleVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysL2OutputOracleVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OasysL2OutputOracleVerifierSession struct {
	Contract     *OasysL2OutputOracleVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OasysL2OutputOracleVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OasysL2OutputOracleVerifierCallerSession struct {
	Contract *OasysL2OutputOracleVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// OasysL2OutputOracleVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OasysL2OutputOracleVerifierTransactorSession struct {
	Contract     *OasysL2OutputOracleVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// OasysL2OutputOracleVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type OasysL2OutputOracleVerifierRaw struct {
	Contract *OasysL2OutputOracleVerifier // Generic contract binding to access the raw methods on
}

// OasysL2OutputOracleVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OasysL2OutputOracleVerifierCallerRaw struct {
	Contract *OasysL2OutputOracleVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// OasysL2OutputOracleVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OasysL2OutputOracleVerifierTransactorRaw struct {
	Contract *OasysL2OutputOracleVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOasysL2OutputOracleVerifier creates a new instance of OasysL2OutputOracleVerifier, bound to a specific deployed contract.
func NewOasysL2OutputOracleVerifier(address common.Address, backend bind.ContractBackend) (*OasysL2OutputOracleVerifier, error) {
	contract, err := bindOasysL2OutputOracleVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleVerifier{OasysL2OutputOracleVerifierCaller: OasysL2OutputOracleVerifierCaller{contract: contract}, OasysL2OutputOracleVerifierTransactor: OasysL2OutputOracleVerifierTransactor{contract: contract}, OasysL2OutputOracleVerifierFilterer: OasysL2OutputOracleVerifierFilterer{contract: contract}}, nil
}

// NewOasysL2OutputOracleVerifierCaller creates a new read-only instance of OasysL2OutputOracleVerifier, bound to a specific deployed contract.
func NewOasysL2OutputOracleVerifierCaller(address common.Address, caller bind.ContractCaller) (*OasysL2OutputOracleVerifierCaller, error) {
	contract, err := bindOasysL2OutputOracleVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleVerifierCaller{contract: contract}, nil
}

// NewOasysL2OutputOracleVerifierTransactor creates a new write-only instance of OasysL2OutputOracleVerifier, bound to a specific deployed contract.
func NewOasysL2OutputOracleVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*OasysL2OutputOracleVerifierTransactor, error) {
	contract, err := bindOasysL2OutputOracleVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleVerifierTransactor{contract: contract}, nil
}

// NewOasysL2OutputOracleVerifierFilterer creates a new log filterer instance of OasysL2OutputOracleVerifier, bound to a specific deployed contract.
func NewOasysL2OutputOracleVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*OasysL2OutputOracleVerifierFilterer, error) {
	contract, err := bindOasysL2OutputOracleVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleVerifierFilterer{contract: contract}, nil
}

// bindOasysL2OutputOracleVerifier binds a generic wrapper to an already deployed contract.
func bindOasysL2OutputOracleVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OasysL2OutputOracleVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OasysL2OutputOracleVerifier.Contract.OasysL2OutputOracleVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.OasysL2OutputOracleVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.OasysL2OutputOracleVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OasysL2OutputOracleVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OasysL2OutputOracleVerifier.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierSession) Version() (string, error) {
	return _OasysL2OutputOracleVerifier.Contract.Version(&_OasysL2OutputOracleVerifier.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierCallerSession) Version() (string, error) {
	return _OasysL2OutputOracleVerifier.Contract.Version(&_OasysL2OutputOracleVerifier.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactor) Approve(opts *bind.TransactOpts, l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.contract.Transact(opts, "approve", l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierSession) Approve(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.Approve(&_OasysL2OutputOracleVerifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve is a paid mutator transaction binding the contract method 0x7d2d9b63.
//
// Solidity: function approve(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactorSession) Approve(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.Approve(&_OasysL2OutputOracleVerifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactor) Approve0(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.contract.Transact(opts, "approve0", stateCommitmentChain, batchHeader, signatures)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierSession) Approve0(stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.Approve0(&_OasysL2OutputOracleVerifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdea2619d.
//
// Solidity: function approve(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactorSession) Approve0(stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.Approve0(&_OasysL2OutputOracleVerifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactor) Reject(opts *bind.TransactOpts, l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.contract.Transact(opts, "reject", l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierSession) Reject(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.Reject(&_OasysL2OutputOracleVerifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject is a paid mutator transaction binding the contract method 0x9ac45572.
//
// Solidity: function reject(address l2OutputOracle, uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactorSession) Reject(l2OutputOracle common.Address, l2OutputIndex *big.Int, l2Output TypesOutputProposal, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.Reject(&_OasysL2OutputOracleVerifier.TransactOpts, l2OutputOracle, l2OutputIndex, l2Output, signatures)
}

// Reject0 is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactor) Reject0(opts *bind.TransactOpts, stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.contract.Transact(opts, "reject0", stateCommitmentChain, batchHeader, signatures)
}

// Reject0 is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierSession) Reject0(stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.Reject0(&_OasysL2OutputOracleVerifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// Reject0 is a paid mutator transaction binding the contract method 0xd17f1fb1.
//
// Solidity: function reject(address stateCommitmentChain, (uint256,bytes32,uint256,uint256,bytes) batchHeader, bytes[] signatures) returns()
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierTransactorSession) Reject0(stateCommitmentChain common.Address, batchHeader LibOVMCodecChainBatchHeader, signatures [][]byte) (*types.Transaction, error) {
	return _OasysL2OutputOracleVerifier.Contract.Reject0(&_OasysL2OutputOracleVerifier.TransactOpts, stateCommitmentChain, batchHeader, signatures)
}

// OasysL2OutputOracleVerifierL2OutputApprovedIterator is returned from FilterL2OutputApproved and is used to iterate over the raw logs and unpacked data for L2OutputApproved events raised by the OasysL2OutputOracleVerifier contract.
type OasysL2OutputOracleVerifierL2OutputApprovedIterator struct {
	Event *OasysL2OutputOracleVerifierL2OutputApproved // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleVerifierL2OutputApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleVerifierL2OutputApproved)
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
		it.Event = new(OasysL2OutputOracleVerifierL2OutputApproved)
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
func (it *OasysL2OutputOracleVerifierL2OutputApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleVerifierL2OutputApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleVerifierL2OutputApproved represents a L2OutputApproved event raised by the OasysL2OutputOracleVerifier contract.
type OasysL2OutputOracleVerifierL2OutputApproved struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	OutputRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterL2OutputApproved is a free log retrieval operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) FilterL2OutputApproved(opts *bind.FilterOpts, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (*OasysL2OutputOracleVerifierL2OutputApprovedIterator, error) {

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

	logs, sub, err := _OasysL2OutputOracleVerifier.contract.FilterLogs(opts, "L2OutputApproved", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleVerifierL2OutputApprovedIterator{contract: _OasysL2OutputOracleVerifier.contract, event: "L2OutputApproved", logs: logs, sub: sub}, nil
}

// WatchL2OutputApproved is a free log subscription operation binding the contract event 0x3ad51d60f5a999fd533e2ea4b5691a58dd622911a02b4ff35afad2848cb9a30f.
//
// Solidity: event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) WatchL2OutputApproved(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleVerifierL2OutputApproved, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OasysL2OutputOracleVerifier.contract.WatchLogs(opts, "L2OutputApproved", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleVerifierL2OutputApproved)
				if err := _OasysL2OutputOracleVerifier.contract.UnpackLog(event, "L2OutputApproved", log); err != nil {
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
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) ParseL2OutputApproved(log types.Log) (*OasysL2OutputOracleVerifierL2OutputApproved, error) {
	event := new(OasysL2OutputOracleVerifierL2OutputApproved)
	if err := _OasysL2OutputOracleVerifier.contract.UnpackLog(event, "L2OutputApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleVerifierL2OutputRejectedIterator is returned from FilterL2OutputRejected and is used to iterate over the raw logs and unpacked data for L2OutputRejected events raised by the OasysL2OutputOracleVerifier contract.
type OasysL2OutputOracleVerifierL2OutputRejectedIterator struct {
	Event *OasysL2OutputOracleVerifierL2OutputRejected // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleVerifierL2OutputRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleVerifierL2OutputRejected)
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
		it.Event = new(OasysL2OutputOracleVerifierL2OutputRejected)
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
func (it *OasysL2OutputOracleVerifierL2OutputRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleVerifierL2OutputRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleVerifierL2OutputRejected represents a L2OutputRejected event raised by the OasysL2OutputOracleVerifier contract.
type OasysL2OutputOracleVerifierL2OutputRejected struct {
	L2OutputOracle common.Address
	L2OutputIndex  *big.Int
	OutputRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterL2OutputRejected is a free log retrieval operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) FilterL2OutputRejected(opts *bind.FilterOpts, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (*OasysL2OutputOracleVerifierL2OutputRejectedIterator, error) {

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

	logs, sub, err := _OasysL2OutputOracleVerifier.contract.FilterLogs(opts, "L2OutputRejected", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleVerifierL2OutputRejectedIterator{contract: _OasysL2OutputOracleVerifier.contract, event: "L2OutputRejected", logs: logs, sub: sub}, nil
}

// WatchL2OutputRejected is a free log subscription operation binding the contract event 0xf7a989ed650c8092db94d143b74020b09046608f0f926d2dfb9387aa03233ca3.
//
// Solidity: event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) WatchL2OutputRejected(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleVerifierL2OutputRejected, l2OutputOracle []common.Address, l2OutputIndex []*big.Int, outputRoot [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OasysL2OutputOracleVerifier.contract.WatchLogs(opts, "L2OutputRejected", l2OutputOracleRule, l2OutputIndexRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleVerifierL2OutputRejected)
				if err := _OasysL2OutputOracleVerifier.contract.UnpackLog(event, "L2OutputRejected", log); err != nil {
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
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) ParseL2OutputRejected(log types.Log) (*OasysL2OutputOracleVerifierL2OutputRejected, error) {
	event := new(OasysL2OutputOracleVerifierL2OutputRejected)
	if err := _OasysL2OutputOracleVerifier.contract.UnpackLog(event, "L2OutputRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleVerifierStateBatchApprovedIterator is returned from FilterStateBatchApproved and is used to iterate over the raw logs and unpacked data for StateBatchApproved events raised by the OasysL2OutputOracleVerifier contract.
type OasysL2OutputOracleVerifierStateBatchApprovedIterator struct {
	Event *OasysL2OutputOracleVerifierStateBatchApproved // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleVerifierStateBatchApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleVerifierStateBatchApproved)
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
		it.Event = new(OasysL2OutputOracleVerifierStateBatchApproved)
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
func (it *OasysL2OutputOracleVerifierStateBatchApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleVerifierStateBatchApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleVerifierStateBatchApproved represents a StateBatchApproved event raised by the OasysL2OutputOracleVerifier contract.
type OasysL2OutputOracleVerifierStateBatchApproved struct {
	StateCommitmentChain common.Address
	BatchIndex           *big.Int
	BatchRoot            [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStateBatchApproved is a free log retrieval operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) FilterStateBatchApproved(opts *bind.FilterOpts, stateCommitmentChain []common.Address, batchIndex []*big.Int) (*OasysL2OutputOracleVerifierStateBatchApprovedIterator, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _OasysL2OutputOracleVerifier.contract.FilterLogs(opts, "StateBatchApproved", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleVerifierStateBatchApprovedIterator{contract: _OasysL2OutputOracleVerifier.contract, event: "StateBatchApproved", logs: logs, sub: sub}, nil
}

// WatchStateBatchApproved is a free log subscription operation binding the contract event 0x4785b50466d5e0124a77ed43985b62f3b454caff4b6cfc6aa94ab02c5e5df1f5.
//
// Solidity: event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) WatchStateBatchApproved(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleVerifierStateBatchApproved, stateCommitmentChain []common.Address, batchIndex []*big.Int) (event.Subscription, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _OasysL2OutputOracleVerifier.contract.WatchLogs(opts, "StateBatchApproved", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleVerifierStateBatchApproved)
				if err := _OasysL2OutputOracleVerifier.contract.UnpackLog(event, "StateBatchApproved", log); err != nil {
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
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) ParseStateBatchApproved(log types.Log) (*OasysL2OutputOracleVerifierStateBatchApproved, error) {
	event := new(OasysL2OutputOracleVerifierStateBatchApproved)
	if err := _OasysL2OutputOracleVerifier.contract.UnpackLog(event, "StateBatchApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleVerifierStateBatchRejectedIterator is returned from FilterStateBatchRejected and is used to iterate over the raw logs and unpacked data for StateBatchRejected events raised by the OasysL2OutputOracleVerifier contract.
type OasysL2OutputOracleVerifierStateBatchRejectedIterator struct {
	Event *OasysL2OutputOracleVerifierStateBatchRejected // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleVerifierStateBatchRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleVerifierStateBatchRejected)
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
		it.Event = new(OasysL2OutputOracleVerifierStateBatchRejected)
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
func (it *OasysL2OutputOracleVerifierStateBatchRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleVerifierStateBatchRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleVerifierStateBatchRejected represents a StateBatchRejected event raised by the OasysL2OutputOracleVerifier contract.
type OasysL2OutputOracleVerifierStateBatchRejected struct {
	StateCommitmentChain common.Address
	BatchIndex           *big.Int
	BatchRoot            [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStateBatchRejected is a free log retrieval operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) FilterStateBatchRejected(opts *bind.FilterOpts, stateCommitmentChain []common.Address, batchIndex []*big.Int) (*OasysL2OutputOracleVerifierStateBatchRejectedIterator, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _OasysL2OutputOracleVerifier.contract.FilterLogs(opts, "StateBatchRejected", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleVerifierStateBatchRejectedIterator{contract: _OasysL2OutputOracleVerifier.contract, event: "StateBatchRejected", logs: logs, sub: sub}, nil
}

// WatchStateBatchRejected is a free log subscription operation binding the contract event 0x2d6c30e420d4abe074221e2e344cec3bd0f565a2e09435a3d9d738f69332a6dd.
//
// Solidity: event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot)
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) WatchStateBatchRejected(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleVerifierStateBatchRejected, stateCommitmentChain []common.Address, batchIndex []*big.Int) (event.Subscription, error) {

	var stateCommitmentChainRule []interface{}
	for _, stateCommitmentChainItem := range stateCommitmentChain {
		stateCommitmentChainRule = append(stateCommitmentChainRule, stateCommitmentChainItem)
	}
	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _OasysL2OutputOracleVerifier.contract.WatchLogs(opts, "StateBatchRejected", stateCommitmentChainRule, batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleVerifierStateBatchRejected)
				if err := _OasysL2OutputOracleVerifier.contract.UnpackLog(event, "StateBatchRejected", log); err != nil {
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
func (_OasysL2OutputOracleVerifier *OasysL2OutputOracleVerifierFilterer) ParseStateBatchRejected(log types.Log) (*OasysL2OutputOracleVerifierStateBatchRejected, error) {
	event := new(OasysL2OutputOracleVerifierStateBatchRejected)
	if err := _OasysL2OutputOracleVerifier.contract.UnpackLog(event, "StateBatchRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
