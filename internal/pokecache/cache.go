package pokecache

import (
	"sync"
	"time"
)

// cacheEntry
type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	interval   time.Duration
	cacheEntry map[string]cacheEntry
	mutex      sync.RWMutex
}

// NewCache()
func NewCache(time time.Duration) *Cache {
	var c = Cache{
		interval:   time,
		cacheEntry: make(map[string]cacheEntry),
		mutex:      sync.RWMutex{},
	}
	c.reapLoop()
	return &c
}

// cache.Add()
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntry[key] = cacheEntry{time.Now(), val}
}

// cache.Get()
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, ok := c.cacheEntry[key]
	return entry.value, ok
}

// cache.reapLoop()
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	go func() {
		for ; ; <-ticker.C {
			c.mutex.Lock()
			for key, entry := range c.cacheEntry {
				if time.Now().After(entry.createdAt) {
					delete(c.cacheEntry, key)
				}
			}
			c.mutex.Unlock()
		}
	}()

}
