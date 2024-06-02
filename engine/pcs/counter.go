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
    "errors"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/concrete"
    "github.com/ethereum/go-ethereum/concrete/api"
    "github.com/ethereum/go-ethereum/concrete/lib"
    "github.com/ethereum/go-ethereum/concrete/utils"
)

//go:embed abi/Counter.json
var counterAbiJson []byte

var (
    ErrMethodNotFound = errors.New("method not found")
    ErrInvalidInput   = errors.New("invalid input")
)

var CounterABI abi.ABI

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

type CounterPrecompile struct {
    lib.BlankPrecompile
}

func (p *CounterPrecompile) IsStatic(input []byte) bool {
    methodID, _ := utils.SplitInput(input)
    method, err := CounterABI.MethodById(methodID)
    if err != nil {
        return false
    }
    return method.IsConstant()
}

func (p *CounterPrecompile) Run(env api.Environment, input []byte) ([]byte, error) {
    methodID, data := utils.SplitInput(input)
    method, err := CounterABI.MethodById(methodID)
    if err != nil {
        return nil, ErrMethodNotFound
    }
    args, err := method.Inputs.Unpack(data)
    if err != nil {
        return nil, ErrInvalidInput
    }

    datastore := lib.NewDatastore(env)
    counterMapping := lib.NewMapping(datastore, "Counter")

    result, err := func() ([]interface{}, error) {
        switch method.Name {
        case "getCount":
            count := counterMapping.Get(&big.Int{}).(*big.Int)
            return []interface{}{count}, nil

        case "increment":
            count := counterMapping.Get(&big.Int{}).(*big.Int)
            count.Add(count, big.NewInt(1))
            counterMapping.Set(&big.Int{}, count)
            env.EmitEvent("CounterIncremented", []interface{}{count})
            return nil, nil

        default:
            return nil, ErrMethodNotFound
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

var _ concrete.Precompile = (*CounterPrecompile)(nil)

