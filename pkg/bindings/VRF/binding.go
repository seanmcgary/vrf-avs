// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VRF

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// VRFRandomnessRequest is an auto generated low-level Go binding around an user-defined struct.
type VRFRandomnessRequest struct {
	Requester   common.Address
	TaskHash    [32]byte
	BlockNumber *big.Int
	Fulfilled   bool
	Result      *big.Int
}

// VRFMetaData contains all meta data concerning the VRF contract.
var VRFMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskMailbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"encodeTaskPayload\",\"inputs\":[{\"name\":\"randomnessType\",\"type\":\"uint8\",\"internalType\":\"enumVRF.RandomnessType\"},{\"name\":\"randomnessParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeVDFParams\",\"inputs\":[{\"name\":\"seed\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeVDFTaskPayload\",\"inputs\":[{\"name\":\"seed\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"executorOperatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRandomnessResult\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structVRF.RandomnessRequest\",\"components\":[{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequestCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onTaskCompleted\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestRandomness\",\"inputs\":[{\"name\":\"seed\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskHashToRequestId\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskMailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskMailbox\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorSet\",\"inputs\":[{\"name\":\"_executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"RandomnessFulfilled\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomnessRequested\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"seed\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidTaskResponse\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequestAlreadyFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequestNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051610fa0380380610fa083398101604081905261002e91610087565b6001600160a01b0391821660805280515f805460209093015163ffffffff16600160a01b026001600160c01b0319909316919093161717905561011b565b80516001600160a01b0381168114610082575f5ffd5b919050565b5f5f8284036060811215610099575f5ffd5b6100a28461006c565b92506040601f19820112156100b5575f5ffd5b50604080519081016001600160401b03811182821017156100e457634e487b7160e01b5f52604160045260245ffd5b6040526100f36020850161006c565b8152604084015163ffffffff8116811461010b575f5ffd5b6020820152919491935090915050565b608051610e5f6101415f395f8181610355015281816103ad01526107b70152610e5f5ff3fe608060405234801561000f575f5ffd5b50600436106100cb575f3560e01c8063ac8d073811610088578063e6dee7ed11610063578063e6dee7ed146102f5578063e7a923bb14610308578063f42a9e1314610350578063fdd154dc1461038f575f5ffd5b8063ac8d073814610267578063c58343ef14610287578063cdae4940146102e2575f5ffd5b80630944a904146100cf57806315505f8c146100e45780632026a4f3146101165780636a36dab91461015b57806381d12c58146101df578063a2785f0b1461025f575b5f5ffd5b6100e26100dd3660046109d6565b6103a2565b005b6101036100f2366004610a1e565b60036020525f908152604090205481565b6040519081526020015b60405180910390f35b5f54610137906001600160a01b03811690600160a01b900463ffffffff1682565b604080516001600160a01b03909316835263ffffffff90911660208301520161010d565b6101c8610169366004610a1e565b5f90815260026020818152604092839020835160a08101855281546001600160a01b031681526001820154928101929092529182015492810192909252600381015460ff16151560608301819052600490910154608090920182905291565b60408051921515835260208301919091520161010d565b61022b6101ed366004610a1e565b600260208190525f9182526040909120805460018201549282015460038301546004909301546001600160a01b039092169392909160ff9091169085565b604080516001600160a01b03909616865260208601949094529284019190915215156060830152608082015260a00161010d565b600154610103565b61027a610275366004610a35565b610509565b60405161010d9190610aa2565b61029a610295366004610a1e565b6105ad565b60405161010d919081516001600160a01b0316815260208083015190820152604080830151908201526060808301511515908201526080918201519181019190915260a00190565b61027a6102f0366004610a35565b61063e565b610103610303366004610a35565b6106a7565b6100e2610316366004610b29565b80515f805460209093015163ffffffff16600160a01b026001600160c01b03199093166001600160a01b0390921691909117919091179055565b6103777f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161010d565b61027a61039d366004610b90565b610908565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103eb57604051635c427cd960e01b815260040160405180910390fd5b5f838152600360205260408120549081900361041a57604051632589d98f60e11b815260040160405180910390fd5b5f818152600260205260409020600381015460ff161561044d5760405163533d99dd60e01b815260040160405180910390fd5b5f61045a84860186610bc6565b905060018151600181111561047157610471610c9d565b1461048f5760405163413041d160e01b815260040160405180910390fd5b5f81602001518060200190518101906104a89190610cb1565b60038401805460ff19166001179055805160048501819055604051919250889186917fbe3f52bb4df8f041a3e9118a0acf71ddeb18adc6fdb619f10c637273eab05e42916104f891815260200190565b60405180910390a350505050505050565b604080516020601f8401819004810282018301835281018381526060925f9291829187908790819085018382808284375f9201829052509390945250506040805180820190915292935091905080600181526020018360405160200161056f9190610cf3565b6040516020818303038152906040528152509050806040516020016105949190610d15565b6040516020818303038152906040529250505092915050565b6105e56040518060a001604052805f6001600160a01b031681526020015f81526020015f81526020015f151581526020015f81525090565b505f90815260026020818152604092839020835160a08101855281546001600160a01b031681526001820154928101929092529182015492810192909252600381015460ff161515606083015260040154608082015290565b604080516020601f8401819004810282018301835281018381526060925f9291829187908790819085018382808284375f92019190915250505091525060405190915061068f908290602001610cf3565b60405160208183030381529060405291505092915050565b5f60015f81546106b690610d57565b9182905550604080516020601f8601819004810282018301835281018581529293505f92909182919087908790819085018382808284375f920182905250939094525050604080518082019091529293509190508060018152602001836040516020016107239190610cf3565b60408051808303601f19018152918152915280516080810182523381525f60208083018290528351808501855282546001600160a01b0381168252600160a01b900463ffffffff1681830152838501529251939450929091606083019161078c91869101610d15565b60408051601f1981840301815291815291525162221dbd60e51b81529091505f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630443b7a0906107ec908590600401610d7b565b6020604051808303815f875af1158015610808573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061082c9190610de4565b6040805160a081018252338082526020808301858152438486019081525f60608601818152608087018281528e83526002808752898420985189546001600160a01b0319166001600160a01b039091161789559451600189015592519387019390935591516003808701805460ff1916921515929092179091559051600490950194909455858152929052908290208890559051919250829187907f152b260fdc6b51380aea47b7a68bd79722f0459b2603c789e9e9f476007e07be906108f6908c908c90610dfb565b60405180910390a45050505092915050565b60605f604051806040016040528086600181111561092857610928610c9d565b815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050915250604051909150610978908290602001610d15565b6040516020818303038152906040529150509392505050565b5f5f83601f8401126109a1575f5ffd5b50813567ffffffffffffffff8111156109b8575f5ffd5b6020830191508360208285010111156109cf575f5ffd5b9250929050565b5f5f5f604084860312156109e8575f5ffd5b83359250602084013567ffffffffffffffff811115610a05575f5ffd5b610a1186828701610991565b9497909650939450505050565b5f60208284031215610a2e575f5ffd5b5035919050565b5f5f60208385031215610a46575f5ffd5b823567ffffffffffffffff811115610a5c575f5ffd5b610a6885828601610991565b90969095509350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610ab46020830184610a74565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff81118282101715610af257610af2610abb565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610b2157610b21610abb565b604052919050565b5f6040828403128015610b3a575f5ffd5b50610b43610acf565b82356001600160a01b0381168114610b59575f5ffd5b8152602083013563ffffffff81168114610b71575f5ffd5b60208201529392505050565b803560028110610b8b575f5ffd5b919050565b5f5f5f60408486031215610ba2575f5ffd5b610bab84610b7d565b9250602084013567ffffffffffffffff811115610a05575f5ffd5b5f60208284031215610bd6575f5ffd5b813567ffffffffffffffff811115610bec575f5ffd5b820160408185031215610bfd575f5ffd5b610c05610acf565b610c0e82610b7d565b8152602082013567ffffffffffffffff811115610c29575f5ffd5b80830192505084601f830112610c3d575f5ffd5b813567ffffffffffffffff811115610c5757610c57610abb565b610c6a601f8201601f1916602001610af8565b818152866020838601011115610c7e575f5ffd5b816020850160208301375f602092820183015290820152949350505050565b634e487b7160e01b5f52602160045260245ffd5b5f6020828403128015610cc2575f5ffd5b506040516020810167ffffffffffffffff81118282101715610ce657610ce6610abb565b6040529151825250919050565b602081525f8251602080840152610d0d6040840182610a74565b949350505050565b602081525f825160028110610d3857634e487b7160e01b5f52602160045260245ffd5b806020840152506020830151604080840152610d0d6060840182610a74565b5f60018201610d7457634e487b7160e01b5f52601160045260245ffd5b5060010190565b6020815260018060a01b0382511660208201526bffffffffffffffffffffffff60208301511660408201525f604083015160018060a01b03815116606084015263ffffffff602082015116608084015250606083015160a080840152610d0d60c0840182610a74565b5f60208284031215610df4575f5ffd5b5051919050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea2646970667358221220853f73a7c5707a67682d1cb4a07d29605020784762600e2955587eec45cfe21a64736f6c634300081b0033",
}

// VRFABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFMetaData.ABI instead.
var VRFABI = VRFMetaData.ABI

// VRFBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VRFMetaData.Bin instead.
var VRFBin = VRFMetaData.Bin

// DeployVRF deploys a new Ethereum contract, binding an instance of VRF to it.
func DeployVRF(auth *bind.TransactOpts, backend bind.ContractBackend, _taskMailbox common.Address, _executorOperatorSet OperatorSet) (common.Address, *types.Transaction, *VRF, error) {
	parsed, err := VRFMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBin), backend, _taskMailbox, _executorOperatorSet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRF{VRFCaller: VRFCaller{contract: contract}, VRFTransactor: VRFTransactor{contract: contract}, VRFFilterer: VRFFilterer{contract: contract}}, nil
}

// VRF is an auto generated Go binding around an Ethereum contract.
type VRF struct {
	VRFCaller     // Read-only binding to the contract
	VRFTransactor // Write-only binding to the contract
	VRFFilterer   // Log filterer for contract events
}

// VRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFSession struct {
	Contract     *VRF              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFCallerSession struct {
	Contract *VRFCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFTransactorSession struct {
	Contract     *VRFTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFRaw struct {
	Contract *VRF // Generic contract binding to access the raw methods on
}

// VRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFCallerRaw struct {
	Contract *VRFCaller // Generic read-only contract binding to access the raw methods on
}

// VRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFTransactorRaw struct {
	Contract *VRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRF creates a new instance of VRF, bound to a specific deployed contract.
func NewVRF(address common.Address, backend bind.ContractBackend) (*VRF, error) {
	contract, err := bindVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRF{VRFCaller: VRFCaller{contract: contract}, VRFTransactor: VRFTransactor{contract: contract}, VRFFilterer: VRFFilterer{contract: contract}}, nil
}

// NewVRFCaller creates a new read-only instance of VRF, bound to a specific deployed contract.
func NewVRFCaller(address common.Address, caller bind.ContractCaller) (*VRFCaller, error) {
	contract, err := bindVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCaller{contract: contract}, nil
}

// NewVRFTransactor creates a new write-only instance of VRF, bound to a specific deployed contract.
func NewVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFTransactor, error) {
	contract, err := bindVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFTransactor{contract: contract}, nil
}

