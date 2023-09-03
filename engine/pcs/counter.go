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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/concrete/api"
	"github.com/ethereum/go-ethereum/concrete/crypto"
	"github.com/ethereum/go-ethereum/concrete/lib"
	"github.com/ethereum/go-ethereum/concrete/precompiles"
	"github.com/ethereum/go-ethereum/concrete/utils"
)

//go:embed abi/Counter.json
var counterAbiJson []byte

var CounterABI abi.ABI
var CounterStorageKey = crypto.Keccak256([]byte("counter"))

func init() {
	var jsonAbi struct {
		ABI abi.ABI `json:"abi"`
	}
	err := json.Unmarshal(counterAbiJson, &jsonAbi)
	if err != nil {
		panic(err)
	}
	CounterABI = jsonAbi.ABI
}

type CounterPC struct {
	lib.BlankPrecompile
}

func (p *CounterPC) IsStatic(input []byte) bool {
	methodID, _ := utils.SplitInput(input)
	method, err := CounterABI.MethodById(methodID)
	if err != nil {
		return false
	}
	return method.IsConstant()
}

func (p *CounterPC) Run(env api.Environment, input []byte) ([]byte, error) {
	methodID, data := utils.SplitInput(input)
	method, err := CounterABI.MethodById(methodID)
	if err != nil {
		return nil, precompiles.ErrMethodNotFound
	}
	args, err := method.Inputs.Unpack(data)
	if err != nil {
		return nil, precompiles.ErrInvalidInput
	}

	counterValue := lib.NewDatastore(env).Value(CounterStorageKey)

	result, err := func() ([]interface{}, error) {
		switch method.Name {

		case "number":
			value := counterValue.Big()
			return []interface{}{value}, nil

		case "setNumber":
			newValue := args[0].(*big.Int)
			counterValue.SetBig(newValue)
			return nil, nil

		case "increment":
			currentValue := counterValue.Big()
			incValue := new(big.Int).Add(currentValue, big.NewInt(1))
			counterValue.SetBig(incValue)
			return nil, nil

		default:
			return nil, precompiles.ErrMethodNotFound
		}
	}()

	if err != nil {
		return nil, err
	}

	if len(method.Outputs) == 0 {

		return method.Outputs.Pack()
	}

	output, err := method.Outputs.Pack(result...)
	if err != nil {
		panic(err) // This will only panic if there is a bug in the precompile
	}
	return output, nil
}

var _ precompiles.Precompile = (*CounterPC)(nil)
