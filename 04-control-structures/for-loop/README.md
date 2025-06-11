Here are **more practical and diverse examples of `for` loops** in Go. These include:

* Countdown loop
* Loop with multiple variables
* Looping over strings (Unicode handling)
* Summing elements in a slice
* Custom step increment
* Nested `for` loop (matrix/grid)

Each example is easy to understand and can be saved as individual `.go` files in this `04-control-structures` directory.

---

## 📄 Example 1: **Countdown Loop**

**File:** `for-countdown.go`

```go
package main

import "fmt"

func main() {
	for i := 5; i >= 1; i-- {
		fmt.Println("Countdown:", i)
	}
	fmt.Println("Liftoff!")
}
```

---

## 📄 Example 2: **For Loop with Multiple Variables**

**File:** `for-multiple-vars.go`

```go
package main

import "fmt"

func main() {
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
}
```

---

## 📄 Example 3: **Summing Elements in a Slice**

**File:** `for-slice-sum.go`

```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)
}
```

---

## 📄 Example 4: **Loop Over Unicode String (Rune Handling)**

**File:** `for-unicode.go`

```go
package main

import "fmt"

func main() {
	word := "Golang 💻"
	for i, r := range word {
		fmt.Printf("Index %d: Rune %q\n", i, r)
	}
}
```

📝 *Note: Using `range` on a string handles Unicode runes correctly.*

---

## 📄 Example 5: **Custom Step Increment (e.g., i += 2)**

**File:** `for-step.go`

```go
package main

import "fmt"

func main() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println(i) // Prints even numbers
	}
}
```

---

## 📄 Example 6: **Nested For Loops (Grid / Matrix Print)**

**File:** `for-nested-grid.go`

```go
package main

import "fmt"

func main() {
	rows := 3
	cols := 4

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
}
```

---

## ✅ Here all link of example

Here’s link:

```markdown
## ✅ More Examples: `for` Loop

- 📄 [Countdown loop → `for-countdown.go`](./for-countdown.go)
- 📄 [For with multiple variables → `for-multiple-vars.go`](./for-multiple-vars.go)
- 📄 [Sum slice elements → `for-slice-sum.go`](./for-slice-sum.go)
- 📄 [Loop over Unicode string → `for-unicode.go`](./for-unicode.go)
- 📄 [Custom step increment → `for-step.go`](./for-step.go)
- 📄 [Nested for loop (grid) → `for-nested-grid.go`](./for-nested-grid.go)
```