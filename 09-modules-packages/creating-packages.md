Here is a **complete tutorial/course on creating and using packages in Go (Golang)**. It includes **all key concepts, patterns, and examples**, organized into **modules**, from basic custom packages to advanced structuring and reusable libraries.

---

# 📘 Course: Creating and Using Packages in Go (Golang)

> Learn how to organize, reuse, and distribute Go code using packages — Go’s fundamental unit of modularity and encapsulation.

---

## 📦 Module 1: What is a Package?

* A **package** is a directory with `.go` files sharing the same `package` name.
* Every Go program starts in the `main` package.

---

## 📁 Module 2: Project Structure for Packages

```
myapp/
├── main.go
└── greet/
    └── greet.go
```

---

## ✨ Module 3: Creating a Basic Package

### 🔹 Lesson 3.1: Define a Custom Package

📄 `greet/greet.go`

```go
package greet

import "fmt"

func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
```

---

### 🔹 Lesson 3.2: Use the Package in main.go

📄 `main.go`

```go
package main

import "myapp/greet" // import local package

func main() {
	greet.Hello("Vijendra")
}
```

> ✅ Run with:

```bash
go run main.go
```

---

## 🧠 Module 4: Package Naming and Visibility

### 🔹 Lesson 4.1: Capitalization = Exported

```go
func Hello()   // Exported (public)
func private() // Unexported (private)
```

Only exported identifiers (capitalized) can be accessed from other packages.

---

## 🗃️ Module 5: Grouping Functions in Packages

### 📄 `mathutil/mathutil.go`

```go
package mathutil

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}
```

### 📄 `main.go`

```go
import "myapp/mathutil"

fmt.Println(mathutil.Add(5, 3))     // 8
fmt.Println(mathutil.Subtract(5, 3)) // 2
```

---

## 🧱 Module 6: Structs and Methods in Packages

📄 `person/person.go`

```go
package person

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hi, I’m " + p.Name
}
```

📄 `main.go`

```go
import "myapp/person"

p := person.Person{Name: "Ravi", Age: 25}
fmt.Println(p.Greet()) // Hi, I’m Ravi
```

---

## 🧩 Module 7: Nested and Sub-Packages

```go
myapp/
└── utils/
    ├── strings/
    │   └── strings.go
    └── math/
        └── math.go
```

* `import "myapp/utils/strings"`
* `import "myapp/utils/math"`

---

## ⚙️ Module 8: init() Function in Packages

📄 `greet/init.go`

```go
package greet

import "fmt"

func init() {
	fmt.Println("Greet package initialized")
}
```

Will run **once** when package is imported.

---

## 🔁 Module 9: Reusability and Testing

📄 `mathutil/mathutil_test.go`

```go
package mathutil

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Fatal("Expected 5")
	}
}
```

> ✅ Run with:

```bash
go test ./mathutil
```

---

## 🌍 Module 10: Using External Packages

```bash
go get github.com/google/uuid
```

📄 `main.go`

```go
import "github.com/google/uuid"

id := uuid.New()
fmt.Println(id)
```

---

## 📦 Module 11: Creating a Reusable Go Module

```bash
go mod init github.com/yourname/mypackage
```

Export with:

```bash
git push https://github.com/yourname/mypackage
```

Then others can import:

```go
import "github.com/yourname/mypackage/..."
```

---

## 📘 Cheatsheet Summary

| Task                       | Example                                   |
| -------------------------- | ----------------------------------------- |
| Create package             | `package mypkg`                           |
| Use package                | `import "project/mypkg"`                  |
| Exported symbol            | Starts with Capital (e.g., `Func`, `Var`) |
| Package folder = namespace | Folder name becomes part of import path   |
| Test package               | `go test ./pkgname`                       |
| Initialize on import       | `func init()`                             |

---

## 🧪 Real-World Exercises

1. ✅ Create a `calc` package with `Add`, `Subtract`, `Multiply`, and `Divide`.
2. ✅ Build a `logger` package with `Info`, `Warn`, and `Error` methods.
3. ✅ Build a `models` package with multiple struct definitions.
4. ✅ Group multiple sub-packages in a `utils` folder.
5. ✅ Build and push your own open-source Go module on GitHub.

---

## 🏗️ Sample Final Project Structure

```
myapp/
├── go.mod
├── main.go
├── greet/
│   └── greet.go
├── mathutil/
│   ├── mathutil.go
│   └── mathutil_test.go
├── person/
│   └── person.go
```
