package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var rwMutex sync.RWMutex

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		// Lock for writing
		rwMutex.Lock()
		counter++
		fmt.Println("Counter Incremented:", counter)
		// Unlock for writing
		rwMutex.Unlock()

		time.Sleep(time.Millisecond * 100) // Simulate some work
	}
}

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		// Lock for reading
		rwMutex.RLock()
		fmt.Println("Counter Value (Read):", counter)
		// Unlock for reading
		rwMutex.RUnlock()

		time.Sleep(time.Millisecond * 50) // Simulate some work
	}
}

func main() {
	var wg sync.WaitGroup

	// Create two goroutines for writing
	wg.Add(2)
	go incrementCounter(&wg)
	go incrementCounter(&wg)

	// Create two goroutines for reading
	wg.Add(2)
	go readCounter(&wg)
	go readCounter(&wg)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final Counter Value:", counter)
}
