Here's a collection of **additional practical examples** for `break` and `continue` in Go. These examples cover:

* Nested loops with `break`
* Skipping even/odd numbers with `continue`
* Breaking on condition in a `range` loop
* Labelled `break` for nested loops

---

### 📄 Example 1: **Break on First Match in a Slice**

**File:** `break-slice-search.go`

```go
package main

import "fmt"

func main() {
	numbers := []int{3, 7, 9, 11, 15}
	target := 9

	for _, num := range numbers {
		if num == target {
			fmt.Println("Found target:", num)
			break
		}
	}
}
```

---

### 📄 Example 2: **Continue to Skip Even Numbers**

**File:** `continue-skip-even.go`

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i) // Will only print odd numbers
	}
}
```

---

### 📄 Example 3: **Break on Condition in Range Loop**

**File:** `break-range.go`

```go
package main

import "fmt"

func main() {
	words := []string{"apple", "banana", "stop", "cherry"}

	for _, word := range words {
		if word == "stop" {
			fmt.Println("Stopping at:", word)
			break
		}
		fmt.Println("Word:", word)
	}
}
```

---

### 📄 Example 4: **Nested Loop with Labelled `break`**

**File:** `labelled-break.go`

```go
package main

import "fmt"

func main() {
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j == 4 {
				fmt.Println("Breaking out of both loops at i*j =", i*j)
				break outer
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}
```

---

### 📄 Example 5: **Nested Loop with Labelled `continue`**

**File:** `labelled-continue.go`

```go
package main

import "fmt"

func main() {
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue outer // Skip to next iteration of outer loop
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}
```

---

## 📝 Linked These Examples

Add them under the **break/continue** section:

```markdown
## ✅ All Examples: break & continue

- 📄 [Break in slice search → `break-slice-search.go`](./break-slice-search.go)
- 📄 [Continue skipping even numbers → `continue-skip-even.go`](./continue-skip-even.go)
- 📄 [Break in range loop → `break-range.go`](./break-range.go)
- 📄 [Labelled break in nested loops → `labelled-break.go`](./labelled-break.go)
- 📄 [Labelled continue in nested loops → `labelled-continue.go`](./labelled-continue.go)
```
