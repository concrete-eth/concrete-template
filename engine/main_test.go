package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/concrete"
	testtool "github.com/ethereum/go-ethereum/concrete/testtool"
)

func TestE2E(t *testing.T) {
	engine := concrete.ConcreteGeth
	setup(engine)
	_, fails := testtool.Test(testtool.TestConfig{
		TestDir: "../sol/test",
		OutDir:  "../out",
	})
	if fails > 0 {
		t.Errorf("TestE2E failed with %d failures", fails)
	}
}
