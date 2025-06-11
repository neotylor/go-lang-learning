package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // skip iteration
		}
		if i == 4 {
			break // exit loop
		}
		fmt.Println(i)
	}
}
