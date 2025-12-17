// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stdcheats

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

// StdCheatsSafeMetaData contains all meta data concerning the StdCheatsSafe contract.
var StdCheatsSafeMetaData = &bind.MetaData{
	ABI: "[]",
}

// StdCheatsSafeABI is the input ABI used to generate the binding from.
// Deprecated: Use StdCheatsSafeMetaData.ABI instead.
var StdCheatsSafeABI = StdCheatsSafeMetaData.ABI

// StdCheatsSafe is an auto generated Go binding around an Ethereum contract.
type StdCheatsSafe struct {
	StdCheatsSafeCaller     // Read-only binding to the contract
	StdCheatsSafeTransactor // Write-only binding to the contract
	StdCheatsSafeFilterer   // Log filterer for contract events
}

// StdCheatsSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdCheatsSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdCheatsSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdCheatsSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdCheatsSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdCheatsSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdCheatsSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdCheatsSafeSession struct {
	Contract     *StdCheatsSafe    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdCheatsSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdCheatsSafeCallerSession struct {
	Contract *StdCheatsSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StdCheatsSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdCheatsSafeTransactorSession struct {
	Contract     *StdCheatsSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StdCheatsSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdCheatsSafeRaw struct {
	Contract *StdCheatsSafe // Generic contract binding to access the raw methods on
}

// StdCheatsSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdCheatsSafeCallerRaw struct {
	Contract *StdCheatsSafeCaller // Generic read-only contract binding to access the raw methods on
}

// StdCheatsSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdCheatsSafeTransactorRaw struct {
	Contract *StdCheatsSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdCheatsSafe creates a new instance of StdCheatsSafe, bound to a specific deployed contract.
func NewStdCheatsSafe(address common.Address, backend bind.ContractBackend) (*StdCheatsSafe, error) {
	contract, err := bindStdCheatsSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdCheatsSafe{StdCheatsSafeCaller: StdCheatsSafeCaller{contract: contract}, StdCheatsSafeTransactor: StdCheatsSafeTransactor{contract: contract}, StdCheatsSafeFilterer: StdCheatsSafeFilterer{contract: contract}}, nil
}

// NewStdCheatsSafeCaller creates a new read-only instance of StdCheatsSafe, bound to a specific deployed contract.
func NewStdCheatsSafeCaller(address common.Address, caller bind.ContractCaller) (*StdCheatsSafeCaller, error) {
	contract, err := bindStdCheatsSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdCheatsSafeCaller{contract: contract}, nil
}

// NewStdCheatsSafeTransactor creates a new write-only instance of StdCheatsSafe, bound to a specific deployed contract.
func NewStdCheatsSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*StdCheatsSafeTransactor, error) {
	contract, err := bindStdCheatsSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdCheatsSafeTransactor{contract: contract}, nil
}

// NewStdCheatsSafeFilterer creates a new log filterer instance of StdCheatsSafe, bound to a specific deployed contract.
func NewStdCheatsSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*StdCheatsSafeFilterer, error) {
	contract, err := bindStdCheatsSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdCheatsSafeFilterer{contract: contract}, nil
}

// bindStdCheatsSafe binds a generic wrapper to an already deployed contract.
func bindStdCheatsSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdCheatsSafeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdCheatsSafe *StdCheatsSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdCheatsSafe.Contract.StdCheatsSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdCheatsSafe *StdCheatsSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdCheatsSafe.Contract.StdCheatsSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdCheatsSafe *StdCheatsSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdCheatsSafe.Contract.StdCheatsSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdCheatsSafe *StdCheatsSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdCheatsSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdCheatsSafe *StdCheatsSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdCheatsSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdCheatsSafe *StdCheatsSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdCheatsSafe.Contract.contract.Transact(opts, method, params...)
}
