package main

import "fmt"

func main() {
	fruits := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cherry",
	}

	for key, value := range fruits {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
