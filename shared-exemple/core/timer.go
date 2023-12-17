package core

import (
	"sync"
	"time"
)

var once sync.Once
var t *time.Ticker

func StartTime(work func()) {

	once.Do(func() {
		t = time.NewTicker(500 * time.Millisecond)
	})

	go func() {
		for {
			work()
			<-t.C
		}
	}()
}
