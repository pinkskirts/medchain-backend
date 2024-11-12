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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prescriptionAddress\",\"type\":\"address\"}],\"name\":\"AddressStored\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllPrescriptionAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPrescriptionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prescriptionAddress\",\"type\":\"address\"}],\"name\":\"getPrescriptionDetailsByAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signingDoctor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"expirationDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"patient\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dosageInstructions\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"medications\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prescriptionAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prescriptionAddress\",\"type\":\"address\"}],\"name\":\"storePrescriptionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// GetAllPrescriptionAddresses is a free data retrieval call binding the contract method 0x544aaf67.
//
// Solidity: function getAllPrescriptionAddresses() view returns(address[])
func (_Storage *StorageCaller) GetAllPrescriptionAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getAllPrescriptionAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllPrescriptionAddresses is a free data retrieval call binding the contract method 0x544aaf67.
//
// Solidity: function getAllPrescriptionAddresses() view returns(address[])
func (_Storage *StorageSession) GetAllPrescriptionAddresses() ([]common.Address, error) {
	return _Storage.Contract.GetAllPrescriptionAddresses(&_Storage.CallOpts)
}

// GetAllPrescriptionAddresses is a free data retrieval call binding the contract method 0x544aaf67.
//
// Solidity: function getAllPrescriptionAddresses() view returns(address[])
func (_Storage *StorageCallerSession) GetAllPrescriptionAddresses() ([]common.Address, error) {
	return _Storage.Contract.GetAllPrescriptionAddresses(&_Storage.CallOpts)
}

