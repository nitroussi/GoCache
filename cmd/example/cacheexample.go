package main

import (
	gocache "GoCache"
	"fmt"
	"time"
)

var (
	cache = gocache.New("test")
)

func main() {

	cache.AddOrUpdate("key", "value", 2)
	v, _ := cache.Fetch("key")
	fmt.Println(v)
	cache.AddOrUpdate("key", "newvalue", 2)
	s, _ := cache.Fetch("key")
	fmt.Println(s)
	time.Sleep(5 * time.Second)
	_, e := cache.Fetch("key")
	if e != nil {
		fmt.Println("cache item already removed from self cleanup")
	}
}
