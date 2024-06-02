package pcs

import (
	_ "embed"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/concrete/api"
	"github.com/ethereum/go-ethereum/concrete/mock"
	"github.com/stretchr/testify/require"
)

func TestCounterPrecompile(t *testing.T) {
	var (
		r    = require.New(t)
		addr = common.HexToAddress("0x1234")
		env  = mock.NewMockEnvironment(addr, api.EnvConfig{}, false, 0)
		pc   = &CounterPrecompile{}
	)
	
	getCount := func() *big.Int {
		input, err := CounterABI.Pack("getCount")
		if err != nil {
			t.Fatal(err)
		}
		output, err := pc.Run(env, input)
		if err != nil {
			t.Fatal(err)
		}
		returns, err := CounterABI.Unpack("getCount", output)
		if err != nil {
			t.Fatal(err)
		}
		count := returns[0].(*big.Int)
		return count
	}

	increment := func() {
		input, err := CounterABI.Pack("increment")
		if err != nil {
			t.Fatal(err)
		}
		_, err = pc.Run(env, input)
		if err != nil {
			t.Fatal(err)
		}
	}

	r.Equal(getCount().Int64(), int64(0))
	increment()
	r.Equal(getCount().Int64(), int64(1))
	increment()
	r.Equal(getCount().Int64(), int64(2))
}
