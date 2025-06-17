package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	println("Learn JSON")

	person := Person{Name: "John", Age: 31, IsAdult: true}

	fmt.Println("Person Data is", person)

	// convert persion into json Encoding (Marshalling)

	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error while marshalling", err)
		return
	}

	fmt.Println("Person Original Json data is", jsonData)
	fmt.Println("Person Json data is", string(jsonData))

	var originalPerson Person

	erro := json.Unmarshal(jsonData, &originalPerson)
	if erro != nil {
		fmt.Println("Erron while unmarshal/Decording")
		return
	}
	fmt.Println("Persion Original data is", originalPerson)
}
