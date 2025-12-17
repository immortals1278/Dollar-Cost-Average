// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vm

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

// VmSafeAccountAccess is an auto generated low-level Go binding around an user-defined struct.
type VmSafeAccountAccess struct {
	ChainInfo       VmSafeChainInfo
	Kind            uint8
	Account         common.Address
	Accessor        common.Address
	Initialized     bool
	OldBalance      *big.Int
	NewBalance      *big.Int
	DeployedCode    []byte
	Value           *big.Int
	Data            []byte
	Reverted        bool
	StorageAccesses []VmSafeStorageAccess
	Depth           uint64
}

// VmSafeChainInfo is an auto generated low-level Go binding around an user-defined struct.
type VmSafeChainInfo struct {
	ForkId  *big.Int
	ChainId *big.Int
}

// VmSafeDirEntry is an auto generated low-level Go binding around an user-defined struct.
type VmSafeDirEntry struct {
	ErrorMessage string
	Path         string
	Depth        uint64
	IsDir        bool
	IsSymlink    bool
}

// VmSafeEthGetLogs is an auto generated low-level Go binding around an user-defined struct.
type VmSafeEthGetLogs struct {
	Emitter          common.Address
	Topics           [][32]byte
	Data             []byte
	BlockHash        [32]byte
	BlockNumber      uint64
	TransactionHash  [32]byte
	TransactionIndex uint64
	LogIndex         *big.Int
	Removed          bool
}

// VmSafeFfiResult is an auto generated low-level Go binding around an user-defined struct.
type VmSafeFfiResult struct {
	ExitCode int32
	Stdout   []byte
	Stderr   []byte
}

// VmSafeFsMetadata is an auto generated low-level Go binding around an user-defined struct.
type VmSafeFsMetadata struct {
	IsDir     bool
	IsSymlink bool
	Length    *big.Int
	ReadOnly  bool
	Modified  *big.Int
	Accessed  *big.Int
	Created   *big.Int
}

// VmSafeGas is an auto generated low-level Go binding around an user-defined struct.
type VmSafeGas struct {
	GasLimit      uint64
	GasTotalUsed  uint64
	GasMemoryUsed uint64
	GasRefunded   int64
	GasRemaining  uint64
}

// VmSafeLog is an auto generated low-level Go binding around an user-defined struct.
type VmSafeLog struct {
	Topics  [][32]byte
	Data    []byte
	Emitter common.Address
}

// VmSafeRpc is an auto generated low-level Go binding around an user-defined struct.
type VmSafeRpc struct {
	Key string
	Url string
}

// VmSafeStorageAccess is an auto generated low-level Go binding around an user-defined struct.
type VmSafeStorageAccess struct {
	Account       common.Address
	Slot          [32]byte
	IsWrite       bool
	PreviousValue [32]byte
	NewValue      [32]byte
	Reverted      bool
}

// VmSafeWallet is an auto generated low-level Go binding around an user-defined struct.
type VmSafeWallet struct {
	Addr       common.Address
	PublicKeyX *big.Int
	PublicKeyY *big.Int
	PrivateKey *big.Int
}

// VmSafeMetaData contains all meta data concerning the VmSafe contract.
var VmSafeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"accesses\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"readSlots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"writeSlots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addr\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"keyAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbs\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbs\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbs\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbs\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbsDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbsDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbsDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbsDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRel\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRel\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRel\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRel\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRelDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRelDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRelDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRelDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"right\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"right\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"right\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"right\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"right\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"right\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"right\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"right\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"right\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"right\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"right\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"right\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"right\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"right\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"right\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"right\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"right\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"right\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"right\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"right\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"right\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"right\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertFalse\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertFalse\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGe\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGe\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGe\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGe\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGt\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGt\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGt\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGt\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLe\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLe\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLe\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLe\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLt\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLt\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLt\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLt\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"right\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"right\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"right\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"right\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"right\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"right\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"right\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"right\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"right\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"right\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"right\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"right\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"right\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"right\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"right\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"right\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"right\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"right\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"right\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"right\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"right\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"right\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertTrue\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertTrue\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assume\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"breakpoint\",\"inputs\":[{\"name\":\"char\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"breakpoint\",\"inputs\":[{\"name\":\"char\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcast\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcast\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcast\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcastRawTransaction\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"computeCreate2Address\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"computeCreate2Address\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"computeCreateAddress\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"copyFile\",\"inputs\":[{\"name\":\"from\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"copied\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"recursive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createWallet\",\"inputs\":[{\"name\":\"walletLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createWallet\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createWallet\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"walletLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"constructorArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deriveKey\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"derivationPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"language\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deriveKey\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"language\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deriveKey\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deriveKey\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"derivationPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"ensNamehash\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"envAddress\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envAddress\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBool\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBool\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBytes\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBytes\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBytes32\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBytes32\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envExists\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envInt\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envInt\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envString\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envString\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envUint\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envUint\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eth_getLogs\",\"inputs\":[{\"name\":\"fromBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"logs\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.EthGetLogs[]\",\"components\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"transactionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactionIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"removed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exists\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ffi\",\"inputs\":[{\"name\":\"commandInput\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fsMetadata\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.FsMetadata\",\"components\":[{\"name\":\"isDir\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isSymlink\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"readOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modified\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"created\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlobBaseFee\",\"inputs\":[],\"outputs\":[{\"name\":\"blobBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"creationBytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeployedCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"runtimeBytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFoundryVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLabel\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"currentLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMappingKeyAndParentOf\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"elementSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"found\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parent\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMappingLength\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mappingSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMappingSlotAt\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mappingSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"idx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRecordedLogs\",\"inputs\":[],\"outputs\":[{\"name\":\"logs\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.Log[]\",\"components\":[{\"name\":\"topics\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"indexOf\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isContext\",\"inputs\":[{\"name\":\"context\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.ForgeContext\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"keyExists\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyExistsJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyExistsToml\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"label\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastCallGas\",\"inputs\":[],\"outputs\":[{\"name\":\"gas\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Gas\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasTotalUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasMemoryUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasRefunded\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"gasRemaining\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"load\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"parseAddress\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseBool\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseBytes\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseBytes32\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseInt\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"abiEncodedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"abiEncodedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonAddress\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonAddressArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBool\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBoolArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBytes\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBytes32\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBytes32Array\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBytesArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonInt\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonIntArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonKeys\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"keys\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonString\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonStringArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonType\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonType\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonTypeArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonUint\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonUintArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseToml\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"abiEncodedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseToml\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"abiEncodedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlAddress\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlAddressArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBool\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBoolArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBytes\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBytes32\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBytes32Array\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBytesArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlInt\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlIntArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlKeys\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"keys\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlString\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlStringArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlUint\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlUintArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseUint\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"pauseGasMetering\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"projectRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prompt\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"promptAddress\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"promptSecret\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"promptSecretUint\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"promptUint\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"randomAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"randomUint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"randomUint\",\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"readDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxDepth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"entries\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.DirEntry[]\",\"components\":[{\"name\":\"errorMessage\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"depth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isDir\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isSymlink\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxDepth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"followLinks\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"entries\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.DirEntry[]\",\"components\":[{\"name\":\"errorMessage\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"depth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isDir\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isSymlink\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"entries\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.DirEntry[]\",\"components\":[{\"name\":\"errorMessage\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"depth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isDir\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isSymlink\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readFileBinary\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readLine\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"line\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readLink\",\"inputs\":[{\"name\":\"linkPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"targetPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"record\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordLogs\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rememberKey\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"keyAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"recursive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"replace\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"from\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"resumeGasMetering\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rpc\",\"inputs\":[{\"name\":\"urlOrAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"method\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rpc\",\"inputs\":[{\"name\":\"method\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rpcUrl\",\"inputs\":[{\"name\":\"rpcAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rpcUrlStructs\",\"inputs\":[],\"outputs\":[{\"name\":\"urls\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.Rpc[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rpcUrls\",\"inputs\":[],\"outputs\":[{\"name\":\"urls\",\"type\":\"string[2][]\",\"internalType\":\"string[2][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"serializeAddress\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeAddress\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBool\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBool\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBytes\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBytes\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBytes32\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBytes32\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeInt\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeInt\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeJson\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeJsonType\",\"inputs\":[{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"serializeJsonType\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeString\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeString\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeUint\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeUint\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeUintToHex\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEnv\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sign\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"sign\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"sign\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sign\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"signCompact\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signCompact\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"signCompact\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"signCompact\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"signP256\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"sleep\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"split\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delimiter\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"outputs\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"startBroadcast\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBroadcast\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBroadcast\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startMappingRecording\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startStateDiffRecording\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopAndReturnStateDiff\",\"inputs\":[],\"outputs\":[{\"name\":\"accountAccesses\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.AccountAccess[]\",\"components\":[{\"name\":\"chainInfo\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.ChainInfo\",\"components\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.AccountAccessKind\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"accessor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"oldBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deployedCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"reverted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"storageAccesses\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.StorageAccess[]\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isWrite\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"previousValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reverted\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"depth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopBroadcast\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopMappingRecording\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toBase64\",\"inputs\":[{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toBase64\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toBase64URL\",\"inputs\":[{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toBase64URL\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toLowercase\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUppercase\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"trim\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tryFfi\",\"inputs\":[{\"name\":\"commandInput\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.FfiResult\",\"components\":[{\"name\":\"exitCode\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"stdout\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"stderr\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unixTime\",\"inputs\":[],\"outputs\":[{\"name\":\"milliseconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeFileBinary\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeLine\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeToml\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeToml\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// VmSafeABI is the input ABI used to generate the binding from.
// Deprecated: Use VmSafeMetaData.ABI instead.
var VmSafeABI = VmSafeMetaData.ABI

// VmSafe is an auto generated Go binding around an Ethereum contract.
type VmSafe struct {
	VmSafeCaller     // Read-only binding to the contract
	VmSafeTransactor // Write-only binding to the contract
	VmSafeFilterer   // Log filterer for contract events
}

// VmSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type VmSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VmSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VmSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VmSafeSession struct {
	Contract     *VmSafe           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VmSafeCallerSession struct {
	Contract *VmSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VmSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VmSafeTransactorSession struct {
	Contract     *VmSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type VmSafeRaw struct {
	Contract *VmSafe // Generic contract binding to access the raw methods on
}

// VmSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VmSafeCallerRaw struct {
	Contract *VmSafeCaller // Generic read-only contract binding to access the raw methods on
}

// VmSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VmSafeTransactorRaw struct {
	Contract *VmSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVmSafe creates a new instance of VmSafe, bound to a specific deployed contract.
func NewVmSafe(address common.Address, backend bind.ContractBackend) (*VmSafe, error) {
	contract, err := bindVmSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VmSafe{VmSafeCaller: VmSafeCaller{contract: contract}, VmSafeTransactor: VmSafeTransactor{contract: contract}, VmSafeFilterer: VmSafeFilterer{contract: contract}}, nil
}

// NewVmSafeCaller creates a new read-only instance of VmSafe, bound to a specific deployed contract.
func NewVmSafeCaller(address common.Address, caller bind.ContractCaller) (*VmSafeCaller, error) {
	contract, err := bindVmSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VmSafeCaller{contract: contract}, nil
}

// NewVmSafeTransactor creates a new write-only instance of VmSafe, bound to a specific deployed contract.
func NewVmSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*VmSafeTransactor, error) {
	contract, err := bindVmSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VmSafeTransactor{contract: contract}, nil
}

// NewVmSafeFilterer creates a new log filterer instance of VmSafe, bound to a specific deployed contract.
func NewVmSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*VmSafeFilterer, error) {
	contract, err := bindVmSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VmSafeFilterer{contract: contract}, nil
}

// bindVmSafe binds a generic wrapper to an already deployed contract.
func bindVmSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VmSafeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmSafe *VmSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmSafe.Contract.VmSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmSafe *VmSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.Contract.VmSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmSafe *VmSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmSafe.Contract.VmSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmSafe *VmSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmSafe *VmSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmSafe *VmSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmSafe.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_VmSafe *VmSafeCaller) Addr(opts *bind.CallOpts, privateKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "addr", privateKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_VmSafe *VmSafeSession) Addr(privateKey *big.Int) (common.Address, error) {
	return _VmSafe.Contract.Addr(&_VmSafe.CallOpts, privateKey)
}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_VmSafe *VmSafeCallerSession) Addr(privateKey *big.Int) (common.Address, error) {
	return _VmSafe.Contract.Addr(&_VmSafe.CallOpts, privateKey)
}

// AssertApproxEqAbs is a free data retrieval call binding the contract method 0x16d207c6.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqAbs(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqAbs", left, right, maxDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs is a free data retrieval call binding the contract method 0x16d207c6.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqAbs(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqAbs(&_VmSafe.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs is a free data retrieval call binding the contract method 0x16d207c6.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqAbs(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqAbs(&_VmSafe.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs0 is a free data retrieval call binding the contract method 0x240f839d.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqAbs0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqAbs0", left, right, maxDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs0 is a free data retrieval call binding the contract method 0x240f839d.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqAbs0(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqAbs0(&_VmSafe.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs0 is a free data retrieval call binding the contract method 0x240f839d.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqAbs0(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqAbs0(&_VmSafe.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs1 is a free data retrieval call binding the contract method 0x8289e621.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqAbs1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqAbs1", left, right, maxDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs1 is a free data retrieval call binding the contract method 0x8289e621.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqAbs1(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqAbs1(&_VmSafe.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbs1 is a free data retrieval call binding the contract method 0x8289e621.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqAbs1(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqAbs1(&_VmSafe.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbs2 is a free data retrieval call binding the contract method 0xf710b062.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqAbs2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqAbs2", left, right, maxDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs2 is a free data retrieval call binding the contract method 0xf710b062.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqAbs2(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqAbs2(&_VmSafe.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbs2 is a free data retrieval call binding the contract method 0xf710b062.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqAbs2(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqAbs2(&_VmSafe.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbsDecimal is a free data retrieval call binding the contract method 0x045c55ce.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqAbsDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqAbsDecimal", left, right, maxDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal is a free data retrieval call binding the contract method 0x045c55ce.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqAbsDecimal(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqAbsDecimal(&_VmSafe.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal is a free data retrieval call binding the contract method 0x045c55ce.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqAbsDecimal(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqAbsDecimal(&_VmSafe.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal0 is a free data retrieval call binding the contract method 0x3d5bc8bc.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqAbsDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqAbsDecimal0", left, right, maxDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal0 is a free data retrieval call binding the contract method 0x3d5bc8bc.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqAbsDecimal0(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqAbsDecimal0(&_VmSafe.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal0 is a free data retrieval call binding the contract method 0x3d5bc8bc.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqAbsDecimal0(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqAbsDecimal0(&_VmSafe.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal1 is a free data retrieval call binding the contract method 0x60429eb2.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqAbsDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqAbsDecimal1", left, right, maxDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal1 is a free data retrieval call binding the contract method 0x60429eb2.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqAbsDecimal1(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqAbsDecimal1(&_VmSafe.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqAbsDecimal1 is a free data retrieval call binding the contract method 0x60429eb2.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqAbsDecimal1(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqAbsDecimal1(&_VmSafe.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqAbsDecimal2 is a free data retrieval call binding the contract method 0x6a5066d4.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqAbsDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqAbsDecimal2", left, right, maxDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal2 is a free data retrieval call binding the contract method 0x6a5066d4.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqAbsDecimal2(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqAbsDecimal2(&_VmSafe.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqAbsDecimal2 is a free data retrieval call binding the contract method 0x6a5066d4.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqAbsDecimal2(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqAbsDecimal2(&_VmSafe.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqRel is a free data retrieval call binding the contract method 0x1ecb7d33.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqRel(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqRel", left, right, maxPercentDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel is a free data retrieval call binding the contract method 0x1ecb7d33.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqRel(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqRel(&_VmSafe.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel is a free data retrieval call binding the contract method 0x1ecb7d33.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqRel(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqRel(&_VmSafe.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel0 is a free data retrieval call binding the contract method 0x8cf25ef4.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqRel0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqRel0", left, right, maxPercentDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel0 is a free data retrieval call binding the contract method 0x8cf25ef4.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqRel0(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqRel0(&_VmSafe.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRel0 is a free data retrieval call binding the contract method 0x8cf25ef4.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqRel0(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqRel0(&_VmSafe.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRel1 is a free data retrieval call binding the contract method 0xef277d72.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqRel1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqRel1", left, right, maxPercentDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel1 is a free data retrieval call binding the contract method 0xef277d72.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqRel1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqRel1(&_VmSafe.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel1 is a free data retrieval call binding the contract method 0xef277d72.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqRel1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqRel1(&_VmSafe.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel2 is a free data retrieval call binding the contract method 0xfea2d14f.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqRel2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqRel2", left, right, maxPercentDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel2 is a free data retrieval call binding the contract method 0xfea2d14f.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqRel2(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqRel2(&_VmSafe.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRel2 is a free data retrieval call binding the contract method 0xfea2d14f.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqRel2(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqRel2(&_VmSafe.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRelDecimal is a free data retrieval call binding the contract method 0x21ed2977.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqRelDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqRelDecimal", left, right, maxPercentDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal is a free data retrieval call binding the contract method 0x21ed2977.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqRelDecimal(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqRelDecimal(&_VmSafe.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal is a free data retrieval call binding the contract method 0x21ed2977.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqRelDecimal(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqRelDecimal(&_VmSafe.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal0 is a free data retrieval call binding the contract method 0x82d6c8fd.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqRelDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqRelDecimal0", left, right, maxPercentDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal0 is a free data retrieval call binding the contract method 0x82d6c8fd.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqRelDecimal0(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqRelDecimal0(&_VmSafe.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertApproxEqRelDecimal0 is a free data retrieval call binding the contract method 0x82d6c8fd.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqRelDecimal0(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqRelDecimal0(&_VmSafe.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertApproxEqRelDecimal1 is a free data retrieval call binding the contract method 0xabbf21cc.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqRelDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqRelDecimal1", left, right, maxPercentDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal1 is a free data retrieval call binding the contract method 0xabbf21cc.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqRelDecimal1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqRelDecimal1(&_VmSafe.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal1 is a free data retrieval call binding the contract method 0xabbf21cc.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqRelDecimal1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertApproxEqRelDecimal1(&_VmSafe.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal2 is a free data retrieval call binding the contract method 0xfccc11c4.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertApproxEqRelDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertApproxEqRelDecimal2", left, right, maxPercentDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal2 is a free data retrieval call binding the contract method 0xfccc11c4.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertApproxEqRelDecimal2(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqRelDecimal2(&_VmSafe.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertApproxEqRelDecimal2 is a free data retrieval call binding the contract method 0xfccc11c4.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertApproxEqRelDecimal2(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertApproxEqRelDecimal2(&_VmSafe.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertEq is a free data retrieval call binding the contract method 0x0cc9ee84.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq(opts *bind.CallOpts, left [][32]byte, right [][32]byte) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq is a free data retrieval call binding the contract method 0x0cc9ee84.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq(left [][32]byte, right [][32]byte) error {
	return _VmSafe.Contract.AssertEq(&_VmSafe.CallOpts, left, right)
}

// AssertEq is a free data retrieval call binding the contract method 0x0cc9ee84.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq(left [][32]byte, right [][32]byte) error {
	return _VmSafe.Contract.AssertEq(&_VmSafe.CallOpts, left, right)
}

// AssertEq0 is a free data retrieval call binding the contract method 0x191f1b30.
//
// Solidity: function assertEq(int256[] left, int256[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq0(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq0 is a free data retrieval call binding the contract method 0x191f1b30.
//
// Solidity: function assertEq(int256[] left, int256[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq0(left []*big.Int, right []*big.Int, error string) error {
	return _VmSafe.Contract.AssertEq0(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq0 is a free data retrieval call binding the contract method 0x191f1b30.
//
// Solidity: function assertEq(int256[] left, int256[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq0(left []*big.Int, right []*big.Int, error string) error {
	return _VmSafe.Contract.AssertEq0(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq1 is a free data retrieval call binding the contract method 0x2f2769d1.
//
// Solidity: function assertEq(address left, address right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq1(opts *bind.CallOpts, left common.Address, right common.Address, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq1", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq1 is a free data retrieval call binding the contract method 0x2f2769d1.
//
// Solidity: function assertEq(address left, address right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq1(left common.Address, right common.Address, error string) error {
	return _VmSafe.Contract.AssertEq1(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq1 is a free data retrieval call binding the contract method 0x2f2769d1.
//
// Solidity: function assertEq(address left, address right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq1(left common.Address, right common.Address, error string) error {
	return _VmSafe.Contract.AssertEq1(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq10 is a free data retrieval call binding the contract method 0x714a2f13.
//
// Solidity: function assertEq(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq10(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq10", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq10 is a free data retrieval call binding the contract method 0x714a2f13.
//
// Solidity: function assertEq(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq10(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertEq10(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq10 is a free data retrieval call binding the contract method 0x714a2f13.
//
// Solidity: function assertEq(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq10(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertEq10(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq11 is a free data retrieval call binding the contract method 0x7c84c69b.
//
// Solidity: function assertEq(bytes32 left, bytes32 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq11(opts *bind.CallOpts, left [32]byte, right [32]byte) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq11", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq11 is a free data retrieval call binding the contract method 0x7c84c69b.
//
// Solidity: function assertEq(bytes32 left, bytes32 right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq11(left [32]byte, right [32]byte) error {
	return _VmSafe.Contract.AssertEq11(&_VmSafe.CallOpts, left, right)
}

// AssertEq11 is a free data retrieval call binding the contract method 0x7c84c69b.
//
// Solidity: function assertEq(bytes32 left, bytes32 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq11(left [32]byte, right [32]byte) error {
	return _VmSafe.Contract.AssertEq11(&_VmSafe.CallOpts, left, right)
}

// AssertEq12 is a free data retrieval call binding the contract method 0x88b44c85.
//
// Solidity: function assertEq(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq12(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq12", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq12 is a free data retrieval call binding the contract method 0x88b44c85.
//
// Solidity: function assertEq(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq12(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertEq12(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq12 is a free data retrieval call binding the contract method 0x88b44c85.
//
// Solidity: function assertEq(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq12(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertEq12(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq13 is a free data retrieval call binding the contract method 0x975d5a12.
//
// Solidity: function assertEq(uint256[] left, uint256[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq13(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq13", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq13 is a free data retrieval call binding the contract method 0x975d5a12.
//
// Solidity: function assertEq(uint256[] left, uint256[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq13(left []*big.Int, right []*big.Int) error {
	return _VmSafe.Contract.AssertEq13(&_VmSafe.CallOpts, left, right)
}

// AssertEq13 is a free data retrieval call binding the contract method 0x975d5a12.
//
// Solidity: function assertEq(uint256[] left, uint256[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq13(left []*big.Int, right []*big.Int) error {
	return _VmSafe.Contract.AssertEq13(&_VmSafe.CallOpts, left, right)
}

// AssertEq14 is a free data retrieval call binding the contract method 0x97624631.
//
// Solidity: function assertEq(bytes left, bytes right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq14(opts *bind.CallOpts, left []byte, right []byte) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq14", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq14 is a free data retrieval call binding the contract method 0x97624631.
//
// Solidity: function assertEq(bytes left, bytes right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq14(left []byte, right []byte) error {
	return _VmSafe.Contract.AssertEq14(&_VmSafe.CallOpts, left, right)
}

// AssertEq14 is a free data retrieval call binding the contract method 0x97624631.
//
// Solidity: function assertEq(bytes left, bytes right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq14(left []byte, right []byte) error {
	return _VmSafe.Contract.AssertEq14(&_VmSafe.CallOpts, left, right)
}

// AssertEq15 is a free data retrieval call binding the contract method 0x98296c54.
//
// Solidity: function assertEq(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq15(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq15", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq15 is a free data retrieval call binding the contract method 0x98296c54.
//
// Solidity: function assertEq(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq15(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertEq15(&_VmSafe.CallOpts, left, right)
}

// AssertEq15 is a free data retrieval call binding the contract method 0x98296c54.
//
// Solidity: function assertEq(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq15(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertEq15(&_VmSafe.CallOpts, left, right)
}

