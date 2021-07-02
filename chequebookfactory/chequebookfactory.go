// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chequebookfactory

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChequeBookABI is the input ABI used to generate the binding from.
const ChequeBookABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Cashouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HardDepositAmountChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"beneficiaryBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"cashout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hardDepositBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paidOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payHardDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBeneficiary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalHardDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPaidOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChequeBookBin is the compiled bytecode used for deploying new contracts.
var ChequeBookBin = "0x608060405234801561001057600080fd5b50610d12806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063b6b55f2511610097578063e0bcf13a11610066578063e0bcf13a146101b9578063f09a4016146101c1578063f6407d0c146101d4578063fc0c546a146101e7576100f5565b8063b6b55f2514610176578063b7b172b31461018b578063b7ec1a331461019e578063c4076876146101a6576100f5565b80635bf58ec7116100d35780635bf58ec71461014057806381f03fcb14610153578063959d15ca14610166578063b69ef8a81461016e576100f5565b80631357e1dc146100fa5780631d143848146101185780635664865a1461012d575b600080fd5b6101026101ef565b60405161010f9190610cd3565b60405180910390f35b6101206101f5565b60405161010f9190610ab2565b61010261013b3660046109ed565b610204565b61010261014e3660046109ed565b610216565b6101026101613660046109ed565b610228565b61010261023a565b610102610240565b610189610184366004610a82565b6102c6565b005b610189610199366004610a39565b61037e565b61010261054b565b6101896101b4366004610a39565b610561565b610102610611565b6101896101cf366004610a07565b610617565b6101896101e2366004610a39565b610694565b6101206106f0565b60045481565b6007546001600160a01b031681565b60026020526000908152604090205481565b60016020526000908152604090205481565b60036020526000908152604090205481565b60065481565b600080546040516370a0823160e01b81526001600160a01b03909116906370a0823190610271903090600401610ab2565b60206040518083038186803b15801561028957600080fd5b505afa15801561029d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c19190610a9a565b905090565b6000546040516323b872dd60e01b81526001600160a01b03909116906323b872dd906102fa90339030908690600401610ac6565b602060405180830381600087803b15801561031457600080fd5b505af1158015610328573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034c9190610a62565b6103715760405162461bcd60e51b815260040161036890610c75565b60405180910390fd5b61037b33826106ff565b50565b6103883382610782565b336000908152600360205260409020546103a290826107f4565b336000908152600360205260409020556004546103bf90826107f4565b6004908155600054604051633950935160e01b81526001600160a01b03909116916339509351916103f4913091869101610aea565b602060405180830381600087803b15801561040e57600080fd5b505af1158015610422573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104469190610a62565b6104625760405162461bcd60e51b815260040161036890610bfd565b6000546040516323b872dd60e01b81526001600160a01b03909116906323b872dd9061049690309086908690600401610ac6565b602060405180830381600087803b1580156104b057600080fd5b505af11580156104c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e89190610a62565b6105045760405162461bcd60e51b815260040161036890610bc6565b336001600160a01b03167f92fec97bff9ddc18bbb7cf738dd42818ad2528ca913ee47fbe149a5cc822ad21838360405161053f929190610aea565b60405180910390a25050565b60006102c160055461055b610240565b90610855565b6000546040516323b872dd60e01b81526001600160a01b03909116906323b872dd9061059590339030908690600401610ac6565b602060405180830381600087803b1580156105af57600080fd5b505af11580156105c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e79190610a62565b6106035760405162461bcd60e51b815260040161036890610c75565b61060d82826108b2565b5050565b60055481565b6001600160a01b03821661063d5760405162461bcd60e51b815260040161036890610b71565b6007546001600160a01b0316156106665760405162461bcd60e51b815260040161036890610b99565b600780546001600160a01b039384166001600160a01b03199182161790915560008054929093169116179055565b61069e338261092e565b6106a882826108b2565b6001600160a01b038216600081815260026020526040908190205490517f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad9161053f91610cd3565b6000546001600160a01b031681565b610707610240565b60055461071490836107f4565b11156107325760405162461bcd60e51b815260040161036890610b3a565b6001600160a01b03821660009081526001602052604090205461075590826107f4565b6001600160a01b03831660009081526001602052604090205560055461077b90826107f4565b6005555050565b8060065410156107a45760405162461bcd60e51b815260040161036890610c9e565b6001600160a01b0382166000908152600260205260409020546107c79082610855565b6001600160a01b0383166000908152600260205260409020556006546107ed9082610855565b6006555050565b60008282018381101561084e576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000828211156108ac576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6108ba610240565b6006546108c790836107f4565b11156108e55760405162461bcd60e51b815260040161036890610b3a565b6001600160a01b03821660009081526002602052604090205461090890826107f4565b6001600160a01b0383166000908152600260205260409020556006546107ed90826107f4565b6001600160a01b0382166000908152600160205260409020548111156109665760405162461bcd60e51b815260040161036890610b03565b8060055410156109885760405162461bcd60e51b815260040161036890610c34565b6001600160a01b0382166000908152600160205260409020546109ab9082610855565b6001600160a01b03831660009081526001602052604090205560055461077b9082610855565b80356001600160a01b03811681146109e857600080fd5b919050565b6000602082840312156109fe578081fd5b61084e826109d1565b60008060408385031215610a19578081fd5b610a22836109d1565b9150610a30602084016109d1565b90509250929050565b60008060408385031215610a4b578182fd5b610a54836109d1565b946020939093013593505050565b600060208284031215610a73578081fd5b8151801515811461084e578182fd5b600060208284031215610a93578081fd5b5035919050565b600060208284031215610aab578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b6020808252601c908201527f616d6f756e742065786365656473206f776e65722062616c616e636500000000604082015260600190565b6020808252601d908201527f616d6f756e74206578636565647320676c6f62616c2062616c616e6365000000604082015260600190565b6020808252600e908201526d34b73b30b634b21034b9b9bab2b960911b604082015260600190565b602080825260139082015272185b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604082015260600190565b60208082526017908201527f636173686f7574207472616e73666572206661696c6564000000000000000000604082015260600190565b60208082526019908201527f696e63726561736520616c6c6f77616e6365206661696c656400000000000000604082015260600190565b60208082526021908201527f616d6f756e74206578636565647320746f74616c2068617264206465706f73696040820152601d60fa1b606082015260800190565b6020808252600f908201526e1d1c985b9cd9995c8819985a5b1959608a1b604082015260600190565b6020808252818101527f616d6f756e74206578636565647320746f74616c2062656e6566696369617279604082015260600190565b9081526020019056fea264697066735822122048c69c48ae8454ebfc51c0054615611553e8c5923fcb284ec676f8f7ebc36df164736f6c63430007060033"

