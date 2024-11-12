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

// PrescriptionMetaData contains all meta data concerning the Prescription contract.
var PrescriptionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signingDoctor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_expirationDate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_patient\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dosageInstructions\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_medications\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"_isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"dosageInstructions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"expirationDate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPrescriptionDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"medications\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"patient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signingDoctor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PrescriptionABI is the input ABI used to generate the binding from.
// Deprecated: Use PrescriptionMetaData.ABI instead.
var PrescriptionABI = PrescriptionMetaData.ABI

// Prescription is an auto generated Go binding around an Ethereum contract.
type Prescription struct {
	PrescriptionCaller     // Read-only binding to the contract
	PrescriptionTransactor // Write-only binding to the contract
	PrescriptionFilterer   // Log filterer for contract events
}

// PrescriptionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrescriptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrescriptionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrescriptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrescriptionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrescriptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrescriptionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrescriptionSession struct {
	Contract     *Prescription     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrescriptionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrescriptionCallerSession struct {
	Contract *PrescriptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PrescriptionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrescriptionTransactorSession struct {
	Contract     *PrescriptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PrescriptionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrescriptionRaw struct {
	Contract *Prescription // Generic contract binding to access the raw methods on
}

// PrescriptionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrescriptionCallerRaw struct {
	Contract *PrescriptionCaller // Generic read-only contract binding to access the raw methods on
}

// PrescriptionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrescriptionTransactorRaw struct {
	Contract *PrescriptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrescription creates a new instance of Prescription, bound to a specific deployed contract.
func NewPrescription(address common.Address, backend bind.ContractBackend) (*Prescription, error) {
	contract, err := bindPrescription(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Prescription{PrescriptionCaller: PrescriptionCaller{contract: contract}, PrescriptionTransactor: PrescriptionTransactor{contract: contract}, PrescriptionFilterer: PrescriptionFilterer{contract: contract}}, nil
}

// NewPrescriptionCaller creates a new read-only instance of Prescription, bound to a specific deployed contract.
func NewPrescriptionCaller(address common.Address, caller bind.ContractCaller) (*PrescriptionCaller, error) {
	contract, err := bindPrescription(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrescriptionCaller{contract: contract}, nil
}

// NewPrescriptionTransactor creates a new write-only instance of Prescription, bound to a specific deployed contract.
func NewPrescriptionTransactor(address common.Address, transactor bind.ContractTransactor) (*PrescriptionTransactor, error) {
	contract, err := bindPrescription(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrescriptionTransactor{contract: contract}, nil
}

// NewPrescriptionFilterer creates a new log filterer instance of Prescription, bound to a specific deployed contract.
func NewPrescriptionFilterer(address common.Address, filterer bind.ContractFilterer) (*PrescriptionFilterer, error) {
	contract, err := bindPrescription(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrescriptionFilterer{contract: contract}, nil
}

// bindPrescription binds a generic wrapper to an already deployed contract.
func bindPrescription(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrescriptionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prescription *PrescriptionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prescription.Contract.PrescriptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prescription *PrescriptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prescription.Contract.PrescriptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prescription *PrescriptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prescription.Contract.PrescriptionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prescription *PrescriptionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prescription.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prescription *PrescriptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prescription.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prescription *PrescriptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prescription.Contract.contract.Transact(opts, method, params...)
}

// DosageInstructions is a free data retrieval call binding the contract method 0xe44f1992.
//
// Solidity: function dosageInstructions() view returns(string)
func (_Prescription *PrescriptionCaller) DosageInstructions(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Prescription.contract.Call(opts, &out, "dosageInstructions")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DosageInstructions is a free data retrieval call binding the contract method 0xe44f1992.
//
// Solidity: function dosageInstructions() view returns(string)
func (_Prescription *PrescriptionSession) DosageInstructions() (string, error) {
	return _Prescription.Contract.DosageInstructions(&_Prescription.CallOpts)
}

// DosageInstructions is a free data retrieval call binding the contract method 0xe44f1992.
//
// Solidity: function dosageInstructions() view returns(string)
func (_Prescription *PrescriptionCallerSession) DosageInstructions() (string, error) {
	return _Prescription.Contract.DosageInstructions(&_Prescription.CallOpts)
}

// ExpirationDate is a free data retrieval call binding the contract method 0x8f620487.
//
// Solidity: function expirationDate() view returns(string)
func (_Prescription *PrescriptionCaller) ExpirationDate(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Prescription.contract.Call(opts, &out, "expirationDate")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExpirationDate is a free data retrieval call binding the contract method 0x8f620487.
//
// Solidity: function expirationDate() view returns(string)
func (_Prescription *PrescriptionSession) ExpirationDate() (string, error) {
	return _Prescription.Contract.ExpirationDate(&_Prescription.CallOpts)
}

// ExpirationDate is a free data retrieval call binding the contract method 0x8f620487.
//
// Solidity: function expirationDate() view returns(string)
func (_Prescription *PrescriptionCallerSession) ExpirationDate() (string, error) {
	return _Prescription.Contract.ExpirationDate(&_Prescription.CallOpts)
}

// GetPrescriptionDetails is a free data retrieval call binding the contract method 0x311e6aec.
//
// Solidity: function getPrescriptionDetails() view returns(address, string, string, string, string[], bool)
func (_Prescription *PrescriptionCaller) GetPrescriptionDetails(opts *bind.CallOpts) (common.Address, string, string, string, []string, bool, error) {
	var out []interface{}
	err := _Prescription.contract.Call(opts, &out, "getPrescriptionDetails")

	if err != nil {
		return *new(common.Address), *new(string), *new(string), *new(string), *new([]string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new([]string)).(*[]string)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// GetPrescriptionDetails is a free data retrieval call binding the contract method 0x311e6aec.
//
// Solidity: function getPrescriptionDetails() view returns(address, string, string, string, string[], bool)
func (_Prescription *PrescriptionSession) GetPrescriptionDetails() (common.Address, string, string, string, []string, bool, error) {
	return _Prescription.Contract.GetPrescriptionDetails(&_Prescription.CallOpts)
}

// GetPrescriptionDetails is a free data retrieval call binding the contract method 0x311e6aec.
//
// Solidity: function getPrescriptionDetails() view returns(address, string, string, string, string[], bool)
func (_Prescription *PrescriptionCallerSession) GetPrescriptionDetails() (common.Address, string, string, string, []string, bool, error) {
	return _Prescription.Contract.GetPrescriptionDetails(&_Prescription.CallOpts)
}

// IsValid is a free data retrieval call binding the contract method 0xbb5d40eb.
//
// Solidity: function isValid() view returns(bool)
func (_Prescription *PrescriptionCaller) IsValid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Prescription.contract.Call(opts, &out, "isValid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0xbb5d40eb.
//
// Solidity: function isValid() view returns(bool)
func (_Prescription *PrescriptionSession) IsValid() (bool, error) {
	return _Prescription.Contract.IsValid(&_Prescription.CallOpts)
}

// IsValid is a free data retrieval call binding the contract method 0xbb5d40eb.
//
// Solidity: function isValid() view returns(bool)
func (_Prescription *PrescriptionCallerSession) IsValid() (bool, error) {
	return _Prescription.Contract.IsValid(&_Prescription.CallOpts)
}

// Medications is a free data retrieval call binding the contract method 0xfa8547f6.
//
// Solidity: function medications(uint256 ) view returns(string)
func (_Prescription *PrescriptionCaller) Medications(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Prescription.contract.Call(opts, &out, "medications", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Medications is a free data retrieval call binding the contract method 0xfa8547f6.
//
// Solidity: function medications(uint256 ) view returns(string)
func (_Prescription *PrescriptionSession) Medications(arg0 *big.Int) (string, error) {
	return _Prescription.Contract.Medications(&_Prescription.CallOpts, arg0)
}

// Medications is a free data retrieval call binding the contract method 0xfa8547f6.
//
// Solidity: function medications(uint256 ) view returns(string)
func (_Prescription *PrescriptionCallerSession) Medications(arg0 *big.Int) (string, error) {
	return _Prescription.Contract.Medications(&_Prescription.CallOpts, arg0)
}

// Patient is a free data retrieval call binding the contract method 0xbd96bd20.
//
// Solidity: function patient() view returns(string)
func (_Prescription *PrescriptionCaller) Patient(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Prescription.contract.Call(opts, &out, "patient")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Patient is a free data retrieval call binding the contract method 0xbd96bd20.
//
// Solidity: function patient() view returns(string)
func (_Prescription *PrescriptionSession) Patient() (string, error) {
	return _Prescription.Contract.Patient(&_Prescription.CallOpts)
}

// Patient is a free data retrieval call binding the contract method 0xbd96bd20.
//
// Solidity: function patient() view returns(string)
func (_Prescription *PrescriptionCallerSession) Patient() (string, error) {
	return _Prescription.Contract.Patient(&_Prescription.CallOpts)
}

// SigningDoctor is a free data retrieval call binding the contract method 0xf8a49443.
//
// Solidity: function signingDoctor() view returns(address)
func (_Prescription *PrescriptionCaller) SigningDoctor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prescription.contract.Call(opts, &out, "signingDoctor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigningDoctor is a free data retrieval call binding the contract method 0xf8a49443.
//
// Solidity: function signingDoctor() view returns(address)
func (_Prescription *PrescriptionSession) SigningDoctor() (common.Address, error) {
	return _Prescription.Contract.SigningDoctor(&_Prescription.CallOpts)
}

// SigningDoctor is a free data retrieval call binding the contract method 0xf8a49443.
//
// Solidity: function signingDoctor() view returns(address)
func (_Prescription *PrescriptionCallerSession) SigningDoctor() (common.Address, error) {
	return _Prescription.Contract.SigningDoctor(&_Prescription.CallOpts)
}
