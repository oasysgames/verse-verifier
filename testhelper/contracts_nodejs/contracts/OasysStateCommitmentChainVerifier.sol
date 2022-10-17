// SPDX-License-Identifier: MIT

pragma solidity ^0.8.2;

import { OasysStateCommitmentChain } from "./OasysStateCommitmentChain.sol";

contract OasysStateCommitmentChainVerifier {
    event StateBatchApproved(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot);

    event StateBatchRejected(address indexed stateCommitmentChain, uint256 indexed batchIndex, bytes32 batchRoot);

    struct ChainBatchHeader {
        uint256 batchIndex;
        bytes32 batchRoot;
        uint256 batchSize;
        uint256 prevTotalElements;
        bytes extraData;
    }

    struct AssertLog {
        address stateCommitmentChain;
        ChainBatchHeader batchHeader;
        bytes signatures;
        bool approve;
    }

    AssertLog[] public assertLogs;

    function approve(
        address stateCommitmentChain,
        ChainBatchHeader memory batchHeader,
        bytes[] memory signatures
    ) external {
        bytes memory _signatures;
        for (uint256 i = 0; i < signatures.length; i++) {
            _signatures = abi.encodePacked(_signatures, signatures[i]);
        }

        assertLogs.push(AssertLog(stateCommitmentChain, batchHeader, _signatures, true));

        OasysStateCommitmentChain(stateCommitmentChain).emitStateBatchVerified(
            batchHeader.batchIndex,
            batchHeader.batchRoot
        );

        emit StateBatchApproved(stateCommitmentChain, batchHeader.batchIndex, batchHeader.batchRoot);
    }

    function reject(
        address stateCommitmentChain,
        ChainBatchHeader memory batchHeader,
        bytes[] memory signatures
    ) external {
        bytes memory _signatures;
        for (uint256 i = 0; i < signatures.length; i++) {
            _signatures = abi.encodePacked(_signatures, signatures[i]);
        }

        assertLogs.push(AssertLog(stateCommitmentChain, batchHeader, _signatures, false));

        OasysStateCommitmentChain(stateCommitmentChain).emitStateBatchFailed(
            batchHeader.batchIndex,
            batchHeader.batchRoot
        );

        emit StateBatchRejected(stateCommitmentChain, batchHeader.batchIndex, batchHeader.batchRoot);
    }
}
