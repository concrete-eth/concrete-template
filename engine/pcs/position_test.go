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

func TestPositionPrecompile(t *testing.T) {
	var (
		r    = require.New(t)
		addr = common.HexToAddress("0x1234")
		env  = mock.NewMockEnvironment(addr, api.EnvConfig{}, false, 0)
		pc   = &PositionPrecompile{}
	)
	setPosition := func(entity int64, coord Coord) {
		input, err := PositionABI.Pack("setPosition", big.NewInt(entity), coord)
		if err != nil {
			t.Fatal(err)
		}
		_, err = pc.Run(env, input)
		if err != nil {
			t.Fatal(err)
		}
	}
	getPosition := func(entity int64) Coord {
		input, err := PositionABI.Pack("getPosition", big.NewInt(entity))
		if err != nil {
			t.Fatal(err)
		}
		output, err := pc.Run(env, input)
		if err != nil {
			t.Fatal(err)
		}
		returns, err := PositionABI.Unpack("getPosition", output)
		if err != nil {
			t.Fatal(err)
		}
		coord := returns[0].(Coord)
		return coord
	}

	r.Equal(getPosition(0), Coord{0, 0})
	setPosition(0, Coord{1, 2})
	r.Equal(getPosition(0), Coord{1, 2})
}
