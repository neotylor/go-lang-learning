package main

import "fmt"

func main() {
	// Declare a variable
	x := 10
	fmt.Println("Value of x:", x)

	// Declare a pointer to x
	var p *int = &x

	// Print the address of x
	fmt.Println("Address of x:", p)

	// Print the value at the address (dereferencing)
	fmt.Println("Value pointed to by p:", *p)

	// Modify the value using the pointer
	*p = 20
	fmt.Println("New value of x after modification through pointer:", x)
}
