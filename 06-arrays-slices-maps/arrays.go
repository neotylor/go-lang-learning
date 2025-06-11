package main

import (
	"fmt"
)

func main() {
	//Learn array in Go lang
	fmt.Println("Start learning Go lang")

	var fruitNames [5]string

	fruitNames[0] = "Apple"
	fruitNames[1] = "Banana"
	fruitNames[2] = "Mango"
	fruitNames[3] = "Strawberry"
	fruitNames[4] = "Grapes"

	fmt.Println("Fruits Name:", fruitNames)

	// create Arrays

	/*
		// Just Create array not initialized
			var numbers = [5]int
			or with shorthand:
			numbers := [5]int

		// Create array with initialized with creation
			var numbers = [5]int{1,2,3,4,5}
			or with shorthand:
			numbers := [5]int{1,2,3,4,5}

	*/
	// var numbers = [5]int{1, 2, 3, 4, 5}
	// OR
	// var numbers [5]int
	// numbers[0] = 1
	// numbers[1] = 2
	// numbers[2] = 3
	// numbers[3] = 4
	// numbers[4] = 5
	//OR
	//  Use Compiler to Infer Length (...)
	numbers := [...]int{1, 2, 3, 4, 5} // Array of length 5, inferred

	fmt.Println("Numbers:", numbers)

}
