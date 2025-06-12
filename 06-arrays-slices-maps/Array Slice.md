<!-- ````markdown -->
# 📘 Go Slice & Array Example – Learning Slice and Array Basics

This simple Go program demonstrates the basics of working with **slices** and **arrays** in Go, including:

- Creating and printing arrays and slices
- Checking their types and lengths
- Appending elements to slices
- Understanding how slices grow and how arrays differ

---

## 🧠 What Are Arrays and Slices?

### 🧩 Array
An **array** in Go is a fixed-size collection of elements of the same type. Its size is part of its type, and cannot be changed after creation.

### 🪄 Slice
A **slice** is a dynamically-sized, flexible view into an array. It has:
- A **pointer** to an underlying array
- A **length**
- A **capacity**

Slices are much more commonly used in Go because they are flexible and easier to work with.

---

## 📄 Code Example

```go
package main

import "fmt"

func main() {
    fmt.Println("Start Slice and Array learning")

    // -----------------------------
    // ✅ Array Example
    // -----------------------------
    var arr = [5]int{1, 2, 3, 4, 5}

    fmt.Println("Array:", arr) // Output: [1 2 3 4 5]
    fmt.Printf("Array Type: %T\n", arr) // Output: [5]int
    fmt.Println("Array Length:", len(arr)) // Output: 5

    // -----------------------------
    // ✅ Slice Example
    // -----------------------------
    numbers := []int{1, 2, 3, 4, 5}

    // Low-level print (not recommended for slices)
    println("Numbers:", numbers) // Output: [5/5]0xc00000c3c0

    // Recommended print using fmt
    fmt.Println("Numbers:", numbers) // Output: [1 2 3 4 5]

    // Print the type of the slice
    fmt.Printf("Numbers array (Slice) Type: %T\n", numbers) // Output: []int

    // Print the length of the slice
    fmt.Println("Numbers's array length:", len(numbers)) // Output: 5

    // Append more elements to the slice
    numbers = append(numbers, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

    // Low-level print again (not recommended)
    println("Updated Numbers:", numbers) // Output: [20/20]0xc000018140

    // Print updated slice contents
    fmt.Println("Updated Numbers:", numbers)

    // Print updated type (still []int)
    fmt.Printf("Updated Numbers array (Slice) Type: %T\n", numbers)

    // Print updated length
    fmt.Println("Updated Numbers's array length:", len(numbers)) // Output: 20
}
````

---

## 🔍 Output Explanation

### For the Slice:

#### `println("Numbers:", numbers)`

This uses Go’s **low-level built-in `println` function**, which doesn't print actual slice contents but instead gives:

```
[length/capacity] memory_address
```

Examples:

* `[5/5]0xc00000c3c0` — 5 elements, capacity 5
* `[20/20]0xc000018140` — After appending, slice has grown with a new backing array

Use `fmt.Println()` instead for readable output.

---

## 📌 Key Concepts Demonstrated

| Concept                  | Description                                                             |
| ------------------------ | ----------------------------------------------------------------------- |
| **Array declaration**    | `var arr = [5]int{1, 2, 3, 4, 5}` — fixed-size, type: `[5]int`          |
| **Slice declaration**    | `numbers := []int{1, 2, 3, 4, 5}` — flexible, dynamic type: `[]int`     |
| **Print values**         | `fmt.Println()` shows array/slice contents                              |
| **Print type**           | `fmt.Printf("%T", variable)` prints the type                            |
| **Length**               | `len()` returns number of elements                                      |
| **Append (slices only)** | `append(numbers, ...)` adds elements, may change capacity               |
| **Slice reallocation**   | When capacity is exceeded, Go creates a new backing array for the slice |

---

## ⚠️ Differences Between Arrays and Slices

| Feature           | Array                                    | Slice                           |
| ----------------- | ---------------------------------------- | ------------------------------- |
| Size              | Fixed                                    | Dynamic                         |
| Type              | Includes size (`[5]int`)                 | Does not include size (`[]int`) |
| Append            | ❌ Not allowed                            | ✅ Allowed                       |
| Flexible Capacity | ❌ No                                     | ✅ Yes                           |
| Use Case          | Rare (used in performance-specific code) | Common (preferred for most use) |

---

## ✅ Best Practices

* ✅ Prefer **slices** for most data-handling tasks in Go
* ✅ Use `fmt.Println()` or `fmt.Printf()` for printing complex types
* ❌ Avoid using `println()` for slices and other complex types
* 🧠 Understand that `append()` may allocate a new array and copy data

---

## 🚀 How to Run

To run the program:

```sh
go run main.go
```

---

## 📚 Additional Learning

* [Go Slices: Usage and Internals](https://go.dev/blog/slices-intro)
* [Effective Go](https://go.dev/doc/effective_go)
* [Arrays vs Slices](https://golang.org/ref/spec#Array_types)

---

Let me know if you'd like a visual diagram or comments embedded in the Go code for beginner understanding.

```

---