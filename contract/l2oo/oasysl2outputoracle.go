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

// TypesOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// OasysL2OutputOracleMetaData contains all meta data concerning the OasysL2OutputOracle contract.
var OasysL2OutputOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submissionInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_finalizationPeriodSeconds\",\"type\":\"uint256\"},{\"internalType\":\"contractIOasysL2OutputOracleVerifier\",\"name\":\"_verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"name\":\"OutputFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"OutputProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"name\":\"OutputVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevNextOutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newNextOutputIndex\",\"type\":\"uint256\"}],\"name\":\"OutputsDeleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHALLENGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FINALIZATION_PERIOD_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMISSION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER\",\"outputs\":[{\"internalType\":\"contractIOasysL2OutputOracleVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"computeL2Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"deleteL2Outputs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"}],\"name\":\"failVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizationPeriodSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"getL2Output\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"getL2OutputAfter\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"getL2OutputIndexAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTimestamp\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"isOutputFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2OracleVerifier\",\"outputs\":[{\"internalType\":\"contractIOasysL2OutputOracleVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextVerifyIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l1BlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l1BlockNumber\",\"type\":\"uint256\"}],\"name\":\"proposeL2Output\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submissionInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.OutputProposal\",\"name\":\"l2Output\",\"type\":\"tuple\"}],\"name\":\"succeedVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifiedL1Timestamp\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620020c8380380620020c883398101604081905262000035916200037c565b8787878787878760008611620000b85760405162461bcd60e51b815260206004820152603460248201527f4c324f75747075744f7261636c653a204c3220626c6f636b2074696d65206d7560448201527f73742062652067726561746572207468616e203000000000000000000000000060648201526084015b60405180910390fd5b60008711620001305760405162461bcd60e51b815260206004820152603a60248201527f4c324f75747075744f7261636c653a207375626d697373696f6e20696e74657260448201527f76616c206d7573742062652067726561746572207468616e20300000000000006064820152608401620000af565b608087905260a08690526001600160a01b0380841660e052821660c0526101008190526200015f858562000183565b5050506001600160a01b039094166101205250620004019950505050505050505050565b6200019a82826200019e60201b620013bf1760201c565b5050565b600054610100900460ff1615808015620001bf5750600054600160ff909116105b80620001ef5750620001dc306200035460201b620015d81760201c565b158015620001ef575060005460ff166001145b620002545760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401620000af565b6000805460ff19166001179055801562000278576000805461ff0019166101001790555b42821115620002fe5760405162461bcd60e51b8152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201526374696d6560e01b608482015260a401620000af565b6002829055600183905580156200034f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6001600160a01b03163b151590565b6001600160a01b03811681146200037957600080fd5b50565b600080600080600080600080610100898b0312156200039a57600080fd5b885197506020890151965060408901519550606089015194506080890151620003c38162000363565b60a08a0151909450620003d68162000363565b60c08a015160e08b01519194509250620003f08162000363565b809150509295985092959890939650565b60805160a05160c05160e0516101005161012051611c1c620004ac600039600081816102350152818161035d015281816109b40152611188015260008181610601015281816106ff015261160a015260008181610564015281816105d00152610c190152600081816102d4015281816103d30152610b3c0152600081816101ee015281816104c2015261112c0152600081816102a3015281816106a901526113800152611c1c6000f3fe6080604052600436106101d75760003560e01c80638878627211610102578063ce5db8d611610095578063dcec334811610064578063dcec334814610685578063e1a41bcf1461069a578063e4a30116146106cd578063f4daa291146106ed57600080fd5b8063ce5db8d6146105f2578063cf8e5cf014610625578063d1de856c14610645578063d88234361461066557600080fd5b8063a25ae557116100d1578063a25ae557146104f9578063a8e4fb9014610555578063b3bd5cf514610588578063bffa7f0f146105be57600080fd5b8063887862721461047d57806389c44cbb1461049357806393991af3146104b35780639aaab648146104e657600080fd5b806369f16eec1161017a5780636dbffb78116101495780636dbffb78146103f557806370872aa5146104255780637f0064201461043b578063859c029d1461045b57600080fd5b806369f16eec146103815780636abcf563146103965780636b405db0146103ab5780636b4d98dd146103c157600080fd5b8063529933df116101b6578063529933df14610291578063534db0e2146102c557806354fd4d50146102f85780635a5a4d9e1461034e57600080fd5b80622134cc146101dc57806308c84e70146102235780634599c7881461027c575b600080fd5b3480156101e857600080fd5b506102107f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b34801561022f57600080fd5b506102577f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161021a565b34801561028857600080fd5b50610210610721565b34801561029d57600080fd5b506102107f000000000000000000000000000000000000000000000000000000000000000081565b3480156102d157600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610257565b34801561030457600080fd5b506103416040518060400160405280600581526020017f312e372e3000000000000000000000000000000000000000000000000000000081525081565b60405161021a9190611963565b34801561035a57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610257565b34801561038d57600080fd5b50610210610794565b3480156103a257600080fd5b50600354610210565b3480156103b757600080fd5b5061021060325481565b3480156103cd57600080fd5b506102577f000000000000000000000000000000000000000000000000000000000000000081565b34801561040157600080fd5b506104156104103660046119d6565b6107a6565b604051901515815260200161021a565b34801561043157600080fd5b5061021060015481565b34801561044757600080fd5b506102106104563660046119d6565b6107b7565b34801561046757600080fd5b5061047b6104763660046119ef565b61099c565b005b34801561048957600080fd5b5061021060025481565b34801561049f57600080fd5b5061047b6104ae3660046119d6565b610b24565b3480156104bf57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610210565b61047b6104f4366004611a45565b610c01565b34801561050557600080fd5b506105196105143660046119d6565b610ffe565b60408051825181526020808401516fffffffffffffffffffffffffffffffff90811691830191909152928201519092169082015260600161021a565b34801561056157600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610257565b34801561059457600080fd5b5061059d611092565b6040516fffffffffffffffffffffffffffffffff909116815260200161021a565b3480156105ca57600080fd5b506102577f000000000000000000000000000000000000000000000000000000000000000081565b3480156105fe57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610210565b34801561063157600080fd5b506105196106403660046119d6565b6110f0565b34801561065157600080fd5b506102106106603660046119d6565b611128565b34801561067157600080fd5b5061047b6106803660046119ef565b611170565b34801561069157600080fd5b5061021061137c565b3480156106a657600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610210565b3480156106d957600080fd5b5061047b6106e8366004611a77565b6113b1565b3480156106f957600080fd5b506102107f000000000000000000000000000000000000000000000000000000000000000081565b6003546000901561078b576003805461073c90600190611ac8565b8154811061074c5761074c611adf565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6001545b905090565b60035460009061078f90600190611ac8565b60006107b1826115f4565b92915050565b60006107c1610721565b8211156108615760405162461bcd60e51b815260206004820152604860248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f7420666f72206120626c6f636b207468617420686173206e6f74206265656e2060648201527f70726f706f736564000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b6003546108fc5760405162461bcd60e51b815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f74206173206e6f206f7574707574732068617665206265656e2070726f706f7360648201527f6564207965740000000000000000000000000000000000000000000000000000608482015260a401610858565b6003546000905b8082101561099557600060026109198385611b0e565b6109239190611b26565b9050846003828154811061093957610939611adf565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16101561098b57610984816001611b0e565b925061098f565b8091505b50610903565b5092915050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a475760405162461bcd60e51b815260206004820152602a60248201527f4f617379734c324f75747075744f7261636c653a2063616c6c6572206973206e60448201527f6f7420616c6c6f776564000000000000000000000000000000000000000000006064820152608401610858565b610a518282611682565b610ac35760405162461bcd60e51b815260206004820152602860248201527f4f617379734c324f75747075744f7261636c653a20696e76616c6964206f757460448201527f70757420726f6f740000000000000000000000000000000000000000000000006064820152608401610858565b610acc826117e1565b610adc6060820160408301611b61565b6fffffffffffffffffffffffffffffffff168160000135837f39dab377091f907ce5cdb2ef4eeea3961f83ca13a3b6daff133ec7f99a3ec34360405160405180910390a45050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610bf55760405162461bcd60e51b815260206004820152604360248201527f4f617379734c324f75747075744f7261636c653a206f6e6c792074686520636860448201527f616c6c656e67657220616464726573732063616e2064656c657465206f75747060648201527f7574730000000000000000000000000000000000000000000000000000000000608482015260a401610858565b610bfe816117e1565b50565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610cd25760405162461bcd60e51b815260206004820152604160248201527f4c324f75747075744f7261636c653a206f6e6c79207468652070726f706f736560448201527f7220616464726573732063616e2070726f706f7365206e6577206f757470757460648201527f7300000000000000000000000000000000000000000000000000000000000000608482015260a401610858565b610cda61137c565b8314610d745760405162461bcd60e51b815260206004820152604860248201527f4c324f75747075744f7261636c653a20626c6f636b206e756d626572206d757360448201527f7420626520657175616c20746f206e65787420657870656374656420626c6f6360648201527f6b206e756d626572000000000000000000000000000000000000000000000000608482015260a401610858565b42610d7e84611128565b10610df15760405162461bcd60e51b815260206004820152603660248201527f4c324f75747075744f7261636c653a2063616e6e6f742070726f706f7365204c60448201527f32206f757470757420696e2074686520667574757265000000000000000000006064820152608401610858565b83610e645760405162461bcd60e51b815260206004820152603a60248201527f4c324f75747075744f7261636c653a204c32206f75747075742070726f706f7360448201527f616c2063616e6e6f7420626520746865207a65726f20686173680000000000006064820152608401610858565b8115610f065781814014610f065760405162461bcd60e51b815260206004820152604960248201527f4c324f75747075744f7261636c653a20626c6f636b206861736820646f65732060448201527f6e6f74206d61746368207468652068617368206174207468652065787065637460648201527f6564206865696768740000000000000000000000000000000000000000000000608482015260a401610858565b82610f1060035490565b857fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e242604051610f4291815260200190565b60405180910390a45050604080516060810182529283526fffffffffffffffffffffffffffffffff4281166020850190815292811691840191825260038054600181018255600091909152935160029094027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810194909455915190518216700100000000000000000000000000000000029116177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c90910155565b60408051606081018252600080825260208201819052918101919091526003828154811061102e5761102e611adf565b600091825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff8082169484019490945270010000000000000000000000000000000090049092169181019190915292915050565b60006032546000146110ea57600360016032546110af9190611ac8565b815481106110bf576110bf611adf565b60009182526020909120600160029092020101546fffffffffffffffffffffffffffffffff16905090565b50600090565b60408051606081018252600080825260208201819052918101919091526003611118836107b7565b8154811061102e5761102e611adf565b60007f0000000000000000000000000000000000000000000000000000000000000000600154836111599190611ac8565b6111639190611b9a565b6002546107b19190611b0e565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461121b5760405162461bcd60e51b815260206004820152602a60248201527f4f617379734c324f75747075744f7261636c653a2063616c6c6572206973206e60448201527f6f7420616c6c6f776564000000000000000000000000000000000000000000006064820152608401610858565b6112258282611682565b6112975760405162461bcd60e51b815260206004820152602860248201527f4f617379734c324f75747075744f7261636c653a20696e76616c6964206f757460448201527f70757420726f6f740000000000000000000000000000000000000000000000006064820152608401610858565b603254821461130e5760405162461bcd60e51b815260206004820152602c60248201527f4f617379734c324f75747075744f7261636c653a20696e76616c6964204c322060448201527f6f757470757420696e64657800000000000000000000000000000000000000006064820152608401610858565b6032805490600061131e83611bd7565b9091555061133490506060820160408301611b61565b6fffffffffffffffffffffffffffffffff168160000135837f98beebb56f3d9f23ca4c2e85cac607b80e030b3ca74fcc7b38f98bdd5146ef6b60405160405180910390a45050565b60007f00000000000000000000000000000000000000000000000000000000000000006113a7610721565b61078f9190611b0e565b6113bb82826113bf565b5050565b600054610100900460ff16158080156113df5750600054600160ff909116105b806113f95750303b1580156113f9575060005460ff166001145b61146b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610858565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156114c957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b428211156115665760405162461bcd60e51b8152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201527f74696d6500000000000000000000000000000000000000000000000000000000608482015260a401610858565b6002829055600183905580156115d357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b600060325482101561160857506001919050565b7f00000000000000000000000000000000000000000000000000000000000000006003838154811061163c5761163c611adf565b600091825260209091206001600290920201015461166c906fffffffffffffffffffffffffffffffff1642611ac8565b111561167a57506001919050565b506000919050565b6000806003848154811061169857611698611adf565b6000918252602091829020604080516060808201835260029490940290920180548084526001909101546fffffffffffffffffffffffffffffffff808216858801527001000000000000000000000000000000008204908116858501528351968701929092527fffffffffffffffffffffffffffffffff00000000000000000000000000000000608091821b8116938701939093521b1660508401529250016040516020818303038152906040528051906020012083600001358460200160208101906117659190611b61565b6117756060870160408801611b61565b6040516020016117c193929190928352608091821b7fffffffffffffffffffffffffffffffff000000000000000000000000000000009081166020850152911b16603082015260400190565b604051602081830303815290604052805190602001201491505092915050565b600354811061187e5760405162461bcd60e51b815260206004820152604860248201527f4f617379734c324f75747075744f7261636c653a2063616e6e6f742064656c6560448201527f7465206f75747075747320616674657220746865206c6174657374206f75747060648201527f757420696e646578000000000000000000000000000000000000000000000000608482015260a401610858565b611887816115f4565b156119205760405162461bcd60e51b815260206004820152604b60248201527f4f617379734c324f75747075744f7261636c653a2063616e6e6f742064656c6560448201527f7465206f7574707574732074686174206861766520616c72656164792062656560648201527f6e2066696e616c697a6564000000000000000000000000000000000000000000608482015260a401610858565b600061192b60035490565b90508160035581817f4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b660405160405180910390a35050565b600060208083528351808285015260005b8181101561199057858101830151858201604001528201611974565b818111156119a2576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b6000602082840312156119e857600080fd5b5035919050565b6000808284036080811215611a0357600080fd5b8335925060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082011215611a3757600080fd5b506020830190509250929050565b60008060008060808587031215611a5b57600080fd5b5050823594602084013594506040840135936060013592509050565b60008060408385031215611a8a57600080fd5b50508035926020909101359150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611ada57611ada611a99565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008219821115611b2157611b21611a99565b500190565b600082611b5c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600060208284031215611b7357600080fd5b81356fffffffffffffffffffffffffffffffff81168114611b9357600080fd5b9392505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611bd257611bd2611a99565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c0857611c08611a99565b506001019056fea164736f6c634300080f000a",
}

