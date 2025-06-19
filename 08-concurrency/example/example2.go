package main

import "sync"

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() //Signal that this goroutine is done
	println("function id is", i)
	println("Function working start")
	// Doing some task
	// time.Sleep(2000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
	println("Function work end and id is", i)
	// wg.Done() //Signal that this goroutine is done
}

func main() {
	println("Learn Goroutines")

	var wg sync.WaitGroup
	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all worker to finish task
	println("Worker stask completed")
}
