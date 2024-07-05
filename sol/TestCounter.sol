// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./CounterPrecompile.sol";

contract CounterTest {
    function setUp() public {}

    function testSetNumber() public {
        CounterPrecompile.setNumber(42);
        require(CounterPrecompile.number() == 42, "number is not 42");
    }

    function testIncrement() public {
        CounterPrecompile.increment();
        require(CounterPrecompile.number() == 1, "number is not 1");
    }
}
