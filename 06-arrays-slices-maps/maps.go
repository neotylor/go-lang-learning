package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a map with string keys and int values
	//Basic Map Creation and Access
	age := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// Access values
	fmt.Println("Alice's age:", age["Alice"])

	// Add a new key-value pair
	age["Charlie"] = 35

	// Delete a key
	delete(age, "Bob")

	// Check if a key exists
	val, ok := age["Bob"]
	if ok {
		fmt.Println("Bob's age:", val)
	} else {
		fmt.Println("Bob not found")
	}

	fmt.Println("Final map:", age)

	// 🔁 Iterating Over a Map
	fruits := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cherry",
	}

	for key, value := range fruits {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// 🧩 Map of Structs
	people := map[string]Person{
		"emp1": {"Alice", 28},
		"emp2": {"Bob", 32},
	}

	for id, person := range people {
		fmt.Printf("ID: %s, Name: %s, Age: %d\n", id, person.Name, person.Age)
	}

	//🧵 Map with Slice Values
	hobbies := map[string][]string{
		"Alice": {"Reading", "Cycling"},
		"Bob":   {"Gaming", "Cooking"},
	}

	hobbies["Charlie"] = []string{"Swimming"}

	for name, hobbyList := range hobbies {
		fmt.Printf("%s's hobbies: %v\n", name, hobbyList)
	}

	//🚨 Handling Map with Non-existent Keys

	score := make(map[string]int)

	// Non-existent key returns zero value
	fmt.Println("Initial score for Dave:", score["Dave"])

	// Use ok-idiom to check existence
	val, ok := score["Dave"]
	fmt.Printf("Exists? %v, Value: %d\n", ok, val)
}
