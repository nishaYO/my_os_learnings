package main

import (
	"fmt"
	"sync"
	"time"
)

func performIOBoundTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Performing IO-bound task...")
	time.Sleep(4 * time.Second)
	fmt.Println("IO-bound task completed")
}

func performSimpleTask() {
	fmt.Println("Performing simple task...")
	fmt.Println("Simple task completed")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go performIOBoundTask(&wg)
	performSimpleTask()
	wg.Wait()
}

// go is a multithreaded language that supports concurrency.
// tasks can be perform in parallel using go routines and synchronized using sync package tools.
// go is good choice for CPU bound tasks because it can run multiple CPU bound tasks at the same time.
