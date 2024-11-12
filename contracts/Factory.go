// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prescriptionAddress\",\"type\":\"address\"}],\"name\":\"PrescriptionCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_expirationDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_patient\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dosageInstructions\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_medications\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"_isValid\",\"type\":\"bool\"}],\"name\":\"createPrescription\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// CreatePrescription is a paid mutator transaction binding the contract method 0x311aa8ca.
//
// Solidity: function createPrescription(string _expirationDate, string _patient, string _dosageInstructions, string[] _medications, bool _isValid) returns()
func (_Factory *FactoryTransactor) CreatePrescription(opts *bind.TransactOpts, _expirationDate string, _patient string, _dosageInstructions string, _medications []string, _isValid bool) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createPrescription", _expirationDate, _patient, _dosageInstructions, _medications, _isValid)
}

// CreatePrescription is a paid mutator transaction binding the contract method 0x311aa8ca.
//
// Solidity: function createPrescription(string _expirationDate, string _patient, string _dosageInstructions, string[] _medications, bool _isValid) returns()
func (_Factory *FactorySession) CreatePrescription(_expirationDate string, _patient string, _dosageInstructions string, _medications []string, _isValid bool) (*types.Transaction, error) {
	return _Factory.Contract.CreatePrescription(&_Factory.TransactOpts, _expirationDate, _patient, _dosageInstructions, _medications, _isValid)
}

// CreatePrescription is a paid mutator transaction binding the contract method 0x311aa8ca.
//
// Solidity: function createPrescription(string _expirationDate, string _patient, string _dosageInstructions, string[] _medications, bool _isValid) returns()
func (_Factory *FactoryTransactorSession) CreatePrescription(_expirationDate string, _patient string, _dosageInstructions string, _medications []string, _isValid bool) (*types.Transaction, error) {
	return _Factory.Contract.CreatePrescription(&_Factory.TransactOpts, _expirationDate, _patient, _dosageInstructions, _medications, _isValid)
}

// FactoryPrescriptionCreatedIterator is returned from FilterPrescriptionCreated and is used to iterate over the raw logs and unpacked data for PrescriptionCreated events raised by the Factory contract.
type FactoryPrescriptionCreatedIterator struct {
	Event *FactoryPrescriptionCreated // Event containing the contract specifics and raw log

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
func (it *FactoryPrescriptionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryPrescriptionCreated)
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
		it.Event = new(FactoryPrescriptionCreated)
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
func (it *FactoryPrescriptionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryPrescriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryPrescriptionCreated represents a PrescriptionCreated event raised by the Factory contract.
type FactoryPrescriptionCreated struct {
	PrescriptionAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPrescriptionCreated is a free log retrieval operation binding the contract event 0x82eea93ef73d2102e1195160b1996b5343eaf036850277e04e9fafe3eb73d928.
//
// Solidity: event PrescriptionCreated(address prescriptionAddress)
func (_Factory *FactoryFilterer) FilterPrescriptionCreated(opts *bind.FilterOpts) (*FactoryPrescriptionCreatedIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "PrescriptionCreated")
	if err != nil {
		return nil, err
	}
	return &FactoryPrescriptionCreatedIterator{contract: _Factory.contract, event: "PrescriptionCreated", logs: logs, sub: sub}, nil
}

// WatchPrescriptionCreated is a free log subscription operation binding the contract event 0x82eea93ef73d2102e1195160b1996b5343eaf036850277e04e9fafe3eb73d928.
//
// Solidity: event PrescriptionCreated(address prescriptionAddress)
func (_Factory *FactoryFilterer) WatchPrescriptionCreated(opts *bind.WatchOpts, sink chan<- *FactoryPrescriptionCreated) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "PrescriptionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryPrescriptionCreated)
				if err := _Factory.contract.UnpackLog(event, "PrescriptionCreated", log); err != nil {
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

// ParsePrescriptionCreated is a log parse operation binding the contract event 0x82eea93ef73d2102e1195160b1996b5343eaf036850277e04e9fafe3eb73d928.
//
// Solidity: event PrescriptionCreated(address prescriptionAddress)
func (_Factory *FactoryFilterer) ParsePrescriptionCreated(log types.Log) (*FactoryPrescriptionCreated, error) {
	event := new(FactoryPrescriptionCreated)
	if err := _Factory.contract.UnpackLog(event, "PrescriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
