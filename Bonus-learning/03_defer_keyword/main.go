package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	println("Starting of the program")
	// println("Middle of the program") // execute this line as possion of code
	result := add(5, 6)
	defer fmt.Println("Result is", result)
	defer println("Middle of the program") // execute this line after all code execute
	println("End of the program")
}

/*
If multiple defer line is in the program then it's go to the LIFO state (stack).

as per code

	println("Starting of the program")
	// println("Middle of the program") // execute this line as possion of code
	result := add(5, 6)
	defer fmt.Println("Result is", result)
	defer println("Middle of the program") // execute this line after all code execute
	println("End of the program")


then Output is

Starting of the program
End of the program
Middle of the program
Result is 11

*/
