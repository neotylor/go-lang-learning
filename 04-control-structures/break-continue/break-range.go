package main

import "fmt"

func main() {
	words := []string{"apple", "banana", "stop", "cherry"}

	for _, word := range words {
		if word == "stop" {
			fmt.Println("Stopping at:", word)
			break
		}
		fmt.Println("Word:", word)
	}
}
