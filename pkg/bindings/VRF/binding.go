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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBN254CertificateVerifierTypesBN254Certificate is an auto generated low-level Go binding around an user-defined struct.
type IBN254CertificateVerifierTypesBN254Certificate struct {
	ReferenceTimestamp uint32
	MessageHash        [32]byte
	Signature          BN254G1Point
	Apk                BN254G2Point
	NonSignerWitnesses []IBN254CertificateVerifierTypesBN254OperatorInfoWitness
}

// IBN254CertificateVerifierTypesBN254OperatorInfoWitness is an auto generated low-level Go binding around an user-defined struct.
type IBN254CertificateVerifierTypesBN254OperatorInfoWitness struct {
	OperatorIndex     uint32
	OperatorInfoProof []byte
	OperatorInfo      IBN254TableCalculatorTypesBN254OperatorInfo
}

// IBN254TableCalculatorTypesBN254OperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IBN254TableCalculatorTypesBN254OperatorInfo struct {
	Pubkey  BN254G1Point
	Weights []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// VRFRandomnessRequest is an auto generated low-level Go binding around an user-defined struct.
type VRFRandomnessRequest struct {
	Requester   common.Address
	BlockNumber *big.Int
	Fulfilled   bool
	Result      *big.Int
}

// VRFMetaData contains all meta data concerning the VRF contract.
var VRFMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskMailbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decodeAndValidateTaskPayload\",\"inputs\":[{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeTaskPayload\",\"inputs\":[{\"name\":\"randomnessType\",\"type\":\"uint8\",\"internalType\":\"enumVRF.RandomnessType\"},{\"name\":\"randomnessParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeVDFParams\",\"inputs\":[{\"name\":\"seed\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeVDFTaskPayload\",\"inputs\":[{\"name\":\"seed\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"executorOperatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRandomnessResult\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structVRF.RandomnessRequest\",\"components\":[{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onTaskCompleted\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestRandomness\",\"inputs\":[{\"name\":\"seed\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requests\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskMailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskMailbox\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorSet\",\"inputs\":[{\"name\":\"_executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatePostTaskCreation\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatePreTaskCreation\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateTaskResultSubmission\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"cert\",\"type\":\"tuple\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBN254CertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBN254TableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"RandomnessFulfilled\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomnessRequested\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seed\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResultSubmissionValidated\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidTaskResponse\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequestAlreadyFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequestNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516116ec3803806116ec83398101604081905261002e91610087565b6001600160a01b0391821660805280515f805460209093015163ffffffff16600160a01b026001600160c01b0319909316919093161717905561011b565b80516001600160a01b0381168114610082575f5ffd5b919050565b5f5f8284036060811215610099575f5ffd5b6100a28461006c565b92506040601f19820112156100b5575f5ffd5b50604080519081016001600160401b03811182821017156100e457634e487b7160e01b5f52604160045260245ffd5b6040526100f36020850161006c565b8152604084015163ffffffff8116811461010b575f5ffd5b6020820152919491935090915050565b6080516115ab6101415f395f81816103280152818161044001526108e901526115ab5ff3fe608060405234801561000f575f5ffd5b50600436106100f0575f3560e01c8063e45c4a0d11610093578063ed0c943f11610063578063ed0c943f14610310578063f42a9e1314610323578063fb1e61ca14610362578063fdd154dc14610422575f5ffd5b8063e45c4a0d1461021e578063e507027a14610294578063e6dee7ed146102a7578063e7a923bb146102c8575f5ffd5b80638679c781116100ce5780638679c781146101665780639d86698514610177578063ac8d0738146101eb578063cdae49401461020b575f5ffd5b80630944a904146100f45780632026a4f314610109578063485e73b814610153575b5f5ffd5b610107610102366004610bd8565b610435565b005b5f5461012a906001600160a01b03811690600160a01b900463ffffffff1682565b604080516001600160a01b03909316835263ffffffff9091166020830152015b60405180910390f35b610107610161366004610fc8565b610597565b6101076101743660046110b2565b50565b6101bb6101853660046110b2565b600160208190525f918252604090912080549181015460028201546003909201546001600160a01b0390931692909160ff169084565b604080516001600160a01b039095168552602085019390935290151591830191909152606082015260800161014a565b6101fe6101f93660046110c9565b610625565b60405161014a9190611135565b6101fe6102193660046110c9565b6106c9565b61027d61022c3660046110b2565b5f90815260016020818152604092839020835160808101855281546001600160a01b031681529281015491830191909152600281015460ff1615159282018390526003015460609091018190529091565b60408051921515835260208301919091520161014a565b6101076102a23660046111a0565b610732565b6102ba6102b53660046110c9565b6107f0565b60405190815260200161014a565b6101076102d63660046111fa565b80515f805460209093015163ffffffff16600160a01b026001600160c01b03199093166001600160a01b0390921691909117919091179055565b61010761031e366004611214565b610a20565b61034a7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161014a565b6103e46103703660046110b2565b604080516080810182525f808252602082018190529181018290526060810191909152505f90815260016020818152604092839020835160808101855281546001600160a01b031681529281015491830191909152600281015460ff16151592820192909252600390910154606082015290565b60405161014a919081516001600160a01b03168152602080830151908201526040808301511515908201526060918201519181019190915260800190565b6101fe610430366004611259565b610b0b565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461047e57604051635c427cd960e01b815260040160405180910390fd5b5f83815260016020526040902080548491906001600160a01b03166104b657604051632589d98f60e11b815260040160405180910390fd5b600281015460ff16156104dc5760405163533d99dd60e01b815260040160405180910390fd5b5f6104e984860186611290565b90506001815160018111156105005761050061130d565b1461051e5760405163413041d160e01b815260040160405180910390fd5b5f81602001518060200190518101906105379190611321565b60028401805460ff1916600117905580516003850181905560405191925085917f9b0aa3f92f46e24caa76b000bdf0dd495b9b390c320cf6585ae10a12b7d09edb916105869190815260200190565b60405180910390a250505050505050565b5f82815260016020526040902080548391906001600160a01b03166105cf57604051632589d98f60e11b815260040160405180910390fd5b600281015460ff16156105f55760405163533d99dd60e01b815260040160405180910390fd5b60405182907f0d5aeffd61ba930c83f1e88ec8bc110c50d1a38cf03190b0272c312d16ca7140905f90a250505050565b604080516020601f8401819004810282018301835281018381526060925f9291829187908790819085018382808284375f9201829052509390945250506040805180820190915292935091905080600181526020018360405160200161068b9190611345565b6040516020818303038152906040528152509050806040516020016106b0919061135f565b6040516020818303038152906040529250505092915050565b604080516020601f8401819004810282018301835281018381526060925f9291829187908790819085018382808284375f92019190915250505091525060405190915061071a908290602001611345565b60405160208183030381529060405291505092915050565b60405163ed0c943f60e01b8152309063ed0c943f90610755908490600401611135565b5f6040518083038186803b15801561076b575f5ffd5b505afa92505050801561077c575060015b6107995760405163413041d160e01b815260040160405180910390fd5b5f5482516001600160a01b0390811691161415806107cd57505f54602083015163ffffffff908116600160a01b9092041614155b156107eb5760405163413041d160e01b815260040160405180910390fd5b505050565b604080516020601f8401819004810282018301835281018381525f92839291829187908790819085018382808284375f920182905250939094525050604080518082019091529293509190508060018152602001836040516020016108559190611345565b60408051808303601f19018152918152915280516080810182523381525f60208083018290528351808501855282546001600160a01b0381168252600160a01b900463ffffffff168183015283850152925193945092909160608301916108be9186910161135f565b60408051601f1981840301815291815291525162221dbd60e51b81529091505f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690630443b7a09061091e9085906004016113a1565b6020604051808303815f875af115801561093a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095e919061140a565b60408051608081018252338082524360208084019182525f8486018181526060860182815288835260019384905291879020955186546001600160a01b0319166001600160a01b03909116178655925191850191909155905160028401805460ff19169115159190911790555160039092019190915590519196508692509082907fadbe210fbb3f1373d938d421ee71b9db071c3d86ccaf8851d6bda23e70c6b53290610a0e908b908b90611421565b60405180910390a35050505092915050565b5f81806020019051810190610a35919061149c565b9050600181516001811115610a4c57610a4c61130d565b14158015610a6c57505f81516001811115610a6957610a6961130d565b14155b15610a8a5760405163413041d160e01b815260040160405180910390fd5b600181516001811115610a9f57610a9f61130d565b03610b07575f8160200151806020019051810190610abd919061150b565b8051519091505f03610ae25760405163413041d160e01b815260040160405180910390fd5b80515161040010156107eb5760405163413041d160e01b815260040160405180910390fd5b5050565b60605f6040518060400160405280866001811115610b2b57610b2b61130d565b815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050915250604051909150610b7b90829060200161135f565b6040516020818303038152906040529150509392505050565b5f5f83601f840112610ba4575f5ffd5b5081356001600160401b03811115610bba575f5ffd5b602083019150836020828501011115610bd1575f5ffd5b9250929050565b5f5f5f60408486031215610bea575f5ffd5b8335925060208401356001600160401b03811115610c06575f5ffd5b610c1286828701610b94565b9497909650939450505050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715610c5557610c55610c1f565b60405290565b604051606081016001600160401b0381118282101715610c5557610c55610c1f565b60405160a081016001600160401b0381118282101715610c5557610c55610c1f565b604051602081016001600160401b0381118282101715610c5557610c55610c1f565b604051601f8201601f191681016001600160401b0381118282101715610ce957610ce9610c1f565b604052919050565b803563ffffffff81168114610d04575f5ffd5b919050565b5f60408284031215610d19575f5ffd5b610d21610c33565b823581526020928301359281019290925250919050565b5f82601f830112610d47575f5ffd5b610d4f610c33565b806040840185811115610d60575f5ffd5b845b81811015610d7a578035845260209384019301610d62565b509095945050505050565b5f6001600160401b03821115610d9d57610d9d610c1f565b5060051b60200190565b5f6001600160401b03821115610dbf57610dbf610c1f565b50601f01601f191660200190565b5f82601f830112610ddc575f5ffd5b8135610def610dea82610da7565b610cc1565b818152846020838601011115610e03575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f82601f830112610e2e575f5ffd5b8135610e3c610dea82610d85565b8082825260208201915060208360051b860101925085831115610e5d575f5ffd5b602085015b83811015610fbe5780356001600160401b03811115610e7f575f5ffd5b86016060818903601f19011215610e94575f5ffd5b610e9c610c5b565b610ea860208301610cf1565b815260408201356001600160401b03811115610ec2575f5ffd5b610ed18a602083860101610dcd565b60208301525060608201356001600160401b03811115610eef575f5ffd5b6020818401019250506060828a031215610f07575f5ffd5b610f0f610c33565b610f198a84610d09565b815260408301356001600160401b03811115610f33575f5ffd5b80840193505089601f840112610f47575f5ffd5b8235610f55610dea82610d85565b8082825260208201915060208360051b87010192508c831115610f76575f5ffd5b6020860195505b82861015610f98578535825260209586019590910190610f7d565b806020850152505050806040830152508085525050602083019250602081019050610e62565b5095945050505050565b5f5f60408385031215610fd9575f5ffd5b8235915060208301356001600160401b03811115610ff5575f5ffd5b8301808503610120811215611008575f5ffd5b611010610c7d565b61101983610cf1565b8152602083810135908201526110328760408501610d09565b60408201526080607f1983011215611048575f5ffd5b611050610c33565b915061105f8760808501610d38565b825261106e8760c08501610d38565b602083015281606082015261010083013591506001600160401b03821115611094575f5ffd5b6110a087838501610e1f565b60808201528093505050509250929050565b5f602082840312156110c2575f5ffd5b5035919050565b5f5f602083850312156110da575f5ffd5b82356001600160401b038111156110ef575f5ffd5b6110fb85828601610b94565b90969095509350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6111476020830184611107565b9392505050565b80356001600160a01b0381168114610d04575f5ffd5b5f60408284031215611174575f5ffd5b61117c610c33565b90506111878261114e565b815261119560208301610cf1565b602082015292915050565b5f5f5f608084860312156111b2575f5ffd5b6111bb8461114e565b92506111ca8560208601611164565b915060608401356001600160401b038111156111e4575f5ffd5b6111f086828701610dcd565b9150509250925092565b5f6040828403121561120a575f5ffd5b6111478383611164565b5f60208284031215611224575f5ffd5b81356001600160401b03811115611239575f5ffd5b61124584828501610dcd565b949350505050565b60028110610174575f5ffd5b5f5f5f6040848603121561126b575f5ffd5b83356112768161124d565b925060208401356001600160401b03811115610c06575f5ffd5b5f602082840312156112a0575f5ffd5b81356001600160401b038111156112b5575f5ffd5b8201604081850312156112c6575f5ffd5b6112ce610c33565b81356112d98161124d565b815260208201356001600160401b038111156112f3575f5ffd5b6112ff86828501610dcd565b602083015250949350505050565b634e487b7160e01b5f52602160045260245ffd5b5f6020828403128015611332575f5ffd5b5061133b610c9f565b9151825250919050565b602081525f82516020808401526112456040840182611107565b602081525f82516002811061138257634e487b7160e01b5f52602160045260245ffd5b8060208401525060208301516040808401526112456060840182611107565b6020815260018060a01b0382511660208201526bffffffffffffffffffffffff60208301511660408201525f604083015160018060a01b03815116606084015263ffffffff602082015116608084015250606083015160a08084015261124560c0840182611107565b5f6020828403121561141a575f5ffd5b5051919050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f82601f83011261145e575f5ffd5b815161146c610dea82610da7565b818152846020838601011115611480575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f602082840312156114ac575f5ffd5b81516001600160401b038111156114c1575f5ffd5b8201604081850312156114d2575f5ffd5b6114da610c33565b81516114e58161124d565b815260208201516001600160401b038111156114ff575f5ffd5b6112ff8682850161144f565b5f6020828403121561151b575f5ffd5b81516001600160401b03811115611530575f5ffd5b820160208185031215611541575f5ffd5b611549610c9f565b81516001600160401b0381111561155e575f5ffd5b61156a8682850161144f565b82525094935050505056fea26469706673582212209a981b3d7bc0eb41cd2f4930a6f23571f640733fd3903b1fcc70f773c3f7457964736f6c634300081b0033",
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

