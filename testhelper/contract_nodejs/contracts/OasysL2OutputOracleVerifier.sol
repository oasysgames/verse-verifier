// SPDX-License-Identifier: MIT

pragma solidity 0.8.2;

import { OasysL2OutputOracle } from "./OasysL2OutputOracle.sol";

contract OasysL2OutputOracleVerifier {
    event L2OutputApproved(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot);

    event L2OutputRejected(address indexed l2OutputOracle, uint256 indexed l2OutputIndex, bytes32 indexed outputRoot);

    struct OutputProposal {
        bytes32 outputRoot;
        uint128 timestamp;
        uint128 l2BlockNumber;
    }

    struct AssertLog {
        address l2OutputOracle;
        uint256 l2OutputIndex;
        OutputProposal l2Output;
        bytes signatures;
        bool approve;
    }

    struct Setting {
        bytes32 outputRoot;
        uint128 l2BlockNumber;
    }

    AssertLog[] public assertLogs;
    Setting public setting;

    function approve(
        address l2OutputOracle,
        uint256 l2OutputIndex,
        OutputProposal calldata l2Output,
        bytes[] calldata signatures
    ) external {
        assertLogs.push(
            AssertLog(
                l2OutputOracle,
                l2OutputIndex,
                l2Output,
                _joinSignatures(signatures),
                true
            )
        );

        OasysL2OutputOracle(l2OutputOracle).emitOutputVerified(
            l2OutputIndex,
            setting.outputRoot,
            setting.l2BlockNumber
        );

        emit L2OutputApproved(l2OutputOracle, l2OutputIndex, l2Output.outputRoot);
    }

    function reject(
        address l2OutputOracle,
        uint256 l2OutputIndex,
        OutputProposal calldata l2Output,
        bytes[] calldata signatures
    ) external {
        assertLogs.push(
            AssertLog(
                l2OutputOracle,
                l2OutputIndex,
                l2Output,
                _joinSignatures(signatures),
                false
            )
        );

        OasysL2OutputOracle(l2OutputOracle).emitOutputFailed(
            l2OutputIndex,
            setting.outputRoot,
            setting.l2BlockNumber
        );

        emit L2OutputRejected(l2OutputOracle, l2OutputIndex, l2Output.outputRoot);
    }

    function setL2ooSetting(Setting calldata _l2ooSetting) external {
        setting = _l2ooSetting;
    }

    function l2ooAssertLogsLen() external view returns (uint256) {
        return assertLogs.length;
    }

    function _joinSignatures(bytes[] calldata signatures) internal pure returns (bytes memory joined) {
        for (uint256 i = 0; i < signatures.length; i++) {
            joined = abi.encodePacked(joined, signatures[i]);
        }
    }
}
