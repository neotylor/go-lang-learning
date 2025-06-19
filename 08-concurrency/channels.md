Here’s a complete **Golang Channels Course**, organized into a structured tutorial format with **lessons, real examples, key concepts, and best practices**.

---

# 📘 Course: Mastering Channels in Go (Golang)

> Learn how to use **channels** to communicate safely across goroutines and build concurrent applications effectively.

---

## 📦 Module 1: Introduction to Channels

### 🔹 Lesson 1.1: What is a Channel?

* A channel is a communication pipe between goroutines.
* Go channels are **typed**: `chan int`, `chan string`, etc.

```go
var ch chan int // Declaration
ch = make(chan int) // Initialization
```

---

## 📥 Module 2: Sending & Receiving

### 🔹 Lesson 2.1: Basic Send and Receive

```go
func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}
```

### 🔹 Lesson 2.2: Direction of Data

* Send: `ch <- value`
* Receive: `val := <-ch`

---

## 🛠️ Module 3: Buffered Channels

### 🔹 Lesson 3.1: Create Buffered Channel

```go
func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

---

## 🚦 Module 4: Channel Synchronization

### 🔹 Lesson 4.1: Blocking Until Done

```go
func worker(ch chan bool) {
	fmt.Println("Working...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
	ch <- true
}

func main() {
	ch := make(chan bool)
	go worker(ch)
	<-ch // wait for signal
}
```

---

## 🚀 Module 5: Directional Channels

### 🔹 Lesson 5.1: Send-only and Receive-only

```go
func send(ch chan<- string) {
	ch <- "message"
}

func receive(ch <-chan string) {
	fmt.Println(<-ch)
}
```

---

## 🔄 Module 6: Range and Close

### 🔹 Lesson 6.1: Close Channel

```go
ch := make(chan int)
go func() {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}()

for v := range ch {
	fmt.Println(v)
}
```

---

## 🧠 Module 7: Select Statement

### 🔹 Lesson 7.1: Wait on Multiple Channels

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { ch1 <- "one" }()
go func() { ch2 <- "two" }()

select {
case msg1 := <-ch1:
	fmt.Println("Received", msg1)
case msg2 := <-ch2:
	fmt.Println("Received", msg2)
}
```

### 🔹 Lesson 7.2: Select with Timeout

```go
ch := make(chan string)

select {
case msg := <-ch:
	fmt.Println(msg)
case <-time.After(1 * time.Second):
	fmt.Println("Timeout!")
}
```

---

## 🔄 Module 8: Channel Iteration Patterns

### 🔹 Lesson 8.1: Fan-Out / Fan-In

```go
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}
```

---

## ⚠️ Module 9: Deadlocks and Pitfalls

### 🔹 Lesson 9.1: Common Mistake — Send on Full Buffered Channel

```go
ch := make(chan int, 1)
ch <- 1
ch <- 2 // fatal error: all goroutines are asleep - deadlock!
```

### 🔹 Lesson 9.2: Receive from Closed Channel

```go
ch := make(chan int, 1)
ch <- 1
close(ch)
fmt.Println(<-ch) // 1
fmt.Println(<-ch) // 0 (zero value)
```

---

## 🔐 Module 10: Best Practices

| Practice                         | Explanation                                          |
| -------------------------------- | ---------------------------------------------------- |
| Always close channels            | When sender is done and no further writes are needed |
| Never close from receiver        | Leads to runtime panic                               |
| Prefer range over manual receive | Clean iteration on closed channels                   |
| Use `select` for multiplexing    | Non-blocking channel operations                      |
| Avoid sharing channels unsafely  | Wrap in struct or function for safety                |

---

## ✅ Module 11: Real-World Patterns

### 🔹 Lesson 11.1: Ping-Pong

```go
func main() {
	ping := make(chan string)
	pong := make(chan string)

	go func() {
		msg := <-ping
		pong <- msg + " pong"
	}()

	ping <- "ping"
	fmt.Println(<-pong)
}
```

### 🔹 Lesson 11.2: Pipeline

```go
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	for n := range square(gen(2, 3, 4)) {
		fmt.Println(n)
	}
}
```

---

## 🧪 Module 12: Exercises

1. ✅ Build a channel-based logger.
2. ✅ Create a system where multiple goroutines generate data and one consumes it.
3. ✅ Implement a timer that notifies every 2 seconds via channel.
4. ✅ Write a pipeline of 3 stages using channels.
5. ✅ Simulate a chat room where users send/receive messages via channels.

---

## 🧵 Channel Type Summary

| Type              | Usage                     |
| ----------------- | ------------------------- |
| `chan T`          | Bidirectional             |
| `chan<- T`        | Send-only                 |
| `<-chan T`        | Receive-only              |
| `make(chan T)`    | Unbuffered channel (sync) |
| `make(chan T, n)` | Buffered channel (async)  |

---

## 📘 Summary Cheatsheet

| Concept           | Syntax / Functionality   |
| ----------------- | ------------------------ |
| Send              | `ch <- value`            |
| Receive           | `val := <-ch`            |
| Buffered          | `make(chan int, 2)`      |
| Close channel     | `close(ch)`              |
| Loop over channel | `for val := range ch {}` |
| Select statement  | `select { case ... }`    |
| Timeout select    | `time.After(duration)`   |
| Directional send  | `func(ch chan<- string)` |
| Directional recv  | `func(ch <-chan string)` |

---