Here! Below is a **comprehensive course/tutorial on Concurrency in Go (Golang)**, structured as a step-by-step **modular guide** with **explanations, code examples, use cases**, and **best practices**.

---

# 📘 Course: Mastering Concurrency in Go (Golang)

> Learn everything about concurrency in Go: goroutines, channels, sync primitives, select statements, worker pools, race conditions, and more.

---

## 📦 Module 1: Introduction to Concurrency

### 🔹 Lesson 1.1: What is Concurrency?

* Concurrency is the ability of a program to do many things at once.
* In Go, this is achieved through:

  * **Goroutines** (lightweight threads)
  * **Channels** (safe communication)
  * **sync primitives**

---

## 🧵 Module 2: Goroutines

### 🔹 Lesson 2.1: Basic Goroutine

```go
func printMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	go printMessage("Hello from Goroutine")
	time.Sleep(1 * time.Second) // Wait for goroutine to finish
}
```

### 🔹 Lesson 2.2: Multiple Goroutines

```go
func main() {
	for i := 1; i <= 5; i++ {
		go fmt.Println("Goroutine:", i)
	}
	time.Sleep(1 * time.Second)
}
```

---

## 📬 Module 3: Channels

### 🔹 Lesson 3.1: Basic Channel Communication

```go
func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello"
	}()
	msg := <-ch
	fmt.Println("Received:", msg)
}
```

### 🔹 Lesson 3.2: Buffered Channels

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
fmt.Println(<-ch)
fmt.Println(<-ch)
```

---

## 🔁 Module 4: Directional Channels

### 🔹 Lesson 4.1: Send-only and Receive-only

```go
func send(ch chan<- string) {
	ch <- "sent"
}

func receive(ch <-chan string) {
	fmt.Println(<-ch)
}
```

---

## 🎛️ Module 5: Select Statement

### 🔹 Lesson 5.1: Using `select` with channels

```go
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "One" }()
	go func() { ch2 <- "Two" }()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
}
```

---

## 🧰 Module 6: WaitGroups

### 🔹 Lesson 6.1: Synchronizing Goroutines

```go
var wg sync.WaitGroup

func worker(id int) {
	defer wg.Done()
	fmt.Println("Worker", id)
}

func main() {
	wg.Add(2)
	go worker(1)
	go worker(2)
	wg.Wait()
}
```

---

## 🛑 Module 7: Mutexes

### 🔹 Lesson 7.1: Prevent Race Conditions

```go
var (
	counter int
	mu      sync.Mutex
)

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Counter:", counter)
}
```

---

## 📊 Module 8: Race Conditions

### 🔹 Lesson 8.1: Demonstrate Race

```go
var count int

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			count++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Count:", count)
}
```

### 🔹 Lesson 8.2: Detect Race (with `go run -race`)

```bash
go run -race main.go
```

---

## 🛠️ Module 9: Worker Pool

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
		fmt.Println("Result:", <-results)
	}
}
```

---

## ⏰ Module 10: Timeout and Ticker

### 🔹 Lesson 10.1: Timeout with `select`

```go
ch := make(chan string)

go func() {
	time.Sleep(2 * time.Second)
	ch <- "done"
}()

select {
case msg := <-ch:
	fmt.Println("Received:", msg)
case <-time.After(1 * time.Second):
	fmt.Println("Timeout!")
}
```

### 🔹 Lesson 10.2: Periodic Execution with `Ticker`

```go
ticker := time.NewTicker(500 * time.Millisecond)
go func() {
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}()
time.Sleep(2 * time.Second)
ticker.Stop()
```

---

## 🧠 Module 11: Channel Closing and Ranging

### 🔹 Lesson 11.1: Close Channel & Range

```go
ch := make(chan int)
go func() {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}()

for val := range ch {
	fmt.Println(val)
}
```

---

## 🔐 Module 12: Once, Cond, and Atomic

### 🔹 sync.Once

```go
var once sync.Once

func initOnlyOnce() {
	fmt.Println("Initialized!")
}

func main() {
	for i := 0; i < 5; i++ {
		go once.Do(initOnlyOnce)
	}
	time.Sleep(1 * time.Second)
}
```

### 🔹 Atomic Counter

```go
var counter int32

func main() {
	for i := 0; i < 1000; i++ {
		go atomic.AddInt32(&counter, 1)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Counter:", counter)
}
```

---

## 🧪 Module 13: Real-World Exercises

1. ✅ Implement a producer-consumer pattern with goroutines.
2. ✅ Build a fixed-size worker pool that handles jobs concurrently.
3. ✅ Use `context.WithTimeout` to cancel goroutines after a limit.
4. ✅ Write a rate limiter using `time.Ticker`.
5. ✅ Safely increment counters from 1000 goroutines using atomic or mutex.

---

## 📘 Summary Cheatsheet

| Feature       | Tool               |
| ------------- | ------------------ |
| Goroutine     | `go func()`        |
| Channel       | `make(chan T)`     |
| Select        | `select { ... }`   |
| WaitGroup     | `sync.WaitGroup`   |
| Mutex         | `sync.Mutex`       |
| Atomic        | `sync/atomic`      |
| Ticker        | `time.NewTicker()` |
| Timeout       | `time.After()`     |
| Once          | `sync.Once`        |
| Close Channel | `close(chan)`      |

---