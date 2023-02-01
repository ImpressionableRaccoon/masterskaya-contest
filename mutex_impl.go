package contest

import (
	"sync/atomic"
	"time"
)

type contextMutex struct {
	locked atomic.Bool
}

func New() Mutex {
	return &contextMutex{}
}

func (mu *contextMutex) Lock() {
	for mu.locked.Load() {
		time.Sleep(time.Microsecond)
	}
	mu.locked.Swap(true)
}

func (mu *contextMutex) LockChannel() <-chan struct{} {
	ch := make(chan struct{}, 1)
	if !mu.locked.Load() {
		mu.Lock()
		ch <- struct{}{}
	}
	return ch
}

func (mu *contextMutex) Unlock() {
	mu.locked.Swap(false)
}
