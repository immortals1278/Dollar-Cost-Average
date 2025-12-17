// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc721

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

// IERC721TokenReceiverMetaData contains all meta data concerning the IERC721TokenReceiver contract.
var IERC721TokenReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IERC721TokenReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721TokenReceiverMetaData.ABI instead.
var IERC721TokenReceiverABI = IERC721TokenReceiverMetaData.ABI

// IERC721TokenReceiver is an auto generated Go binding around an Ethereum contract.
type IERC721TokenReceiver struct {
	IERC721TokenReceiverCaller     // Read-only binding to the contract
	IERC721TokenReceiverTransactor // Write-only binding to the contract
	IERC721TokenReceiverFilterer   // Log filterer for contract events
}

// IERC721TokenReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721TokenReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721TokenReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721TokenReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721TokenReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721TokenReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721TokenReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721TokenReceiverSession struct {
	Contract     *IERC721TokenReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IERC721TokenReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721TokenReceiverCallerSession struct {
	Contract *IERC721TokenReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IERC721TokenReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TokenReceiverTransactorSession struct {
	Contract     *IERC721TokenReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IERC721TokenReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721TokenReceiverRaw struct {
	Contract *IERC721TokenReceiver // Generic contract binding to access the raw methods on
}

// IERC721TokenReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721TokenReceiverCallerRaw struct {
	Contract *IERC721TokenReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721TokenReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TokenReceiverTransactorRaw struct {
	Contract *IERC721TokenReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721TokenReceiver creates a new instance of IERC721TokenReceiver, bound to a specific deployed contract.
func NewIERC721TokenReceiver(address common.Address, backend bind.ContractBackend) (*IERC721TokenReceiver, error) {
	contract, err := bindIERC721TokenReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721TokenReceiver{IERC721TokenReceiverCaller: IERC721TokenReceiverCaller{contract: contract}, IERC721TokenReceiverTransactor: IERC721TokenReceiverTransactor{contract: contract}, IERC721TokenReceiverFilterer: IERC721TokenReceiverFilterer{contract: contract}}, nil
}

// NewIERC721TokenReceiverCaller creates a new read-only instance of IERC721TokenReceiver, bound to a specific deployed contract.
func NewIERC721TokenReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC721TokenReceiverCaller, error) {
	contract, err := bindIERC721TokenReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721TokenReceiverCaller{contract: contract}, nil
}

// NewIERC721TokenReceiverTransactor creates a new write-only instance of IERC721TokenReceiver, bound to a specific deployed contract.
func NewIERC721TokenReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721TokenReceiverTransactor, error) {
	contract, err := bindIERC721TokenReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721TokenReceiverTransactor{contract: contract}, nil
}

// NewIERC721TokenReceiverFilterer creates a new log filterer instance of IERC721TokenReceiver, bound to a specific deployed contract.
func NewIERC721TokenReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721TokenReceiverFilterer, error) {
	contract, err := bindIERC721TokenReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721TokenReceiverFilterer{contract: contract}, nil
}

// bindIERC721TokenReceiver binds a generic wrapper to an already deployed contract.
func bindIERC721TokenReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC721TokenReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721TokenReceiver *IERC721TokenReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721TokenReceiver.Contract.IERC721TokenReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721TokenReceiver *IERC721TokenReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721TokenReceiver.Contract.IERC721TokenReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721TokenReceiver *IERC721TokenReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721TokenReceiver.Contract.IERC721TokenReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721TokenReceiver *IERC721TokenReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721TokenReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721TokenReceiver *IERC721TokenReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721TokenReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721TokenReceiver *IERC721TokenReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721TokenReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_IERC721TokenReceiver *IERC721TokenReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC721TokenReceiver.contract.Transact(opts, "onERC721Received", _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_IERC721TokenReceiver *IERC721TokenReceiverSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC721TokenReceiver.Contract.OnERC721Received(&_IERC721TokenReceiver.TransactOpts, _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_IERC721TokenReceiver *IERC721TokenReceiverTransactorSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC721TokenReceiver.Contract.OnERC721Received(&_IERC721TokenReceiver.TransactOpts, _operator, _from, _tokenId, _data)
}