// DeployChequeBook deploys a new Ethereum contract, binding an instance of ChequeBook to it.
func DeployChequeBook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChequeBook, error) {
	parsed, err := abi.JSON(strings.NewReader(ChequeBookABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChequeBookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChequeBook{ChequeBookCaller: ChequeBookCaller{contract: contract}, ChequeBookTransactor: ChequeBookTransactor{contract: contract}, ChequeBookFilterer: ChequeBookFilterer{contract: contract}}, nil
}

// ChequeBook is an auto generated Go binding around an Ethereum contract.
type ChequeBook struct {
	ChequeBookCaller     // Read-only binding to the contract
	ChequeBookTransactor // Write-only binding to the contract
	ChequeBookFilterer   // Log filterer for contract events
}

// ChequeBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChequeBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChequeBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChequeBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChequeBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChequeBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChequeBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChequeBookSession struct {
	Contract     *ChequeBook       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChequeBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChequeBookCallerSession struct {
	Contract *ChequeBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ChequeBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChequeBookTransactorSession struct {
	Contract     *ChequeBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ChequeBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChequeBookRaw struct {
	Contract *ChequeBook // Generic contract binding to access the raw methods on
}

// ChequeBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChequeBookCallerRaw struct {
	Contract *ChequeBookCaller // Generic read-only contract binding to access the raw methods on
}

// ChequeBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChequeBookTransactorRaw struct {
	Contract *ChequeBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChequeBook creates a new instance of ChequeBook, bound to a specific deployed contract.
func NewChequeBook(address common.Address, backend bind.ContractBackend) (*ChequeBook, error) {
	contract, err := bindChequeBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChequeBook{ChequeBookCaller: ChequeBookCaller{contract: contract}, ChequeBookTransactor: ChequeBookTransactor{contract: contract}, ChequeBookFilterer: ChequeBookFilterer{contract: contract}}, nil
}

// NewChequeBookCaller creates a new read-only instance of ChequeBook, bound to a specific deployed contract.
func NewChequeBookCaller(address common.Address, caller bind.ContractCaller) (*ChequeBookCaller, error) {
	contract, err := bindChequeBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChequeBookCaller{contract: contract}, nil
}

// NewChequeBookTransactor creates a new write-only instance of ChequeBook, bound to a specific deployed contract.
func NewChequeBookTransactor(address common.Address, transactor bind.ContractTransactor) (*ChequeBookTransactor, error) {
	contract, err := bindChequeBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChequeBookTransactor{contract: contract}, nil
}

// NewChequeBookFilterer creates a new log filterer instance of ChequeBook, bound to a specific deployed contract.
func NewChequeBookFilterer(address common.Address, filterer bind.ContractFilterer) (*ChequeBookFilterer, error) {
	contract, err := bindChequeBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChequeBookFilterer{contract: contract}, nil
}

// bindChequeBook binds a generic wrapper to an already deployed contract.
func bindChequeBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChequeBookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChequeBook *ChequeBookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChequeBook.Contract.ChequeBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChequeBook *ChequeBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChequeBook.Contract.ChequeBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChequeBook *ChequeBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChequeBook.Contract.ChequeBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChequeBook *ChequeBookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChequeBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChequeBook *ChequeBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChequeBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChequeBook *ChequeBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChequeBook.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_ChequeBook *ChequeBookCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_ChequeBook *ChequeBookSession) Balance() (*big.Int, error) {
	return _ChequeBook.Contract.Balance(&_ChequeBook.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_ChequeBook *ChequeBookCallerSession) Balance() (*big.Int, error) {
	return _ChequeBook.Contract.Balance(&_ChequeBook.CallOpts)
}

// BeneficiaryBalances is a free data retrieval call binding the contract method 0x5664865a.
//
// Solidity: function beneficiaryBalances(address ) view returns(uint256)
func (_ChequeBook *ChequeBookCaller) BeneficiaryBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "beneficiaryBalances", arg0)
	return *ret0, err
}

// BeneficiaryBalances is a free data retrieval call binding the contract method 0x5664865a.
//
// Solidity: function beneficiaryBalances(address ) view returns(uint256)
func (_ChequeBook *ChequeBookSession) BeneficiaryBalances(arg0 common.Address) (*big.Int, error) {
	return _ChequeBook.Contract.BeneficiaryBalances(&_ChequeBook.CallOpts, arg0)
}

// BeneficiaryBalances is a free data retrieval call binding the contract method 0x5664865a.
//
// Solidity: function beneficiaryBalances(address ) view returns(uint256)
func (_ChequeBook *ChequeBookCallerSession) BeneficiaryBalances(arg0 common.Address) (*big.Int, error) {
	return _ChequeBook.Contract.BeneficiaryBalances(&_ChequeBook.CallOpts, arg0)
}

// HardDepositBalances is a free data retrieval call binding the contract method 0x5bf58ec7.
//
// Solidity: function hardDepositBalances(address ) view returns(uint256)
func (_ChequeBook *ChequeBookCaller) HardDepositBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "hardDepositBalances", arg0)
	return *ret0, err
}

// HardDepositBalances is a free data retrieval call binding the contract method 0x5bf58ec7.
//
// Solidity: function hardDepositBalances(address ) view returns(uint256)
func (_ChequeBook *ChequeBookSession) HardDepositBalances(arg0 common.Address) (*big.Int, error) {
	return _ChequeBook.Contract.HardDepositBalances(&_ChequeBook.CallOpts, arg0)
}

// HardDepositBalances is a free data retrieval call binding the contract method 0x5bf58ec7.
//
// Solidity: function hardDepositBalances(address ) view returns(uint256)
func (_ChequeBook *ChequeBookCallerSession) HardDepositBalances(arg0 common.Address) (*big.Int, error) {
	return _ChequeBook.Contract.HardDepositBalances(&_ChequeBook.CallOpts, arg0)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_ChequeBook *ChequeBookCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_ChequeBook *ChequeBookSession) Issuer() (common.Address, error) {
	return _ChequeBook.Contract.Issuer(&_ChequeBook.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_ChequeBook *ChequeBookCallerSession) Issuer() (common.Address, error) {
	return _ChequeBook.Contract.Issuer(&_ChequeBook.CallOpts)
}

// LiquidBalance is a free data retrieval call binding the contract method 0xb7ec1a33.
//
// Solidity: function liquidBalance() view returns(uint256)
func (_ChequeBook *ChequeBookCaller) LiquidBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "liquidBalance")
	return *ret0, err
}

// LiquidBalance is a free data retrieval call binding the contract method 0xb7ec1a33.
//
// Solidity: function liquidBalance() view returns(uint256)
func (_ChequeBook *ChequeBookSession) LiquidBalance() (*big.Int, error) {
	return _ChequeBook.Contract.LiquidBalance(&_ChequeBook.CallOpts)
}

// LiquidBalance is a free data retrieval call binding the contract method 0xb7ec1a33.
//
// Solidity: function liquidBalance() view returns(uint256)
func (_ChequeBook *ChequeBookCallerSession) LiquidBalance() (*big.Int, error) {
	return _ChequeBook.Contract.LiquidBalance(&_ChequeBook.CallOpts)
}

// PaidOut is a free data retrieval call binding the contract method 0x81f03fcb.
//
// Solidity: function paidOut(address ) view returns(uint256)
func (_ChequeBook *ChequeBookCaller) PaidOut(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "paidOut", arg0)
	return *ret0, err
}

// PaidOut is a free data retrieval call binding the contract method 0x81f03fcb.
//
// Solidity: function paidOut(address ) view returns(uint256)
func (_ChequeBook *ChequeBookSession) PaidOut(arg0 common.Address) (*big.Int, error) {
	return _ChequeBook.Contract.PaidOut(&_ChequeBook.CallOpts, arg0)
}

