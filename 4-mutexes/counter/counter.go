package counter

import "sync"

type ClickCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewClickCounter() *ClickCounter {
	return &ClickCounter{
		counts: make(map[string]int),
	}
}

// Increment increases click count for given short code safely
func (c *ClickCounter) Increment(code string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[code]++
}

// Get returns click count for given short code safely
func (c *ClickCounter) Get(code string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[code]
}