// OasysL2OutputOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OasysL2OutputOracleMetaData.ABI instead.
var OasysL2OutputOracleABI = OasysL2OutputOracleMetaData.ABI

// OasysL2OutputOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OasysL2OutputOracleMetaData.Bin instead.
var OasysL2OutputOracleBin = OasysL2OutputOracleMetaData.Bin

// DeployOasysL2OutputOracle deploys a new Ethereum contract, binding an instance of OasysL2OutputOracle to it.
func DeployOasysL2OutputOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _submissionInterval *big.Int, _l2BlockTime *big.Int, _startingBlockNumber *big.Int, _startingTimestamp *big.Int, _proposer common.Address, _challenger common.Address, _finalizationPeriodSeconds *big.Int, _verifier common.Address) (common.Address, *types.Transaction, *OasysL2OutputOracle, error) {
	parsed, err := OasysL2OutputOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OasysL2OutputOracleBin), backend, _submissionInterval, _l2BlockTime, _startingBlockNumber, _startingTimestamp, _proposer, _challenger, _finalizationPeriodSeconds, _verifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OasysL2OutputOracle{OasysL2OutputOracleCaller: OasysL2OutputOracleCaller{contract: contract}, OasysL2OutputOracleTransactor: OasysL2OutputOracleTransactor{contract: contract}, OasysL2OutputOracleFilterer: OasysL2OutputOracleFilterer{contract: contract}}, nil
}