// PaidOut is a free data retrieval call binding the contract method 0x81f03fcb.
//
// Solidity: function paidOut(address ) view returns(uint256)
func (_ChequeBook *ChequeBookCallerSession) PaidOut(arg0 common.Address) (*big.Int, error) {
	return _ChequeBook.Contract.PaidOut(&_ChequeBook.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ChequeBook *ChequeBookCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ChequeBook *ChequeBookSession) Token() (common.Address, error) {
	return _ChequeBook.Contract.Token(&_ChequeBook.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ChequeBook *ChequeBookCallerSession) Token() (common.Address, error) {
	return _ChequeBook.Contract.Token(&_ChequeBook.CallOpts)
}

// TotalBeneficiary is a free data retrieval call binding the contract method 0x959d15ca.
//
// Solidity: function totalBeneficiary() view returns(uint256)
func (_ChequeBook *ChequeBookCaller) TotalBeneficiary(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "totalBeneficiary")
	return *ret0, err
}

// TotalBeneficiary is a free data retrieval call binding the contract method 0x959d15ca.
//
// Solidity: function totalBeneficiary() view returns(uint256)
func (_ChequeBook *ChequeBookSession) TotalBeneficiary() (*big.Int, error) {
	return _ChequeBook.Contract.TotalBeneficiary(&_ChequeBook.CallOpts)
}

// TotalBeneficiary is a free data retrieval call binding the contract method 0x959d15ca.
//
// Solidity: function totalBeneficiary() view returns(uint256)
func (_ChequeBook *ChequeBookCallerSession) TotalBeneficiary() (*big.Int, error) {
	return _ChequeBook.Contract.TotalBeneficiary(&_ChequeBook.CallOpts)
}

// TotalHardDeposit is a free data retrieval call binding the contract method 0xe0bcf13a.
//
// Solidity: function totalHardDeposit() view returns(uint256)
func (_ChequeBook *ChequeBookCaller) TotalHardDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "totalHardDeposit")
	return *ret0, err
}

// TotalHardDeposit is a free data retrieval call binding the contract method 0xe0bcf13a.
//
// Solidity: function totalHardDeposit() view returns(uint256)
func (_ChequeBook *ChequeBookSession) TotalHardDeposit() (*big.Int, error) {
	return _ChequeBook.Contract.TotalHardDeposit(&_ChequeBook.CallOpts)
}

// TotalHardDeposit is a free data retrieval call binding the contract method 0xe0bcf13a.
//
// Solidity: function totalHardDeposit() view returns(uint256)
func (_ChequeBook *ChequeBookCallerSession) TotalHardDeposit() (*big.Int, error) {
	return _ChequeBook.Contract.TotalHardDeposit(&_ChequeBook.CallOpts)
}

// TotalPaidOut is a free data retrieval call binding the contract method 0x1357e1dc.
//
// Solidity: function totalPaidOut() view returns(uint256)
func (_ChequeBook *ChequeBookCaller) TotalPaidOut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChequeBook.contract.Call(opts, out, "totalPaidOut")
	return *ret0, err
}

// TotalPaidOut is a free data retrieval call binding the contract method 0x1357e1dc.
//
// Solidity: function totalPaidOut() view returns(uint256)
func (_ChequeBook *ChequeBookSession) TotalPaidOut() (*big.Int, error) {
	return _ChequeBook.Contract.TotalPaidOut(&_ChequeBook.CallOpts)
}

// TotalPaidOut is a free data retrieval call binding the contract method 0x1357e1dc.
//
// Solidity: function totalPaidOut() view returns(uint256)
func (_ChequeBook *ChequeBookCallerSession) TotalPaidOut() (*big.Int, error) {
	return _ChequeBook.Contract.TotalPaidOut(&_ChequeBook.CallOpts)
}

// Cashout is a paid mutator transaction binding the contract method 0xb7b172b3.
//
// Solidity: function cashout(address _recipient, uint256 _amount) returns()
func (_ChequeBook *ChequeBookTransactor) Cashout(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.contract.Transact(opts, "cashout", _recipient, _amount)
}

// Cashout is a paid mutator transaction binding the contract method 0xb7b172b3.
//
// Solidity: function cashout(address _recipient, uint256 _amount) returns()
func (_ChequeBook *ChequeBookSession) Cashout(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.Contract.Cashout(&_ChequeBook.TransactOpts, _recipient, _amount)
}

