package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Id        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

func main() {
	println("Learn CRUN in GO lang")
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
