package sync

import "sync"

// NewCounter  NewCounter
func NewCounter() *Counter {
	return &Counter{}
}

//Counter ... does Counter very important, not random something
type Counter struct {
	mu    sync.Mutex
	value int
}

//Inc ...
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

//Value ...
func (c *Counter) Value() int {
	return c.value
}
