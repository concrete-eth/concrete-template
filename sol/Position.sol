// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

struct Coord {
    int32 x;
    int32 y;
}

interface Position {
    function setPosition(uint entity, Coord memory coord) external;

    function getPosition(uint entity) external view returns (Coord memory);
}
