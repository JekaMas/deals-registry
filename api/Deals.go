// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DealsABI is the input ABI used to generate the binding from.
const DealsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"GetDeals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hub\",\"type\":\"address\"},{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_specHash\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_workTime\",\"type\":\"uint256\"}],\"name\":\"OpenDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"GetDealInfo\",\"outputs\":[{\"name\":\"specHach\",\"type\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"hub\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"workTime\",\"type\":\"uint256\"},{\"name\":\"endTIme\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AcceptDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CloseDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetDealsAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DealOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DealAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DealClosed\",\"type\":\"event\"}]"

// DealsBin is the compiled bytecode used for deploying new contracts.
const DealsBin = `0x60606040526000600155341561001457600080fd5b604051602080610c3b833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b610be3806100586000396000f3006060604052361561005c5763ffffffff60e060020a600035041663087ca1d881146100615780633b24475f146100d45780635ad5f6ae146101045780635be20de21461016d578063ca6ad1bf14610185578063e45ea8d31461019d575b600080fd5b341561006c57600080fd5b610080600160a060020a03600435166101c2565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156100c05780820151818401525b6020016100a7565b505050509050019250505060405180910390f35b34156100df57600080fd5b610102600160a060020a0360043581169060243516604435606435608435610246565b005b341561010f57600080fd5b61011a6004356104d4565b604051978852600160a060020a039687166020890152949095166040808801919091526060870193909352608086019190915260a085015260c084019290925260e0830152610100909101905180910390f35b341561017857600080fd5b610102600435610530565b005b341561019057600080fd5b6101026004356105f3565b005b34156101a857600080fd5b6101b0610b22565b60405190815260200160405180910390f35b6101ca610b5a565b6004600083600160a060020a0316600160a060020a0316815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561023957602002820191906000526020600020905b815481526020019060010190808311610225575b505050505090505b919050565b83600160a060020a031633600160a060020a031614151561026657600080fd5b60018054810190556101006040519081016040528084815260200185600160a060020a0316815260200186600160a060020a031681526020018381526020016000815260200182815260200160008152602001600181525060026000600154815260200190815260200160002060008201518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201516007909101555060008054600160a060020a0316906323b872dd908690309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156103e657600080fd5b6102c65a03f115156103f757600080fd5b50505060405180519050151561040c57600080fd5b60015460009081526003602052604090205461042e908363ffffffff610b2916565b60018054600090815260036020908152604080832094909455600160a060020a0388168252600490529190912080549091810161046b8382610b6c565b916000526020600020900160005b6001549091909150555083600160a060020a031685600160a060020a03167f873cb35202fef184c9f8ee23c04e36dc38f3e26fb285224ca574a837be97684860015460405190815260200160405180910390a35b5050505050565b60008181526002602081905260409091208054600182015492820154600383015460048401546005850154600686015460078701549597600160a060020a03908116979516959394929391929091905b50919395975091939597565b6000818152600260208190526040909120015433600160a060020a0390811691161461055b57600080fd5b60008181526002602052604090206007015460011461057957600080fd5b60008181526002602081905260409182902060078101829055426004820181905560058201540160068201556001810154910154600160a060020a03918216929116907f3a38edea6028913403c74ce8433c90eca94f4ca074d318d8cb77be5290ba4f159084905190815260200160405180910390a35b50565b6000818152600260208190526040822060070154141561098e57600082815260026020526040902060060154421115610709576000805483825260026020819052604080842091820154600390920154600160a060020a039384169463a9059cbb94909316929091516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156106a357600080fd5b6102c65a03f115156106b457600080fd5b5050506040518051905015156106c957600080fd5b60008281526002602090815260408083206003908101549252909120546106f59163ffffffff610b4316565b600083815260036020526040902055610926565b60008281526002602052604090206001015433600160a060020a0390811691161461073357600080fd5b6000828152600260205260409020600581015460039091015481151561075557fe5b6000848152600260208190526040808320600481015484549190930154959094044292909203919091029450600160a060020a039283169363a9059cbb9316918591516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156107df57600080fd5b6102c65a03f115156107f057600080fd5b50505060405180519050151561080557600080fd5b600082815260036020526040902054610824908263ffffffff610b4316565b600083815260036020818152604080842094909455825460029091528383206001810154920154600160a060020a039182169463a9059cbb9493909216929086900391516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156108af57600080fd5b6102c65a03f115156108c057600080fd5b5050506040518051905015156108d557600080fd5b60008281526002602090815260408083206003908101549252909120546109049183900363ffffffff610b4316565b6000838152600360209081526040808320939093556002905220426006909101555b600082815260026020819052604091829020600360078201556001810154910154600160a060020a03918216929116907f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb8909085905190815260200160405180910390a3610b1c565b6000828152600260205260409020600701546001141561005c5760008281526002602052604090206001015433600160a060020a039081169116146109d257600080fd5b60008054838252600260205260408083206001810154600390910154600160a060020a039384169463a9059cbb9490921692909190516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610a4f57600080fd5b6102c65a03f11515610a6057600080fd5b505050604051805190501515610a7557600080fd5b6000828152600260209081526040808320600390810154925290912054610aa19163ffffffff610b4316565b6000838152600360208181526040808420949094556002908190529183902060078101919091556001810154910154600160a060020a03918216929116907f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb8909085905190815260200160405180910390a3610b1c565b600080fd5b5b5b5050565b6001545b90565b600082820183811015610b3857fe5b8091505b5092915050565b600082821115610b4f57fe5b508082035b92915050565b60206040519081016040526000815290565b815481835581811511610b9057600083815260209020610b90918101908301610b96565b5b505050565b610b2691905b80821115610bb05760008155600101610b9c565b5090565b905600a165627a7a7230582026bd8fd2c48104e5aa2b6a16796c8e88ee8473f7647926c45c93a9f02100ba980029`

