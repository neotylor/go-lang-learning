package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "Hello from Goroutine!" // send data to channel
}

func main() {
	ch := make(chan string) // create an unbuffered channel

	go sendData(ch) // start goroutine

	message := <-ch // receive data from channel
	fmt.Println("Received:", message)
}

/*
 Output:
	Received: Hello from Goroutine!


 🔑 Key Concepts:
| Concept               | Description                                         |
| --------------------- | --------------------------------------------------- |
| `make(chan T)`        | Creates a new channel of type `T`                   |
| `ch <- val`           | Sends `val` into the channel `ch`                   |
| `<-ch`                | Receives a value from the channel `ch`              |
| Goroutines + channels | Enables safe communication between concurrent tasks |

*/
