// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package revert

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

// AbortContext is an auto generated low-level Go binding around an user-defined struct.
type AbortContext struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}

// AbortableMetaData contains all meta data concerning the Abortable contract.
var AbortableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onAbort\",\"inputs\":[{\"name\":\"abortContext\",\"type\":\"tuple\",\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// AbortableABI is the input ABI used to generate the binding from.
// Deprecated: Use AbortableMetaData.ABI instead.
var AbortableABI = AbortableMetaData.ABI

// Abortable is an auto generated Go binding around an Ethereum contract.
type Abortable struct {
	AbortableCaller     // Read-only binding to the contract
	AbortableTransactor // Write-only binding to the contract
	AbortableFilterer   // Log filterer for contract events
}

// AbortableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbortableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbortableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbortableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbortableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbortableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbortableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbortableSession struct {
	Contract     *Abortable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbortableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbortableCallerSession struct {
	Contract *AbortableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AbortableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbortableTransactorSession struct {
	Contract     *AbortableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AbortableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbortableRaw struct {
	Contract *Abortable // Generic contract binding to access the raw methods on
}

// AbortableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbortableCallerRaw struct {
	Contract *AbortableCaller // Generic read-only contract binding to access the raw methods on
}

// AbortableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbortableTransactorRaw struct {
	Contract *AbortableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbortable creates a new instance of Abortable, bound to a specific deployed contract.
func NewAbortable(address common.Address, backend bind.ContractBackend) (*Abortable, error) {
	contract, err := bindAbortable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abortable{AbortableCaller: AbortableCaller{contract: contract}, AbortableTransactor: AbortableTransactor{contract: contract}, AbortableFilterer: AbortableFilterer{contract: contract}}, nil
}

// NewAbortableCaller creates a new read-only instance of Abortable, bound to a specific deployed contract.
func NewAbortableCaller(address common.Address, caller bind.ContractCaller) (*AbortableCaller, error) {
	contract, err := bindAbortable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbortableCaller{contract: contract}, nil
}

// NewAbortableTransactor creates a new write-only instance of Abortable, bound to a specific deployed contract.
func NewAbortableTransactor(address common.Address, transactor bind.ContractTransactor) (*AbortableTransactor, error) {
	contract, err := bindAbortable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbortableTransactor{contract: contract}, nil
}

// NewAbortableFilterer creates a new log filterer instance of Abortable, bound to a specific deployed contract.
func NewAbortableFilterer(address common.Address, filterer bind.ContractFilterer) (*AbortableFilterer, error) {
	contract, err := bindAbortable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbortableFilterer{contract: contract}, nil
}

// bindAbortable binds a generic wrapper to an already deployed contract.
func bindAbortable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbortableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abortable *AbortableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abortable.Contract.AbortableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abortable *AbortableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abortable.Contract.AbortableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abortable *AbortableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abortable.Contract.AbortableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abortable *AbortableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abortable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abortable *AbortableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abortable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abortable *AbortableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abortable.Contract.contract.Transact(opts, method, params...)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_Abortable *AbortableTransactor) OnAbort(opts *bind.TransactOpts, abortContext AbortContext) (*types.Transaction, error) {
	return _Abortable.contract.Transact(opts, "onAbort", abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_Abortable *AbortableSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _Abortable.Contract.OnAbort(&_Abortable.TransactOpts, abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_Abortable *AbortableTransactorSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _Abortable.Contract.OnAbort(&_Abortable.TransactOpts, abortContext)
}
