package main

import (
	sharedexemple "shared-value/shared-exemple"
	"time"
)

// / channel
func executeChannel() {

	r := sharedexemple.WorkChannel()
	for i := 0; i < 10000; i++ {
		print("case in channel::")
		println(r.GetStatus())
		time.Sleep(10 * time.Millisecond)
	}
}

// / ref
func executeRef() {
	r := sharedexemple.WorkRef()
	for i := 0; i < 10000; i++ {
		print("case in ref::")
		println(r.GetStatus())
		time.Sleep(10 * time.Millisecond)
	}
}

// / mutex read and write
func executeMutex() {
	r := sharedexemple.WorkMutex()
	for i := 0; i < 10000; i++ {
		print("case in ref::")
		println(r.GetStatus())
		time.Sleep(10 * time.Millisecond)
	}
}

/// atomic ref

func executeAtomic() {
	r := sharedexemple.WorkAtomic()
	for i := 0; i < 10000; i++ {
		print("case in ref::")
		println(r.GetStatus())
		time.Sleep(10 * time.Millisecond)
	}
}
func main() {
	executeAtomic()
}