// Cashout is a paid mutator transaction binding the contract method 0xb7b172b3.
//
// Solidity: function cashout(address _recipient, uint256 _amount) returns()
func (_ChequeBook *ChequeBookTransactorSession) Cashout(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.Contract.Cashout(&_ChequeBook.TransactOpts, _recipient, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_ChequeBook *ChequeBookTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_ChequeBook *ChequeBookSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.Contract.Deposit(&_ChequeBook.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_ChequeBook *ChequeBookTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.Contract.Deposit(&_ChequeBook.TransactOpts, _amount)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _issuer, address _token) returns()
func (_ChequeBook *ChequeBookTransactor) Init(opts *bind.TransactOpts, _issuer common.Address, _token common.Address) (*types.Transaction, error) {
	return _ChequeBook.contract.Transact(opts, "init", _issuer, _token)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _issuer, address _token) returns()
func (_ChequeBook *ChequeBookSession) Init(_issuer common.Address, _token common.Address) (*types.Transaction, error) {
	return _ChequeBook.Contract.Init(&_ChequeBook.TransactOpts, _issuer, _token)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _issuer, address _token) returns()
func (_ChequeBook *ChequeBookTransactorSession) Init(_issuer common.Address, _token common.Address) (*types.Transaction, error) {
	return _ChequeBook.Contract.Init(&_ChequeBook.TransactOpts, _issuer, _token)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(address _recipient, uint256 _amount) returns()
func (_ChequeBook *ChequeBookTransactor) Pay(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.contract.Transact(opts, "pay", _recipient, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(address _recipient, uint256 _amount) returns()
func (_ChequeBook *ChequeBookSession) Pay(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.Contract.Pay(&_ChequeBook.TransactOpts, _recipient, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(address _recipient, uint256 _amount) returns()
func (_ChequeBook *ChequeBookTransactorSession) Pay(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.Contract.Pay(&_ChequeBook.TransactOpts, _recipient, _amount)
}

// PayHardDeposit is a paid mutator transaction binding the contract method 0xf6407d0c.
//
// Solidity: function payHardDeposit(address _recipient, uint256 _amount) returns()
func (_ChequeBook *ChequeBookTransactor) PayHardDeposit(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.contract.Transact(opts, "payHardDeposit", _recipient, _amount)
}

// PayHardDeposit is a paid mutator transaction binding the contract method 0xf6407d0c.
//
// Solidity: function payHardDeposit(address _recipient, uint256 _amount) returns()
func (_ChequeBook *ChequeBookSession) PayHardDeposit(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.Contract.PayHardDeposit(&_ChequeBook.TransactOpts, _recipient, _amount)
}

// PayHardDeposit is a paid mutator transaction binding the contract method 0xf6407d0c.
//
// Solidity: function payHardDeposit(address _recipient, uint256 _amount) returns()
func (_ChequeBook *ChequeBookTransactorSession) PayHardDeposit(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChequeBook.Contract.PayHardDeposit(&_ChequeBook.TransactOpts, _recipient, _amount)
}

// ChequeBookCashoutedIterator is returned from FilterCashouted and is used to iterate over the raw logs and unpacked data for Cashouted events raised by the ChequeBook contract.
type ChequeBookCashoutedIterator struct {
	Event *ChequeBookCashouted // Event containing the contract specifics and raw log

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
func (it *ChequeBookCashoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChequeBookCashouted)
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
		it.Event = new(ChequeBookCashouted)
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
func (it *ChequeBookCashoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChequeBookCashoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChequeBookCashouted represents a Cashouted event raised by the ChequeBook contract.
type ChequeBookCashouted struct {
	Beneficiary common.Address
	Recipient   common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCashouted is a free log retrieval operation binding the contract event 0x92fec97bff9ddc18bbb7cf738dd42818ad2528ca913ee47fbe149a5cc822ad21.
//
// Solidity: event Cashouted(address indexed beneficiary, address recipient, uint256 amount)
func (_ChequeBook *ChequeBookFilterer) FilterCashouted(opts *bind.FilterOpts, beneficiary []common.Address) (*ChequeBookCashoutedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ChequeBook.contract.FilterLogs(opts, "Cashouted", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ChequeBookCashoutedIterator{contract: _ChequeBook.contract, event: "Cashouted", logs: logs, sub: sub}, nil
}

// WatchCashouted is a free log subscription operation binding the contract event 0x92fec97bff9ddc18bbb7cf738dd42818ad2528ca913ee47fbe149a5cc822ad21.
//
// Solidity: event Cashouted(address indexed beneficiary, address recipient, uint256 amount)
func (_ChequeBook *ChequeBookFilterer) WatchCashouted(opts *bind.WatchOpts, sink chan<- *ChequeBookCashouted, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ChequeBook.contract.WatchLogs(opts, "Cashouted", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChequeBookCashouted)
				if err := _ChequeBook.contract.UnpackLog(event, "Cashouted", log); err != nil {
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

// ParseCashouted is a log parse operation binding the contract event 0x92fec97bff9ddc18bbb7cf738dd42818ad2528ca913ee47fbe149a5cc822ad21.
//
// Solidity: event Cashouted(address indexed beneficiary, address recipient, uint256 amount)
func (_ChequeBook *ChequeBookFilterer) ParseCashouted(log types.Log) (*ChequeBookCashouted, error) {
	event := new(ChequeBookCashouted)
	if err := _ChequeBook.contract.UnpackLog(event, "Cashouted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChequeBookHardDepositAmountChangedIterator is returned from FilterHardDepositAmountChanged and is used to iterate over the raw logs and unpacked data for HardDepositAmountChanged events raised by the ChequeBook contract.
type ChequeBookHardDepositAmountChangedIterator struct {
	Event *ChequeBookHardDepositAmountChanged // Event containing the contract specifics and raw log

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
func (it *ChequeBookHardDepositAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChequeBookHardDepositAmountChanged)
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
		it.Event = new(ChequeBookHardDepositAmountChanged)
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
func (it *ChequeBookHardDepositAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChequeBookHardDepositAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChequeBookHardDepositAmountChanged represents a HardDepositAmountChanged event raised by the ChequeBook contract.
type ChequeBookHardDepositAmountChanged struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHardDepositAmountChanged is a free log retrieval operation binding the contract event 0x2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad.
//
// Solidity: event HardDepositAmountChanged(address indexed beneficiary, uint256 amount)
func (_ChequeBook *ChequeBookFilterer) FilterHardDepositAmountChanged(opts *bind.FilterOpts, beneficiary []common.Address) (*ChequeBookHardDepositAmountChangedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ChequeBook.contract.FilterLogs(opts, "HardDepositAmountChanged", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ChequeBookHardDepositAmountChangedIterator{contract: _ChequeBook.contract, event: "HardDepositAmountChanged", logs: logs, sub: sub}, nil
}

// WatchHardDepositAmountChanged is a free log subscription operation binding the contract event 0x2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad.
//
// Solidity: event HardDepositAmountChanged(address indexed beneficiary, uint256 amount)
func (_ChequeBook *ChequeBookFilterer) WatchHardDepositAmountChanged(opts *bind.WatchOpts, sink chan<- *ChequeBookHardDepositAmountChanged, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ChequeBook.contract.WatchLogs(opts, "HardDepositAmountChanged", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChequeBookHardDepositAmountChanged)
				if err := _ChequeBook.contract.UnpackLog(event, "HardDepositAmountChanged", log); err != nil {
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

// ParseHardDepositAmountChanged is a log parse operation binding the contract event 0x2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad.
//
// Solidity: event HardDepositAmountChanged(address indexed beneficiary, uint256 amount)
func (_ChequeBook *ChequeBookFilterer) ParseHardDepositAmountChanged(log types.Log) (*ChequeBookHardDepositAmountChanged, error) {
	event := new(ChequeBookHardDepositAmountChanged)
	if err := _ChequeBook.contract.UnpackLog(event, "HardDepositAmountChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChequeBookFactoryABI is the input ABI used to generate the binding from.
const ChequeBookFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ERC20Address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ChequeBookDeployed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"deployChequeBook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployedContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChequeBookFactoryBin is the compiled bytecode used for deploying new contracts.
var ChequeBookFactoryBin = "0x608060405234801561001057600080fd5b506040516111d13803806111d183398101604081905261002f91610129565b600180546001600160a01b0319166001600160a01b0383161790556040516000906100599061011c565b604051809103906000f080158015610075573d6000803e3d6000fd5b506040517ff09a40160000000000000000000000000000000000000000000000000000000081529091506001600160a01b0382169063f09a4016906100c290600190600090600401610157565b600060405180830381600087803b1580156100dc57600080fd5b505af11580156100f0573d6000803e3d6000fd5b5050600280546001600160a01b0319166001600160a01b03949094169390931790925550610171915050565b610d328061049f83390190565b60006020828403121561013a578081fd5b81516001600160a01b0381168114610150578182fd5b9392505050565b6001600160a01b0392831681529116602082015260400190565b61031f806101806000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634e36003414610051578063a6021ace1461007a578063b35fbb2714610082578063ee97f7f31461008a575b600080fd5b61006461005f36600461026c565b610092565b60405161007191906102a2565b60405180910390f35b61006461019a565b6100646101a9565b6100646101b8565b60025460405160009182916100d6916001600160a01b0316906100bb90339087906020016102b6565b604051602081830303815290604052805190602001206101c7565b60015460405163784d200b60e11b81529192506001600160a01b038084169263f09a40169261010c9289929116906004016102cf565b600060405180830381600087803b15801561012657600080fd5b505af115801561013a573d6000803e3d6000fd5b5050600080546001600160a01b0319166001600160a01b03851617905550506040517f7a3ceb4fc57b5e88d7a21720f0b69536d1d1c1084b320499456655ac02e81897906101899083906102a2565b60405180910390a190505b92915050565b6001546001600160a01b031681565b6000546001600160a01b031681565b6002546001600160a01b031681565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528360601b60148201526e5af43d82803e903d91602b57fd5bf360881b6028820152826037826000f59150506001600160a01b038116610194576040805162461bcd60e51b815260206004820152601760248201527f455243313136373a2063726561746532206661696c6564000000000000000000604482015290519081900360640190fd5b6000806040838503121561027e578182fd5b82356001600160a01b0381168114610294578283fd5b946020939093013593505050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b039283168152911660208201526040019056fea2646970667358221220a12cf80e83ba3e55ac5fef44b4dd70b18704d45d181763230a14b99444560fcb64736f6c63430007060033608060405234801561001057600080fd5b50610d12806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063b6b55f2511610097578063e0bcf13a11610066578063e0bcf13a146101b9578063f09a4016146101c1578063f6407d0c146101d4578063fc0c546a146101e7576100f5565b8063b6b55f2514610176578063b7b172b31461018b578063b7ec1a331461019e578063c4076876146101a6576100f5565b80635bf58ec7116100d35780635bf58ec71461014057806381f03fcb14610153578063959d15ca14610166578063b69ef8a81461016e576100f5565b80631357e1dc146100fa5780631d143848146101185780635664865a1461012d575b600080fd5b6101026101ef565b60405161010f9190610cd3565b60405180910390f35b6101206101f5565b60405161010f9190610ab2565b61010261013b3660046109ed565b610204565b61010261014e3660046109ed565b610216565b6101026101613660046109ed565b610228565b61010261023a565b610102610240565b610189610184366004610a82565b6102c6565b005b610189610199366004610a39565b61037e565b61010261054b565b6101896101b4366004610a39565b610561565b610102610611565b6101896101cf366004610a07565b610617565b6101896101e2366004610a39565b610694565b6101206106f0565b60045481565b6007546001600160a01b031681565b60026020526000908152604090205481565b60016020526000908152604090205481565b60036020526000908152604090205481565b60065481565b600080546040516370a0823160e01b81526001600160a01b03909116906370a0823190610271903090600401610ab2565b60206040518083038186803b15801561028957600080fd5b505afa15801561029d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c19190610a9a565b905090565b6000546040516323b872dd60e01b81526001600160a01b03909116906323b872dd906102fa90339030908690600401610ac6565b602060405180830381600087803b15801561031457600080fd5b505af1158015610328573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034c9190610a62565b6103715760405162461bcd60e51b815260040161036890610c75565b60405180910390fd5b61037b33826106ff565b50565b6103883382610782565b336000908152600360205260409020546103a290826107f4565b336000908152600360205260409020556004546103bf90826107f4565b6004908155600054604051633950935160e01b81526001600160a01b03909116916339509351916103f4913091869101610aea565b602060405180830381600087803b15801561040e57600080fd5b505af1158015610422573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104469190610a62565b6104625760405162461bcd60e51b815260040161036890610bfd565b6000546040516323b872dd60e01b81526001600160a01b03909116906323b872dd9061049690309086908690600401610ac6565b602060405180830381600087803b1580156104b057600080fd5b505af11580156104c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e89190610a62565b6105045760405162461bcd60e51b815260040161036890610bc6565b336001600160a01b03167f92fec97bff9ddc18bbb7cf738dd42818ad2528ca913ee47fbe149a5cc822ad21838360405161053f929190610aea565b60405180910390a25050565b60006102c160055461055b610240565b90610855565b6000546040516323b872dd60e01b81526001600160a01b03909116906323b872dd9061059590339030908690600401610ac6565b602060405180830381600087803b1580156105af57600080fd5b505af11580156105c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e79190610a62565b6106035760405162461bcd60e51b815260040161036890610c75565b61060d82826108b2565b5050565b60055481565b6001600160a01b03821661063d5760405162461bcd60e51b815260040161036890610b71565b6007546001600160a01b0316156106665760405162461bcd60e51b815260040161036890610b99565b600780546001600160a01b039384166001600160a01b03199182161790915560008054929093169116179055565b61069e338261092e565b6106a882826108b2565b6001600160a01b038216600081815260026020526040908190205490517f2506c43272ded05d095b91dbba876e66e46888157d3e078db5691496e96c5fad9161053f91610cd3565b6000546001600160a01b031681565b610707610240565b60055461071490836107f4565b11156107325760405162461bcd60e51b815260040161036890610b3a565b6001600160a01b03821660009081526001602052604090205461075590826107f4565b6001600160a01b03831660009081526001602052604090205560055461077b90826107f4565b6005555050565b8060065410156107a45760405162461bcd60e51b815260040161036890610c9e565b6001600160a01b0382166000908152600260205260409020546107c79082610855565b6001600160a01b0383166000908152600260205260409020556006546107ed9082610855565b6006555050565b60008282018381101561084e576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000828211156108ac576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6108ba610240565b6006546108c790836107f4565b11156108e55760405162461bcd60e51b815260040161036890610b3a565b6001600160a01b03821660009081526002602052604090205461090890826107f4565b6001600160a01b0383166000908152600260205260409020556006546107ed90826107f4565b6001600160a01b0382166000908152600160205260409020548111156109665760405162461bcd60e51b815260040161036890610b03565b8060055410156109885760405162461bcd60e51b815260040161036890610c34565b6001600160a01b0382166000908152600160205260409020546109ab9082610855565b6001600160a01b03831660009081526001602052604090205560055461077b9082610855565b80356001600160a01b03811681146109e857600080fd5b919050565b6000602082840312156109fe578081fd5b61084e826109d1565b60008060408385031215610a19578081fd5b610a22836109d1565b9150610a30602084016109d1565b90509250929050565b60008060408385031215610a4b578182fd5b610a54836109d1565b946020939093013593505050565b600060208284031215610a73578081fd5b8151801515811461084e578182fd5b600060208284031215610a93578081fd5b5035919050565b600060208284031215610aab578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b6020808252601c908201527f616d6f756e742065786365656473206f776e65722062616c616e636500000000604082015260600190565b6020808252601d908201527f616d6f756e74206578636565647320676c6f62616c2062616c616e6365000000604082015260600190565b6020808252600e908201526d34b73b30b634b21034b9b9bab2b960911b604082015260600190565b602080825260139082015272185b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604082015260600190565b60208082526017908201527f636173686f7574207472616e73666572206661696c6564000000000000000000604082015260600190565b60208082526019908201527f696e63726561736520616c6c6f77616e6365206661696c656400000000000000604082015260600190565b60208082526021908201527f616d6f756e74206578636565647320746f74616c2068617264206465706f73696040820152601d60fa1b606082015260800190565b6020808252600f908201526e1d1c985b9cd9995c8819985a5b1959608a1b604082015260600190565b6020808252818101527f616d6f756e74206578636565647320746f74616c2062656e6566696369617279604082015260600190565b9081526020019056fea264697066735822122048c69c48ae8454ebfc51c0054615611553e8c5923fcb284ec676f8f7ebc36df164736f6c63430007060033"

// DeployChequeBookFactory deploys a new Ethereum contract, binding an instance of ChequeBookFactory to it.
func DeployChequeBookFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _ERC20Address common.Address) (common.Address, *types.Transaction, *ChequeBookFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ChequeBookFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChequeBookFactoryBin), backend, _ERC20Address)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChequeBookFactory{ChequeBookFactoryCaller: ChequeBookFactoryCaller{contract: contract}, ChequeBookFactoryTransactor: ChequeBookFactoryTransactor{contract: contract}, ChequeBookFactoryFilterer: ChequeBookFactoryFilterer{contract: contract}}, nil
}

// ChequeBookFactory is an auto generated Go binding around an Ethereum contract.
type ChequeBookFactory struct {
	ChequeBookFactoryCaller     // Read-only binding to the contract
	ChequeBookFactoryTransactor // Write-only binding to the contract
	ChequeBookFactoryFilterer   // Log filterer for contract events
}

// ChequeBookFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChequeBookFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChequeBookFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChequeBookFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChequeBookFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChequeBookFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChequeBookFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChequeBookFactorySession struct {
	Contract     *ChequeBookFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ChequeBookFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChequeBookFactoryCallerSession struct {
	Contract *ChequeBookFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ChequeBookFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChequeBookFactoryTransactorSession struct {
	Contract     *ChequeBookFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ChequeBookFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChequeBookFactoryRaw struct {
	Contract *ChequeBookFactory // Generic contract binding to access the raw methods on
}

// ChequeBookFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChequeBookFactoryCallerRaw struct {
	Contract *ChequeBookFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ChequeBookFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChequeBookFactoryTransactorRaw struct {
	Contract *ChequeBookFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChequeBookFactory creates a new instance of ChequeBookFactory, bound to a specific deployed contract.
func NewChequeBookFactory(address common.Address, backend bind.ContractBackend) (*ChequeBookFactory, error) {
	contract, err := bindChequeBookFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChequeBookFactory{ChequeBookFactoryCaller: ChequeBookFactoryCaller{contract: contract}, ChequeBookFactoryTransactor: ChequeBookFactoryTransactor{contract: contract}, ChequeBookFactoryFilterer: ChequeBookFactoryFilterer{contract: contract}}, nil
}

// NewChequeBookFactoryCaller creates a new read-only instance of ChequeBookFactory, bound to a specific deployed contract.
func NewChequeBookFactoryCaller(address common.Address, caller bind.ContractCaller) (*ChequeBookFactoryCaller, error) {
	contract, err := bindChequeBookFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChequeBookFactoryCaller{contract: contract}, nil
}

// NewChequeBookFactoryTransactor creates a new write-only instance of ChequeBookFactory, bound to a specific deployed contract.
func NewChequeBookFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ChequeBookFactoryTransactor, error) {
	contract, err := bindChequeBookFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChequeBookFactoryTransactor{contract: contract}, nil
}

// NewChequeBookFactoryFilterer creates a new log filterer instance of ChequeBookFactory, bound to a specific deployed contract.
func NewChequeBookFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ChequeBookFactoryFilterer, error) {
	contract, err := bindChequeBookFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChequeBookFactoryFilterer{contract: contract}, nil
}

// bindChequeBookFactory binds a generic wrapper to an already deployed contract.
func bindChequeBookFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChequeBookFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChequeBookFactory *ChequeBookFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChequeBookFactory.Contract.ChequeBookFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChequeBookFactory *ChequeBookFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChequeBookFactory.Contract.ChequeBookFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChequeBookFactory *ChequeBookFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChequeBookFactory.Contract.ChequeBookFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChequeBookFactory *ChequeBookFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChequeBookFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChequeBookFactory *ChequeBookFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChequeBookFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChequeBookFactory *ChequeBookFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChequeBookFactory.Contract.contract.Transact(opts, method, params...)
}

// ERC20Address is a free data retrieval call binding the contract method 0xa6021ace.
//
// Solidity: function ERC20Address() view returns(address)
func (_ChequeBookFactory *ChequeBookFactoryCaller) ERC20Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChequeBookFactory.contract.Call(opts, out, "ERC20Address")
	return *ret0, err
}

// ERC20Address is a free data retrieval call binding the contract method 0xa6021ace.
//
// Solidity: function ERC20Address() view returns(address)
func (_ChequeBookFactory *ChequeBookFactorySession) ERC20Address() (common.Address, error) {
	return _ChequeBookFactory.Contract.ERC20Address(&_ChequeBookFactory.CallOpts)
}

// ERC20Address is a free data retrieval call binding the contract method 0xa6021ace.
//
// Solidity: function ERC20Address() view returns(address)
func (_ChequeBookFactory *ChequeBookFactoryCallerSession) ERC20Address() (common.Address, error) {
	return _ChequeBookFactory.Contract.ERC20Address(&_ChequeBookFactory.CallOpts)
}

// DeployedContractAddress is a free data retrieval call binding the contract method 0xb35fbb27.
//
// Solidity: function deployedContractAddress() view returns(address)
func (_ChequeBookFactory *ChequeBookFactoryCaller) DeployedContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChequeBookFactory.contract.Call(opts, out, "deployedContractAddress")
	return *ret0, err
}

// DeployedContractAddress is a free data retrieval call binding the contract method 0xb35fbb27.
//
// Solidity: function deployedContractAddress() view returns(address)
func (_ChequeBookFactory *ChequeBookFactorySession) DeployedContractAddress() (common.Address, error) {
	return _ChequeBookFactory.Contract.DeployedContractAddress(&_ChequeBookFactory.CallOpts)
}

// DeployedContractAddress is a free data retrieval call binding the contract method 0xb35fbb27.
//
// Solidity: function deployedContractAddress() view returns(address)
func (_ChequeBookFactory *ChequeBookFactoryCallerSession) DeployedContractAddress() (common.Address, error) {
	return _ChequeBookFactory.Contract.DeployedContractAddress(&_ChequeBookFactory.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ChequeBookFactory *ChequeBookFactoryCaller) Master(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChequeBookFactory.contract.Call(opts, out, "master")
	return *ret0, err
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ChequeBookFactory *ChequeBookFactorySession) Master() (common.Address, error) {
	return _ChequeBookFactory.Contract.Master(&_ChequeBookFactory.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ChequeBookFactory *ChequeBookFactoryCallerSession) Master() (common.Address, error) {
	return _ChequeBookFactory.Contract.Master(&_ChequeBookFactory.CallOpts)
}

// DeployChequeBook is a paid mutator transaction binding the contract method 0x4e360034.
//
// Solidity: function deployChequeBook(address issuer, bytes32 salt) returns(address)
func (_ChequeBookFactory *ChequeBookFactoryTransactor) DeployChequeBook(opts *bind.TransactOpts, issuer common.Address, salt [32]byte) (*types.Transaction, error) {
	return _ChequeBookFactory.contract.Transact(opts, "deployChequeBook", issuer, salt)
}

// DeployChequeBook is a paid mutator transaction binding the contract method 0x4e360034.
//
// Solidity: function deployChequeBook(address issuer, bytes32 salt) returns(address)
func (_ChequeBookFactory *ChequeBookFactorySession) DeployChequeBook(issuer common.Address, salt [32]byte) (*types.Transaction, error) {
	return _ChequeBookFactory.Contract.DeployChequeBook(&_ChequeBookFactory.TransactOpts, issuer, salt)
}

// DeployChequeBook is a paid mutator transaction binding the contract method 0x4e360034.
//
// Solidity: function deployChequeBook(address issuer, bytes32 salt) returns(address)
func (_ChequeBookFactory *ChequeBookFactoryTransactorSession) DeployChequeBook(issuer common.Address, salt [32]byte) (*types.Transaction, error) {
	return _ChequeBookFactory.Contract.DeployChequeBook(&_ChequeBookFactory.TransactOpts, issuer, salt)
}

// ChequeBookFactoryChequeBookDeployedIterator is returned from FilterChequeBookDeployed and is used to iterate over the raw logs and unpacked data for ChequeBookDeployed events raised by the ChequeBookFactory contract.
type ChequeBookFactoryChequeBookDeployedIterator struct {
	Event *ChequeBookFactoryChequeBookDeployed // Event containing the contract specifics and raw log

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
func (it *ChequeBookFactoryChequeBookDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChequeBookFactoryChequeBookDeployed)
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
		it.Event = new(ChequeBookFactoryChequeBookDeployed)
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
func (it *ChequeBookFactoryChequeBookDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChequeBookFactoryChequeBookDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChequeBookFactoryChequeBookDeployed represents a ChequeBookDeployed event raised by the ChequeBookFactory contract.
type ChequeBookFactoryChequeBookDeployed struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChequeBookDeployed is a free log retrieval operation binding the contract event 0x7a3ceb4fc57b5e88d7a21720f0b69536d1d1c1084b320499456655ac02e81897.
//
// Solidity: event ChequeBookDeployed(address contractAddress)
func (_ChequeBookFactory *ChequeBookFactoryFilterer) FilterChequeBookDeployed(opts *bind.FilterOpts) (*ChequeBookFactoryChequeBookDeployedIterator, error) {

	logs, sub, err := _ChequeBookFactory.contract.FilterLogs(opts, "ChequeBookDeployed")
	if err != nil {
		return nil, err
	}
	return &ChequeBookFactoryChequeBookDeployedIterator{contract: _ChequeBookFactory.contract, event: "ChequeBookDeployed", logs: logs, sub: sub}, nil
}

// WatchChequeBookDeployed is a free log subscription operation binding the contract event 0x7a3ceb4fc57b5e88d7a21720f0b69536d1d1c1084b320499456655ac02e81897.
//
// Solidity: event ChequeBookDeployed(address contractAddress)
func (_ChequeBookFactory *ChequeBookFactoryFilterer) WatchChequeBookDeployed(opts *bind.WatchOpts, sink chan<- *ChequeBookFactoryChequeBookDeployed) (event.Subscription, error) {

	logs, sub, err := _ChequeBookFactory.contract.WatchLogs(opts, "ChequeBookDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChequeBookFactoryChequeBookDeployed)
				if err := _ChequeBookFactory.contract.UnpackLog(event, "ChequeBookDeployed", log); err != nil {
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

// ParseChequeBookDeployed is a log parse operation binding the contract event 0x7a3ceb4fc57b5e88d7a21720f0b69536d1d1c1084b320499456655ac02e81897.
//
// Solidity: event ChequeBookDeployed(address contractAddress)
func (_ChequeBookFactory *ChequeBookFactoryFilterer) ParseChequeBookDeployed(log types.Log) (*ChequeBookFactoryChequeBookDeployed, error) {
	event := new(ChequeBookFactoryChequeBookDeployed)
	if err := _ChequeBookFactory.contract.UnpackLog(event, "ChequeBookDeployed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ClonesABI is the input ABI used to generate the binding from.
const ClonesABI = "[]"

// ClonesBin is the compiled bytecode used for deploying new contracts.
var ClonesBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ffcd0cce12da8866c3fe00d55b70c6376e63149ecbdfd94498e93496e17523a564736f6c63430007060033"

// DeployClones deploys a new Ethereum contract, binding an instance of Clones to it.
func DeployClones(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Clones, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClonesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// Clones is an auto generated Go binding around an Ethereum contract.
type Clones struct {
	ClonesCaller     // Read-only binding to the contract
	ClonesTransactor // Write-only binding to the contract
	ClonesFilterer   // Log filterer for contract events
}

// ClonesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClonesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClonesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClonesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClonesSession struct {
	Contract     *Clones           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClonesCallerSession struct {
	Contract *ClonesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClonesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClonesTransactorSession struct {
	Contract     *ClonesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClonesRaw struct {
	Contract *Clones // Generic contract binding to access the raw methods on
}

// ClonesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClonesCallerRaw struct {
	Contract *ClonesCaller // Generic read-only contract binding to access the raw methods on
}

// ClonesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClonesTransactorRaw struct {
	Contract *ClonesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClones creates a new instance of Clones, bound to a specific deployed contract.
func NewClones(address common.Address, backend bind.ContractBackend) (*Clones, error) {
	contract, err := bindClones(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// NewClonesCaller creates a new read-only instance of Clones, bound to a specific deployed contract.
func NewClonesCaller(address common.Address, caller bind.ContractCaller) (*ClonesCaller, error) {
	contract, err := bindClones(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesCaller{contract: contract}, nil
}

// NewClonesTransactor creates a new write-only instance of Clones, bound to a specific deployed contract.
func NewClonesTransactor(address common.Address, transactor bind.ContractTransactor) (*ClonesTransactor, error) {
	contract, err := bindClones(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesTransactor{contract: contract}, nil
}

// NewClonesFilterer creates a new log filterer instance of Clones, bound to a specific deployed contract.
func NewClonesFilterer(address common.Address, filterer bind.ContractFilterer) (*ClonesFilterer, error) {
	contract, err := bindClones(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClonesFilterer{contract: contract}, nil
}

// bindClones binds a generic wrapper to an already deployed contract.
func bindClones(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.ClonesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x60806040523480156200001157600080fd5b5060405162000c8e38038062000c8e833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405250508251620001b491506003906020850190620001e0565b508051620001ca906004906020840190620001e0565b50506005805460ff19166012179055506200028c565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000218576000855562000263565b82601f106200023357805160ff191683800117855562000263565b8280016001018555821562000263579182015b828111156200026357825182559160200191906001019062000246565b506200027192915062000275565b5090565b5b8082111562000271576000815560010162000276565b6109f2806200029c6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063395093511161007157806339509351146101d957806370a082311461020557806395d89b411461022b578063a457c2d714610233578063a9059cbb1461025f578063dd62ed3e1461028b576100a9565b806306fdde03146100ae578063095ea7b31461012b57806318160ddd1461016b57806323b872dd14610185578063313ce567146101bb575b600080fd5b6100b66102b9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f05781810151838201526020016100d8565b50505050905090810190601f16801561011d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101576004803603604081101561014157600080fd5b506001600160a01b03813516906020013561034f565b604080519115158252519081900360200190f35b61017361036c565b60408051918252519081900360200190f35b6101576004803603606081101561019b57600080fd5b506001600160a01b03813581169160208101359091169060400135610372565b6101c36103f9565b6040805160ff9092168252519081900360200190f35b610157600480360360408110156101ef57600080fd5b506001600160a01b038135169060200135610402565b6101736004803603602081101561021b57600080fd5b50356001600160a01b0316610450565b6100b661046b565b6101576004803603604081101561024957600080fd5b506001600160a01b0381351690602001356104cc565b6101576004803603604081101561027557600080fd5b506001600160a01b038135169060200135610534565b610173600480360360408110156102a157600080fd5b506001600160a01b0381358116916020013516610548565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103455780601f1061031a57610100808354040283529160200191610345565b820191906000526020600020905b81548152906001019060200180831161032857829003601f168201915b5050505050905090565b600061036361035c610573565b8484610577565b50600192915050565b60025490565b600061037f848484610663565b6103ef8461038b610573565b6103ea85604051806060016040528060288152602001610927602891396001600160a01b038a166000908152600160205260408120906103c9610573565b6001600160a01b0316815260208101919091526040016000205491906107be565b610577565b5060019392505050565b60055460ff1690565b600061036361040f610573565b846103ea8560016000610420610573565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490610855565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103455780601f1061031a57610100808354040283529160200191610345565b60006103636104d9610573565b846103ea856040518060600160405280602581526020016109986025913960016000610503610573565b6001600160a01b03908116825260208083019390935260409182016000908120918d168152925290205491906107be565b6000610363610541610573565b8484610663565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166105bc5760405162461bcd60e51b81526004018080602001828103825260248152602001806109746024913960400191505060405180910390fd5b6001600160a01b0382166106015760405162461bcd60e51b81526004018080602001828103825260228152602001806108df6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166106a85760405162461bcd60e51b815260040180806020018281038252602581526020018061094f6025913960400191505060405180910390fd5b6001600160a01b0382166106ed5760405162461bcd60e51b81526004018080602001828103825260238152602001806108bc6023913960400191505060405180910390fd5b6106f88383836108b6565b61073581604051806060016040528060268152602001610901602691396001600160a01b03861660009081526020819052604090205491906107be565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546107649082610855565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561084d5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156108125781810151838201526020016107fa565b50505050905090810190601f16801561083f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000828201838110156108af576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220c3e3a4dc4e2c07bd355c332282b075a2f45cf38b85dd09afa790a3339189d95364736f6c63430007060033"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IChequeBookABI is the input ABI used to generate the binding from.
const IChequeBookABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"cashout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payHardDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChequeBook is an auto generated Go binding around an Ethereum contract.
type IChequeBook struct {
	IChequeBookCaller     // Read-only binding to the contract
	IChequeBookTransactor // Write-only binding to the contract
	IChequeBookFilterer   // Log filterer for contract events
}

// IChequeBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChequeBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChequeBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChequeBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChequeBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChequeBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChequeBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChequeBookSession struct {
	Contract     *IChequeBook      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IChequeBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChequeBookCallerSession struct {
	Contract *IChequeBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IChequeBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChequeBookTransactorSession struct {
	Contract     *IChequeBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IChequeBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChequeBookRaw struct {
	Contract *IChequeBook // Generic contract binding to access the raw methods on
}

// IChequeBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChequeBookCallerRaw struct {
	Contract *IChequeBookCaller // Generic read-only contract binding to access the raw methods on
}

// IChequeBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChequeBookTransactorRaw struct {
	Contract *IChequeBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChequeBook creates a new instance of IChequeBook, bound to a specific deployed contract.
func NewIChequeBook(address common.Address, backend bind.ContractBackend) (*IChequeBook, error) {
	contract, err := bindIChequeBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChequeBook{IChequeBookCaller: IChequeBookCaller{contract: contract}, IChequeBookTransactor: IChequeBookTransactor{contract: contract}, IChequeBookFilterer: IChequeBookFilterer{contract: contract}}, nil
}

// NewIChequeBookCaller creates a new read-only instance of IChequeBook, bound to a specific deployed contract.
func NewIChequeBookCaller(address common.Address, caller bind.ContractCaller) (*IChequeBookCaller, error) {
	contract, err := bindIChequeBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChequeBookCaller{contract: contract}, nil
}

// NewIChequeBookTransactor creates a new write-only instance of IChequeBook, bound to a specific deployed contract.
func NewIChequeBookTransactor(address common.Address, transactor bind.ContractTransactor) (*IChequeBookTransactor, error) {
	contract, err := bindIChequeBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChequeBookTransactor{contract: contract}, nil
}

// NewIChequeBookFilterer creates a new log filterer instance of IChequeBook, bound to a specific deployed contract.
func NewIChequeBookFilterer(address common.Address, filterer bind.ContractFilterer) (*IChequeBookFilterer, error) {
	contract, err := bindIChequeBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChequeBookFilterer{contract: contract}, nil
}

// bindIChequeBook binds a generic wrapper to an already deployed contract.
func bindIChequeBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChequeBookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChequeBook *IChequeBookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChequeBook.Contract.IChequeBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChequeBook *IChequeBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChequeBook.Contract.IChequeBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChequeBook *IChequeBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChequeBook.Contract.IChequeBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChequeBook *IChequeBookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChequeBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChequeBook *IChequeBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChequeBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChequeBook *IChequeBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChequeBook.Contract.contract.Transact(opts, method, params...)
}

// Cashout is a paid mutator transaction binding the contract method 0xb7b172b3.
//
// Solidity: function cashout(address _recipient, uint256 _amount) returns()
func (_IChequeBook *IChequeBookTransactor) Cashout(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.contract.Transact(opts, "cashout", _recipient, _amount)
}

// Cashout is a paid mutator transaction binding the contract method 0xb7b172b3.
//
// Solidity: function cashout(address _recipient, uint256 _amount) returns()
func (_IChequeBook *IChequeBookSession) Cashout(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.Contract.Cashout(&_IChequeBook.TransactOpts, _recipient, _amount)
}

// Cashout is a paid mutator transaction binding the contract method 0xb7b172b3.
//
// Solidity: function cashout(address _recipient, uint256 _amount) returns()
func (_IChequeBook *IChequeBookTransactorSession) Cashout(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.Contract.Cashout(&_IChequeBook.TransactOpts, _recipient, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_IChequeBook *IChequeBookTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_IChequeBook *IChequeBookSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.Contract.Deposit(&_IChequeBook.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_IChequeBook *IChequeBookTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.Contract.Deposit(&_IChequeBook.TransactOpts, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(address _recipient, uint256 _amount) returns()
func (_IChequeBook *IChequeBookTransactor) Pay(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.contract.Transact(opts, "pay", _recipient, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(address _recipient, uint256 _amount) returns()
func (_IChequeBook *IChequeBookSession) Pay(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.Contract.Pay(&_IChequeBook.TransactOpts, _recipient, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(address _recipient, uint256 _amount) returns()
func (_IChequeBook *IChequeBookTransactorSession) Pay(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.Contract.Pay(&_IChequeBook.TransactOpts, _recipient, _amount)
}

// PayHardDeposit is a paid mutator transaction binding the contract method 0xf6407d0c.
//
// Solidity: function payHardDeposit(address _recipient, uint256 _amount) returns()
func (_IChequeBook *IChequeBookTransactor) PayHardDeposit(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.contract.Transact(opts, "payHardDeposit", _recipient, _amount)
}

// PayHardDeposit is a paid mutator transaction binding the contract method 0xf6407d0c.
//
// Solidity: function payHardDeposit(address _recipient, uint256 _amount) returns()
func (_IChequeBook *IChequeBookSession) PayHardDeposit(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.Contract.PayHardDeposit(&_IChequeBook.TransactOpts, _recipient, _amount)
}

// PayHardDeposit is a paid mutator transaction binding the contract method 0xf6407d0c.
//
// Solidity: function payHardDeposit(address _recipient, uint256 _amount) returns()
func (_IChequeBook *IChequeBookTransactorSession) PayHardDeposit(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IChequeBook.Contract.PayHardDeposit(&_IChequeBook.TransactOpts, _recipient, _amount)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
var MathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200d2fa3b9a7102b76f2cf18c39a0c05939114993046d701329167a2b24f16613664736f6c63430007060033"

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d8fad7392ab5d778f028019f4297c5a5dba3653906a836c4ed5dd1e118eee5f464736f6c63430007060033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
