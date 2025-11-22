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
	cacheEntry map[string]cacheEntry
	mutex      *sync.RWMutex
}
