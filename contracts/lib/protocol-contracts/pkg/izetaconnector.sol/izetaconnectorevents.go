// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izetaconnector

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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// IZetaConnectorEventsMetaData contains all meta data concerning the IZetaConnectorEvents contract.
var IZetaConnectorEventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"UpdatedZetaConnectorTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false}]",
}

// IZetaConnectorEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IZetaConnectorEventsMetaData.ABI instead.
var IZetaConnectorEventsABI = IZetaConnectorEventsMetaData.ABI

// IZetaConnectorEvents is an auto generated Go binding around an Ethereum contract.
type IZetaConnectorEvents struct {
	IZetaConnectorEventsCaller     // Read-only binding to the contract
	IZetaConnectorEventsTransactor // Write-only binding to the contract
	IZetaConnectorEventsFilterer   // Log filterer for contract events
}

// IZetaConnectorEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZetaConnectorEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaConnectorEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZetaConnectorEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaConnectorEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZetaConnectorEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaConnectorEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZetaConnectorEventsSession struct {
	Contract     *IZetaConnectorEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IZetaConnectorEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZetaConnectorEventsCallerSession struct {
	Contract *IZetaConnectorEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IZetaConnectorEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZetaConnectorEventsTransactorSession struct {
	Contract     *IZetaConnectorEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IZetaConnectorEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZetaConnectorEventsRaw struct {
	Contract *IZetaConnectorEvents // Generic contract binding to access the raw methods on
}

// IZetaConnectorEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZetaConnectorEventsCallerRaw struct {
	Contract *IZetaConnectorEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IZetaConnectorEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZetaConnectorEventsTransactorRaw struct {
	Contract *IZetaConnectorEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZetaConnectorEvents creates a new instance of IZetaConnectorEvents, bound to a specific deployed contract.
func NewIZetaConnectorEvents(address common.Address, backend bind.ContractBackend) (*IZetaConnectorEvents, error) {
	contract, err := bindIZetaConnectorEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEvents{IZetaConnectorEventsCaller: IZetaConnectorEventsCaller{contract: contract}, IZetaConnectorEventsTransactor: IZetaConnectorEventsTransactor{contract: contract}, IZetaConnectorEventsFilterer: IZetaConnectorEventsFilterer{contract: contract}}, nil
}

// NewIZetaConnectorEventsCaller creates a new read-only instance of IZetaConnectorEvents, bound to a specific deployed contract.
func NewIZetaConnectorEventsCaller(address common.Address, caller bind.ContractCaller) (*IZetaConnectorEventsCaller, error) {
	contract, err := bindIZetaConnectorEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsCaller{contract: contract}, nil
}

// NewIZetaConnectorEventsTransactor creates a new write-only instance of IZetaConnectorEvents, bound to a specific deployed contract.
func NewIZetaConnectorEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IZetaConnectorEventsTransactor, error) {
	contract, err := bindIZetaConnectorEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsTransactor{contract: contract}, nil
}

// NewIZetaConnectorEventsFilterer creates a new log filterer instance of IZetaConnectorEvents, bound to a specific deployed contract.
func NewIZetaConnectorEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IZetaConnectorEventsFilterer, error) {
	contract, err := bindIZetaConnectorEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsFilterer{contract: contract}, nil
}

// bindIZetaConnectorEvents binds a generic wrapper to an already deployed contract.
func bindIZetaConnectorEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZetaConnectorEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaConnectorEvents *IZetaConnectorEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaConnectorEvents.Contract.IZetaConnectorEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaConnectorEvents *IZetaConnectorEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaConnectorEvents.Contract.IZetaConnectorEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaConnectorEvents *IZetaConnectorEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaConnectorEvents.Contract.IZetaConnectorEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaConnectorEvents *IZetaConnectorEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaConnectorEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaConnectorEvents *IZetaConnectorEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaConnectorEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaConnectorEvents *IZetaConnectorEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaConnectorEvents.Contract.contract.Transact(opts, method, params...)
}

