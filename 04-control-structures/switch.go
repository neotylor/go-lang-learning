package main

import "fmt"

func main() {
	day := 3
	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// default:
	// 	fmt.Println("Another day")
	// }

	switch day {
	case 1, 2, 3:
		fmt.Println("Start of the week")
	}

	char := 'a'

	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("It's a vowel")
	default:
		fmt.Println("It's a consonant")
	}

	score := 85

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}
}
