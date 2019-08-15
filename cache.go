package gocache

import (
	"errors"
	"sync"
)

//CacheItem holds the interface type to be cached
type CacheItem struct {
	Item interface{}
}

//Cache holds the cached items
type Cache struct {
	Name  string
	Items map[string]*CacheItem
	sync.RWMutex
}

//New creates a pointer to a Cache type
func New(name string) *Cache {
	c := &Cache{
		Name:  name,
		Items: make(map[string]*CacheItem),
	}

	return c
}

//AddOrUpdate adds an item to the cache if it doesnt exist, if it does exist, its value is updated
func (c *Cache) AddOrUpdate(key string, data interface{}) {
	c.Lock()
	defer c.Unlock()

	if i, ok := c.Items[key]; ok {
		i.Item = data
	}

	x := &CacheItem{
		Item: data,
	}

	c.Items[key] = x

}

//Fetch gets an item from the cache, if it doesnt exist an error is returned
func (c *Cache) Fetch(key string) (interface{}, error) {
	c.Lock()
	defer c.Unlock()

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
