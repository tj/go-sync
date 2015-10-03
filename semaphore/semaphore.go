package semaphore

type Semaphore chan struct{}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

func (s Semaphore) Wait() {
	for i := 0; i < cap(s); i++ {
		s <- struct{}{}
	}
}
