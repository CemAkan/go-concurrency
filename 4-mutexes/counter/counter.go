package counter

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// Config holds configuration for the ClickCounter.
type Config struct {
	CleanupInterval time.Duration
	MaxEntries      int
}

// LoadConfig loads configuration from environment variables with defaults.
func LoadConfig() Config {
	cleanupIntv := 5 * time.Minute
	maxEntries := 10000

	if val, ok := os.LookupEnv("CLICK_COUNTER_CLEANUP_INTERVAL"); ok {
		if d, err := time.ParseDuration(val); err == nil {
			cleanupIntv = d
		} else {
			log.Printf("Invalid CLICK_COUNTER_CLEANUP_INTERVAL: %v", err)
		}
	}

	if val, ok := os.LookupEnv("CLICK_COUNTER_MAX_ENTRIES"); ok {
		if i, err := strconv.Atoi(val); err == nil && i > 0 {
			maxEntries = i
		} else {
			log.Printf("Invalid CLICK_COUNTER_MAX_ENTRIES: %v", err)
		}
	}

	return Config{
		CleanupInterval: cleanupIntv,
		MaxEntries:      maxEntries,
	}
}

// ClickCounter is a thread-safe in-memory counter for URL short codes.
type ClickCounter struct {
	mu      sync.RWMutex
	counts  map[string]int64
	config  Config
	quit    chan struct{}
	running bool
}

// NewClickCounter creates and returns a configured ClickCounter instance.
func NewClickCounter(cfg Config) *ClickCounter {
	cc := &ClickCounter{
		counts:  make(map[string]int64),
		config:  cfg,
		quit:    make(chan struct{}),
		running: true,
	}

	// Start cleanup routine to remove old entries or limit map size
	go cc.cleanupLoop()

	return cc
}

// Increment safely increments the click count for the given short code.
func (c *ClickCounter) Increment(code string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[code]++
}

// Get safely returns the click count for the given short code.
func (c *ClickCounter) Get(code string) int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counts[code]
}

// Stop gracefully stops the cleanup goroutine.
func (c *ClickCounter) Stop() {
	if c.running {
		close(c.quit)
		c.running = false
	}
}

// cleanupLoop periodically cleans up the map to control memory usage.
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

// cleanUp removes entries to keep map size under MaxEntries.
func (c *ClickCounter) cleanUp() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.counts) <= c.config.MaxEntries {
		return
	}

	// Simple cleanup: remove random keys until under limit
	// (for demo purposes; customize as needed)
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
