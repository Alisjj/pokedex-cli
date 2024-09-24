package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	a := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}

	go a.reapLoop(interval)

	return a
}

func (c *Cache) Add(key string, val []byte) {
	ent := cacheEntry{time.Now(), val}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = ent

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[key]

	if ok {
		return entry.val, ok
	}
	return entry.val, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entries {
			rem := time.Since(v.createdAt)
			if rem > interval {

				delete(c.entries, k)

			}
		}
		c.mu.Unlock()
	}

}
