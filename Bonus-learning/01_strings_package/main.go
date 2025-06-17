package main

import (
	"fmt"
	"strings"
)

func main() {
	println("Learn Strings Package")

	// spite string to array(slice)
	fruit_names := "apple,banana,orange,mango"
	fruit_slice := strings.Split(fruit_names, ",")

	fmt.Println("Original String", fruit_names)
	fmt.Println("Split by ,", fruit_slice)

	//count string from string

	str_1 := "one two three four two two five tow"

	wordCount := strings.Count(str_1, "two")

	fmt.Println("Original string is", str_1, "count of word 'two' is", wordCount)

	//Remove white space from start or end

	str_2 := "                 Hello, Go!       "

	trimmedString := strings.TrimSpace(str_2)
	fmt.Printf("Original string is '%s'\nTrimmed string '%s'\n", str_2, trimmedString)

	//Join Strings

	str_3 := "Neo"
	str_4 := "Tylor"

	jointString := strings.Join([]string{str_3, str_4}, " ")

	fmt.Println("First string is", str_3)
	fmt.Println("Second string is", str_4)
	fmt.Println("Joint string is", jointString)
}
