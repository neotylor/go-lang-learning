package main

import "fmt"

func main() {
	println("Learn control structures (Range loop).")
	//Over maps
	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}

	/*
		📝 Notes:
		You can ignore values using _, like for _, v := range nums.
	*/
}