// AssertEq16 is a free data retrieval call binding the contract method 0xc1fa1ed0.
//
// Solidity: function assertEq(bytes32 left, bytes32 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq16(opts *bind.CallOpts, left [32]byte, right [32]byte, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq16", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq16 is a free data retrieval call binding the contract method 0xc1fa1ed0.
//
// Solidity: function assertEq(bytes32 left, bytes32 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq16(left [32]byte, right [32]byte, error string) error {
	return _VmSafe.Contract.AssertEq16(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq16 is a free data retrieval call binding the contract method 0xc1fa1ed0.
//
// Solidity: function assertEq(bytes32 left, bytes32 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq16(left [32]byte, right [32]byte, error string) error {
	return _VmSafe.Contract.AssertEq16(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq17 is a free data retrieval call binding the contract method 0xcf1c049c.
//
// Solidity: function assertEq(string[] left, string[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq17(opts *bind.CallOpts, left []string, right []string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq17", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq17 is a free data retrieval call binding the contract method 0xcf1c049c.
//
// Solidity: function assertEq(string[] left, string[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq17(left []string, right []string) error {
	return _VmSafe.Contract.AssertEq17(&_VmSafe.CallOpts, left, right)
}

// AssertEq17 is a free data retrieval call binding the contract method 0xcf1c049c.
//
// Solidity: function assertEq(string[] left, string[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq17(left []string, right []string) error {
	return _VmSafe.Contract.AssertEq17(&_VmSafe.CallOpts, left, right)
}

// AssertEq18 is a free data retrieval call binding the contract method 0xe03e9177.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq18(opts *bind.CallOpts, left [][32]byte, right [][32]byte, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq18", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq18 is a free data retrieval call binding the contract method 0xe03e9177.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq18(left [][32]byte, right [][32]byte, error string) error {
	return _VmSafe.Contract.AssertEq18(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq18 is a free data retrieval call binding the contract method 0xe03e9177.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq18(left [][32]byte, right [][32]byte, error string) error {
	return _VmSafe.Contract.AssertEq18(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq19 is a free data retrieval call binding the contract method 0xe24fed00.
//
// Solidity: function assertEq(bytes left, bytes right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq19(opts *bind.CallOpts, left []byte, right []byte, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq19", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq19 is a free data retrieval call binding the contract method 0xe24fed00.
//
// Solidity: function assertEq(bytes left, bytes right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq19(left []byte, right []byte, error string) error {
	return _VmSafe.Contract.AssertEq19(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq19 is a free data retrieval call binding the contract method 0xe24fed00.
//
// Solidity: function assertEq(bytes left, bytes right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq19(left []byte, right []byte, error string) error {
	return _VmSafe.Contract.AssertEq19(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq2 is a free data retrieval call binding the contract method 0x36f656d8.
//
// Solidity: function assertEq(string left, string right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq2(opts *bind.CallOpts, left string, right string, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq2 is a free data retrieval call binding the contract method 0x36f656d8.
//
// Solidity: function assertEq(string left, string right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq2(left string, right string, error string) error {
	return _VmSafe.Contract.AssertEq2(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq2 is a free data retrieval call binding the contract method 0x36f656d8.
//
// Solidity: function assertEq(string left, string right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq2(left string, right string, error string) error {
	return _VmSafe.Contract.AssertEq2(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq20 is a free data retrieval call binding the contract method 0xe48a8f8d.
//
// Solidity: function assertEq(bool[] left, bool[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq20(opts *bind.CallOpts, left []bool, right []bool, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq20", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq20 is a free data retrieval call binding the contract method 0xe48a8f8d.
//
// Solidity: function assertEq(bool[] left, bool[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq20(left []bool, right []bool, error string) error {
	return _VmSafe.Contract.AssertEq20(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq20 is a free data retrieval call binding the contract method 0xe48a8f8d.
//
// Solidity: function assertEq(bool[] left, bool[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq20(left []bool, right []bool, error string) error {
	return _VmSafe.Contract.AssertEq20(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq21 is a free data retrieval call binding the contract method 0xe5fb9b4a.
//
// Solidity: function assertEq(bytes[] left, bytes[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq21(opts *bind.CallOpts, left [][]byte, right [][]byte) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq21", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq21 is a free data retrieval call binding the contract method 0xe5fb9b4a.
//
// Solidity: function assertEq(bytes[] left, bytes[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq21(left [][]byte, right [][]byte) error {
	return _VmSafe.Contract.AssertEq21(&_VmSafe.CallOpts, left, right)
}

// AssertEq21 is a free data retrieval call binding the contract method 0xe5fb9b4a.
//
// Solidity: function assertEq(bytes[] left, bytes[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq21(left [][]byte, right [][]byte) error {
	return _VmSafe.Contract.AssertEq21(&_VmSafe.CallOpts, left, right)
}

// AssertEq22 is a free data retrieval call binding the contract method 0xeff6b27d.
//
// Solidity: function assertEq(string[] left, string[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq22(opts *bind.CallOpts, left []string, right []string, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq22", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq22 is a free data retrieval call binding the contract method 0xeff6b27d.
//
// Solidity: function assertEq(string[] left, string[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq22(left []string, right []string, error string) error {
	return _VmSafe.Contract.AssertEq22(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq22 is a free data retrieval call binding the contract method 0xeff6b27d.
//
// Solidity: function assertEq(string[] left, string[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq22(left []string, right []string, error string) error {
	return _VmSafe.Contract.AssertEq22(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq23 is a free data retrieval call binding the contract method 0xf320d963.
//
// Solidity: function assertEq(string left, string right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq23(opts *bind.CallOpts, left string, right string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq23", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq23 is a free data retrieval call binding the contract method 0xf320d963.
//
// Solidity: function assertEq(string left, string right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq23(left string, right string) error {
	return _VmSafe.Contract.AssertEq23(&_VmSafe.CallOpts, left, right)
}

// AssertEq23 is a free data retrieval call binding the contract method 0xf320d963.
//
// Solidity: function assertEq(string left, string right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq23(left string, right string) error {
	return _VmSafe.Contract.AssertEq23(&_VmSafe.CallOpts, left, right)
}

// AssertEq24 is a free data retrieval call binding the contract method 0xf413f0b6.
//
// Solidity: function assertEq(bytes[] left, bytes[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq24(opts *bind.CallOpts, left [][]byte, right [][]byte, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq24", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq24 is a free data retrieval call binding the contract method 0xf413f0b6.
//
// Solidity: function assertEq(bytes[] left, bytes[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq24(left [][]byte, right [][]byte, error string) error {
	return _VmSafe.Contract.AssertEq24(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq24 is a free data retrieval call binding the contract method 0xf413f0b6.
//
// Solidity: function assertEq(bytes[] left, bytes[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq24(left [][]byte, right [][]byte, error string) error {
	return _VmSafe.Contract.AssertEq24(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq25 is a free data retrieval call binding the contract method 0xf7fe3477.
//
// Solidity: function assertEq(bool left, bool right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq25(opts *bind.CallOpts, left bool, right bool) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq25", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq25 is a free data retrieval call binding the contract method 0xf7fe3477.
//
// Solidity: function assertEq(bool left, bool right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq25(left bool, right bool) error {
	return _VmSafe.Contract.AssertEq25(&_VmSafe.CallOpts, left, right)
}

// AssertEq25 is a free data retrieval call binding the contract method 0xf7fe3477.
//
// Solidity: function assertEq(bool left, bool right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq25(left bool, right bool) error {
	return _VmSafe.Contract.AssertEq25(&_VmSafe.CallOpts, left, right)
}

// AssertEq26 is a free data retrieval call binding the contract method 0xfe74f05b.
//
// Solidity: function assertEq(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq26(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq26", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq26 is a free data retrieval call binding the contract method 0xfe74f05b.
//
// Solidity: function assertEq(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq26(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertEq26(&_VmSafe.CallOpts, left, right)
}

// AssertEq26 is a free data retrieval call binding the contract method 0xfe74f05b.
//
// Solidity: function assertEq(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq26(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertEq26(&_VmSafe.CallOpts, left, right)
}

// AssertEq3 is a free data retrieval call binding the contract method 0x3868ac34.
//
// Solidity: function assertEq(address[] left, address[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq3(opts *bind.CallOpts, left []common.Address, right []common.Address) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq3", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq3 is a free data retrieval call binding the contract method 0x3868ac34.
//
// Solidity: function assertEq(address[] left, address[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq3(left []common.Address, right []common.Address) error {
	return _VmSafe.Contract.AssertEq3(&_VmSafe.CallOpts, left, right)
}

// AssertEq3 is a free data retrieval call binding the contract method 0x3868ac34.
//
// Solidity: function assertEq(address[] left, address[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq3(left []common.Address, right []common.Address) error {
	return _VmSafe.Contract.AssertEq3(&_VmSafe.CallOpts, left, right)
}

// AssertEq4 is a free data retrieval call binding the contract method 0x3e9173c5.
//
// Solidity: function assertEq(address[] left, address[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq4(opts *bind.CallOpts, left []common.Address, right []common.Address, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq4", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq4 is a free data retrieval call binding the contract method 0x3e9173c5.
//
// Solidity: function assertEq(address[] left, address[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq4(left []common.Address, right []common.Address, error string) error {
	return _VmSafe.Contract.AssertEq4(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq4 is a free data retrieval call binding the contract method 0x3e9173c5.
//
// Solidity: function assertEq(address[] left, address[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq4(left []common.Address, right []common.Address, error string) error {
	return _VmSafe.Contract.AssertEq4(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq5 is a free data retrieval call binding the contract method 0x4db19e7e.
//
// Solidity: function assertEq(bool left, bool right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq5(opts *bind.CallOpts, left bool, right bool, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq5", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq5 is a free data retrieval call binding the contract method 0x4db19e7e.
//
// Solidity: function assertEq(bool left, bool right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq5(left bool, right bool, error string) error {
	return _VmSafe.Contract.AssertEq5(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq5 is a free data retrieval call binding the contract method 0x4db19e7e.
//
// Solidity: function assertEq(bool left, bool right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq5(left bool, right bool, error string) error {
	return _VmSafe.Contract.AssertEq5(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq6 is a free data retrieval call binding the contract method 0x515361f6.
//
// Solidity: function assertEq(address left, address right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq6(opts *bind.CallOpts, left common.Address, right common.Address) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq6", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq6 is a free data retrieval call binding the contract method 0x515361f6.
//
// Solidity: function assertEq(address left, address right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq6(left common.Address, right common.Address) error {
	return _VmSafe.Contract.AssertEq6(&_VmSafe.CallOpts, left, right)
}

// AssertEq6 is a free data retrieval call binding the contract method 0x515361f6.
//
// Solidity: function assertEq(address left, address right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq6(left common.Address, right common.Address) error {
	return _VmSafe.Contract.AssertEq6(&_VmSafe.CallOpts, left, right)
}

// AssertEq7 is a free data retrieval call binding the contract method 0x5d18c73a.
//
// Solidity: function assertEq(uint256[] left, uint256[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq7(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq7", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq7 is a free data retrieval call binding the contract method 0x5d18c73a.
//
// Solidity: function assertEq(uint256[] left, uint256[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEq7(left []*big.Int, right []*big.Int, error string) error {
	return _VmSafe.Contract.AssertEq7(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq7 is a free data retrieval call binding the contract method 0x5d18c73a.
//
// Solidity: function assertEq(uint256[] left, uint256[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq7(left []*big.Int, right []*big.Int, error string) error {
	return _VmSafe.Contract.AssertEq7(&_VmSafe.CallOpts, left, right, error)
}

// AssertEq8 is a free data retrieval call binding the contract method 0x707df785.
//
// Solidity: function assertEq(bool[] left, bool[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq8(opts *bind.CallOpts, left []bool, right []bool) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq8", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq8 is a free data retrieval call binding the contract method 0x707df785.
//
// Solidity: function assertEq(bool[] left, bool[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq8(left []bool, right []bool) error {
	return _VmSafe.Contract.AssertEq8(&_VmSafe.CallOpts, left, right)
}

// AssertEq8 is a free data retrieval call binding the contract method 0x707df785.
//
// Solidity: function assertEq(bool[] left, bool[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq8(left []bool, right []bool) error {
	return _VmSafe.Contract.AssertEq8(&_VmSafe.CallOpts, left, right)
}

// AssertEq9 is a free data retrieval call binding the contract method 0x711043ac.
//
// Solidity: function assertEq(int256[] left, int256[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertEq9(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEq9", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq9 is a free data retrieval call binding the contract method 0x711043ac.
//
// Solidity: function assertEq(int256[] left, int256[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertEq9(left []*big.Int, right []*big.Int) error {
	return _VmSafe.Contract.AssertEq9(&_VmSafe.CallOpts, left, right)
}

// AssertEq9 is a free data retrieval call binding the contract method 0x711043ac.
//
// Solidity: function assertEq(int256[] left, int256[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEq9(left []*big.Int, right []*big.Int) error {
	return _VmSafe.Contract.AssertEq9(&_VmSafe.CallOpts, left, right)
}

// AssertEqDecimal is a free data retrieval call binding the contract method 0x27af7d9c.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertEqDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEqDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal is a free data retrieval call binding the contract method 0x27af7d9c.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertEqDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertEqDecimal is a free data retrieval call binding the contract method 0x27af7d9c.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertEqDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertEqDecimal0 is a free data retrieval call binding the contract method 0x48016c04.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertEqDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEqDecimal0", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal0 is a free data retrieval call binding the contract method 0x48016c04.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertEqDecimal0(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertEqDecimal0 is a free data retrieval call binding the contract method 0x48016c04.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertEqDecimal0(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertEqDecimal1 is a free data retrieval call binding the contract method 0x7e77b0c5.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEqDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEqDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal1 is a free data retrieval call binding the contract method 0x7e77b0c5.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertEqDecimal1(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertEqDecimal1 is a free data retrieval call binding the contract method 0x7e77b0c5.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertEqDecimal1(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertEqDecimal2 is a free data retrieval call binding the contract method 0xd0cbbdef.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertEqDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertEqDecimal2", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal2 is a free data retrieval call binding the contract method 0xd0cbbdef.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertEqDecimal2(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertEqDecimal2 is a free data retrieval call binding the contract method 0xd0cbbdef.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertEqDecimal2(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertFalse(opts *bind.CallOpts, condition bool, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertFalse", condition, error)

	if err != nil {
		return err
	}

	return err

}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertFalse(condition bool, error string) error {
	return _VmSafe.Contract.AssertFalse(&_VmSafe.CallOpts, condition, error)
}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertFalse(condition bool, error string) error {
	return _VmSafe.Contract.AssertFalse(&_VmSafe.CallOpts, condition, error)
}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_VmSafe *VmSafeCaller) AssertFalse0(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertFalse0", condition)

	if err != nil {
		return err
	}

	return err

}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_VmSafe *VmSafeSession) AssertFalse0(condition bool) error {
	return _VmSafe.Contract.AssertFalse0(&_VmSafe.CallOpts, condition)
}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertFalse0(condition bool) error {
	return _VmSafe.Contract.AssertFalse0(&_VmSafe.CallOpts, condition)
}

// AssertGe is a free data retrieval call binding the contract method 0x0a30b771.
//
// Solidity: function assertGe(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertGe(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGe", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGe is a free data retrieval call binding the contract method 0x0a30b771.
//
// Solidity: function assertGe(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertGe(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertGe(&_VmSafe.CallOpts, left, right)
}

// AssertGe is a free data retrieval call binding the contract method 0x0a30b771.
//
// Solidity: function assertGe(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGe(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertGe(&_VmSafe.CallOpts, left, right)
}

// AssertGe0 is a free data retrieval call binding the contract method 0xa84328dd.
//
// Solidity: function assertGe(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertGe0(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGe0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGe0 is a free data retrieval call binding the contract method 0xa84328dd.
//
// Solidity: function assertGe(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertGe0(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertGe0(&_VmSafe.CallOpts, left, right, error)
}

// AssertGe0 is a free data retrieval call binding the contract method 0xa84328dd.
//
// Solidity: function assertGe(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGe0(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertGe0(&_VmSafe.CallOpts, left, right, error)
}

// AssertGe1 is a free data retrieval call binding the contract method 0xa8d4d1d9.
//
// Solidity: function assertGe(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertGe1(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGe1", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGe1 is a free data retrieval call binding the contract method 0xa8d4d1d9.
//
// Solidity: function assertGe(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertGe1(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertGe1(&_VmSafe.CallOpts, left, right)
}

// AssertGe1 is a free data retrieval call binding the contract method 0xa8d4d1d9.
//
// Solidity: function assertGe(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGe1(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertGe1(&_VmSafe.CallOpts, left, right)
}

// AssertGe2 is a free data retrieval call binding the contract method 0xe25242c0.
//
// Solidity: function assertGe(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertGe2(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGe2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGe2 is a free data retrieval call binding the contract method 0xe25242c0.
//
// Solidity: function assertGe(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertGe2(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertGe2(&_VmSafe.CallOpts, left, right, error)
}

// AssertGe2 is a free data retrieval call binding the contract method 0xe25242c0.
//
// Solidity: function assertGe(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGe2(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertGe2(&_VmSafe.CallOpts, left, right, error)
}

// AssertGeDecimal is a free data retrieval call binding the contract method 0x3d1fe08a.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertGeDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGeDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal is a free data retrieval call binding the contract method 0x3d1fe08a.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertGeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertGeDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertGeDecimal is a free data retrieval call binding the contract method 0x3d1fe08a.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertGeDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertGeDecimal0 is a free data retrieval call binding the contract method 0x5df93c9b.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertGeDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGeDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal0 is a free data retrieval call binding the contract method 0x5df93c9b.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertGeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertGeDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal0 is a free data retrieval call binding the contract method 0x5df93c9b.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertGeDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal1 is a free data retrieval call binding the contract method 0x8bff9133.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertGeDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGeDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal1 is a free data retrieval call binding the contract method 0x8bff9133.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertGeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertGeDecimal1(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal1 is a free data retrieval call binding the contract method 0x8bff9133.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertGeDecimal1(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal2 is a free data retrieval call binding the contract method 0xdc28c0f1.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertGeDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGeDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal2 is a free data retrieval call binding the contract method 0xdc28c0f1.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertGeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertGeDecimal2(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertGeDecimal2 is a free data retrieval call binding the contract method 0xdc28c0f1.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertGeDecimal2(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertGt is a free data retrieval call binding the contract method 0x5a362d45.
//
// Solidity: function assertGt(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertGt(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGt", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGt is a free data retrieval call binding the contract method 0x5a362d45.
//
// Solidity: function assertGt(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertGt(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertGt(&_VmSafe.CallOpts, left, right)
}

// AssertGt is a free data retrieval call binding the contract method 0x5a362d45.
//
// Solidity: function assertGt(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGt(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertGt(&_VmSafe.CallOpts, left, right)
}

// AssertGt0 is a free data retrieval call binding the contract method 0xd9a3c4d2.
//
// Solidity: function assertGt(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertGt0(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGt0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGt0 is a free data retrieval call binding the contract method 0xd9a3c4d2.
//
// Solidity: function assertGt(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertGt0(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertGt0(&_VmSafe.CallOpts, left, right, error)
}

// AssertGt0 is a free data retrieval call binding the contract method 0xd9a3c4d2.
//
// Solidity: function assertGt(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGt0(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertGt0(&_VmSafe.CallOpts, left, right, error)
}

// AssertGt1 is a free data retrieval call binding the contract method 0xdb07fcd2.
//
// Solidity: function assertGt(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertGt1(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGt1", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGt1 is a free data retrieval call binding the contract method 0xdb07fcd2.
//
// Solidity: function assertGt(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertGt1(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertGt1(&_VmSafe.CallOpts, left, right)
}

// AssertGt1 is a free data retrieval call binding the contract method 0xdb07fcd2.
//
// Solidity: function assertGt(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGt1(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertGt1(&_VmSafe.CallOpts, left, right)
}

// AssertGt2 is a free data retrieval call binding the contract method 0xf8d33b9b.
//
// Solidity: function assertGt(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertGt2(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGt2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGt2 is a free data retrieval call binding the contract method 0xf8d33b9b.
//
// Solidity: function assertGt(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertGt2(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertGt2(&_VmSafe.CallOpts, left, right, error)
}

// AssertGt2 is a free data retrieval call binding the contract method 0xf8d33b9b.
//
// Solidity: function assertGt(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGt2(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertGt2(&_VmSafe.CallOpts, left, right, error)
}

// AssertGtDecimal is a free data retrieval call binding the contract method 0x04a5c7ab.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertGtDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGtDecimal", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal is a free data retrieval call binding the contract method 0x04a5c7ab.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertGtDecimal(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertGtDecimal(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal is a free data retrieval call binding the contract method 0x04a5c7ab.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGtDecimal(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertGtDecimal(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal0 is a free data retrieval call binding the contract method 0x64949a8d.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertGtDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGtDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal0 is a free data retrieval call binding the contract method 0x64949a8d.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertGtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertGtDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal0 is a free data retrieval call binding the contract method 0x64949a8d.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertGtDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal1 is a free data retrieval call binding the contract method 0x78611f0e.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertGtDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGtDecimal1", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal1 is a free data retrieval call binding the contract method 0x78611f0e.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertGtDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertGtDecimal1(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertGtDecimal1 is a free data retrieval call binding the contract method 0x78611f0e.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGtDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertGtDecimal1(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertGtDecimal2 is a free data retrieval call binding the contract method 0xeccd2437.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertGtDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertGtDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal2 is a free data retrieval call binding the contract method 0xeccd2437.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertGtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertGtDecimal2(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertGtDecimal2 is a free data retrieval call binding the contract method 0xeccd2437.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertGtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertGtDecimal2(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertLe is a free data retrieval call binding the contract method 0x4dfe692c.
//
// Solidity: function assertLe(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertLe(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLe", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLe is a free data retrieval call binding the contract method 0x4dfe692c.
//
// Solidity: function assertLe(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertLe(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertLe(&_VmSafe.CallOpts, left, right, error)
}

// AssertLe is a free data retrieval call binding the contract method 0x4dfe692c.
//
// Solidity: function assertLe(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLe(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertLe(&_VmSafe.CallOpts, left, right, error)
}

// AssertLe0 is a free data retrieval call binding the contract method 0x8466f415.
//
// Solidity: function assertLe(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertLe0(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLe0", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLe0 is a free data retrieval call binding the contract method 0x8466f415.
//
// Solidity: function assertLe(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertLe0(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertLe0(&_VmSafe.CallOpts, left, right)
}

// AssertLe0 is a free data retrieval call binding the contract method 0x8466f415.
//
// Solidity: function assertLe(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLe0(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertLe0(&_VmSafe.CallOpts, left, right)
}

// AssertLe1 is a free data retrieval call binding the contract method 0x95fd154e.
//
// Solidity: function assertLe(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertLe1(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLe1", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLe1 is a free data retrieval call binding the contract method 0x95fd154e.
//
// Solidity: function assertLe(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertLe1(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertLe1(&_VmSafe.CallOpts, left, right)
}

// AssertLe1 is a free data retrieval call binding the contract method 0x95fd154e.
//
// Solidity: function assertLe(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLe1(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertLe1(&_VmSafe.CallOpts, left, right)
}

// AssertLe2 is a free data retrieval call binding the contract method 0xd17d4b0d.
//
// Solidity: function assertLe(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertLe2(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLe2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLe2 is a free data retrieval call binding the contract method 0xd17d4b0d.
//
// Solidity: function assertLe(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertLe2(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertLe2(&_VmSafe.CallOpts, left, right, error)
}

// AssertLe2 is a free data retrieval call binding the contract method 0xd17d4b0d.
//
// Solidity: function assertLe(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLe2(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertLe2(&_VmSafe.CallOpts, left, right, error)
}

