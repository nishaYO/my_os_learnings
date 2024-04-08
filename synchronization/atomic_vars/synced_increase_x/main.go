package main

import (
	"sync"
	"sync/atomic"
)

// this solves the problem of racing condition in increase_x dir
// it uses atomic vars to ensure the adding operation is done all at once by each goroutine.
// this script produces deterministic output every time.

func main() {
	var wg sync.WaitGroup

	var x int32 // Use int32 for atomic operations
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&x, 1) // Atomic operation to increment x
		}()
	}
	wg.Wait()
	println(atomic.LoadInt32(&x)) // Atomic load to retrieve the value of x
}
