// SPDX-License-Identifier: MIT

pragma solidity 0.8.2;

contract OasysStateCommitmentChain {
    event StateBatchAppended(
        uint256 indexed _batchIndex,
        bytes32 _batchRoot,
        uint256 _batchSize,
        uint256 _prevTotalElements,
        bytes _extraData
    );

    event StateBatchDeleted(uint256 indexed _batchIndex, bytes32 _batchRoot);

    event StateBatchVerified(uint256 indexed _batchIndex, bytes32 _batchRoot);

    event StateBatchFailed(uint256 indexed _batchIndex, bytes32 _batchRoot);

    event OtherEvent(uint256 indexed _batchIndex);

    uint256 public nextIndex;

    function emitStateBatchAppended(
        uint256 batchIndex,
        bytes32 batchRoot,
        uint256 batchSize,
        uint256 prevTotalElements,
        bytes memory extraData
    ) external {
        emit StateBatchAppended(batchIndex, batchRoot, batchSize, prevTotalElements, extraData);
    }

    function emitStateBatchDeleted(uint256 batchIndex, bytes32 batchRoot) external {
        emit StateBatchDeleted(batchIndex, batchRoot);
    }

    function emitStateBatchVerified(uint256 batchIndex, bytes32 batchRoot) external {
        require(batchIndex == nextIndex, "Invalid batch index.");

        nextIndex++;

        emit StateBatchVerified(batchIndex, batchRoot);
    }

    function emitStateBatchFailed(uint256 batchIndex, bytes32 batchRoot) external {
        require(batchIndex >= nextIndex, "Invalid batch index.");

        emit StateBatchFailed(batchIndex, batchRoot);
    }

    function emitOtherEvent(uint256 batchIndex) external {
        emit OtherEvent(batchIndex);
    }
}
