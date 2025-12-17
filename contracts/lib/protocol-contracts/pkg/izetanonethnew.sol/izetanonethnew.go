// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izetanonethnew

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

// IZetaNonEthNewMetaData contains all meta data concerning the IZetaNonEthNew contract.
var IZetaNonEthNewMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"mintee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IZetaNonEthNewABI is the input ABI used to generate the binding from.
// Deprecated: Use IZetaNonEthNewMetaData.ABI instead.
var IZetaNonEthNewABI = IZetaNonEthNewMetaData.ABI

// IZetaNonEthNew is an auto generated Go binding around an Ethereum contract.
type IZetaNonEthNew struct {
	IZetaNonEthNewCaller     // Read-only binding to the contract
	IZetaNonEthNewTransactor // Write-only binding to the contract
	IZetaNonEthNewFilterer   // Log filterer for contract events
}

// IZetaNonEthNewCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZetaNonEthNewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaNonEthNewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZetaNonEthNewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaNonEthNewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZetaNonEthNewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaNonEthNewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZetaNonEthNewSession struct {
	Contract     *IZetaNonEthNew   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZetaNonEthNewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZetaNonEthNewCallerSession struct {
	Contract *IZetaNonEthNewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IZetaNonEthNewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZetaNonEthNewTransactorSession struct {
	Contract     *IZetaNonEthNewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IZetaNonEthNewRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZetaNonEthNewRaw struct {
	Contract *IZetaNonEthNew // Generic contract binding to access the raw methods on
}

// IZetaNonEthNewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZetaNonEthNewCallerRaw struct {
	Contract *IZetaNonEthNewCaller // Generic read-only contract binding to access the raw methods on
}

// IZetaNonEthNewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZetaNonEthNewTransactorRaw struct {
	Contract *IZetaNonEthNewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZetaNonEthNew creates a new instance of IZetaNonEthNew, bound to a specific deployed contract.
func NewIZetaNonEthNew(address common.Address, backend bind.ContractBackend) (*IZetaNonEthNew, error) {
	contract, err := bindIZetaNonEthNew(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthNew{IZetaNonEthNewCaller: IZetaNonEthNewCaller{contract: contract}, IZetaNonEthNewTransactor: IZetaNonEthNewTransactor{contract: contract}, IZetaNonEthNewFilterer: IZetaNonEthNewFilterer{contract: contract}}, nil
}

// NewIZetaNonEthNewCaller creates a new read-only instance of IZetaNonEthNew, bound to a specific deployed contract.
func NewIZetaNonEthNewCaller(address common.Address, caller bind.ContractCaller) (*IZetaNonEthNewCaller, error) {
	contract, err := bindIZetaNonEthNew(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthNewCaller{contract: contract}, nil
}

// NewIZetaNonEthNewTransactor creates a new write-only instance of IZetaNonEthNew, bound to a specific deployed contract.
func NewIZetaNonEthNewTransactor(address common.Address, transactor bind.ContractTransactor) (*IZetaNonEthNewTransactor, error) {
	contract, err := bindIZetaNonEthNew(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthNewTransactor{contract: contract}, nil
}

// NewIZetaNonEthNewFilterer creates a new log filterer instance of IZetaNonEthNew, bound to a specific deployed contract.
func NewIZetaNonEthNewFilterer(address common.Address, filterer bind.ContractFilterer) (*IZetaNonEthNewFilterer, error) {
	contract, err := bindIZetaNonEthNew(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthNewFilterer{contract: contract}, nil
}

// bindIZetaNonEthNew binds a generic wrapper to an already deployed contract.
func bindIZetaNonEthNew(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZetaNonEthNewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaNonEthNew *IZetaNonEthNewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaNonEthNew.Contract.IZetaNonEthNewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaNonEthNew *IZetaNonEthNewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.IZetaNonEthNewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaNonEthNew *IZetaNonEthNewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.IZetaNonEthNewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaNonEthNew *IZetaNonEthNewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaNonEthNew.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaNonEthNew *IZetaNonEthNewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaNonEthNew *IZetaNonEthNewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZetaNonEthNew *IZetaNonEthNewCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZetaNonEthNew.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZetaNonEthNew *IZetaNonEthNewSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZetaNonEthNew.Contract.Allowance(&_IZetaNonEthNew.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZetaNonEthNew *IZetaNonEthNewCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZetaNonEthNew.Contract.Allowance(&_IZetaNonEthNew.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZetaNonEthNew *IZetaNonEthNewCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZetaNonEthNew.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZetaNonEthNew *IZetaNonEthNewSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZetaNonEthNew.Contract.BalanceOf(&_IZetaNonEthNew.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZetaNonEthNew *IZetaNonEthNewCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZetaNonEthNew.Contract.BalanceOf(&_IZetaNonEthNew.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZetaNonEthNew *IZetaNonEthNewCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZetaNonEthNew.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZetaNonEthNew *IZetaNonEthNewSession) TotalSupply() (*big.Int, error) {
	return _IZetaNonEthNew.Contract.TotalSupply(&_IZetaNonEthNew.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZetaNonEthNew *IZetaNonEthNewCallerSession) TotalSupply() (*big.Int, error) {
	return _IZetaNonEthNew.Contract.TotalSupply(&_IZetaNonEthNew.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IZetaNonEthNew *IZetaNonEthNewTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IZetaNonEthNew *IZetaNonEthNewSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.Approve(&_IZetaNonEthNew.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IZetaNonEthNew *IZetaNonEthNewTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.Approve(&_IZetaNonEthNew.TransactOpts, spender, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IZetaNonEthNew *IZetaNonEthNewTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IZetaNonEthNew *IZetaNonEthNewSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.BurnFrom(&_IZetaNonEthNew.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IZetaNonEthNew *IZetaNonEthNewTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.BurnFrom(&_IZetaNonEthNew.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IZetaNonEthNew *IZetaNonEthNewTransactor) Mint(opts *bind.TransactOpts, mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IZetaNonEthNew.contract.Transact(opts, "mint", mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IZetaNonEthNew *IZetaNonEthNewSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.Mint(&_IZetaNonEthNew.TransactOpts, mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IZetaNonEthNew *IZetaNonEthNewTransactorSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.Mint(&_IZetaNonEthNew.TransactOpts, mintee, value, internalSendHash)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IZetaNonEthNew *IZetaNonEthNewTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IZetaNonEthNew *IZetaNonEthNewSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.Transfer(&_IZetaNonEthNew.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IZetaNonEthNew *IZetaNonEthNewTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.Transfer(&_IZetaNonEthNew.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IZetaNonEthNew *IZetaNonEthNewTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IZetaNonEthNew *IZetaNonEthNewSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.TransferFrom(&_IZetaNonEthNew.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IZetaNonEthNew *IZetaNonEthNewTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthNew.Contract.TransferFrom(&_IZetaNonEthNew.TransactOpts, from, to, value)
}

// IZetaNonEthNewApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IZetaNonEthNew contract.
type IZetaNonEthNewApprovalIterator struct {
	Event *IZetaNonEthNewApproval // Event containing the contract specifics and raw log

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
func (it *IZetaNonEthNewApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaNonEthNewApproval)
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
		it.Event = new(IZetaNonEthNewApproval)
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
func (it *IZetaNonEthNewApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaNonEthNewApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaNonEthNewApproval represents a Approval event raised by the IZetaNonEthNew contract.
type IZetaNonEthNewApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IZetaNonEthNew *IZetaNonEthNewFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IZetaNonEthNewApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IZetaNonEthNew.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthNewApprovalIterator{contract: _IZetaNonEthNew.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IZetaNonEthNew *IZetaNonEthNewFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IZetaNonEthNewApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IZetaNonEthNew.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaNonEthNewApproval)
				if err := _IZetaNonEthNew.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IZetaNonEthNew *IZetaNonEthNewFilterer) ParseApproval(log types.Log) (*IZetaNonEthNewApproval, error) {
	event := new(IZetaNonEthNewApproval)
	if err := _IZetaNonEthNew.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaNonEthNewTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IZetaNonEthNew contract.
type IZetaNonEthNewTransferIterator struct {
	Event *IZetaNonEthNewTransfer // Event containing the contract specifics and raw log

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
func (it *IZetaNonEthNewTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaNonEthNewTransfer)
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
		it.Event = new(IZetaNonEthNewTransfer)
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
func (it *IZetaNonEthNewTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaNonEthNewTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaNonEthNewTransfer represents a Transfer event raised by the IZetaNonEthNew contract.
type IZetaNonEthNewTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IZetaNonEthNew *IZetaNonEthNewFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IZetaNonEthNewTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaNonEthNew.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthNewTransferIterator{contract: _IZetaNonEthNew.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IZetaNonEthNew *IZetaNonEthNewFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IZetaNonEthNewTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaNonEthNew.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaNonEthNewTransfer)
				if err := _IZetaNonEthNew.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IZetaNonEthNew *IZetaNonEthNewFilterer) ParseTransfer(log types.Log) (*IZetaNonEthNewTransfer, error) {
	event := new(IZetaNonEthNewTransfer)
	if err := _IZetaNonEthNew.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