// DeployDeals deploys a new Ethereum contract, binding an instance of Deals to it.
func DeployDeals(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Deals, error) {
	parsed, err := abi.JSON(strings.NewReader(DealsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DealsBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Deals{DealsCaller: DealsCaller{contract: contract}, DealsTransactor: DealsTransactor{contract: contract}}, nil
}

// Deals is an auto generated Go binding around an Ethereum contract.
type Deals struct {
	DealsCaller     // Read-only binding to the contract
	DealsTransactor // Write-only binding to the contract
}

// DealsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DealsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DealsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DealsSession struct {
	Contract     *Deals            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DealsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DealsCallerSession struct {
	Contract *DealsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DealsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DealsTransactorSession struct {
	Contract     *DealsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DealsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DealsRaw struct {
	Contract *Deals // Generic contract binding to access the raw methods on
}

// DealsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DealsCallerRaw struct {
	Contract *DealsCaller // Generic read-only contract binding to access the raw methods on
}

// DealsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DealsTransactorRaw struct {
	Contract *DealsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeals creates a new instance of Deals, bound to a specific deployed contract.
func NewDeals(address common.Address, backend bind.ContractBackend) (*Deals, error) {
	contract, err := bindDeals(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deals{DealsCaller: DealsCaller{contract: contract}, DealsTransactor: DealsTransactor{contract: contract}}, nil
}

// NewDealsCaller creates a new read-only instance of Deals, bound to a specific deployed contract.
func NewDealsCaller(address common.Address, caller bind.ContractCaller) (*DealsCaller, error) {
	contract, err := bindDeals(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DealsCaller{contract: contract}, nil
}

// NewDealsTransactor creates a new write-only instance of Deals, bound to a specific deployed contract.
func NewDealsTransactor(address common.Address, transactor bind.ContractTransactor) (*DealsTransactor, error) {
	contract, err := bindDeals(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DealsTransactor{contract: contract}, nil
}

// bindDeals binds a generic wrapper to an already deployed contract.
func bindDeals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DealsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deals *DealsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deals.Contract.DealsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deals *DealsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deals.Contract.DealsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deals *DealsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deals.Contract.DealsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deals *DealsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deals *DealsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deals *DealsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deals.Contract.contract.Transact(opts, method, params...)
}

// GetDealInfo is a free data retrieval call binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) constant returns(specHach uint256, client address, hub address, price uint256, startTime uint256, workTime uint256, endTIme uint256, status uint256)
func (_Deals *DealsCaller) GetDealInfo(opts *bind.CallOpts, dealIndex *big.Int) (struct {
	SpecHach  *big.Int
	Client    common.Address
	Hub       common.Address
	Price     *big.Int
	StartTime *big.Int
	WorkTime  *big.Int
	EndTIme   *big.Int
	Status    *big.Int
}, error) {
	ret := new(struct {
		SpecHach  *big.Int
		Client    common.Address
		Hub       common.Address
		Price     *big.Int
		StartTime *big.Int
		WorkTime  *big.Int
		EndTIme   *big.Int
		Status    *big.Int
	})
	out := ret
	err := _Deals.contract.Call(opts, out, "GetDealInfo", dealIndex)
	return *ret, err
}

// GetDealInfo is a free data retrieval call binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) constant returns(specHach uint256, client address, hub address, price uint256, startTime uint256, workTime uint256, endTIme uint256, status uint256)
func (_Deals *DealsSession) GetDealInfo(dealIndex *big.Int) (struct {
	SpecHach  *big.Int
	Client    common.Address
	Hub       common.Address
	Price     *big.Int
	StartTime *big.Int
	WorkTime  *big.Int
	EndTIme   *big.Int
	Status    *big.Int
}, error) {
	return _Deals.Contract.GetDealInfo(&_Deals.CallOpts, dealIndex)
}

// GetDealInfo is a free data retrieval call binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) constant returns(specHach uint256, client address, hub address, price uint256, startTime uint256, workTime uint256, endTIme uint256, status uint256)
func (_Deals *DealsCallerSession) GetDealInfo(dealIndex *big.Int) (struct {
	SpecHach  *big.Int
	Client    common.Address
	Hub       common.Address
	Price     *big.Int
	StartTime *big.Int
	WorkTime  *big.Int
	EndTIme   *big.Int
	Status    *big.Int
}, error) {
	return _Deals.Contract.GetDealInfo(&_Deals.CallOpts, dealIndex)
}

// GetDeals is a free data retrieval call binding the contract method 0x087ca1d8.
//
// Solidity: function GetDeals(addr address) constant returns(uint256[])
func (_Deals *DealsCaller) GetDeals(opts *bind.CallOpts, addr common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Deals.contract.Call(opts, out, "GetDeals", addr)
	return *ret0, err
}

// GetDeals is a free data retrieval call binding the contract method 0x087ca1d8.
//
// Solidity: function GetDeals(addr address) constant returns(uint256[])
func (_Deals *DealsSession) GetDeals(addr common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDeals(&_Deals.CallOpts, addr)
}

// GetDeals is a free data retrieval call binding the contract method 0x087ca1d8.
//
// Solidity: function GetDeals(addr address) constant returns(uint256[])
func (_Deals *DealsCallerSession) GetDeals(addr common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDeals(&_Deals.CallOpts, addr)
}

// GetDealsAmount is a free data retrieval call binding the contract method 0xe45ea8d3.
//
// Solidity: function GetDealsAmount() constant returns(uint256)
func (_Deals *DealsCaller) GetDealsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Deals.contract.Call(opts, out, "GetDealsAmount")
	return *ret0, err
}

// GetDealsAmount is a free data retrieval call binding the contract method 0xe45ea8d3.
//
// Solidity: function GetDealsAmount() constant returns(uint256)
func (_Deals *DealsSession) GetDealsAmount() (*big.Int, error) {
	return _Deals.Contract.GetDealsAmount(&_Deals.CallOpts)
}

// GetDealsAmount is a free data retrieval call binding the contract method 0xe45ea8d3.
//
// Solidity: function GetDealsAmount() constant returns(uint256)
func (_Deals *DealsCallerSession) GetDealsAmount() (*big.Int, error) {
	return _Deals.Contract.GetDealsAmount(&_Deals.CallOpts)
}

// AcceptDeal is a paid mutator transaction binding the contract method 0x5be20de2.
//
// Solidity: function AcceptDeal(id uint256) returns()
func (_Deals *DealsTransactor) AcceptDeal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "AcceptDeal", id)
}

// AcceptDeal is a paid mutator transaction binding the contract method 0x5be20de2.
//
// Solidity: function AcceptDeal(id uint256) returns()
func (_Deals *DealsSession) AcceptDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.AcceptDeal(&_Deals.TransactOpts, id)
}

// AcceptDeal is a paid mutator transaction binding the contract method 0x5be20de2.
//
// Solidity: function AcceptDeal(id uint256) returns()
func (_Deals *DealsTransactorSession) AcceptDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.AcceptDeal(&_Deals.TransactOpts, id)
}

