// Copyright 2023 The concrete-geth Authors
//
// The concrete-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The concrete library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the concrete library. If not, see <http://www.gnu.org/licenses/>.

package pcs

import (
	_ "embed"
	"encoding/json"
	"math/big"

	datamod "github.com/concrete-eth/concrete-template/engine/pcs/codegen"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/concrete/api"
	"github.com/ethereum/go-ethereum/concrete/lib"
	"github.com/ethereum/go-ethereum/concrete/precompiles"
	"github.com/ethereum/go-ethereum/concrete/utils"
)

//go:embed abi/Position.json
var positionAbiJson []byte

var PositionABI abi.ABI

type Coord = struct {
	X int32 "json:\"x\""
	Y int32 "json:\"y\""
}

func init() {
	var jsonAbi struct {
		ABI abi.ABI `json:"abi"`
	}
	err := json.Unmarshal(positionAbiJson, &jsonAbi)
	if err != nil {
		panic(err)
	}
	PositionABI = jsonAbi.ABI
}

type PositionPrecompile struct {
	lib.BlankPrecompile
}

func (p *PositionPrecompile) IsStatic(input []byte) bool {
	methodID, _ := utils.SplitInput(input)
	method, err := PositionABI.MethodById(methodID)
	if err != nil {
		return false
	}
	return method.IsConstant()
}

func (p *PositionPrecompile) Run(env api.Environment, input []byte) ([]byte, error) {
	methodID, data := utils.SplitInput(input)
	method, err := PositionABI.MethodById(methodID)
	if err != nil {
		return nil, precompiles.ErrMethodNotFound
	}
	args, err := method.Inputs.Unpack(data)
	if err != nil {
		return nil, precompiles.ErrInvalidInput
	}

	datastore := lib.NewDatastore(env)
	positionMapping := datamod.NewPosition(datastore)

	result, err := func() ([]interface{}, error) {
		switch method.Name {

		case "setPosition":
			entity := args[0].(*big.Int)
			coord := args[1].(Coord)
			position := positionMapping.Get(entity)
			// position.Set(coord.X, coord.Y)
			position.Set(coord.X, coord.Y)
			return nil, nil

		case "getPosition":
			entity := args[0].(*big.Int)
			position := positionMapping.Get(entity)
			coord := Coord{position.GetX(), position.GetY()}
			return []interface{}{coord}, nil

		default:
			return nil, precompiles.ErrMethodNotFound
		}
	}()

	if err != nil {
		return nil, err
	}

	var output []byte
	if len(method.Outputs) == 0 {
		output, err = method.Outputs.Pack()
	} else {
		output, err = method.Outputs.Pack(result...)
	}
	if err != nil {
		panic(err) // This will only panic if there is a bug in the precompile
	}
	return output, nil
}

var _ precompiles.Precompile = (*PositionPrecompile)(nil)