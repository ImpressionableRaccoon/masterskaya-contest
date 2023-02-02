package contest

type contextMutex struct {
	ch chan struct{}
}

func New() Mutex {
	return &contextMutex{
		ch: make(chan struct{}, 1),
	}
}

func (mu *contextMutex) Lock() {
	mu.ch <- struct{}{}
}

func (mu *contextMutex) LockChannel() <-chan struct{} {
	ch := make(chan struct{}, 1)
	select {
	case mu.ch <- struct{}{}:
		ch <- struct{}{}
	default:
	}
	return ch
}

func (mu *contextMutex) Unlock() {
	<-mu.ch
}
