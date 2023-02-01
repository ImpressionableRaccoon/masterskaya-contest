package contest

import (
	"time"
)

type contextMutex struct {
	locked bool
}

func New() Mutex {
	return &contextMutex{}
}

func (mu *contextMutex) Lock() {
	for mu.locked {
		time.Sleep(time.Microsecond)
	}
	mu.locked = true
}

func (mu *contextMutex) LockChannel() <-chan struct{} {
	ch := make(chan struct{}, 1)
	if !mu.locked {
		mu.Lock()
		ch <- struct{}{}
	}
	return ch
}

func (mu *contextMutex) Unlock() {
	mu.locked = false
}
