package main

import "fmt"

func main() {
	age := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	fmt.Println("Alice's age:", age["Alice"])
	age["Charlie"] = 35
	delete(age, "Bob")

	val, ok := age["Bob"]
	if ok {
		fmt.Println("Bob's age:", val)
	} else {
		fmt.Println("Bob not found")
	}

	fmt.Println("Final map:", age)
}
