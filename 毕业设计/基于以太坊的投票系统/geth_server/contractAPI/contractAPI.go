// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAPI

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

// VoteSystemVote is an auto generated low-level Go binding around an user-defined struct.
type VoteSystemVote struct {
	Voted bool
	Idx   uint8
	Time  uint64
}

// VoteSystemMetaData contains all meta data concerning the VoteSystem contract.
var VoteSystemMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getVote\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"idx\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"}],\"internalType\":\"structVoteSystem.Vote\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotingInfo\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"contents\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"et\",\"type\":\"uint64\"}],\"name\":\"initVotingOptions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"idx\",\"type\":\"uint8\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VoteSystemABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteSystemMetaData.ABI instead.
var VoteSystemABI = VoteSystemMetaData.ABI

// VoteSystem is an auto generated Go binding around an Ethereum contract.
type VoteSystem struct {
	VoteSystemCaller     // Read-only binding to the contract
	VoteSystemTransactor // Write-only binding to the contract
	VoteSystemFilterer   // Log filterer for contract events
}

// VoteSystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteSystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteSystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteSystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSystemSession struct {
	Contract     *VoteSystem       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteSystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteSystemCallerSession struct {
	Contract *VoteSystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VoteSystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteSystemTransactorSession struct {
	Contract     *VoteSystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VoteSystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteSystemRaw struct {
	Contract *VoteSystem // Generic contract binding to access the raw methods on
}

// VoteSystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteSystemCallerRaw struct {
	Contract *VoteSystemCaller // Generic read-only contract binding to access the raw methods on
}

// VoteSystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteSystemTransactorRaw struct {
	Contract *VoteSystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoteSystem creates a new instance of VoteSystem, bound to a specific deployed contract.
func NewVoteSystem(address common.Address, backend bind.ContractBackend) (*VoteSystem, error) {
	contract, err := bindVoteSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoteSystem{VoteSystemCaller: VoteSystemCaller{contract: contract}, VoteSystemTransactor: VoteSystemTransactor{contract: contract}, VoteSystemFilterer: VoteSystemFilterer{contract: contract}}, nil
}

// NewVoteSystemCaller creates a new read-only instance of VoteSystem, bound to a specific deployed contract.
func NewVoteSystemCaller(address common.Address, caller bind.ContractCaller) (*VoteSystemCaller, error) {
	contract, err := bindVoteSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteSystemCaller{contract: contract}, nil
}

// NewVoteSystemTransactor creates a new write-only instance of VoteSystem, bound to a specific deployed contract.
func NewVoteSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteSystemTransactor, error) {
	contract, err := bindVoteSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteSystemTransactor{contract: contract}, nil
}

// NewVoteSystemFilterer creates a new log filterer instance of VoteSystem, bound to a specific deployed contract.
func NewVoteSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteSystemFilterer, error) {
	contract, err := bindVoteSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteSystemFilterer{contract: contract}, nil
}

// bindVoteSystem binds a generic wrapper to an already deployed contract.
func bindVoteSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoteSystemMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoteSystem *VoteSystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoteSystem.Contract.VoteSystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoteSystem *VoteSystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoteSystem.Contract.VoteSystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoteSystem *VoteSystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoteSystem.Contract.VoteSystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoteSystem *VoteSystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoteSystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoteSystem *VoteSystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoteSystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoteSystem *VoteSystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoteSystem.Contract.contract.Transact(opts, method, params...)
}

// GetVote is a free data retrieval call binding the contract method 0x0242f351.
//
// Solidity: function getVote() view returns((bool,uint8,uint64))
func (_VoteSystem *VoteSystemCaller) GetVote(opts *bind.CallOpts) (VoteSystemVote, error) {
	var out []interface{}
	err := _VoteSystem.contract.Call(opts, &out, "getVote")

	if err != nil {
		return *new(VoteSystemVote), err
	}

	out0 := *abi.ConvertType(out[0], new(VoteSystemVote)).(*VoteSystemVote)

	return out0, err

}

// GetVote is a free data retrieval call binding the contract method 0x0242f351.
//
// Solidity: function getVote() view returns((bool,uint8,uint64))
func (_VoteSystem *VoteSystemSession) GetVote() (VoteSystemVote, error) {
	return _VoteSystem.Contract.GetVote(&_VoteSystem.CallOpts)
}

