// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iupgradeableproxy

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

// IUpgradeableProxyMetaData contains all meta data concerning the IUpgradeableProxy contract.
var IUpgradeableProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"}]",
}

// IUpgradeableProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use IUpgradeableProxyMetaData.ABI instead.
var IUpgradeableProxyABI = IUpgradeableProxyMetaData.ABI

// IUpgradeableProxy is an auto generated Go binding around an Ethereum contract.
type IUpgradeableProxy struct {
	IUpgradeableProxyCaller     // Read-only binding to the contract
	IUpgradeableProxyTransactor // Write-only binding to the contract
	IUpgradeableProxyFilterer   // Log filterer for contract events
}

// IUpgradeableProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUpgradeableProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradeableProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUpgradeableProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradeableProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUpgradeableProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradeableProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUpgradeableProxySession struct {
	Contract     *IUpgradeableProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IUpgradeableProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUpgradeableProxyCallerSession struct {
	Contract *IUpgradeableProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IUpgradeableProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUpgradeableProxyTransactorSession struct {
	Contract     *IUpgradeableProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IUpgradeableProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUpgradeableProxyRaw struct {
	Contract *IUpgradeableProxy // Generic contract binding to access the raw methods on
}

// IUpgradeableProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUpgradeableProxyCallerRaw struct {
	Contract *IUpgradeableProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IUpgradeableProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUpgradeableProxyTransactorRaw struct {
	Contract *IUpgradeableProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUpgradeableProxy creates a new instance of IUpgradeableProxy, bound to a specific deployed contract.
func NewIUpgradeableProxy(address common.Address, backend bind.ContractBackend) (*IUpgradeableProxy, error) {
	contract, err := bindIUpgradeableProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUpgradeableProxy{IUpgradeableProxyCaller: IUpgradeableProxyCaller{contract: contract}, IUpgradeableProxyTransactor: IUpgradeableProxyTransactor{contract: contract}, IUpgradeableProxyFilterer: IUpgradeableProxyFilterer{contract: contract}}, nil
}

// NewIUpgradeableProxyCaller creates a new read-only instance of IUpgradeableProxy, bound to a specific deployed contract.
func NewIUpgradeableProxyCaller(address common.Address, caller bind.ContractCaller) (*IUpgradeableProxyCaller, error) {
	contract, err := bindIUpgradeableProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradeableProxyCaller{contract: contract}, nil
}

// NewIUpgradeableProxyTransactor creates a new write-only instance of IUpgradeableProxy, bound to a specific deployed contract.
func NewIUpgradeableProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IUpgradeableProxyTransactor, error) {
	contract, err := bindIUpgradeableProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradeableProxyTransactor{contract: contract}, nil
}

// NewIUpgradeableProxyFilterer creates a new log filterer instance of IUpgradeableProxy, bound to a specific deployed contract.
func NewIUpgradeableProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IUpgradeableProxyFilterer, error) {
	contract, err := bindIUpgradeableProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUpgradeableProxyFilterer{contract: contract}, nil
}

// bindIUpgradeableProxy binds a generic wrapper to an already deployed contract.
func bindIUpgradeableProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradeableProxy *IUpgradeableProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUpgradeableProxy.Contract.IUpgradeableProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradeableProxy *IUpgradeableProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradeableProxy.Contract.IUpgradeableProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradeableProxy *IUpgradeableProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradeableProxy.Contract.IUpgradeableProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradeableProxy *IUpgradeableProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUpgradeableProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradeableProxy *IUpgradeableProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradeableProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradeableProxy *IUpgradeableProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradeableProxy.Contract.contract.Transact(opts, method, params...)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_IUpgradeableProxy *IUpgradeableProxyTransactor) UpgradeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradeableProxy.contract.Transact(opts, "upgradeTo", arg0)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_IUpgradeableProxy *IUpgradeableProxySession) UpgradeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradeableProxy.Contract.UpgradeTo(&_IUpgradeableProxy.TransactOpts, arg0)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_IUpgradeableProxy *IUpgradeableProxyTransactorSession) UpgradeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradeableProxy.Contract.UpgradeTo(&_IUpgradeableProxy.TransactOpts, arg0)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address , bytes ) payable returns()
func (_IUpgradeableProxy *IUpgradeableProxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, arg0 common.Address, arg1 []byte) (*types.Transaction, error) {
	return _IUpgradeableProxy.contract.Transact(opts, "upgradeToAndCall", arg0, arg1)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address , bytes ) payable returns()
func (_IUpgradeableProxy *IUpgradeableProxySession) UpgradeToAndCall(arg0 common.Address, arg1 []byte) (*types.Transaction, error) {
	return _IUpgradeableProxy.Contract.UpgradeToAndCall(&_IUpgradeableProxy.TransactOpts, arg0, arg1)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address , bytes ) payable returns()
func (_IUpgradeableProxy *IUpgradeableProxyTransactorSession) UpgradeToAndCall(arg0 common.Address, arg1 []byte) (*types.Transaction, error) {
	return _IUpgradeableProxy.Contract.UpgradeToAndCall(&_IUpgradeableProxy.TransactOpts, arg0, arg1)
}
