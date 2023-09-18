package main

import (
	"testing"

	testtool "github.com/ethereum/go-ethereum/concrete/testtool"
)

func TestE2E(t *testing.T) {
	registry := newRegistry()
	_, fails := testtool.Test(registry, testtool.TestConfig{
		TestDir: "../sol/test",
		OutDir:  "../out",
	})
	if fails > 0 {
		t.Errorf("TestE2E failed with %d failures", fails)
	}
}