// IZetaConnectorEventsUpdatedZetaConnectorTSSAddressIterator is returned from FilterUpdatedZetaConnectorTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedZetaConnectorTSSAddress events raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsUpdatedZetaConnectorTSSAddressIterator struct {
	Event *IZetaConnectorEventsUpdatedZetaConnectorTSSAddress // Event containing the contract specifics and raw log

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
func (it *IZetaConnectorEventsUpdatedZetaConnectorTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaConnectorEventsUpdatedZetaConnectorTSSAddress)
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
		it.Event = new(IZetaConnectorEventsUpdatedZetaConnectorTSSAddress)
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
func (it *IZetaConnectorEventsUpdatedZetaConnectorTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaConnectorEventsUpdatedZetaConnectorTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaConnectorEventsUpdatedZetaConnectorTSSAddress represents a UpdatedZetaConnectorTSSAddress event raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsUpdatedZetaConnectorTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedZetaConnectorTSSAddress is a free log retrieval operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) FilterUpdatedZetaConnectorTSSAddress(opts *bind.FilterOpts) (*IZetaConnectorEventsUpdatedZetaConnectorTSSAddressIterator, error) {

	logs, sub, err := _IZetaConnectorEvents.contract.FilterLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsUpdatedZetaConnectorTSSAddressIterator{contract: _IZetaConnectorEvents.contract, event: "UpdatedZetaConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedZetaConnectorTSSAddress is a free log subscription operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) WatchUpdatedZetaConnectorTSSAddress(opts *bind.WatchOpts, sink chan<- *IZetaConnectorEventsUpdatedZetaConnectorTSSAddress) (event.Subscription, error) {

	logs, sub, err := _IZetaConnectorEvents.contract.WatchLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaConnectorEventsUpdatedZetaConnectorTSSAddress)
				if err := _IZetaConnectorEvents.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
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

// ParseUpdatedZetaConnectorTSSAddress is a log parse operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) ParseUpdatedZetaConnectorTSSAddress(log types.Log) (*IZetaConnectorEventsUpdatedZetaConnectorTSSAddress, error) {
	event := new(IZetaConnectorEventsUpdatedZetaConnectorTSSAddress)
	if err := _IZetaConnectorEvents.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaConnectorEventsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawnIterator struct {
	Event *IZetaConnectorEventsWithdrawn // Event containing the contract specifics and raw log

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
func (it *IZetaConnectorEventsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaConnectorEventsWithdrawn)
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
		it.Event = new(IZetaConnectorEventsWithdrawn)
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
func (it *IZetaConnectorEventsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaConnectorEventsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaConnectorEventsWithdrawn represents a Withdrawn event raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*IZetaConnectorEventsWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsWithdrawnIterator{contract: _IZetaConnectorEvents.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IZetaConnectorEventsWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaConnectorEventsWithdrawn)
				if err := _IZetaConnectorEvents.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) ParseWithdrawn(log types.Log) (*IZetaConnectorEventsWithdrawn, error) {
	event := new(IZetaConnectorEventsWithdrawn)
	if err := _IZetaConnectorEvents.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaConnectorEventsWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawnAndCalledIterator struct {
	Event *IZetaConnectorEventsWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *IZetaConnectorEventsWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaConnectorEventsWithdrawnAndCalled)
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
		it.Event = new(IZetaConnectorEventsWithdrawnAndCalled)
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
func (it *IZetaConnectorEventsWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaConnectorEventsWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaConnectorEventsWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawnAndCalled struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address) (*IZetaConnectorEventsWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsWithdrawnAndCalledIterator{contract: _IZetaConnectorEvents.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *IZetaConnectorEventsWithdrawnAndCalled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaConnectorEventsWithdrawnAndCalled)
				if err := _IZetaConnectorEvents.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) ParseWithdrawnAndCalled(log types.Log) (*IZetaConnectorEventsWithdrawnAndCalled, error) {
	event := new(IZetaConnectorEventsWithdrawnAndCalled)
	if err := _IZetaConnectorEvents.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaConnectorEventsWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawnAndRevertedIterator struct {
	Event *IZetaConnectorEventsWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *IZetaConnectorEventsWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaConnectorEventsWithdrawnAndReverted)
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
		it.Event = new(IZetaConnectorEventsWithdrawnAndReverted)
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
func (it *IZetaConnectorEventsWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaConnectorEventsWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaConnectorEventsWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawnAndReverted struct {
	To            common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address) (*IZetaConnectorEventsWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsWithdrawnAndRevertedIterator{contract: _IZetaConnectorEvents.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *IZetaConnectorEventsWithdrawnAndReverted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaConnectorEventsWithdrawnAndReverted)
				if err := _IZetaConnectorEvents.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) ParseWithdrawnAndReverted(log types.Log) (*IZetaConnectorEventsWithdrawnAndReverted, error) {
	event := new(IZetaConnectorEventsWithdrawnAndReverted)
	if err := _IZetaConnectorEvents.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
