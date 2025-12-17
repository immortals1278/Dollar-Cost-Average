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

// TestBaseMetaData contains all meta data concerning the TestBase contract.
var TestBaseMetaData = &bind.MetaData{
	ABI: "[]",
}

// TestBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use TestBaseMetaData.ABI instead.
var TestBaseABI = TestBaseMetaData.ABI

// TestBase is an auto generated Go binding around an Ethereum contract.
type TestBase struct {
	TestBaseCaller     // Read-only binding to the contract
	TestBaseTransactor // Write-only binding to the contract
	TestBaseFilterer   // Log filterer for contract events
}

// TestBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestBaseSession struct {
	Contract     *TestBase         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestBaseCallerSession struct {
	Contract *TestBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TestBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestBaseTransactorSession struct {
	Contract     *TestBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestBaseRaw struct {
	Contract *TestBase // Generic contract binding to access the raw methods on
}

// TestBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestBaseCallerRaw struct {
	Contract *TestBaseCaller // Generic read-only contract binding to access the raw methods on
}

// TestBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestBaseTransactorRaw struct {
	Contract *TestBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestBase creates a new instance of TestBase, bound to a specific deployed contract.
func NewTestBase(address common.Address, backend bind.ContractBackend) (*TestBase, error) {
	contract, err := bindTestBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestBase{TestBaseCaller: TestBaseCaller{contract: contract}, TestBaseTransactor: TestBaseTransactor{contract: contract}, TestBaseFilterer: TestBaseFilterer{contract: contract}}, nil
}

// NewTestBaseCaller creates a new read-only instance of TestBase, bound to a specific deployed contract.
func NewTestBaseCaller(address common.Address, caller bind.ContractCaller) (*TestBaseCaller, error) {
	contract, err := bindTestBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestBaseCaller{contract: contract}, nil
}

// NewTestBaseTransactor creates a new write-only instance of TestBase, bound to a specific deployed contract.
func NewTestBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*TestBaseTransactor, error) {
	contract, err := bindTestBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestBaseTransactor{contract: contract}, nil
}

// NewTestBaseFilterer creates a new log filterer instance of TestBase, bound to a specific deployed contract.
func NewTestBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*TestBaseFilterer, error) {
	contract, err := bindTestBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestBaseFilterer{contract: contract}, nil
}

// bindTestBase binds a generic wrapper to an already deployed contract.
func bindTestBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestBase *TestBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestBase.Contract.TestBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestBase *TestBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestBase.Contract.TestBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestBase *TestBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestBase.Contract.TestBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestBase *TestBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestBase *TestBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestBase *TestBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestBase.Contract.contract.Transact(opts, method, params...)
}
