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

func TestCounterPC(t *testing.T) {
	var (
		r    = require.New(t)
		addr = common.HexToAddress("0x1234")
		env  = mock.NewMockEnvironment(addr, api.EnvConfig{}, false, 0)
		pc   = &CounterPC{}
	)
	number := func() int64 {
		input, err := CounterABI.Pack("number")
		if err != nil {
			t.Fatal(err)
		}
		output, err := pc.Run(env, input)
		if err != nil {
			t.Fatal(err)
		}
		returns, err := CounterABI.Unpack("number", output)
		if err != nil {
			t.Fatal(err)
		}
		return returns[0].(*big.Int).Int64()
	}
	setNumber := func(val int64) {
		input, err := CounterABI.Pack("setNumber", big.NewInt(val))
		if err != nil {
			t.Fatal(err)
		}
		output, err := pc.Run(env, input)
		if err != nil {
			t.Fatal(err)
		}
		_, err = CounterABI.Unpack("setNumber", output)
		if err != nil {
			t.Fatal(err)
		}
	}
	increment := func() {
		input, err := CounterABI.Pack("increment")
		if err != nil {
			t.Fatal(err)
		}
		output, err := pc.Run(env, input)
		if err != nil {
			t.Fatal(err)
		}
		_, err = CounterABI.Unpack("increment", output)
		if err != nil {
			t.Fatal(err)
		}
	}

	r.Equal(int64(0), number())
	setNumber(10)
	r.Equal(int64(10), number())
	increment()
	r.Equal(int64(11), number())
}
