// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface ICounter {
    function number() external view returns (uint256);
    function setNumber(uint256 x) external;
    function increment() external;
}
