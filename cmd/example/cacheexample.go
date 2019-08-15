package main

import (
	gocache "GoCache"
	"fmt"
	"time"
)

var (
	cache = gocache.New("test", 3)
)

func main() {

	cache.AddOrUpdate("key", "value")
	v, _ := cache.Fetch("key")
	fmt.Println(v)
	cache.AddOrUpdate("key", "newvalue")
	s, _ := cache.Fetch("key")
	fmt.Println(s)
	time.Sleep(5 * time.Second)
	_, e := cache.Fetch("key")
	if e != nil {
		fmt.Println("cache item already removed from self cleanup")
	}
}