// AssertLeDecimal is a free data retrieval call binding the contract method 0x11d1364a.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertLeDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLeDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal is a free data retrieval call binding the contract method 0x11d1364a.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertLeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertLeDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertLeDecimal is a free data retrieval call binding the contract method 0x11d1364a.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertLeDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertLeDecimal0 is a free data retrieval call binding the contract method 0x7fefbbe0.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertLeDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLeDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal0 is a free data retrieval call binding the contract method 0x7fefbbe0.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertLeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertLeDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal0 is a free data retrieval call binding the contract method 0x7fefbbe0.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertLeDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal1 is a free data retrieval call binding the contract method 0xaa5cf788.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertLeDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLeDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal1 is a free data retrieval call binding the contract method 0xaa5cf788.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertLeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertLeDecimal1(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal1 is a free data retrieval call binding the contract method 0xaa5cf788.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertLeDecimal1(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal2 is a free data retrieval call binding the contract method 0xc304aab7.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertLeDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLeDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal2 is a free data retrieval call binding the contract method 0xc304aab7.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertLeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertLeDecimal2(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertLeDecimal2 is a free data retrieval call binding the contract method 0xc304aab7.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertLeDecimal2(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertLt is a free data retrieval call binding the contract method 0x3e914080.
//
// Solidity: function assertLt(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertLt(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLt", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLt is a free data retrieval call binding the contract method 0x3e914080.
//
// Solidity: function assertLt(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertLt(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertLt(&_VmSafe.CallOpts, left, right)
}

// AssertLt is a free data retrieval call binding the contract method 0x3e914080.
//
// Solidity: function assertLt(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLt(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertLt(&_VmSafe.CallOpts, left, right)
}

// AssertLt0 is a free data retrieval call binding the contract method 0x65d5c135.
//
// Solidity: function assertLt(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertLt0(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLt0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLt0 is a free data retrieval call binding the contract method 0x65d5c135.
//
// Solidity: function assertLt(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertLt0(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertLt0(&_VmSafe.CallOpts, left, right, error)
}

// AssertLt0 is a free data retrieval call binding the contract method 0x65d5c135.
//
// Solidity: function assertLt(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLt0(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertLt0(&_VmSafe.CallOpts, left, right, error)
}

// AssertLt1 is a free data retrieval call binding the contract method 0x9ff531e3.
//
// Solidity: function assertLt(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertLt1(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLt1", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLt1 is a free data retrieval call binding the contract method 0x9ff531e3.
//
// Solidity: function assertLt(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertLt1(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertLt1(&_VmSafe.CallOpts, left, right, error)
}

// AssertLt1 is a free data retrieval call binding the contract method 0x9ff531e3.
//
// Solidity: function assertLt(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLt1(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertLt1(&_VmSafe.CallOpts, left, right, error)
}

// AssertLt2 is a free data retrieval call binding the contract method 0xb12fc005.
//
// Solidity: function assertLt(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertLt2(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLt2", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLt2 is a free data retrieval call binding the contract method 0xb12fc005.
//
// Solidity: function assertLt(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertLt2(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertLt2(&_VmSafe.CallOpts, left, right)
}

// AssertLt2 is a free data retrieval call binding the contract method 0xb12fc005.
//
// Solidity: function assertLt(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLt2(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertLt2(&_VmSafe.CallOpts, left, right)
}

