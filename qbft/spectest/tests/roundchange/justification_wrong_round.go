package roundchange

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// JustificationWrongRound tests a single prepare justification with round != prepared round
func JustificationWrongRound() *tests.MsgProcessingSpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = 5
	ks := testingutils.Testing4SharesSet()

	prepareMsgs := []*qbft.SignedMessage{
		testingutils.TestingPrepareMessageWithRound(ks.Shares[1], types.OperatorID(1), 2),
		testingutils.TestingPrepareMessageWithRound(ks.Shares[2], types.OperatorID(2), 2),
		testingutils.TestingPrepareMessageWithRound(ks.Shares[3], types.OperatorID(3), 2),
	}
	msgs := []*qbft.SignedMessage{
		testingutils.TestingRoundChangeMessageWithRoundAndRC(ks.Shares[1], types.OperatorID(1), 5,
			testingutils.MarshalJustifications(prepareMsgs)),
	}

	return &tests.MsgProcessingSpecTest{
		Name:           "round change justification wrong round",
		Pre:            pre,
		PostRoot:       "00eeb94859eac77f900d7984d6f46fc0c1a769edaa36fffc34caebc08d99f478",
		InputMessages:  msgs,
		OutputMessages: []*qbft.SignedMessage{},
		ExpectedError:  "invalid signed message: round change justification invalid: wrong msg round",
	}
}
