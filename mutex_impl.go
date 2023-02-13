package contest

type contestMutex struct {
	ch chan struct{}
}

func New() Mutex {
	mu := &contestMutex{
		ch: make(chan struct{}, 1),
	}
	mu.ch <- struct{}{}
	return mu
}

func (mu *contestMutex) Lock() {
	<-mu.ch
}

func (mu *contestMutex) LockChannel() <-chan struct{} {
	return mu.ch
}

func (mu *contestMutex) Unlock() {
	mu.ch <- struct{}{}
}
