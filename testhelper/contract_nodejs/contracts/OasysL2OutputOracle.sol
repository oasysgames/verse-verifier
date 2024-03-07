// SPDX-License-Identifier: MIT

pragma solidity 0.8.2;

contract OasysL2OutputOracle {
    event OutputProposed(
        bytes32 indexed outputRoot,
        uint256 indexed l2OutputIndex,
        uint256 indexed l2BlockNumber,
        uint256 l1Timestamp
    );

    event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex);

    event OutputVerified(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber);

    event OutputFailed(uint256 indexed l2OutputIndex, bytes32 indexed outputRoot, uint128 indexed l2BlockNumber);

    uint256 public nextVerifyIndex;

    function setNextVerifyIndex(uint256 val) external {
        nextVerifyIndex = val;
    }

    function emitOutputProposed(
        bytes32 outputRoot,
        uint256 l2OutputIndex,
        uint256 l2BlockNumber,
        uint256 l1Timestamp
    ) external {
        emit OutputProposed(outputRoot, l2OutputIndex, l2BlockNumber, l1Timestamp);
    }

    function emitOutputsDeleted(uint256 prevNextOutputIndex, uint256 newNextOutputIndex) external {
        emit OutputsDeleted(prevNextOutputIndex, newNextOutputIndex);
    }

    function emitOutputVerified(uint256 l2OutputIndex, bytes32 outputRoot, uint128 l2BlockNumber) external {
        require(l2OutputIndex == nextVerifyIndex, "L2OutputOracle: Invalid L2 output index");

        nextVerifyIndex++;

        emit OutputVerified(l2OutputIndex, outputRoot, l2BlockNumber);
    }

    function emitOutputFailed(uint256 l2OutputIndex, bytes32 outputRoot, uint128 l2BlockNumber) external {
        require(l2OutputIndex >= nextVerifyIndex, "L2OutputOracle: Invalid L2 output index");

        emit OutputFailed(l2OutputIndex, outputRoot, l2BlockNumber);
    }
}
