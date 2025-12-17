// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20custodyechidnatest

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

// ERC20CustodyEchidnaTestMetaData contains all meta data concerning the ERC20CustodyEchidnaTest contract.
var ERC20CustodyEchidnaTestMetaData = &bind.MetaData{
	ABI: "[]",
}

// ERC20CustodyEchidnaTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20CustodyEchidnaTestMetaData.ABI instead.
var ERC20CustodyEchidnaTestABI = ERC20CustodyEchidnaTestMetaData.ABI

// ERC20CustodyEchidnaTest is an auto generated Go binding around an Ethereum contract.
type ERC20CustodyEchidnaTest struct {
	ERC20CustodyEchidnaTestCaller     // Read-only binding to the contract
	ERC20CustodyEchidnaTestTransactor // Write-only binding to the contract
	ERC20CustodyEchidnaTestFilterer   // Log filterer for contract events
}

// ERC20CustodyEchidnaTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20CustodyEchidnaTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyEchidnaTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20CustodyEchidnaTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyEchidnaTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20CustodyEchidnaTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyEchidnaTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20CustodyEchidnaTestSession struct {
	Contract     *ERC20CustodyEchidnaTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20CustodyEchidnaTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CustodyEchidnaTestCallerSession struct {
	Contract *ERC20CustodyEchidnaTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ERC20CustodyEchidnaTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20CustodyEchidnaTestTransactorSession struct {
	Contract     *ERC20CustodyEchidnaTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ERC20CustodyEchidnaTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20CustodyEchidnaTestRaw struct {
	Contract *ERC20CustodyEchidnaTest // Generic contract binding to access the raw methods on
}

// ERC20CustodyEchidnaTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CustodyEchidnaTestCallerRaw struct {
	Contract *ERC20CustodyEchidnaTestCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20CustodyEchidnaTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20CustodyEchidnaTestTransactorRaw struct {
	Contract *ERC20CustodyEchidnaTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20CustodyEchidnaTest creates a new instance of ERC20CustodyEchidnaTest, bound to a specific deployed contract.
func NewERC20CustodyEchidnaTest(address common.Address, backend bind.ContractBackend) (*ERC20CustodyEchidnaTest, error) {
	contract, err := bindERC20CustodyEchidnaTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyEchidnaTest{ERC20CustodyEchidnaTestCaller: ERC20CustodyEchidnaTestCaller{contract: contract}, ERC20CustodyEchidnaTestTransactor: ERC20CustodyEchidnaTestTransactor{contract: contract}, ERC20CustodyEchidnaTestFilterer: ERC20CustodyEchidnaTestFilterer{contract: contract}}, nil
}

// NewERC20CustodyEchidnaTestCaller creates a new read-only instance of ERC20CustodyEchidnaTest, bound to a specific deployed contract.
func NewERC20CustodyEchidnaTestCaller(address common.Address, caller bind.ContractCaller) (*ERC20CustodyEchidnaTestCaller, error) {
	contract, err := bindERC20CustodyEchidnaTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyEchidnaTestCaller{contract: contract}, nil
}

// NewERC20CustodyEchidnaTestTransactor creates a new write-only instance of ERC20CustodyEchidnaTest, bound to a specific deployed contract.
func NewERC20CustodyEchidnaTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20CustodyEchidnaTestTransactor, error) {
	contract, err := bindERC20CustodyEchidnaTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyEchidnaTestTransactor{contract: contract}, nil
}

// NewERC20CustodyEchidnaTestFilterer creates a new log filterer instance of ERC20CustodyEchidnaTest, bound to a specific deployed contract.
func NewERC20CustodyEchidnaTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20CustodyEchidnaTestFilterer, error) {
	contract, err := bindERC20CustodyEchidnaTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyEchidnaTestFilterer{contract: contract}, nil
}

// bindERC20CustodyEchidnaTest binds a generic wrapper to an already deployed contract.
func bindERC20CustodyEchidnaTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20CustodyEchidnaTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CustodyEchidnaTest *ERC20CustodyEchidnaTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CustodyEchidnaTest.Contract.ERC20CustodyEchidnaTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CustodyEchidnaTest *ERC20CustodyEchidnaTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyEchidnaTest.Contract.ERC20CustodyEchidnaTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CustodyEchidnaTest *ERC20CustodyEchidnaTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CustodyEchidnaTest.Contract.ERC20CustodyEchidnaTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CustodyEchidnaTest *ERC20CustodyEchidnaTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CustodyEchidnaTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CustodyEchidnaTest *ERC20CustodyEchidnaTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyEchidnaTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CustodyEchidnaTest *ERC20CustodyEchidnaTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CustodyEchidnaTest.Contract.contract.Transact(opts, method, params...)
}
