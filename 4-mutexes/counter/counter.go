package counter

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"mutexExercise/internal/config"
)

type ClickCounter struct {
	mu              sync.RWMutex
	counts          map[string]int64
	quit            chan struct{}
	running         bool
	cleanupInterval time.Duration
	maxEntries      int
}

func NewClickCounter() *ClickCounter {
	cc := &ClickCounter{
		counts:  make(map[string]int64),
		quit:    make(chan struct{}),
		running: true,
	}

	cleanupSecStr := config.GetEnv("CLICK_COUNTER_CLEANUP_INTERVAL", "300")
	cleanupSec, _ := strconv.Atoi(cleanupSecStr)

	cc.cleanupInterval = time.Duration(cleanupSec) * time.Second

	maxEntriesStr := config.GetEnv("CLICK_COUNTER_MAX_ENTRIES", "10000")
	maxEntries, _ := strconv.Atoi(maxEntriesStr)
	cc.maxEntries = maxEntries

	// Modern rand seed
	src := rand.NewSource(time.Now().UnixNano())
	rand.New(src) // sadece seed için

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
	ticker := time.NewTicker(c.cleanupInterval)
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

	if len(c.counts) <= c.maxEntries {
		return
	}

	removed := 0
	for k := range c.counts {
		delete(c.counts, k)
		removed++
		if len(c.counts) <= c.maxEntries {
			break
		}
	}

	log.Printf("ClickCounter cleanup: removed %d entries, current size %d", removed, len(c.counts))
}
