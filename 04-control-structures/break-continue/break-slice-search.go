package main

import "fmt"

func main() {
	numbers := []int{3, 7, 9, 11, 15}
	target := 9

	for _, num := range numbers {
		if num == target {
			fmt.Println("Found target:", num)
			break
		}
	}
}
