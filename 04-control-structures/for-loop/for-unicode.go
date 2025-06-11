package main

import "fmt"

func main() {
	word := "Golang 💻"
	for i, r := range word {
		fmt.Printf("Index %d: Rune %q\n", i, r)
	}
}