// OasysL2OutputOracle is an auto generated Go binding around an Ethereum contract.
type OasysL2OutputOracle struct {
	OasysL2OutputOracleCaller     // Read-only binding to the contract
	OasysL2OutputOracleTransactor // Write-only binding to the contract
	OasysL2OutputOracleFilterer   // Log filterer for contract events
}

// OasysL2OutputOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OasysL2OutputOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysL2OutputOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OasysL2OutputOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysL2OutputOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OasysL2OutputOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasysL2OutputOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OasysL2OutputOracleSession struct {
	Contract     *OasysL2OutputOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OasysL2OutputOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OasysL2OutputOracleCallerSession struct {
	Contract *OasysL2OutputOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OasysL2OutputOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OasysL2OutputOracleTransactorSession struct {
	Contract     *OasysL2OutputOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OasysL2OutputOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OasysL2OutputOracleRaw struct {
	Contract *OasysL2OutputOracle // Generic contract binding to access the raw methods on
}

// OasysL2OutputOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OasysL2OutputOracleCallerRaw struct {
	Contract *OasysL2OutputOracleCaller // Generic read-only contract binding to access the raw methods on
}

// OasysL2OutputOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OasysL2OutputOracleTransactorRaw struct {
	Contract *OasysL2OutputOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOasysL2OutputOracle creates a new instance of OasysL2OutputOracle, bound to a specific deployed contract.
func NewOasysL2OutputOracle(address common.Address, backend bind.ContractBackend) (*OasysL2OutputOracle, error) {
	contract, err := bindOasysL2OutputOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracle{OasysL2OutputOracleCaller: OasysL2OutputOracleCaller{contract: contract}, OasysL2OutputOracleTransactor: OasysL2OutputOracleTransactor{contract: contract}, OasysL2OutputOracleFilterer: OasysL2OutputOracleFilterer{contract: contract}}, nil
}

// NewOasysL2OutputOracleCaller creates a new read-only instance of OasysL2OutputOracle, bound to a specific deployed contract.
func NewOasysL2OutputOracleCaller(address common.Address, caller bind.ContractCaller) (*OasysL2OutputOracleCaller, error) {
	contract, err := bindOasysL2OutputOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleCaller{contract: contract}, nil
}

// NewOasysL2OutputOracleTransactor creates a new write-only instance of OasysL2OutputOracle, bound to a specific deployed contract.
func NewOasysL2OutputOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OasysL2OutputOracleTransactor, error) {
	contract, err := bindOasysL2OutputOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleTransactor{contract: contract}, nil
}

