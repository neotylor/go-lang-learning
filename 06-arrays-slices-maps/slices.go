package main

import "fmt"

func main() {
	fmt.Println("Start Slice learning")
	// Slice is struct (pointer to array)

	numbers := []int{1, 2, 3, 4, 5}
	println("Numbers:", numbers)    //Numbers: [5/5]0xc00000c3c0
	println("Numbers:", &numbers)   //Numbers: 0xc000049e58
	println("Numbers:", numbers[1]) // Numbers: 2

	fmt.Println("Numbers:", numbers) //Numbers: [1 2 3 4 5]

	fmt.Printf("Numbers array (Slice) Type: %T\n", numbers) // Numbers array (Slice) Type: []int

	fmt.Println("Numbers's array lenght:", len(numbers)) //Numbers's array lenght: 5
	fmt.Println("Capacity:", cap(numbers))               // Capacity: 5

	// Add more element in array (slice). append method is not available in traditional array.
	numbers = append(numbers, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	println("Updated Numbers:", numbers) //[20/20]0xc000018140

	println("Updated Numbers:", &numbers) // Updated Numbers: 0xc000049e58

	println("Updated Numbers:", numbers[6]) // Numbers: 7

	fmt.Println("Updated Numbers:", numbers) //Numbers: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]

	fmt.Printf("Updated Numbers array (Slice) Type: %T\n", numbers) // Numbers array (Slice) Type: []int

	fmt.Println("Updated Numbers's array lenght:", len(numbers)) //Numbers's array lenght: 20

	fmt.Println("Capacity:", cap(numbers)) //Capacity: 20

	//Empty array
	name := []string{}

	fmt.Println("Name slice", name)         //Name slice []
	fmt.Println("Name lenght", len(name))   //Name lenght 0
	fmt.Println("Name capacity", cap(name)) //Name capacity 0

	// Create Slice using make method
	// Create array with custome lenght and capacity
	//Empty array
	newNumbers := make([]int, 3, 5)

	fmt.Println("New Numbers with make method slice", newNumbers)         //New Numbers with make method slice [0 0 0]
	fmt.Println("New Numbers with make method lenght", len(newNumbers))   //New Numbers with make method lenght 3
	fmt.Println("New Numbers with make method capacity", cap(newNumbers)) //New Numbers with make method capacity 5

	newNumbers = append(newNumbers, 1)

	fmt.Println("New Numbers with make method slice", newNumbers)         //New Numbers with make method slice [0 0 0 1]
	fmt.Println("New Numbers with make method lenght", len(newNumbers))   //New Numbers with make method lenght 4
	fmt.Println("New Numbers with make method capacity", cap(newNumbers)) //New Numbers with make method capacity 5

	newNumbers = append(newNumbers, 1, 2, 3, 4, 5)

	fmt.Println("New Numbers with make method slice", newNumbers)         //New Numbers with make method slice [0 0 0 1 1 2 3 4 5]
	fmt.Println("New Numbers with make method lenght", len(newNumbers))   //New Numbers with make method lenght 9
	fmt.Println("New Numbers with make method capacity", cap(newNumbers)) //New Numbers with make method capacity 10

	newNumbers = append(newNumbers, 1, 2)

	fmt.Println("New Numbers with make method slice", newNumbers)         //New Numbers with make method slice [0 0 0 1 1 2 3 4 5 1 2]
	fmt.Println("New Numbers with make method lenght", len(newNumbers))   //New Numbers with make method lenght 11
	fmt.Println("New Numbers with make method capacity", cap(newNumbers)) //New Numbers with make method capacity 20

	newNumbers = append(newNumbers, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("New Numbers with make method slice", newNumbers)         //New Numbers with make method slice [0 0 0 1 1 2 3 4 5 1 2 1 2 3 4 5 6 7 8 9 10]
	fmt.Println("New Numbers with make method lenght", len(newNumbers))   //New Numbers with make method lenght 21
	fmt.Println("New Numbers with make method capacity", cap(newNumbers)) //New Numbers with make method capacity 40

	//Error on create empty slice/Array
	// var emptySlice = []string

	// Can create empty slice using make method
	// var emptySlice = make([]string, 0)
	//OR short hand
	emptySlice := make([]string, 0)

	fmt.Println("Empty Slice", emptySlice)

}
