# 📍 Understanding Maps in Go

## 🚀 What is a Map?

A **map** in Go is an **unordered collection of key-value pairs**. It allows you to **look up values** by their keys efficiently, similar to a dictionary in Python or an object in JavaScript.

---

## 🧪 How to Create a Map

```go
age := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
````

Here:

* `string` is the **key type**
* `int` is the **value type**
* `"Alice"` is a key, and `25` is the value associated with it

---

## 🛠️ Common Map Operations

### ✅ Add or Update a Value

```go
age["Charlie"] = 35
```

### ❌ Delete a Key

```go
delete(age, "Bob")
```

### 🔍 Check if a Key Exists

```go
val, ok := age["Bob"]
if ok {
    fmt.Println("Bob exists with age", val)
} else {
    fmt.Println("Bob not found")
}
```

### 🔁 Iterate Over a Map

```go
for name, age := range age {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

---

## 🧩 Advanced Map Examples

### Map of Structs

```go
type Person struct {
    Name string
    Age  int
}

people := map[string]Person{
    "p1": {"Alice", 28},
    "p2": {"Bob", 32},
}
```

### Map with Slice Values

```go
hobbies := map[string][]string{
    "Alice": {"Reading", "Biking"},
}
```

---

## ⚠️ Important Notes

* The **zero value** for a map is `nil`. You need to use `make()` or assign a literal to initialize it.

  ```go
  m := make(map[string]int)
  ```

* Keys must be **comparable types** (e.g., strings, numbers, booleans, pointers).

* Maps are **not safe for concurrent use** without proper synchronization (use `sync.Map` for that).

# Go Map Examples

This project contains multiple examples demonstrating the use of maps in Go.

## Files

- `basic_map.go`: Basic operations on maps (create, update, delete, check existence).
- `map_iteration.go`: How to iterate over a map using `range`.
- `map_of_structs.go`: Maps with struct values.
- `map_with_slices.go`: Maps with slice values.
- `map_key_check.go`: Checking for the existence of keys in a map.

---

## 🧪 Example Usage

```bash
go run basic_map.go
go run map_iteration.go
```

---

## 📚 Further Reading

* Go Official Maps Docs: [https://golang.org/ref/spec#Map\_types](https://golang.org/ref/spec#Map_types)
* Effective Go (Maps): [https://golang.org/doc/effective\_go#maps](https://golang.org/doc/effective_go#maps)

---

Happy Coding! 🚀
