package sharedexemple

import (
	coreFunctions "shared-value/shared-exemple/core"
	"sync/atomic"
	"time"
)

type modelAtomic struct {
	currentState atomic.Bool
	state        iOperation
}

func WorkAtomic() *modelAtomic {
	v := &modelAtomic{}
	coreFunctions.StartTime(func() {
		time.Sleep(100 * time.Millisecond)
		tempState := &tempState{}
		v.currentState.Swap(tempState.enable())
	})
	return v
}

func (m *modelAtomic) GetStatus() bool {
	return m.currentState.Load()
}
