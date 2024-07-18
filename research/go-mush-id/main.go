package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	mid := NewMushID(DefaultSeedFn, DefaultHashFn)

	st := time.Now()
	count := 10
	coll := make(map[string]int, 0)
	for i := 0; i < count; i++ {
		id := mid.ID()
		fmt.Println(id)
		coll[id] = 0
	}

	if len(coll) == count {
		fmt.Printf("Created %d unique ID in %v successfully!\n", count, time.Since(st))
		os.Exit(0)
	}
	
	fmt.Printf("Failed, target was %d unique IDs but we got collision, map size: %d\ntime: %v\n", 
		count, len(coll), time.Since(st))
}