// AssertLtDecimal is a free data retrieval call binding the contract method 0x2077337e.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertLtDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLtDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal is a free data retrieval call binding the contract method 0x2077337e.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertLtDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertLtDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertLtDecimal is a free data retrieval call binding the contract method 0x2077337e.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLtDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertLtDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertLtDecimal0 is a free data retrieval call binding the contract method 0x40f0b4e0.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertLtDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLtDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal0 is a free data retrieval call binding the contract method 0x40f0b4e0.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertLtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertLtDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal0 is a free data retrieval call binding the contract method 0x40f0b4e0.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertLtDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal1 is a free data retrieval call binding the contract method 0xa972d037.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertLtDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLtDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal1 is a free data retrieval call binding the contract method 0xa972d037.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertLtDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertLtDecimal1(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal1 is a free data retrieval call binding the contract method 0xa972d037.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLtDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertLtDecimal1(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal2 is a free data retrieval call binding the contract method 0xdbe8d88b.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertLtDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertLtDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal2 is a free data retrieval call binding the contract method 0xdbe8d88b.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertLtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertLtDecimal2(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertLtDecimal2 is a free data retrieval call binding the contract method 0xdbe8d88b.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertLtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertLtDecimal2(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertNotEq is a free data retrieval call binding the contract method 0x0603ea68.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq(opts *bind.CallOpts, left [][32]byte, right [][32]byte) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq is a free data retrieval call binding the contract method 0x0603ea68.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq(left [][32]byte, right [][32]byte) error {
	return _VmSafe.Contract.AssertNotEq(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq is a free data retrieval call binding the contract method 0x0603ea68.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq(left [][32]byte, right [][32]byte) error {
	return _VmSafe.Contract.AssertNotEq(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq0 is a free data retrieval call binding the contract method 0x0b72f4ef.
//
// Solidity: function assertNotEq(int256[] left, int256[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq0(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq0", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq0 is a free data retrieval call binding the contract method 0x0b72f4ef.
//
// Solidity: function assertNotEq(int256[] left, int256[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq0(left []*big.Int, right []*big.Int) error {
	return _VmSafe.Contract.AssertNotEq0(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq0 is a free data retrieval call binding the contract method 0x0b72f4ef.
//
// Solidity: function assertNotEq(int256[] left, int256[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq0(left []*big.Int, right []*big.Int) error {
	return _VmSafe.Contract.AssertNotEq0(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq1 is a free data retrieval call binding the contract method 0x1091a261.
//
// Solidity: function assertNotEq(bool left, bool right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq1(opts *bind.CallOpts, left bool, right bool, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq1", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq1 is a free data retrieval call binding the contract method 0x1091a261.
//
// Solidity: function assertNotEq(bool left, bool right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq1(left bool, right bool, error string) error {
	return _VmSafe.Contract.AssertNotEq1(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq1 is a free data retrieval call binding the contract method 0x1091a261.
//
// Solidity: function assertNotEq(bool left, bool right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq1(left bool, right bool, error string) error {
	return _VmSafe.Contract.AssertNotEq1(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq10 is a free data retrieval call binding the contract method 0x6a8237b3.
//
// Solidity: function assertNotEq(string left, string right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq10(opts *bind.CallOpts, left string, right string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq10", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq10 is a free data retrieval call binding the contract method 0x6a8237b3.
//
// Solidity: function assertNotEq(string left, string right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq10(left string, right string) error {
	return _VmSafe.Contract.AssertNotEq10(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq10 is a free data retrieval call binding the contract method 0x6a8237b3.
//
// Solidity: function assertNotEq(string left, string right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq10(left string, right string) error {
	return _VmSafe.Contract.AssertNotEq10(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq11 is a free data retrieval call binding the contract method 0x72c7e0b5.
//
// Solidity: function assertNotEq(address[] left, address[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq11(opts *bind.CallOpts, left []common.Address, right []common.Address, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq11", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq11 is a free data retrieval call binding the contract method 0x72c7e0b5.
//
// Solidity: function assertNotEq(address[] left, address[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq11(left []common.Address, right []common.Address, error string) error {
	return _VmSafe.Contract.AssertNotEq11(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq11 is a free data retrieval call binding the contract method 0x72c7e0b5.
//
// Solidity: function assertNotEq(address[] left, address[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq11(left []common.Address, right []common.Address, error string) error {
	return _VmSafe.Contract.AssertNotEq11(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq12 is a free data retrieval call binding the contract method 0x78bdcea7.
//
// Solidity: function assertNotEq(string left, string right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq12(opts *bind.CallOpts, left string, right string, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq12", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq12 is a free data retrieval call binding the contract method 0x78bdcea7.
//
// Solidity: function assertNotEq(string left, string right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq12(left string, right string, error string) error {
	return _VmSafe.Contract.AssertNotEq12(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq12 is a free data retrieval call binding the contract method 0x78bdcea7.
//
// Solidity: function assertNotEq(string left, string right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq12(left string, right string, error string) error {
	return _VmSafe.Contract.AssertNotEq12(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq13 is a free data retrieval call binding the contract method 0x8775a591.
//
// Solidity: function assertNotEq(address left, address right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq13(opts *bind.CallOpts, left common.Address, right common.Address, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq13", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq13 is a free data retrieval call binding the contract method 0x8775a591.
//
// Solidity: function assertNotEq(address left, address right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq13(left common.Address, right common.Address, error string) error {
	return _VmSafe.Contract.AssertNotEq13(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq13 is a free data retrieval call binding the contract method 0x8775a591.
//
// Solidity: function assertNotEq(address left, address right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq13(left common.Address, right common.Address, error string) error {
	return _VmSafe.Contract.AssertNotEq13(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq14 is a free data retrieval call binding the contract method 0x898e83fc.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq14(opts *bind.CallOpts, left [32]byte, right [32]byte) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq14", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq14 is a free data retrieval call binding the contract method 0x898e83fc.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq14(left [32]byte, right [32]byte) error {
	return _VmSafe.Contract.AssertNotEq14(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq14 is a free data retrieval call binding the contract method 0x898e83fc.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq14(left [32]byte, right [32]byte) error {
	return _VmSafe.Contract.AssertNotEq14(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq15 is a free data retrieval call binding the contract method 0x9507540e.
//
// Solidity: function assertNotEq(bytes left, bytes right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq15(opts *bind.CallOpts, left []byte, right []byte, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq15", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq15 is a free data retrieval call binding the contract method 0x9507540e.
//
// Solidity: function assertNotEq(bytes left, bytes right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq15(left []byte, right []byte, error string) error {
	return _VmSafe.Contract.AssertNotEq15(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq15 is a free data retrieval call binding the contract method 0x9507540e.
//
// Solidity: function assertNotEq(bytes left, bytes right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq15(left []byte, right []byte, error string) error {
	return _VmSafe.Contract.AssertNotEq15(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq16 is a free data retrieval call binding the contract method 0x98f9bdbd.
//
// Solidity: function assertNotEq(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq16(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq16", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq16 is a free data retrieval call binding the contract method 0x98f9bdbd.
//
// Solidity: function assertNotEq(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq16(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEq16(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq16 is a free data retrieval call binding the contract method 0x98f9bdbd.
//
// Solidity: function assertNotEq(uint256 left, uint256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq16(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEq16(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq17 is a free data retrieval call binding the contract method 0x9a7fbd8f.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq17(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq17", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq17 is a free data retrieval call binding the contract method 0x9a7fbd8f.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq17(left []*big.Int, right []*big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEq17(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq17 is a free data retrieval call binding the contract method 0x9a7fbd8f.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq17(left []*big.Int, right []*big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEq17(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq18 is a free data retrieval call binding the contract method 0xb12e1694.
//
// Solidity: function assertNotEq(address left, address right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq18(opts *bind.CallOpts, left common.Address, right common.Address) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq18", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq18 is a free data retrieval call binding the contract method 0xb12e1694.
//
// Solidity: function assertNotEq(address left, address right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq18(left common.Address, right common.Address) error {
	return _VmSafe.Contract.AssertNotEq18(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq18 is a free data retrieval call binding the contract method 0xb12e1694.
//
// Solidity: function assertNotEq(address left, address right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq18(left common.Address, right common.Address) error {
	return _VmSafe.Contract.AssertNotEq18(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq19 is a free data retrieval call binding the contract method 0xb2332f51.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq19(opts *bind.CallOpts, left [32]byte, right [32]byte, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq19", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq19 is a free data retrieval call binding the contract method 0xb2332f51.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq19(left [32]byte, right [32]byte, error string) error {
	return _VmSafe.Contract.AssertNotEq19(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq19 is a free data retrieval call binding the contract method 0xb2332f51.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq19(left [32]byte, right [32]byte, error string) error {
	return _VmSafe.Contract.AssertNotEq19(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq2 is a free data retrieval call binding the contract method 0x1dcd1f68.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq2(opts *bind.CallOpts, left [][]byte, right [][]byte, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq2 is a free data retrieval call binding the contract method 0x1dcd1f68.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq2(left [][]byte, right [][]byte, error string) error {
	return _VmSafe.Contract.AssertNotEq2(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq2 is a free data retrieval call binding the contract method 0x1dcd1f68.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq2(left [][]byte, right [][]byte, error string) error {
	return _VmSafe.Contract.AssertNotEq2(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq20 is a free data retrieval call binding the contract method 0xb67187f3.
//
// Solidity: function assertNotEq(string[] left, string[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq20(opts *bind.CallOpts, left []string, right []string, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq20", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq20 is a free data retrieval call binding the contract method 0xb67187f3.
//
// Solidity: function assertNotEq(string[] left, string[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq20(left []string, right []string, error string) error {
	return _VmSafe.Contract.AssertNotEq20(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq20 is a free data retrieval call binding the contract method 0xb67187f3.
//
// Solidity: function assertNotEq(string[] left, string[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq20(left []string, right []string, error string) error {
	return _VmSafe.Contract.AssertNotEq20(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq21 is a free data retrieval call binding the contract method 0xb7909320.
//
// Solidity: function assertNotEq(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq21(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq21", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq21 is a free data retrieval call binding the contract method 0xb7909320.
//
// Solidity: function assertNotEq(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq21(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertNotEq21(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq21 is a free data retrieval call binding the contract method 0xb7909320.
//
// Solidity: function assertNotEq(uint256 left, uint256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq21(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertNotEq21(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq22 is a free data retrieval call binding the contract method 0xb873634c.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq22(opts *bind.CallOpts, left [][32]byte, right [][32]byte, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq22", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq22 is a free data retrieval call binding the contract method 0xb873634c.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq22(left [][32]byte, right [][32]byte, error string) error {
	return _VmSafe.Contract.AssertNotEq22(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq22 is a free data retrieval call binding the contract method 0xb873634c.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq22(left [][32]byte, right [][32]byte, error string) error {
	return _VmSafe.Contract.AssertNotEq22(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq23 is a free data retrieval call binding the contract method 0xbdfacbe8.
//
// Solidity: function assertNotEq(string[] left, string[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq23(opts *bind.CallOpts, left []string, right []string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq23", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq23 is a free data retrieval call binding the contract method 0xbdfacbe8.
//
// Solidity: function assertNotEq(string[] left, string[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq23(left []string, right []string) error {
	return _VmSafe.Contract.AssertNotEq23(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq23 is a free data retrieval call binding the contract method 0xbdfacbe8.
//
// Solidity: function assertNotEq(string[] left, string[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq23(left []string, right []string) error {
	return _VmSafe.Contract.AssertNotEq23(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq24 is a free data retrieval call binding the contract method 0xd3977322.
//
// Solidity: function assertNotEq(int256[] left, int256[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq24(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq24", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq24 is a free data retrieval call binding the contract method 0xd3977322.
//
// Solidity: function assertNotEq(int256[] left, int256[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq24(left []*big.Int, right []*big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEq24(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq24 is a free data retrieval call binding the contract method 0xd3977322.
//
// Solidity: function assertNotEq(int256[] left, int256[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq24(left []*big.Int, right []*big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEq24(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq25 is a free data retrieval call binding the contract method 0xedecd035.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq25(opts *bind.CallOpts, left [][]byte, right [][]byte) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq25", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq25 is a free data retrieval call binding the contract method 0xedecd035.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq25(left [][]byte, right [][]byte) error {
	return _VmSafe.Contract.AssertNotEq25(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq25 is a free data retrieval call binding the contract method 0xedecd035.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq25(left [][]byte, right [][]byte) error {
	return _VmSafe.Contract.AssertNotEq25(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq26 is a free data retrieval call binding the contract method 0xf4c004e3.
//
// Solidity: function assertNotEq(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq26(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq26", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq26 is a free data retrieval call binding the contract method 0xf4c004e3.
//
// Solidity: function assertNotEq(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq26(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertNotEq26(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq26 is a free data retrieval call binding the contract method 0xf4c004e3.
//
// Solidity: function assertNotEq(int256 left, int256 right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq26(left *big.Int, right *big.Int) error {
	return _VmSafe.Contract.AssertNotEq26(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq3 is a free data retrieval call binding the contract method 0x236e4d66.
//
// Solidity: function assertNotEq(bool left, bool right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq3(opts *bind.CallOpts, left bool, right bool) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq3", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq3 is a free data retrieval call binding the contract method 0x236e4d66.
//
// Solidity: function assertNotEq(bool left, bool right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq3(left bool, right bool) error {
	return _VmSafe.Contract.AssertNotEq3(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq3 is a free data retrieval call binding the contract method 0x236e4d66.
//
// Solidity: function assertNotEq(bool left, bool right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq3(left bool, right bool) error {
	return _VmSafe.Contract.AssertNotEq3(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq4 is a free data retrieval call binding the contract method 0x286fafea.
//
// Solidity: function assertNotEq(bool[] left, bool[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq4(opts *bind.CallOpts, left []bool, right []bool) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq4", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq4 is a free data retrieval call binding the contract method 0x286fafea.
//
// Solidity: function assertNotEq(bool[] left, bool[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq4(left []bool, right []bool) error {
	return _VmSafe.Contract.AssertNotEq4(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq4 is a free data retrieval call binding the contract method 0x286fafea.
//
// Solidity: function assertNotEq(bool[] left, bool[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq4(left []bool, right []bool) error {
	return _VmSafe.Contract.AssertNotEq4(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq5 is a free data retrieval call binding the contract method 0x3cf78e28.
//
// Solidity: function assertNotEq(bytes left, bytes right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq5(opts *bind.CallOpts, left []byte, right []byte) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq5", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq5 is a free data retrieval call binding the contract method 0x3cf78e28.
//
// Solidity: function assertNotEq(bytes left, bytes right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq5(left []byte, right []byte) error {
	return _VmSafe.Contract.AssertNotEq5(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq5 is a free data retrieval call binding the contract method 0x3cf78e28.
//
// Solidity: function assertNotEq(bytes left, bytes right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq5(left []byte, right []byte) error {
	return _VmSafe.Contract.AssertNotEq5(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq6 is a free data retrieval call binding the contract method 0x46d0b252.
//
// Solidity: function assertNotEq(address[] left, address[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq6(opts *bind.CallOpts, left []common.Address, right []common.Address) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq6", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq6 is a free data retrieval call binding the contract method 0x46d0b252.
//
// Solidity: function assertNotEq(address[] left, address[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq6(left []common.Address, right []common.Address) error {
	return _VmSafe.Contract.AssertNotEq6(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq6 is a free data retrieval call binding the contract method 0x46d0b252.
//
// Solidity: function assertNotEq(address[] left, address[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq6(left []common.Address, right []common.Address) error {
	return _VmSafe.Contract.AssertNotEq6(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq7 is a free data retrieval call binding the contract method 0x4724c5b9.
//
// Solidity: function assertNotEq(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq7(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq7", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq7 is a free data retrieval call binding the contract method 0x4724c5b9.
//
// Solidity: function assertNotEq(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq7(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEq7(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq7 is a free data retrieval call binding the contract method 0x4724c5b9.
//
// Solidity: function assertNotEq(int256 left, int256 right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq7(left *big.Int, right *big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEq7(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq8 is a free data retrieval call binding the contract method 0x56f29cba.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq8(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq8", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq8 is a free data retrieval call binding the contract method 0x56f29cba.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq8(left []*big.Int, right []*big.Int) error {
	return _VmSafe.Contract.AssertNotEq8(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq8 is a free data retrieval call binding the contract method 0x56f29cba.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq8(left []*big.Int, right []*big.Int) error {
	return _VmSafe.Contract.AssertNotEq8(&_VmSafe.CallOpts, left, right)
}

// AssertNotEq9 is a free data retrieval call binding the contract method 0x62c6f9fb.
//
// Solidity: function assertNotEq(bool[] left, bool[] right, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEq9(opts *bind.CallOpts, left []bool, right []bool, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEq9", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq9 is a free data retrieval call binding the contract method 0x62c6f9fb.
//
// Solidity: function assertNotEq(bool[] left, bool[] right, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEq9(left []bool, right []bool, error string) error {
	return _VmSafe.Contract.AssertNotEq9(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEq9 is a free data retrieval call binding the contract method 0x62c6f9fb.
//
// Solidity: function assertNotEq(bool[] left, bool[] right, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEq9(left []bool, right []bool, error string) error {
	return _VmSafe.Contract.AssertNotEq9(&_VmSafe.CallOpts, left, right, error)
}

// AssertNotEqDecimal is a free data retrieval call binding the contract method 0x14e75680.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEqDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEqDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal is a free data retrieval call binding the contract method 0x14e75680.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertNotEqDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal is a free data retrieval call binding the contract method 0x14e75680.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertNotEqDecimal(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal0 is a free data retrieval call binding the contract method 0x33949f0b.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEqDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEqDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal0 is a free data retrieval call binding the contract method 0x33949f0b.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEqDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertNotEqDecimal0 is a free data retrieval call binding the contract method 0x33949f0b.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEqDecimal0(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertNotEqDecimal1 is a free data retrieval call binding the contract method 0x669efca7.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEqDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEqDecimal1", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal1 is a free data retrieval call binding the contract method 0x669efca7.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertNotEqDecimal1(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal1 is a free data retrieval call binding the contract method 0x669efca7.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _VmSafe.Contract.AssertNotEqDecimal1(&_VmSafe.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal2 is a free data retrieval call binding the contract method 0xf5a55558.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertNotEqDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertNotEqDecimal2", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal2 is a free data retrieval call binding the contract method 0xf5a55558.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertNotEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEqDecimal2(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertNotEqDecimal2 is a free data retrieval call binding the contract method 0xf5a55558.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertNotEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _VmSafe.Contract.AssertNotEqDecimal2(&_VmSafe.CallOpts, left, right, decimals, error)
}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_VmSafe *VmSafeCaller) AssertTrue(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertTrue", condition)

	if err != nil {
		return err
	}

	return err

}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_VmSafe *VmSafeSession) AssertTrue(condition bool) error {
	return _VmSafe.Contract.AssertTrue(&_VmSafe.CallOpts, condition)
}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertTrue(condition bool) error {
	return _VmSafe.Contract.AssertTrue(&_VmSafe.CallOpts, condition)
}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string error) pure returns()
func (_VmSafe *VmSafeCaller) AssertTrue0(opts *bind.CallOpts, condition bool, error string) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assertTrue0", condition, error)

	if err != nil {
		return err
	}

	return err

}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string error) pure returns()
func (_VmSafe *VmSafeSession) AssertTrue0(condition bool, error string) error {
	return _VmSafe.Contract.AssertTrue0(&_VmSafe.CallOpts, condition, error)
}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string error) pure returns()
func (_VmSafe *VmSafeCallerSession) AssertTrue0(condition bool, error string) error {
	return _VmSafe.Contract.AssertTrue0(&_VmSafe.CallOpts, condition, error)
}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_VmSafe *VmSafeCaller) Assume(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "assume", condition)

	if err != nil {
		return err
	}

	return err

}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_VmSafe *VmSafeSession) Assume(condition bool) error {
	return _VmSafe.Contract.Assume(&_VmSafe.CallOpts, condition)
}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_VmSafe *VmSafeCallerSession) Assume(condition bool) error {
	return _VmSafe.Contract.Assume(&_VmSafe.CallOpts, condition)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_VmSafe *VmSafeCaller) ComputeCreate2Address(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "computeCreate2Address", salt, initCodeHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_VmSafe *VmSafeSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _VmSafe.Contract.ComputeCreate2Address(&_VmSafe.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_VmSafe *VmSafeCallerSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _VmSafe.Contract.ComputeCreate2Address(&_VmSafe.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_VmSafe *VmSafeCaller) ComputeCreate2Address0(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "computeCreate2Address0", salt, initCodeHash, deployer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_VmSafe *VmSafeSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _VmSafe.Contract.ComputeCreate2Address0(&_VmSafe.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_VmSafe *VmSafeCallerSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _VmSafe.Contract.ComputeCreate2Address0(&_VmSafe.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_VmSafe *VmSafeCaller) ComputeCreateAddress(opts *bind.CallOpts, deployer common.Address, nonce *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "computeCreateAddress", deployer, nonce)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_VmSafe *VmSafeSession) ComputeCreateAddress(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _VmSafe.Contract.ComputeCreateAddress(&_VmSafe.CallOpts, deployer, nonce)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_VmSafe *VmSafeCallerSession) ComputeCreateAddress(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _VmSafe.Contract.ComputeCreateAddress(&_VmSafe.CallOpts, deployer, nonce)
}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeCaller) DeriveKey(opts *bind.CallOpts, mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "deriveKey", mnemonic, derivationPath, index, language)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeSession) DeriveKey(mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	return _VmSafe.Contract.DeriveKey(&_VmSafe.CallOpts, mnemonic, derivationPath, index, language)
}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeCallerSession) DeriveKey(mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	return _VmSafe.Contract.DeriveKey(&_VmSafe.CallOpts, mnemonic, derivationPath, index, language)
}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeCaller) DeriveKey0(opts *bind.CallOpts, mnemonic string, index uint32, language string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "deriveKey0", mnemonic, index, language)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeSession) DeriveKey0(mnemonic string, index uint32, language string) (*big.Int, error) {
	return _VmSafe.Contract.DeriveKey0(&_VmSafe.CallOpts, mnemonic, index, language)
}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeCallerSession) DeriveKey0(mnemonic string, index uint32, language string) (*big.Int, error) {
	return _VmSafe.Contract.DeriveKey0(&_VmSafe.CallOpts, mnemonic, index, language)
}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeCaller) DeriveKey1(opts *bind.CallOpts, mnemonic string, index uint32) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "deriveKey1", mnemonic, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeSession) DeriveKey1(mnemonic string, index uint32) (*big.Int, error) {
	return _VmSafe.Contract.DeriveKey1(&_VmSafe.CallOpts, mnemonic, index)
}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeCallerSession) DeriveKey1(mnemonic string, index uint32) (*big.Int, error) {
	return _VmSafe.Contract.DeriveKey1(&_VmSafe.CallOpts, mnemonic, index)
}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeCaller) DeriveKey2(opts *bind.CallOpts, mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "deriveKey2", mnemonic, derivationPath, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeSession) DeriveKey2(mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	return _VmSafe.Contract.DeriveKey2(&_VmSafe.CallOpts, mnemonic, derivationPath, index)
}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_VmSafe *VmSafeCallerSession) DeriveKey2(mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	return _VmSafe.Contract.DeriveKey2(&_VmSafe.CallOpts, mnemonic, derivationPath, index)
}

// EnsNamehash is a free data retrieval call binding the contract method 0x8c374c65.
//
// Solidity: function ensNamehash(string name) pure returns(bytes32)
func (_VmSafe *VmSafeCaller) EnsNamehash(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "ensNamehash", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnsNamehash is a free data retrieval call binding the contract method 0x8c374c65.
//
// Solidity: function ensNamehash(string name) pure returns(bytes32)
func (_VmSafe *VmSafeSession) EnsNamehash(name string) ([32]byte, error) {
	return _VmSafe.Contract.EnsNamehash(&_VmSafe.CallOpts, name)
}

// EnsNamehash is a free data retrieval call binding the contract method 0x8c374c65.
//
// Solidity: function ensNamehash(string name) pure returns(bytes32)
func (_VmSafe *VmSafeCallerSession) EnsNamehash(name string) ([32]byte, error) {
	return _VmSafe.Contract.EnsNamehash(&_VmSafe.CallOpts, name)
}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_VmSafe *VmSafeCaller) EnvAddress(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_VmSafe *VmSafeSession) EnvAddress(name string) (common.Address, error) {
	return _VmSafe.Contract.EnvAddress(&_VmSafe.CallOpts, name)
}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_VmSafe *VmSafeCallerSession) EnvAddress(name string) (common.Address, error) {
	return _VmSafe.Contract.EnvAddress(&_VmSafe.CallOpts, name)
}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_VmSafe *VmSafeCaller) EnvAddress0(opts *bind.CallOpts, name string, delim string) ([]common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envAddress0", name, delim)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_VmSafe *VmSafeSession) EnvAddress0(name string, delim string) ([]common.Address, error) {
	return _VmSafe.Contract.EnvAddress0(&_VmSafe.CallOpts, name, delim)
}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_VmSafe *VmSafeCallerSession) EnvAddress0(name string, delim string) ([]common.Address, error) {
	return _VmSafe.Contract.EnvAddress0(&_VmSafe.CallOpts, name, delim)
}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_VmSafe *VmSafeCaller) EnvBool(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envBool", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_VmSafe *VmSafeSession) EnvBool(name string) (bool, error) {
	return _VmSafe.Contract.EnvBool(&_VmSafe.CallOpts, name)
}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_VmSafe *VmSafeCallerSession) EnvBool(name string) (bool, error) {
	return _VmSafe.Contract.EnvBool(&_VmSafe.CallOpts, name)
}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_VmSafe *VmSafeCaller) EnvBool0(opts *bind.CallOpts, name string, delim string) ([]bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envBool0", name, delim)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_VmSafe *VmSafeSession) EnvBool0(name string, delim string) ([]bool, error) {
	return _VmSafe.Contract.EnvBool0(&_VmSafe.CallOpts, name, delim)
}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_VmSafe *VmSafeCallerSession) EnvBool0(name string, delim string) ([]bool, error) {
	return _VmSafe.Contract.EnvBool0(&_VmSafe.CallOpts, name, delim)
}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_VmSafe *VmSafeCaller) EnvBytes(opts *bind.CallOpts, name string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envBytes", name)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_VmSafe *VmSafeSession) EnvBytes(name string) ([]byte, error) {
	return _VmSafe.Contract.EnvBytes(&_VmSafe.CallOpts, name)
}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_VmSafe *VmSafeCallerSession) EnvBytes(name string) ([]byte, error) {
	return _VmSafe.Contract.EnvBytes(&_VmSafe.CallOpts, name)
}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_VmSafe *VmSafeCaller) EnvBytes0(opts *bind.CallOpts, name string, delim string) ([][]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envBytes0", name, delim)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_VmSafe *VmSafeSession) EnvBytes0(name string, delim string) ([][]byte, error) {
	return _VmSafe.Contract.EnvBytes0(&_VmSafe.CallOpts, name, delim)
}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_VmSafe *VmSafeCallerSession) EnvBytes0(name string, delim string) ([][]byte, error) {
	return _VmSafe.Contract.EnvBytes0(&_VmSafe.CallOpts, name, delim)
}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_VmSafe *VmSafeCaller) EnvBytes32(opts *bind.CallOpts, name string, delim string) ([][32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envBytes32", name, delim)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_VmSafe *VmSafeSession) EnvBytes32(name string, delim string) ([][32]byte, error) {
	return _VmSafe.Contract.EnvBytes32(&_VmSafe.CallOpts, name, delim)
}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_VmSafe *VmSafeCallerSession) EnvBytes32(name string, delim string) ([][32]byte, error) {
	return _VmSafe.Contract.EnvBytes32(&_VmSafe.CallOpts, name, delim)
}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_VmSafe *VmSafeCaller) EnvBytes320(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envBytes320", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_VmSafe *VmSafeSession) EnvBytes320(name string) ([32]byte, error) {
	return _VmSafe.Contract.EnvBytes320(&_VmSafe.CallOpts, name)
}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_VmSafe *VmSafeCallerSession) EnvBytes320(name string) ([32]byte, error) {
	return _VmSafe.Contract.EnvBytes320(&_VmSafe.CallOpts, name)
}

// EnvExists is a free data retrieval call binding the contract method 0xce8365f9.
//
// Solidity: function envExists(string name) view returns(bool result)
func (_VmSafe *VmSafeCaller) EnvExists(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envExists", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvExists is a free data retrieval call binding the contract method 0xce8365f9.
//
// Solidity: function envExists(string name) view returns(bool result)
func (_VmSafe *VmSafeSession) EnvExists(name string) (bool, error) {
	return _VmSafe.Contract.EnvExists(&_VmSafe.CallOpts, name)
}

// EnvExists is a free data retrieval call binding the contract method 0xce8365f9.
//
// Solidity: function envExists(string name) view returns(bool result)
func (_VmSafe *VmSafeCallerSession) EnvExists(name string) (bool, error) {
	return _VmSafe.Contract.EnvExists(&_VmSafe.CallOpts, name)
}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_VmSafe *VmSafeCaller) EnvInt(opts *bind.CallOpts, name string, delim string) ([]*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envInt", name, delim)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_VmSafe *VmSafeSession) EnvInt(name string, delim string) ([]*big.Int, error) {
	return _VmSafe.Contract.EnvInt(&_VmSafe.CallOpts, name, delim)
}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_VmSafe *VmSafeCallerSession) EnvInt(name string, delim string) ([]*big.Int, error) {
	return _VmSafe.Contract.EnvInt(&_VmSafe.CallOpts, name, delim)
}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_VmSafe *VmSafeCaller) EnvInt0(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envInt0", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_VmSafe *VmSafeSession) EnvInt0(name string) (*big.Int, error) {
	return _VmSafe.Contract.EnvInt0(&_VmSafe.CallOpts, name)
}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_VmSafe *VmSafeCallerSession) EnvInt0(name string) (*big.Int, error) {
	return _VmSafe.Contract.EnvInt0(&_VmSafe.CallOpts, name)
}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_VmSafe *VmSafeCaller) EnvOr(opts *bind.CallOpts, name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr", name, delim, defaultValue)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_VmSafe *VmSafeSession) EnvOr(name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	return _VmSafe.Contract.EnvOr(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_VmSafe *VmSafeCallerSession) EnvOr(name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	return _VmSafe.Contract.EnvOr(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_VmSafe *VmSafeCaller) EnvOr0(opts *bind.CallOpts, name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr0", name, delim, defaultValue)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_VmSafe *VmSafeSession) EnvOr0(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _VmSafe.Contract.EnvOr0(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_VmSafe *VmSafeCallerSession) EnvOr0(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _VmSafe.Contract.EnvOr0(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_VmSafe *VmSafeCaller) EnvOr1(opts *bind.CallOpts, name string, defaultValue bool) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr1", name, defaultValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_VmSafe *VmSafeSession) EnvOr1(name string, defaultValue bool) (bool, error) {
	return _VmSafe.Contract.EnvOr1(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_VmSafe *VmSafeCallerSession) EnvOr1(name string, defaultValue bool) (bool, error) {
	return _VmSafe.Contract.EnvOr1(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_VmSafe *VmSafeCaller) EnvOr10(opts *bind.CallOpts, name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr10", name, delim, defaultValue)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_VmSafe *VmSafeSession) EnvOr10(name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	return _VmSafe.Contract.EnvOr10(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_VmSafe *VmSafeCallerSession) EnvOr10(name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	return _VmSafe.Contract.EnvOr10(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_VmSafe *VmSafeCaller) EnvOr11(opts *bind.CallOpts, name string, defaultValue string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr11", name, defaultValue)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_VmSafe *VmSafeSession) EnvOr11(name string, defaultValue string) (string, error) {
	return _VmSafe.Contract.EnvOr11(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_VmSafe *VmSafeCallerSession) EnvOr11(name string, defaultValue string) (string, error) {
	return _VmSafe.Contract.EnvOr11(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_VmSafe *VmSafeCaller) EnvOr12(opts *bind.CallOpts, name string, delim string, defaultValue []bool) ([]bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr12", name, delim, defaultValue)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_VmSafe *VmSafeSession) EnvOr12(name string, delim string, defaultValue []bool) ([]bool, error) {
	return _VmSafe.Contract.EnvOr12(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_VmSafe *VmSafeCallerSession) EnvOr12(name string, delim string, defaultValue []bool) ([]bool, error) {
	return _VmSafe.Contract.EnvOr12(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_VmSafe *VmSafeCaller) EnvOr2(opts *bind.CallOpts, name string, defaultValue common.Address) (common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr2", name, defaultValue)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_VmSafe *VmSafeSession) EnvOr2(name string, defaultValue common.Address) (common.Address, error) {
	return _VmSafe.Contract.EnvOr2(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_VmSafe *VmSafeCallerSession) EnvOr2(name string, defaultValue common.Address) (common.Address, error) {
	return _VmSafe.Contract.EnvOr2(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_VmSafe *VmSafeCaller) EnvOr3(opts *bind.CallOpts, name string, defaultValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr3", name, defaultValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_VmSafe *VmSafeSession) EnvOr3(name string, defaultValue *big.Int) (*big.Int, error) {
	return _VmSafe.Contract.EnvOr3(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_VmSafe *VmSafeCallerSession) EnvOr3(name string, defaultValue *big.Int) (*big.Int, error) {
	return _VmSafe.Contract.EnvOr3(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_VmSafe *VmSafeCaller) EnvOr4(opts *bind.CallOpts, name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr4", name, delim, defaultValue)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_VmSafe *VmSafeSession) EnvOr4(name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	return _VmSafe.Contract.EnvOr4(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_VmSafe *VmSafeCallerSession) EnvOr4(name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	return _VmSafe.Contract.EnvOr4(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_VmSafe *VmSafeCaller) EnvOr5(opts *bind.CallOpts, name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr5", name, delim, defaultValue)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_VmSafe *VmSafeSession) EnvOr5(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _VmSafe.Contract.EnvOr5(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_VmSafe *VmSafeCallerSession) EnvOr5(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _VmSafe.Contract.EnvOr5(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_VmSafe *VmSafeCaller) EnvOr6(opts *bind.CallOpts, name string, delim string, defaultValue []string) ([]string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr6", name, delim, defaultValue)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_VmSafe *VmSafeSession) EnvOr6(name string, delim string, defaultValue []string) ([]string, error) {
	return _VmSafe.Contract.EnvOr6(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_VmSafe *VmSafeCallerSession) EnvOr6(name string, delim string, defaultValue []string) ([]string, error) {
	return _VmSafe.Contract.EnvOr6(&_VmSafe.CallOpts, name, delim, defaultValue)
}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_VmSafe *VmSafeCaller) EnvOr7(opts *bind.CallOpts, name string, defaultValue []byte) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr7", name, defaultValue)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_VmSafe *VmSafeSession) EnvOr7(name string, defaultValue []byte) ([]byte, error) {
	return _VmSafe.Contract.EnvOr7(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_VmSafe *VmSafeCallerSession) EnvOr7(name string, defaultValue []byte) ([]byte, error) {
	return _VmSafe.Contract.EnvOr7(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_VmSafe *VmSafeCaller) EnvOr8(opts *bind.CallOpts, name string, defaultValue [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr8", name, defaultValue)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_VmSafe *VmSafeSession) EnvOr8(name string, defaultValue [32]byte) ([32]byte, error) {
	return _VmSafe.Contract.EnvOr8(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_VmSafe *VmSafeCallerSession) EnvOr8(name string, defaultValue [32]byte) ([32]byte, error) {
	return _VmSafe.Contract.EnvOr8(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_VmSafe *VmSafeCaller) EnvOr9(opts *bind.CallOpts, name string, defaultValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envOr9", name, defaultValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_VmSafe *VmSafeSession) EnvOr9(name string, defaultValue *big.Int) (*big.Int, error) {
	return _VmSafe.Contract.EnvOr9(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_VmSafe *VmSafeCallerSession) EnvOr9(name string, defaultValue *big.Int) (*big.Int, error) {
	return _VmSafe.Contract.EnvOr9(&_VmSafe.CallOpts, name, defaultValue)
}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_VmSafe *VmSafeCaller) EnvString(opts *bind.CallOpts, name string, delim string) ([]string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envString", name, delim)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_VmSafe *VmSafeSession) EnvString(name string, delim string) ([]string, error) {
	return _VmSafe.Contract.EnvString(&_VmSafe.CallOpts, name, delim)
}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_VmSafe *VmSafeCallerSession) EnvString(name string, delim string) ([]string, error) {
	return _VmSafe.Contract.EnvString(&_VmSafe.CallOpts, name, delim)
}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_VmSafe *VmSafeCaller) EnvString0(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envString0", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_VmSafe *VmSafeSession) EnvString0(name string) (string, error) {
	return _VmSafe.Contract.EnvString0(&_VmSafe.CallOpts, name)
}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_VmSafe *VmSafeCallerSession) EnvString0(name string) (string, error) {
	return _VmSafe.Contract.EnvString0(&_VmSafe.CallOpts, name)
}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_VmSafe *VmSafeCaller) EnvUint(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envUint", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_VmSafe *VmSafeSession) EnvUint(name string) (*big.Int, error) {
	return _VmSafe.Contract.EnvUint(&_VmSafe.CallOpts, name)
}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_VmSafe *VmSafeCallerSession) EnvUint(name string) (*big.Int, error) {
	return _VmSafe.Contract.EnvUint(&_VmSafe.CallOpts, name)
}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_VmSafe *VmSafeCaller) EnvUint0(opts *bind.CallOpts, name string, delim string) ([]*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "envUint0", name, delim)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_VmSafe *VmSafeSession) EnvUint0(name string, delim string) ([]*big.Int, error) {
	return _VmSafe.Contract.EnvUint0(&_VmSafe.CallOpts, name, delim)
}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_VmSafe *VmSafeCallerSession) EnvUint0(name string, delim string) ([]*big.Int, error) {
	return _VmSafe.Contract.EnvUint0(&_VmSafe.CallOpts, name, delim)
}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_VmSafe *VmSafeCaller) FsMetadata(opts *bind.CallOpts, path string) (VmSafeFsMetadata, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "fsMetadata", path)

	if err != nil {
		return *new(VmSafeFsMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeFsMetadata)).(*VmSafeFsMetadata)

	return out0, err

}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_VmSafe *VmSafeSession) FsMetadata(path string) (VmSafeFsMetadata, error) {
	return _VmSafe.Contract.FsMetadata(&_VmSafe.CallOpts, path)
}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_VmSafe *VmSafeCallerSession) FsMetadata(path string) (VmSafeFsMetadata, error) {
	return _VmSafe.Contract.FsMetadata(&_VmSafe.CallOpts, path)
}

// GetBlobBaseFee is a free data retrieval call binding the contract method 0x1f6d6ef7.
//
// Solidity: function getBlobBaseFee() view returns(uint256 blobBaseFee)
func (_VmSafe *VmSafeCaller) GetBlobBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "getBlobBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlobBaseFee is a free data retrieval call binding the contract method 0x1f6d6ef7.
//
// Solidity: function getBlobBaseFee() view returns(uint256 blobBaseFee)
func (_VmSafe *VmSafeSession) GetBlobBaseFee() (*big.Int, error) {
	return _VmSafe.Contract.GetBlobBaseFee(&_VmSafe.CallOpts)
}

// GetBlobBaseFee is a free data retrieval call binding the contract method 0x1f6d6ef7.
//
// Solidity: function getBlobBaseFee() view returns(uint256 blobBaseFee)
func (_VmSafe *VmSafeCallerSession) GetBlobBaseFee() (*big.Int, error) {
	return _VmSafe.Contract.GetBlobBaseFee(&_VmSafe.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_VmSafe *VmSafeCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_VmSafe *VmSafeSession) GetBlockNumber() (*big.Int, error) {
	return _VmSafe.Contract.GetBlockNumber(&_VmSafe.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_VmSafe *VmSafeCallerSession) GetBlockNumber() (*big.Int, error) {
	return _VmSafe.Contract.GetBlockNumber(&_VmSafe.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_VmSafe *VmSafeCaller) GetBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "getBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_VmSafe *VmSafeSession) GetBlockTimestamp() (*big.Int, error) {
	return _VmSafe.Contract.GetBlockTimestamp(&_VmSafe.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_VmSafe *VmSafeCallerSession) GetBlockTimestamp() (*big.Int, error) {
	return _VmSafe.Contract.GetBlockTimestamp(&_VmSafe.CallOpts)
}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_VmSafe *VmSafeCaller) GetCode(opts *bind.CallOpts, artifactPath string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "getCode", artifactPath)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_VmSafe *VmSafeSession) GetCode(artifactPath string) ([]byte, error) {
	return _VmSafe.Contract.GetCode(&_VmSafe.CallOpts, artifactPath)
}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_VmSafe *VmSafeCallerSession) GetCode(artifactPath string) ([]byte, error) {
	return _VmSafe.Contract.GetCode(&_VmSafe.CallOpts, artifactPath)
}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_VmSafe *VmSafeCaller) GetDeployedCode(opts *bind.CallOpts, artifactPath string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "getDeployedCode", artifactPath)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_VmSafe *VmSafeSession) GetDeployedCode(artifactPath string) ([]byte, error) {
	return _VmSafe.Contract.GetDeployedCode(&_VmSafe.CallOpts, artifactPath)
}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_VmSafe *VmSafeCallerSession) GetDeployedCode(artifactPath string) ([]byte, error) {
	return _VmSafe.Contract.GetDeployedCode(&_VmSafe.CallOpts, artifactPath)
}

// GetFoundryVersion is a free data retrieval call binding the contract method 0xea991bb5.
//
// Solidity: function getFoundryVersion() view returns(string version)
func (_VmSafe *VmSafeCaller) GetFoundryVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "getFoundryVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFoundryVersion is a free data retrieval call binding the contract method 0xea991bb5.
//
// Solidity: function getFoundryVersion() view returns(string version)
func (_VmSafe *VmSafeSession) GetFoundryVersion() (string, error) {
	return _VmSafe.Contract.GetFoundryVersion(&_VmSafe.CallOpts)
}

// GetFoundryVersion is a free data retrieval call binding the contract method 0xea991bb5.
//
// Solidity: function getFoundryVersion() view returns(string version)
func (_VmSafe *VmSafeCallerSession) GetFoundryVersion() (string, error) {
	return _VmSafe.Contract.GetFoundryVersion(&_VmSafe.CallOpts)
}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_VmSafe *VmSafeCaller) GetLabel(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "getLabel", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_VmSafe *VmSafeSession) GetLabel(account common.Address) (string, error) {
	return _VmSafe.Contract.GetLabel(&_VmSafe.CallOpts, account)
}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_VmSafe *VmSafeCallerSession) GetLabel(account common.Address) (string, error) {
	return _VmSafe.Contract.GetLabel(&_VmSafe.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_VmSafe *VmSafeCaller) GetNonce(opts *bind.CallOpts, account common.Address) (uint64, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "getNonce", account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_VmSafe *VmSafeSession) GetNonce(account common.Address) (uint64, error) {
	return _VmSafe.Contract.GetNonce(&_VmSafe.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_VmSafe *VmSafeCallerSession) GetNonce(account common.Address) (uint64, error) {
	return _VmSafe.Contract.GetNonce(&_VmSafe.CallOpts, account)
}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string input, string key) pure returns(uint256)
func (_VmSafe *VmSafeCaller) IndexOf(opts *bind.CallOpts, input string, key string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "indexOf", input, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string input, string key) pure returns(uint256)
func (_VmSafe *VmSafeSession) IndexOf(input string, key string) (*big.Int, error) {
	return _VmSafe.Contract.IndexOf(&_VmSafe.CallOpts, input, key)
}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string input, string key) pure returns(uint256)
func (_VmSafe *VmSafeCallerSession) IndexOf(input string, key string) (*big.Int, error) {
	return _VmSafe.Contract.IndexOf(&_VmSafe.CallOpts, input, key)
}

// IsContext is a free data retrieval call binding the contract method 0x64af255d.
//
// Solidity: function isContext(uint8 context) view returns(bool result)
func (_VmSafe *VmSafeCaller) IsContext(opts *bind.CallOpts, context uint8) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "isContext", context)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContext is a free data retrieval call binding the contract method 0x64af255d.
//
// Solidity: function isContext(uint8 context) view returns(bool result)
func (_VmSafe *VmSafeSession) IsContext(context uint8) (bool, error) {
	return _VmSafe.Contract.IsContext(&_VmSafe.CallOpts, context)
}

// IsContext is a free data retrieval call binding the contract method 0x64af255d.
//
// Solidity: function isContext(uint8 context) view returns(bool result)
func (_VmSafe *VmSafeCallerSession) IsContext(context uint8) (bool, error) {
	return _VmSafe.Contract.IsContext(&_VmSafe.CallOpts, context)
}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_VmSafe *VmSafeCaller) KeyExists(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "keyExists", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_VmSafe *VmSafeSession) KeyExists(json string, key string) (bool, error) {
	return _VmSafe.Contract.KeyExists(&_VmSafe.CallOpts, json, key)
}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_VmSafe *VmSafeCallerSession) KeyExists(json string, key string) (bool, error) {
	return _VmSafe.Contract.KeyExists(&_VmSafe.CallOpts, json, key)
}

// KeyExistsJson is a free data retrieval call binding the contract method 0xdb4235f6.
//
// Solidity: function keyExistsJson(string json, string key) view returns(bool)
func (_VmSafe *VmSafeCaller) KeyExistsJson(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "keyExistsJson", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExistsJson is a free data retrieval call binding the contract method 0xdb4235f6.
//
// Solidity: function keyExistsJson(string json, string key) view returns(bool)
func (_VmSafe *VmSafeSession) KeyExistsJson(json string, key string) (bool, error) {
	return _VmSafe.Contract.KeyExistsJson(&_VmSafe.CallOpts, json, key)
}

// KeyExistsJson is a free data retrieval call binding the contract method 0xdb4235f6.
//
// Solidity: function keyExistsJson(string json, string key) view returns(bool)
func (_VmSafe *VmSafeCallerSession) KeyExistsJson(json string, key string) (bool, error) {
	return _VmSafe.Contract.KeyExistsJson(&_VmSafe.CallOpts, json, key)
}

// KeyExistsToml is a free data retrieval call binding the contract method 0x600903ad.
//
// Solidity: function keyExistsToml(string toml, string key) view returns(bool)
func (_VmSafe *VmSafeCaller) KeyExistsToml(opts *bind.CallOpts, toml string, key string) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "keyExistsToml", toml, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExistsToml is a free data retrieval call binding the contract method 0x600903ad.
//
// Solidity: function keyExistsToml(string toml, string key) view returns(bool)
func (_VmSafe *VmSafeSession) KeyExistsToml(toml string, key string) (bool, error) {
	return _VmSafe.Contract.KeyExistsToml(&_VmSafe.CallOpts, toml, key)
}

// KeyExistsToml is a free data retrieval call binding the contract method 0x600903ad.
//
// Solidity: function keyExistsToml(string toml, string key) view returns(bool)
func (_VmSafe *VmSafeCallerSession) KeyExistsToml(toml string, key string) (bool, error) {
	return _VmSafe.Contract.KeyExistsToml(&_VmSafe.CallOpts, toml, key)
}

// LastCallGas is a free data retrieval call binding the contract method 0x2b589b28.
//
// Solidity: function lastCallGas() view returns((uint64,uint64,uint64,int64,uint64) gas)
func (_VmSafe *VmSafeCaller) LastCallGas(opts *bind.CallOpts) (VmSafeGas, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "lastCallGas")

	if err != nil {
		return *new(VmSafeGas), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeGas)).(*VmSafeGas)

	return out0, err

}

// LastCallGas is a free data retrieval call binding the contract method 0x2b589b28.
//
// Solidity: function lastCallGas() view returns((uint64,uint64,uint64,int64,uint64) gas)
func (_VmSafe *VmSafeSession) LastCallGas() (VmSafeGas, error) {
	return _VmSafe.Contract.LastCallGas(&_VmSafe.CallOpts)
}

// LastCallGas is a free data retrieval call binding the contract method 0x2b589b28.
//
// Solidity: function lastCallGas() view returns((uint64,uint64,uint64,int64,uint64) gas)
func (_VmSafe *VmSafeCallerSession) LastCallGas() (VmSafeGas, error) {
	return _VmSafe.Contract.LastCallGas(&_VmSafe.CallOpts)
}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_VmSafe *VmSafeCaller) Load(opts *bind.CallOpts, target common.Address, slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "load", target, slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_VmSafe *VmSafeSession) Load(target common.Address, slot [32]byte) ([32]byte, error) {
	return _VmSafe.Contract.Load(&_VmSafe.CallOpts, target, slot)
}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_VmSafe *VmSafeCallerSession) Load(target common.Address, slot [32]byte) ([32]byte, error) {
	return _VmSafe.Contract.Load(&_VmSafe.CallOpts, target, slot)
}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_VmSafe *VmSafeCaller) ParseAddress(opts *bind.CallOpts, stringifiedValue string) (common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseAddress", stringifiedValue)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_VmSafe *VmSafeSession) ParseAddress(stringifiedValue string) (common.Address, error) {
	return _VmSafe.Contract.ParseAddress(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_VmSafe *VmSafeCallerSession) ParseAddress(stringifiedValue string) (common.Address, error) {
	return _VmSafe.Contract.ParseAddress(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_VmSafe *VmSafeCaller) ParseBool(opts *bind.CallOpts, stringifiedValue string) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseBool", stringifiedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_VmSafe *VmSafeSession) ParseBool(stringifiedValue string) (bool, error) {
	return _VmSafe.Contract.ParseBool(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_VmSafe *VmSafeCallerSession) ParseBool(stringifiedValue string) (bool, error) {
	return _VmSafe.Contract.ParseBool(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_VmSafe *VmSafeCaller) ParseBytes(opts *bind.CallOpts, stringifiedValue string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseBytes", stringifiedValue)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_VmSafe *VmSafeSession) ParseBytes(stringifiedValue string) ([]byte, error) {
	return _VmSafe.Contract.ParseBytes(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_VmSafe *VmSafeCallerSession) ParseBytes(stringifiedValue string) ([]byte, error) {
	return _VmSafe.Contract.ParseBytes(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_VmSafe *VmSafeCaller) ParseBytes32(opts *bind.CallOpts, stringifiedValue string) ([32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseBytes32", stringifiedValue)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_VmSafe *VmSafeSession) ParseBytes32(stringifiedValue string) ([32]byte, error) {
	return _VmSafe.Contract.ParseBytes32(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_VmSafe *VmSafeCallerSession) ParseBytes32(stringifiedValue string) ([32]byte, error) {
	return _VmSafe.Contract.ParseBytes32(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_VmSafe *VmSafeCaller) ParseInt(opts *bind.CallOpts, stringifiedValue string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseInt", stringifiedValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_VmSafe *VmSafeSession) ParseInt(stringifiedValue string) (*big.Int, error) {
	return _VmSafe.Contract.ParseInt(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_VmSafe *VmSafeCallerSession) ParseInt(stringifiedValue string) (*big.Int, error) {
	return _VmSafe.Contract.ParseInt(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeCaller) ParseJson(opts *bind.CallOpts, json string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJson", json)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeSession) ParseJson(json string) ([]byte, error) {
	return _VmSafe.Contract.ParseJson(&_VmSafe.CallOpts, json)
}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeCallerSession) ParseJson(json string) ([]byte, error) {
	return _VmSafe.Contract.ParseJson(&_VmSafe.CallOpts, json)
}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeCaller) ParseJson0(opts *bind.CallOpts, json string, key string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJson0", json, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeSession) ParseJson0(json string, key string) ([]byte, error) {
	return _VmSafe.Contract.ParseJson0(&_VmSafe.CallOpts, json, key)
}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeCallerSession) ParseJson0(json string, key string) ([]byte, error) {
	return _VmSafe.Contract.ParseJson0(&_VmSafe.CallOpts, json, key)
}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_VmSafe *VmSafeCaller) ParseJsonAddress(opts *bind.CallOpts, json string, key string) (common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonAddress", json, key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_VmSafe *VmSafeSession) ParseJsonAddress(json string, key string) (common.Address, error) {
	return _VmSafe.Contract.ParseJsonAddress(&_VmSafe.CallOpts, json, key)
}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_VmSafe *VmSafeCallerSession) ParseJsonAddress(json string, key string) (common.Address, error) {
	return _VmSafe.Contract.ParseJsonAddress(&_VmSafe.CallOpts, json, key)
}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_VmSafe *VmSafeCaller) ParseJsonAddressArray(opts *bind.CallOpts, json string, key string) ([]common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonAddressArray", json, key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_VmSafe *VmSafeSession) ParseJsonAddressArray(json string, key string) ([]common.Address, error) {
	return _VmSafe.Contract.ParseJsonAddressArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_VmSafe *VmSafeCallerSession) ParseJsonAddressArray(json string, key string) ([]common.Address, error) {
	return _VmSafe.Contract.ParseJsonAddressArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_VmSafe *VmSafeCaller) ParseJsonBool(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonBool", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_VmSafe *VmSafeSession) ParseJsonBool(json string, key string) (bool, error) {
	return _VmSafe.Contract.ParseJsonBool(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_VmSafe *VmSafeCallerSession) ParseJsonBool(json string, key string) (bool, error) {
	return _VmSafe.Contract.ParseJsonBool(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_VmSafe *VmSafeCaller) ParseJsonBoolArray(opts *bind.CallOpts, json string, key string) ([]bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonBoolArray", json, key)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_VmSafe *VmSafeSession) ParseJsonBoolArray(json string, key string) ([]bool, error) {
	return _VmSafe.Contract.ParseJsonBoolArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_VmSafe *VmSafeCallerSession) ParseJsonBoolArray(json string, key string) ([]bool, error) {
	return _VmSafe.Contract.ParseJsonBoolArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_VmSafe *VmSafeCaller) ParseJsonBytes(opts *bind.CallOpts, json string, key string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonBytes", json, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_VmSafe *VmSafeSession) ParseJsonBytes(json string, key string) ([]byte, error) {
	return _VmSafe.Contract.ParseJsonBytes(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_VmSafe *VmSafeCallerSession) ParseJsonBytes(json string, key string) ([]byte, error) {
	return _VmSafe.Contract.ParseJsonBytes(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_VmSafe *VmSafeCaller) ParseJsonBytes32(opts *bind.CallOpts, json string, key string) ([32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonBytes32", json, key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_VmSafe *VmSafeSession) ParseJsonBytes32(json string, key string) ([32]byte, error) {
	return _VmSafe.Contract.ParseJsonBytes32(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_VmSafe *VmSafeCallerSession) ParseJsonBytes32(json string, key string) ([32]byte, error) {
	return _VmSafe.Contract.ParseJsonBytes32(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_VmSafe *VmSafeCaller) ParseJsonBytes32Array(opts *bind.CallOpts, json string, key string) ([][32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonBytes32Array", json, key)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_VmSafe *VmSafeSession) ParseJsonBytes32Array(json string, key string) ([][32]byte, error) {
	return _VmSafe.Contract.ParseJsonBytes32Array(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_VmSafe *VmSafeCallerSession) ParseJsonBytes32Array(json string, key string) ([][32]byte, error) {
	return _VmSafe.Contract.ParseJsonBytes32Array(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_VmSafe *VmSafeCaller) ParseJsonBytesArray(opts *bind.CallOpts, json string, key string) ([][]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonBytesArray", json, key)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_VmSafe *VmSafeSession) ParseJsonBytesArray(json string, key string) ([][]byte, error) {
	return _VmSafe.Contract.ParseJsonBytesArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_VmSafe *VmSafeCallerSession) ParseJsonBytesArray(json string, key string) ([][]byte, error) {
	return _VmSafe.Contract.ParseJsonBytesArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_VmSafe *VmSafeCaller) ParseJsonInt(opts *bind.CallOpts, json string, key string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonInt", json, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_VmSafe *VmSafeSession) ParseJsonInt(json string, key string) (*big.Int, error) {
	return _VmSafe.Contract.ParseJsonInt(&_VmSafe.CallOpts, json, key)
}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_VmSafe *VmSafeCallerSession) ParseJsonInt(json string, key string) (*big.Int, error) {
	return _VmSafe.Contract.ParseJsonInt(&_VmSafe.CallOpts, json, key)
}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_VmSafe *VmSafeCaller) ParseJsonIntArray(opts *bind.CallOpts, json string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonIntArray", json, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_VmSafe *VmSafeSession) ParseJsonIntArray(json string, key string) ([]*big.Int, error) {
	return _VmSafe.Contract.ParseJsonIntArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_VmSafe *VmSafeCallerSession) ParseJsonIntArray(json string, key string) ([]*big.Int, error) {
	return _VmSafe.Contract.ParseJsonIntArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_VmSafe *VmSafeCaller) ParseJsonKeys(opts *bind.CallOpts, json string, key string) ([]string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonKeys", json, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_VmSafe *VmSafeSession) ParseJsonKeys(json string, key string) ([]string, error) {
	return _VmSafe.Contract.ParseJsonKeys(&_VmSafe.CallOpts, json, key)
}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_VmSafe *VmSafeCallerSession) ParseJsonKeys(json string, key string) ([]string, error) {
	return _VmSafe.Contract.ParseJsonKeys(&_VmSafe.CallOpts, json, key)
}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_VmSafe *VmSafeCaller) ParseJsonString(opts *bind.CallOpts, json string, key string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonString", json, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_VmSafe *VmSafeSession) ParseJsonString(json string, key string) (string, error) {
	return _VmSafe.Contract.ParseJsonString(&_VmSafe.CallOpts, json, key)
}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_VmSafe *VmSafeCallerSession) ParseJsonString(json string, key string) (string, error) {
	return _VmSafe.Contract.ParseJsonString(&_VmSafe.CallOpts, json, key)
}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_VmSafe *VmSafeCaller) ParseJsonStringArray(opts *bind.CallOpts, json string, key string) ([]string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonStringArray", json, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_VmSafe *VmSafeSession) ParseJsonStringArray(json string, key string) ([]string, error) {
	return _VmSafe.Contract.ParseJsonStringArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_VmSafe *VmSafeCallerSession) ParseJsonStringArray(json string, key string) ([]string, error) {
	return _VmSafe.Contract.ParseJsonStringArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonType is a free data retrieval call binding the contract method 0xa9da313b.
//
// Solidity: function parseJsonType(string json, string typeDescription) pure returns(bytes)
func (_VmSafe *VmSafeCaller) ParseJsonType(opts *bind.CallOpts, json string, typeDescription string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonType", json, typeDescription)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonType is a free data retrieval call binding the contract method 0xa9da313b.
//
// Solidity: function parseJsonType(string json, string typeDescription) pure returns(bytes)
func (_VmSafe *VmSafeSession) ParseJsonType(json string, typeDescription string) ([]byte, error) {
	return _VmSafe.Contract.ParseJsonType(&_VmSafe.CallOpts, json, typeDescription)
}

// ParseJsonType is a free data retrieval call binding the contract method 0xa9da313b.
//
// Solidity: function parseJsonType(string json, string typeDescription) pure returns(bytes)
func (_VmSafe *VmSafeCallerSession) ParseJsonType(json string, typeDescription string) ([]byte, error) {
	return _VmSafe.Contract.ParseJsonType(&_VmSafe.CallOpts, json, typeDescription)
}

// ParseJsonType0 is a free data retrieval call binding the contract method 0xe3f5ae33.
//
// Solidity: function parseJsonType(string json, string key, string typeDescription) pure returns(bytes)
func (_VmSafe *VmSafeCaller) ParseJsonType0(opts *bind.CallOpts, json string, key string, typeDescription string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonType0", json, key, typeDescription)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonType0 is a free data retrieval call binding the contract method 0xe3f5ae33.
//
// Solidity: function parseJsonType(string json, string key, string typeDescription) pure returns(bytes)
func (_VmSafe *VmSafeSession) ParseJsonType0(json string, key string, typeDescription string) ([]byte, error) {
	return _VmSafe.Contract.ParseJsonType0(&_VmSafe.CallOpts, json, key, typeDescription)
}

// ParseJsonType0 is a free data retrieval call binding the contract method 0xe3f5ae33.
//
// Solidity: function parseJsonType(string json, string key, string typeDescription) pure returns(bytes)
func (_VmSafe *VmSafeCallerSession) ParseJsonType0(json string, key string, typeDescription string) ([]byte, error) {
	return _VmSafe.Contract.ParseJsonType0(&_VmSafe.CallOpts, json, key, typeDescription)
}

// ParseJsonTypeArray is a free data retrieval call binding the contract method 0x0175d535.
//
// Solidity: function parseJsonTypeArray(string json, string key, string typeDescription) pure returns(bytes)
func (_VmSafe *VmSafeCaller) ParseJsonTypeArray(opts *bind.CallOpts, json string, key string, typeDescription string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonTypeArray", json, key, typeDescription)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonTypeArray is a free data retrieval call binding the contract method 0x0175d535.
//
// Solidity: function parseJsonTypeArray(string json, string key, string typeDescription) pure returns(bytes)
func (_VmSafe *VmSafeSession) ParseJsonTypeArray(json string, key string, typeDescription string) ([]byte, error) {
	return _VmSafe.Contract.ParseJsonTypeArray(&_VmSafe.CallOpts, json, key, typeDescription)
}

// ParseJsonTypeArray is a free data retrieval call binding the contract method 0x0175d535.
//
// Solidity: function parseJsonTypeArray(string json, string key, string typeDescription) pure returns(bytes)
func (_VmSafe *VmSafeCallerSession) ParseJsonTypeArray(json string, key string, typeDescription string) ([]byte, error) {
	return _VmSafe.Contract.ParseJsonTypeArray(&_VmSafe.CallOpts, json, key, typeDescription)
}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_VmSafe *VmSafeCaller) ParseJsonUint(opts *bind.CallOpts, json string, key string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonUint", json, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_VmSafe *VmSafeSession) ParseJsonUint(json string, key string) (*big.Int, error) {
	return _VmSafe.Contract.ParseJsonUint(&_VmSafe.CallOpts, json, key)
}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_VmSafe *VmSafeCallerSession) ParseJsonUint(json string, key string) (*big.Int, error) {
	return _VmSafe.Contract.ParseJsonUint(&_VmSafe.CallOpts, json, key)
}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_VmSafe *VmSafeCaller) ParseJsonUintArray(opts *bind.CallOpts, json string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseJsonUintArray", json, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_VmSafe *VmSafeSession) ParseJsonUintArray(json string, key string) ([]*big.Int, error) {
	return _VmSafe.Contract.ParseJsonUintArray(&_VmSafe.CallOpts, json, key)
}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_VmSafe *VmSafeCallerSession) ParseJsonUintArray(json string, key string) ([]*big.Int, error) {
	return _VmSafe.Contract.ParseJsonUintArray(&_VmSafe.CallOpts, json, key)
}

// ParseToml is a free data retrieval call binding the contract method 0x37736e08.
//
// Solidity: function parseToml(string toml, string key) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeCaller) ParseToml(opts *bind.CallOpts, toml string, key string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseToml", toml, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseToml is a free data retrieval call binding the contract method 0x37736e08.
//
// Solidity: function parseToml(string toml, string key) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeSession) ParseToml(toml string, key string) ([]byte, error) {
	return _VmSafe.Contract.ParseToml(&_VmSafe.CallOpts, toml, key)
}

// ParseToml is a free data retrieval call binding the contract method 0x37736e08.
//
// Solidity: function parseToml(string toml, string key) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeCallerSession) ParseToml(toml string, key string) ([]byte, error) {
	return _VmSafe.Contract.ParseToml(&_VmSafe.CallOpts, toml, key)
}

// ParseToml0 is a free data retrieval call binding the contract method 0x592151f0.
//
// Solidity: function parseToml(string toml) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeCaller) ParseToml0(opts *bind.CallOpts, toml string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseToml0", toml)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseToml0 is a free data retrieval call binding the contract method 0x592151f0.
//
// Solidity: function parseToml(string toml) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeSession) ParseToml0(toml string) ([]byte, error) {
	return _VmSafe.Contract.ParseToml0(&_VmSafe.CallOpts, toml)
}

// ParseToml0 is a free data retrieval call binding the contract method 0x592151f0.
//
// Solidity: function parseToml(string toml) pure returns(bytes abiEncodedData)
func (_VmSafe *VmSafeCallerSession) ParseToml0(toml string) ([]byte, error) {
	return _VmSafe.Contract.ParseToml0(&_VmSafe.CallOpts, toml)
}

// ParseTomlAddress is a free data retrieval call binding the contract method 0x65e7c844.
//
// Solidity: function parseTomlAddress(string toml, string key) pure returns(address)
func (_VmSafe *VmSafeCaller) ParseTomlAddress(opts *bind.CallOpts, toml string, key string) (common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlAddress", toml, key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseTomlAddress is a free data retrieval call binding the contract method 0x65e7c844.
//
// Solidity: function parseTomlAddress(string toml, string key) pure returns(address)
func (_VmSafe *VmSafeSession) ParseTomlAddress(toml string, key string) (common.Address, error) {
	return _VmSafe.Contract.ParseTomlAddress(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlAddress is a free data retrieval call binding the contract method 0x65e7c844.
//
// Solidity: function parseTomlAddress(string toml, string key) pure returns(address)
func (_VmSafe *VmSafeCallerSession) ParseTomlAddress(toml string, key string) (common.Address, error) {
	return _VmSafe.Contract.ParseTomlAddress(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlAddressArray is a free data retrieval call binding the contract method 0x65c428e7.
//
// Solidity: function parseTomlAddressArray(string toml, string key) pure returns(address[])
func (_VmSafe *VmSafeCaller) ParseTomlAddressArray(opts *bind.CallOpts, toml string, key string) ([]common.Address, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlAddressArray", toml, key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ParseTomlAddressArray is a free data retrieval call binding the contract method 0x65c428e7.
//
// Solidity: function parseTomlAddressArray(string toml, string key) pure returns(address[])
func (_VmSafe *VmSafeSession) ParseTomlAddressArray(toml string, key string) ([]common.Address, error) {
	return _VmSafe.Contract.ParseTomlAddressArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlAddressArray is a free data retrieval call binding the contract method 0x65c428e7.
//
// Solidity: function parseTomlAddressArray(string toml, string key) pure returns(address[])
func (_VmSafe *VmSafeCallerSession) ParseTomlAddressArray(toml string, key string) ([]common.Address, error) {
	return _VmSafe.Contract.ParseTomlAddressArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBool is a free data retrieval call binding the contract method 0xd30dced6.
//
// Solidity: function parseTomlBool(string toml, string key) pure returns(bool)
func (_VmSafe *VmSafeCaller) ParseTomlBool(opts *bind.CallOpts, toml string, key string) (bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlBool", toml, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseTomlBool is a free data retrieval call binding the contract method 0xd30dced6.
//
// Solidity: function parseTomlBool(string toml, string key) pure returns(bool)
func (_VmSafe *VmSafeSession) ParseTomlBool(toml string, key string) (bool, error) {
	return _VmSafe.Contract.ParseTomlBool(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBool is a free data retrieval call binding the contract method 0xd30dced6.
//
// Solidity: function parseTomlBool(string toml, string key) pure returns(bool)
func (_VmSafe *VmSafeCallerSession) ParseTomlBool(toml string, key string) (bool, error) {
	return _VmSafe.Contract.ParseTomlBool(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBoolArray is a free data retrieval call binding the contract method 0x127cfe9a.
//
// Solidity: function parseTomlBoolArray(string toml, string key) pure returns(bool[])
func (_VmSafe *VmSafeCaller) ParseTomlBoolArray(opts *bind.CallOpts, toml string, key string) ([]bool, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlBoolArray", toml, key)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// ParseTomlBoolArray is a free data retrieval call binding the contract method 0x127cfe9a.
//
// Solidity: function parseTomlBoolArray(string toml, string key) pure returns(bool[])
func (_VmSafe *VmSafeSession) ParseTomlBoolArray(toml string, key string) ([]bool, error) {
	return _VmSafe.Contract.ParseTomlBoolArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBoolArray is a free data retrieval call binding the contract method 0x127cfe9a.
//
// Solidity: function parseTomlBoolArray(string toml, string key) pure returns(bool[])
func (_VmSafe *VmSafeCallerSession) ParseTomlBoolArray(toml string, key string) ([]bool, error) {
	return _VmSafe.Contract.ParseTomlBoolArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBytes is a free data retrieval call binding the contract method 0xd77bfdb9.
//
// Solidity: function parseTomlBytes(string toml, string key) pure returns(bytes)
func (_VmSafe *VmSafeCaller) ParseTomlBytes(opts *bind.CallOpts, toml string, key string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlBytes", toml, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseTomlBytes is a free data retrieval call binding the contract method 0xd77bfdb9.
//
// Solidity: function parseTomlBytes(string toml, string key) pure returns(bytes)
func (_VmSafe *VmSafeSession) ParseTomlBytes(toml string, key string) ([]byte, error) {
	return _VmSafe.Contract.ParseTomlBytes(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBytes is a free data retrieval call binding the contract method 0xd77bfdb9.
//
// Solidity: function parseTomlBytes(string toml, string key) pure returns(bytes)
func (_VmSafe *VmSafeCallerSession) ParseTomlBytes(toml string, key string) ([]byte, error) {
	return _VmSafe.Contract.ParseTomlBytes(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBytes32 is a free data retrieval call binding the contract method 0x8e214810.
//
// Solidity: function parseTomlBytes32(string toml, string key) pure returns(bytes32)
func (_VmSafe *VmSafeCaller) ParseTomlBytes32(opts *bind.CallOpts, toml string, key string) ([32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlBytes32", toml, key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseTomlBytes32 is a free data retrieval call binding the contract method 0x8e214810.
//
// Solidity: function parseTomlBytes32(string toml, string key) pure returns(bytes32)
func (_VmSafe *VmSafeSession) ParseTomlBytes32(toml string, key string) ([32]byte, error) {
	return _VmSafe.Contract.ParseTomlBytes32(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBytes32 is a free data retrieval call binding the contract method 0x8e214810.
//
// Solidity: function parseTomlBytes32(string toml, string key) pure returns(bytes32)
func (_VmSafe *VmSafeCallerSession) ParseTomlBytes32(toml string, key string) ([32]byte, error) {
	return _VmSafe.Contract.ParseTomlBytes32(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBytes32Array is a free data retrieval call binding the contract method 0x3e716f81.
//
// Solidity: function parseTomlBytes32Array(string toml, string key) pure returns(bytes32[])
func (_VmSafe *VmSafeCaller) ParseTomlBytes32Array(opts *bind.CallOpts, toml string, key string) ([][32]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlBytes32Array", toml, key)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ParseTomlBytes32Array is a free data retrieval call binding the contract method 0x3e716f81.
//
// Solidity: function parseTomlBytes32Array(string toml, string key) pure returns(bytes32[])
func (_VmSafe *VmSafeSession) ParseTomlBytes32Array(toml string, key string) ([][32]byte, error) {
	return _VmSafe.Contract.ParseTomlBytes32Array(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBytes32Array is a free data retrieval call binding the contract method 0x3e716f81.
//
// Solidity: function parseTomlBytes32Array(string toml, string key) pure returns(bytes32[])
func (_VmSafe *VmSafeCallerSession) ParseTomlBytes32Array(toml string, key string) ([][32]byte, error) {
	return _VmSafe.Contract.ParseTomlBytes32Array(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBytesArray is a free data retrieval call binding the contract method 0xb197c247.
//
// Solidity: function parseTomlBytesArray(string toml, string key) pure returns(bytes[])
func (_VmSafe *VmSafeCaller) ParseTomlBytesArray(opts *bind.CallOpts, toml string, key string) ([][]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlBytesArray", toml, key)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ParseTomlBytesArray is a free data retrieval call binding the contract method 0xb197c247.
//
// Solidity: function parseTomlBytesArray(string toml, string key) pure returns(bytes[])
func (_VmSafe *VmSafeSession) ParseTomlBytesArray(toml string, key string) ([][]byte, error) {
	return _VmSafe.Contract.ParseTomlBytesArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlBytesArray is a free data retrieval call binding the contract method 0xb197c247.
//
// Solidity: function parseTomlBytesArray(string toml, string key) pure returns(bytes[])
func (_VmSafe *VmSafeCallerSession) ParseTomlBytesArray(toml string, key string) ([][]byte, error) {
	return _VmSafe.Contract.ParseTomlBytesArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlInt is a free data retrieval call binding the contract method 0xc1350739.
//
// Solidity: function parseTomlInt(string toml, string key) pure returns(int256)
func (_VmSafe *VmSafeCaller) ParseTomlInt(opts *bind.CallOpts, toml string, key string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlInt", toml, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseTomlInt is a free data retrieval call binding the contract method 0xc1350739.
//
// Solidity: function parseTomlInt(string toml, string key) pure returns(int256)
func (_VmSafe *VmSafeSession) ParseTomlInt(toml string, key string) (*big.Int, error) {
	return _VmSafe.Contract.ParseTomlInt(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlInt is a free data retrieval call binding the contract method 0xc1350739.
//
// Solidity: function parseTomlInt(string toml, string key) pure returns(int256)
func (_VmSafe *VmSafeCallerSession) ParseTomlInt(toml string, key string) (*big.Int, error) {
	return _VmSafe.Contract.ParseTomlInt(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlIntArray is a free data retrieval call binding the contract method 0xd3522ae6.
//
// Solidity: function parseTomlIntArray(string toml, string key) pure returns(int256[])
func (_VmSafe *VmSafeCaller) ParseTomlIntArray(opts *bind.CallOpts, toml string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlIntArray", toml, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseTomlIntArray is a free data retrieval call binding the contract method 0xd3522ae6.
//
// Solidity: function parseTomlIntArray(string toml, string key) pure returns(int256[])
func (_VmSafe *VmSafeSession) ParseTomlIntArray(toml string, key string) ([]*big.Int, error) {
	return _VmSafe.Contract.ParseTomlIntArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlIntArray is a free data retrieval call binding the contract method 0xd3522ae6.
//
// Solidity: function parseTomlIntArray(string toml, string key) pure returns(int256[])
func (_VmSafe *VmSafeCallerSession) ParseTomlIntArray(toml string, key string) ([]*big.Int, error) {
	return _VmSafe.Contract.ParseTomlIntArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlKeys is a free data retrieval call binding the contract method 0x812a44b2.
//
// Solidity: function parseTomlKeys(string toml, string key) pure returns(string[] keys)
func (_VmSafe *VmSafeCaller) ParseTomlKeys(opts *bind.CallOpts, toml string, key string) ([]string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlKeys", toml, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseTomlKeys is a free data retrieval call binding the contract method 0x812a44b2.
//
// Solidity: function parseTomlKeys(string toml, string key) pure returns(string[] keys)
func (_VmSafe *VmSafeSession) ParseTomlKeys(toml string, key string) ([]string, error) {
	return _VmSafe.Contract.ParseTomlKeys(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlKeys is a free data retrieval call binding the contract method 0x812a44b2.
//
// Solidity: function parseTomlKeys(string toml, string key) pure returns(string[] keys)
func (_VmSafe *VmSafeCallerSession) ParseTomlKeys(toml string, key string) ([]string, error) {
	return _VmSafe.Contract.ParseTomlKeys(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlString is a free data retrieval call binding the contract method 0x8bb8dd43.
//
// Solidity: function parseTomlString(string toml, string key) pure returns(string)
func (_VmSafe *VmSafeCaller) ParseTomlString(opts *bind.CallOpts, toml string, key string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlString", toml, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ParseTomlString is a free data retrieval call binding the contract method 0x8bb8dd43.
//
// Solidity: function parseTomlString(string toml, string key) pure returns(string)
func (_VmSafe *VmSafeSession) ParseTomlString(toml string, key string) (string, error) {
	return _VmSafe.Contract.ParseTomlString(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlString is a free data retrieval call binding the contract method 0x8bb8dd43.
//
// Solidity: function parseTomlString(string toml, string key) pure returns(string)
func (_VmSafe *VmSafeCallerSession) ParseTomlString(toml string, key string) (string, error) {
	return _VmSafe.Contract.ParseTomlString(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlStringArray is a free data retrieval call binding the contract method 0x9f629281.
//
// Solidity: function parseTomlStringArray(string toml, string key) pure returns(string[])
func (_VmSafe *VmSafeCaller) ParseTomlStringArray(opts *bind.CallOpts, toml string, key string) ([]string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlStringArray", toml, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseTomlStringArray is a free data retrieval call binding the contract method 0x9f629281.
//
// Solidity: function parseTomlStringArray(string toml, string key) pure returns(string[])
func (_VmSafe *VmSafeSession) ParseTomlStringArray(toml string, key string) ([]string, error) {
	return _VmSafe.Contract.ParseTomlStringArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlStringArray is a free data retrieval call binding the contract method 0x9f629281.
//
// Solidity: function parseTomlStringArray(string toml, string key) pure returns(string[])
func (_VmSafe *VmSafeCallerSession) ParseTomlStringArray(toml string, key string) ([]string, error) {
	return _VmSafe.Contract.ParseTomlStringArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlUint is a free data retrieval call binding the contract method 0xcc7b0487.
//
// Solidity: function parseTomlUint(string toml, string key) pure returns(uint256)
func (_VmSafe *VmSafeCaller) ParseTomlUint(opts *bind.CallOpts, toml string, key string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlUint", toml, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseTomlUint is a free data retrieval call binding the contract method 0xcc7b0487.
//
// Solidity: function parseTomlUint(string toml, string key) pure returns(uint256)
func (_VmSafe *VmSafeSession) ParseTomlUint(toml string, key string) (*big.Int, error) {
	return _VmSafe.Contract.ParseTomlUint(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlUint is a free data retrieval call binding the contract method 0xcc7b0487.
//
// Solidity: function parseTomlUint(string toml, string key) pure returns(uint256)
func (_VmSafe *VmSafeCallerSession) ParseTomlUint(toml string, key string) (*big.Int, error) {
	return _VmSafe.Contract.ParseTomlUint(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlUintArray is a free data retrieval call binding the contract method 0xb5df27c8.
//
// Solidity: function parseTomlUintArray(string toml, string key) pure returns(uint256[])
func (_VmSafe *VmSafeCaller) ParseTomlUintArray(opts *bind.CallOpts, toml string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseTomlUintArray", toml, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseTomlUintArray is a free data retrieval call binding the contract method 0xb5df27c8.
//
// Solidity: function parseTomlUintArray(string toml, string key) pure returns(uint256[])
func (_VmSafe *VmSafeSession) ParseTomlUintArray(toml string, key string) ([]*big.Int, error) {
	return _VmSafe.Contract.ParseTomlUintArray(&_VmSafe.CallOpts, toml, key)
}

// ParseTomlUintArray is a free data retrieval call binding the contract method 0xb5df27c8.
//
// Solidity: function parseTomlUintArray(string toml, string key) pure returns(uint256[])
func (_VmSafe *VmSafeCallerSession) ParseTomlUintArray(toml string, key string) ([]*big.Int, error) {
	return _VmSafe.Contract.ParseTomlUintArray(&_VmSafe.CallOpts, toml, key)
}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_VmSafe *VmSafeCaller) ParseUint(opts *bind.CallOpts, stringifiedValue string) (*big.Int, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "parseUint", stringifiedValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_VmSafe *VmSafeSession) ParseUint(stringifiedValue string) (*big.Int, error) {
	return _VmSafe.Contract.ParseUint(&_VmSafe.CallOpts, stringifiedValue)
}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_VmSafe *VmSafeCallerSession) ParseUint(stringifiedValue string) (*big.Int, error) {
	return _VmSafe.Contract.ParseUint(&_VmSafe.CallOpts, stringifiedValue)
}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_VmSafe *VmSafeCaller) ProjectRoot(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "projectRoot")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_VmSafe *VmSafeSession) ProjectRoot() (string, error) {
	return _VmSafe.Contract.ProjectRoot(&_VmSafe.CallOpts)
}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_VmSafe *VmSafeCallerSession) ProjectRoot() (string, error) {
	return _VmSafe.Contract.ProjectRoot(&_VmSafe.CallOpts)
}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_VmSafe *VmSafeCaller) ReadDir(opts *bind.CallOpts, path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "readDir", path, maxDepth)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_VmSafe *VmSafeSession) ReadDir(path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	return _VmSafe.Contract.ReadDir(&_VmSafe.CallOpts, path, maxDepth)
}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_VmSafe *VmSafeCallerSession) ReadDir(path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	return _VmSafe.Contract.ReadDir(&_VmSafe.CallOpts, path, maxDepth)
}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_VmSafe *VmSafeCaller) ReadDir0(opts *bind.CallOpts, path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "readDir0", path, maxDepth, followLinks)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_VmSafe *VmSafeSession) ReadDir0(path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	return _VmSafe.Contract.ReadDir0(&_VmSafe.CallOpts, path, maxDepth, followLinks)
}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_VmSafe *VmSafeCallerSession) ReadDir0(path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	return _VmSafe.Contract.ReadDir0(&_VmSafe.CallOpts, path, maxDepth, followLinks)
}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_VmSafe *VmSafeCaller) ReadDir1(opts *bind.CallOpts, path string) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "readDir1", path)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_VmSafe *VmSafeSession) ReadDir1(path string) ([]VmSafeDirEntry, error) {
	return _VmSafe.Contract.ReadDir1(&_VmSafe.CallOpts, path)
}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_VmSafe *VmSafeCallerSession) ReadDir1(path string) ([]VmSafeDirEntry, error) {
	return _VmSafe.Contract.ReadDir1(&_VmSafe.CallOpts, path)
}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_VmSafe *VmSafeCaller) ReadFile(opts *bind.CallOpts, path string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "readFile", path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_VmSafe *VmSafeSession) ReadFile(path string) (string, error) {
	return _VmSafe.Contract.ReadFile(&_VmSafe.CallOpts, path)
}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_VmSafe *VmSafeCallerSession) ReadFile(path string) (string, error) {
	return _VmSafe.Contract.ReadFile(&_VmSafe.CallOpts, path)
}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_VmSafe *VmSafeCaller) ReadFileBinary(opts *bind.CallOpts, path string) ([]byte, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "readFileBinary", path)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_VmSafe *VmSafeSession) ReadFileBinary(path string) ([]byte, error) {
	return _VmSafe.Contract.ReadFileBinary(&_VmSafe.CallOpts, path)
}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_VmSafe *VmSafeCallerSession) ReadFileBinary(path string) ([]byte, error) {
	return _VmSafe.Contract.ReadFileBinary(&_VmSafe.CallOpts, path)
}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_VmSafe *VmSafeCaller) ReadLine(opts *bind.CallOpts, path string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "readLine", path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_VmSafe *VmSafeSession) ReadLine(path string) (string, error) {
	return _VmSafe.Contract.ReadLine(&_VmSafe.CallOpts, path)
}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_VmSafe *VmSafeCallerSession) ReadLine(path string) (string, error) {
	return _VmSafe.Contract.ReadLine(&_VmSafe.CallOpts, path)
}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_VmSafe *VmSafeCaller) ReadLink(opts *bind.CallOpts, linkPath string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "readLink", linkPath)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_VmSafe *VmSafeSession) ReadLink(linkPath string) (string, error) {
	return _VmSafe.Contract.ReadLink(&_VmSafe.CallOpts, linkPath)
}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_VmSafe *VmSafeCallerSession) ReadLink(linkPath string) (string, error) {
	return _VmSafe.Contract.ReadLink(&_VmSafe.CallOpts, linkPath)
}

// Replace is a free data retrieval call binding the contract method 0xe00ad03e.
//
// Solidity: function replace(string input, string from, string to) pure returns(string output)
func (_VmSafe *VmSafeCaller) Replace(opts *bind.CallOpts, input string, from string, to string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "replace", input, from, to)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Replace is a free data retrieval call binding the contract method 0xe00ad03e.
//
// Solidity: function replace(string input, string from, string to) pure returns(string output)
func (_VmSafe *VmSafeSession) Replace(input string, from string, to string) (string, error) {
	return _VmSafe.Contract.Replace(&_VmSafe.CallOpts, input, from, to)
}

// Replace is a free data retrieval call binding the contract method 0xe00ad03e.
//
// Solidity: function replace(string input, string from, string to) pure returns(string output)
func (_VmSafe *VmSafeCallerSession) Replace(input string, from string, to string) (string, error) {
	return _VmSafe.Contract.Replace(&_VmSafe.CallOpts, input, from, to)
}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_VmSafe *VmSafeCaller) RpcUrl(opts *bind.CallOpts, rpcAlias string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "rpcUrl", rpcAlias)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_VmSafe *VmSafeSession) RpcUrl(rpcAlias string) (string, error) {
	return _VmSafe.Contract.RpcUrl(&_VmSafe.CallOpts, rpcAlias)
}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_VmSafe *VmSafeCallerSession) RpcUrl(rpcAlias string) (string, error) {
	return _VmSafe.Contract.RpcUrl(&_VmSafe.CallOpts, rpcAlias)
}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_VmSafe *VmSafeCaller) RpcUrlStructs(opts *bind.CallOpts) ([]VmSafeRpc, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "rpcUrlStructs")

	if err != nil {
		return *new([]VmSafeRpc), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeRpc)).(*[]VmSafeRpc)

	return out0, err

}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_VmSafe *VmSafeSession) RpcUrlStructs() ([]VmSafeRpc, error) {
	return _VmSafe.Contract.RpcUrlStructs(&_VmSafe.CallOpts)
}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_VmSafe *VmSafeCallerSession) RpcUrlStructs() ([]VmSafeRpc, error) {
	return _VmSafe.Contract.RpcUrlStructs(&_VmSafe.CallOpts)
}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_VmSafe *VmSafeCaller) RpcUrls(opts *bind.CallOpts) ([][2]string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "rpcUrls")

	if err != nil {
		return *new([][2]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][2]string)).(*[][2]string)

	return out0, err

}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_VmSafe *VmSafeSession) RpcUrls() ([][2]string, error) {
	return _VmSafe.Contract.RpcUrls(&_VmSafe.CallOpts)
}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_VmSafe *VmSafeCallerSession) RpcUrls() ([][2]string, error) {
	return _VmSafe.Contract.RpcUrls(&_VmSafe.CallOpts)
}

// SerializeJsonType is a free data retrieval call binding the contract method 0x6d4f96a6.
//
// Solidity: function serializeJsonType(string typeDescription, bytes value) pure returns(string json)
func (_VmSafe *VmSafeCaller) SerializeJsonType(opts *bind.CallOpts, typeDescription string, value []byte) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "serializeJsonType", typeDescription, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SerializeJsonType is a free data retrieval call binding the contract method 0x6d4f96a6.
//
// Solidity: function serializeJsonType(string typeDescription, bytes value) pure returns(string json)
func (_VmSafe *VmSafeSession) SerializeJsonType(typeDescription string, value []byte) (string, error) {
	return _VmSafe.Contract.SerializeJsonType(&_VmSafe.CallOpts, typeDescription, value)
}

// SerializeJsonType is a free data retrieval call binding the contract method 0x6d4f96a6.
//
// Solidity: function serializeJsonType(string typeDescription, bytes value) pure returns(string json)
func (_VmSafe *VmSafeCallerSession) SerializeJsonType(typeDescription string, value []byte) (string, error) {
	return _VmSafe.Contract.SerializeJsonType(&_VmSafe.CallOpts, typeDescription, value)
}

// Sign is a free data retrieval call binding the contract method 0x799cd333.
//
// Solidity: function sign(bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeCaller) Sign(opts *bind.CallOpts, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "sign", digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign is a free data retrieval call binding the contract method 0x799cd333.
//
// Solidity: function sign(bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeSession) Sign(digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _VmSafe.Contract.Sign(&_VmSafe.CallOpts, digest)
}

// Sign is a free data retrieval call binding the contract method 0x799cd333.
//
// Solidity: function sign(bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeCallerSession) Sign(digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _VmSafe.Contract.Sign(&_VmSafe.CallOpts, digest)
}

// Sign0 is a free data retrieval call binding the contract method 0x8c1aa205.
//
// Solidity: function sign(address signer, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeCaller) Sign0(opts *bind.CallOpts, signer common.Address, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "sign0", signer, digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign0 is a free data retrieval call binding the contract method 0x8c1aa205.
//
// Solidity: function sign(address signer, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeSession) Sign0(signer common.Address, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _VmSafe.Contract.Sign0(&_VmSafe.CallOpts, signer, digest)
}

// Sign0 is a free data retrieval call binding the contract method 0x8c1aa205.
//
// Solidity: function sign(address signer, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeCallerSession) Sign0(signer common.Address, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _VmSafe.Contract.Sign0(&_VmSafe.CallOpts, signer, digest)
}

// Sign2 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeCaller) Sign2(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "sign2", privateKey, digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign2 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeSession) Sign2(privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _VmSafe.Contract.Sign2(&_VmSafe.CallOpts, privateKey, digest)
}

// Sign2 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeCallerSession) Sign2(privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _VmSafe.Contract.Sign2(&_VmSafe.CallOpts, privateKey, digest)
}

// SignCompact0 is a free data retrieval call binding the contract method 0x8e2f97bf.
//
// Solidity: function signCompact(address signer, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeCaller) SignCompact0(opts *bind.CallOpts, signer common.Address, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "signCompact0", signer, digest)

	outstruct := new(struct {
		R  [32]byte
		Vs [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Vs = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignCompact0 is a free data retrieval call binding the contract method 0x8e2f97bf.
//
// Solidity: function signCompact(address signer, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeSession) SignCompact0(signer common.Address, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _VmSafe.Contract.SignCompact0(&_VmSafe.CallOpts, signer, digest)
}

// SignCompact0 is a free data retrieval call binding the contract method 0x8e2f97bf.
//
// Solidity: function signCompact(address signer, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeCallerSession) SignCompact0(signer common.Address, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _VmSafe.Contract.SignCompact0(&_VmSafe.CallOpts, signer, digest)
}

// SignCompact1 is a free data retrieval call binding the contract method 0xa282dc4b.
//
// Solidity: function signCompact(bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeCaller) SignCompact1(opts *bind.CallOpts, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "signCompact1", digest)

	outstruct := new(struct {
		R  [32]byte
		Vs [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Vs = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignCompact1 is a free data retrieval call binding the contract method 0xa282dc4b.
//
// Solidity: function signCompact(bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeSession) SignCompact1(digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _VmSafe.Contract.SignCompact1(&_VmSafe.CallOpts, digest)
}

// SignCompact1 is a free data retrieval call binding the contract method 0xa282dc4b.
//
// Solidity: function signCompact(bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeCallerSession) SignCompact1(digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _VmSafe.Contract.SignCompact1(&_VmSafe.CallOpts, digest)
}

// SignCompact2 is a free data retrieval call binding the contract method 0xcc2a781f.
//
// Solidity: function signCompact(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeCaller) SignCompact2(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "signCompact2", privateKey, digest)

	outstruct := new(struct {
		R  [32]byte
		Vs [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Vs = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignCompact2 is a free data retrieval call binding the contract method 0xcc2a781f.
//
// Solidity: function signCompact(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeSession) SignCompact2(privateKey *big.Int, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _VmSafe.Contract.SignCompact2(&_VmSafe.CallOpts, privateKey, digest)
}

// SignCompact2 is a free data retrieval call binding the contract method 0xcc2a781f.
//
// Solidity: function signCompact(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeCallerSession) SignCompact2(privateKey *big.Int, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _VmSafe.Contract.SignCompact2(&_VmSafe.CallOpts, privateKey, digest)
}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_VmSafe *VmSafeCaller) SignP256(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "signP256", privateKey, digest)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_VmSafe *VmSafeSession) SignP256(privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	return _VmSafe.Contract.SignP256(&_VmSafe.CallOpts, privateKey, digest)
}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_VmSafe *VmSafeCallerSession) SignP256(privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	return _VmSafe.Contract.SignP256(&_VmSafe.CallOpts, privateKey, digest)
}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string input, string delimiter) pure returns(string[] outputs)
func (_VmSafe *VmSafeCaller) Split(opts *bind.CallOpts, input string, delimiter string) ([]string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "split", input, delimiter)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string input, string delimiter) pure returns(string[] outputs)
func (_VmSafe *VmSafeSession) Split(input string, delimiter string) ([]string, error) {
	return _VmSafe.Contract.Split(&_VmSafe.CallOpts, input, delimiter)
}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string input, string delimiter) pure returns(string[] outputs)
func (_VmSafe *VmSafeCallerSession) Split(input string, delimiter string) ([]string, error) {
	return _VmSafe.Contract.Split(&_VmSafe.CallOpts, input, delimiter)
}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_VmSafe *VmSafeCaller) ToBase64(opts *bind.CallOpts, data string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toBase64", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_VmSafe *VmSafeSession) ToBase64(data string) (string, error) {
	return _VmSafe.Contract.ToBase64(&_VmSafe.CallOpts, data)
}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_VmSafe *VmSafeCallerSession) ToBase64(data string) (string, error) {
	return _VmSafe.Contract.ToBase64(&_VmSafe.CallOpts, data)
}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_VmSafe *VmSafeCaller) ToBase640(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toBase640", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_VmSafe *VmSafeSession) ToBase640(data []byte) (string, error) {
	return _VmSafe.Contract.ToBase640(&_VmSafe.CallOpts, data)
}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_VmSafe *VmSafeCallerSession) ToBase640(data []byte) (string, error) {
	return _VmSafe.Contract.ToBase640(&_VmSafe.CallOpts, data)
}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_VmSafe *VmSafeCaller) ToBase64URL(opts *bind.CallOpts, data string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toBase64URL", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_VmSafe *VmSafeSession) ToBase64URL(data string) (string, error) {
	return _VmSafe.Contract.ToBase64URL(&_VmSafe.CallOpts, data)
}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_VmSafe *VmSafeCallerSession) ToBase64URL(data string) (string, error) {
	return _VmSafe.Contract.ToBase64URL(&_VmSafe.CallOpts, data)
}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_VmSafe *VmSafeCaller) ToBase64URL0(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toBase64URL0", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_VmSafe *VmSafeSession) ToBase64URL0(data []byte) (string, error) {
	return _VmSafe.Contract.ToBase64URL0(&_VmSafe.CallOpts, data)
}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_VmSafe *VmSafeCallerSession) ToBase64URL0(data []byte) (string, error) {
	return _VmSafe.Contract.ToBase64URL0(&_VmSafe.CallOpts, data)
}

// ToLowercase is a free data retrieval call binding the contract method 0x50bb0884.
//
// Solidity: function toLowercase(string input) pure returns(string output)
func (_VmSafe *VmSafeCaller) ToLowercase(opts *bind.CallOpts, input string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toLowercase", input)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToLowercase is a free data retrieval call binding the contract method 0x50bb0884.
//
// Solidity: function toLowercase(string input) pure returns(string output)
func (_VmSafe *VmSafeSession) ToLowercase(input string) (string, error) {
	return _VmSafe.Contract.ToLowercase(&_VmSafe.CallOpts, input)
}

// ToLowercase is a free data retrieval call binding the contract method 0x50bb0884.
//
// Solidity: function toLowercase(string input) pure returns(string output)
func (_VmSafe *VmSafeCallerSession) ToLowercase(input string) (string, error) {
	return _VmSafe.Contract.ToLowercase(&_VmSafe.CallOpts, input)
}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCaller) ToString(opts *bind.CallOpts, value common.Address) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeSession) ToString(value common.Address) (string, error) {
	return _VmSafe.Contract.ToString(&_VmSafe.CallOpts, value)
}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCallerSession) ToString(value common.Address) (string, error) {
	return _VmSafe.Contract.ToString(&_VmSafe.CallOpts, value)
}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCaller) ToString0(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toString0", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeSession) ToString0(value *big.Int) (string, error) {
	return _VmSafe.Contract.ToString0(&_VmSafe.CallOpts, value)
}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCallerSession) ToString0(value *big.Int) (string, error) {
	return _VmSafe.Contract.ToString0(&_VmSafe.CallOpts, value)
}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCaller) ToString1(opts *bind.CallOpts, value []byte) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toString1", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeSession) ToString1(value []byte) (string, error) {
	return _VmSafe.Contract.ToString1(&_VmSafe.CallOpts, value)
}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCallerSession) ToString1(value []byte) (string, error) {
	return _VmSafe.Contract.ToString1(&_VmSafe.CallOpts, value)
}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCaller) ToString2(opts *bind.CallOpts, value bool) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toString2", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeSession) ToString2(value bool) (string, error) {
	return _VmSafe.Contract.ToString2(&_VmSafe.CallOpts, value)
}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCallerSession) ToString2(value bool) (string, error) {
	return _VmSafe.Contract.ToString2(&_VmSafe.CallOpts, value)
}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCaller) ToString3(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toString3", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeSession) ToString3(value *big.Int) (string, error) {
	return _VmSafe.Contract.ToString3(&_VmSafe.CallOpts, value)
}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCallerSession) ToString3(value *big.Int) (string, error) {
	return _VmSafe.Contract.ToString3(&_VmSafe.CallOpts, value)
}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCaller) ToString4(opts *bind.CallOpts, value [32]byte) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toString4", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeSession) ToString4(value [32]byte) (string, error) {
	return _VmSafe.Contract.ToString4(&_VmSafe.CallOpts, value)
}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_VmSafe *VmSafeCallerSession) ToString4(value [32]byte) (string, error) {
	return _VmSafe.Contract.ToString4(&_VmSafe.CallOpts, value)
}

// ToUppercase is a free data retrieval call binding the contract method 0x074ae3d7.
//
// Solidity: function toUppercase(string input) pure returns(string output)
func (_VmSafe *VmSafeCaller) ToUppercase(opts *bind.CallOpts, input string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "toUppercase", input)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToUppercase is a free data retrieval call binding the contract method 0x074ae3d7.
//
// Solidity: function toUppercase(string input) pure returns(string output)
func (_VmSafe *VmSafeSession) ToUppercase(input string) (string, error) {
	return _VmSafe.Contract.ToUppercase(&_VmSafe.CallOpts, input)
}

// ToUppercase is a free data retrieval call binding the contract method 0x074ae3d7.
//
// Solidity: function toUppercase(string input) pure returns(string output)
func (_VmSafe *VmSafeCallerSession) ToUppercase(input string) (string, error) {
	return _VmSafe.Contract.ToUppercase(&_VmSafe.CallOpts, input)
}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string input) pure returns(string output)
func (_VmSafe *VmSafeCaller) Trim(opts *bind.CallOpts, input string) (string, error) {
	var out []interface{}
	err := _VmSafe.contract.Call(opts, &out, "trim", input)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string input) pure returns(string output)
