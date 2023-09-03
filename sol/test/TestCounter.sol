// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "../Counter.sol";

address constant COUNTER_ADDRESS = address(0x80);

contract CounterWrapper is Counter {
    function number() public view override returns (uint256) {
        (bool success, bytes memory data) = COUNTER_ADDRESS.staticcall(
            abi.encodeWithSignature("number()")
        );
        require(success);
        return abi.decode(data, (uint256));
    }

    function setNumber(uint256 newNumber) public override {
        (bool success, ) = COUNTER_ADDRESS.call(
            abi.encodeWithSignature("setNumber(uint256)", newNumber)
        );
        require(success);
    }

    function increment() public override {
        (bool success, ) = COUNTER_ADDRESS.call(
            abi.encodeWithSignature("increment()")
        );
        require(success);
    }
}

contract TestCounter {
    Counter internal counter;

    function setUp() public {
        counter = new CounterWrapper();
    }

    function testNumber() public view {
        require(counter.number() == 0);
    }

    function testSetNumber() public {
        counter.setNumber(10);
        require(counter.number() == 10);
    }

    function testIncrement() public {
        counter.increment();
        require(counter.number() == 1);
    }
}
