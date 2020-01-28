package gocache

import (
	"errors"
	"sync"
	"time"
)

//CacheItem holds the interface type to be cached
type CacheItem struct {
	Item interface{}
	sync.RWMutex
	Lifetime   int
	LastUpdate time.Time
}

//Cache holds the cached items
type Cache struct {
	Name  string
	Items map[string]*CacheItem
	sync.RWMutex
}

//New creates a pointer to a Cache type
func New(name string) *Cache {
	return &Cache{
		Name:  name,
		Items: make(map[string]*CacheItem),
	}
}

//AddOrUpdate adds an item to the cache if it doesnt exist, if it does exist, its value is updated.
//Lifetime is a value in seconds, a lifetime of 0 means no cache.
func (c *Cache) AddOrUpdate(key string, data interface{}, lifetime int) interface{} {

	c.RLock()

	cacheItem, ok := c.Items[key]

	//update the value if it already exists
	if ok {
		existingCacheItem := cacheItem.Item
		cacheItem.Item = data
		cacheItem.LastUpdate = time.Now()
		c.RUnlock()
		return existingCacheItem
	}

	//we need to create a new cache item
	c.RUnlock()
	x := &CacheItem{
		Item:       data,
		LastUpdate: time.Now(),
	}
	c.Lock()
	c.Items[key] = x
	c.Unlock()

	if lifetime != 0 {
		go func() {
			time.Sleep(time.Duration(lifetime) * time.Second)
			c.Remove(key)
		}()
	}

	return nil

}

//Fetch gets an item from the cache, if it doesnt exist an error is returned
func (c *Cache) Fetch(key string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	if i, ok := c.Items[key]; ok {
		return i.Item, nil
	}

	return nil, errors.New("Cache item doesnt exist for key")
}

//Remove an item from the cache by key
func (c *Cache) Remove(key string) {
	c.Lock()
	defer c.Unlock()

	delete(c.Items, key)
}
