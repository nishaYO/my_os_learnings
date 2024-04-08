package main

import "sync"

// here we have a slice which is accessed by 5 goroutines and each of them is adding their id to the slice. 
// this is again causing racing conditions and hence non deterministic output.
// try running the executable multiple times to see it in action!

func main() {
	var wg sync.WaitGroup

	// Shared slice
	var slice []int

	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			slice = append(slice, id)
		}(i)
	}

	wg.Wait()

	println("Slice contents:", slice)
}
