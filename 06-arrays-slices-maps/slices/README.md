# 📘 Go Language - Slices

Slices are powerful, flexible, and commonly used in Go. They provide a more dynamic alternative to arrays, allowing you to work with collections of data efficiently.

---

## ✅ Slice Basics

```go
nums := []int{10, 20, 30} // A slice of integers
fmt.Println(nums[0])      // Access first element
fmt.Println(len(nums))    // Length of the slice
```

📄 [View Code → `slice-basics.go`](./slice-basics.go)

---

## ✅ Create Slice from Array

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // → [2, 3, 4]
```

📄 [View Code → `slice-from-array.go`](./slice-from-array.go)

---

## ✅ Using `make()` to Create Slice

```go
s := make([]int, 3)       // Length 3, Capacity 3
s := make([]int, 3, 5)    // Length 3, Capacity 5
```

📄 [View Code → `slice-make.go`](./slice-make.go)

---

## ✅ Appending Elements

```go
s := []int{1, 2}
s = append(s, 3, 4)       // → [1, 2, 3, 4]
```

📄 [View Code → `slice-append.go`](./slice-append.go)

---

## ✅ Copying Slices

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```

📄 [View Code → `slice-copy.go`](./slice-copy.go)

---

## ✅ Looping Over a Slice

```go
nums := []int{10, 20, 30}
for i, val := range nums {
	fmt.Printf("Index: %d, Value: %d\n", i, val)
}
```

📄 [View Code → `slice-loop.go`](./slice-loop.go)

---

## ✅ Unicode-Friendly String Looping

```go
text := "Go 💻"
for i, r := range text {
	fmt.Printf("Index: %d, Rune: %q\n", i, r)
}
```

📄 [View Code → `slice-unicode.go`](./slice-unicode.go)

---

## ✅ Growing Capacity Example

```go
s := []int{}
for i := 0; i < 10; i++ {
	s = append(s, i)
	fmt.Printf("Len: %d, Cap: %d\n", len(s), cap(s))
}
```

📄 [View Code → `slice-growth.go`](./slice-growth.go)

---

## ✅ Summary Table

| Feature       | Slice                           | Array               |
| ------------- | ------------------------------- | ------------------- |
| Size          | Dynamic (can grow/shrink)       | Fixed               |
| Declaration   | `[]int{}` or `make([]int, ...)` | `[5]int{}`          |
| Operations    | `append()`, `copy()`, `len()`   | Only basic indexing |
| Preferred Use | ✅ Most use cases                | Rare (low-level)    |

---

## 📂 File Structure

```
05-slices/
├── slice-append.go
├── slice-basics.go
├── slice-copy.go
├── slice-from-array.go
├── slice-growth.go
├── slice-loop.go
├── slice-make.go
├── slice-unicode.go
└── README.md
```

---

## 🌐 GitHub Link (Replace with your own)

🔗 [GitHub Repo → go-lang-learning/05-slices](https://github.com/yourusername/go-lang-learning/tree/main/05-slices)

---

```

---