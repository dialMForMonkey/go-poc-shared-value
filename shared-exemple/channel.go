package sharedexemple

import (
	coreFunctions "shared-value/shared-exemple/core"
	"time"
)

type iOperation interface {
	enable() bool
}

type model struct {
	currentState bool
	state        chan iOperation
}

type tempState struct {
}

func (t *tempState) enable() bool {
	return coreFunctions.RandomResult()
}

func WorkChannel() *model {
	v := &model{
		currentState: false,
	}
	coreFunctions.StartTime(func() {
		time.Sleep(100 * time.Millisecond)
		tempState := &tempState{}
		v.state = make(chan iOperation)
		v.state <- tempState
	})
	return v
}

func (m *model) GetStatus() bool {
	if m.state != nil {
		value, ok := <-m.state
		if ok {
			m.currentState = value.enable()
			close(m.state)
		}
	}
	return m.currentState
}
