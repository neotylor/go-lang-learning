package main

import "fmt"

func main() {
	println("Learn control structures (Range loop).")
	//Range maps
	nums := []int{2, 4, 6}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	/*
		📝 Notes:
		You can ignore values using _, like for _, v := range nums.
	*/
}
