package main

import (
	"sync"
)

// this is a simple script that increases the value of x by 1 each time in the loop.
// for this 100 go routines are spawned up, but this causes racing conditions
// the binary file "./out" will produce non deterministic output due to racing conditions.

func main() {
	var wg sync.WaitGroup
	// increase value of x to 1 100 times
	var x = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			x++
		}()
	}
	wg.Wait()
	println(x)
}
