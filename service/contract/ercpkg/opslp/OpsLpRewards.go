// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opslp

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

// OpslpABI is the input ABI used to generate the binding from.
const OpslpABI = "[{\"inputs\":[{\"internalType\":\"contractOpsErc20\",\"name\":\"_ops\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_opsPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_feeAddr\",\"type\":\"address\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"acct\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetFeeRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_migrAddr\",\"type\":\"address\"}],\"name\":\"SetMigrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdateMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"acct\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdateOpsPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BONUS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeAddr\",\"type\":\"address\"}],\"name\":\"fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ops\",\"outputs\":[{\"internalType\":\"contractOpsErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"opsPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingOps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accOpsPerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeRatio_\",\"type\":\"uint256\"}],\"name\":\"setFeeRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiplierNumber\",\"type\":\"uint256\"}],\"name\":\"updateMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_opsPerBlock\",\"type\":\"uint256\"}],\"name\":\"updateOpsPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Opslp is an auto generated Go binding around an Ethereum contract.
type Opslp struct {
	OpslpCaller     // Read-only binding to the contract
	OpslpTransactor // Write-only binding to the contract
	OpslpFilterer   // Log filterer for contract events
}

// OpslpCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpslpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpslpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpslpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpslpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpslpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpslpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpslpSession struct {
	Contract     *Opslp            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpslpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpslpCallerSession struct {
	Contract *OpslpCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OpslpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpslpTransactorSession struct {
	Contract     *OpslpTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpslpRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpslpRaw struct {
	Contract *Opslp // Generic contract binding to access the raw methods on
}

// OpslpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpslpCallerRaw struct {
	Contract *OpslpCaller // Generic read-only contract binding to access the raw methods on
}

// OpslpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpslpTransactorRaw struct {
	Contract *OpslpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpslp creates a new instance of Opslp, bound to a specific deployed contract.
func NewOpslp(address common.Address, backend bind.ContractBackend) (*Opslp, error) {
	contract, err := bindOpslp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Opslp{OpslpCaller: OpslpCaller{contract: contract}, OpslpTransactor: OpslpTransactor{contract: contract}, OpslpFilterer: OpslpFilterer{contract: contract}}, nil
}

// NewOpslpCaller creates a new read-only instance of Opslp, bound to a specific deployed contract.
func NewOpslpCaller(address common.Address, caller bind.ContractCaller) (*OpslpCaller, error) {
	contract, err := bindOpslp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpslpCaller{contract: contract}, nil
}

// NewOpslpTransactor creates a new write-only instance of Opslp, bound to a specific deployed contract.
func NewOpslpTransactor(address common.Address, transactor bind.ContractTransactor) (*OpslpTransactor, error) {
	contract, err := bindOpslp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpslpTransactor{contract: contract}, nil
}

// NewOpslpFilterer creates a new log filterer instance of Opslp, bound to a specific deployed contract.
func NewOpslpFilterer(address common.Address, filterer bind.ContractFilterer) (*OpslpFilterer, error) {
	contract, err := bindOpslp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpslpFilterer{contract: contract}, nil
}

// bindOpslp binds a generic wrapper to an already deployed contract.
func bindOpslp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpslpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opslp *OpslpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opslp.Contract.OpslpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opslp *OpslpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opslp.Contract.OpslpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opslp *OpslpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opslp.Contract.OpslpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opslp *OpslpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opslp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opslp *OpslpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opslp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opslp *OpslpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opslp.Contract.contract.Transact(opts, method, params...)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Opslp *OpslpCaller) BONUSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "BONUS_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Opslp *OpslpSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Opslp.Contract.BONUSMULTIPLIER(&_Opslp.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Opslp *OpslpCallerSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Opslp.Contract.BONUSMULTIPLIER(&_Opslp.CallOpts)
}

// FeeAddr is a free data retrieval call binding the contract method 0x39e7fddc.
//
// Solidity: function feeAddr() view returns(address)
func (_Opslp *OpslpCaller) FeeAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "feeAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddr is a free data retrieval call binding the contract method 0x39e7fddc.
//
// Solidity: function feeAddr() view returns(address)
func (_Opslp *OpslpSession) FeeAddr() (common.Address, error) {
	return _Opslp.Contract.FeeAddr(&_Opslp.CallOpts)
}

