package main

func main() {
	println("Functions in Go")
	// Function calls
	printMessage("Hello, World!")
	sum, multiply := addWithMultiply(5, 10)
	println("Sum:", sum, "Multiplication:", multiply)
}

// funtion prototype 
	
/*
	// You can also define a function with multiple return values
	func addWithMultiply(a int, b int) (int, int) {
		return a + b, a * b
	}

*/
// Function to print a message
func printMessage(message string) {
	println(message)
}

func addWithMultiply(a int, b int) (int, int) {
	return a + b, a * b
}