func (_VmSafe *VmSafeSession) Trim(input string) (string, error) {
	return _VmSafe.Contract.Trim(&_VmSafe.CallOpts, input)
}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string input) pure returns(string output)
func (_VmSafe *VmSafeCallerSession) Trim(input string) (string, error) {
	return _VmSafe.Contract.Trim(&_VmSafe.CallOpts, input)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_VmSafe *VmSafeTransactor) Accesses(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "accesses", target)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_VmSafe *VmSafeSession) Accesses(target common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.Accesses(&_VmSafe.TransactOpts, target)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_VmSafe *VmSafeTransactorSession) Accesses(target common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.Accesses(&_VmSafe.TransactOpts, target)
}

// Breakpoint is a paid mutator transaction binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) returns()
func (_VmSafe *VmSafeTransactor) Breakpoint(opts *bind.TransactOpts, char string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "breakpoint", char)
}

// Breakpoint is a paid mutator transaction binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) returns()
func (_VmSafe *VmSafeSession) Breakpoint(char string) (*types.Transaction, error) {
	return _VmSafe.Contract.Breakpoint(&_VmSafe.TransactOpts, char)
}

// Breakpoint is a paid mutator transaction binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) returns()
func (_VmSafe *VmSafeTransactorSession) Breakpoint(char string) (*types.Transaction, error) {
	return _VmSafe.Contract.Breakpoint(&_VmSafe.TransactOpts, char)
}

