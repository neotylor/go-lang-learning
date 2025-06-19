Absolutely! Here's a comprehensive and structured **course/tutorial on Goroutines in Go (Golang)** — organized into **modules**, each containing **clear explanations and runnable code examples** that explore all essential and advanced use cases of goroutines.

---

# 📘 Course: Mastering Goroutines in Go

> Learn everything about Goroutines — the backbone of Go’s concurrency model — including creation, synchronization, error handling, and performance.

---

## 📦 Module 1: Introduction to Goroutines

### 🔹 Lesson 1.1: What is a Goroutine?

* A **goroutine** is a lightweight thread managed by the Go runtime.
* Use the `go` keyword to start a new goroutine.

```go
go someFunction()
```

### 🔹 Lesson 1.2: Minimal Goroutine Example

```go
func sayHello() {
	fmt.Println("Hello from Goroutine")
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second) // Wait for goroutine to complete
}
```

---

## 🚀 Module 2: Running Functions Concurrently

### 🔹 Lesson 2.1: Multiple Goroutines

```go
func printNumber(n int) {
	fmt.Println("Number:", n)
}

func main() {
	for i := 0; i < 5; i++ {
		go printNumber(i)
	}
	time.Sleep(1 * time.Second)
}
```

---

## 📥 Module 3: Synchronization with WaitGroup

### 🔹 Lesson 3.1: Use `sync.WaitGroup`

```go
var wg sync.WaitGroup

func worker(id int) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}
```

---

## 📬 Module 4: Communication with Channels

### 🔹 Lesson 4.1: Passing Data from Goroutine

```go
func sendMessage(ch chan string) {
	ch <- "Hello from goroutine"
}

func main() {
	ch := make(chan string)
	go sendMessage(ch)
	msg := <-ch
	fmt.Println(msg)
}
```

---

## 🔁 Module 5: Anonymous and Inline Goroutines

```go
func main() {
	go func(name string) {
		fmt.Println("Hello", name)
	}("GoLang")

	time.Sleep(1 * time.Second)
}
```

---

## 🧠 Module 6: Goroutines with Closures (Loop Bug)

### 🔹 Lesson 6.1: Common Mistake

```go
for i := 0; i < 5; i++ {
	go func() {
		fmt.Println("Incorrect i:", i) // might print 5 five times
	}()
}
```

### ✅ Fix with Parameter

```go
for i := 0; i < 5; i++ {
	go func(n int) {
		fmt.Println("Correct i:", n)
	}(i)
}
```

---

## ⛔ Module 7: Goroutine Leaks

### 🔹 Lesson 7.1: Example of Leak

```go
func leakyGoroutine(ch chan int) {
	for {
		val := <-ch // blocks forever if no value sent
		fmt.Println(val)
	}
}
```

✅ **Always ensure goroutines exit properly.**

---

## 🔂 Module 8: Goroutines and Select Statement

```go
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "One"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Two"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}
```

---

## ⏱️ Module 9: Goroutines with Timers & Tickers

### 🔹 Lesson 9.1: Using `time.After`

```go
go func() {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Finished after 2 seconds")
	}
}()
```

### 🔹 Lesson 9.2: Using `time.Ticker`

```go
ticker := time.NewTicker(1 * time.Second)
go func() {
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}()
time.Sleep(3 * time.Second)
ticker.Stop()
```

---

## 💥 Module 10: Panic Recovery in Goroutines

```go
func safeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		fn()
	}()
}

func main() {
	safeGo(func() {
		panic("something went wrong")
	})
	time.Sleep(1 * time.Second)
}
```

---

## 🔬 Module 11: Performance Testing

### 🔹 Lesson 11.1: Using `GOMAXPROCS`

```go
runtime.GOMAXPROCS(1) // Limit to 1 OS thread
```

### 🔹 Lesson 11.2: Benchmarks

Use `go test -bench .` to measure performance when spawning goroutines.

---

## 🧪 Module 12: Real-World Exercises

1. ✅ Build a downloader that fetches multiple URLs concurrently.
2. ✅ Use goroutines to compute parts of a large sum in parallel.
3. ✅ Implement a worker pool with goroutines.
4. ✅ Create a producer-consumer model with channels and goroutines.
5. ✅ Use a `context.WithCancel` to stop a long-running goroutine gracefully.

---

## ✅ Module 13: Best Practices

| Tip                                    | Why it matters            |
| -------------------------------------- | ------------------------- |
| Avoid blocking goroutines indefinitely | Prevent leaks             |
| Use WaitGroup or Channels to sync      | Avoid Sleep hacks         |
| Handle panics inside goroutines        | Prevent crashing          |
| Limit goroutines with pools            | Prevent memory exhaustion |
| Use Go Race Detector                   | Detect data races         |

---

## 📘 Summary Cheatsheet

| Task                           | Snippet                             |
| ------------------------------ | ----------------------------------- |
| Start goroutine                | `go function()`                     |
| Wait for completion            | `sync.WaitGroup`                    |
| Communicate between goroutines | `chan`                              |
| Handle panic safely            | `defer recover()`                   |
| Prevent goroutine leak         | Use `context` or `select` with exit |
| Range-safe goroutine           | `go func(n int) { }(i)`             |

---