// NewVRFFilterer creates a new log filterer instance of VRF, bound to a specific deployed contract.
func NewVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFFilterer, error) {
	contract, err := bindVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFFilterer{contract: contract}, nil
}

// bindVRF binds a generic wrapper to an already deployed contract.
func bindVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRF *VRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.VRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRF *VRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRF *VRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRF *VRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRF *VRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRF *VRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transact(opts, method, params...)
}

// EncodeTaskPayload is a free data retrieval call binding the contract method 0xfdd154dc.
//
// Solidity: function encodeTaskPayload(uint8 randomnessType, bytes randomnessParams) pure returns(bytes)
func (_VRF *VRFCaller) EncodeTaskPayload(opts *bind.CallOpts, randomnessType uint8, randomnessParams []byte) ([]byte, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "encodeTaskPayload", randomnessType, randomnessParams)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTaskPayload is a free data retrieval call binding the contract method 0xfdd154dc.
//
// Solidity: function encodeTaskPayload(uint8 randomnessType, bytes randomnessParams) pure returns(bytes)
func (_VRF *VRFSession) EncodeTaskPayload(randomnessType uint8, randomnessParams []byte) ([]byte, error) {
	return _VRF.Contract.EncodeTaskPayload(&_VRF.CallOpts, randomnessType, randomnessParams)
}

