# 📘 Go Language - Control Structures

This document provides clear examples and explanations of Go control structures, including conditional statements, loops, and `switch` cases.

---

## ✅ 1. `if`, `else if`, `else`

📄 [View Code → `if-else.go`](./if-else.go)

```go
package main

import "fmt"

func main() {
	x := 10

	if x > 0 {
		fmt.Println("x is positive")
	} else if x < 0 {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is zero")
	}
}
```

**Notes:**

* No parentheses around conditions.
* `{}` braces are required.
* You can initialize variables within `if`:
  `if y := someFunc(); y > 10 { ... }`

---

## ✅ 2. `switch` Statement

📄 [View Code → `switch.go`](./switch.go)

### ➤ Basic Usage

```go
package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}
}
```

### ➤ Multiple Values in One Case

```go
switch day {
	case 1, 2, 3:
		fmt.Println("Start of the week")
}
```

### ➤ Conditional Switch

```go
temp := 30

switch {
	case temp > 35:
		fmt.Println("Too hot")
	case temp > 25:
		fmt.Println("Warm")
	default:
		fmt.Println("Cool")
}
```

**Notes:**

* No `break` needed (it's implicit).
* Use `fallthrough` to continue to next case.

---

## ✅ 3. `for` Loop

📄 [View Code → `for-loop.go`](./for-loop.go)

### ➤ Classic For Loop

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

### ➤ While-Style Loop

📄 [View Code → `while-style.go`](./for-loop/while-style.go)

```go
i := 0
for i < 5 {
	fmt.Println(i)
	i++
}
```

### ➤ Infinite Loop

📄 [View Code → `infinite-loop.go`](./for-loop/infinite-loop.go)

```go
for {
	fmt.Println("Infinite loop")
	break // Required to exit
}
```

---

## ✅ 4. `for` with `range`

📄 [View Code → `range-slice.go`](./range-slice.go)

### ➤ Looping Over Slice

```go
nums := []int{2, 4, 6}
for index, value := range nums {
	fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

### ➤ Looping Over Map

📄 [View Code → `range-map.go`](./range-map.go)

```go
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
	fmt.Printf("%s: %d\n", key, value)
}
```

### ➤ Ignoring Index or Value

```go
for _, value := range nums {
	fmt.Println(value)
}
```

---

## ✅ 5. `break` and `continue`

📄 [View Code → `break-continue.go`](./break-continue.go)

```go
for i := 1; i <= 5; i++ {
	if i == 3 {
		continue // Skip this iteration
	}
	if i == 4 {
		break // Exit loop
	}
	fmt.Println(i)
}
```

---

## ✅ Summary

| Structure            | Example                           | Notes                   | Link                                     | 
| -------------------- | --------------------------------- | ----------------------- | ---------------------------------------- |
| `if`, `else`         | `if x > 5 {}`                     | No `()` needed          | [if-else.go](./if-else.go)               |
| `switch`             | `switch value {}`                 | No `break` needed       | [switch.go](./switch.go)                 |
| `for`                | `for i := 0; i < n; i++ {}`       | Only loop in Go         | [for-loop.go](./for-loop/for-loop.go)             |
| `range`              | `for k, v := range collection {}` | Works with slices, maps | [range-slice.go](./range-slice.go)       |
| `break` / `continue` | Used inside loops                 | Control loop flow       | [break-continue.go](./break-continue/break-continue.go) |

---

## 📂 Our File Structure

```
04-control-structures/
├── break-continue  //learn after loop
|	├── break-continue.go
|	└── README.md
├── for-loop
|	├── for-loop.go
|	├── infinite-loop.go
|	├── while-style.go
|	└── README.md
├── if-else.go
├── range-map.go
├── range-loop.go
├── range-slice.go
├── switch.go
└── README.md


```

> 🔗 Continue learning in [05-functions](../05-functions)

---

## 🌐 GitHub Repository

🔗 [GitHub: go-lang-learning](https://github.com/neotylor/go-lang-learning/tree/master/04-control-structures)

---

Happy coding! 🎉

```




