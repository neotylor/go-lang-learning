package main

func main() {
	println("Functions in Go")
	// Function calls
	printMessage("Hello, World!")
	sum := add(5, 10)
	println("Sum:", sum)

	// Using a function as a variable
	multiplyFunc := multiply
	result := multiplyFunc(3, 4)
	println("Multiplication Result:", result)
}

// funtion prototype 
/*
	func functionName(parameterType parameterName) returnType {
		// function body
		return value
	}
	// parameterType can be omitted if no parameters are needed
	// returnType can be omitted if no value is returned
*/
	
/*
	// You can define a function with no parameters and no return value
	func noParamsNoReturn() {
		println("This function has no parameters and no return value")
	}
	
	// You can define a function with no parameters but with a return value
	func noParamsWithReturn() string {
		return "This function has no parameters but returns a value"
	}

	// You can define a function with parameters but no return value
	func withParamsNoReturn(param1 string, param2 int) {
		println("This function has parameters but no return value")
		println("Parameter 1:", param1)
		println("Parameter 2:", param2)
	}

	// You can define a function with parameters and a return value
	func withParamsWithReturn(param1 string, param2 int) string {
		return "This function has parameters and returns a value: " + param1 + " " + string(param2)
	}

	// You can also define a function with multiple return values
	func multipleReturns(a int, b int) (int, int) {
		return a + b, a * b
	}

	// You can also define a function with named return values
	func namedReturns(a int, b int) (sum int, product int) {
		sum = a + b
		product = a * b
		return // named return values can be returned without specifying them again
	}

*/
// Function to print a message
func printMessage(message string) {
	println(message)
}

// Function to add two integers
/*
	// You can define parameters in a single line
	// func add(a, b int) int {
	func add(a, b int) int {
		return a + b
	}
*/ 
// Alternatively, you can define parameters separately
func add(a int, b int) int {
	return a + b
}

// Function to multiply two integers
func multiply(a int, b int) int {
	return a * b
}