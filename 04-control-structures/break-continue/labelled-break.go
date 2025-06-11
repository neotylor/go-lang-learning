package main

import "fmt"

func main() {
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j == 4 {
				fmt.Println("Breaking out of both loops at i*j =", i*j)
				break outer
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}