// NewOasysL2OutputOracleFilterer creates a new log filterer instance of OasysL2OutputOracle, bound to a specific deployed contract.
func NewOasysL2OutputOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OasysL2OutputOracleFilterer, error) {
	contract, err := bindOasysL2OutputOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleFilterer{contract: contract}, nil
}

// bindOasysL2OutputOracle binds a generic wrapper to an already deployed contract.
func bindOasysL2OutputOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OasysL2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OasysL2OutputOracle *OasysL2OutputOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OasysL2OutputOracle.Contract.OasysL2OutputOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OasysL2OutputOracle *OasysL2OutputOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.OasysL2OutputOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OasysL2OutputOracle *OasysL2OutputOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.OasysL2OutputOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OasysL2OutputOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.contract.Transact(opts, method, params...)
}

// CHALLENGER is a free data retrieval call binding the contract method 0x6b4d98dd.
//
// Solidity: function CHALLENGER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) CHALLENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "CHALLENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CHALLENGER is a free data retrieval call binding the contract method 0x6b4d98dd.
//
// Solidity: function CHALLENGER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) CHALLENGER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.CHALLENGER(&_OasysL2OutputOracle.CallOpts)
}

// CHALLENGER is a free data retrieval call binding the contract method 0x6b4d98dd.
//
// Solidity: function CHALLENGER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) CHALLENGER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.CHALLENGER(&_OasysL2OutputOracle.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) FINALIZATIONPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "FINALIZATION_PERIOD_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.FINALIZATIONPERIODSECONDS(&_OasysL2OutputOracle.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.FINALIZATIONPERIODSECONDS(&_OasysL2OutputOracle.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) L2BLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "L2_BLOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) L2BLOCKTIME() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.L2BLOCKTIME(&_OasysL2OutputOracle.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) L2BLOCKTIME() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.L2BLOCKTIME(&_OasysL2OutputOracle.CallOpts)
}

// PROPOSER is a free data retrieval call binding the contract method 0xbffa7f0f.
//
// Solidity: function PROPOSER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) PROPOSER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "PROPOSER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROPOSER is a free data retrieval call binding the contract method 0xbffa7f0f.
//
// Solidity: function PROPOSER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) PROPOSER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.PROPOSER(&_OasysL2OutputOracle.CallOpts)
}

