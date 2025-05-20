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

	// Cleanup interval
	cleanupSecStr := config.GetEnv("CLICK_COUNTER_CLEANUP_INTERVAL", "300")
	cleanupSec, err := strconv.Atoi(cleanupSecStr)
	if err != nil || cleanupSec <= 0 {
		log.Printf("Invalid CLICK_COUNTER_CLEANUP_INTERVAL value '%s', using default 300 seconds", cleanupSecStr)
		cleanupSec = 300
	}
	cc.cleanupInterval = time.Duration(cleanupSec) * time.Second

	// Max entries
	maxEntriesStr := config.GetEnv("CLICK_COUNTER_MAX_ENTRIES", "10000")
	maxEntries, err := strconv.Atoi(maxEntriesStr)
	if err != nil || maxEntries <= 0 {
		log.Printf("Invalid CLICK_COUNTER_MAX_ENTRIES value '%s', using default 10000", maxEntriesStr)
		maxEntries = 10000
	}
	cc.maxEntries = maxEntries

	// Seed random
	src := rand.NewSource(time.Now().UnixNano())
	rand.New(src) // sadece seed için kullanılıyor

	// Başlatıcı goroutine: düzenli cleanup döngüsü
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