// Breakpoint0 is a paid mutator transaction binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) returns()
func (_VmSafe *VmSafeTransactor) Breakpoint0(opts *bind.TransactOpts, char string, value bool) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "breakpoint0", char, value)
}

// Breakpoint0 is a paid mutator transaction binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) returns()
func (_VmSafe *VmSafeSession) Breakpoint0(char string, value bool) (*types.Transaction, error) {
	return _VmSafe.Contract.Breakpoint0(&_VmSafe.TransactOpts, char, value)
}

// Breakpoint0 is a paid mutator transaction binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) returns()
func (_VmSafe *VmSafeTransactorSession) Breakpoint0(char string, value bool) (*types.Transaction, error) {
	return _VmSafe.Contract.Breakpoint0(&_VmSafe.TransactOpts, char, value)
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_VmSafe *VmSafeTransactor) Broadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "broadcast")
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_VmSafe *VmSafeSession) Broadcast() (*types.Transaction, error) {
	return _VmSafe.Contract.Broadcast(&_VmSafe.TransactOpts)
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_VmSafe *VmSafeTransactorSession) Broadcast() (*types.Transaction, error) {
	return _VmSafe.Contract.Broadcast(&_VmSafe.TransactOpts)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_VmSafe *VmSafeTransactor) Broadcast0(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "broadcast0", signer)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_VmSafe *VmSafeSession) Broadcast0(signer common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.Broadcast0(&_VmSafe.TransactOpts, signer)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_VmSafe *VmSafeTransactorSession) Broadcast0(signer common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.Broadcast0(&_VmSafe.TransactOpts, signer)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_VmSafe *VmSafeTransactor) Broadcast1(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "broadcast1", privateKey)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_VmSafe *VmSafeSession) Broadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.Broadcast1(&_VmSafe.TransactOpts, privateKey)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_VmSafe *VmSafeTransactorSession) Broadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.Broadcast1(&_VmSafe.TransactOpts, privateKey)
}