// DecodeAndValidateTaskPayload is a free data retrieval call binding the contract method 0xed0c943f.
//
// Solidity: function decodeAndValidateTaskPayload(bytes payload) pure returns()
func (_VRF *VRFCaller) DecodeAndValidateTaskPayload(opts *bind.CallOpts, payload []byte) error {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "decodeAndValidateTaskPayload", payload)

	if err != nil {
		return err
	}

	return err

}

// DecodeAndValidateTaskPayload is a free data retrieval call binding the contract method 0xed0c943f.
//
// Solidity: function decodeAndValidateTaskPayload(bytes payload) pure returns()
func (_VRF *VRFSession) DecodeAndValidateTaskPayload(payload []byte) error {
	return _VRF.Contract.DecodeAndValidateTaskPayload(&_VRF.CallOpts, payload)
}

// DecodeAndValidateTaskPayload is a free data retrieval call binding the contract method 0xed0c943f.
//
// Solidity: function decodeAndValidateTaskPayload(bytes payload) pure returns()
func (_VRF *VRFCallerSession) DecodeAndValidateTaskPayload(payload []byte) error {
	return _VRF.Contract.DecodeAndValidateTaskPayload(&_VRF.CallOpts, payload)
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

// GetRandomnessResult is a free data retrieval call binding the contract method 0xe45c4a0d.
//
// Solidity: function getRandomnessResult(bytes32 requestId) view returns(bool fulfilled, uint256 result)
func (_VRF *VRFCaller) GetRandomnessResult(opts *bind.CallOpts, requestId [32]byte) (struct {
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

// GetRandomnessResult is a free data retrieval call binding the contract method 0xe45c4a0d.
//
// Solidity: function getRandomnessResult(bytes32 requestId) view returns(bool fulfilled, uint256 result)
func (_VRF *VRFSession) GetRandomnessResult(requestId [32]byte) (struct {
	Fulfilled bool
	Result    *big.Int
}, error) {
	return _VRF.Contract.GetRandomnessResult(&_VRF.CallOpts, requestId)
}

// GetRandomnessResult is a free data retrieval call binding the contract method 0xe45c4a0d.
//
// Solidity: function getRandomnessResult(bytes32 requestId) view returns(bool fulfilled, uint256 result)
func (_VRF *VRFCallerSession) GetRandomnessResult(requestId [32]byte) (struct {
	Fulfilled bool
	Result    *big.Int
}, error) {
	return _VRF.Contract.GetRandomnessResult(&_VRF.CallOpts, requestId)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 requestId) view returns((address,uint256,bool,uint256) request)
func (_VRF *VRFCaller) GetRequest(opts *bind.CallOpts, requestId [32]byte) (VRFRandomnessRequest, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "getRequest", requestId)

	if err != nil {
		return *new(VRFRandomnessRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(VRFRandomnessRequest)).(*VRFRandomnessRequest)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 requestId) view returns((address,uint256,bool,uint256) request)
func (_VRF *VRFSession) GetRequest(requestId [32]byte) (VRFRandomnessRequest, error) {
	return _VRF.Contract.GetRequest(&_VRF.CallOpts, requestId)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 requestId) view returns((address,uint256,bool,uint256) request)
func (_VRF *VRFCallerSession) GetRequest(requestId [32]byte) (VRFRandomnessRequest, error) {
	return _VRF.Contract.GetRequest(&_VRF.CallOpts, requestId)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(address requester, uint256 blockNumber, bool fulfilled, uint256 result)
func (_VRF *VRFCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Requester   common.Address
	BlockNumber *big.Int
	Fulfilled   bool
	Result      *big.Int
}, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Requester   common.Address
		BlockNumber *big.Int
		Fulfilled   bool
		Result      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Result = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(address requester, uint256 blockNumber, bool fulfilled, uint256 result)
func (_VRF *VRFSession) Requests(arg0 [32]byte) (struct {
	Requester   common.Address
	BlockNumber *big.Int
	Fulfilled   bool
	Result      *big.Int
}, error) {
	return _VRF.Contract.Requests(&_VRF.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(address requester, uint256 blockNumber, bool fulfilled, uint256 result)
func (_VRF *VRFCallerSession) Requests(arg0 [32]byte) (struct {
	Requester   common.Address
	BlockNumber *big.Int
	Fulfilled   bool
	Result      *big.Int
}, error) {
	return _VRF.Contract.Requests(&_VRF.CallOpts, arg0)
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

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0xe507027a.
//
// Solidity: function validatePreTaskCreation(address caller, (address,uint32) operatorSet, bytes payload) view returns()
func (_VRF *VRFCaller) ValidatePreTaskCreation(opts *bind.CallOpts, caller common.Address, operatorSet OperatorSet, payload []byte) error {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "validatePreTaskCreation", caller, operatorSet, payload)

	if err != nil {
		return err
	}

	return err

}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0xe507027a.
//
// Solidity: function validatePreTaskCreation(address caller, (address,uint32) operatorSet, bytes payload) view returns()
func (_VRF *VRFSession) ValidatePreTaskCreation(caller common.Address, operatorSet OperatorSet, payload []byte) error {
	return _VRF.Contract.ValidatePreTaskCreation(&_VRF.CallOpts, caller, operatorSet, payload)
}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0xe507027a.
//
// Solidity: function validatePreTaskCreation(address caller, (address,uint32) operatorSet, bytes payload) view returns()
func (_VRF *VRFCallerSession) ValidatePreTaskCreation(caller common.Address, operatorSet OperatorSet, payload []byte) error {
	return _VRF.Contract.ValidatePreTaskCreation(&_VRF.CallOpts, caller, operatorSet, payload)
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
// Solidity: function requestRandomness(bytes seed) returns(bytes32 requestId)
func (_VRF *VRFTransactor) RequestRandomness(opts *bind.TransactOpts, seed []byte) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "requestRandomness", seed)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xe6dee7ed.
//
// Solidity: function requestRandomness(bytes seed) returns(bytes32 requestId)
func (_VRF *VRFSession) RequestRandomness(seed []byte) (*types.Transaction, error) {
	return _VRF.Contract.RequestRandomness(&_VRF.TransactOpts, seed)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xe6dee7ed.
//
// Solidity: function requestRandomness(bytes seed) returns(bytes32 requestId)
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

// ValidatePostTaskCreation is a paid mutator transaction binding the contract method 0x8679c781.
//
// Solidity: function validatePostTaskCreation(bytes32 taskHash) returns()
func (_VRF *VRFTransactor) ValidatePostTaskCreation(opts *bind.TransactOpts, taskHash [32]byte) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "validatePostTaskCreation", taskHash)
}

// ValidatePostTaskCreation is a paid mutator transaction binding the contract method 0x8679c781.
//
// Solidity: function validatePostTaskCreation(bytes32 taskHash) returns()
func (_VRF *VRFSession) ValidatePostTaskCreation(taskHash [32]byte) (*types.Transaction, error) {
	return _VRF.Contract.ValidatePostTaskCreation(&_VRF.TransactOpts, taskHash)
}

// ValidatePostTaskCreation is a paid mutator transaction binding the contract method 0x8679c781.
//
// Solidity: function validatePostTaskCreation(bytes32 taskHash) returns()
func (_VRF *VRFTransactorSession) ValidatePostTaskCreation(taskHash [32]byte) (*types.Transaction, error) {
	return _VRF.Contract.ValidatePostTaskCreation(&_VRF.TransactOpts, taskHash)
}

// ValidateTaskResultSubmission is a paid mutator transaction binding the contract method 0x485e73b8.
//
// Solidity: function validateTaskResultSubmission(bytes32 taskHash, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns()
func (_VRF *VRFTransactor) ValidateTaskResultSubmission(opts *bind.TransactOpts, taskHash [32]byte, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "validateTaskResultSubmission", taskHash, cert)
}

// ValidateTaskResultSubmission is a paid mutator transaction binding the contract method 0x485e73b8.
//
// Solidity: function validateTaskResultSubmission(bytes32 taskHash, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns()
func (_VRF *VRFSession) ValidateTaskResultSubmission(taskHash [32]byte, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _VRF.Contract.ValidateTaskResultSubmission(&_VRF.TransactOpts, taskHash, cert)
}

// ValidateTaskResultSubmission is a paid mutator transaction binding the contract method 0x485e73b8.
//
// Solidity: function validateTaskResultSubmission(bytes32 taskHash, (uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint32,bytes,((uint256,uint256),uint256[]))[]) cert) returns()
func (_VRF *VRFTransactorSession) ValidateTaskResultSubmission(taskHash [32]byte, cert IBN254CertificateVerifierTypesBN254Certificate) (*types.Transaction, error) {
	return _VRF.Contract.ValidateTaskResultSubmission(&_VRF.TransactOpts, taskHash, cert)
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
	RequestId [32]byte
	Result    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRandomnessFulfilled is a free log retrieval operation binding the contract event 0x9b0aa3f92f46e24caa76b000bdf0dd495b9b390c320cf6585ae10a12b7d09edb.
//
// Solidity: event RandomnessFulfilled(bytes32 indexed requestId, uint256 result)
func (_VRF *VRFFilterer) FilterRandomnessFulfilled(opts *bind.FilterOpts, requestId [][32]byte) (*VRFRandomnessFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "RandomnessFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFRandomnessFulfilledIterator{contract: _VRF.contract, event: "RandomnessFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomnessFulfilled is a free log subscription operation binding the contract event 0x9b0aa3f92f46e24caa76b000bdf0dd495b9b390c320cf6585ae10a12b7d09edb.
//
// Solidity: event RandomnessFulfilled(bytes32 indexed requestId, uint256 result)
func (_VRF *VRFFilterer) WatchRandomnessFulfilled(opts *bind.WatchOpts, sink chan<- *VRFRandomnessFulfilled, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "RandomnessFulfilled", requestIdRule)
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

// ParseRandomnessFulfilled is a log parse operation binding the contract event 0x9b0aa3f92f46e24caa76b000bdf0dd495b9b390c320cf6585ae10a12b7d09edb.
//
// Solidity: event RandomnessFulfilled(bytes32 indexed requestId, uint256 result)
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
	RequestId [32]byte
	Requester common.Address
	Seed      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRandomnessRequested is a free log retrieval operation binding the contract event 0xadbe210fbb3f1373d938d421ee71b9db071c3d86ccaf8851d6bda23e70c6b532.
//
// Solidity: event RandomnessRequested(bytes32 indexed requestId, address indexed requester, bytes seed)
func (_VRF *VRFFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestId [][32]byte, requester []common.Address) (*VRFRandomnessRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "RandomnessRequested", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFRandomnessRequestedIterator{contract: _VRF.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

// WatchRandomnessRequested is a free log subscription operation binding the contract event 0xadbe210fbb3f1373d938d421ee71b9db071c3d86ccaf8851d6bda23e70c6b532.
//
// Solidity: event RandomnessRequested(bytes32 indexed requestId, address indexed requester, bytes seed)
func (_VRF *VRFFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFRandomnessRequested, requestId [][32]byte, requester []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "RandomnessRequested", requestIdRule, requesterRule)
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

// ParseRandomnessRequested is a log parse operation binding the contract event 0xadbe210fbb3f1373d938d421ee71b9db071c3d86ccaf8851d6bda23e70c6b532.
//
// Solidity: event RandomnessRequested(bytes32 indexed requestId, address indexed requester, bytes seed)
func (_VRF *VRFFilterer) ParseRandomnessRequested(log types.Log) (*VRFRandomnessRequested, error) {
	event := new(VRFRandomnessRequested)
	if err := _VRF.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFTaskResultSubmissionValidatedIterator is returned from FilterTaskResultSubmissionValidated and is used to iterate over the raw logs and unpacked data for TaskResultSubmissionValidated events raised by the VRF contract.
type VRFTaskResultSubmissionValidatedIterator struct {
	Event *VRFTaskResultSubmissionValidated // Event containing the contract specifics and raw log

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
func (it *VRFTaskResultSubmissionValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFTaskResultSubmissionValidated)
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
		it.Event = new(VRFTaskResultSubmissionValidated)
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
func (it *VRFTaskResultSubmissionValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFTaskResultSubmissionValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFTaskResultSubmissionValidated represents a TaskResultSubmissionValidated event raised by the VRF contract.
type VRFTaskResultSubmissionValidated struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskResultSubmissionValidated is a free log retrieval operation binding the contract event 0x0d5aeffd61ba930c83f1e88ec8bc110c50d1a38cf03190b0272c312d16ca7140.
//
// Solidity: event TaskResultSubmissionValidated(bytes32 indexed requestId)
func (_VRF *VRFFilterer) FilterTaskResultSubmissionValidated(opts *bind.FilterOpts, requestId [][32]byte) (*VRFTaskResultSubmissionValidatedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "TaskResultSubmissionValidated", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFTaskResultSubmissionValidatedIterator{contract: _VRF.contract, event: "TaskResultSubmissionValidated", logs: logs, sub: sub}, nil
}

// WatchTaskResultSubmissionValidated is a free log subscription operation binding the contract event 0x0d5aeffd61ba930c83f1e88ec8bc110c50d1a38cf03190b0272c312d16ca7140.
//
// Solidity: event TaskResultSubmissionValidated(bytes32 indexed requestId)
func (_VRF *VRFFilterer) WatchTaskResultSubmissionValidated(opts *bind.WatchOpts, sink chan<- *VRFTaskResultSubmissionValidated, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "TaskResultSubmissionValidated", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFTaskResultSubmissionValidated)
				if err := _VRF.contract.UnpackLog(event, "TaskResultSubmissionValidated", log); err != nil {
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

// ParseTaskResultSubmissionValidated is a log parse operation binding the contract event 0x0d5aeffd61ba930c83f1e88ec8bc110c50d1a38cf03190b0272c312d16ca7140.
//
// Solidity: event TaskResultSubmissionValidated(bytes32 indexed requestId)
func (_VRF *VRFFilterer) ParseTaskResultSubmissionValidated(log types.Log) (*VRFTaskResultSubmissionValidated, error) {
	event := new(VRFTaskResultSubmissionValidated)
	if err := _VRF.contract.UnpackLog(event, "TaskResultSubmissionValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
