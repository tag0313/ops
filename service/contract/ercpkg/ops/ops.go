// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ops

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OpsABI is the input ABI used to generate the binding from.
const OpsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addMinter\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delMinter\",\"type\":\"address\"}],\"name\":\"delMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinterLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OpsBin is the compiled bytecode used for deploying new contracts.
var OpsBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600381526020017f4f707300000000000000000000000000000000000000000000000000000000008152506040518060400160405280600981526020017f4f707320546f6b656e000000000000000000000000000000000000000000000081525081600390805190602001906200009692919062000461565b508060049080519060200190620000af92919062000461565b506012600560006101000a81548160ff021916908360ff16021790555050506000620000e0620001ed60201b60201c565b905080600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350601260ff16600a0a6301406f4002601260ff16600a0a621e8480021115620001a657600080fd5b620001d3620001ba620001ed60201b60201c565b601260ff16600a0a621e848002620001f560201b60201c565b601260ff16600a0a6301406f400260088190555062000510565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000299576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b620002ad60008383620003d360201b60201c565b620002c981600254620003d860201b62001ccd1790919060201c565b60028190555062000327816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054620003d860201b62001ccd1790919060201c565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b505050565b60008082840190508381101562000457576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004a457805160ff1916838001178555620004d5565b82800160010185558215620004d5579182015b82811115620004d4578251825591602001919060010190620004b7565b5b509050620004e49190620004e8565b5090565b6200050d91905b8082111562000509576000816000905550600101620004ef565b5090565b90565b6125d780620005206000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c806370a08231116100c3578063a457c2d71161007c578063a457c2d7146106cf578063a9059cbb14610735578063aa271e1a1461079b578063dd62ed3e146107f7578063f26539861461086f578063f2fde38b1461088d57610158565b806370a08231146104f6578063715018a61461054e57806379cc6790146105585780638da5cb5b146105a657806395d89b41146105f0578063983b2d561461067357610158565b80632ff2e9dc116101155780632ff2e9dc14610364578063313ce5671461038257806339509351146103a657806340c10f191461040c57806342966c681461045a5780635b7121f81461048857610158565b80630323aac71461015d57806306fdde031461017b578063095ea7b3146101fe57806318160ddd1461026457806323338b881461028257806323b872dd146102de575b600080fd5b6101656108d1565b6040518082815260200191505060405180910390f35b6101836108e2565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101c35780820151818401526020810190506101a8565b50505050905090810190601f1680156101f05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61024a6004803603604081101561021457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610984565b604051808215151515815260200191505060405180910390f35b61026c6109a2565b6040518082815260200191505060405180910390f35b6102c46004803603602081101561029857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109ac565b604051808215151515815260200191505060405180910390f35b61034a600480360360608110156102f457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b12565b604051808215151515815260200191505060405180910390f35b61036c610beb565b6040518082815260200191505060405180910390f35b61038a610bfb565b604051808260ff1660ff16815260200191505060405180910390f35b6103f2600480360360408110156103bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c12565b604051808215151515815260200191505060405180910390f35b6104586004803603604081101561042257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cc5565b005b6104866004803603602081101561047057600080fd5b8101908080359060200190929190505050610ddf565b005b6104b46004803603602081101561049e57600080fd5b8101908080359060200190929190505050610df3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105386004803603602081101561050c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f36565b6040518082815260200191505060405180910390f35b610556610f7e565b005b6105a46004803603604081101561056e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110ee565b005b6105ae611150565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105f861117a565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561063857808201518184015260208101905061061d565b50505050905090810190601f1680156106655780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106b56004803603602081101561068957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061121c565b604051808215151515815260200191505060405180910390f35b61071b600480360360408110156106e557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611382565b604051808215151515815260200191505060405180910390f35b6107816004803603604081101561074b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061144f565b604051808215151515815260200191505060405180910390f35b6107dd600480360360208110156107b157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061146d565b604051808215151515815260200191505060405180910390f35b6108596004803603604081101561080d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611481565b6040518082815260200191505060405180910390f35b610877611508565b6040518082815260200191505060405180910390f35b6108cf600480360360208110156108a357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611519565b005b60006108dd600661170e565b905090565b606060038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561097a5780601f1061094f5761010080835404028352916020019161097a565b820191906000526020600020905b81548152906001019060200180831161095d57829003601f168201915b5050505050905090565b6000610998610991611723565b848461172b565b6001905092915050565b6000600254905090565b60006109b6611723565b73ffffffffffffffffffffffffffffffffffffffff166109d4611150565b73ffffffffffffffffffffffffffffffffffffffff1614610a5d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b00576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f5f64656c4d696e74657220697320746865207a65726f2061646472657373000081525060200191505060405180910390fd5b610b0b600683611922565b9050919050565b6000610b1f848484611952565b610be084610b2b611723565b610bdb856040518060600160405280602881526020016124c760289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610b91611723565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c139092919063ffffffff16565b61172b565b600190509392505050565b601260ff16600a0a621e84800281565b6000600560009054906101000a900460ff16905090565b6000610cbb610c1f611723565b84610cb68560016000610c30611723565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611ccd90919063ffffffff16565b61172b565b6001905092915050565b610cce3361146d565b610d40576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f63616c6c6572206973206e6f7420746865206d696e746572000000000000000081525060200191505060405180910390fd5b600854610d5d82610d4f6109a2565b611ccd90919063ffffffff16565b1115610dd1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f45524332304361707065643a206361702065786365656465640000000000000081525060200191505060405180910390fd5b610ddb8282611d55565b5050565b610df0610dea611723565b82611f1c565b50565b6000610dfd611723565b73ffffffffffffffffffffffffffffffffffffffff16610e1b611150565b73ffffffffffffffffffffffffffffffffffffffff1614610ea4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6001610eae6108d1565b03821115610f24576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f696e646578206f7574206f6620626f756e64730000000000000000000000000081525060200191505060405180910390fd5b610f2f6006836120e0565b9050919050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610f86611723565b73ffffffffffffffffffffffffffffffffffffffff16610fa4611150565b73ffffffffffffffffffffffffffffffffffffffff161461102d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600061112d826040518060600160405280602481526020016124ef6024913961111e86611119611723565b611481565b611c139092919063ffffffff16565b90506111418361113b611723565b8361172b565b61114b8383611f1c565b505050565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112125780601f106111e757610100808354040283529160200191611212565b820191906000526020600020905b8154815290600101906020018083116111f557829003601f168201915b5050505050905090565b6000611226611723565b73ffffffffffffffffffffffffffffffffffffffff16611244611150565b73ffffffffffffffffffffffffffffffffffffffff16146112cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611370576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f205f6164644d696e74657220697320746865207a65726f20616464726573730081525060200191505060405180910390fd5b61137b6006836120fa565b9050919050565b600061144561138f611723565b846114408560405180606001604052806025815260200161257d60259139600160006113b9611723565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c139092919063ffffffff16565b61172b565b6001905092915050565b600061146361145c611723565b8484611952565b6001905092915050565b600061147a60068361212a565b9050919050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b601260ff16600a0a6301406f400281565b611521611723565b73ffffffffffffffffffffffffffffffffffffffff1661153f611150565b73ffffffffffffffffffffffffffffffffffffffff16146115c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561164e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806124596026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600061171c8260000161215a565b9050919050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156117b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806125596024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611837576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061247f6022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600061194a836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61216b565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156119d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806125346025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611a5e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806124146023913960400191505060405180910390fd5b611a69838383612253565b611ad4816040518060600160405280602681526020016124a1602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c139092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611b67816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611ccd90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290611cc0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611c85578082015181840152602081019050611c6a565b50505050905090810190601f168015611cb25780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082840390509392505050565b600080828401905083811015611d4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611df8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b611e0460008383612253565b611e1981600254611ccd90919063ffffffff16565b600281905550611e70816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611ccd90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611fa2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806125136021913960400191505060405180910390fd5b611fae82600083612253565b61201981604051806060016040528060228152602001612437602291396000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c139092919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506120708160025461225890919063ffffffff16565b600281905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b60006120ef83600001836122db565b60001c905092915050565b6000612122836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61235e565b905092915050565b6000612152836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6123ce565b905092915050565b600081600001805490509050919050565b6000808360010160008481526020019081526020016000205490506000811461224757600060018203905060006001866000018054905003905060008660000182815481106121b657fe5b90600052602060002001549050808760000184815481106121d357fe5b906000526020600020018190555060018301876001016000838152602001908152602001600020819055508660000180548061220b57fe5b6001900381819060005260206000200160009055905586600101600087815260200190815260200160002060009055600194505050505061224d565b60009150505b92915050565b505050565b6000828211156122d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b818303905092915050565b60008183600001805490501161233c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806123f26022913960400191505060405180910390fd5b82600001828154811061234b57fe5b9060005260206000200154905092915050565b600061236a83836123ce565b6123c35782600001829080600181540180825580915050600190039060005260206000200160009091909190915055826000018054905083600101600084815260200190815260200160002081905550600190506123c8565b600090505b92915050565b60008083600101600084815260200190815260200160002054141590509291505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e647345524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e63654f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122048f2ceaded1d79eeefd304c0cad15cdb9c50bce45541b03a9ddc58038573200f64736f6c63430006020033"

