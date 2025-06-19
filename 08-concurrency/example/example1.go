package main

import "time"

func sayHello() {
	println("Function sayHello start")
	println("Hello Function")
	time.Sleep(2000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
	println("Function sayHello end")
}
func sayHi() {
	println("Hi Codies :)")
	time.Sleep(1000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
}
func main() {
	println("Learn Goroutines")
	go sayHello()
	go sayHi()

	// wait for a moment to allow the goroutine to finish
	time.Sleep(1000 * time.Millisecond) // Simulating some work (wait and hold for 3 sec)
}

/*
example 1
func sayHello() {
	println("Function sayHello start")
	println("Hello Function")
	println("Function sayHello end")
}
func sayHi() {
	println("Hi Codies :)")
}
func main() {
	println("Learn Goroutines")
	sayHello()
	sayHi()
}

output:
Learn Goroutines
Function sayHello start
Hello Function
Function sayHello end
Hi Codies :)
===================================

Example 2
func sayHello() {
	println("Function sayHello start")
	println("Hello Function")
	time.Sleep(2000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
	println("Function sayHello end")
}
func sayHi() {
	println("Hi Codies :)")
}
func main() {
	println("Learn Goroutines")
	sayHello()
	sayHi()
}

Output:
Learn Goroutines
Function sayHello start
Hello Function 					// hould 2 sec then run next code
Function sayHello end
Hi Codies :)
===================================

Example 3
func sayHello() {
	println("Function sayHello start")
	println("Hello Function")
	time.Sleep(2000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
	println("Function sayHello end")
}
func sayHi() {
	println("Hi Codies :)")
}
func main() {
	println("Learn Goroutines")
	go sayHello() // switch to next task till then this task done. if main funtion finish before this function execution then it's funtion outout not visible
	sayHi()
}

Output:
Learn Goroutines
Hi Codies :)
===================================

Example 4
func sayHello() {
	println("Function sayHello start")
	println("Hello Function")
	time.Sleep(2000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
	println("Function sayHello end")
}
func sayHi() {
	println("Hi Codies :)")
}
func main() {
	println("Learn Goroutines")
	go sayHello()
	sayHi()

	// wait for a moment to allow the goroutine to finish
	time.Sleep(1000 * time.Millisecond) // Simulating some work (wait and hold for 1 sec)
}

Output:
Learn Goroutines
Hi Codies :)
Function sayHello start
Hello Function //this statement show due to main function wait 1 sec
===================================

Example 5
func sayHello() {
	println("Function sayHello start")
	println("Hello Function")
	time.Sleep(2000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
	println("Function sayHello end")
}
func sayHi() {
	println("Hi Codies :)")
}
func main() {
	println("Learn Goroutines")
	go sayHello()
	sayHi()

	// wait for a moment to allow the goroutine to finish
	time.Sleep(3000 * time.Millisecond) // Simulating some work (wait and hold for 3 sec)
}

Output:
Learn Goroutines
Hi Codies :)
Function sayHello start      	-|
Hello Function					-| this line execute becouse main funtion wait 3 sec and this function execute completely in with in 3 sec.
Function sayHello end			-|
===================================

Example 6
func sayHello() {
	println("Function sayHello start")
	println("Hello Function")
	time.Sleep(2000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
	println("Function sayHello end")
}
func sayHi() {
	println("Hi Codies :)")
}
func main() {
	println("Learn Goroutines")
	go sayHello()
	go sayHi()

	// wait for a moment to allow the goroutine to finish
	time.Sleep(1000 * time.Millisecond) // Simulating some work (wait and hold for 1 sec)
}

Outout:
Learn Goroutines
Hi Codies :)
Function sayHello start
Hello Function
===================================

Example 7
func sayHello() {
	println("Function sayHello start")
	println("Hello Function")
	time.Sleep(2000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
	println("Function sayHello end")
}
func sayHi() {
	println("Hi Codies :)")
	time.Sleep(1000 * time.Millisecond) // Simulating some work (wait and hold for 2 sec)
}
func main() {
	println("Learn Goroutines")
	go sayHello()
	go sayHi()

	// wait for a moment to allow the goroutine to finish
	time.Sleep(1000 * time.Millisecond) // Simulating some work (wait and hold for 3 sec)
}

Output
Learn Goroutines
Function sayHello start
Hello Function
Hi Codies :)
===================================

*/
