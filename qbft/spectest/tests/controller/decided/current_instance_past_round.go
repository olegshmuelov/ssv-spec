package decided

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/herumi/bls-eth-go-binary/bls"
)

// CurrentInstancePastRound tests a decided msg received for current running instance for a past round
func CurrentInstancePastRound() *tests.ControllerSpecTest {
	ks := testingutils.Testing4SharesSet()

	rcMsgs := []*qbft.SignedMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[1], 1, 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[2], 2, 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[3], 3, 2),
	}

	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessage(ks.Shares[1], 1),

		testingutils.TestingPrepareMessage(ks.Shares[1], 1),
		testingutils.TestingPrepareMessage(ks.Shares[2], 2),
	}
	msgs = append(msgs, rcMsgs...)
	msgs = append(msgs, []*qbft.SignedMessage{
		testingutils.TestingProposalMessageWithRound(ks.Shares[1], 1, 2),

		testingutils.TestingPrepareMessageWithRound(ks.Shares[1], 1, 2),
		testingutils.TestingPrepareMessageWithRound(ks.Shares[2], 2, 2),
		testingutils.TestingPrepareMessageWithRound(ks.Shares[3], 3, 2),

		testingutils.TestingCommitMessageWithRound(ks.Shares[1], 1, 2),
		testingutils.TestingCommitMessageWithRound(ks.Shares[2], 2, 2),
		testingutils.TestingCommitMultiSignerMessage([]*bls.SecretKey{ks.Shares[1], ks.Shares[2], ks.Shares[3]}, []types.OperatorID{1, 2, 3}),
	}...)

	return &tests.ControllerSpecTest{
		Name: "decide current instance past round",
		RunInstanceData: []*tests.RunInstanceData{
			{
				InputValue:    []byte{1, 2, 3, 4},
				InputMessages: msgs,
				ExpectedDecidedState: tests.DecidedState{
					DecidedCnt: 1,
					DecidedVal: []byte{1, 2, 3, 4},
				},
				ControllerPostRoot: "03bd86e8a9e695b939266da2d0f02421d2fa8574bad5417fe318d5603dd54017",
			},
		},
	}
}
