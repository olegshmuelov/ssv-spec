package synccommittee

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/ssv"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// HappyFlow tests a full valcheck + post valcheck + duty sig reconstruction flow
func HappyFlow() *tests.MsgProcessingSpecTest {
	ks := testingutils.Testing4SharesSet()
	dr := testingutils.SyncCommitteeRunner(ks)

	msgs := []*types.SSVMessage{
		testingutils.SSVMsgSyncCommittee(testingutils.SignQBFTMsg(ks.Shares[1], 1, &qbft.Message{
			MsgType:    qbft.ProposalMsgType,
			Height:     qbft.FirstHeight,
			Round:      qbft.FirstRound,
			Identifier: testingutils.SyncCommitteeMsgID,
			Data:       testingutils.ProposalDataBytes(testingutils.TestSyncCommitteeConsensusDataByts, nil, nil),
		}), nil),
		testingutils.SSVMsgSyncCommittee(testingutils.SignQBFTMsg(ks.Shares[1], 1, &qbft.Message{
			MsgType:    qbft.PrepareMsgType,
			Height:     qbft.FirstHeight,
			Round:      qbft.FirstRound,
			Identifier: testingutils.SyncCommitteeMsgID,
			Data:       testingutils.PrepareDataBytes(testingutils.TestSyncCommitteeConsensusDataByts),
		}), nil),
		testingutils.SSVMsgSyncCommittee(testingutils.SignQBFTMsg(ks.Shares[2], 2, &qbft.Message{
			MsgType:    qbft.PrepareMsgType,
			Height:     qbft.FirstHeight,
			Round:      qbft.FirstRound,
			Identifier: testingutils.SyncCommitteeMsgID,
			Data:       testingutils.PrepareDataBytes(testingutils.TestSyncCommitteeConsensusDataByts),
		}), nil),
		testingutils.SSVMsgSyncCommittee(testingutils.SignQBFTMsg(ks.Shares[3], 3, &qbft.Message{
			MsgType:    qbft.PrepareMsgType,
			Height:     qbft.FirstHeight,
			Round:      qbft.FirstRound,
			Identifier: testingutils.SyncCommitteeMsgID,
			Data:       testingutils.PrepareDataBytes(testingutils.TestSyncCommitteeConsensusDataByts),
		}), nil),
		testingutils.SSVMsgSyncCommittee(testingutils.SignQBFTMsg(ks.Shares[1], 1, &qbft.Message{
			MsgType:    qbft.CommitMsgType,
			Height:     qbft.FirstHeight,
			Round:      qbft.FirstRound,
			Identifier: testingutils.SyncCommitteeMsgID,
			Data:       testingutils.CommitDataBytes(testingutils.TestSyncCommitteeConsensusDataByts),
		}), nil),
		testingutils.SSVMsgSyncCommittee(testingutils.SignQBFTMsg(ks.Shares[2], 2, &qbft.Message{
			MsgType:    qbft.CommitMsgType,
			Height:     qbft.FirstHeight,
			Round:      qbft.FirstRound,
			Identifier: testingutils.SyncCommitteeMsgID,
			Data:       testingutils.CommitDataBytes(testingutils.TestSyncCommitteeConsensusDataByts),
		}), nil),
		testingutils.SSVMsgSyncCommittee(testingutils.SignQBFTMsg(ks.Shares[3], 3, &qbft.Message{
			MsgType:    qbft.CommitMsgType,
			Height:     qbft.FirstHeight,
			Round:      qbft.FirstRound,
			Identifier: testingutils.SyncCommitteeMsgID,
			Data:       testingutils.CommitDataBytes(testingutils.TestSyncCommitteeConsensusDataByts),
		}), nil),

		testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[1], 1)),
		testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[2], 2)),
		testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[3], 3)),
	}

	return &tests.MsgProcessingSpecTest{
		Name:                    "sync committee happy flow",
		Runner:                  dr,
		Duty:                    testingutils.TestingSyncCommitteeDuty,
		Messages:                msgs,
		PostDutyRunnerStateRoot: "453b88b4df1fe24d5adaba650bde1a2bd8a10bcc97956085cebd4f1aee684f30",
		OutputMessages: []*ssv.SignedPartialSignatureMessage{
			testingutils.PostConsensusSyncCommitteeMsg(testingutils.Testing4SharesSet().Shares[1], 1),
		},
	}
}