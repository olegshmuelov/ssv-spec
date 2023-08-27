package testingutils

import (
	"time"

	"github.com/bloxapp/ssv-spec/qbft"
)

type TimerState struct {
	Timeouts int
	Round    qbft.Round
}

type TestQBFTTimer struct {
	State TimerState
}

func NewTestingTimer() qbft.Timer {
	return &TestQBFTTimer{
		State: TimerState{},
	}
}

func (t *TestQBFTTimer) TimeoutForRound(dutyStartTime time.Time, round qbft.Round) {
	t.State.Timeouts++
	t.State.Round = round
}