// FeeAddr is a free data retrieval call binding the contract method 0x39e7fddc.
//
// Solidity: function feeAddr() view returns(address)
func (_Opslp *OpslpCallerSession) FeeAddr() (common.Address, error) {
	return _Opslp.Contract.FeeAddr(&_Opslp.CallOpts)
}

// FeeRatio is a free data retrieval call binding the contract method 0x41744dd4.
//
// Solidity: function feeRatio() view returns(uint256)
func (_Opslp *OpslpCaller) FeeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "feeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRatio is a free data retrieval call binding the contract method 0x41744dd4.
//
// Solidity: function feeRatio() view returns(uint256)
func (_Opslp *OpslpSession) FeeRatio() (*big.Int, error) {
	return _Opslp.Contract.FeeRatio(&_Opslp.CallOpts)
}

// FeeRatio is a free data retrieval call binding the contract method 0x41744dd4.
//
// Solidity: function feeRatio() view returns(uint256)
func (_Opslp *OpslpCallerSession) FeeRatio() (*big.Int, error) {
	return _Opslp.Contract.FeeRatio(&_Opslp.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Opslp *OpslpCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Opslp *OpslpSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Opslp.Contract.GetMultiplier(&_Opslp.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Opslp *OpslpCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Opslp.Contract.GetMultiplier(&_Opslp.CallOpts, _from, _to)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Opslp *OpslpCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Opslp *OpslpSession) Migrator() (common.Address, error) {
	return _Opslp.Contract.Migrator(&_Opslp.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Opslp *OpslpCallerSession) Migrator() (common.Address, error) {
	return _Opslp.Contract.Migrator(&_Opslp.CallOpts)
}

// Ops is a free data retrieval call binding the contract method 0xe70abe92.
//
// Solidity: function ops() view returns(address)
func (_Opslp *OpslpCaller) Ops(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "ops")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ops is a free data retrieval call binding the contract method 0xe70abe92.
//
// Solidity: function ops() view returns(address)
func (_Opslp *OpslpSession) Ops() (common.Address, error) {
	return _Opslp.Contract.Ops(&_Opslp.CallOpts)
}

// Ops is a free data retrieval call binding the contract method 0xe70abe92.
//
// Solidity: function ops() view returns(address)
func (_Opslp *OpslpCallerSession) Ops() (common.Address, error) {
	return _Opslp.Contract.Ops(&_Opslp.CallOpts)
}

// OpsPerBlock is a free data retrieval call binding the contract method 0x611592e8.
//
// Solidity: function opsPerBlock() view returns(uint256)
func (_Opslp *OpslpCaller) OpsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "opsPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpsPerBlock is a free data retrieval call binding the contract method 0x611592e8.
//
// Solidity: function opsPerBlock() view returns(uint256)
func (_Opslp *OpslpSession) OpsPerBlock() (*big.Int, error) {
	return _Opslp.Contract.OpsPerBlock(&_Opslp.CallOpts)
}

// OpsPerBlock is a free data retrieval call binding the contract method 0x611592e8.
//
// Solidity: function opsPerBlock() view returns(uint256)
func (_Opslp *OpslpCallerSession) OpsPerBlock() (*big.Int, error) {
	return _Opslp.Contract.OpsPerBlock(&_Opslp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Opslp *OpslpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Opslp *OpslpSession) Owner() (common.Address, error) {
	return _Opslp.Contract.Owner(&_Opslp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Opslp *OpslpCallerSession) Owner() (common.Address, error) {
	return _Opslp.Contract.Owner(&_Opslp.CallOpts)
}

// PendingOps is a free data retrieval call binding the contract method 0xb2fdc061.
//
// Solidity: function pendingOps(uint256 _pid, address _user) view returns(uint256)
func (_Opslp *OpslpCaller) PendingOps(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "pendingOps", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingOps is a free data retrieval call binding the contract method 0xb2fdc061.
//
// Solidity: function pendingOps(uint256 _pid, address _user) view returns(uint256)
func (_Opslp *OpslpSession) PendingOps(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Opslp.Contract.PendingOps(&_Opslp.CallOpts, _pid, _user)
}

// PendingOps is a free data retrieval call binding the contract method 0xb2fdc061.
//
// Solidity: function pendingOps(uint256 _pid, address _user) view returns(uint256)
func (_Opslp *OpslpCallerSession) PendingOps(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Opslp.Contract.PendingOps(&_Opslp.CallOpts, _pid, _user)
}

// PoolExist is a free data retrieval call binding the contract method 0x89345efb.
//
// Solidity: function poolExist(address ) view returns(bool)
func (_Opslp *OpslpCaller) PoolExist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "poolExist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolExist is a free data retrieval call binding the contract method 0x89345efb.
//
// Solidity: function poolExist(address ) view returns(bool)
func (_Opslp *OpslpSession) PoolExist(arg0 common.Address) (bool, error) {
	return _Opslp.Contract.PoolExist(&_Opslp.CallOpts, arg0)
}

// PoolExist is a free data retrieval call binding the contract method 0x89345efb.
//
// Solidity: function poolExist(address ) view returns(bool)
func (_Opslp *OpslpCallerSession) PoolExist(arg0 common.Address) (bool, error) {
	return _Opslp.Contract.PoolExist(&_Opslp.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accOpsPerShare)
func (_Opslp *OpslpCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccOpsPerShare  *big.Int
}, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccOpsPerShare  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccOpsPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accOpsPerShare)
func (_Opslp *OpslpSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccOpsPerShare  *big.Int
}, error) {
	return _Opslp.Contract.PoolInfo(&_Opslp.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accOpsPerShare)
func (_Opslp *OpslpCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccOpsPerShare  *big.Int
}, error) {
	return _Opslp.Contract.PoolInfo(&_Opslp.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Opslp *OpslpCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Opslp *OpslpSession) PoolLength() (*big.Int, error) {
	return _Opslp.Contract.PoolLength(&_Opslp.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Opslp *OpslpCallerSession) PoolLength() (*big.Int, error) {
	return _Opslp.Contract.PoolLength(&_Opslp.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Opslp *OpslpCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Opslp *OpslpSession) StartBlock() (*big.Int, error) {
	return _Opslp.Contract.StartBlock(&_Opslp.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Opslp *OpslpCallerSession) StartBlock() (*big.Int, error) {
	return _Opslp.Contract.StartBlock(&_Opslp.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Opslp *OpslpCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Opslp *OpslpSession) TotalAllocPoint() (*big.Int, error) {
	return _Opslp.Contract.TotalAllocPoint(&_Opslp.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Opslp *OpslpCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Opslp.Contract.TotalAllocPoint(&_Opslp.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Opslp *OpslpCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Opslp.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Opslp *OpslpSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Opslp.Contract.UserInfo(&_Opslp.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Opslp *OpslpCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Opslp.Contract.UserInfo(&_Opslp.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Opslp *OpslpTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Opslp *OpslpSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Opslp.Contract.Add(&_Opslp.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Opslp *OpslpTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Opslp.Contract.Add(&_Opslp.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Opslp *OpslpTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Opslp *OpslpSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.Deposit(&_Opslp.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Opslp *OpslpTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.Deposit(&_Opslp.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Opslp *OpslpTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Opslp *OpslpSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.EmergencyWithdraw(&_Opslp.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Opslp *OpslpTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.EmergencyWithdraw(&_Opslp.TransactOpts, _pid)
}

// Fee is a paid mutator transaction binding the contract method 0x6fcca69b.
//
// Solidity: function fee(address _feeAddr) returns()
func (_Opslp *OpslpTransactor) Fee(opts *bind.TransactOpts, _feeAddr common.Address) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "fee", _feeAddr)
}

// Fee is a paid mutator transaction binding the contract method 0x6fcca69b.
//
// Solidity: function fee(address _feeAddr) returns()
func (_Opslp *OpslpSession) Fee(_feeAddr common.Address) (*types.Transaction, error) {
	return _Opslp.Contract.Fee(&_Opslp.TransactOpts, _feeAddr)
}

// Fee is a paid mutator transaction binding the contract method 0x6fcca69b.
//
// Solidity: function fee(address _feeAddr) returns()
func (_Opslp *OpslpTransactorSession) Fee(_feeAddr common.Address) (*types.Transaction, error) {
	return _Opslp.Contract.Fee(&_Opslp.TransactOpts, _feeAddr)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Opslp *OpslpTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Opslp *OpslpSession) MassUpdatePools() (*types.Transaction, error) {
	return _Opslp.Contract.MassUpdatePools(&_Opslp.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Opslp *OpslpTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Opslp.Contract.MassUpdatePools(&_Opslp.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Opslp *OpslpTransactor) Migrate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "migrate", _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Opslp *OpslpSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.Migrate(&_Opslp.TransactOpts, _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Opslp *OpslpTransactorSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.Migrate(&_Opslp.TransactOpts, _pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Opslp *OpslpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Opslp *OpslpSession) RenounceOwnership() (*types.Transaction, error) {
	return _Opslp.Contract.RenounceOwnership(&_Opslp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Opslp *OpslpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Opslp.Contract.RenounceOwnership(&_Opslp.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Opslp *OpslpTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Opslp *OpslpSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Opslp.Contract.Set(&_Opslp.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Opslp *OpslpTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Opslp.Contract.Set(&_Opslp.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetFeeRatio is a paid mutator transaction binding the contract method 0x19f4ff2f.
//
// Solidity: function setFeeRatio(uint256 feeRatio_) returns()
func (_Opslp *OpslpTransactor) SetFeeRatio(opts *bind.TransactOpts, feeRatio_ *big.Int) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "setFeeRatio", feeRatio_)
}

// SetFeeRatio is a paid mutator transaction binding the contract method 0x19f4ff2f.
//
// Solidity: function setFeeRatio(uint256 feeRatio_) returns()
func (_Opslp *OpslpSession) SetFeeRatio(feeRatio_ *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.SetFeeRatio(&_Opslp.TransactOpts, feeRatio_)
}

// SetFeeRatio is a paid mutator transaction binding the contract method 0x19f4ff2f.
//
// Solidity: function setFeeRatio(uint256 feeRatio_) returns()
func (_Opslp *OpslpTransactorSession) SetFeeRatio(feeRatio_ *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.SetFeeRatio(&_Opslp.TransactOpts, feeRatio_)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Opslp *OpslpTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Opslp *OpslpSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Opslp.Contract.SetMigrator(&_Opslp.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Opslp *OpslpTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Opslp.Contract.SetMigrator(&_Opslp.TransactOpts, _migrator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Opslp *OpslpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Opslp *OpslpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Opslp.Contract.TransferOwnership(&_Opslp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Opslp *OpslpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Opslp.Contract.TransferOwnership(&_Opslp.TransactOpts, newOwner)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Opslp *OpslpTransactor) UpdateMultiplier(opts *bind.TransactOpts, multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "updateMultiplier", multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Opslp *OpslpSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.UpdateMultiplier(&_Opslp.TransactOpts, multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Opslp *OpslpTransactorSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.UpdateMultiplier(&_Opslp.TransactOpts, multiplierNumber)
}

// UpdateOpsPerBlock is a paid mutator transaction binding the contract method 0x581d93ce.
//
// Solidity: function updateOpsPerBlock(uint256 _opsPerBlock) returns()
func (_Opslp *OpslpTransactor) UpdateOpsPerBlock(opts *bind.TransactOpts, _opsPerBlock *big.Int) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "updateOpsPerBlock", _opsPerBlock)
}

// UpdateOpsPerBlock is a paid mutator transaction binding the contract method 0x581d93ce.
//
// Solidity: function updateOpsPerBlock(uint256 _opsPerBlock) returns()
func (_Opslp *OpslpSession) UpdateOpsPerBlock(_opsPerBlock *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.UpdateOpsPerBlock(&_Opslp.TransactOpts, _opsPerBlock)
}

// UpdateOpsPerBlock is a paid mutator transaction binding the contract method 0x581d93ce.
//
// Solidity: function updateOpsPerBlock(uint256 _opsPerBlock) returns()
func (_Opslp *OpslpTransactorSession) UpdateOpsPerBlock(_opsPerBlock *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.UpdateOpsPerBlock(&_Opslp.TransactOpts, _opsPerBlock)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Opslp *OpslpTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Opslp *OpslpSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.UpdatePool(&_Opslp.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Opslp *OpslpTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.UpdatePool(&_Opslp.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Opslp *OpslpTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Opslp.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Opslp *OpslpSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.Withdraw(&_Opslp.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Opslp *OpslpTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Opslp.Contract.Withdraw(&_Opslp.TransactOpts, _pid, _amount)
}

// OpslpDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Opslp contract.
type OpslpDepositIterator struct {
	Event *OpslpDeposit // Event containing the contract specifics and raw log

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
func (it *OpslpDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpslpDeposit)
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
		it.Event = new(OpslpDeposit)
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
func (it *OpslpDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpslpDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpslpDeposit represents a Deposit event raised by the Opslp contract.
type OpslpDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Opslp *OpslpFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*OpslpDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Opslp.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &OpslpDepositIterator{contract: _Opslp.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Opslp *OpslpFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *OpslpDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Opslp.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpslpDeposit)
				if err := _Opslp.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Opslp *OpslpFilterer) ParseDeposit(log types.Log) (*OpslpDeposit, error) {
	event := new(OpslpDeposit)
	if err := _Opslp.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpslpEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Opslp contract.
type OpslpEmergencyWithdrawIterator struct {
	Event *OpslpEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *OpslpEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpslpEmergencyWithdraw)
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
		it.Event = new(OpslpEmergencyWithdraw)
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
func (it *OpslpEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpslpEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpslpEmergencyWithdraw represents a EmergencyWithdraw event raised by the Opslp contract.
type OpslpEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Opslp *OpslpFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*OpslpEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Opslp.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &OpslpEmergencyWithdrawIterator{contract: _Opslp.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Opslp *OpslpFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *OpslpEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Opslp.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpslpEmergencyWithdraw)
				if err := _Opslp.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Opslp *OpslpFilterer) ParseEmergencyWithdraw(log types.Log) (*OpslpEmergencyWithdraw, error) {
	event := new(OpslpEmergencyWithdraw)
	if err := _Opslp.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpslpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Opslp contract.
type OpslpOwnershipTransferredIterator struct {
	Event *OpslpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OpslpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpslpOwnershipTransferred)
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
		it.Event = new(OpslpOwnershipTransferred)
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
func (it *OpslpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpslpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpslpOwnershipTransferred represents a OwnershipTransferred event raised by the Opslp contract.
type OpslpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Opslp *OpslpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OpslpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Opslp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OpslpOwnershipTransferredIterator{contract: _Opslp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Opslp *OpslpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OpslpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Opslp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpslpOwnershipTransferred)
				if err := _Opslp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Opslp *OpslpFilterer) ParseOwnershipTransferred(log types.Log) (*OpslpOwnershipTransferred, error) {
	event := new(OpslpOwnershipTransferred)
	if err := _Opslp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpslpSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the Opslp contract.
type OpslpSetFeeIterator struct {
	Event *OpslpSetFee // Event containing the contract specifics and raw log

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
func (it *OpslpSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpslpSetFee)
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
		it.Event = new(OpslpSetFee)
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
func (it *OpslpSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpslpSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpslpSetFee represents a SetFee event raised by the Opslp contract.
type OpslpSetFee struct {
	User    common.Address
	FeeAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0xc581d32b46aa9c8bff3b0b4636c086d84ab1a2469ab2fed0210e0426bf9a90a1.
//
// Solidity: event SetFee(address indexed user, address indexed _feeAddr)
func (_Opslp *OpslpFilterer) FilterSetFee(opts *bind.FilterOpts, user []common.Address, _feeAddr []common.Address) (*OpslpSetFeeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var _feeAddrRule []interface{}
	for _, _feeAddrItem := range _feeAddr {
		_feeAddrRule = append(_feeAddrRule, _feeAddrItem)
	}

	logs, sub, err := _Opslp.contract.FilterLogs(opts, "SetFee", userRule, _feeAddrRule)
	if err != nil {
		return nil, err
	}
	return &OpslpSetFeeIterator{contract: _Opslp.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0xc581d32b46aa9c8bff3b0b4636c086d84ab1a2469ab2fed0210e0426bf9a90a1.
//
// Solidity: event SetFee(address indexed user, address indexed _feeAddr)
func (_Opslp *OpslpFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *OpslpSetFee, user []common.Address, _feeAddr []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var _feeAddrRule []interface{}
	for _, _feeAddrItem := range _feeAddr {
		_feeAddrRule = append(_feeAddrRule, _feeAddrItem)
	}

	logs, sub, err := _Opslp.contract.WatchLogs(opts, "SetFee", userRule, _feeAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpslpSetFee)
				if err := _Opslp.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0xc581d32b46aa9c8bff3b0b4636c086d84ab1a2469ab2fed0210e0426bf9a90a1.
//
// Solidity: event SetFee(address indexed user, address indexed _feeAddr)
func (_Opslp *OpslpFilterer) ParseSetFee(log types.Log) (*OpslpSetFee, error) {
	event := new(OpslpSetFee)
	if err := _Opslp.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpslpSetFeeRatioIterator is returned from FilterSetFeeRatio and is used to iterate over the raw logs and unpacked data for SetFeeRatio events raised by the Opslp contract.
type OpslpSetFeeRatioIterator struct {
	Event *OpslpSetFeeRatio // Event containing the contract specifics and raw log

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
func (it *OpslpSetFeeRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpslpSetFeeRatio)
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
		it.Event = new(OpslpSetFeeRatio)
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
func (it *OpslpSetFeeRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpslpSetFeeRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpslpSetFeeRatio represents a SetFeeRatio event raised by the Opslp contract.
type OpslpSetFeeRatio struct {
	Acct   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFeeRatio is a free log retrieval operation binding the contract event 0x2f407797e17e05867e2657ee390ad55cf8dbf1a504056f7b965c645c4821cffd.
//
// Solidity: event SetFeeRatio(address indexed acct, uint256 amount)
func (_Opslp *OpslpFilterer) FilterSetFeeRatio(opts *bind.FilterOpts, acct []common.Address) (*OpslpSetFeeRatioIterator, error) {

	var acctRule []interface{}
	for _, acctItem := range acct {
		acctRule = append(acctRule, acctItem)
	}

	logs, sub, err := _Opslp.contract.FilterLogs(opts, "SetFeeRatio", acctRule)
	if err != nil {
		return nil, err
	}
	return &OpslpSetFeeRatioIterator{contract: _Opslp.contract, event: "SetFeeRatio", logs: logs, sub: sub}, nil
}

// WatchSetFeeRatio is a free log subscription operation binding the contract event 0x2f407797e17e05867e2657ee390ad55cf8dbf1a504056f7b965c645c4821cffd.
//
// Solidity: event SetFeeRatio(address indexed acct, uint256 amount)
func (_Opslp *OpslpFilterer) WatchSetFeeRatio(opts *bind.WatchOpts, sink chan<- *OpslpSetFeeRatio, acct []common.Address) (event.Subscription, error) {

	var acctRule []interface{}
	for _, acctItem := range acct {
		acctRule = append(acctRule, acctItem)
	}

	logs, sub, err := _Opslp.contract.WatchLogs(opts, "SetFeeRatio", acctRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpslpSetFeeRatio)
				if err := _Opslp.contract.UnpackLog(event, "SetFeeRatio", log); err != nil {
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

// ParseSetFeeRatio is a log parse operation binding the contract event 0x2f407797e17e05867e2657ee390ad55cf8dbf1a504056f7b965c645c4821cffd.
//
// Solidity: event SetFeeRatio(address indexed acct, uint256 amount)
func (_Opslp *OpslpFilterer) ParseSetFeeRatio(log types.Log) (*OpslpSetFeeRatio, error) {
	event := new(OpslpSetFeeRatio)
	if err := _Opslp.contract.UnpackLog(event, "SetFeeRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpslpSetMigratorIterator is returned from FilterSetMigrator and is used to iterate over the raw logs and unpacked data for SetMigrator events raised by the Opslp contract.
type OpslpSetMigratorIterator struct {
	Event *OpslpSetMigrator // Event containing the contract specifics and raw log

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
func (it *OpslpSetMigratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpslpSetMigrator)
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
		it.Event = new(OpslpSetMigrator)
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
func (it *OpslpSetMigratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpslpSetMigratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpslpSetMigrator represents a SetMigrator event raised by the Opslp contract.
type OpslpSetMigrator struct {
	Userr    common.Address
	MigrAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMigrator is a free log retrieval operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address indexed userr, address indexed _migrAddr)
func (_Opslp *OpslpFilterer) FilterSetMigrator(opts *bind.FilterOpts, userr []common.Address, _migrAddr []common.Address) (*OpslpSetMigratorIterator, error) {

	var userrRule []interface{}
	for _, userrItem := range userr {
		userrRule = append(userrRule, userrItem)
	}
	var _migrAddrRule []interface{}
	for _, _migrAddrItem := range _migrAddr {
		_migrAddrRule = append(_migrAddrRule, _migrAddrItem)
	}

	logs, sub, err := _Opslp.contract.FilterLogs(opts, "SetMigrator", userrRule, _migrAddrRule)
	if err != nil {
		return nil, err
	}
	return &OpslpSetMigratorIterator{contract: _Opslp.contract, event: "SetMigrator", logs: logs, sub: sub}, nil
}

// WatchSetMigrator is a free log subscription operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address indexed userr, address indexed _migrAddr)
func (_Opslp *OpslpFilterer) WatchSetMigrator(opts *bind.WatchOpts, sink chan<- *OpslpSetMigrator, userr []common.Address, _migrAddr []common.Address) (event.Subscription, error) {

	var userrRule []interface{}
	for _, userrItem := range userr {
		userrRule = append(userrRule, userrItem)
	}
	var _migrAddrRule []interface{}
	for _, _migrAddrItem := range _migrAddr {
		_migrAddrRule = append(_migrAddrRule, _migrAddrItem)
	}

	logs, sub, err := _Opslp.contract.WatchLogs(opts, "SetMigrator", userrRule, _migrAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpslpSetMigrator)
				if err := _Opslp.contract.UnpackLog(event, "SetMigrator", log); err != nil {
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

// ParseSetMigrator is a log parse operation binding the contract event 0xd8ad954fe808212ab9ed7139873e40807dff7995fe36e3d6cdeb8fa00fcebf10.
//
// Solidity: event SetMigrator(address indexed userr, address indexed _migrAddr)
func (_Opslp *OpslpFilterer) ParseSetMigrator(log types.Log) (*OpslpSetMigrator, error) {
	event := new(OpslpSetMigrator)
	if err := _Opslp.contract.UnpackLog(event, "SetMigrator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpslpUpdateMultiplierIterator is returned from FilterUpdateMultiplier and is used to iterate over the raw logs and unpacked data for UpdateMultiplier events raised by the Opslp contract.
type OpslpUpdateMultiplierIterator struct {
	Event *OpslpUpdateMultiplier // Event containing the contract specifics and raw log

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
func (it *OpslpUpdateMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpslpUpdateMultiplier)
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
		it.Event = new(OpslpUpdateMultiplier)
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
func (it *OpslpUpdateMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpslpUpdateMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpslpUpdateMultiplier represents a UpdateMultiplier event raised by the Opslp contract.
type OpslpUpdateMultiplier struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateMultiplier is a free log retrieval operation binding the contract event 0xcf12649cc9200d10a84b242df5627dc75aae1025c6d2a3c73b71206d06a44cdc.
//
// Solidity: event UpdateMultiplier(address indexed user, uint256 amount)
func (_Opslp *OpslpFilterer) FilterUpdateMultiplier(opts *bind.FilterOpts, user []common.Address) (*OpslpUpdateMultiplierIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Opslp.contract.FilterLogs(opts, "UpdateMultiplier", userRule)
	if err != nil {
		return nil, err
	}
	return &OpslpUpdateMultiplierIterator{contract: _Opslp.contract, event: "UpdateMultiplier", logs: logs, sub: sub}, nil
}

// WatchUpdateMultiplier is a free log subscription operation binding the contract event 0xcf12649cc9200d10a84b242df5627dc75aae1025c6d2a3c73b71206d06a44cdc.
//
// Solidity: event UpdateMultiplier(address indexed user, uint256 amount)
func (_Opslp *OpslpFilterer) WatchUpdateMultiplier(opts *bind.WatchOpts, sink chan<- *OpslpUpdateMultiplier, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Opslp.contract.WatchLogs(opts, "UpdateMultiplier", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpslpUpdateMultiplier)
				if err := _Opslp.contract.UnpackLog(event, "UpdateMultiplier", log); err != nil {
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

// ParseUpdateMultiplier is a log parse operation binding the contract event 0xcf12649cc9200d10a84b242df5627dc75aae1025c6d2a3c73b71206d06a44cdc.
//
// Solidity: event UpdateMultiplier(address indexed user, uint256 amount)
func (_Opslp *OpslpFilterer) ParseUpdateMultiplier(log types.Log) (*OpslpUpdateMultiplier, error) {
	event := new(OpslpUpdateMultiplier)
	if err := _Opslp.contract.UnpackLog(event, "UpdateMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpslpUpdateOpsPerBlockIterator is returned from FilterUpdateOpsPerBlock and is used to iterate over the raw logs and unpacked data for UpdateOpsPerBlock events raised by the Opslp contract.
type OpslpUpdateOpsPerBlockIterator struct {
	Event *OpslpUpdateOpsPerBlock // Event containing the contract specifics and raw log

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
func (it *OpslpUpdateOpsPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpslpUpdateOpsPerBlock)
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
		it.Event = new(OpslpUpdateOpsPerBlock)
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
func (it *OpslpUpdateOpsPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpslpUpdateOpsPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpslpUpdateOpsPerBlock represents a UpdateOpsPerBlock event raised by the Opslp contract.
type OpslpUpdateOpsPerBlock struct {
	Acct   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateOpsPerBlock is a free log retrieval operation binding the contract event 0x2cedd8b9af0aa41a2d2388fb9632e7642d1bab90da05227eefc9af80ebe174b7.
//
// Solidity: event UpdateOpsPerBlock(address indexed acct, uint256 amount)
func (_Opslp *OpslpFilterer) FilterUpdateOpsPerBlock(opts *bind.FilterOpts, acct []common.Address) (*OpslpUpdateOpsPerBlockIterator, error) {

	var acctRule []interface{}
	for _, acctItem := range acct {
		acctRule = append(acctRule, acctItem)
	}

	logs, sub, err := _Opslp.contract.FilterLogs(opts, "UpdateOpsPerBlock", acctRule)
	if err != nil {
		return nil, err
	}
	return &OpslpUpdateOpsPerBlockIterator{contract: _Opslp.contract, event: "UpdateOpsPerBlock", logs: logs, sub: sub}, nil
}

// WatchUpdateOpsPerBlock is a free log subscription operation binding the contract event 0x2cedd8b9af0aa41a2d2388fb9632e7642d1bab90da05227eefc9af80ebe174b7.
//
// Solidity: event UpdateOpsPerBlock(address indexed acct, uint256 amount)
func (_Opslp *OpslpFilterer) WatchUpdateOpsPerBlock(opts *bind.WatchOpts, sink chan<- *OpslpUpdateOpsPerBlock, acct []common.Address) (event.Subscription, error) {

	var acctRule []interface{}
	for _, acctItem := range acct {
		acctRule = append(acctRule, acctItem)
	}

	logs, sub, err := _Opslp.contract.WatchLogs(opts, "UpdateOpsPerBlock", acctRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpslpUpdateOpsPerBlock)
				if err := _Opslp.contract.UnpackLog(event, "UpdateOpsPerBlock", log); err != nil {
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

// ParseUpdateOpsPerBlock is a log parse operation binding the contract event 0x2cedd8b9af0aa41a2d2388fb9632e7642d1bab90da05227eefc9af80ebe174b7.
//
// Solidity: event UpdateOpsPerBlock(address indexed acct, uint256 amount)
func (_Opslp *OpslpFilterer) ParseUpdateOpsPerBlock(log types.Log) (*OpslpUpdateOpsPerBlock, error) {
	event := new(OpslpUpdateOpsPerBlock)
	if err := _Opslp.contract.UnpackLog(event, "UpdateOpsPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpslpWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Opslp contract.
type OpslpWithdrawIterator struct {
	Event *OpslpWithdraw // Event containing the contract specifics and raw log

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
func (it *OpslpWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpslpWithdraw)
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
		it.Event = new(OpslpWithdraw)
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
func (it *OpslpWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpslpWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpslpWithdraw represents a Withdraw event raised by the Opslp contract.
type OpslpWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Opslp *OpslpFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*OpslpWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Opslp.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &OpslpWithdrawIterator{contract: _Opslp.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Opslp *OpslpFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *OpslpWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Opslp.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpslpWithdraw)
				if err := _Opslp.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Opslp *OpslpFilterer) ParseWithdraw(log types.Log) (*OpslpWithdraw, error) {
	event := new(OpslpWithdraw)
	if err := _Opslp.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
