package tests

import (
	"encoding/hex"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/stretchr/testify/require"
	"testing"
)

type MsgProcessingSpecTest struct {
	Name           string
	Pre            *qbft.Instance
	PostRoot       string
	InputMessages  []*qbft.SignedMessage
	OutputMessages []*qbft.SignedMessage
	ExpectedError  string
}

func (test *MsgProcessingSpecTest) Run(t *testing.T) {
	var lastErr error
	for _, msg := range test.InputMessages {
		_, _, _, err := test.Pre.ProcessMsg(msg)
		if err != nil {
			lastErr = err
		}
	}

	if len(test.ExpectedError) != 0 {
		require.EqualError(t, lastErr, test.ExpectedError)
	} else {
		require.NoError(t, lastErr)
	}

	postRoot, err := test.Pre.State.GetRoot()
	require.NoError(t, err)

	// test output message
	if len(test.OutputMessages) > 0 {
		broadcastedMsgs := test.Pre.Config.GetNetwork().(*testingutils.TestingNetwork).BroadcastedMsgs
		require.Len(t, broadcastedMsgs, len(test.OutputMessages))

		for i, msg := range test.OutputMessages {
			r1, _ := msg.GetRoot()
			r2, _ := broadcastedMsgs[i].GetRoot()
			require.EqualValues(t, r1, r2)
		}
	}

	require.EqualValues(t, test.PostRoot, hex.EncodeToString(postRoot), "post root not valid")
}

func (test *MsgProcessingSpecTest) TestName() string {
	return test.Name
}