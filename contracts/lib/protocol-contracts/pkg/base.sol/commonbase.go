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

// CommonBaseMetaData contains all meta data concerning the CommonBase contract.
var CommonBaseMetaData = &bind.MetaData{
	ABI: "[]",
}

// CommonBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use CommonBaseMetaData.ABI instead.
var CommonBaseABI = CommonBaseMetaData.ABI

// CommonBase is an auto generated Go binding around an Ethereum contract.
type CommonBase struct {
	CommonBaseCaller     // Read-only binding to the contract
	CommonBaseTransactor // Write-only binding to the contract
	CommonBaseFilterer   // Log filterer for contract events
}

// CommonBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommonBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommonBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommonBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommonBaseSession struct {
	Contract     *CommonBase       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommonBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommonBaseCallerSession struct {
	Contract *CommonBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CommonBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommonBaseTransactorSession struct {
	Contract     *CommonBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CommonBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommonBaseRaw struct {
	Contract *CommonBase // Generic contract binding to access the raw methods on
}

// CommonBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommonBaseCallerRaw struct {
	Contract *CommonBaseCaller // Generic read-only contract binding to access the raw methods on
}

// CommonBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommonBaseTransactorRaw struct {
	Contract *CommonBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommonBase creates a new instance of CommonBase, bound to a specific deployed contract.
func NewCommonBase(address common.Address, backend bind.ContractBackend) (*CommonBase, error) {
	contract, err := bindCommonBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommonBase{CommonBaseCaller: CommonBaseCaller{contract: contract}, CommonBaseTransactor: CommonBaseTransactor{contract: contract}, CommonBaseFilterer: CommonBaseFilterer{contract: contract}}, nil
}

// NewCommonBaseCaller creates a new read-only instance of CommonBase, bound to a specific deployed contract.
func NewCommonBaseCaller(address common.Address, caller bind.ContractCaller) (*CommonBaseCaller, error) {
	contract, err := bindCommonBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommonBaseCaller{contract: contract}, nil
}

// NewCommonBaseTransactor creates a new write-only instance of CommonBase, bound to a specific deployed contract.
func NewCommonBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*CommonBaseTransactor, error) {
	contract, err := bindCommonBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommonBaseTransactor{contract: contract}, nil
}

// NewCommonBaseFilterer creates a new log filterer instance of CommonBase, bound to a specific deployed contract.
func NewCommonBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*CommonBaseFilterer, error) {
	contract, err := bindCommonBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommonBaseFilterer{contract: contract}, nil
}

// bindCommonBase binds a generic wrapper to an already deployed contract.
func bindCommonBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommonBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommonBase *CommonBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommonBase.Contract.CommonBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommonBase *CommonBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommonBase.Contract.CommonBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommonBase *CommonBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommonBase.Contract.CommonBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommonBase *CommonBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommonBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommonBase *CommonBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommonBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommonBase *CommonBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommonBase.Contract.contract.Transact(opts, method, params...)
}
