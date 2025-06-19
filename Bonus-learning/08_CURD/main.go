package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	Id        int    `json:"id,omitempty"` // ,omitempty is optional, depending on the API
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

func getRequest() {
	println("Learn Get method")
	todoRes, todoErr := http.Get("https://dummyjson.com/todos/1")
	if todoErr != nil {
		fmt.Println("Error while fatch todo list", todoErr)
		return
	}
	defer todoRes.Body.Close()

	if todoRes.StatusCode != http.StatusOK {
		fmt.Println("API getting error status is", todoRes.Status)
	}

	// todoList, todoDataErr := ioutil.ReadAll(todoRes.Body)

	// if todoDataErr != nil {
	// 	println("Error while Reading Response of todo list", todoDataErr)
	// 	return
	// }

	// fmt.Println("Todo List", string(todoList))

	//Other way to get data from api response
	var todoList Todo

	todoErr = json.NewDecoder(todoRes.Body).Decode(&todoList)
	if todoErr != nil {
		fmt.Println("Error while decode body", todoErr)
		return
	}
	fmt.Println("Todo List", todoList)
	fmt.Println("Todo title", todoList.Todo)
	fmt.Println("Todo Completed", todoList.Completed)
}

func postRequest() {
	println("Learn Post method")
	postUrl := "https://dummyjson.com/todos/add"
	todo := Todo{
		Todo:      "This is a new todo",
		Completed: false,
		UserId:    23,
	}

	// convert the Todo struct to JSON
	jsonData, err_ := json.Marshal(todo)
	if err_ != nil {
		fmt.Println("Error while marshal todo", err_)
		return
	}

	//convert json data to string
	jsonString := string(jsonData)

	// convert string data to response
	jsonReader := strings.NewReader(jsonString)

	res, err := http.Post(postUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending todo request", err)
		return
	}

	defer res.Body.Close()

	// convert response to readable data

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response Data", string(data))
	fmt.Println("Response status", res.Status)
}

func putRequest() {
	println("Learn Put method")
	putUrl := "https://dummyjson.com/todos/1"
	todo := Todo{
		Todo:      "This is a new todo",
		Completed: false,
		UserId:    23,
	}

	// convert the Todo struct to JSON
	jsonData, err_ := json.Marshal(todo)
	if err_ != nil {
		fmt.Println("Error while marshal todo", err_)
		return
	}

	//convert json data to string
	jsonString := string(jsonData)

	// convert string data to response
	jsonReader := strings.NewReader(jsonString)

	// Create put request
	req, err := http.NewRequest(http.MethodPut, putUrl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	//Send the Request
	client := http.Client{}

	httpRes, resErr := client.Do(req)

	if resErr != nil {
		fmt.Println("Error sending PUT request", resErr)
		return
	}

	defer httpRes.Body.Close()

	fmt.Println("Response status", httpRes.Status, req)

	// convert response to readable data
	data, _ := ioutil.ReadAll(httpRes.Body)
	fmt.Println("Response Data", string(data))
}

func deleteRequest() {
	println("Learn Delete method")
	deleteUrl := "https://dummyjson.com/todos/1"

	// Create put request
	req, err := http.NewRequest(http.MethodDelete, deleteUrl, nil)
	if err != nil {
		fmt.Println("Error creating Delete request", err)
		return
	}

	//Send the Request
	client := http.Client{}

	httpRes, resErr := client.Do(req)

	if resErr != nil {
		fmt.Println("Error sending Delete request", resErr)
		return
	}

	defer httpRes.Body.Close()

	fmt.Println("Response status", httpRes.Status, "\n", req)
}
func main() {
	println("Learn CRUN in GO lang")
	// getRequest()
	// postRequest()
	// putRequest()
	deleteRequest()
}
