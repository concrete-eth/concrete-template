package main

import (
	"fmt"
	"os"

	"github.com/concrete-eth/concrete-template/engine/pcs"
	"github.com/ethereum/go-ethereum/cmd/geth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/concrete"
	"github.com/urfave/cli/v2"
)

func newRegistry() concrete.PrecompileRegistry {
	var (
		addr  = common.HexToAddress("0x80")
		pc    = &pcs.CounterPrecompile{}
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
