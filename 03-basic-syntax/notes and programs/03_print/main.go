package main

import "fmt"
func main() {
	// This is the main entry point for the print package.
	// The actual functionality will be implemented in the print package.
	fmt.Println("Hello, World!");
	fmt.Println("This is the print package.");
	fmt.Println("You can use this package to print messages."); // all three line have new lines due to fmt.Println
	fmt.Println("Feel free to modify the code as needed.", "Add more functionality or change the messages."); // This line has two arguments, which will be printed with a space in between.
	data := "This is a variable holding some data.";
	fmt.Println(data); // This line prints the variable data.

	name:="John Doe";
	age:= 30;
	height:= 5.9;

	fmt.Println("Name:", name, "Age:", age, "Height:", height); // This line prints multiple variables in one line. with a space in between.
	fmt.Println("You can also use fmt.Printf for formatted output.");
	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height); // This line uses fmt.Printf for formatted output.
	fmt.Printf("Name: %s, type: %T, Age: %d, type: %T, Height: %.1f, type: %T,\n", name, name, age, age, height, height); // This line uses fmt.Printf for formatted output and it's types.
	fmt.Printf("Name: %d, Age: %d, Height: %.1f\n", name, age, height);  // This line will cause a compile error because name is a string, not an int.

	fmt.Println("This is the end of the main function.");
}