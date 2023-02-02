package contest

type contestMutex struct {
	ch chan struct{}
}

func New() Mutex {
	return &contestMutex{
		ch: make(chan struct{}, 1),
	}
}

func (mu *contestMutex) Lock() {
	mu.ch <- struct{}{}
}

func (mu *contestMutex) LockChannel() <-chan struct{} {
	ch := make(chan struct{}, 1)
	select {
	case mu.ch <- struct{}{}:
		ch <- struct{}{}
	default:
	}
	return ch
}

func (mu *contestMutex) Unlock() {
	<-mu.ch
}
