## Tutorial: Adding a method to a precompile

### Introduction

This repository holds the code for a minimal concrete appchain with a stateful precompile with the following interface:

```solidity
// ################ ICounter.sol ################

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface ICounter {
    function number() external view returns (uint256);
    function setNumber(uint256 x) external;
    function increment() external;
}
```

The counter logic is implemented in `engine/main`. You'll notice it's written in Go instead of Solidity. Concrete precompiles run natively as part of the execution engine, not in the EVM, so they are written in languages suited for native execution.

Before we start, make sure you have installed the concrete CLI and you can successfully run the tests. See how to do this in the [Readme.md](README.md) "Getting started" section.

### Adding a `multiply(uint254 m)` method

We are going to add a `multiply` method to the precompile that will take a value, multiply it with the value held by the counter, and store the result back in the counter so it can be retrieved with `number()`.

First, add the method to the `ICounter.Sol:ICounter` interface so it looks like this:

```solidity
interface ICounter {
    function number() external view returns (uint256);
    function setNumber(uint256 x) external;
    function increment() external;
    function multiply(uint256 m) external; // add this
}
```

Now simply run `yarn`.

This will change two files:

- `sol/CounterPrecompile.sol` will now include a method to call `multiply`
- And `engine/ICounter.abi.json` will include the corresponding ABI

Once this is done, we need to implement the logic for the method in `engine/main.go`. You'll see the new ABI is already embedded in this file so all we have to do is add the multiplication logic in the switch statement in `Counter.Run`.

```go
switch method.Name {
case "number":
    return number.Bytes32().Bytes(), nil
case "setNumber":
    number.SetBytes32(common.BytesToHash(data))
    return nil, nil
case "increment":
    numberUint254 := number.Uint256()
    numberUint254.Add(numberUint254, new(uint256.Int).SetUint64(1))
    number.SetUint256(numberUint254)
    return nil, nil
case "multiply":
    // Read the input as a uint256
    multiplier := new(uint256.Int).SetBytes(data)
    // Load the current number
    numberUint254 := number.Uint256()
    // Multiply the number by the multiplier
    numberUint254.Mul(numberUint254, multiplier)
    // Store the new number
    number.SetUint256(numberUint254)
    // Return nothing
    return nil, nil
}
```

### Testing the new method

All we have left to do is make sure our new method actually works as expected.

Normal testing frameworks like foundry or hardhat don't work out of the box because the chains they run tests on don't include the precompile we have created.

We could test the precompile with Go, but then we would miss any issues that could arise in between the EVM and the precompile.

Fortunately, we can use concrete's custom framework for testing our precompile end-to-end using Solidity. You can see how this works under the hood by inspecting `engine/main_test.go` and the logic it imports from the concrete repository.

Now, go to `sol/TestCounter.sol` and add a test for our new method.

```solidity
function testMultiply() public {
    // Set an initial value for the counter
    CounterPrecompile.setNumber(2);
    // Call multiply with 3 as the multiplier
    CounterPrecompile.multiply(3);
    // Assert that number() returns the expected value
    require(CounterPrecompile.number() == 6, "number is not 6");
}
```

Note that, in order to be run as tests, method names must begin with "test". If, instead, we wanted to test that a method fails when it is expected to fail, we would prefix the name with "testFail".

Run the test with `yarn test`. You should see something like this:

```
=== RUN   TestE2E
    testtool.go:209:
        Running tests for sol/TestCounter.sol:CounterTest
=== RUN   TestE2E/testIncrement
    testtool.go:87: Gas used: 44685
=== RUN   TestE2E/testMultiply
    testtool.go:87: Gas used: 45599
=== RUN   TestE2E/testSetNumber
    testtool.go:87: Gas used: 44758
--- PASS: TestE2E (0.00s)
    --- PASS: TestE2E/testIncrement (0.00s)
    --- PASS: TestE2E/testMultiply (0.00s)
    --- PASS: TestE2E/testSetNumber (0.00s)
PASS
ok  	github.com/concrete-eth/concrete-template/engine	0.507s
```

Sweet!

We have added a multiply method to a concrete stateful precompile written in Go and tested it end-to-end with a Solidity test.

Checkout the [quickstart](https://github.com/concrete-eth/concrete-template/tree/quickstart) branch to see all the changes we made in this tutorial.

### Extra: What is CounterPrecompile.sol for?

You notice in the Solidity tests we call the precompile through a library instead of doing something like

```solidity
ICounter counter = ICounter(pcAddr);
counter.Number(42);
```

This is because solidity expects all addresses we call methods on through high-level functions to contain EVM bytecode and will preemptively revert if they don't. Precompiles, of course, don't have any EVM bytecode, since they don't run in the EVM! So the concrete CLI autogenerates a library that does low level calls to the precompile to circumvent this requirement.
