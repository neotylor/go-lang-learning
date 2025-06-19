package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(message, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Call one function in a goroutine
	go printMessage("Goroutine")

	// Call another function normally
	printMessage("Main Function")

	// Wait for the goroutine to finish
	time.Sleep(2 * time.Second)
	fmt.Println("Main function ends")
}

/*
Output (approximate):

Main Function 1
Goroutine 1
Main Function 2
Goroutine 2
Main Function 3
Goroutine 3
Main function ends


🔑 Key Concepts:
go printMessage("Goroutine"): runs the function concurrently.

time.Sleep: allows time for the goroutine to complete (in real apps, use sync.WaitGroup).

Goroutines are lightweight — you can spawn thousands of them.
*/
