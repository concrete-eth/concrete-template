package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/cmd/geth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/concrete"
	"github.com/ethereum/go-ethereum/concrete/api"
	"github.com/ethereum/go-ethereum/concrete/lib"
	"github.com/ethereum/go-ethereum/concrete/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/urfave/cli/v2"
)

//go:embed ICounter.abi.json
var ABIJsonStr string

var (
	CounterKey = crypto.Keccak256([]byte("template.counter.v1"))
	ABI, _     = abi.JSON(strings.NewReader(ABIJsonStr))
)

type Counter struct {
	lib.BlankPrecompile
}

func (c *Counter) IsStatic(input []byte) bool {
	if method, err := ABI.MethodById(input[:4]); err == nil {
		return method.IsConstant()
	}
	return false
}

func (c *Counter) Run(API api.Environment, input []byte) ([]byte, error) {
	id, data := utils.SplitInput(input)
	method, err := ABI.MethodById(id)
	if err != nil {
		return nil, fmt.Errorf("method not found")
	}

	ds := lib.NewDatastore(API)
	number := ds.Get(CounterKey)

	switch method.Name {
	case "number":
		return number.Bytes32().Bytes(), nil
	case "setNumber":
		number.SetBytes32(common.BytesToHash(data))
		return nil, nil
	case "increment":
		numberUint254 := number.Uint256()
		numberUint254.Add(numberUint254, new(uint256.Int).SetUint64(1))
		number.SetUint256(numberUint254)
		return nil, nil
	}

	return nil, fmt.Errorf("method not found")
}

func newRegistry() concrete.PrecompileRegistry {
	var (
		addr  = common.HexToAddress("0x80")
		pc    = &Counter{}
		block = uint64(0)
	)
	registry := concrete.NewRegistry()
	registry.AddPrecompile(block, addr, pc)
	return registry
}

func newGeth() *cli.App {
	registry := newRegistry()
	return geth.NewConcreteGethApp(registry)
}

func main() {
	app := newGeth()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
