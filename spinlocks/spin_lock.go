package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type SpinLock struct {
	locked atomic.Int32
}

func NewSpinLock() sync.Locker {
	var lock SpinLock
	return &lock
}

func (l *SpinLock) Lock() {
	for !l.locked.CompareAndSwap(0, 1) {
		// Yield the processor so another runnable goroutine can execute.
		// This avoids monopolizing the CPU while repeatedly attempting
		// to acquire the lock.
		runtime.Gosched()
	}
}

func (l *SpinLock) Unlock() {
	l.locked.Store(0)
}