// BroadcastRawTransaction is a paid mutator transaction binding the contract method 0x8c0c72e0.
//
// Solidity: function broadcastRawTransaction(bytes data) returns()
func (_VmSafe *VmSafeTransactor) BroadcastRawTransaction(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "broadcastRawTransaction", data)
}

// BroadcastRawTransaction is a paid mutator transaction binding the contract method 0x8c0c72e0.
//
// Solidity: function broadcastRawTransaction(bytes data) returns()
func (_VmSafe *VmSafeSession) BroadcastRawTransaction(data []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.BroadcastRawTransaction(&_VmSafe.TransactOpts, data)
}

// BroadcastRawTransaction is a paid mutator transaction binding the contract method 0x8c0c72e0.
//
// Solidity: function broadcastRawTransaction(bytes data) returns()
func (_VmSafe *VmSafeTransactorSession) BroadcastRawTransaction(data []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.BroadcastRawTransaction(&_VmSafe.TransactOpts, data)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_VmSafe *VmSafeTransactor) CloseFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "closeFile", path)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_VmSafe *VmSafeSession) CloseFile(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.CloseFile(&_VmSafe.TransactOpts, path)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_VmSafe *VmSafeTransactorSession) CloseFile(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.CloseFile(&_VmSafe.TransactOpts, path)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_VmSafe *VmSafeTransactor) CopyFile(opts *bind.TransactOpts, from string, to string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "copyFile", from, to)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_VmSafe *VmSafeSession) CopyFile(from string, to string) (*types.Transaction, error) {
	return _VmSafe.Contract.CopyFile(&_VmSafe.TransactOpts, from, to)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_VmSafe *VmSafeTransactorSession) CopyFile(from string, to string) (*types.Transaction, error) {
	return _VmSafe.Contract.CopyFile(&_VmSafe.TransactOpts, from, to)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_VmSafe *VmSafeTransactor) CreateDir(opts *bind.TransactOpts, path string, recursive bool) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "createDir", path, recursive)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_VmSafe *VmSafeSession) CreateDir(path string, recursive bool) (*types.Transaction, error) {
	return _VmSafe.Contract.CreateDir(&_VmSafe.TransactOpts, path, recursive)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_VmSafe *VmSafeTransactorSession) CreateDir(path string, recursive bool) (*types.Transaction, error) {
	return _VmSafe.Contract.CreateDir(&_VmSafe.TransactOpts, path, recursive)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_VmSafe *VmSafeTransactor) CreateWallet(opts *bind.TransactOpts, walletLabel string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "createWallet", walletLabel)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_VmSafe *VmSafeSession) CreateWallet(walletLabel string) (*types.Transaction, error) {
	return _VmSafe.Contract.CreateWallet(&_VmSafe.TransactOpts, walletLabel)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_VmSafe *VmSafeTransactorSession) CreateWallet(walletLabel string) (*types.Transaction, error) {
	return _VmSafe.Contract.CreateWallet(&_VmSafe.TransactOpts, walletLabel)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_VmSafe *VmSafeTransactor) CreateWallet0(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "createWallet0", privateKey)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_VmSafe *VmSafeSession) CreateWallet0(privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.CreateWallet0(&_VmSafe.TransactOpts, privateKey)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_VmSafe *VmSafeTransactorSession) CreateWallet0(privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.CreateWallet0(&_VmSafe.TransactOpts, privateKey)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_VmSafe *VmSafeTransactor) CreateWallet1(opts *bind.TransactOpts, privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "createWallet1", privateKey, walletLabel)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_VmSafe *VmSafeSession) CreateWallet1(privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _VmSafe.Contract.CreateWallet1(&_VmSafe.TransactOpts, privateKey, walletLabel)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_VmSafe *VmSafeTransactorSession) CreateWallet1(privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _VmSafe.Contract.CreateWallet1(&_VmSafe.TransactOpts, privateKey, walletLabel)
}

// DeployCode is a paid mutator transaction binding the contract method 0x29ce9dde.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs) returns(address deployedAddress)
func (_VmSafe *VmSafeTransactor) DeployCode(opts *bind.TransactOpts, artifactPath string, constructorArgs []byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "deployCode", artifactPath, constructorArgs)
}

// DeployCode is a paid mutator transaction binding the contract method 0x29ce9dde.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs) returns(address deployedAddress)
func (_VmSafe *VmSafeSession) DeployCode(artifactPath string, constructorArgs []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.DeployCode(&_VmSafe.TransactOpts, artifactPath, constructorArgs)
}

// DeployCode is a paid mutator transaction binding the contract method 0x29ce9dde.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs) returns(address deployedAddress)
func (_VmSafe *VmSafeTransactorSession) DeployCode(artifactPath string, constructorArgs []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.DeployCode(&_VmSafe.TransactOpts, artifactPath, constructorArgs)
}

// DeployCode0 is a paid mutator transaction binding the contract method 0x9a8325a0.
//
// Solidity: function deployCode(string artifactPath) returns(address deployedAddress)
func (_VmSafe *VmSafeTransactor) DeployCode0(opts *bind.TransactOpts, artifactPath string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "deployCode0", artifactPath)
}

// DeployCode0 is a paid mutator transaction binding the contract method 0x9a8325a0.
//
// Solidity: function deployCode(string artifactPath) returns(address deployedAddress)
func (_VmSafe *VmSafeSession) DeployCode0(artifactPath string) (*types.Transaction, error) {
	return _VmSafe.Contract.DeployCode0(&_VmSafe.TransactOpts, artifactPath)
}

// DeployCode0 is a paid mutator transaction binding the contract method 0x9a8325a0.
//
// Solidity: function deployCode(string artifactPath) returns(address deployedAddress)
func (_VmSafe *VmSafeTransactorSession) DeployCode0(artifactPath string) (*types.Transaction, error) {
	return _VmSafe.Contract.DeployCode0(&_VmSafe.TransactOpts, artifactPath)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_VmSafe *VmSafeTransactor) EthGetLogs(opts *bind.TransactOpts, fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "eth_getLogs", fromBlock, toBlock, target, topics)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_VmSafe *VmSafeSession) EthGetLogs(fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.EthGetLogs(&_VmSafe.TransactOpts, fromBlock, toBlock, target, topics)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_VmSafe *VmSafeTransactorSession) EthGetLogs(fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.EthGetLogs(&_VmSafe.TransactOpts, fromBlock, toBlock, target, topics)
}

// Exists is a paid mutator transaction binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) returns(bool result)
func (_VmSafe *VmSafeTransactor) Exists(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "exists", path)
}

// Exists is a paid mutator transaction binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) returns(bool result)
func (_VmSafe *VmSafeSession) Exists(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.Exists(&_VmSafe.TransactOpts, path)
}

// Exists is a paid mutator transaction binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) returns(bool result)
func (_VmSafe *VmSafeTransactorSession) Exists(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.Exists(&_VmSafe.TransactOpts, path)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_VmSafe *VmSafeTransactor) Ffi(opts *bind.TransactOpts, commandInput []string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "ffi", commandInput)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_VmSafe *VmSafeSession) Ffi(commandInput []string) (*types.Transaction, error) {
	return _VmSafe.Contract.Ffi(&_VmSafe.TransactOpts, commandInput)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_VmSafe *VmSafeTransactorSession) Ffi(commandInput []string) (*types.Transaction, error) {
	return _VmSafe.Contract.Ffi(&_VmSafe.TransactOpts, commandInput)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_VmSafe *VmSafeTransactor) GetMappingKeyAndParentOf(opts *bind.TransactOpts, target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "getMappingKeyAndParentOf", target, elementSlot)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_VmSafe *VmSafeSession) GetMappingKeyAndParentOf(target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.GetMappingKeyAndParentOf(&_VmSafe.TransactOpts, target, elementSlot)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_VmSafe *VmSafeTransactorSession) GetMappingKeyAndParentOf(target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.GetMappingKeyAndParentOf(&_VmSafe.TransactOpts, target, elementSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_VmSafe *VmSafeTransactor) GetMappingLength(opts *bind.TransactOpts, target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "getMappingLength", target, mappingSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_VmSafe *VmSafeSession) GetMappingLength(target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.GetMappingLength(&_VmSafe.TransactOpts, target, mappingSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_VmSafe *VmSafeTransactorSession) GetMappingLength(target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.GetMappingLength(&_VmSafe.TransactOpts, target, mappingSlot)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_VmSafe *VmSafeTransactor) GetMappingSlotAt(opts *bind.TransactOpts, target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "getMappingSlotAt", target, mappingSlot, idx)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_VmSafe *VmSafeSession) GetMappingSlotAt(target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.GetMappingSlotAt(&_VmSafe.TransactOpts, target, mappingSlot, idx)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_VmSafe *VmSafeTransactorSession) GetMappingSlotAt(target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.GetMappingSlotAt(&_VmSafe.TransactOpts, target, mappingSlot, idx)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_VmSafe *VmSafeTransactor) GetNonce0(opts *bind.TransactOpts, wallet VmSafeWallet) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "getNonce0", wallet)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_VmSafe *VmSafeSession) GetNonce0(wallet VmSafeWallet) (*types.Transaction, error) {
	return _VmSafe.Contract.GetNonce0(&_VmSafe.TransactOpts, wallet)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_VmSafe *VmSafeTransactorSession) GetNonce0(wallet VmSafeWallet) (*types.Transaction, error) {
	return _VmSafe.Contract.GetNonce0(&_VmSafe.TransactOpts, wallet)
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_VmSafe *VmSafeTransactor) GetRecordedLogs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "getRecordedLogs")
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_VmSafe *VmSafeSession) GetRecordedLogs() (*types.Transaction, error) {
	return _VmSafe.Contract.GetRecordedLogs(&_VmSafe.TransactOpts)
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_VmSafe *VmSafeTransactorSession) GetRecordedLogs() (*types.Transaction, error) {
	return _VmSafe.Contract.GetRecordedLogs(&_VmSafe.TransactOpts)
}

// IsDir is a paid mutator transaction binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) returns(bool result)
func (_VmSafe *VmSafeTransactor) IsDir(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "isDir", path)
}

// IsDir is a paid mutator transaction binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) returns(bool result)
func (_VmSafe *VmSafeSession) IsDir(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.IsDir(&_VmSafe.TransactOpts, path)
}

// IsDir is a paid mutator transaction binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) returns(bool result)
func (_VmSafe *VmSafeTransactorSession) IsDir(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.IsDir(&_VmSafe.TransactOpts, path)
}

// IsFile is a paid mutator transaction binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) returns(bool result)
func (_VmSafe *VmSafeTransactor) IsFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "isFile", path)
}

// IsFile is a paid mutator transaction binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) returns(bool result)
func (_VmSafe *VmSafeSession) IsFile(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.IsFile(&_VmSafe.TransactOpts, path)
}

// IsFile is a paid mutator transaction binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) returns(bool result)
func (_VmSafe *VmSafeTransactorSession) IsFile(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.IsFile(&_VmSafe.TransactOpts, path)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_VmSafe *VmSafeTransactor) Label(opts *bind.TransactOpts, account common.Address, newLabel string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "label", account, newLabel)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_VmSafe *VmSafeSession) Label(account common.Address, newLabel string) (*types.Transaction, error) {
	return _VmSafe.Contract.Label(&_VmSafe.TransactOpts, account, newLabel)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_VmSafe *VmSafeTransactorSession) Label(account common.Address, newLabel string) (*types.Transaction, error) {
	return _VmSafe.Contract.Label(&_VmSafe.TransactOpts, account, newLabel)
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_VmSafe *VmSafeTransactor) PauseGasMetering(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "pauseGasMetering")
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_VmSafe *VmSafeSession) PauseGasMetering() (*types.Transaction, error) {
	return _VmSafe.Contract.PauseGasMetering(&_VmSafe.TransactOpts)
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_VmSafe *VmSafeTransactorSession) PauseGasMetering() (*types.Transaction, error) {
	return _VmSafe.Contract.PauseGasMetering(&_VmSafe.TransactOpts)
}

// Prompt is a paid mutator transaction binding the contract method 0x47eaf474.
//
// Solidity: function prompt(string promptText) returns(string input)
func (_VmSafe *VmSafeTransactor) Prompt(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "prompt", promptText)
}

// Prompt is a paid mutator transaction binding the contract method 0x47eaf474.
//
// Solidity: function prompt(string promptText) returns(string input)
func (_VmSafe *VmSafeSession) Prompt(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.Prompt(&_VmSafe.TransactOpts, promptText)
}

// Prompt is a paid mutator transaction binding the contract method 0x47eaf474.
//
// Solidity: function prompt(string promptText) returns(string input)
func (_VmSafe *VmSafeTransactorSession) Prompt(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.Prompt(&_VmSafe.TransactOpts, promptText)
}

// PromptAddress is a paid mutator transaction binding the contract method 0x62ee05f4.
//
// Solidity: function promptAddress(string promptText) returns(address)
func (_VmSafe *VmSafeTransactor) PromptAddress(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "promptAddress", promptText)
}

// PromptAddress is a paid mutator transaction binding the contract method 0x62ee05f4.
//
// Solidity: function promptAddress(string promptText) returns(address)
func (_VmSafe *VmSafeSession) PromptAddress(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.PromptAddress(&_VmSafe.TransactOpts, promptText)
}

// PromptAddress is a paid mutator transaction binding the contract method 0x62ee05f4.
//
// Solidity: function promptAddress(string promptText) returns(address)
func (_VmSafe *VmSafeTransactorSession) PromptAddress(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.PromptAddress(&_VmSafe.TransactOpts, promptText)
}

// PromptSecret is a paid mutator transaction binding the contract method 0x1e279d41.
//
// Solidity: function promptSecret(string promptText) returns(string input)
func (_VmSafe *VmSafeTransactor) PromptSecret(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "promptSecret", promptText)
}

// PromptSecret is a paid mutator transaction binding the contract method 0x1e279d41.
//
// Solidity: function promptSecret(string promptText) returns(string input)
func (_VmSafe *VmSafeSession) PromptSecret(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.PromptSecret(&_VmSafe.TransactOpts, promptText)
}

// PromptSecret is a paid mutator transaction binding the contract method 0x1e279d41.
//
// Solidity: function promptSecret(string promptText) returns(string input)
func (_VmSafe *VmSafeTransactorSession) PromptSecret(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.PromptSecret(&_VmSafe.TransactOpts, promptText)
}

// PromptSecretUint is a paid mutator transaction binding the contract method 0x69ca02b7.
//
// Solidity: function promptSecretUint(string promptText) returns(uint256)
func (_VmSafe *VmSafeTransactor) PromptSecretUint(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "promptSecretUint", promptText)
}

// PromptSecretUint is a paid mutator transaction binding the contract method 0x69ca02b7.
//
// Solidity: function promptSecretUint(string promptText) returns(uint256)
func (_VmSafe *VmSafeSession) PromptSecretUint(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.PromptSecretUint(&_VmSafe.TransactOpts, promptText)
}

