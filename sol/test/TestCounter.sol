// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "../ICounter.sol";
import "../codegen/CounterPrecompile.sol";

contract TestCounter {
    function setUp() public {}

    function testNumber() public view {
        require(CounterPrecompile.number() == 0);
    }

    function testSetNumber() public {
        CounterPrecompile.setNumber(10);
        require(CounterPrecompile.number() == 10);
    }

    function testIncrement() public {
        CounterPrecompile.increment();
        require(CounterPrecompile.number() == 1);
    }
}
