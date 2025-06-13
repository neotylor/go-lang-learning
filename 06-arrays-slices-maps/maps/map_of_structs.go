package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	people := map[string]Person{
		"emp1": {"Alice", 28},
		"emp2": {"Bob", 32},
	}

	for id, person := range people {
		fmt.Printf("ID: %s, Name: %s, Age: %d\n", id, person.Name, person.Age)
	}
}
