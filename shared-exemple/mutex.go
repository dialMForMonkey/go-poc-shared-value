package sharedexemple

import (
	coreFunctions "shared-value/shared-exemple/core"
	"sync"
	"time"
)

type modelMutex struct {
	currentState bool
	mr           sync.RWMutex
	state        iOperation
}

func WorkMutex() *modelMutex {
	v := &modelMutex{
		currentState: false,
	}
	coreFunctions.StartTime(func() {
		time.Sleep(100 * time.Millisecond)
		tempState := &tempState{}
		v.mr.Lock()
		v.state = tempState
		v.mr.Unlock()
	})
	return v
}

func (m *modelMutex) GetStatus() bool {
	m.mr.RLock()
	if m.state != nil {
		m.currentState = m.state.enable()
	}
	m.mr.RUnlock()
	return m.currentState
}
