package main

import (
	sharedexemple "shared-value/shared-exemple"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	r := sharedexemple.WorkChannel()
	for i := 0; i < 10; i++ {
		r.GetStatus()
		time.Sleep(10 * time.Millisecond)
	}
}

func BenchmarkChannel(b *testing.B) {
	r := sharedexemple.WorkChannel()
	for i := 0; i < b.N; i++ {
		r.GetStatus()
	}
}
func TestRef(t *testing.T) {
	r := sharedexemple.WorkRef()
	for i := 0; i < 10; i++ {
		r.GetStatus()
		time.Sleep(10 * time.Millisecond)
	}
}

func BenchmarkRef(b *testing.B) {
	r := sharedexemple.WorkRef()
	for i := 0; i < b.N; i++ {
		r.GetStatus()
	}
}

func TestMutex(t *testing.T) {
	r := sharedexemple.WorkMutex()
	for i := 0; i < 10; i++ {
		r.GetStatus()
		time.Sleep(10 * time.Millisecond)
	}
}

func BenchmarkMutex(b *testing.B) {
	r := sharedexemple.WorkMutex()
	for i := 0; i < b.N; i++ {
		r.GetStatus()
	}
}

func TestAtomic(t *testing.T) {
	r := sharedexemple.WorkAtomic()
	for i := 0; i < 10; i++ {
		r.GetStatus()
		time.Sleep(10 * time.Millisecond)
	}
}

func BenchmarkAtomic(b *testing.B) {
	r := sharedexemple.WorkAtomic()
	for i := 0; i < b.N; i++ {
		r.GetStatus()
	}
}
