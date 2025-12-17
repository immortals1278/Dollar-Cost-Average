// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package base

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

// ScriptBaseMetaData contains all meta data concerning the ScriptBase contract.
var ScriptBaseMetaData = &bind.MetaData{
	ABI: "[]",
}

// ScriptBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use ScriptBaseMetaData.ABI instead.
var ScriptBaseABI = ScriptBaseMetaData.ABI

// ScriptBase is an auto generated Go binding around an Ethereum contract.
type ScriptBase struct {
	ScriptBaseCaller     // Read-only binding to the contract
	ScriptBaseTransactor // Write-only binding to the contract
	ScriptBaseFilterer   // Log filterer for contract events
}

// ScriptBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScriptBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScriptBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScriptBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScriptBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScriptBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScriptBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScriptBaseSession struct {
	Contract     *ScriptBase       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScriptBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScriptBaseCallerSession struct {
	Contract *ScriptBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ScriptBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScriptBaseTransactorSession struct {
	Contract     *ScriptBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ScriptBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScriptBaseRaw struct {
	Contract *ScriptBase // Generic contract binding to access the raw methods on
}

// ScriptBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScriptBaseCallerRaw struct {
	Contract *ScriptBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ScriptBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScriptBaseTransactorRaw struct {
	Contract *ScriptBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScriptBase creates a new instance of ScriptBase, bound to a specific deployed contract.
func NewScriptBase(address common.Address, backend bind.ContractBackend) (*ScriptBase, error) {
	contract, err := bindScriptBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScriptBase{ScriptBaseCaller: ScriptBaseCaller{contract: contract}, ScriptBaseTransactor: ScriptBaseTransactor{contract: contract}, ScriptBaseFilterer: ScriptBaseFilterer{contract: contract}}, nil
}

// NewScriptBaseCaller creates a new read-only instance of ScriptBase, bound to a specific deployed contract.
func NewScriptBaseCaller(address common.Address, caller bind.ContractCaller) (*ScriptBaseCaller, error) {
	contract, err := bindScriptBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScriptBaseCaller{contract: contract}, nil
}

// NewScriptBaseTransactor creates a new write-only instance of ScriptBase, bound to a specific deployed contract.
func NewScriptBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ScriptBaseTransactor, error) {
	contract, err := bindScriptBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScriptBaseTransactor{contract: contract}, nil
}

// NewScriptBaseFilterer creates a new log filterer instance of ScriptBase, bound to a specific deployed contract.
func NewScriptBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ScriptBaseFilterer, error) {
	contract, err := bindScriptBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScriptBaseFilterer{contract: contract}, nil
}

// bindScriptBase binds a generic wrapper to an already deployed contract.
func bindScriptBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScriptBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScriptBase *ScriptBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScriptBase.Contract.ScriptBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScriptBase *ScriptBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScriptBase.Contract.ScriptBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScriptBase *ScriptBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScriptBase.Contract.ScriptBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScriptBase *ScriptBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScriptBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScriptBase *ScriptBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScriptBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScriptBase *ScriptBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScriptBase.Contract.contract.Transact(opts, method, params...)
}
