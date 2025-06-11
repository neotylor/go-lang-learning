package main

import "fmt"

func main() {
	println("Learn control structures (loop).")
	//While-like loop
	i := 0
	for i < 5 {
		fmt.Print(i)
		i++
	}
	//add new line
	fmt.Println()
	/*
		Notes:

			for is the only loop keyword in Go.

			Acts as while and do-while with different syntaxes.
	*/

}
