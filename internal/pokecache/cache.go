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

// NewCache
func NewCache(time time.Duration) *Cache {
	return &Cache{
		interval:   time,
		cacheEntry: make(map[string]cacheEntry),
		mutex:      sync.RWMutex{},
	}
}

// cache.Add()
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntry[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, ok := c.cacheEntry[key]
	return entry.value, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(1 * time.Second)
	c.mutex.Lock()
	for range ticker.C {
		for key, entry := range c.cacheEntry {
			if time.Now().After(entry.createdAt.Add(c.interval)) {
				delete(c.cacheEntry, key)
			}
		}
	}
}
