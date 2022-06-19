package sync

import "sync"

type Counter struct {
	// mut sync.Mutex
	sync.Mutex // sync.Mutex embedded into the struct
	value      int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	// c.mut.Lock()
	// defer c.mut.Unlock()
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
