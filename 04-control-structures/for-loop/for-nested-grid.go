package main

import "fmt"

func main() {
	rows := 3
	cols := 4

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
}
