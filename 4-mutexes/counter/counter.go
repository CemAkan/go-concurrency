package counter

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"mutexExercise/internal/config"
)

type ClickCounter struct {
	mu      sync.RWMutex
	counts  map[string]int64
	config  config.Config
	quit    chan struct{}
	running bool
}

func NewClickCounter(cfg config.Config) *ClickCounter {
	cc := &ClickCounter{
		counts:  make(map[string]int64),
		config:  cfg,
		quit:    make(chan struct{}),
		running: true,
	}

	// Modern rand usage with rand.NewSource
	src := rand.NewSource(time.Now().UnixNano())
	rand.New(src) // not used here, but seed is set without deprecated Seed()

	// Start cleanup routine to remove old entries or limit map size
	go cc.cleanupLoop()

	return cc
}

func (c *ClickCounter) Increment(code string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[code]++
}

func (c *ClickCounter) Get(code string) int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counts[code]
}

func (c *ClickCounter) Stop() {
	if c.running {
		close(c.quit)
		c.running = false
	}
}

func (c *ClickCounter) cleanupLoop() {
	ticker := time.NewTicker(c.config.CleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanUp()
		case <-c.quit:
			return
		}
	}
}

func (c *ClickCounter) cleanUp() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.counts) <= c.config.MaxEntries {
		return
	}

	removed := 0
	for k := range c.counts {
		delete(c.counts, k)
		removed++
		if len(c.counts) <= c.config.MaxEntries {
			break
		}
	}

	log.Printf("ClickCounter cleanup: removed %d entries, current size %d", removed, len(c.counts))
}
