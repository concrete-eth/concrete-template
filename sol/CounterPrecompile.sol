// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/* Autogenerated file. Do not edit manually. */

import "./ICounter.sol";

library CounterPrecompile {
    address constant precompileAddress = address(0x0000000000000000000000000000000000000080);

    function number() internal view returns (uint256) {
        (bool success, bytes memory data) = precompileAddress.staticcall(
            abi.encodeWithSignature("number()")
        );
        require(success);
        return abi.decode(data, (uint256));
    }

    function setNumber(uint256 x) internal {
        (bool success, ) = precompileAddress.call(
            abi.encodeWithSignature("setNumber(uint256)", x)
        );
        require(success);
    }

    function increment() internal {
        (bool success, ) = precompileAddress.call(
            abi.encodeWithSignature("increment()")
        );
        require(success);
    }
}