// GetVote is a free data retrieval call binding the contract method 0x0242f351.
//
// Solidity: function getVote() view returns((bool,uint8,uint64))
func (_VoteSystem *VoteSystemCallerSession) GetVote() (VoteSystemVote, error) {
	return _VoteSystem.Contract.GetVote(&_VoteSystem.CallOpts)
}

// GetVotingInfo is a free data retrieval call binding the contract method 0x359af4b2.
//
// Solidity: function getVotingInfo() view returns(string[], uint32[], uint64)
func (_VoteSystem *VoteSystemCaller) GetVotingInfo(opts *bind.CallOpts) ([]string, []uint32, uint64, error) {
	var out []interface{}
	err := _VoteSystem.contract.Call(opts, &out, "getVotingInfo")

	if err != nil {
		return *new([]string), *new([]uint32), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new([]uint32)).(*[]uint32)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// GetVotingInfo is a free data retrieval call binding the contract method 0x359af4b2.
//
// Solidity: function getVotingInfo() view returns(string[], uint32[], uint64)
func (_VoteSystem *VoteSystemSession) GetVotingInfo() ([]string, []uint32, uint64, error) {
	return _VoteSystem.Contract.GetVotingInfo(&_VoteSystem.CallOpts)
}

// GetVotingInfo is a free data retrieval call binding the contract method 0x359af4b2.
//
// Solidity: function getVotingInfo() view returns(string[], uint32[], uint64)
func (_VoteSystem *VoteSystemCallerSession) GetVotingInfo() ([]string, []uint32, uint64, error) {
	return _VoteSystem.Contract.GetVotingInfo(&_VoteSystem.CallOpts)
}

// InitVotingOptions is a paid mutator transaction binding the contract method 0xb1aec8c8.
//
// Solidity: function initVotingOptions(string[] contents, uint64 et) returns(bool)
func (_VoteSystem *VoteSystemTransactor) InitVotingOptions(opts *bind.TransactOpts, contents []string, et uint64) (*types.Transaction, error) {
	return _VoteSystem.contract.Transact(opts, "initVotingOptions", contents, et)
}

// InitVotingOptions is a paid mutator transaction binding the contract method 0xb1aec8c8.
//
// Solidity: function initVotingOptions(string[] contents, uint64 et) returns(bool)
func (_VoteSystem *VoteSystemSession) InitVotingOptions(contents []string, et uint64) (*types.Transaction, error) {
	return _VoteSystem.Contract.InitVotingOptions(&_VoteSystem.TransactOpts, contents, et)
}

// InitVotingOptions is a paid mutator transaction binding the contract method 0xb1aec8c8.
//
// Solidity: function initVotingOptions(string[] contents, uint64 et) returns(bool)
func (_VoteSystem *VoteSystemTransactorSession) InitVotingOptions(contents []string, et uint64) (*types.Transaction, error) {
	return _VoteSystem.Contract.InitVotingOptions(&_VoteSystem.TransactOpts, contents, et)
}

// Vote is a paid mutator transaction binding the contract method 0xb3f98adc.
//
// Solidity: function vote(uint8 idx) returns(bool)
func (_VoteSystem *VoteSystemTransactor) Vote(opts *bind.TransactOpts, idx uint8) (*types.Transaction, error) {
	return _VoteSystem.contract.Transact(opts, "vote", idx)
}

// Vote is a paid mutator transaction binding the contract method 0xb3f98adc.
//
// Solidity: function vote(uint8 idx) returns(bool)
func (_VoteSystem *VoteSystemSession) Vote(idx uint8) (*types.Transaction, error) {
	return _VoteSystem.Contract.Vote(&_VoteSystem.TransactOpts, idx)
}

// Vote is a paid mutator transaction binding the contract method 0xb3f98adc.
//
// Solidity: function vote(uint8 idx) returns(bool)
func (_VoteSystem *VoteSystemTransactorSession) Vote(idx uint8) (*types.Transaction, error) {
	return _VoteSystem.Contract.Vote(&_VoteSystem.TransactOpts, idx)
}
