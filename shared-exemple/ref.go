package sharedexemple

import (
	coreFunctions "shared-value/shared-exemple/core"
	"time"
)

type modelRef struct {
	currentState bool
	state        iOperation
}

func WorkRef() *modelRef {
	v := &modelRef{
		currentState: false,
	}
	coreFunctions.StartTime(func() {
		time.Sleep(100 * time.Millisecond)
		tempState := &tempState{}
		v.state = tempState
	})
	return v
}

func (m *modelRef) GetStatus() bool {
	if m.state != nil {
		m.currentState = m.state.enable()
	}
	return m.currentState
}
