// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "../counter_contract.sol";
import "../codegen/PositionPrecompile.sol";

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";

contract TestCounter {
    Counter counter = Counter(DeployedAddresses.Counter());

    function testInitialCount() public {
        uint256 expected = 0;
        Assert.equal(counter.getCount(), expected, "Initial count should be 0");
    }

    function testIncrement() public {
        counter.increment();
        uint256 expected = 1;
        Assert.equal(counter.getCount(), expected, "Count should be 1 after increment");
    }
}

