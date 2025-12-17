// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iupgradeablebeacon

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

// IUpgradeableBeaconMetaData contains all meta data concerning the IUpgradeableBeacon contract.
var IUpgradeableBeaconMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IUpgradeableBeaconABI is the input ABI used to generate the binding from.
// Deprecated: Use IUpgradeableBeaconMetaData.ABI instead.
var IUpgradeableBeaconABI = IUpgradeableBeaconMetaData.ABI

// IUpgradeableBeacon is an auto generated Go binding around an Ethereum contract.
type IUpgradeableBeacon struct {
	IUpgradeableBeaconCaller     // Read-only binding to the contract
	IUpgradeableBeaconTransactor // Write-only binding to the contract
	IUpgradeableBeaconFilterer   // Log filterer for contract events
}

// IUpgradeableBeaconCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUpgradeableBeaconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradeableBeaconTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUpgradeableBeaconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradeableBeaconFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUpgradeableBeaconFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradeableBeaconSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUpgradeableBeaconSession struct {
	Contract     *IUpgradeableBeacon // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IUpgradeableBeaconCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUpgradeableBeaconCallerSession struct {
	Contract *IUpgradeableBeaconCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IUpgradeableBeaconTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUpgradeableBeaconTransactorSession struct {
	Contract     *IUpgradeableBeaconTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IUpgradeableBeaconRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUpgradeableBeaconRaw struct {
	Contract *IUpgradeableBeacon // Generic contract binding to access the raw methods on
}

// IUpgradeableBeaconCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUpgradeableBeaconCallerRaw struct {
	Contract *IUpgradeableBeaconCaller // Generic read-only contract binding to access the raw methods on
}

// IUpgradeableBeaconTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUpgradeableBeaconTransactorRaw struct {
	Contract *IUpgradeableBeaconTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUpgradeableBeacon creates a new instance of IUpgradeableBeacon, bound to a specific deployed contract.
func NewIUpgradeableBeacon(address common.Address, backend bind.ContractBackend) (*IUpgradeableBeacon, error) {
	contract, err := bindIUpgradeableBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUpgradeableBeacon{IUpgradeableBeaconCaller: IUpgradeableBeaconCaller{contract: contract}, IUpgradeableBeaconTransactor: IUpgradeableBeaconTransactor{contract: contract}, IUpgradeableBeaconFilterer: IUpgradeableBeaconFilterer{contract: contract}}, nil
}

// NewIUpgradeableBeaconCaller creates a new read-only instance of IUpgradeableBeacon, bound to a specific deployed contract.
func NewIUpgradeableBeaconCaller(address common.Address, caller bind.ContractCaller) (*IUpgradeableBeaconCaller, error) {
	contract, err := bindIUpgradeableBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradeableBeaconCaller{contract: contract}, nil
}

// NewIUpgradeableBeaconTransactor creates a new write-only instance of IUpgradeableBeacon, bound to a specific deployed contract.
func NewIUpgradeableBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*IUpgradeableBeaconTransactor, error) {
	contract, err := bindIUpgradeableBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradeableBeaconTransactor{contract: contract}, nil
}

// NewIUpgradeableBeaconFilterer creates a new log filterer instance of IUpgradeableBeacon, bound to a specific deployed contract.
func NewIUpgradeableBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*IUpgradeableBeaconFilterer, error) {
	contract, err := bindIUpgradeableBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUpgradeableBeaconFilterer{contract: contract}, nil
}

// bindIUpgradeableBeacon binds a generic wrapper to an already deployed contract.
func bindIUpgradeableBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUpgradeableBeaconMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradeableBeacon *IUpgradeableBeaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUpgradeableBeacon.Contract.IUpgradeableBeaconCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradeableBeacon *IUpgradeableBeaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradeableBeacon.Contract.IUpgradeableBeaconTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradeableBeacon *IUpgradeableBeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradeableBeacon.Contract.IUpgradeableBeaconTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradeableBeacon *IUpgradeableBeaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUpgradeableBeacon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradeableBeacon *IUpgradeableBeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradeableBeacon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradeableBeacon *IUpgradeableBeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradeableBeacon.Contract.contract.Transact(opts, method, params...)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_IUpgradeableBeacon *IUpgradeableBeaconTransactor) UpgradeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradeableBeacon.contract.Transact(opts, "upgradeTo", arg0)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_IUpgradeableBeacon *IUpgradeableBeaconSession) UpgradeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradeableBeacon.Contract.UpgradeTo(&_IUpgradeableBeacon.TransactOpts, arg0)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address ) returns()
func (_IUpgradeableBeacon *IUpgradeableBeaconTransactorSession) UpgradeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradeableBeacon.Contract.UpgradeTo(&_IUpgradeableBeacon.TransactOpts, arg0)
}
