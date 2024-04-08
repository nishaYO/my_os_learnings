package main

import (
	"sync"
)

// this solves the problem of shared_slice dir by utilizing mutex
// in this script only one go routine can access the slice at a point of time.
// this prevents racing conditions and hence deterministic output.
// try running the executable multiple times to see it in action!

func main() {
	var wg sync.WaitGroup

	// Shared slice
	var slice []int

	var mu sync.Mutex

	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			slice = append(slice, id)
		}(i)
	}

	wg.Wait()
	println("Slice contents:", slice)
}