// EncodeTaskPayload is a free data retrieval call binding the contract method 0xfdd154dc.
//
// Solidity: function encodeTaskPayload(uint8 randomnessType, bytes randomnessParams) pure returns(bytes)
func (_VRF *VRFCallerSession) EncodeTaskPayload(randomnessType uint8, randomnessParams []byte) ([]byte, error) {
	return _VRF.Contract.EncodeTaskPayload(&_VRF.CallOpts, randomnessType, randomnessParams)
}

// EncodeVDFParams is a free data retrieval call binding the contract method 0xcdae4940.
//
// Solidity: function encodeVDFParams(bytes seed) pure returns(bytes)
func (_VRF *VRFCaller) EncodeVDFParams(opts *bind.CallOpts, seed []byte) ([]byte, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "encodeVDFParams", seed)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeVDFParams is a free data retrieval call binding the contract method 0xcdae4940.
//
// Solidity: function encodeVDFParams(bytes seed) pure returns(bytes)
func (_VRF *VRFSession) EncodeVDFParams(seed []byte) ([]byte, error) {
	return _VRF.Contract.EncodeVDFParams(&_VRF.CallOpts, seed)
}

// EncodeVDFParams is a free data retrieval call binding the contract method 0xcdae4940.
//
// Solidity: function encodeVDFParams(bytes seed) pure returns(bytes)
func (_VRF *VRFCallerSession) EncodeVDFParams(seed []byte) ([]byte, error) {
	return _VRF.Contract.EncodeVDFParams(&_VRF.CallOpts, seed)
}

// EncodeVDFTaskPayload is a free data retrieval call binding the contract method 0xac8d0738.
//
// Solidity: function encodeVDFTaskPayload(bytes seed) pure returns(bytes)
func (_VRF *VRFCaller) EncodeVDFTaskPayload(opts *bind.CallOpts, seed []byte) ([]byte, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "encodeVDFTaskPayload", seed)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeVDFTaskPayload is a free data retrieval call binding the contract method 0xac8d0738.
//
// Solidity: function encodeVDFTaskPayload(bytes seed) pure returns(bytes)
func (_VRF *VRFSession) EncodeVDFTaskPayload(seed []byte) ([]byte, error) {
	return _VRF.Contract.EncodeVDFTaskPayload(&_VRF.CallOpts, seed)
}

// EncodeVDFTaskPayload is a free data retrieval call binding the contract method 0xac8d0738.
//
// Solidity: function encodeVDFTaskPayload(bytes seed) pure returns(bytes)
func (_VRF *VRFCallerSession) EncodeVDFTaskPayload(seed []byte) ([]byte, error) {
	return _VRF.Contract.EncodeVDFTaskPayload(&_VRF.CallOpts, seed)
}

