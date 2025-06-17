package main

import (
	"fmt"
	"net/url"
)

func main() {
	println("Learn URL handling")
	rawURL := "https://dummyjson.com/todos?limit=3&skip=10"
	// rawURL := "neotylor"

	fmt.Printf("Type of URL: %T\n", rawURL)

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error while parse URL", err)
		return
	}
	fmt.Printf("Type of URL: %T\n", parsedURL)

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

	//Modifying URL

	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "page=50"

	// Constracting a URL string from a URL object
	newURL := parsedURL.String()

	fmt.Println("New URL is", newURL)
}
