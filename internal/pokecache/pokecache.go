package pokecache

import (
	"time"
	"sync"
)


type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Define the Cache struct with a map and mutex
type Cache struct {
	entries map[string]cacheEntry 
	mu      sync.Mutex            
}


func NewCache(interval time.Duration) *Cache {
    // Initialize empty map
    entries := make(map[string]cacheEntry)
    
    // Create new cache with map
    cache := &Cache{
        entries: entries,
        mu: sync.Mutex{},
    }
    
    go cache.reapLoop(interval)
    // How would you do this?
    
    return cache
}


//TO-DO : Create a cache.Add() method that adds a new entry to the cache. It should take a key (a string) and a val (a []byte).

//TO-DO : Create a cache.Get() method that gets an entry from the cache. It should take a key (a string) and return a []byte and a bool. The bool should be true if the entry was found and false if it wasn't.


func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    
    for {
        select {
        case <-ticker.C:
            c.mu.Lock()
            // For each entry in the cache
			for key, entry := range c.entries {
				// time.Since(entry.createdAt) tells us how much time has passed
				// since this entry was created
				timePassed := time.Since(entry.createdAt)
				
				// if more time has passed than our interval...
				if timePassed > interval {
					// delete this entry from the map using the key
					delete(c.entries, key)
				}
			}
            c.mu.Unlock()
        }
    }
}

