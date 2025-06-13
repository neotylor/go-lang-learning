package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a new instance of the struct
	p1 := Person{Name: "Alice", Age: 30}

	// Access fields
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)

	// Modify fields
	p1.Age = 31
	fmt.Println("Updated Age:", p1.Age)
}

/*
### 🏁 Output:


Name: Alice
Age: 30
Updated Age: 31
*/
