package main

import ( 
	"fmt"
	// "bufio"
	// "os"
	// "strings" // Uncomment if you need to manipulate strings later
)

func main() {
	// This is the main entry point for the application.
	// The actual implementation will be added later.
	// For now, we can keep this empty or add a simple print statement.
	println("This is the main function of the application. Implementation will follow.")
	fmt.Println("Welcome to the User Input Application!")	
	// fmt.Println("Please enter your input below:")
	// var userInput string
	// Wait for user input
	// fmt.Scanln(&userInput) // This will wait for user input from the console.
	// fmt.Println("You entered:", userInput)
	// Here we can process the user input as needed.
	// fmt.Println("Please enter any additional information or commands:")
	// var additionalInput string
	// fmt.Scan(&additionalInput) // Wait for additional input
	// fmt.Println("You entered:", additionalInput)
	// Simulate some processing
	// fmt.Println("Input processed successfully!")
	
	//difference between Scanln and Scan
	// Scanln reads a line of input until a newline character is encountered.
	// Scan reads input until the first space or newline character is encountered.
	
	// Using bufio to read a line of input from the user (Most Common)
	// reader := bufio.NewReader(os.Stdin)
	// // Read a line of input from the user
	// fmt.Println("Please enter a line of text:")
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// 	return
	// }
	// // Print the input received
	// fmt.Println("You entered:", input)

	// Using ReadLine (Low-Level, Less Common)
	// line, err := reader.ReadLine()
	// if err != nil {	
	// 	fmt.Println("Error reading line:", err)
	// 	return
	// }
	// // Print the line received
	// fmt.Println("You entered (ReadLine):", string(line))
	// // Using fmt.Scanf (Less Common, More Structured)
	fmt.Println("Please enter a number:")
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		fmt.Println("Error reading number:", err)
		return
	}
	// Print the number received
	fmt.Println("You entered the number:", number)

	// Final message before exiting
	fmt.Println("Thank you for your input!")
	fmt.Println("Exiting the application. Goodbye!")
}