// DeployOps deploys a new Ethereum contract, binding an instance of Ops to it.
func DeployOps(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ops, error) {
	parsed, err := abi.JSON(strings.NewReader(OpsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OpsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ops{OpsCaller: OpsCaller{contract: contract}, OpsTransactor: OpsTransactor{contract: contract}, OpsFilterer: OpsFilterer{contract: contract}}, nil
}

// Ops is an auto generated Go binding around an Ethereum contract.
type Ops struct {
	OpsCaller     // Read-only binding to the contract
	OpsTransactor // Write-only binding to the contract
	OpsFilterer   // Log filterer for contract events
}

// OpsCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpsSession struct {
	Contract     *Ops              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpsCallerSession struct {
	Contract *OpsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OpsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpsTransactorSession struct {
	Contract     *OpsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpsRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpsRaw struct {
	Contract *Ops // Generic contract binding to access the raw methods on
}

// OpsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpsCallerRaw struct {
	Contract *OpsCaller // Generic read-only contract binding to access the raw methods on
}

// OpsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpsTransactorRaw struct {
	Contract *OpsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOps creates a new instance of Ops, bound to a specific deployed contract.
func NewOps(address common.Address, backend bind.ContractBackend) (*Ops, error) {
	contract, err := bindOps(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ops{OpsCaller: OpsCaller{contract: contract}, OpsTransactor: OpsTransactor{contract: contract}, OpsFilterer: OpsFilterer{contract: contract}}, nil
}

// NewOpsCaller creates a new read-only instance of Ops, bound to a specific deployed contract.
func NewOpsCaller(address common.Address, caller bind.ContractCaller) (*OpsCaller, error) {
	contract, err := bindOps(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpsCaller{contract: contract}, nil
}

// NewOpsTransactor creates a new write-only instance of Ops, bound to a specific deployed contract.
func NewOpsTransactor(address common.Address, transactor bind.ContractTransactor) (*OpsTransactor, error) {
	contract, err := bindOps(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpsTransactor{contract: contract}, nil
}

// NewOpsFilterer creates a new log filterer instance of Ops, bound to a specific deployed contract.
func NewOpsFilterer(address common.Address, filterer bind.ContractFilterer) (*OpsFilterer, error) {
	contract, err := bindOps(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpsFilterer{contract: contract}, nil
}

// bindOps binds a generic wrapper to an already deployed contract.
func bindOps(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ops *OpsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ops.Contract.OpsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ops *OpsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ops.Contract.OpsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ops *OpsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ops.Contract.OpsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ops *OpsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ops.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ops *OpsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ops.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ops *OpsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ops.Contract.contract.Transact(opts, method, params...)
}

// INITIALCAP is a free data retrieval call binding the contract method 0xf2653986.
//
// Solidity: function INITIAL_CAP() view returns(uint256)
func (_Ops *OpsCaller) INITIALCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "INITIAL_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALCAP is a free data retrieval call binding the contract method 0xf2653986.
//
// Solidity: function INITIAL_CAP() view returns(uint256)
func (_Ops *OpsSession) INITIALCAP() (*big.Int, error) {
	return _Ops.Contract.INITIALCAP(&_Ops.CallOpts)
}

// INITIALCAP is a free data retrieval call binding the contract method 0xf2653986.
//
// Solidity: function INITIAL_CAP() view returns(uint256)
func (_Ops *OpsCallerSession) INITIALCAP() (*big.Int, error) {
	return _Ops.Contract.INITIALCAP(&_Ops.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Ops *OpsCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "INITIAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Ops *OpsSession) INITIALSUPPLY() (*big.Int, error) {
	return _Ops.Contract.INITIALSUPPLY(&_Ops.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Ops *OpsCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _Ops.Contract.INITIALSUPPLY(&_Ops.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ops *OpsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ops *OpsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ops.Contract.Allowance(&_Ops.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ops *OpsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ops.Contract.Allowance(&_Ops.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ops *OpsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ops *OpsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ops.Contract.BalanceOf(&_Ops.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ops *OpsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ops.Contract.BalanceOf(&_Ops.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ops *OpsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ops *OpsSession) Decimals() (uint8, error) {
	return _Ops.Contract.Decimals(&_Ops.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ops *OpsCallerSession) Decimals() (uint8, error) {
	return _Ops.Contract.Decimals(&_Ops.CallOpts)
}

// GetMinter is a free data retrieval call binding the contract method 0x5b7121f8.
//
// Solidity: function getMinter(uint256 _index) view returns(address)
func (_Ops *OpsCaller) GetMinter(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "getMinter", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMinter is a free data retrieval call binding the contract method 0x5b7121f8.
//
// Solidity: function getMinter(uint256 _index) view returns(address)
func (_Ops *OpsSession) GetMinter(_index *big.Int) (common.Address, error) {
	return _Ops.Contract.GetMinter(&_Ops.CallOpts, _index)
}

// GetMinter is a free data retrieval call binding the contract method 0x5b7121f8.
//
// Solidity: function getMinter(uint256 _index) view returns(address)
func (_Ops *OpsCallerSession) GetMinter(_index *big.Int) (common.Address, error) {
	return _Ops.Contract.GetMinter(&_Ops.CallOpts, _index)
}

// GetMinterLength is a free data retrieval call binding the contract method 0x0323aac7.
//
// Solidity: function getMinterLength() view returns(uint256)
func (_Ops *OpsCaller) GetMinterLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "getMinterLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinterLength is a free data retrieval call binding the contract method 0x0323aac7.
//
// Solidity: function getMinterLength() view returns(uint256)
func (_Ops *OpsSession) GetMinterLength() (*big.Int, error) {
	return _Ops.Contract.GetMinterLength(&_Ops.CallOpts)
}

// GetMinterLength is a free data retrieval call binding the contract method 0x0323aac7.
//
// Solidity: function getMinterLength() view returns(uint256)
func (_Ops *OpsCallerSession) GetMinterLength() (*big.Int, error) {
	return _Ops.Contract.GetMinterLength(&_Ops.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Ops *OpsCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Ops *OpsSession) IsMinter(account common.Address) (bool, error) {
	return _Ops.Contract.IsMinter(&_Ops.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Ops *OpsCallerSession) IsMinter(account common.Address) (bool, error) {
	return _Ops.Contract.IsMinter(&_Ops.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ops *OpsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ops *OpsSession) Name() (string, error) {
	return _Ops.Contract.Name(&_Ops.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ops *OpsCallerSession) Name() (string, error) {
	return _Ops.Contract.Name(&_Ops.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ops *OpsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ops *OpsSession) Owner() (common.Address, error) {
	return _Ops.Contract.Owner(&_Ops.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ops *OpsCallerSession) Owner() (common.Address, error) {
	return _Ops.Contract.Owner(&_Ops.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ops *OpsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ops *OpsSession) Symbol() (string, error) {
	return _Ops.Contract.Symbol(&_Ops.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ops *OpsCallerSession) Symbol() (string, error) {
	return _Ops.Contract.Symbol(&_Ops.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ops *OpsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ops.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ops *OpsSession) TotalSupply() (*big.Int, error) {
	return _Ops.Contract.TotalSupply(&_Ops.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ops *OpsCallerSession) TotalSupply() (*big.Int, error) {
	return _Ops.Contract.TotalSupply(&_Ops.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _addMinter) returns(bool)
func (_Ops *OpsTransactor) AddMinter(opts *bind.TransactOpts, _addMinter common.Address) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "addMinter", _addMinter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _addMinter) returns(bool)
func (_Ops *OpsSession) AddMinter(_addMinter common.Address) (*types.Transaction, error) {
	return _Ops.Contract.AddMinter(&_Ops.TransactOpts, _addMinter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _addMinter) returns(bool)
func (_Ops *OpsTransactorSession) AddMinter(_addMinter common.Address) (*types.Transaction, error) {
	return _Ops.Contract.AddMinter(&_Ops.TransactOpts, _addMinter)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ops *OpsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ops *OpsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.Approve(&_Ops.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ops *OpsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.Approve(&_Ops.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Ops *OpsTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Ops *OpsSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.Burn(&_Ops.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Ops *OpsTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.Burn(&_Ops.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Ops *OpsTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Ops *OpsSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.BurnFrom(&_Ops.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Ops *OpsTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.BurnFrom(&_Ops.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Ops *OpsTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Ops *OpsSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.DecreaseAllowance(&_Ops.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Ops *OpsTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.DecreaseAllowance(&_Ops.TransactOpts, spender, subtractedValue)
}

// DelMinter is a paid mutator transaction binding the contract method 0x23338b88.
//
// Solidity: function delMinter(address _delMinter) returns(bool)
func (_Ops *OpsTransactor) DelMinter(opts *bind.TransactOpts, _delMinter common.Address) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "delMinter", _delMinter)
}

// DelMinter is a paid mutator transaction binding the contract method 0x23338b88.
//
// Solidity: function delMinter(address _delMinter) returns(bool)
func (_Ops *OpsSession) DelMinter(_delMinter common.Address) (*types.Transaction, error) {
	return _Ops.Contract.DelMinter(&_Ops.TransactOpts, _delMinter)
}

// DelMinter is a paid mutator transaction binding the contract method 0x23338b88.
//
// Solidity: function delMinter(address _delMinter) returns(bool)
func (_Ops *OpsTransactorSession) DelMinter(_delMinter common.Address) (*types.Transaction, error) {
	return _Ops.Contract.DelMinter(&_Ops.TransactOpts, _delMinter)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Ops *OpsTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Ops *OpsSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.IncreaseAllowance(&_Ops.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Ops *OpsTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.IncreaseAllowance(&_Ops.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Ops *OpsTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Ops *OpsSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.Mint(&_Ops.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Ops *OpsTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.Mint(&_Ops.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ops *OpsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ops *OpsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ops.Contract.RenounceOwnership(&_Ops.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ops *OpsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ops.Contract.RenounceOwnership(&_Ops.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Ops *OpsTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Ops *OpsSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.Transfer(&_Ops.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Ops *OpsTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.Transfer(&_Ops.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Ops *OpsTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Ops *OpsSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.TransferFrom(&_Ops.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Ops *OpsTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ops.Contract.TransferFrom(&_Ops.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ops *OpsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ops.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ops *OpsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ops.Contract.TransferOwnership(&_Ops.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ops *OpsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ops.Contract.TransferOwnership(&_Ops.TransactOpts, newOwner)
}

// OpsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ops contract.
type OpsApprovalIterator struct {
	Event *OpsApproval // Event containing the contract specifics and raw log

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
func (it *OpsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpsApproval)
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
		it.Event = new(OpsApproval)
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
func (it *OpsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpsApproval represents a Approval event raised by the Ops contract.
type OpsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ops *OpsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OpsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ops.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OpsApprovalIterator{contract: _Ops.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ops *OpsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OpsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ops.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpsApproval)
				if err := _Ops.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ops *OpsFilterer) ParseApproval(log types.Log) (*OpsApproval, error) {
	event := new(OpsApproval)
	if err := _Ops.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ops contract.
type OpsOwnershipTransferredIterator struct {
	Event *OpsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OpsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpsOwnershipTransferred)
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
		it.Event = new(OpsOwnershipTransferred)
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
func (it *OpsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpsOwnershipTransferred represents a OwnershipTransferred event raised by the Ops contract.
type OpsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ops *OpsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OpsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ops.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OpsOwnershipTransferredIterator{contract: _Ops.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ops *OpsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OpsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ops.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpsOwnershipTransferred)
				if err := _Ops.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ops *OpsFilterer) ParseOwnershipTransferred(log types.Log) (*OpsOwnershipTransferred, error) {
	event := new(OpsOwnershipTransferred)
	if err := _Ops.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ops contract.
type OpsTransferIterator struct {
	Event *OpsTransfer // Event containing the contract specifics and raw log

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
func (it *OpsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpsTransfer)
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
		it.Event = new(OpsTransfer)
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
func (it *OpsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpsTransfer represents a Transfer event raised by the Ops contract.
type OpsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ops *OpsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OpsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ops.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OpsTransferIterator{contract: _Ops.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ops *OpsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OpsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ops.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpsTransfer)
				if err := _Ops.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ops *OpsFilterer) ParseTransfer(log types.Log) (*OpsTransfer, error) {
	event := new(OpsTransfer)
	if err := _Ops.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