// ExecutorOperatorSet is a free data retrieval call binding the contract method 0x2026a4f3.
//
// Solidity: function executorOperatorSet() view returns(address avs, uint32 id)
func (_VRF *VRFCaller) ExecutorOperatorSet(opts *bind.CallOpts) (struct {
	Avs common.Address
	Id  uint32
}, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "executorOperatorSet")

	outstruct := new(struct {
		Avs common.Address
		Id  uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Avs = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// ExecutorOperatorSet is a free data retrieval call binding the contract method 0x2026a4f3.
//
// Solidity: function executorOperatorSet() view returns(address avs, uint32 id)
func (_VRF *VRFSession) ExecutorOperatorSet() (struct {
	Avs common.Address
	Id  uint32
}, error) {
	return _VRF.Contract.ExecutorOperatorSet(&_VRF.CallOpts)
}

// ExecutorOperatorSet is a free data retrieval call binding the contract method 0x2026a4f3.
//
// Solidity: function executorOperatorSet() view returns(address avs, uint32 id)
func (_VRF *VRFCallerSession) ExecutorOperatorSet() (struct {
	Avs common.Address
	Id  uint32
}, error) {
	return _VRF.Contract.ExecutorOperatorSet(&_VRF.CallOpts)
}

// GetRandomnessResult is a free data retrieval call binding the contract method 0x6a36dab9.
//
// Solidity: function getRandomnessResult(uint256 requestId) view returns(bool fulfilled, uint256 result)
func (_VRF *VRFCaller) GetRandomnessResult(opts *bind.CallOpts, requestId *big.Int) (struct {
	Fulfilled bool
	Result    *big.Int
}, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "getRandomnessResult", requestId)

	outstruct := new(struct {
		Fulfilled bool
		Result    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Result = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRandomnessResult is a free data retrieval call binding the contract method 0x6a36dab9.
//
// Solidity: function getRandomnessResult(uint256 requestId) view returns(bool fulfilled, uint256 result)
func (_VRF *VRFSession) GetRandomnessResult(requestId *big.Int) (struct {
	Fulfilled bool
	Result    *big.Int
}, error) {
	return _VRF.Contract.GetRandomnessResult(&_VRF.CallOpts, requestId)
}

// GetRandomnessResult is a free data retrieval call binding the contract method 0x6a36dab9.
//
// Solidity: function getRandomnessResult(uint256 requestId) view returns(bool fulfilled, uint256 result)
func (_VRF *VRFCallerSession) GetRandomnessResult(requestId *big.Int) (struct {
	Fulfilled bool
	Result    *big.Int
}, error) {
	return _VRF.Contract.GetRandomnessResult(&_VRF.CallOpts, requestId)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns((address,bytes32,uint256,bool,uint256) request)
func (_VRF *VRFCaller) GetRequest(opts *bind.CallOpts, requestId *big.Int) (VRFRandomnessRequest, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "getRequest", requestId)

	if err != nil {
		return *new(VRFRandomnessRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(VRFRandomnessRequest)).(*VRFRandomnessRequest)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns((address,bytes32,uint256,bool,uint256) request)
func (_VRF *VRFSession) GetRequest(requestId *big.Int) (VRFRandomnessRequest, error) {
	return _VRF.Contract.GetRequest(&_VRF.CallOpts, requestId)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns((address,bytes32,uint256,bool,uint256) request)
func (_VRF *VRFCallerSession) GetRequest(requestId *big.Int) (VRFRandomnessRequest, error) {
	return _VRF.Contract.GetRequest(&_VRF.CallOpts, requestId)
}

// GetRequestCounter is a free data retrieval call binding the contract method 0xa2785f0b.
//
// Solidity: function getRequestCounter() view returns(uint256)
func (_VRF *VRFCaller) GetRequestCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "getRequestCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestCounter is a free data retrieval call binding the contract method 0xa2785f0b.
//
// Solidity: function getRequestCounter() view returns(uint256)
func (_VRF *VRFSession) GetRequestCounter() (*big.Int, error) {
	return _VRF.Contract.GetRequestCounter(&_VRF.CallOpts)
}

// GetRequestCounter is a free data retrieval call binding the contract method 0xa2785f0b.
//
// Solidity: function getRequestCounter() view returns(uint256)
func (_VRF *VRFCallerSession) GetRequestCounter() (*big.Int, error) {
	return _VRF.Contract.GetRequestCounter(&_VRF.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(address requester, bytes32 taskHash, uint256 blockNumber, bool fulfilled, uint256 result)
func (_VRF *VRFCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requester   common.Address
	TaskHash    [32]byte
	BlockNumber *big.Int
	Fulfilled   bool
	Result      *big.Int
}, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Requester   common.Address
		TaskHash    [32]byte
		BlockNumber *big.Int
		Fulfilled   bool
		Result      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TaskHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.BlockNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Result = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(address requester, bytes32 taskHash, uint256 blockNumber, bool fulfilled, uint256 result)
func (_VRF *VRFSession) Requests(arg0 *big.Int) (struct {
	Requester   common.Address
	TaskHash    [32]byte
	BlockNumber *big.Int
	Fulfilled   bool
	Result      *big.Int
}, error) {
	return _VRF.Contract.Requests(&_VRF.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(address requester, bytes32 taskHash, uint256 blockNumber, bool fulfilled, uint256 result)
func (_VRF *VRFCallerSession) Requests(arg0 *big.Int) (struct {
	Requester   common.Address
	TaskHash    [32]byte
	BlockNumber *big.Int
	Fulfilled   bool
	Result      *big.Int
}, error) {
	return _VRF.Contract.Requests(&_VRF.CallOpts, arg0)
}

// TaskHashToRequestId is a free data retrieval call binding the contract method 0x15505f8c.
//
// Solidity: function taskHashToRequestId(bytes32 ) view returns(uint256)
func (_VRF *VRFCaller) TaskHashToRequestId(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "taskHashToRequestId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskHashToRequestId is a free data retrieval call binding the contract method 0x15505f8c.
//
// Solidity: function taskHashToRequestId(bytes32 ) view returns(uint256)
func (_VRF *VRFSession) TaskHashToRequestId(arg0 [32]byte) (*big.Int, error) {
	return _VRF.Contract.TaskHashToRequestId(&_VRF.CallOpts, arg0)
}

// TaskHashToRequestId is a free data retrieval call binding the contract method 0x15505f8c.
//
// Solidity: function taskHashToRequestId(bytes32 ) view returns(uint256)
func (_VRF *VRFCallerSession) TaskHashToRequestId(arg0 [32]byte) (*big.Int, error) {
	return _VRF.Contract.TaskHashToRequestId(&_VRF.CallOpts, arg0)
}

// TaskMailbox is a free data retrieval call binding the contract method 0xf42a9e13.
//
// Solidity: function taskMailbox() view returns(address)
func (_VRF *VRFCaller) TaskMailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "taskMailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskMailbox is a free data retrieval call binding the contract method 0xf42a9e13.
//
// Solidity: function taskMailbox() view returns(address)
func (_VRF *VRFSession) TaskMailbox() (common.Address, error) {
	return _VRF.Contract.TaskMailbox(&_VRF.CallOpts)
}

// TaskMailbox is a free data retrieval call binding the contract method 0xf42a9e13.
//
// Solidity: function taskMailbox() view returns(address)
func (_VRF *VRFCallerSession) TaskMailbox() (common.Address, error) {
	return _VRF.Contract.TaskMailbox(&_VRF.CallOpts)
}

// OnTaskCompleted is a paid mutator transaction binding the contract method 0x0944a904.
//
// Solidity: function onTaskCompleted(bytes32 taskHash, bytes result) returns()
func (_VRF *VRFTransactor) OnTaskCompleted(opts *bind.TransactOpts, taskHash [32]byte, result []byte) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "onTaskCompleted", taskHash, result)
}

// OnTaskCompleted is a paid mutator transaction binding the contract method 0x0944a904.
//
// Solidity: function onTaskCompleted(bytes32 taskHash, bytes result) returns()
func (_VRF *VRFSession) OnTaskCompleted(taskHash [32]byte, result []byte) (*types.Transaction, error) {
	return _VRF.Contract.OnTaskCompleted(&_VRF.TransactOpts, taskHash, result)
}

// OnTaskCompleted is a paid mutator transaction binding the contract method 0x0944a904.
//
// Solidity: function onTaskCompleted(bytes32 taskHash, bytes result) returns()
func (_VRF *VRFTransactorSession) OnTaskCompleted(taskHash [32]byte, result []byte) (*types.Transaction, error) {
	return _VRF.Contract.OnTaskCompleted(&_VRF.TransactOpts, taskHash, result)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xe6dee7ed.
//
// Solidity: function requestRandomness(bytes seed) returns(uint256 requestId)
func (_VRF *VRFTransactor) RequestRandomness(opts *bind.TransactOpts, seed []byte) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "requestRandomness", seed)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xe6dee7ed.
//
// Solidity: function requestRandomness(bytes seed) returns(uint256 requestId)
func (_VRF *VRFSession) RequestRandomness(seed []byte) (*types.Transaction, error) {
	return _VRF.Contract.RequestRandomness(&_VRF.TransactOpts, seed)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xe6dee7ed.
//
// Solidity: function requestRandomness(bytes seed) returns(uint256 requestId)
func (_VRF *VRFTransactorSession) RequestRandomness(seed []byte) (*types.Transaction, error) {
	return _VRF.Contract.RequestRandomness(&_VRF.TransactOpts, seed)
}

// UpdateOperatorSet is a paid mutator transaction binding the contract method 0xe7a923bb.
//
// Solidity: function updateOperatorSet((address,uint32) _executorOperatorSet) returns()
func (_VRF *VRFTransactor) UpdateOperatorSet(opts *bind.TransactOpts, _executorOperatorSet OperatorSet) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "updateOperatorSet", _executorOperatorSet)
}

// UpdateOperatorSet is a paid mutator transaction binding the contract method 0xe7a923bb.
//
// Solidity: function updateOperatorSet((address,uint32) _executorOperatorSet) returns()
func (_VRF *VRFSession) UpdateOperatorSet(_executorOperatorSet OperatorSet) (*types.Transaction, error) {
	return _VRF.Contract.UpdateOperatorSet(&_VRF.TransactOpts, _executorOperatorSet)
}

// UpdateOperatorSet is a paid mutator transaction binding the contract method 0xe7a923bb.
//
// Solidity: function updateOperatorSet((address,uint32) _executorOperatorSet) returns()
func (_VRF *VRFTransactorSession) UpdateOperatorSet(_executorOperatorSet OperatorSet) (*types.Transaction, error) {
	return _VRF.Contract.UpdateOperatorSet(&_VRF.TransactOpts, _executorOperatorSet)
}

// VRFRandomnessFulfilledIterator is returned from FilterRandomnessFulfilled and is used to iterate over the raw logs and unpacked data for RandomnessFulfilled events raised by the VRF contract.
type VRFRandomnessFulfilledIterator struct {
	Event *VRFRandomnessFulfilled // Event containing the contract specifics and raw log

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
func (it *VRFRandomnessFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRandomnessFulfilled)
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
		it.Event = new(VRFRandomnessFulfilled)
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
func (it *VRFRandomnessFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFRandomnessFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFRandomnessFulfilled represents a RandomnessFulfilled event raised by the VRF contract.
type VRFRandomnessFulfilled struct {
	RequestId *big.Int
	TaskHash  [32]byte
	Result    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRandomnessFulfilled is a free log retrieval operation binding the contract event 0xbe3f52bb4df8f041a3e9118a0acf71ddeb18adc6fdb619f10c637273eab05e42.
//
// Solidity: event RandomnessFulfilled(uint256 indexed requestId, bytes32 indexed taskHash, uint256 result)
func (_VRF *VRFFilterer) FilterRandomnessFulfilled(opts *bind.FilterOpts, requestId []*big.Int, taskHash [][32]byte) (*VRFRandomnessFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "RandomnessFulfilled", requestIdRule, taskHashRule)
	if err != nil {
		return nil, err
	}
	return &VRFRandomnessFulfilledIterator{contract: _VRF.contract, event: "RandomnessFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomnessFulfilled is a free log subscription operation binding the contract event 0xbe3f52bb4df8f041a3e9118a0acf71ddeb18adc6fdb619f10c637273eab05e42.
//
// Solidity: event RandomnessFulfilled(uint256 indexed requestId, bytes32 indexed taskHash, uint256 result)
func (_VRF *VRFFilterer) WatchRandomnessFulfilled(opts *bind.WatchOpts, sink chan<- *VRFRandomnessFulfilled, requestId []*big.Int, taskHash [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "RandomnessFulfilled", requestIdRule, taskHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFRandomnessFulfilled)
				if err := _VRF.contract.UnpackLog(event, "RandomnessFulfilled", log); err != nil {
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

// ParseRandomnessFulfilled is a log parse operation binding the contract event 0xbe3f52bb4df8f041a3e9118a0acf71ddeb18adc6fdb619f10c637273eab05e42.
//
// Solidity: event RandomnessFulfilled(uint256 indexed requestId, bytes32 indexed taskHash, uint256 result)
func (_VRF *VRFFilterer) ParseRandomnessFulfilled(log types.Log) (*VRFRandomnessFulfilled, error) {
	event := new(VRFRandomnessFulfilled)
	if err := _VRF.contract.UnpackLog(event, "RandomnessFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFRandomnessRequestedIterator is returned from FilterRandomnessRequested and is used to iterate over the raw logs and unpacked data for RandomnessRequested events raised by the VRF contract.
type VRFRandomnessRequestedIterator struct {
	Event *VRFRandomnessRequested // Event containing the contract specifics and raw log

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
func (it *VRFRandomnessRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRandomnessRequested)
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
		it.Event = new(VRFRandomnessRequested)
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
func (it *VRFRandomnessRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFRandomnessRequested represents a RandomnessRequested event raised by the VRF contract.
type VRFRandomnessRequested struct {
	RequestId *big.Int
	Requester common.Address
	TaskHash  [32]byte
	Seed      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRandomnessRequested is a free log retrieval operation binding the contract event 0x152b260fdc6b51380aea47b7a68bd79722f0459b2603c789e9e9f476007e07be.
//
// Solidity: event RandomnessRequested(uint256 indexed requestId, address indexed requester, bytes32 indexed taskHash, bytes seed)
func (_VRF *VRFFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestId []*big.Int, requester []common.Address, taskHash [][32]byte) (*VRFRandomnessRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "RandomnessRequested", requestIdRule, requesterRule, taskHashRule)
	if err != nil {
		return nil, err
	}
	return &VRFRandomnessRequestedIterator{contract: _VRF.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

// WatchRandomnessRequested is a free log subscription operation binding the contract event 0x152b260fdc6b51380aea47b7a68bd79722f0459b2603c789e9e9f476007e07be.
//
// Solidity: event RandomnessRequested(uint256 indexed requestId, address indexed requester, bytes32 indexed taskHash, bytes seed)
func (_VRF *VRFFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFRandomnessRequested, requestId []*big.Int, requester []common.Address, taskHash [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "RandomnessRequested", requestIdRule, requesterRule, taskHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFRandomnessRequested)
				if err := _VRF.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

// ParseRandomnessRequested is a log parse operation binding the contract event 0x152b260fdc6b51380aea47b7a68bd79722f0459b2603c789e9e9f476007e07be.
//
// Solidity: event RandomnessRequested(uint256 indexed requestId, address indexed requester, bytes32 indexed taskHash, bytes seed)
func (_VRF *VRFFilterer) ParseRandomnessRequested(log types.Log) (*VRFRandomnessRequested, error) {
	event := new(VRFRandomnessRequested)
	if err := _VRF.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
