package pokecache

import (
	"sync"
	"time"
)

// NewCache()
func NewCache(interval time.Duration) Cache {
	var c = Cache{
		cacheEntry: make(map[string]cacheEntry),
		mutex:      &sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return c
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
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		c.mutex.Lock()
		for key, entry := range c.cacheEntry {
			if time.Now().After(entry.createdAt) {
				delete(c.cacheEntry, key)
			}
		}
		c.mutex.Unlock()
	}

}
