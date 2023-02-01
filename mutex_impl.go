package contest

import (
	"sync/atomic"
	"time"
)

type contextMutex struct {
	locked uint32
}

func New() Mutex {
	return &contextMutex{}
}

func (mu *contextMutex) Lock() {
	for atomic.LoadUint32(&mu.locked) != 0 {
		time.Sleep(time.Microsecond)
	}
	atomic.StoreUint32(&mu.locked, 1)
}

func (mu *contextMutex) LockChannel() <-chan struct{} {
	ch := make(chan struct{}, 1)
	if atomic.LoadUint32(&mu.locked) == 0 {
		mu.Lock()
		ch <- struct{}{}
	}
	return ch
}

func (mu *contextMutex) Unlock() {
	atomic.StoreUint32(&mu.locked, 0)
}