// PROPOSER is a free data retrieval call binding the contract method 0xbffa7f0f.
//
// Solidity: function PROPOSER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) PROPOSER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.PROPOSER(&_OasysL2OutputOracle.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) SUBMISSIONINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "SUBMISSION_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.SUBMISSIONINTERVAL(&_OasysL2OutputOracle.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.SUBMISSIONINTERVAL(&_OasysL2OutputOracle.CallOpts)
}

// VERIFIER is a free data retrieval call binding the contract method 0x08c84e70.
//
// Solidity: function VERIFIER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) VERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VERIFIER is a free data retrieval call binding the contract method 0x08c84e70.
//
// Solidity: function VERIFIER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) VERIFIER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.VERIFIER(&_OasysL2OutputOracle.CallOpts)
}

// VERIFIER is a free data retrieval call binding the contract method 0x08c84e70.
//
// Solidity: function VERIFIER() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) VERIFIER() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.VERIFIER(&_OasysL2OutputOracle.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) Challenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "challenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) Challenger() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.Challenger(&_OasysL2OutputOracle.CallOpts)
}

// Challenger is a free data retrieval call binding the contract method 0x534db0e2.
//
// Solidity: function challenger() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) Challenger() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.Challenger(&_OasysL2OutputOracle.CallOpts)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) ComputeL2Timestamp(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "computeL2Timestamp", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.ComputeL2Timestamp(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.ComputeL2Timestamp(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// FinalizationPeriodSeconds is a free data retrieval call binding the contract method 0xce5db8d6.
//
// Solidity: function finalizationPeriodSeconds() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) FinalizationPeriodSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "finalizationPeriodSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalizationPeriodSeconds is a free data retrieval call binding the contract method 0xce5db8d6.
//
// Solidity: function finalizationPeriodSeconds() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) FinalizationPeriodSeconds() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.FinalizationPeriodSeconds(&_OasysL2OutputOracle.CallOpts)
}

// FinalizationPeriodSeconds is a free data retrieval call binding the contract method 0xce5db8d6.
//
// Solidity: function finalizationPeriodSeconds() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) FinalizationPeriodSeconds() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.FinalizationPeriodSeconds(&_OasysL2OutputOracle.CallOpts)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) GetL2Output(opts *bind.CallOpts, _l2OutputIndex *big.Int) (TypesOutputProposal, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "getL2Output", _l2OutputIndex)

	if err != nil {
		return *new(TypesOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesOutputProposal)).(*TypesOutputProposal)

	return out0, err

}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) GetL2Output(_l2OutputIndex *big.Int) (TypesOutputProposal, error) {
	return _OasysL2OutputOracle.Contract.GetL2Output(&_OasysL2OutputOracle.CallOpts, _l2OutputIndex)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) GetL2Output(_l2OutputIndex *big.Int) (TypesOutputProposal, error) {
	return _OasysL2OutputOracle.Contract.GetL2Output(&_OasysL2OutputOracle.CallOpts, _l2OutputIndex)
}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) GetL2OutputAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (TypesOutputProposal, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "getL2OutputAfter", _l2BlockNumber)

	if err != nil {
		return *new(TypesOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesOutputProposal)).(*TypesOutputProposal)

	return out0, err

}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) GetL2OutputAfter(_l2BlockNumber *big.Int) (TypesOutputProposal, error) {
	return _OasysL2OutputOracle.Contract.GetL2OutputAfter(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((bytes32,uint128,uint128))
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) GetL2OutputAfter(_l2BlockNumber *big.Int) (TypesOutputProposal, error) {
	return _OasysL2OutputOracle.Contract.GetL2OutputAfter(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) GetL2OutputIndexAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "getL2OutputIndexAfter", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) GetL2OutputIndexAfter(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.GetL2OutputIndexAfter(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) GetL2OutputIndexAfter(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.GetL2OutputIndexAfter(&_OasysL2OutputOracle.CallOpts, _l2BlockNumber)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 l2OutputIndex) view returns(bool)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) IsOutputFinalized(opts *bind.CallOpts, l2OutputIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "isOutputFinalized", l2OutputIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 l2OutputIndex) view returns(bool)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) IsOutputFinalized(l2OutputIndex *big.Int) (bool, error) {
	return _OasysL2OutputOracle.Contract.IsOutputFinalized(&_OasysL2OutputOracle.CallOpts, l2OutputIndex)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 l2OutputIndex) view returns(bool)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) IsOutputFinalized(l2OutputIndex *big.Int) (bool, error) {
	return _OasysL2OutputOracle.Contract.IsOutputFinalized(&_OasysL2OutputOracle.CallOpts, l2OutputIndex)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) L2BlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "l2BlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) L2BlockTime() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.L2BlockTime(&_OasysL2OutputOracle.CallOpts)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) L2BlockTime() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.L2BlockTime(&_OasysL2OutputOracle.CallOpts)
}

// L2OracleVerifier is a free data retrieval call binding the contract method 0x5a5a4d9e.
//
// Solidity: function l2OracleVerifier() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) L2OracleVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "l2OracleVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2OracleVerifier is a free data retrieval call binding the contract method 0x5a5a4d9e.
//
// Solidity: function l2OracleVerifier() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) L2OracleVerifier() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.L2OracleVerifier(&_OasysL2OutputOracle.CallOpts)
}

