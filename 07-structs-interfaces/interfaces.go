package main

import "fmt"

// Define an interface
type Animal interface {
	Speak() string
}

// Define a type that implements the interface
type Dog struct{}

// Implement the Speak method for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Another type
type Cat struct{}

// Implement the Speak method for Cat
func (c Cat) Speak() string {
	return "Meow!"
}

// Function that accepts the interface
func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	var dog Animal = Dog{}
	var cat Animal = Cat{}

	MakeAnimalSpeak(dog)
	MakeAnimalSpeak(cat)
}

/*
### 🏁 Output:


Woof!
Meow!
*/