// GetPrescriptionAddress is a free data retrieval call binding the contract method 0x54760284.
//
// Solidity: function getPrescriptionAddress(uint256 index) view returns(address)
func (_Storage *StorageCaller) GetPrescriptionAddress(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPrescriptionAddress", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPrescriptionAddress is a free data retrieval call binding the contract method 0x54760284.
//
// Solidity: function getPrescriptionAddress(uint256 index) view returns(address)
func (_Storage *StorageSession) GetPrescriptionAddress(index *big.Int) (common.Address, error) {
	return _Storage.Contract.GetPrescriptionAddress(&_Storage.CallOpts, index)
}

// GetPrescriptionAddress is a free data retrieval call binding the contract method 0x54760284.
//
// Solidity: function getPrescriptionAddress(uint256 index) view returns(address)
func (_Storage *StorageCallerSession) GetPrescriptionAddress(index *big.Int) (common.Address, error) {
	return _Storage.Contract.GetPrescriptionAddress(&_Storage.CallOpts, index)
}

// GetPrescriptionDetailsByAddress is a free data retrieval call binding the contract method 0x6498f28a.
//
// Solidity: function getPrescriptionDetailsByAddress(address _prescriptionAddress) view returns(address signingDoctor, string expirationDate, string patient, string dosageInstructions, string[] medications, bool isValid)
func (_Storage *StorageCaller) GetPrescriptionDetailsByAddress(opts *bind.CallOpts, _prescriptionAddress common.Address) (struct {
	SigningDoctor      common.Address
	ExpirationDate     string
	Patient            string
	DosageInstructions string
	Medications        []string
	IsValid            bool
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPrescriptionDetailsByAddress", _prescriptionAddress)

	outstruct := new(struct {
		SigningDoctor      common.Address
		ExpirationDate     string
		Patient            string
		DosageInstructions string
		Medications        []string
		IsValid            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SigningDoctor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ExpirationDate = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Patient = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.DosageInstructions = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Medications = *abi.ConvertType(out[4], new([]string)).(*[]string)
	outstruct.IsValid = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetPrescriptionDetailsByAddress is a free data retrieval call binding the contract method 0x6498f28a.
//
// Solidity: function getPrescriptionDetailsByAddress(address _prescriptionAddress) view returns(address signingDoctor, string expirationDate, string patient, string dosageInstructions, string[] medications, bool isValid)
func (_Storage *StorageSession) GetPrescriptionDetailsByAddress(_prescriptionAddress common.Address) (struct {
	SigningDoctor      common.Address
	ExpirationDate     string
	Patient            string
	DosageInstructions string
	Medications        []string
	IsValid            bool
}, error) {
	return _Storage.Contract.GetPrescriptionDetailsByAddress(&_Storage.CallOpts, _prescriptionAddress)
}

// GetPrescriptionDetailsByAddress is a free data retrieval call binding the contract method 0x6498f28a.
//
// Solidity: function getPrescriptionDetailsByAddress(address _prescriptionAddress) view returns(address signingDoctor, string expirationDate, string patient, string dosageInstructions, string[] medications, bool isValid)
func (_Storage *StorageCallerSession) GetPrescriptionDetailsByAddress(_prescriptionAddress common.Address) (struct {
	SigningDoctor      common.Address
	ExpirationDate     string
	Patient            string
	DosageInstructions string
	Medications        []string
	IsValid            bool
}, error) {
	return _Storage.Contract.GetPrescriptionDetailsByAddress(&_Storage.CallOpts, _prescriptionAddress)
}

// PrescriptionAddresses is a free data retrieval call binding the contract method 0x8fb0ab00.
//
// Solidity: function prescriptionAddresses(uint256 ) view returns(address)
func (_Storage *StorageCaller) PrescriptionAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "prescriptionAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrescriptionAddresses is a free data retrieval call binding the contract method 0x8fb0ab00.
//
// Solidity: function prescriptionAddresses(uint256 ) view returns(address)
func (_Storage *StorageSession) PrescriptionAddresses(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.PrescriptionAddresses(&_Storage.CallOpts, arg0)
}

// PrescriptionAddresses is a free data retrieval call binding the contract method 0x8fb0ab00.
//
// Solidity: function prescriptionAddresses(uint256 ) view returns(address)
func (_Storage *StorageCallerSession) PrescriptionAddresses(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.PrescriptionAddresses(&_Storage.CallOpts, arg0)
}

// StorePrescriptionAddress is a paid mutator transaction binding the contract method 0xee032462.
//
// Solidity: function storePrescriptionAddress(address _prescriptionAddress) returns()
func (_Storage *StorageTransactor) StorePrescriptionAddress(opts *bind.TransactOpts, _prescriptionAddress common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "storePrescriptionAddress", _prescriptionAddress)
}

// StorePrescriptionAddress is a paid mutator transaction binding the contract method 0xee032462.
//
// Solidity: function storePrescriptionAddress(address _prescriptionAddress) returns()
func (_Storage *StorageSession) StorePrescriptionAddress(_prescriptionAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.StorePrescriptionAddress(&_Storage.TransactOpts, _prescriptionAddress)
}

// StorePrescriptionAddress is a paid mutator transaction binding the contract method 0xee032462.
//
// Solidity: function storePrescriptionAddress(address _prescriptionAddress) returns()
func (_Storage *StorageTransactorSession) StorePrescriptionAddress(_prescriptionAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.StorePrescriptionAddress(&_Storage.TransactOpts, _prescriptionAddress)
}

// StorageAddressStoredIterator is returned from FilterAddressStored and is used to iterate over the raw logs and unpacked data for AddressStored events raised by the Storage contract.
type StorageAddressStoredIterator struct {
	Event *StorageAddressStored // Event containing the contract specifics and raw log

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
func (it *StorageAddressStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAddressStored)
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
		it.Event = new(StorageAddressStored)
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
func (it *StorageAddressStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAddressStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAddressStored represents a AddressStored event raised by the Storage contract.
type StorageAddressStored struct {
	PrescriptionAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAddressStored is a free log retrieval operation binding the contract event 0x03bfc5d7eb2573c750c90a0e020d3d1124a447e5acd965a323ad0abc11f8f8b4.
//
// Solidity: event AddressStored(address indexed prescriptionAddress)
func (_Storage *StorageFilterer) FilterAddressStored(opts *bind.FilterOpts, prescriptionAddress []common.Address) (*StorageAddressStoredIterator, error) {

	var prescriptionAddressRule []interface{}
	for _, prescriptionAddressItem := range prescriptionAddress {
		prescriptionAddressRule = append(prescriptionAddressRule, prescriptionAddressItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AddressStored", prescriptionAddressRule)
	if err != nil {
		return nil, err
	}
	return &StorageAddressStoredIterator{contract: _Storage.contract, event: "AddressStored", logs: logs, sub: sub}, nil
}

// WatchAddressStored is a free log subscription operation binding the contract event 0x03bfc5d7eb2573c750c90a0e020d3d1124a447e5acd965a323ad0abc11f8f8b4.
//
// Solidity: event AddressStored(address indexed prescriptionAddress)
func (_Storage *StorageFilterer) WatchAddressStored(opts *bind.WatchOpts, sink chan<- *StorageAddressStored, prescriptionAddress []common.Address) (event.Subscription, error) {

	var prescriptionAddressRule []interface{}
	for _, prescriptionAddressItem := range prescriptionAddress {
		prescriptionAddressRule = append(prescriptionAddressRule, prescriptionAddressItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AddressStored", prescriptionAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAddressStored)
				if err := _Storage.contract.UnpackLog(event, "AddressStored", log); err != nil {
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

// ParseAddressStored is a log parse operation binding the contract event 0x03bfc5d7eb2573c750c90a0e020d3d1124a447e5acd965a323ad0abc11f8f8b4.
//
// Solidity: event AddressStored(address indexed prescriptionAddress)
func (_Storage *StorageFilterer) ParseAddressStored(log types.Log) (*StorageAddressStored, error) {
	event := new(StorageAddressStored)
	if err := _Storage.contract.UnpackLog(event, "AddressStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
