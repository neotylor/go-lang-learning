Here is a **complete tutorial/course-style guide** on the **`defer` keyword in Go**, covering **all possible use cases**, organized into modules with **examples, use cases**, and **best practices**.

---

# 📘 Course: Mastering `defer` in Go (Golang)

## 🧠 What is `defer`?

* The `defer` statement **delays the execution** of a function **until the surrounding function returns**.
* Commonly used for: **clean-up**, **unlocking**, **closing files**, **recovering from panics**, etc.

---

## 🧭 Module 1: Basics of `defer`

### 🔹 Lesson 1.1: Basic Syntax

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

**Output:**

```
hello
world
```

✅ The deferred call runs **after** `main()` finishes.

---

## ⏳ Module 2: Execution Order

### 🔹 Lesson 2.1: Last-In-First-Out (LIFO)

```go
func main() {
    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")
}
```

**Output:**

```
third
second
first
```

✅ Deferred functions are executed in **reverse order**.

---

## 🧮 Module 3: With Parameters

### 🔹 Lesson 3.1: Arguments are evaluated **immediately**, not at `defer` execution.

```go
func main() {
    x := 10
    defer fmt.Println("Deferred:", x)
    x = 20
    fmt.Println("Changed:", x)
}
```

**Output:**

```
Changed: 20
Deferred: 10
```

✅ `x` is captured at the point of **defer declaration**.

---

## 🧹 Module 4: Common Use Cases

### 🔹 Lesson 4.1: File Handling

```go
file, err := os.Open("example.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

✅ Ensures the file is closed **when function returns**, even if errors occur.

---

### 🔹 Lesson 4.2: Mutex Unlock

```go
mu := sync.Mutex{}
mu.Lock()
defer mu.Unlock()
```

✅ Ensures you **unlock the mutex** no matter how the function exits.

---

### 🔹 Lesson 4.3: Database Connection Cleanup (Simplified)

```go
db, _ := sql.Open("sqlite3", "example.db")
defer db.Close()
```

✅ Ensures DB connections are **gracefully released**.

---

## ⚠️ Module 5: Defer and Return Values

### 🔹 Lesson 5.1: Named Return Values + Defer

```go
func test() (result int) {
    defer func() {
        result += 1
    }()
    return 5
}
```

**Output:** `6`

✅ Deferred function can **modify named return values**.

---

### 🔹 Lesson 5.2: No Named Return Value

```go
func test() int {
    result := 5
    defer func() {
        result += 1
    }()
    return result
}
```

**Output:** `5`

❌ It **doesn’t modify** the return value unless it's named.

---

## 🛑 Module 6: Defer + Panic + Recover

### 🔹 Lesson 6.1: Graceful Recovery

```go
func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    fmt.Println(a / b)
}
```

```go
safeDivide(10, 0)
```

✅ Use `defer` to handle **panics** gracefully with `recover`.

---

## 🧪 Module 7: Advanced Patterns

### 🔹 Lesson 7.1: Defer Loop Bug (Common Mistake)

```go
for i := 0; i < 3; i++ {
    defer fmt.Println("Deferred:", i)
}
```

**Output:**

```
Deferred: 2
Deferred: 1
Deferred: 0
```

✅ Works, but…

### 🔹 Lesson 7.2: Deferred Closures Capture Last Value

```go
for i := 0; i < 3; i++ {
    defer func() {
        fmt.Println("Captured i:", i)
    }()
}
```

**Output:**

```
Captured i: 3
Captured i: 3
Captured i: 3
```

❌ Common bug — all closures capture the **same final value** of `i`.

### ✅ Fix:

```go
for i := 0; i < 3; i++ {
    i := i // new variable in each iteration
    defer func() {
        fmt.Println("Fixed i:", i)
    }()
}
```

---

## 🔁 Module 8: Performance & Considerations

### ⚠️ Defer is not free

```go
start := time.Now()
for i := 0; i < 100000; i++ {
    defer func() {}()
}
fmt.Println("Elapsed:", time.Since(start))
```

✅ Defer adds **overhead**, especially in tight loops.

---

## 🎓 Final Module: Best Practices

| ✅ Do                                         | ❌ Don’t                                 |
| -------------------------------------------- | --------------------------------------- |
| Use for file/mutex cleanup                   | Use in tight loops                      |
| Handle panic with defer                      | Rely on defer to fix logic bugs         |
| Use named return with defer-modified returns | Assume all return values can be changed |

---

## ✅ Summary

| Concept     | Usage                                   |
| ----------- | --------------------------------------- |
| Clean-up    | File close, DB disconnect, mutex unlock |
| Safety      | Panic + recover                         |
| Order       | LIFO                                    |
| Arguments   | Evaluated immediately                   |
| Performance | Adds runtime cost                       |

---

## 🔨 Bonus: Practice Challenges

1. Implement a **safe HTTP handler** using `defer + recover`.
2. Measure **execution time** of a function using `defer`.
3. Create a **logging wrapper** that logs entry/exit with `defer`.

```go
func loggedFunc() {
    start := time.Now()
    defer func() {
        fmt.Println("Execution time:", time.Since(start))
    }()
    // do work
}
```

---