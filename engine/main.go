package main

import (
	"github.com/concrete-eth/concrete-template/engine/pcs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/concrete"
)

func setup(engine concrete.ConcreteApp) {
	addr := common.HexToAddress("0x80")
	pc := &pcs.CounterPC{}
	engine.AddPrecompile(addr, pc)
}

func main() {
	engine := concrete.ConcreteGeth
	setup(engine)
	engine.Run()
}
