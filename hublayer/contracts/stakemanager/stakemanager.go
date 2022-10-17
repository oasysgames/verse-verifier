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
)

// StakemanagerMetaData contains all meta data concerning the Stakemanager contract.
var StakemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyJoined\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountMismatched\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBlockProducer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotLastBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAsOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedValidator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"ValidatorActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"ValidatorDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"until\",\"type\":\"uint256\"}],\"name\":\"ValidatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorSlashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"activateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowlist\",\"outputs\":[{\"internalType\":\"contractIAllowlist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"claimCommissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"claimUnstakes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"environment\",\"outputs\":[{\"internalType\":\"contractIEnvironment\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getBlockAndSlashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getCommissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commissions\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getStakerStakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"oasStakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"woasStakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"soasStakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getTotalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amounts\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getUnstakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"oasUnstakes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"woasUnstakes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"soasUnstakes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"candidate\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getValidatorOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getValidatorStakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"howMany\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"candidates\",\"type\":\"bool[]\"},{\"internalType\":\"uint256\",\"name\":\"newCursor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEnvironment\",\"name\":\"_environment\",\"type\":\"address\"},{\"internalType\":\"contractIAllowlist\",\"name\":\"_allowlist\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"joinValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeUpdates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimCommission\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506179bc80620000216000396000f3fe6080604052600436106101ee5760003560e01c80636b2b33691161010d578063ad71bd36116100a0578063e1aca3411161006f578063e1aca341146107cf578063f3621e43146107f8578063f65a5ed214610821578063f8d6b1ab1461085e578063fa52c7d814610887576101ee565b8063ad71bd36146106ea578063cbc0fac614610728578063d1f18ee114610751578063dbd61d8714610792576101ee565b80637befa74f116100dc5780637befa74f1461062957806388325234146106455780639168ae7214610684578063ac7475ed146106c1576101ee565b80636b2b336914610557578063724319911461058057806374e2b63c146105c15780637b520aa8146105ec576101ee565b80632b42ed8c1161018557806345367f231161015457806345367f231461047557806346dfce7b146104b2578063485cc955146104f15780635efc766e1461051a576101ee565b80632b42ed8c146103a35780632b47da52146103e457806333f32d781461040f578063428e85621461044c576101ee565b8063195afea1116101c1578063195afea1146102ad5780631c1b4f3a146102ea5780632168e8b4146103275780632222636714610365576101ee565b806302fb4d85146101f3578063158ef93e1461021c5780631903cf1614610247578063190b925714610270575b600080fd5b3480156101ff57600080fd5b5061021a60048036038101906102159190616753565b6108c6565b005b34801561022857600080fd5b50610231610d1c565b60405161023e91906167ae565b60405180910390f35b34801561025357600080fd5b5061026e60048036038101906102699190616922565b610d2d565b005b34801561027c57600080fd5b506102976004803603810190610292919061697e565b611122565b6040516102a491906169ba565b60405180910390f35b3480156102b957600080fd5b506102d460048036038101906102cf9190616753565b611146565b6040516102e191906169ba565b60405180910390f35b3480156102f657600080fd5b50610311600480360381019061030c919061697e565b6111c8565b60405161031e91906169ba565b60405180910390f35b34801561033357600080fd5b5061034e600480360381019061034991906169d5565b6111ec565b60405161035c929190616ad3565b60405180910390f35b34801561037157600080fd5b5061038c60048036038101906103879190616753565b611310565b60405161039a929190616b03565b60405180910390f35b3480156103af57600080fd5b506103ca60048036038101906103c59190616b2c565b611412565b6040516103db959493929190616c51565b60405180910390f35b3480156103f057600080fd5b506103f96117f7565b6040516104069190616d1f565b60405180910390f35b34801561041b57600080fd5b5061043660048036038101906104319190616dfd565b61181d565b60405161044391906169ba565b60405180910390f35b34801561045857600080fd5b50610473600480360381019061046e9190616922565b611a40565b005b34801561048157600080fd5b5061049c6004803603810190610497919061697e565b611e35565b6040516104a991906169ba565b60405180910390f35b3480156104be57600080fd5b506104d960048036038101906104d49190616b2c565b611ef6565b6040516104e893929190616e59565b60405180910390f35b3480156104fd57600080fd5b5061051860048036038101906105139190616f1a565b6122a1565b005b34801561052657600080fd5b50610541600480360381019061053c919061697e565b6123eb565b60405161054e9190616f69565b60405180910390f35b34801561056357600080fd5b5061057e60048036038101906105799190616f84565b61242a565b005b34801561058c57600080fd5b506105a760048036038101906105a29190616fb1565b612631565b6040516105b89594939291906170c2565b60405180910390f35b3480156105cd57600080fd5b506105d6612b11565b6040516105e39190617152565b60405180910390f35b3480156105f857600080fd5b50610613600480360381019061060e9190616f84565b612b37565b6040516106209190616f69565b60405180910390f35b610643600480360381019061063e9190617192565b612b6a565b005b34801561065157600080fd5b5061066c60048036038101906106679190616f84565b613021565b60405161067b939291906171e5565b60405180910390f35b34801561069057600080fd5b506106ab60048036038101906106a69190616f84565b61311e565b6040516106b89190616f69565b60405180910390f35b3480156106cd57600080fd5b506106e860048036038101906106e39190616f84565b61315c565b005b3480156106f657600080fd5b50610711600480360381019061070c91906169d5565b6132f9565b60405161071f929190616ad3565b60405180910390f35b34801561073457600080fd5b5061074f600480360381019061074a9190616753565b61341d565b005b34801561075d57600080fd5b5061077860048036038101906107739190616753565b613561565b60405161078995949392919061721c565b60405180910390f35b34801561079e57600080fd5b506107b960048036038101906107b4919061726f565b613783565b6040516107c691906169ba565b60405180910390f35b3480156107db57600080fd5b506107f660048036038101906107f19190617192565b613846565b005b34801561080457600080fd5b5061081f600480360381019061081a919061726f565b613cba565b005b34801561082d57600080fd5b506108486004803603810190610843919061697e565b613f0a565b6040516108559190616f69565b60405180910390f35b34801561086a57600080fd5b5061088560048036038101906108809190616f84565b613f49565b005b34801561089357600080fd5b506108ae60048036038101906108a99190616f84565b61408a565b6040516108bd939291906172c2565b60405180910390f35b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156109ef576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a54576040517f1cf4735900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060046000600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000610c30600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633fa4f2456040518163ffffffff1660e01b815260040161012060405180830381865afa158015610b69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8d91906173f3565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1e9190617421565b86856140f4909392919063ffffffff16565b90508160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda60405160405180910390a26000811115610d15578160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167feb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e9582580282604051610d0c91906169ba565b60405180910390a25b5050505050565b60008054906101000a900460ff1681565b816000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610e2157508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610e58576040517f0809490800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600073ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610f22576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fb3919061747a565b15610fea576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110cd600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561105a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061107e9190617421565b85600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061421a9092919063ffffffff16565b8473ffffffffffffffffffffffffffffffffffffffff167fc11dfc9c24621433bb10587dc4bbae26a33a4aff53914e0d4c9fddf224a8cb7d8560405161111391906174a7565b60405180910390a25050505050565b6002818154811061113257600080fd5b906000526020600020016000915090505481565b60006111bd600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061422c9092919063ffffffff16565b508091505092915050565b600381815481106111d857600080fd5b906000526020600020016000915090505481565b6060600061120084846005805490506143ed565b80925081945050508267ffffffffffffffff811115611222576112216167df565b5b6040519080825280602002602001820160405280156112505781602001602082028036833780820191505090505b50915060005b83811015611308576005818661126c91906174f8565b8154811061127d5761127c61754e565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168382815181106112bb576112ba61754e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806113009061757d565b915050611256565b509250929050565b600080611403600084116113b457600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561138b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113af9190617421565b6113b6565b835b600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061442a90919063ffffffff16565b80925081935050509250929050565b606080606080600080600760008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600089116114fa57600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156114d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114f59190617421565b6114fc565b885b985061150e88886005805490506143ed565b80935081985050508667ffffffffffffffff8111156115305761152f6167df565b5b60405190808252806020026020018201604052801561155e5781602001602082028036833780820191505090505b5095508667ffffffffffffffff81111561157b5761157a6167df565b5b6040519080825280602002602001820160405280156115a95781602001602082028036833780820191505090505b5094508667ffffffffffffffff8111156115c6576115c56167df565b5b6040519080825280602002602001820160405280156115f45781602001602082028036833780820191505090505b5093508667ffffffffffffffff811115611611576116106167df565b5b60405190808252806020026020018201604052801561163f5781602001602082028036833780820191505090505b50925060005b878110156117ea576005818a61165b91906174f8565b8154811061166c5761166b61754e565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168782815181106116aa576116a961754e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506117168782815181106116fa576116f961754e565b5b602002602001015160008c85614464909392919063ffffffff16565b8682815181106117295761172861754e565b5b60200260200101818152505061176787828151811061174b5761174a61754e565b5b602002602001015160018c85614464909392919063ffffffff16565b85828151811061177a5761177961754e565b5b6020026020010181815250506117b887828151811061179c5761179b61754e565b5b602002602001015160028c85614464909392919063ffffffff16565b8482815181106117cb576117ca61754e565b5b60200260200101818152505080806117e29061757d565b915050611645565b5050945094509450945094565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600183600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611890573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118b49190617421565b6118be91906175c6565b6118c891906175c6565b905060008451905060005b84811015611a37576001836118e891906174f8565b925060008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fcbb371b856040518263ffffffff1660e01b815260040161194691906169ba565b61012060405180830381865afa158015611964573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061198891906173f3565b905060005b83811015611a2257611a028286600460008c86815181106119b1576119b061754e565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061456f9092919063ffffffff16565b86611a0d91906174f8565b95508080611a1a9061757d565b91505061198d565b50508080611a2f9061757d565b9150506118d3565b50505092915050565b816000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015611b3457508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15611b6b576040517f0809490800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600073ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611c35576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ca2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cc6919061747a565b15611cfd576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611de0600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d919190617421565b85600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206146b99092919063ffffffff16565b8473ffffffffffffffffffffffffffffffffffffffff167f0ad9bf1b8c026a174a2f30954417002a6ea00c9b08c1b8846c7951c687be809585604051611e2691906174a7565b60405180910390a25050505050565b6000808211611ed457600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ecf9190617421565b611ed6565b815b9150611eef60038360026146cb9092919063ffffffff16565b9050919050565b606080600080600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008711611fdb57600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fd69190617421565b611fdd565b865b9650611ff1868683600701805490506143ed565b80935081965050508467ffffffffffffffff811115612013576120126167df565b5b6040519080825280602002602001820160405280156120415781602001602082028036833780820191505090505b5093508467ffffffffffffffff81111561205e5761205d6167df565b5b60405190808252806020026020018201604052801561208c5781602001602082028036833780820191505090505b50925060005b858110156122955760006007600084600701848b6120b091906174f8565b815481106120c1576120c061754e565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168683815181106121605761215f61754e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506121d68360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660028b84614464909392919063ffffffff16565b6122128460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660018c85614464909392919063ffffffff16565b61224e8560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660008d86614464909392919063ffffffff16565b61225891906174f8565b61226291906174f8565b8583815181106122755761227461754e565b5b60200260200101818152505050808061228d9061757d565b915050612092565b50509450945094915050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612306576040517f1cf4735900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900460ff161561234b576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff02191690831515021790555081600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600581815481106123fb57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663322433e3336040518263ffffffff1660e01b81526004016124859190616f69565b602060405180830381865afa1580156124a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124c6919061747a565b6124fc576040517f8460af8a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61254d81600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206147ad90919063ffffffff16565b6005339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b606080606080600080600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156126a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126cb9190617421565b9050600089116126db57806126dd565b885b985060008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fcbb371b8b6040518263ffffffff1660e01b815260040161273b91906169ba565b61012060405180830381865afa158015612759573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061277d91906173f3565b905061278f89896005805490506143ed565b80945081995050508767ffffffffffffffff8111156127b1576127b06167df565b5b6040519080825280602002602001820160405280156127df5781602001602082028036833780820191505090505b5096508767ffffffffffffffff8111156127fc576127fb6167df565b5b60405190808252806020026020018201604052801561282a5781602001602082028036833780820191505090505b5095508767ffffffffffffffff811115612847576128466167df565b5b6040519080825280602002602001820160405280156128755781602001602082028036833780820191505090505b5094508767ffffffffffffffff811115612892576128916167df565b5b6040519080825280602002602001820160405280156128c05781602001602082028036833780820191505090505b50935060005b88811015612b03576000600460006005848e6128e291906174f8565b815481106128f3576128f261754e565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168983815181106129925761299161754e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16888381518110612a0457612a0361754e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050612a518c8261488790919063ffffffff16565b878381518110612a6457612a6361754e565b5b602002602001018181525050612a838c826148ac90919063ffffffff16565b158015612aa05750612a9e8c826148d990919063ffffffff16565b155b8015612aca57508260c00151878381518110612abf57612abe61754e565b5b602002602001015110155b868381518110612add57612adc61754e565b5b602002602001019015159081151581525050508080612afb9061757d565b9150506128c6565b505050939792965093509350565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60066020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b82600073ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415612c34576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ca1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cc5919061747a565b15612cfc576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415612d37576040517ff792180a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612dec60036001600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612dab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dcf9190617421565b612dd991906174f8565b846002614906909392919063ffffffff16565b612df783338461495b565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415612f3a57338160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506008339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b612fb3600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020868685614ad590949392919063ffffffff16565b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8fc656e319452025372383dc27d933046d412b8253de50a10739eeaa59862be68686604051613012929190617671565b60405180910390a35050505050565b600080600080600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506130a0600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600083614cda9092919063ffffffff16565b93506130da600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600183614cda9092919063ffffffff16565b9250613114600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600283614cda9092919063ffffffff16565b9150509193909250565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081565b33600073ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613226576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61327782600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020614f2990919063ffffffff16565b33600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6060600061330d84846008805490506143ed565b80925081945050508267ffffffffffffffff81111561332f5761332e6167df565b5b60405190808252806020026020018201604052801561335d5781602001602082028036833780820191505090505b50915060005b83811015613415576008818661337991906174f8565b8154811061338a5761338961754e565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168382815181106133c8576133c761754e565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808061340d9061757d565b915050613363565b509250929050565b81600073ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156134e7576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61355c600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206150619092919063ffffffff16565b505050565b600080600080600080600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156135d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135fb9190617421565b90506000871161360b578061360d565b865b965060008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fcbb371b896040518263ffffffff1660e01b815260040161366b91906169ba565b61012060405180830381865afa158015613689573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136ad91906173f3565b90506000600460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905061370589826148ac90919063ffffffff16565b15965061371b89826148d990919063ffffffff16565b9550613730898261488790919063ffffffff16565b935086801561373d575085155b801561374d57508160c001518410155b94508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1697505050509295509295909350565b600061383a600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084600760008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206150bd909392919063ffffffff16565b50809150509392505050565b82600073ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613910576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b33600073ffffffffffffffffffffffffffffffffffffffff16600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156139da576040517fcf83d93d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a6b919061747a565b15613aa2576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000831415613add576040517ff792180a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613b9260036001600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613b51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b759190617421565b613b7f91906174f8565b856002615398909392919063ffffffff16565b50613c4a600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208686600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061544590949392919063ffffffff16565b92508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ff2812c3df2511a467cbe777b1ee98b1ddb9918bb0f09568a269d2fb58233cb528686604051613cab929190617671565b60405180910390a35050505050565b81600073ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613d84576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600073ffffffffffffffffffffffffffffffffffffffff16600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613e4e576040517fcf83d93d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613f03600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085600760008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061574c909392919063ffffffff16565b5050505050565b60088181548110613f1a57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600073ffffffffffffffffffffffffffffffffffffffff16600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415614013576040517fcf83d93d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614086600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061580b90919063ffffffff16565b5050565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060154905083565b60008085600901600085815260200190815260200160002054141561412e5781856009016000858152602001908152602001600020819055505b6000600186600a0160008681526020019081526020016000205461415291906174f8565b90508086600a016000868152602001908152602001600020819055508460e0015181101580156141ae575085600301600060018661419091906174f8565b815260200190815260200160002060009054906101000a900460ff16155b1561421157846101000151846141c491906174f8565b91505b818410156142105783806141da9061757d565b945050600186600301600086815260200190815260200160002060006101000a81548160ff0219169083151502179055506141c7565b5b50949350505050565b6142278383836000615833565b505050565b60008084600601549050600060018573ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614285573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142a99190617421565b6142b391906175c6565b905060008414806142ce57508082856142cc91906174f8565b115b156142e25781816142df91906175c6565b93505b60005b848110156143e3576001836142fa91906174f8565b925060008673ffffffffffffffffffffffffffffffffffffffff1663fcbb371b856040518263ffffffff1660e01b815260040161433791906169ba565b61012060405180830381865afa158015614355573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061437991906173f3565b9050600061438889838761456f565b9050600081141561439a5750506143d0565b60008260a0015114156143ae5750506143d0565b6143c0818360a00151606460196158e7565b866143cb91906174f8565b955050505b80806143db9061757d565b9150506142e5565b5050935093915050565b6000808284866143fd91906174f8565b1061441157848361440e91906175c6565b93505b83848661441e91906174f8565b91509150935093915050565b6000808360090160008481526020019081526020016000205484600a01600085815260200190815260200160002054915091509250929050565b6000614565856002016000856002811115614482576144816175fa565b5b6002811115614494576144936175fa565b5b815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020838760010160008760028111156144f9576144f86175fa565b5b600281111561450b5761450a6175fa565b5b815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206146cb9092919063ffffffff16565b9050949350505050565b600061457b84836148ac565b8061458c575061458b84836148d9565b5b1561459a57600090506146b2565b60006145a68584614887565b905060008114156145bb5760009150506146b2565b60006019600a6145cb91906177cd565b6145dc866080015160646019615920565b836145e79190617818565b6145f191906178a1565b90506000811415614607576000925050506146b2565b61462a8560600151866040015161461e9190617818565b6301e133806019615920565b816146359190617818565b90506019600a61464591906177cd565b8161465091906178a1565b9050600086600a01600086815260200190815260200160002054905060008111156146ab5760008760090160008781526020019081526020016000205490506146a783838361469f91906175c6565b8360196158e7565b9250505b8193505050505b9392505050565b6146c68383836001615833565b505050565b600080848054905090506000811480614701575082856000815481106146f4576146f361754e565b5b9060005260206000200154115b156147105760009150506147a6565b828560018361471f91906175c6565b815481106147305761472f61754e565b5b906000526020600020015411614772578360018261474e91906175c6565b8154811061475f5761475e61754e565b5b90600052602060002001549150506147a6565b60006147818685600085615977565b90508481815481106147965761479561754e565b5b9060005260206000200154925050505b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff168260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614836576040517e3b268200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b338260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506148838282614f29565b5050565b60006148a48360050183856004016146cb9092919063ffffffff16565b905092915050565b600082600201600083815260200190815260200160002060009054906101000a900460ff16905092915050565b600082600301600083815260200190815260200160002060009054906101000a900460ff16905092915050565b614911848484615a3c565b80836001868054905061492491906175c6565b815481106149355761493461754e565b5b90600052602060002001600082825461494e91906174f8565b9250508190555050505050565b6000600281111561496f5761496e6175fa565b5b836002811115614982576149816175fa565b5b14156149c6578034146149c1576040517f1fcb60ca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614ad0565b60003414614a00576040517fa745ac8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000614a0b84615b4f565b73ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401614a47939291906172c2565b6020604051808303816000875af1158015614a66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a8a919061747a565b905080614ace57836040517f1b6a68ea000000000000000000000000000000000000000000000000000000008152600401614ac591906178d2565b60405180910390fd5b505b505050565b614c98856002016000846002811115614af157614af06175fa565b5b6002811115614b0357614b026175fa565b5b815260200190815260200160002060008560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060018673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614bbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614be39190617421565b614bed91906174f8565b83886001016000876002811115614c0757614c066175fa565b5b6002811115614c1957614c186175fa565b5b815260200190815260200160002060008860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020614906909392919063ffffffff16565b614cd3848660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168386615c19909392919063ffffffff16565b5050505050565b600080846003016000846002811115614cf657614cf56175fa565b5b6002811115614d0857614d076175fa565b5b81526020019081526020016000208054905090506000811415614d2f576000915050614f22565b60008473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614d7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614da09190617421565b90506000600183614db191906175c6565b9050600081118015614e17575081876003016000876002811115614dd857614dd76175fa565b5b6002811115614dea57614de96175fa565b5b81526020019081526020016000208281548110614e0a57614e0961754e565b5b9060005260206000200154115b15614e2b578080614e27906178ed565b9150505b81876003016000876002811115614e4557614e446175fa565b5b6002811115614e5757614e566175fa565b5b81526020019081526020016000208281548110614e7757614e7661754e565b5b90600052602060002001541115614e945760009350505050614f22565b600080600090505b828111614f1957886004016000886002811115614ebc57614ebb6175fa565b5b6002811115614ece57614ecd6175fa565b5b81526020019081526020016000208181548110614eee57614eed61754e565b5b906000526020600020015482614f0491906174f8565b91508080614f119061757d565b915050614e9c565b50809450505050505b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415614f90576040517f7138356f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561501a576040517fe037058f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b808260010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60008061506f85858561422c565b9150915080856006018190555060008211156150b6576150b560008660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684615dc9565b5b5050505050565b6000808560050160008560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600060018673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615177573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061519b9190617421565b6151a591906175c6565b905060008414806151c057508082856151be91906174f8565b115b156151d45781816151d191906175c6565b93505b60005b8481101561538d576001836151ec91906174f8565b92506000615221898860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600287614464565b6152528a8960000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600188614464565b6152838b8a60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600089614464565b61528d91906174f8565b61529791906174f8565b905060008114156152a8575061537a565b60006153398973ffffffffffffffffffffffffffffffffffffffff1663fcbb371b876040518263ffffffff1660e01b81526004016152e691906169ba565b61012060405180830381865afa158015615304573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061532891906173f3565b868a615f399092919063ffffffff16565b9050600081141561534b57505061537a565b61536a8183615363888c61488790919063ffffffff16565b60196158e7565b8661537591906174f8565b955050505b80806153859061757d565b9150506151d7565b505094509492505050565b60006153a5858585615a3c565b6000858054905090506000856001836153be91906175c6565b815481106153cf576153ce61754e565b5b90600052602060002001549050808411156153ea57806153ec565b835b9350600084111561543857838660018461540691906175c6565b815481106154175761541661754e565b5b90600052602060002001600082825461543091906175c6565b925050819055505b8392505050949350505050565b6000808573ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615493573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906154b79190617421565b905060006154eb888760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168785614464565b9050600061552b898860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168860018761552691906174f8565b614464565b9050615682896002016000886002811115615549576155486175fa565b5b600281111561555b5761555a6175fa565b5b815260200190815260200160002060008960000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001856155d791906174f8565b878c60010160008b60028111156155f1576155f06175fa565b5b6002811115615603576156026175fa565b5b815260200190815260200160002060008c60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615398909392919063ffffffff16565b945060008514156156995760009350505050615743565b6156ae888689615f9a9092919063ffffffff16565b60008590506000838311156156eb5783836156c991906175c6565b90508087106156d857806156da565b865b905080826156e891906175c6565b91505b6000821115615701576157008b8b8a85616037565b5b600081111561573a57615739888c60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683615dc9565b5b86955050505050505b95945050505050565b60008061575b868686866150bd565b91509150808660050160008660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008211156158035761580260008760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684615dc9565b5b505050505050565b615817828260016162b4565b615823828260026162b4565b61582f828260006162b4565b5050565b60008251905060005b818110156158df5760008482815181106158595761585861754e565b5b602002602001015190508581118015615898575083151587600201600083815260200190815260200160002060009054906101000a900460ff16151514155b156158cb578387600201600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5080806158d79061757d565b91505061583c565b505050505050565b600081600a6158f691906177cd565b615901858585615920565b8661590c9190617818565b61591691906178a1565b9050949350505050565b60008060018361593091906174f8565b600a61593c91906177cd565b856159479190617818565b9050600a6005858361595991906178a1565b61596391906174f8565b61596d91906178a1565b9150509392505050565b6000818314156159955760018261598e91906175c6565b9050615a34565b6000600283856159a591906174f8565b6159af91906178a1565b9050848682815481106159c5576159c461754e565b5b906000526020600020015411156159ea576159e286868684615977565b915050615a34565b848682815481106159fe576159fd61754e565b5b90600052602060002001541015615a2f57615a278686600184615a2191906174f8565b86615977565b915050615a34565b809150505b949350505050565b6000838054905090506000811415615a9857838290806001815401808255809150506001900390600052602060002001600090919091909150558260018160018154018082558091505003906000526020600020505050615b4a565b600084600183615aa891906175c6565b81548110615ab957615ab861754e565b5b90600052602060002001549050828114615b4757848390806001815401808255809150506001900390600052602060002001600090919091909150558384600184615b0491906175c6565b81548110615b1557615b1461754e565b5b906000526020600020015490806001815401808255809150506001900390600052602060002001600090919091909150555b50505b505050565b600060016002811115615b6557615b646175fa565b5b826002811115615b7857615b776175fa565b5b1415615b9a577352000000000000000000000000000000000000019050615c14565b600280811115615bad57615bac6175fa565b5b826002811115615bc057615bbf6175fa565b5b1415615be2577352000000000000000000000000000000000000029050615c14565b6040517f8698bf3700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b8360080160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16615d2c5760018460080160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555083600701829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b615dc38460050160018573ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615d80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615da49190617421565b615dae91906174f8565b8387600401614906909392919063ffffffff16565b50505050565b6000806002811115615dde57615ddd6175fa565b5b846002811115615df157615df06175fa565b5b1415615e68578273ffffffffffffffffffffffffffffffffffffffff1682604051615e1b90617948565b60006040518083038185875af1925050503d8060008114615e58576040519150601f19603f3d011682016040523d82523d6000602084013e615e5d565b606091505b505080915050615ef1565b615e7184615b4f565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b8152600401615eab92919061795d565b6020604051808303816000875af1158015615eca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615eee919061747a565b90505b80615f3357836040517f1b6a68ea000000000000000000000000000000000000000000000000000000008152600401615f2a91906178d2565b60405180910390fd5b50505050565b600080615f4785858561456f565b90506000811415615f5c576000915050615f93565b60008460a001511415615f725780915050615f93565b615f84818560a00151606460196158e7565b81615f8f91906175c6565b9150505b9392505050565b6160318360050160018473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615fee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906160129190617421565b61601c91906174f8565b8386600401615398909392919063ffffffff16565b50505050565b600060018473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015616086573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906160aa9190617421565b6160b491906174f8565b905060008560030160008560028111156160d1576160d06175fa565b5b60028111156160e3576160e26175fa565b5b8152602001908152602001600020805490509050600081148061616757508186600301600086600281111561611b5761611a6175fa565b5b600281111561612d5761612c6175fa565b5b815260200190815260200160002060018361614891906175c6565b815481106161595761615861754e565b5b906000526020600020015414155b1561623357856003016000856002811115616185576161846175fa565b5b6002811115616197576161966175fa565b5b81526020019081526020016000208290806001815401808255809150506001900390600052602060002001600090919091909150558560040160008560028111156161e5576161e46175fa565b5b60028111156161f7576161f66175fa565b5b815260200190815260200160002083908060018154018082558091505060019003906000526020600020016000909190919091505550506162ae565b8286600401600086600281111561624d5761624c6175fa565b5b600281111561625f5761625e6175fa565b5b815260200190815260200160002060018361627a91906175c6565b8154811061628b5761628a61754e565b5b9060005260206000200160008282546162a491906174f8565b9250508190555050505b50505050565b60006162c1848484614cda565b905060008114156162d2575061661b565b60008460030160008460028111156162ed576162ec6175fa565b5b60028111156162ff576162fe6175fa565b5b81526020019081526020016000208054905090508373ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561635e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906163829190617421565b85600301600085600281111561639b5761639a6175fa565b5b60028111156163ad576163ac6175fa565b5b81526020019081526020016000206001836163c891906175c6565b815481106163d9576163d861754e565b5b90600052602060002001541161647857846003016000846002811115616402576164016175fa565b5b6002811115616414576164136175fa565b5b8152602001908152602001600020600061642e9190616620565b846004016000846002811115616447576164466175fa565b5b6002811115616459576164586175fa565b5b815260200190815260200160002060006164739190616620565b6165e9565b604051806020016040528086600301600086600281111561649c5761649b6175fa565b5b60028111156164ae576164ad6175fa565b5b81526020019081526020016000206001846164c991906175c6565b815481106164da576164d961754e565b5b9060005260206000200154815250856003016000856002811115616501576165006175fa565b5b6002811115616513576165126175fa565b5b815260200190815260200160002090600161652f929190616641565b506040518060200160405280866004016000866002811115616554576165536175fa565b5b6002811115616566576165656175fa565b5b815260200190815260200160002060018461658191906175c6565b815481106165925761659161754e565b5b90600052602060002001548152508560040160008560028111156165b9576165b86175fa565b5b60028111156165cb576165ca6175fa565b5b81526020019081526020016000209060016165e7929190616641565b505b616618838660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684615dc9565b50505b505050565b508054600082559060005260206000209081019061663e919061668e565b50565b82805482825590600052602060002090810192821561667d579160200282015b8281111561667c578251825591602001919060010190616661565b5b50905061668a919061668e565b5090565b5b808211156166a757600081600090555060010161668f565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006166ea826166bf565b9050919050565b6166fa816166df565b811461670557600080fd5b50565b600081359050616717816166f1565b92915050565b6000819050919050565b6167308161671d565b811461673b57600080fd5b50565b60008135905061674d81616727565b92915050565b6000806040838503121561676a576167696166b5565b5b600061677885828601616708565b92505060206167898582860161673e565b9150509250929050565b60008115159050919050565b6167a881616793565b82525050565b60006020820190506167c3600083018461679f565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b616817826167ce565b810181811067ffffffffffffffff82111715616836576168356167df565b5b80604052505050565b60006168496166ab565b9050616855828261680e565b919050565b600067ffffffffffffffff821115616875576168746167df565b5b602082029050602081019050919050565b600080fd5b600061689e6168998461685a565b61683f565b905080838252602082019050602084028301858111156168c1576168c0616886565b5b835b818110156168ea57806168d6888261673e565b8452602084019350506020810190506168c3565b5050509392505050565b600082601f830112616909576169086167c9565b5b813561691984826020860161688b565b91505092915050565b60008060408385031215616939576169386166b5565b5b600061694785828601616708565b925050602083013567ffffffffffffffff811115616968576169676166ba565b5b616974858286016168f4565b9150509250929050565b600060208284031215616994576169936166b5565b5b60006169a28482850161673e565b91505092915050565b6169b48161671d565b82525050565b60006020820190506169cf60008301846169ab565b92915050565b600080604083850312156169ec576169eb6166b5565b5b60006169fa8582860161673e565b9250506020616a0b8582860161673e565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b616a4a816166df565b82525050565b6000616a5c8383616a41565b60208301905092915050565b6000602082019050919050565b6000616a8082616a15565b616a8a8185616a20565b9350616a9583616a31565b8060005b83811015616ac6578151616aad8882616a50565b9750616ab883616a68565b925050600181019050616a99565b5085935050505092915050565b60006040820190508181036000830152616aed8185616a75565b9050616afc60208301846169ab565b9392505050565b6000604082019050616b1860008301856169ab565b616b2560208301846169ab565b9392505050565b60008060008060808587031215616b4657616b456166b5565b5b6000616b5487828801616708565b9450506020616b658782880161673e565b9350506040616b768782880161673e565b9250506060616b878782880161673e565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b616bc88161671d565b82525050565b6000616bda8383616bbf565b60208301905092915050565b6000602082019050919050565b6000616bfe82616b93565b616c088185616b9e565b9350616c1383616baf565b8060005b83811015616c44578151616c2b8882616bce565b9750616c3683616be6565b925050600181019050616c17565b5085935050505092915050565b600060a0820190508181036000830152616c6b8188616a75565b90508181036020830152616c7f8187616bf3565b90508181036040830152616c938186616bf3565b90508181036060830152616ca78185616bf3565b9050616cb660808301846169ab565b9695505050505050565b6000819050919050565b6000616ce5616ce0616cdb846166bf565b616cc0565b6166bf565b9050919050565b6000616cf782616cca565b9050919050565b6000616d0982616cec565b9050919050565b616d1981616cfe565b82525050565b6000602082019050616d346000830184616d10565b92915050565b600067ffffffffffffffff821115616d5557616d546167df565b5b602082029050602081019050919050565b6000616d79616d7484616d3a565b61683f565b90508083825260208201905060208402830185811115616d9c57616d9b616886565b5b835b81811015616dc55780616db18882616708565b845260208401935050602081019050616d9e565b5050509392505050565b600082601f830112616de457616de36167c9565b5b8135616df4848260208601616d66565b91505092915050565b60008060408385031215616e1457616e136166b5565b5b600083013567ffffffffffffffff811115616e3257616e316166ba565b5b616e3e85828601616dcf565b9250506020616e4f8582860161673e565b9150509250929050565b60006060820190508181036000830152616e738186616a75565b90508181036020830152616e878185616bf3565b9050616e9660408301846169ab565b949350505050565b6000616ea9826166df565b9050919050565b616eb981616e9e565b8114616ec457600080fd5b50565b600081359050616ed681616eb0565b92915050565b6000616ee7826166df565b9050919050565b616ef781616edc565b8114616f0257600080fd5b50565b600081359050616f1481616eee565b92915050565b60008060408385031215616f3157616f306166b5565b5b6000616f3f85828601616ec7565b9250506020616f5085828601616f05565b9150509250929050565b616f63816166df565b82525050565b6000602082019050616f7e6000830184616f5a565b92915050565b600060208284031215616f9a57616f996166b5565b5b6000616fa884828501616708565b91505092915050565b600080600060608486031215616fca57616fc96166b5565b5b6000616fd88682870161673e565b9350506020616fe98682870161673e565b9250506040616ffa8682870161673e565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61703981616793565b82525050565b600061704b8383617030565b60208301905092915050565b6000602082019050919050565b600061706f82617004565b617079818561700f565b935061708483617020565b8060005b838110156170b557815161709c888261703f565b97506170a783617057565b925050600181019050617088565b5085935050505092915050565b600060a08201905081810360008301526170dc8188616a75565b905081810360208301526170f08187616a75565b905081810360408301526171048186616bf3565b905081810360608301526171188185617064565b905061712760808301846169ab565b9695505050505050565b600061713c82616cec565b9050919050565b61714c81617131565b82525050565b60006020820190506171676000830184617143565b92915050565b6003811061717a57600080fd5b50565b60008135905061718c8161716d565b92915050565b6000806000606084860312156171ab576171aa6166b5565b5b60006171b986828701616708565b93505060206171ca8682870161717d565b92505060406171db8682870161673e565b9150509250925092565b60006060820190506171fa60008301866169ab565b61720760208301856169ab565b61721460408301846169ab565b949350505050565b600060a0820190506172316000830188616f5a565b61723e602083018761679f565b61724b604083018661679f565b617258606083018561679f565b61726560808301846169ab565b9695505050505050565b600080600060608486031215617288576172876166b5565b5b600061729686828701616708565b93505060206172a786828701616708565b92505060406172b88682870161673e565b9150509250925092565b60006060820190506172d76000830186616f5a565b6172e46020830185616f5a565b6172f160408301846169ab565b949350505050565b600080fd5b60008151905061730d81616727565b92915050565b6000610120828403121561732a576173296172f9565b5b61733561012061683f565b90506000617345848285016172fe565b6000830152506020617359848285016172fe565b602083015250604061736d848285016172fe565b6040830152506060617381848285016172fe565b6060830152506080617395848285016172fe565b60808301525060a06173a9848285016172fe565b60a08301525060c06173bd848285016172fe565b60c08301525060e06173d1848285016172fe565b60e0830152506101006173e6848285016172fe565b6101008301525092915050565b6000610120828403121561740a576174096166b5565b5b600061741884828501617313565b91505092915050565b600060208284031215617437576174366166b5565b5b6000617445848285016172fe565b91505092915050565b61745781616793565b811461746257600080fd5b50565b6000815190506174748161744e565b92915050565b6000602082840312156174905761748f6166b5565b5b600061749e84828501617465565b91505092915050565b600060208201905081810360008301526174c18184616bf3565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006175038261671d565b915061750e8361671d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115617543576175426174c9565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006175888261671d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156175bb576175ba6174c9565b5b600182019050919050565b60006175d18261671d565b91506175dc8361671d565b9250828210156175ef576175ee6174c9565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6003811061763a576176396175fa565b5b50565b600081905061764b82617629565b919050565b600061765b8261763d565b9050919050565b61766b81617650565b82525050565b60006040820190506176866000830185617662565b61769360208301846169ab565b9392505050565b60008160011c9050919050565b6000808291508390505b60018511156176f1578086048111156176cd576176cc6174c9565b5b60018516156176dc5780820291505b80810290506176ea8561769a565b94506176b1565b94509492505050565b60008261770a57600190506177c6565b8161771857600090506177c6565b816001811461772e576002811461773857617767565b60019150506177c6565b60ff84111561774a576177496174c9565b5b8360020a915084821115617761576177606174c9565b5b506177c6565b5060208310610133831016604e8410600b841016171561779c5782820a905083811115617797576177966174c9565b5b6177c6565b6177a984848460016176a7565b925090508184048111156177c0576177bf6174c9565b5b81810290505b9392505050565b60006177d88261671d565b91506177e38361671d565b92506178107fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846176fa565b905092915050565b60006178238261671d565b915061782e8361671d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615617867576178666174c9565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006178ac8261671d565b91506178b78361671d565b9250826178c7576178c6617872565b5b828204905092915050565b60006020820190506178e76000830184617662565b92915050565b60006178f88261671d565b9150600082141561790c5761790b6174c9565b5b600182039050919050565b600081905092915050565b50565b6000617932600083617917565b915061793d82617922565b600082019050919050565b600061795382617925565b9150819050919050565b60006040820190506179726000830185616f5a565b61797f60208301846169ab565b939250505056fea2646970667358221220e44937d2f6845ac0c5064f33e62e24b574c5cc928890846ae5abeaab791b751e64736f6c634300080c0033",
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
	parsed, err := abi.JSON(strings.NewReader(StakemanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
// Solidity: function getValidatorInfo(address validator, uint256 epoch) view returns(address operator, bool active, bool jailed, bool candidate, uint256 stakes)
func (_Stakemanager *StakemanagerCaller) GetValidatorInfo(opts *bind.CallOpts, validator common.Address, epoch *big.Int) (struct {
	Operator  common.Address
	Active    bool
	Jailed    bool
	Candidate bool
	Stakes    *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorInfo", validator, epoch)

	outstruct := new(struct {
		Operator  common.Address
		Active    bool
		Jailed    bool
		Candidate bool
		Stakes    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Active = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Jailed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Candidate = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Stakes = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xd1f18ee1.
//
// Solidity: function getValidatorInfo(address validator, uint256 epoch) view returns(address operator, bool active, bool jailed, bool candidate, uint256 stakes)
func (_Stakemanager *StakemanagerSession) GetValidatorInfo(validator common.Address, epoch *big.Int) (struct {
	Operator  common.Address
	Active    bool
	Jailed    bool
	Candidate bool
	Stakes    *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorInfo(&_Stakemanager.CallOpts, validator, epoch)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xd1f18ee1.
//
// Solidity: function getValidatorInfo(address validator, uint256 epoch) view returns(address operator, bool active, bool jailed, bool candidate, uint256 stakes)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorInfo(validator common.Address, epoch *big.Int) (struct {
	Operator  common.Address
	Active    bool
	Jailed    bool
	Candidate bool
	Stakes    *big.Int
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

// GetValidatorStakes is a free data retrieval call binding the contract method 0x46dfce7b.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] _stakers, uint256[] stakes, uint256 newCursor)
func (_Stakemanager *StakemanagerCaller) GetValidatorStakes(opts *bind.CallOpts, validator common.Address, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Stakers   []common.Address
	Stakes    []*big.Int
	NewCursor *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorStakes", validator, epoch, cursor, howMany)

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

// GetValidatorStakes is a free data retrieval call binding the contract method 0x46dfce7b.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] _stakers, uint256[] stakes, uint256 newCursor)
func (_Stakemanager *StakemanagerSession) GetValidatorStakes(validator common.Address, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Stakers   []common.Address
	Stakes    []*big.Int
	NewCursor *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorStakes(&_Stakemanager.CallOpts, validator, epoch, cursor, howMany)
}

// GetValidatorStakes is a free data retrieval call binding the contract method 0x46dfce7b.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] _stakers, uint256[] stakes, uint256 newCursor)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorStakes(validator common.Address, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Stakers   []common.Address
	Stakes    []*big.Int
	NewCursor *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorStakes(&_Stakemanager.CallOpts, validator, epoch, cursor, howMany)
}

// GetValidators is a free data retrieval call binding the contract method 0x72431991.
//
// Solidity: function getValidators(uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] owners, address[] operators, uint256[] stakes, bool[] candidates, uint256 newCursor)
func (_Stakemanager *StakemanagerCaller) GetValidators(opts *bind.CallOpts, epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Owners     []common.Address
	Operators  []common.Address
	Stakes     []*big.Int
	Candidates []bool
	NewCursor  *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidators", epoch, cursor, howMany)

	outstruct := new(struct {
		Owners     []common.Address
		Operators  []common.Address
		Stakes     []*big.Int
		Candidates []bool
		NewCursor  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Operators = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Candidates = *abi.ConvertType(out[3], new([]bool)).(*[]bool)
	outstruct.NewCursor = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidators is a free data retrieval call binding the contract method 0x72431991.
//
// Solidity: function getValidators(uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] owners, address[] operators, uint256[] stakes, bool[] candidates, uint256 newCursor)
func (_Stakemanager *StakemanagerSession) GetValidators(epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Owners     []common.Address
	Operators  []common.Address
	Stakes     []*big.Int
	Candidates []bool
	NewCursor  *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidators(&_Stakemanager.CallOpts, epoch, cursor, howMany)
}

// GetValidators is a free data retrieval call binding the contract method 0x72431991.
//
// Solidity: function getValidators(uint256 epoch, uint256 cursor, uint256 howMany) view returns(address[] owners, address[] operators, uint256[] stakes, bool[] candidates, uint256 newCursor)
func (_Stakemanager *StakemanagerCallerSession) GetValidators(epoch *big.Int, cursor *big.Int, howMany *big.Int) (struct {
	Owners     []common.Address
	Operators  []common.Address
	Stakes     []*big.Int
	Candidates []bool
	NewCursor  *big.Int
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
// Solidity: function validators(address ) view returns(address owner, address operator, uint256 lastClaimCommission)
func (_Stakemanager *StakemanagerCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	LastClaimCommission *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Owner               common.Address
		Operator            common.Address
		LastClaimCommission *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.LastClaimCommission = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, uint256 lastClaimCommission)
func (_Stakemanager *StakemanagerSession) Validators(arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	LastClaimCommission *big.Int
}, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, uint256 lastClaimCommission)
func (_Stakemanager *StakemanagerCallerSession) Validators(arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	LastClaimCommission *big.Int
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

// ClaimCommissions is a paid mutator transaction binding the contract method 0xcbc0fac6.
//
// Solidity: function claimCommissions(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimCommissions(opts *bind.TransactOpts, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimCommissions", validator, epochs)
}

// ClaimCommissions is a paid mutator transaction binding the contract method 0xcbc0fac6.
//
// Solidity: function claimCommissions(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerSession) ClaimCommissions(validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimCommissions(&_Stakemanager.TransactOpts, validator, epochs)
}

// ClaimCommissions is a paid mutator transaction binding the contract method 0xcbc0fac6.
//
// Solidity: function claimCommissions(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimCommissions(validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimCommissions(&_Stakemanager.TransactOpts, validator, epochs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf3621e43.
//
// Solidity: function claimRewards(address staker, address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimRewards(opts *bind.TransactOpts, staker common.Address, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimRewards", staker, validator, epochs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf3621e43.
//
// Solidity: function claimRewards(address staker, address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerSession) ClaimRewards(staker common.Address, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimRewards(&_Stakemanager.TransactOpts, staker, validator, epochs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf3621e43.
//
// Solidity: function claimRewards(address staker, address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimRewards(staker common.Address, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimRewards(&_Stakemanager.TransactOpts, staker, validator, epochs)
}

// ClaimUnstakes is a paid mutator transaction binding the contract method 0xf8d6b1ab.
//
// Solidity: function claimUnstakes(address staker) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimUnstakes(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimUnstakes", staker)
}

// ClaimUnstakes is a paid mutator transaction binding the contract method 0xf8d6b1ab.
//
// Solidity: function claimUnstakes(address staker) returns()
func (_Stakemanager *StakemanagerSession) ClaimUnstakes(staker common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimUnstakes(&_Stakemanager.TransactOpts, staker)
}

// ClaimUnstakes is a paid mutator transaction binding the contract method 0xf8d6b1ab.
//
// Solidity: function claimUnstakes(address staker) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimUnstakes(staker common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimUnstakes(&_Stakemanager.TransactOpts, staker)
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
// Solidity: function unstake(address validator, uint8 token, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) Unstake(opts *bind.TransactOpts, validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "unstake", validator, token, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xe1aca341.
//
// Solidity: function unstake(address validator, uint8 token, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) Unstake(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, validator, token, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xe1aca341.
//
// Solidity: function unstake(address validator, uint8 token, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) Unstake(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, validator, token, amount)
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
