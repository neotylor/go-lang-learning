package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	println("Learn Web Request/services")
	res, err := http.Get("https://dummyjson.com/todos")

	if err != nil {
		fmt.Println("Error getting GET response", err)
		return
	}
	// resource free
	defer res.Body.Close()

	fmt.Printf("Type of response: %T\n", res)

	data, err_ := ioutil.ReadAll(res.Body)
	if err_ != nil {
		fmt.Println("Error getting while read response body", err_)
		return
	}

	fmt.Println("response", string(data))
}
