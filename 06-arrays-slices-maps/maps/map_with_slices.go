package main

import "fmt"

func main() {
	hobbies := map[string][]string{
		"Alice": {"Reading", "Cycling"},
		"Bob":   {"Gaming", "Cooking"},
	}

	hobbies["Charlie"] = []string{"Swimming"}

	for name, hobbyList := range hobbies {
		fmt.Printf("%s's hobbies: %v\n", name, hobbyList)
	}
}
