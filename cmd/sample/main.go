package main

import (
	"fmt"

	"github.com/Tolstyak122/homeworks/pkg/cache/private"
	"github.com/Tolstyak122/homeworks/pkg/cache/simple"
)

func main() {
	cache1 := simple.GetCacheInstance()
	cache1.Set("key1", "val1")
	fmt.Println(cache1.Get("key12"))
	cache2 := private.GetCacheInstance()
	cache2.Set("key1", "val2")
	cache3 := private.GetCacheInstance()
	cache3.Set("key1", "val3")
	val1, _ := cache2.Get("key1")
	val2, _ := cache3.Get("key1")
	fmt.Println(val1, val2)
}
