// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRewardRate\",\"type\":\"uint256\"}],\"name\":\"RewardRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRewardRate\",\"type\":\"uint256\"}],\"name\":\"updateRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawStakedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405262c5c100670de0b6b3a764000061001b9190610254565b600455348015610029575f80fd5b50604051611a81380380611a81833981810160405281019061004b91906102e2565b8060015f819055505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100c3575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100ba9190610341565b60405180910390fd5b6100d28161015b60201b60201c565b508260025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505061035a565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61025e8261021e565b91506102698361021e565b92508261027957610278610227565b5b828204905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102b182610288565b9050919050565b6102c1816102a7565b81146102cb575f80fd5b50565b5f815190506102dc816102b8565b92915050565b5f805f606084860312156102f9576102f8610284565b5b5f610306868287016102ce565b9350506020610317868287016102ce565b9250506040610328868287016102ce565b9150509250925092565b61033b816102a7565b82525050565b5f6020820190506103545f830184610332565b92915050565b61171a806103675f395ff3fe608060405234801561000f575f80fd5b5060043610610108575f3560e01c80638b876347116100a0578063adca9b721161006f578063adca9b721461028e578063c8f33c91146102ac578063cd3daf9d146102ca578063df136d65146102e8578063f2fde38b1461030657610108565b80638b876347146102085780638da5cb5b146102385780639ef3a26114610256578063a694fc3a1461027257610108565b806360217267116100dc57806360217267146101925780636374cc00146101c2578063715018a6146101e05780637b0a47ee146101ea57610108565b80628cc2621461010c5780630700037d1461013c5780630a57336a1461016c5780633d18b91214610188575b5f80fd5b6101266004803603810190610121919061116c565b610322565b60405161013391906111af565b60405180910390f35b6101566004803603810190610151919061116c565b610439565b60405161016391906111af565b60405180910390f35b610186600480360381019061018191906111f2565b61044e565b005b610190610793565b005b6101ac60048036038101906101a7919061116c565b610a2f565b6040516101b991906111af565b60405180910390f35b6101ca610a44565b6040516101d79190611278565b60405180910390f35b6101e8610a69565b005b6101f2610a7c565b6040516101ff91906111af565b60405180910390f35b610222600480360381019061021d919061116c565b610a82565b60405161022f91906111af565b60405180910390f35b610240610a97565b60405161024d91906112a0565b60405180910390f35b610270600480360381019061026b91906111f2565b610abf565b005b61028c600480360381019061028791906111f2565b610b08565b005b610296610dcf565b6040516102a39190611278565b60405180910390f35b6102b4610df4565b6040516102c191906111af565b60405180910390f35b6102d2610dfa565b6040516102df91906111af565b60405180910390f35b6102f0610e89565b6040516102fd91906111af565b60405180910390f35b610320600480360381019061031b919061116c565b610e8f565b005b5f61043260095f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610424670de0b6b3a76400006104166103ca600a5f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546103bc610dfa565b610f1390919063ffffffff16565b60085f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610f2890919063ffffffff16565b610f3d90919063ffffffff16565b610f5290919063ffffffff16565b9050919050565b6009602052805f5260405f205f915090505481565b610456610f67565b3361045f610dfa565b6006819055504260078190555061047581610322565b60095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600654600a5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f821161053c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053390611313565b60405180910390fd5b8160085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156105bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b39061137b565b60405180910390fd5b6105d182600554610f1390919063ffffffff16565b6005819055506106278260085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610f1390919063ffffffff16565b60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550813373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d560405160405180910390a360025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401610708929190611399565b6020604051808303815f875af1158015610724573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061074891906113f5565b610787576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077e9061146a565b60405180910390fd5b50610790610fb4565b50565b61079b610f67565b336107a4610dfa565b600681905550426007819055506107ba81610322565b60095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600654600a5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f60095f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f81116108c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b9906114d2565b60405180910390fd5b5f60095f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550803373ffffffffffffffffffffffffffffffffffffffff167ffc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe60405160405180910390a360035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016109a4929190611399565b6020604051808303815f875af11580156109c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109e491906113f5565b610a23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1a9061146a565b60405180910390fd5b5050610a2d610fb4565b565b6008602052805f5260405f205f915090505481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610a71610fbd565b610a7a5f611044565b565b60045481565b600a602052805f5260405f205f915090505481565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610ac7610fbd565b806004819055507f41d466ebd06fb97e7786086ac8b69b7eb7da798592036251291d34e9791cde0181604051610afd91906111af565b60405180910390a150565b610b10610f67565b33610b19610dfa565b60068190555042600781905550610b2f81610322565b60095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600654600a5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f8211610bf6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bed90611313565b60405180910390fd5b610c0b82600554610f5290919063ffffffff16565b600581905550610c618260085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610f5290919063ffffffff16565b60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550813373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d60405160405180910390a360025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401610d44939291906114f0565b6020604051808303815f875af1158015610d60573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d8491906113f5565b610dc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dba9061146a565b60405180910390fd5b50610dcc610fb4565b50565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b5f8060055403610e0e576006549050610e86565b5f610e2460075442610f1390919063ffffffff16565b90505f610e3c82600454610f2890919063ffffffff16565b9050610e81610e70600554610e62670de0b6b3a764000085610f2890919063ffffffff16565b610f3d90919063ffffffff16565b600654610f5290919063ffffffff16565b925050505b90565b60065481565b610e97610fbd565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f07575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610efe91906112a0565b60405180910390fd5b610f1081611044565b50565b5f8183610f209190611552565b905092915050565b5f8183610f359190611585565b905092915050565b5f8183610f4a91906115f3565b905092915050565b5f8183610f5f9190611623565b905092915050565b60025f5403610fab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa2906116a0565b60405180910390fd5b60025f81905550565b60015f81905550565b610fc5611107565b73ffffffffffffffffffffffffffffffffffffffff16610fe3610a97565b73ffffffffffffffffffffffffffffffffffffffff161461104257611006611107565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161103991906112a0565b60405180910390fd5b565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61113b82611112565b9050919050565b61114b81611131565b8114611155575f80fd5b50565b5f8135905061116681611142565b92915050565b5f602082840312156111815761118061110e565b5b5f61118e84828501611158565b91505092915050565b5f819050919050565b6111a981611197565b82525050565b5f6020820190506111c25f8301846111a0565b92915050565b6111d181611197565b81146111db575f80fd5b50565b5f813590506111ec816111c8565b92915050565b5f602082840312156112075761120661110e565b5b5f611214848285016111de565b91505092915050565b5f819050919050565b5f61124061123b61123684611112565b61121d565b611112565b9050919050565b5f61125182611226565b9050919050565b5f61126282611247565b9050919050565b61127281611258565b82525050565b5f60208201905061128b5f830184611269565b92915050565b61129a81611131565b82525050565b5f6020820190506112b35f830184611291565b92915050565b5f82825260208201905092915050565b7f416d6f756e74206d7573742062652067726561746572207468616e207a65726f5f82015250565b5f6112fd6020836112b9565b9150611308826112c9565b602082019050919050565b5f6020820190508181035f83015261132a816112f1565b9050919050565b7f496e73756666696369656e74207374616b656420616d6f756e740000000000005f82015250565b5f611365601a836112b9565b915061137082611331565b602082019050919050565b5f6020820190508181035f83015261139281611359565b9050919050565b5f6040820190506113ac5f830185611291565b6113b960208301846111a0565b9392505050565b5f8115159050919050565b6113d4816113c0565b81146113de575f80fd5b50565b5f815190506113ef816113cb565b92915050565b5f6020828403121561140a5761140961110e565b5b5f611417848285016113e1565b91505092915050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f611454600f836112b9565b915061145f82611420565b602082019050919050565b5f6020820190508181035f83015261148181611448565b9050919050565b7f4e6f207265776172647320746f20636c61696d000000000000000000000000005f82015250565b5f6114bc6013836112b9565b91506114c782611488565b602082019050919050565b5f6020820190508181035f8301526114e9816114b0565b9050919050565b5f6060820190506115035f830186611291565b6115106020830185611291565b61151d60408301846111a0565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61155c82611197565b915061156783611197565b925082820390508181111561157f5761157e611525565b5b92915050565b5f61158f82611197565b915061159a83611197565b92508282026115a881611197565b915082820484148315176115bf576115be611525565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6115fd82611197565b915061160883611197565b925082611618576116176115c6565b5b828204905092915050565b5f61162d82611197565b915061163883611197565b92508282019050808211156116505761164f611525565b5b92915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c005f82015250565b5f61168a601f836112b9565b915061169582611656565b602082019050919050565b5f6020820190508181035f8301526116b78161167e565b905091905056fea2646970667358221220501f6c4bc5f5d2838c1f250e099bde43ceb14b3fc65e9ac930a1935ad3da5a3d64736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, stakingToken common.Address, rewardToken common.Address, initialOwner common.Address) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, stakingToken, rewardToken, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Staking *StakingCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Staking *StakingSession) Earned(account common.Address) (*big.Int, error) {
	return _Staking.Contract.Earned(&_Staking.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Staking *StakingCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _Staking.Contract.Earned(&_Staking.CallOpts, account)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Staking *StakingCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Staking *StakingSession) LastUpdateTime() (*big.Int, error) {
	return _Staking.Contract.LastUpdateTime(&_Staking.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Staking *StakingCallerSession) LastUpdateTime() (*big.Int, error) {
	return _Staking.Contract.LastUpdateTime(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Staking *StakingCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Staking *StakingSession) RewardPerToken() (*big.Int, error) {
	return _Staking.Contract.RewardPerToken(&_Staking.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Staking *StakingCallerSession) RewardPerToken() (*big.Int, error) {
	return _Staking.Contract.RewardPerToken(&_Staking.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Staking *StakingCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Staking *StakingSession) RewardPerTokenStored() (*big.Int, error) {
	return _Staking.Contract.RewardPerTokenStored(&_Staking.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Staking *StakingCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _Staking.Contract.RewardPerTokenStored(&_Staking.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Staking *StakingCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Staking *StakingSession) RewardRate() (*big.Int, error) {
	return _Staking.Contract.RewardRate(&_Staking.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Staking *StakingCallerSession) RewardRate() (*big.Int, error) {
	return _Staking.Contract.RewardRate(&_Staking.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Staking *StakingCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Staking *StakingSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.Rewards(&_Staking.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Staking *StakingCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.Rewards(&_Staking.CallOpts, arg0)
}

// SRewardToken is a free data retrieval call binding the contract method 0x6374cc00.
//
// Solidity: function s_rewardToken() view returns(address)
func (_Staking *StakingCaller) SRewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "s_rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SRewardToken is a free data retrieval call binding the contract method 0x6374cc00.
//
// Solidity: function s_rewardToken() view returns(address)
func (_Staking *StakingSession) SRewardToken() (common.Address, error) {
	return _Staking.Contract.SRewardToken(&_Staking.CallOpts)
}

// SRewardToken is a free data retrieval call binding the contract method 0x6374cc00.
//
// Solidity: function s_rewardToken() view returns(address)
func (_Staking *StakingCallerSession) SRewardToken() (common.Address, error) {
	return _Staking.Contract.SRewardToken(&_Staking.CallOpts)
}

// SStakingToken is a free data retrieval call binding the contract method 0xadca9b72.
//
// Solidity: function s_stakingToken() view returns(address)
func (_Staking *StakingCaller) SStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "s_stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SStakingToken is a free data retrieval call binding the contract method 0xadca9b72.
//
// Solidity: function s_stakingToken() view returns(address)
func (_Staking *StakingSession) SStakingToken() (common.Address, error) {
	return _Staking.Contract.SStakingToken(&_Staking.CallOpts)
}

// SStakingToken is a free data retrieval call binding the contract method 0xadca9b72.
//
// Solidity: function s_stakingToken() view returns(address)
func (_Staking *StakingCallerSession) SStakingToken() (common.Address, error) {
	return _Staking.Contract.SStakingToken(&_Staking.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address ) view returns(uint256)
func (_Staking *StakingCaller) StakedBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakedBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address ) view returns(uint256)
func (_Staking *StakingSession) StakedBalance(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.StakedBalance(&_Staking.CallOpts, arg0)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address ) view returns(uint256)
func (_Staking *StakingCallerSession) StakedBalance(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.StakedBalance(&_Staking.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Staking *StakingCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Staking *StakingSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.UserRewardPerTokenPaid(&_Staking.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Staking *StakingCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.UserRewardPerTokenPaid(&_Staking.CallOpts, arg0)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Staking *StakingTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Staking *StakingSession) GetReward() (*types.Transaction, error) {
	return _Staking.Contract.GetReward(&_Staking.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Staking *StakingTransactorSession) GetReward() (*types.Transaction, error) {
	return _Staking.Contract.GetReward(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Staking *StakingSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Staking *StakingTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x9ef3a261.
//
// Solidity: function updateRewardRate(uint256 newRewardRate) returns()
func (_Staking *StakingTransactor) UpdateRewardRate(opts *bind.TransactOpts, newRewardRate *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateRewardRate", newRewardRate)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x9ef3a261.
//
// Solidity: function updateRewardRate(uint256 newRewardRate) returns()
func (_Staking *StakingSession) UpdateRewardRate(newRewardRate *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateRewardRate(&_Staking.TransactOpts, newRewardRate)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x9ef3a261.
//
// Solidity: function updateRewardRate(uint256 newRewardRate) returns()
func (_Staking *StakingTransactorSession) UpdateRewardRate(newRewardRate *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateRewardRate(&_Staking.TransactOpts, newRewardRate)
}

// WithdrawStakedTokens is a paid mutator transaction binding the contract method 0x0a57336a.
//
// Solidity: function withdrawStakedTokens(uint256 amount) returns()
func (_Staking *StakingTransactor) WithdrawStakedTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdrawStakedTokens", amount)
}

// WithdrawStakedTokens is a paid mutator transaction binding the contract method 0x0a57336a.
//
// Solidity: function withdrawStakedTokens(uint256 amount) returns()
func (_Staking *StakingSession) WithdrawStakedTokens(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawStakedTokens(&_Staking.TransactOpts, amount)
}

// WithdrawStakedTokens is a paid mutator transaction binding the contract method 0x0a57336a.
//
// Solidity: function withdrawStakedTokens(uint256 amount) returns()
func (_Staking *StakingTransactorSession) WithdrawStakedTokens(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawStakedTokens(&_Staking.TransactOpts, amount)
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardRateUpdatedIterator is returned from FilterRewardRateUpdated and is used to iterate over the raw logs and unpacked data for RewardRateUpdated events raised by the Staking contract.
type StakingRewardRateUpdatedIterator struct {
	Event *StakingRewardRateUpdated // Event containing the contract specifics and raw log

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
func (it *StakingRewardRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardRateUpdated)
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
		it.Event = new(StakingRewardRateUpdated)
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
func (it *StakingRewardRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardRateUpdated represents a RewardRateUpdated event raised by the Staking contract.
type StakingRewardRateUpdated struct {
	NewRewardRate *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardRateUpdated is a free log retrieval operation binding the contract event 0x41d466ebd06fb97e7786086ac8b69b7eb7da798592036251291d34e9791cde01.
//
// Solidity: event RewardRateUpdated(uint256 newRewardRate)
func (_Staking *StakingFilterer) FilterRewardRateUpdated(opts *bind.FilterOpts) (*StakingRewardRateUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardRateUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingRewardRateUpdatedIterator{contract: _Staking.contract, event: "RewardRateUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardRateUpdated is a free log subscription operation binding the contract event 0x41d466ebd06fb97e7786086ac8b69b7eb7da798592036251291d34e9791cde01.
//
// Solidity: event RewardRateUpdated(uint256 newRewardRate)
func (_Staking *StakingFilterer) WatchRewardRateUpdated(opts *bind.WatchOpts, sink chan<- *StakingRewardRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardRateUpdated)
				if err := _Staking.contract.UnpackLog(event, "RewardRateUpdated", log); err != nil {
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

// ParseRewardRateUpdated is a log parse operation binding the contract event 0x41d466ebd06fb97e7786086ac8b69b7eb7da798592036251291d34e9791cde01.
//
// Solidity: event RewardRateUpdated(uint256 newRewardRate)
func (_Staking *StakingFilterer) ParseRewardRateUpdated(log types.Log) (*StakingRewardRateUpdated, error) {
	event := new(StakingRewardRateUpdated)
	if err := _Staking.contract.UnpackLog(event, "RewardRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Staking contract.
type StakingRewardsClaimedIterator struct {
	Event *StakingRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *StakingRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardsClaimed)
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
		it.Event = new(StakingRewardsClaimed)
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
func (it *StakingRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardsClaimed represents a RewardsClaimed event raised by the Staking contract.
type StakingRewardsClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 indexed amount)
func (_Staking *StakingFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*StakingRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardsClaimed", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardsClaimedIterator{contract: _Staking.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 indexed amount)
func (_Staking *StakingFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *StakingRewardsClaimed, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardsClaimed", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardsClaimed)
				if err := _Staking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 indexed amount)
func (_Staking *StakingFilterer) ParseRewardsClaimed(log types.Log) (*StakingRewardsClaimed, error) {
	event := new(StakingRewardsClaimed)
	if err := _Staking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Staking contract.
type StakingStakedIterator struct {
	Event *StakingStaked // Event containing the contract specifics and raw log

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
func (it *StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStaked)
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
		it.Event = new(StakingStaked)
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
func (it *StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStaked represents a Staked event raised by the Staking contract.
type StakingStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 indexed amount)
func (_Staking *StakingFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*StakingStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Staked", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakedIterator{contract: _Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 indexed amount)
func (_Staking *StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingStaked, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Staked", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStaked)
				if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 indexed amount)
func (_Staking *StakingFilterer) ParseStaked(log types.Log) (*StakingStaked, error) {
	event := new(StakingStaked)
	if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Staking contract.
type StakingWithdrawnIterator struct {
	Event *StakingWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdrawn)
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
		it.Event = new(StakingWithdrawn)
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
func (it *StakingWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdrawn represents a Withdrawn event raised by the Staking contract.
type StakingWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed amount)
func (_Staking *StakingFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*StakingWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Withdrawn", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawnIterator{contract: _Staking.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed amount)
func (_Staking *StakingFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingWithdrawn, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Withdrawn", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdrawn)
				if err := _Staking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed amount)
func (_Staking *StakingFilterer) ParseWithdrawn(log types.Log) (*StakingWithdrawn, error) {
	event := new(StakingWithdrawn)
	if err := _Staking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
