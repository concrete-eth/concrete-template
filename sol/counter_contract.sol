// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Counter {
    uint256 private count;

    // Event to emit when the counter is incremented
    event CounterIncremented(uint256 newCount);

    // Function to get the current count
    function getCount() public view returns (uint256) {
        return count;
    }

    // Function to increment the count
    function increment() public {
        count += 1;
        emit CounterIncremented(count);
    }
}