// L2OracleVerifier is a free data retrieval call binding the contract method 0x5a5a4d9e.
//
// Solidity: function l2OracleVerifier() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) L2OracleVerifier() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.L2OracleVerifier(&_OasysL2OutputOracle.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) LatestBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.LatestBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.LatestBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) LatestOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "latestOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) LatestOutputIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.LatestOutputIndex(&_OasysL2OutputOracle.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) LatestOutputIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.LatestOutputIndex(&_OasysL2OutputOracle.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) NextBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "nextBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) NextBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) NextBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) NextOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "nextOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) NextOutputIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextOutputIndex(&_OasysL2OutputOracle.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) NextOutputIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextOutputIndex(&_OasysL2OutputOracle.CallOpts)
}

// NextVerifyIndex is a free data retrieval call binding the contract method 0x6b405db0.
//
// Solidity: function nextVerifyIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) NextVerifyIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "nextVerifyIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextVerifyIndex is a free data retrieval call binding the contract method 0x6b405db0.
//
// Solidity: function nextVerifyIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) NextVerifyIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextVerifyIndex(&_OasysL2OutputOracle.CallOpts)
}

// NextVerifyIndex is a free data retrieval call binding the contract method 0x6b405db0.
//
// Solidity: function nextVerifyIndex() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) NextVerifyIndex() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.NextVerifyIndex(&_OasysL2OutputOracle.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) Proposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "proposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) Proposer() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.Proposer(&_OasysL2OutputOracle.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) Proposer() (common.Address, error) {
	return _OasysL2OutputOracle.Contract.Proposer(&_OasysL2OutputOracle.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) StartingBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "startingBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) StartingBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.StartingBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) StartingBlockNumber() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.StartingBlockNumber(&_OasysL2OutputOracle.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) StartingTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "startingTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) StartingTimestamp() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.StartingTimestamp(&_OasysL2OutputOracle.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) StartingTimestamp() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.StartingTimestamp(&_OasysL2OutputOracle.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) SubmissionInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "submissionInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) SubmissionInterval() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.SubmissionInterval(&_OasysL2OutputOracle.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) SubmissionInterval() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.SubmissionInterval(&_OasysL2OutputOracle.CallOpts)
}

// VerifiedL1Timestamp is a free data retrieval call binding the contract method 0xb3bd5cf5.
//
// Solidity: function verifiedL1Timestamp() view returns(uint128)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) VerifiedL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "verifiedL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifiedL1Timestamp is a free data retrieval call binding the contract method 0xb3bd5cf5.
//
// Solidity: function verifiedL1Timestamp() view returns(uint128)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) VerifiedL1Timestamp() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.VerifiedL1Timestamp(&_OasysL2OutputOracle.CallOpts)
}

// VerifiedL1Timestamp is a free data retrieval call binding the contract method 0xb3bd5cf5.
//
// Solidity: function verifiedL1Timestamp() view returns(uint128)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) VerifiedL1Timestamp() (*big.Int, error) {
	return _OasysL2OutputOracle.Contract.VerifiedL1Timestamp(&_OasysL2OutputOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OasysL2OutputOracle *OasysL2OutputOracleCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OasysL2OutputOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) Version() (string, error) {
	return _OasysL2OutputOracle.Contract.Version(&_OasysL2OutputOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OasysL2OutputOracle *OasysL2OutputOracleCallerSession) Version() (string, error) {
	return _OasysL2OutputOracle.Contract.Version(&_OasysL2OutputOracle.CallOpts)
}

// DeleteL2Outputs is a paid mutator transaction binding the contract method 0x89c44cbb.
//
// Solidity: function deleteL2Outputs(uint256 l2OutputIndex) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) DeleteL2Outputs(opts *bind.TransactOpts, l2OutputIndex *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "deleteL2Outputs", l2OutputIndex)
}

// DeleteL2Outputs is a paid mutator transaction binding the contract method 0x89c44cbb.
//
// Solidity: function deleteL2Outputs(uint256 l2OutputIndex) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) DeleteL2Outputs(l2OutputIndex *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.DeleteL2Outputs(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex)
}

// DeleteL2Outputs is a paid mutator transaction binding the contract method 0x89c44cbb.
//
// Solidity: function deleteL2Outputs(uint256 l2OutputIndex) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) DeleteL2Outputs(l2OutputIndex *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.DeleteL2Outputs(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex)
}

// FailVerification is a paid mutator transaction binding the contract method 0x859c029d.
//
// Solidity: function failVerification(uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) FailVerification(opts *bind.TransactOpts, l2OutputIndex *big.Int, l2Output TypesOutputProposal) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "failVerification", l2OutputIndex, l2Output)
}

// FailVerification is a paid mutator transaction binding the contract method 0x859c029d.
//
// Solidity: function failVerification(uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) FailVerification(l2OutputIndex *big.Int, l2Output TypesOutputProposal) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.FailVerification(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex, l2Output)
}

// FailVerification is a paid mutator transaction binding the contract method 0x859c029d.
//
// Solidity: function failVerification(uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) FailVerification(l2OutputIndex *big.Int, l2Output TypesOutputProposal) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.FailVerification(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex, l2Output)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) Initialize(opts *bind.TransactOpts, _startingBlockNumber *big.Int, _startingTimestamp *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "initialize", _startingBlockNumber, _startingTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) Initialize(_startingBlockNumber *big.Int, _startingTimestamp *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.Initialize(&_OasysL2OutputOracle.TransactOpts, _startingBlockNumber, _startingTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) Initialize(_startingBlockNumber *big.Int, _startingTimestamp *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.Initialize(&_OasysL2OutputOracle.TransactOpts, _startingBlockNumber, _startingTimestamp)
}