// CloseDeal is a paid mutator transaction binding the contract method 0xca6ad1bf.
//
// Solidity: function CloseDeal(id uint256) returns()
func (_Deals *DealsTransactor) CloseDeal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "CloseDeal", id)
}

// CloseDeal is a paid mutator transaction binding the contract method 0xca6ad1bf.
//
// Solidity: function CloseDeal(id uint256) returns()
func (_Deals *DealsSession) CloseDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.CloseDeal(&_Deals.TransactOpts, id)
}

// CloseDeal is a paid mutator transaction binding the contract method 0xca6ad1bf.
//
// Solidity: function CloseDeal(id uint256) returns()
func (_Deals *DealsTransactorSession) CloseDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.CloseDeal(&_Deals.TransactOpts, id)
}

// OpenDeal is a paid mutator transaction binding the contract method 0x3b24475f.
//
// Solidity: function OpenDeal(_hub address, _client address, _specHash uint256, _price uint256, _workTime uint256) returns()
func (_Deals *DealsTransactor) OpenDeal(opts *bind.TransactOpts, _hub common.Address, _client common.Address, _specHash *big.Int, _price *big.Int, _workTime *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "OpenDeal", _hub, _client, _specHash, _price, _workTime)
}

// OpenDeal is a paid mutator transaction binding the contract method 0x3b24475f.
//
// Solidity: function OpenDeal(_hub address, _client address, _specHash uint256, _price uint256, _workTime uint256) returns()
func (_Deals *DealsSession) OpenDeal(_hub common.Address, _client common.Address, _specHash *big.Int, _price *big.Int, _workTime *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.OpenDeal(&_Deals.TransactOpts, _hub, _client, _specHash, _price, _workTime)
}

// OpenDeal is a paid mutator transaction binding the contract method 0x3b24475f.
//
// Solidity: function OpenDeal(_hub address, _client address, _specHash uint256, _price uint256, _workTime uint256) returns()
func (_Deals *DealsTransactorSession) OpenDeal(_hub common.Address, _client common.Address, _specHash *big.Int, _price *big.Int, _workTime *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.OpenDeal(&_Deals.TransactOpts, _hub, _client, _specHash, _price, _workTime)
}