// PromptSecretUint is a paid mutator transaction binding the contract method 0x69ca02b7.
//
// Solidity: function promptSecretUint(string promptText) returns(uint256)
func (_VmSafe *VmSafeTransactorSession) PromptSecretUint(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.PromptSecretUint(&_VmSafe.TransactOpts, promptText)
}

// PromptUint is a paid mutator transaction binding the contract method 0x652fd489.
//
// Solidity: function promptUint(string promptText) returns(uint256)
func (_VmSafe *VmSafeTransactor) PromptUint(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "promptUint", promptText)
}

// PromptUint is a paid mutator transaction binding the contract method 0x652fd489.
//
// Solidity: function promptUint(string promptText) returns(uint256)
func (_VmSafe *VmSafeSession) PromptUint(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.PromptUint(&_VmSafe.TransactOpts, promptText)
}

// PromptUint is a paid mutator transaction binding the contract method 0x652fd489.
//
// Solidity: function promptUint(string promptText) returns(uint256)
func (_VmSafe *VmSafeTransactorSession) PromptUint(promptText string) (*types.Transaction, error) {
	return _VmSafe.Contract.PromptUint(&_VmSafe.TransactOpts, promptText)
}

// RandomAddress is a paid mutator transaction binding the contract method 0xd5bee9f5.
//
// Solidity: function randomAddress() returns(address)
func (_VmSafe *VmSafeTransactor) RandomAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "randomAddress")
}

// RandomAddress is a paid mutator transaction binding the contract method 0xd5bee9f5.
//
// Solidity: function randomAddress() returns(address)
func (_VmSafe *VmSafeSession) RandomAddress() (*types.Transaction, error) {
	return _VmSafe.Contract.RandomAddress(&_VmSafe.TransactOpts)
}

// RandomAddress is a paid mutator transaction binding the contract method 0xd5bee9f5.
//
// Solidity: function randomAddress() returns(address)
func (_VmSafe *VmSafeTransactorSession) RandomAddress() (*types.Transaction, error) {
	return _VmSafe.Contract.RandomAddress(&_VmSafe.TransactOpts)
}

// RandomUint is a paid mutator transaction binding the contract method 0x25124730.
//
// Solidity: function randomUint() returns(uint256)
func (_VmSafe *VmSafeTransactor) RandomUint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "randomUint")
}

// RandomUint is a paid mutator transaction binding the contract method 0x25124730.
//
// Solidity: function randomUint() returns(uint256)
func (_VmSafe *VmSafeSession) RandomUint() (*types.Transaction, error) {
	return _VmSafe.Contract.RandomUint(&_VmSafe.TransactOpts)
}

// RandomUint is a paid mutator transaction binding the contract method 0x25124730.
//
// Solidity: function randomUint() returns(uint256)
func (_VmSafe *VmSafeTransactorSession) RandomUint() (*types.Transaction, error) {
	return _VmSafe.Contract.RandomUint(&_VmSafe.TransactOpts)
}

// RandomUint0 is a paid mutator transaction binding the contract method 0xd61b051b.
//
// Solidity: function randomUint(uint256 min, uint256 max) returns(uint256)
func (_VmSafe *VmSafeTransactor) RandomUint0(opts *bind.TransactOpts, min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "randomUint0", min, max)
}

// RandomUint0 is a paid mutator transaction binding the contract method 0xd61b051b.
//
// Solidity: function randomUint(uint256 min, uint256 max) returns(uint256)
func (_VmSafe *VmSafeSession) RandomUint0(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.RandomUint0(&_VmSafe.TransactOpts, min, max)
}

// RandomUint0 is a paid mutator transaction binding the contract method 0xd61b051b.
//
// Solidity: function randomUint(uint256 min, uint256 max) returns(uint256)
func (_VmSafe *VmSafeTransactorSession) RandomUint0(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.RandomUint0(&_VmSafe.TransactOpts, min, max)
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_VmSafe *VmSafeTransactor) Record(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "record")
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_VmSafe *VmSafeSession) Record() (*types.Transaction, error) {
	return _VmSafe.Contract.Record(&_VmSafe.TransactOpts)
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_VmSafe *VmSafeTransactorSession) Record() (*types.Transaction, error) {
	return _VmSafe.Contract.Record(&_VmSafe.TransactOpts)
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_VmSafe *VmSafeTransactor) RecordLogs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "recordLogs")
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_VmSafe *VmSafeSession) RecordLogs() (*types.Transaction, error) {
	return _VmSafe.Contract.RecordLogs(&_VmSafe.TransactOpts)
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_VmSafe *VmSafeTransactorSession) RecordLogs() (*types.Transaction, error) {
	return _VmSafe.Contract.RecordLogs(&_VmSafe.TransactOpts)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_VmSafe *VmSafeTransactor) RememberKey(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "rememberKey", privateKey)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_VmSafe *VmSafeSession) RememberKey(privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.RememberKey(&_VmSafe.TransactOpts, privateKey)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_VmSafe *VmSafeTransactorSession) RememberKey(privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.RememberKey(&_VmSafe.TransactOpts, privateKey)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_VmSafe *VmSafeTransactor) RemoveDir(opts *bind.TransactOpts, path string, recursive bool) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "removeDir", path, recursive)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_VmSafe *VmSafeSession) RemoveDir(path string, recursive bool) (*types.Transaction, error) {
	return _VmSafe.Contract.RemoveDir(&_VmSafe.TransactOpts, path, recursive)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_VmSafe *VmSafeTransactorSession) RemoveDir(path string, recursive bool) (*types.Transaction, error) {
	return _VmSafe.Contract.RemoveDir(&_VmSafe.TransactOpts, path, recursive)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_VmSafe *VmSafeTransactor) RemoveFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "removeFile", path)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_VmSafe *VmSafeSession) RemoveFile(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.RemoveFile(&_VmSafe.TransactOpts, path)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_VmSafe *VmSafeTransactorSession) RemoveFile(path string) (*types.Transaction, error) {
	return _VmSafe.Contract.RemoveFile(&_VmSafe.TransactOpts, path)
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_VmSafe *VmSafeTransactor) ResumeGasMetering(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "resumeGasMetering")
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_VmSafe *VmSafeSession) ResumeGasMetering() (*types.Transaction, error) {
	return _VmSafe.Contract.ResumeGasMetering(&_VmSafe.TransactOpts)
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_VmSafe *VmSafeTransactorSession) ResumeGasMetering() (*types.Transaction, error) {
	return _VmSafe.Contract.ResumeGasMetering(&_VmSafe.TransactOpts)
}

// Rpc is a paid mutator transaction binding the contract method 0x0199a220.
//
// Solidity: function rpc(string urlOrAlias, string method, string params) returns(bytes data)
func (_VmSafe *VmSafeTransactor) Rpc(opts *bind.TransactOpts, urlOrAlias string, method string, params string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "rpc", urlOrAlias, method, params)
}

// Rpc is a paid mutator transaction binding the contract method 0x0199a220.
//
// Solidity: function rpc(string urlOrAlias, string method, string params) returns(bytes data)
func (_VmSafe *VmSafeSession) Rpc(urlOrAlias string, method string, params string) (*types.Transaction, error) {
	return _VmSafe.Contract.Rpc(&_VmSafe.TransactOpts, urlOrAlias, method, params)
}

// Rpc is a paid mutator transaction binding the contract method 0x0199a220.
//
// Solidity: function rpc(string urlOrAlias, string method, string params) returns(bytes data)
func (_VmSafe *VmSafeTransactorSession) Rpc(urlOrAlias string, method string, params string) (*types.Transaction, error) {
	return _VmSafe.Contract.Rpc(&_VmSafe.TransactOpts, urlOrAlias, method, params)
}

// Rpc0 is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_VmSafe *VmSafeTransactor) Rpc0(opts *bind.TransactOpts, method string, params string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "rpc0", method, params)
}

// Rpc0 is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_VmSafe *VmSafeSession) Rpc0(method string, params string) (*types.Transaction, error) {
	return _VmSafe.Contract.Rpc0(&_VmSafe.TransactOpts, method, params)
}

// Rpc0 is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_VmSafe *VmSafeTransactorSession) Rpc0(method string, params string) (*types.Transaction, error) {
	return _VmSafe.Contract.Rpc0(&_VmSafe.TransactOpts, method, params)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeAddress(opts *bind.TransactOpts, objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeAddress", objectKey, valueKey, values)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_VmSafe *VmSafeSession) SerializeAddress(objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeAddress(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeAddress(objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeAddress(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeAddress0(opts *bind.TransactOpts, objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeAddress0", objectKey, valueKey, value)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeAddress0(objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeAddress0(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeAddress0(objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeAddress0(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeBool(opts *bind.TransactOpts, objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeBool", objectKey, valueKey, values)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_VmSafe *VmSafeSession) SerializeBool(objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBool(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeBool(objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBool(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeBool0(opts *bind.TransactOpts, objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeBool0", objectKey, valueKey, value)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeBool0(objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBool0(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeBool0(objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBool0(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeBytes(opts *bind.TransactOpts, objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeBytes", objectKey, valueKey, values)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_VmSafe *VmSafeSession) SerializeBytes(objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBytes(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeBytes(objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBytes(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeBytes0(opts *bind.TransactOpts, objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeBytes0", objectKey, valueKey, value)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeBytes0(objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBytes0(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeBytes0(objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBytes0(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeBytes32(opts *bind.TransactOpts, objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeBytes32", objectKey, valueKey, values)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_VmSafe *VmSafeSession) SerializeBytes32(objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBytes32(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeBytes32(objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBytes32(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeBytes320(opts *bind.TransactOpts, objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeBytes320", objectKey, valueKey, value)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeBytes320(objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBytes320(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeBytes320(objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeBytes320(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeInt(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeInt", objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeInt(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeInt(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeInt(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeInt(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeInt0(opts *bind.TransactOpts, objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeInt0", objectKey, valueKey, values)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_VmSafe *VmSafeSession) SerializeInt0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeInt0(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeInt0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeInt0(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeJson(opts *bind.TransactOpts, objectKey string, value string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeJson", objectKey, value)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeJson(objectKey string, value string) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeJson(&_VmSafe.TransactOpts, objectKey, value)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeJson(objectKey string, value string) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeJson(&_VmSafe.TransactOpts, objectKey, value)
}

// SerializeJsonType0 is a paid mutator transaction binding the contract method 0x6f93bccb.
//
// Solidity: function serializeJsonType(string objectKey, string valueKey, string typeDescription, bytes value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeJsonType0(opts *bind.TransactOpts, objectKey string, valueKey string, typeDescription string, value []byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeJsonType0", objectKey, valueKey, typeDescription, value)
}

// SerializeJsonType0 is a paid mutator transaction binding the contract method 0x6f93bccb.
//
// Solidity: function serializeJsonType(string objectKey, string valueKey, string typeDescription, bytes value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeJsonType0(objectKey string, valueKey string, typeDescription string, value []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeJsonType0(&_VmSafe.TransactOpts, objectKey, valueKey, typeDescription, value)
}

// SerializeJsonType0 is a paid mutator transaction binding the contract method 0x6f93bccb.
//
// Solidity: function serializeJsonType(string objectKey, string valueKey, string typeDescription, bytes value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeJsonType0(objectKey string, valueKey string, typeDescription string, value []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeJsonType0(&_VmSafe.TransactOpts, objectKey, valueKey, typeDescription, value)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeString(opts *bind.TransactOpts, objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeString", objectKey, valueKey, values)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_VmSafe *VmSafeSession) SerializeString(objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeString(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeString(objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeString(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeString0(opts *bind.TransactOpts, objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeString0", objectKey, valueKey, value)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeString0(objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeString0(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeString0(objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeString0(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeUint(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeUint", objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeUint(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeUint(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeUint(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeUint(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeUint0(opts *bind.TransactOpts, objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeUint0", objectKey, valueKey, values)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_VmSafe *VmSafeSession) SerializeUint0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeUint0(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeUint0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeUint0(&_VmSafe.TransactOpts, objectKey, valueKey, values)
}

// SerializeUintToHex is a paid mutator transaction binding the contract method 0xae5a2ae8.
//
// Solidity: function serializeUintToHex(string objectKey, string valueKey, uint256 value) returns(string json)
func (_VmSafe *VmSafeTransactor) SerializeUintToHex(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "serializeUintToHex", objectKey, valueKey, value)
}

// SerializeUintToHex is a paid mutator transaction binding the contract method 0xae5a2ae8.
//
// Solidity: function serializeUintToHex(string objectKey, string valueKey, uint256 value) returns(string json)
func (_VmSafe *VmSafeSession) SerializeUintToHex(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeUintToHex(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SerializeUintToHex is a paid mutator transaction binding the contract method 0xae5a2ae8.
//
// Solidity: function serializeUintToHex(string objectKey, string valueKey, uint256 value) returns(string json)
func (_VmSafe *VmSafeTransactorSession) SerializeUintToHex(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.SerializeUintToHex(&_VmSafe.TransactOpts, objectKey, valueKey, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_VmSafe *VmSafeTransactor) SetEnv(opts *bind.TransactOpts, name string, value string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "setEnv", name, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_VmSafe *VmSafeSession) SetEnv(name string, value string) (*types.Transaction, error) {
	return _VmSafe.Contract.SetEnv(&_VmSafe.TransactOpts, name, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_VmSafe *VmSafeTransactorSession) SetEnv(name string, value string) (*types.Transaction, error) {
	return _VmSafe.Contract.SetEnv(&_VmSafe.TransactOpts, name, value)
}

// Sign1 is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeTransactor) Sign1(opts *bind.TransactOpts, wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "sign1", wallet, digest)
}

// Sign1 is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeSession) Sign1(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.Sign1(&_VmSafe.TransactOpts, wallet, digest)
}

// Sign1 is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_VmSafe *VmSafeTransactorSession) Sign1(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.Sign1(&_VmSafe.TransactOpts, wallet, digest)
}

// SignCompact is a paid mutator transaction binding the contract method 0x3d0e292f.
//
// Solidity: function signCompact((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeTransactor) SignCompact(opts *bind.TransactOpts, wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "signCompact", wallet, digest)
}

// SignCompact is a paid mutator transaction binding the contract method 0x3d0e292f.
//
// Solidity: function signCompact((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeSession) SignCompact(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SignCompact(&_VmSafe.TransactOpts, wallet, digest)
}

// SignCompact is a paid mutator transaction binding the contract method 0x3d0e292f.
//
// Solidity: function signCompact((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(bytes32 r, bytes32 vs)
func (_VmSafe *VmSafeTransactorSession) SignCompact(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _VmSafe.Contract.SignCompact(&_VmSafe.TransactOpts, wallet, digest)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_VmSafe *VmSafeTransactor) Sleep(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "sleep", duration)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_VmSafe *VmSafeSession) Sleep(duration *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.Sleep(&_VmSafe.TransactOpts, duration)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_VmSafe *VmSafeTransactorSession) Sleep(duration *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.Sleep(&_VmSafe.TransactOpts, duration)
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_VmSafe *VmSafeTransactor) StartBroadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "startBroadcast")
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_VmSafe *VmSafeSession) StartBroadcast() (*types.Transaction, error) {
	return _VmSafe.Contract.StartBroadcast(&_VmSafe.TransactOpts)
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_VmSafe *VmSafeTransactorSession) StartBroadcast() (*types.Transaction, error) {
	return _VmSafe.Contract.StartBroadcast(&_VmSafe.TransactOpts)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_VmSafe *VmSafeTransactor) StartBroadcast0(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "startBroadcast0", signer)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_VmSafe *VmSafeSession) StartBroadcast0(signer common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.StartBroadcast0(&_VmSafe.TransactOpts, signer)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_VmSafe *VmSafeTransactorSession) StartBroadcast0(signer common.Address) (*types.Transaction, error) {
	return _VmSafe.Contract.StartBroadcast0(&_VmSafe.TransactOpts, signer)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_VmSafe *VmSafeTransactor) StartBroadcast1(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "startBroadcast1", privateKey)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_VmSafe *VmSafeSession) StartBroadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.StartBroadcast1(&_VmSafe.TransactOpts, privateKey)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_VmSafe *VmSafeTransactorSession) StartBroadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _VmSafe.Contract.StartBroadcast1(&_VmSafe.TransactOpts, privateKey)
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_VmSafe *VmSafeTransactor) StartMappingRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "startMappingRecording")
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_VmSafe *VmSafeSession) StartMappingRecording() (*types.Transaction, error) {
	return _VmSafe.Contract.StartMappingRecording(&_VmSafe.TransactOpts)
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_VmSafe *VmSafeTransactorSession) StartMappingRecording() (*types.Transaction, error) {
	return _VmSafe.Contract.StartMappingRecording(&_VmSafe.TransactOpts)
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_VmSafe *VmSafeTransactor) StartStateDiffRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "startStateDiffRecording")
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_VmSafe *VmSafeSession) StartStateDiffRecording() (*types.Transaction, error) {
	return _VmSafe.Contract.StartStateDiffRecording(&_VmSafe.TransactOpts)
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_VmSafe *VmSafeTransactorSession) StartStateDiffRecording() (*types.Transaction, error) {
	return _VmSafe.Contract.StartStateDiffRecording(&_VmSafe.TransactOpts)
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[],uint64)[] accountAccesses)
func (_VmSafe *VmSafeTransactor) StopAndReturnStateDiff(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "stopAndReturnStateDiff")
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[],uint64)[] accountAccesses)
func (_VmSafe *VmSafeSession) StopAndReturnStateDiff() (*types.Transaction, error) {
	return _VmSafe.Contract.StopAndReturnStateDiff(&_VmSafe.TransactOpts)
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[],uint64)[] accountAccesses)
func (_VmSafe *VmSafeTransactorSession) StopAndReturnStateDiff() (*types.Transaction, error) {
	return _VmSafe.Contract.StopAndReturnStateDiff(&_VmSafe.TransactOpts)
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_VmSafe *VmSafeTransactor) StopBroadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "stopBroadcast")
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_VmSafe *VmSafeSession) StopBroadcast() (*types.Transaction, error) {
	return _VmSafe.Contract.StopBroadcast(&_VmSafe.TransactOpts)
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_VmSafe *VmSafeTransactorSession) StopBroadcast() (*types.Transaction, error) {
	return _VmSafe.Contract.StopBroadcast(&_VmSafe.TransactOpts)
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_VmSafe *VmSafeTransactor) StopMappingRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "stopMappingRecording")
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_VmSafe *VmSafeSession) StopMappingRecording() (*types.Transaction, error) {
	return _VmSafe.Contract.StopMappingRecording(&_VmSafe.TransactOpts)
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_VmSafe *VmSafeTransactorSession) StopMappingRecording() (*types.Transaction, error) {
	return _VmSafe.Contract.StopMappingRecording(&_VmSafe.TransactOpts)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_VmSafe *VmSafeTransactor) TryFfi(opts *bind.TransactOpts, commandInput []string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "tryFfi", commandInput)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_VmSafe *VmSafeSession) TryFfi(commandInput []string) (*types.Transaction, error) {
	return _VmSafe.Contract.TryFfi(&_VmSafe.TransactOpts, commandInput)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_VmSafe *VmSafeTransactorSession) TryFfi(commandInput []string) (*types.Transaction, error) {
	return _VmSafe.Contract.TryFfi(&_VmSafe.TransactOpts, commandInput)
}

// UnixTime is a paid mutator transaction binding the contract method 0x625387dc.
//
// Solidity: function unixTime() returns(uint256 milliseconds)
func (_VmSafe *VmSafeTransactor) UnixTime(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "unixTime")
}

// UnixTime is a paid mutator transaction binding the contract method 0x625387dc.
//
// Solidity: function unixTime() returns(uint256 milliseconds)
func (_VmSafe *VmSafeSession) UnixTime() (*types.Transaction, error) {
	return _VmSafe.Contract.UnixTime(&_VmSafe.TransactOpts)
}

// UnixTime is a paid mutator transaction binding the contract method 0x625387dc.
//
// Solidity: function unixTime() returns(uint256 milliseconds)
func (_VmSafe *VmSafeTransactorSession) UnixTime() (*types.Transaction, error) {
	return _VmSafe.Contract.UnixTime(&_VmSafe.TransactOpts)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_VmSafe *VmSafeTransactor) WriteFile(opts *bind.TransactOpts, path string, data string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "writeFile", path, data)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_VmSafe *VmSafeSession) WriteFile(path string, data string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteFile(&_VmSafe.TransactOpts, path, data)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_VmSafe *VmSafeTransactorSession) WriteFile(path string, data string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteFile(&_VmSafe.TransactOpts, path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_VmSafe *VmSafeTransactor) WriteFileBinary(opts *bind.TransactOpts, path string, data []byte) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "writeFileBinary", path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_VmSafe *VmSafeSession) WriteFileBinary(path string, data []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteFileBinary(&_VmSafe.TransactOpts, path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_VmSafe *VmSafeTransactorSession) WriteFileBinary(path string, data []byte) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteFileBinary(&_VmSafe.TransactOpts, path, data)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_VmSafe *VmSafeTransactor) WriteJson(opts *bind.TransactOpts, json string, path string, valueKey string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "writeJson", json, path, valueKey)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_VmSafe *VmSafeSession) WriteJson(json string, path string, valueKey string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteJson(&_VmSafe.TransactOpts, json, path, valueKey)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_VmSafe *VmSafeTransactorSession) WriteJson(json string, path string, valueKey string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteJson(&_VmSafe.TransactOpts, json, path, valueKey)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_VmSafe *VmSafeTransactor) WriteJson0(opts *bind.TransactOpts, json string, path string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "writeJson0", json, path)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_VmSafe *VmSafeSession) WriteJson0(json string, path string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteJson0(&_VmSafe.TransactOpts, json, path)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_VmSafe *VmSafeTransactorSession) WriteJson0(json string, path string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteJson0(&_VmSafe.TransactOpts, json, path)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_VmSafe *VmSafeTransactor) WriteLine(opts *bind.TransactOpts, path string, data string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "writeLine", path, data)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_VmSafe *VmSafeSession) WriteLine(path string, data string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteLine(&_VmSafe.TransactOpts, path, data)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_VmSafe *VmSafeTransactorSession) WriteLine(path string, data string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteLine(&_VmSafe.TransactOpts, path, data)
}

// WriteToml is a paid mutator transaction binding the contract method 0x51ac6a33.
//
// Solidity: function writeToml(string json, string path, string valueKey) returns()
func (_VmSafe *VmSafeTransactor) WriteToml(opts *bind.TransactOpts, json string, path string, valueKey string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "writeToml", json, path, valueKey)
}

// WriteToml is a paid mutator transaction binding the contract method 0x51ac6a33.
//
// Solidity: function writeToml(string json, string path, string valueKey) returns()
func (_VmSafe *VmSafeSession) WriteToml(json string, path string, valueKey string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteToml(&_VmSafe.TransactOpts, json, path, valueKey)
}

// WriteToml is a paid mutator transaction binding the contract method 0x51ac6a33.
//
// Solidity: function writeToml(string json, string path, string valueKey) returns()
func (_VmSafe *VmSafeTransactorSession) WriteToml(json string, path string, valueKey string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteToml(&_VmSafe.TransactOpts, json, path, valueKey)
}

// WriteToml0 is a paid mutator transaction binding the contract method 0xc0865ba7.
//
// Solidity: function writeToml(string json, string path) returns()
func (_VmSafe *VmSafeTransactor) WriteToml0(opts *bind.TransactOpts, json string, path string) (*types.Transaction, error) {
	return _VmSafe.contract.Transact(opts, "writeToml0", json, path)
}

// WriteToml0 is a paid mutator transaction binding the contract method 0xc0865ba7.
//
// Solidity: function writeToml(string json, string path) returns()
func (_VmSafe *VmSafeSession) WriteToml0(json string, path string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteToml0(&_VmSafe.TransactOpts, json, path)
}

// WriteToml0 is a paid mutator transaction binding the contract method 0xc0865ba7.
//
// Solidity: function writeToml(string json, string path) returns()
func (_VmSafe *VmSafeTransactorSession) WriteToml0(json string, path string) (*types.Transaction, error) {
	return _VmSafe.Contract.WriteToml0(&_VmSafe.TransactOpts, json, path)
}
