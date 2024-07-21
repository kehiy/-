package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)
func 
main() {
	st := time.Now()
	count := 1_000_000
	coll := make(map[string]int, 0)
	for i := 0; i < count; i++ {
		id := uuid.New()
		fmt.Println(id.String())
		coll[id.String()] = 0
	}

	if len(coll) == count {
		fmt.Printf("Created %d unique ID in %v successfully!\n", count, time.Since(st))
		os.Exit(0)
	}
	
	fmt.Printf("Failed, target was %d unique IDs but we got collision, map size: %d\ntime: %v\n", 
		count, len(coll), time.Since(st))
}
