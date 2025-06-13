package main

import "fmt"

func main() {
	score := make(map[string]int)

	fmt.Println("Initial score for Dave:", score["Dave"])

	val, ok := score["Dave"]
	fmt.Printf("Exists? %v, Value: %d\n", ok, val)
}
