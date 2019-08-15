package main

import (
	gocache "GoCache"
	"fmt"
)

var (
	cache = gocache.New("test")
)

func main() {

	cache.AddOrUpdate("key", "value")
	v, _ := cache.Fetch("key")
	fmt.Println(v)
	cache.AddOrUpdate("key", "newvalue")
	s, _ := cache.Fetch("key")
	fmt.Println(s)
	cache.Remove("key")
}
