package main

import (
	"LRU-Cache-in-golang/lru"
	"fmt"
)

func main(){
	cacheSize := 3
	cache := lru.NewCache(cacheSize)
	fmt.Println("Get 1 : ", cache.GetFromCache(1))
	cache.Print()
	cache.AddToCache(2)
	cache.Print()
	cache.AddToCache(3)
	cache.Print()
	cache.AddToCache(4)
	cache.Print()
	fmt.Println("Get 2 : ", cache.GetFromCache(2))
	cache.Print()
	cache.AddToCache(5)
	cache.Print()
	fmt.Println("Get 2 : ", cache.GetFromCache(2))
	cache.Print()
}

