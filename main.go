package main

import (
	"fmt"
	"github.com/vttrz/cache-lru/cache"
)

func main() {
	lru := cache.NewCacheLRU(5)
	lru.Set("a", 1)
	val := lru.Get("a")

	fmt.Printf("value: %v", val)
}