// ProposeL2Output is a paid mutator transaction binding the contract method 0x9aaab648.
//
// Solidity: function proposeL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber) payable returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) ProposeL2Output(opts *bind.TransactOpts, _outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "proposeL2Output", _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
}

// ProposeL2Output is a paid mutator transaction binding the contract method 0x9aaab648.
//
// Solidity: function proposeL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber) payable returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) ProposeL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.ProposeL2Output(&_OasysL2OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
}

// ProposeL2Output is a paid mutator transaction binding the contract method 0x9aaab648.
//
// Solidity: function proposeL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber) payable returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) ProposeL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.ProposeL2Output(&_OasysL2OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
}

// SucceedVerification is a paid mutator transaction binding the contract method 0xd8823436.
//
// Solidity: function succeedVerification(uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactor) SucceedVerification(opts *bind.TransactOpts, l2OutputIndex *big.Int, l2Output TypesOutputProposal) (*types.Transaction, error) {
	return _OasysL2OutputOracle.contract.Transact(opts, "succeedVerification", l2OutputIndex, l2Output)
}

// SucceedVerification is a paid mutator transaction binding the contract method 0xd8823436.
//
// Solidity: function succeedVerification(uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleSession) SucceedVerification(l2OutputIndex *big.Int, l2Output TypesOutputProposal) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.SucceedVerification(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex, l2Output)
}

// SucceedVerification is a paid mutator transaction binding the contract method 0xd8823436.
//
// Solidity: function succeedVerification(uint256 l2OutputIndex, (bytes32,uint128,uint128) l2Output) returns()
func (_OasysL2OutputOracle *OasysL2OutputOracleTransactorSession) SucceedVerification(l2OutputIndex *big.Int, l2Output TypesOutputProposal) (*types.Transaction, error) {
	return _OasysL2OutputOracle.Contract.SucceedVerification(&_OasysL2OutputOracle.TransactOpts, l2OutputIndex, l2Output)
}

