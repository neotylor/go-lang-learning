package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.Create("example.txt")

		if err != nil {
			fmt.Println("Error while creating files", err)
			return
		}

		// free resource after create successfully file
		//most required coding snippet when you engage any resource
		defer file.Close()

		println("Successfully created file.")

		content := "Magna laborum cillum magna velit. Dolor magna dolor occaecat non Lorem. Irure reprehenderit in tempor id duis esse. Dolor quis ex quis id occaecat in reprehenderit exercitation deserunt sit culpa consectetur officia. Est laboris aute exercitation irure reprehenderit amet occaecat enim dolore eiusmod velit. Nostrud in sunt laborum aliqua anim incididunt minim aliquip ut ex dolore sit dolor."

		contentByte, error := io.WriteString(file, content)

		if error != nil {
			println("Error while adding content on file", error)
			return
		}
		println("Content added successfully.", contentByte)

	*/

	/* file, err := os.Open("example.txt")

	if err != nil {
		println("Error while opening file", err) //Error while opening file (0xa29aa8,0xc0000241b0)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	//Read the file content into the buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file", err)
		}

		//Process the read content
		fmt.Println(string(buffer[:n]))
	} */

	// ioutil.ReadFile is deprecated: As of Go 1.16, this function simply calls [os.ReadFile].deprecateddefault
	// package ioutil ("io/ioutil")
	// content, err := ioutil.ReadFile("example.txt") // deprecated by ioutil
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}
	//content is a Array of slice so we need to use string()
	fmt.Println(string(content))
}
