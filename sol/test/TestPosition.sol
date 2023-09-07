// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "../Position.sol";
import "../codegen/PositionPrecompile.sol";

contract TestPosition {
    function setUp() public {}
    
    function testPosition() public {
        Coord memory coord;
        coord = PositionPrecompile.getPosition(0);
        require(coord.x == 0 && coord.y == 0, "Position should be (0, 0)");
        PositionPrecompile.setPosition(0, Coord(1, 2));
        coord = PositionPrecompile.getPosition(0);
        require(coord.x == 1 && coord.y == 2, "Position should be (1, 2)");
    }
}