// OasysL2OutputOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleInitializedIterator struct {
	Event *OasysL2OutputOracleInitialized // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleInitialized)
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
		it.Event = new(OasysL2OutputOracleInitialized)
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
func (it *OasysL2OutputOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleInitialized represents a Initialized event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*OasysL2OutputOracleInitializedIterator, error) {

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleInitializedIterator{contract: _OasysL2OutputOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleInitialized)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseInitialized(log types.Log) (*OasysL2OutputOracleInitialized, error) {
	event := new(OasysL2OutputOracleInitialized)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleOutputFailedIterator is returned from FilterOutputFailed and is used to iterate over the raw logs and unpacked data for OutputFailed events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputFailedIterator struct {
	Event *OasysL2OutputOracleOutputFailed // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleOutputFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleOutputFailed)
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
		it.Event = new(OasysL2OutputOracleOutputFailed)
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
func (it *OasysL2OutputOracleOutputFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleOutputFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleOutputFailed represents a OutputFailed event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputFailed struct {
	L2OutputIndex *big.Int
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputFailed is a free log retrieval operation binding the contract event 0x39dab377091f907ce5cdb2ef4eeea3961f83ca13a3b6daff133ec7f99a3ec343.
//
// Solidity: event OutputFailed(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterOutputFailed(opts *bind.FilterOpts, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (*OasysL2OutputOracleOutputFailedIterator, error) {

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

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "OutputFailed", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleOutputFailedIterator{contract: _OasysL2OutputOracle.contract, event: "OutputFailed", logs: logs, sub: sub}, nil
}

// WatchOutputFailed is a free log subscription operation binding the contract event 0x39dab377091f907ce5cdb2ef4eeea3961f83ca13a3b6daff133ec7f99a3ec343.
//
// Solidity: event OutputFailed(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchOutputFailed(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleOutputFailed, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "OutputFailed", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleOutputFailed)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputFailed", log); err != nil {
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
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseOutputFailed(log types.Log) (*OasysL2OutputOracleOutputFailed, error) {
	event := new(OasysL2OutputOracleOutputFailed)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleOutputProposedIterator is returned from FilterOutputProposed and is used to iterate over the raw logs and unpacked data for OutputProposed events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputProposedIterator struct {
	Event *OasysL2OutputOracleOutputProposed // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleOutputProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleOutputProposed)
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
		it.Event = new(OasysL2OutputOracleOutputProposed)
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
func (it *OasysL2OutputOracleOutputProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleOutputProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleOutputProposed represents a OutputProposed event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputProposed struct {
	OutputRoot    [32]byte
	L2OutputIndex *big.Int
	L2BlockNumber *big.Int
	L1Timestamp   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputProposed is a free log retrieval operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterOutputProposed(opts *bind.FilterOpts, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (*OasysL2OutputOracleOutputProposedIterator, error) {

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

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleOutputProposedIterator{contract: _OasysL2OutputOracle.contract, event: "OutputProposed", logs: logs, sub: sub}, nil
}

// WatchOutputProposed is a free log subscription operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchOutputProposed(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleOutputProposed, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleOutputProposed)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputProposed", log); err != nil {
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
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseOutputProposed(log types.Log) (*OasysL2OutputOracleOutputProposed, error) {
	event := new(OasysL2OutputOracleOutputProposed)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleOutputVerifiedIterator is returned from FilterOutputVerified and is used to iterate over the raw logs and unpacked data for OutputVerified events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputVerifiedIterator struct {
	Event *OasysL2OutputOracleOutputVerified // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleOutputVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleOutputVerified)
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
		it.Event = new(OasysL2OutputOracleOutputVerified)
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
func (it *OasysL2OutputOracleOutputVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleOutputVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleOutputVerified represents a OutputVerified event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputVerified struct {
	L2OutputIndex *big.Int
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputVerified is a free log retrieval operation binding the contract event 0x98beebb56f3d9f23ca4c2e85cac607b80e030b3ca74fcc7b38f98bdd5146ef6b.
//
// Solidity: event OutputVerified(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterOutputVerified(opts *bind.FilterOpts, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (*OasysL2OutputOracleOutputVerifiedIterator, error) {

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

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "OutputVerified", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleOutputVerifiedIterator{contract: _OasysL2OutputOracle.contract, event: "OutputVerified", logs: logs, sub: sub}, nil
}

// WatchOutputVerified is a free log subscription operation binding the contract event 0x98beebb56f3d9f23ca4c2e85cac607b80e030b3ca74fcc7b38f98bdd5146ef6b.
//
// Solidity: event OutputVerified(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchOutputVerified(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleOutputVerified, l2OutputIndex []*big.Int, outputRoot [][32]byte, l2BlockNumber []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "OutputVerified", l2OutputIndexRule, outputRootRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleOutputVerified)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputVerified", log); err != nil {
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
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseOutputVerified(log types.Log) (*OasysL2OutputOracleOutputVerified, error) {
	event := new(OasysL2OutputOracleOutputVerified)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OasysL2OutputOracleOutputsDeletedIterator is returned from FilterOutputsDeleted and is used to iterate over the raw logs and unpacked data for OutputsDeleted events raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputsDeletedIterator struct {
	Event *OasysL2OutputOracleOutputsDeleted // Event containing the contract specifics and raw log

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
func (it *OasysL2OutputOracleOutputsDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OasysL2OutputOracleOutputsDeleted)
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
		it.Event = new(OasysL2OutputOracleOutputsDeleted)
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
func (it *OasysL2OutputOracleOutputsDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OasysL2OutputOracleOutputsDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OasysL2OutputOracleOutputsDeleted represents a OutputsDeleted event raised by the OasysL2OutputOracle contract.
type OasysL2OutputOracleOutputsDeleted struct {
	PrevNextOutputIndex *big.Int
	NewNextOutputIndex  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterOutputsDeleted is a free log retrieval operation binding the contract event 0x4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b6.
//
// Solidity: event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) FilterOutputsDeleted(opts *bind.FilterOpts, prevNextOutputIndex []*big.Int, newNextOutputIndex []*big.Int) (*OasysL2OutputOracleOutputsDeletedIterator, error) {

	var prevNextOutputIndexRule []interface{}
	for _, prevNextOutputIndexItem := range prevNextOutputIndex {
		prevNextOutputIndexRule = append(prevNextOutputIndexRule, prevNextOutputIndexItem)
	}
	var newNextOutputIndexRule []interface{}
	for _, newNextOutputIndexItem := range newNextOutputIndex {
		newNextOutputIndexRule = append(newNextOutputIndexRule, newNextOutputIndexItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.FilterLogs(opts, "OutputsDeleted", prevNextOutputIndexRule, newNextOutputIndexRule)
	if err != nil {
		return nil, err
	}
	return &OasysL2OutputOracleOutputsDeletedIterator{contract: _OasysL2OutputOracle.contract, event: "OutputsDeleted", logs: logs, sub: sub}, nil
}

// WatchOutputsDeleted is a free log subscription operation binding the contract event 0x4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b6.
//
// Solidity: event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex)
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) WatchOutputsDeleted(opts *bind.WatchOpts, sink chan<- *OasysL2OutputOracleOutputsDeleted, prevNextOutputIndex []*big.Int, newNextOutputIndex []*big.Int) (event.Subscription, error) {

	var prevNextOutputIndexRule []interface{}
	for _, prevNextOutputIndexItem := range prevNextOutputIndex {
		prevNextOutputIndexRule = append(prevNextOutputIndexRule, prevNextOutputIndexItem)
	}
	var newNextOutputIndexRule []interface{}
	for _, newNextOutputIndexItem := range newNextOutputIndex {
		newNextOutputIndexRule = append(newNextOutputIndexRule, newNextOutputIndexItem)
	}

	logs, sub, err := _OasysL2OutputOracle.contract.WatchLogs(opts, "OutputsDeleted", prevNextOutputIndexRule, newNextOutputIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OasysL2OutputOracleOutputsDeleted)
				if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputsDeleted", log); err != nil {
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
func (_OasysL2OutputOracle *OasysL2OutputOracleFilterer) ParseOutputsDeleted(log types.Log) (*OasysL2OutputOracleOutputsDeleted, error) {
	event := new(OasysL2OutputOracleOutputsDeleted)
	if err := _OasysL2OutputOracle.contract.UnpackLog(event, "OutputsDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
