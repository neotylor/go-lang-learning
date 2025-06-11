package main

import (
	"fmt" 
)
func main() {
	// Error handling is a critical aspect of programming in Go.
	// It allows developers to manage unexpected situations gracefully.

	// In Go, errors are treated as values. Functions that can fail typically return an error as the last return value.
	// The convention is to check if the error is nil to determine if the operation was successful.

	// Example of error handling:
	// func doSomething() (result string, err error) {
	//     // Perform some operation
	//     if failureCondition {
	//         return "", fmt.Errorf("an error occurred")
	//     }
	//     return "success", nil
	// }

	// Always handle errors appropriately to avoid panics and ensure program stability.
	// Use the "errors" package to create and handle errors.
	// Example:
	// answer, err := divide(10, 0)
	// if err != nil {
	// 	println("Error:", err.Error())
	// } else {
	// 	println("Result:", answer)
	// }
	
	// Example on err not use in function, then throws an error "declared and not used: err"
	// ans, err := divide(10, 0)
	// println("Result:", ans)


	// learn _ = err // This line is used to ignore the error, but it's not a good practice.
	ans, _ := divide(10, 0)
	println("Result:", ans)

	/*
		_ is a blank identifier in Go, used to ignore values that are not needed.
		It allows you to call a function and ignore its return value without causing a compile-time error.
		However, using the blank identifier for error handling is not recommended,
		as it can lead to unhandled errors and unexpected behavior in your program.

		error is a built-in interface type in Go that represents an error condition.
		It is commonly used in functions to indicate whether an operation was successful or if an error occurred.
		Using the blank identifier for error handling is not recommended,
		as it can lead to unhandled errors and unexpected behavior in your program.

		nil is a predeclared identifier in Go that represents a zero value for pointers, interfaces, maps, slices, channels, and function types.
		It is used to indicate the absence of a value or to signify that a variable has not been initialized.
		Using nil in error handling is common, as it indicates that no error occurred.
=	*/
}
/*
	explain the error handling in Go:
	Error handling in Go is done using the built-in `error` type, 
	which is an interface that represents an error condition. 
	Functions that can fail typically return an error as their last return value. 
	If the operation is successful, 
	the error will be `nil`. If there is an error, it will contain a descriptive message.

	prototype of error handling function:
	func functionName(parameters) (returnType, error) {
		// Perform some operation
		if someCondition {
			return value, fmt.Errorf("an error occurred: %s", details)
		}
		return value, nil
	}
*/
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

