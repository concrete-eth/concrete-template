package main

import (
	"testing"

	testtool "github.com/ethereum/go-ethereum/concrete/testtool"
)

func TestE2E(t *testing.T) {
	registry := newRegistry()
	testtool.Test(t, registry, testtool.TestConfig{
		Contract: "../sol/TestCounter.sol:CounterTest",
		OutDir:   "../out",
	})
}
