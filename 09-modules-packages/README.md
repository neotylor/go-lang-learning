Here's a **complete, structured course/tutorial** on using **Modules + Packages** in Go, covering **all possible examples** — from basics to advanced real-world structuring.

---

# 📘 Course: Mastering Go Modules & Packages

> Learn how to organize, structure, and manage Go applications using **modules and packages** with best practices and real examples.

---

## 🧩 Overview

* **Package**: A reusable unit of Go code (folder + `package` declaration).
* **Module**: A collection of packages with a `go.mod` file at the root.
* Go Modules are the standard for dependency and project management (since Go 1.11+).

---

## 📦 Module 1: Setup and Initialization

### ✅ Lesson 1.1: Creating a New Module

```bash
mkdir myapp && cd myapp
go mod init github.com/yourname/myapp
```

✅ Creates a `go.mod` file:

```go
module github.com/yourname/myapp

go 1.22
```

---

## 🗂️ Module 2: Creating and Using Packages

### 📁 Structure

```
myapp/
├── go.mod
├── main.go
└── greet/
    └── greet.go
```

### 📄 greet/greet.go

```go
package greet

func Hello(name string) string {
	return "Hello, " + name
}
```

### 📄 main.go

```go
package main

import (
	"fmt"
	"github.com/yourname/myapp/greet"
)

func main() {
	fmt.Println(greet.Hello("Vijendra"))
}
```

---

## 🧱 Module 3: Multiple Packages in One Module

### 📁 Structure

```
myapp/
├── go.mod
├── main.go
├── mathutil/
│   └── math.go
└── stringutil/
    └── string.go
```

### 📄 mathutil/math.go

```go
package mathutil

func Add(a, b int) int {
	return a + b
}
```

### 📄 stringutil/string.go

```go
package stringutil

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
```

### 📄 main.go

```go
import (
	"fmt"
	"github.com/yourname/myapp/mathutil"
	"github.com/yourname/myapp/stringutil"
)

func main() {
	fmt.Println(mathutil.Add(2, 3))
	fmt.Println(stringutil.ToUpper("hello"))
}
```

---

## 🧠 Module 4: Exported vs Unexported

| Name        | Access     | Example          |
| ----------- | ---------- | ---------------- |
| Capitalized | ✅ Exported | `Add`, `ToUpper` |
| lowercase   | ❌ Private  | `helperFunc`     |

Only **exported names** (capitalized) can be accessed from other packages.

---

## 📁 Module 5: init() Functions in Packages

```go
package greet

import "fmt"

func init() {
	fmt.Println("greet package initialized")
}
```

This runs **automatically** when the package is imported.

---

## 🔄 Module 6: Internal Packages (Encapsulation)

```
myapp/
├── go.mod
├── main.go
└── internal/
    └── helper/
        └── helper.go
```

### `internal/helper/helper.go`

```go
package helper

func HiddenFunc() string {
	return "Internal Only"
}
```

✅ This package can **only be imported** by code inside `myapp/`.

❌ External code **cannot access** `internal/helper`.

---

## 🔁 Module 7: Using External Modules (Dependencies)

```bash
go get github.com/google/uuid
```

```go
import "github.com/google/uuid"

id := uuid.New()
```

Use `go mod tidy` to clean unused modules.

---

## 🔂 Module 8: Nested Modules (Monorepo Structure)

### ✅ Create two modules in same repo:

```
monorepo/
├── shared/
│   ├── go.mod (github.com/yourname/shared)
│   └── shared.go
└── service/
    ├── go.mod (github.com/yourname/service)
    └── main.go
```

Use `replace` in `service/go.mod`:

```go
replace github.com/yourname/shared => ../shared
```

Now `service` can import `github.com/yourname/shared`.

---

## 🔗 Module 9: Versioning and Publishing Modules

```bash
git tag v1.0.0
git push origin v1.0.0
```

Update imports in another project:

```bash
go get github.com/yourname/mypkg@v1.0.0
```

For **v2+**, use:

* folder: `mypkg/v2`
* go.mod: `module github.com/yourname/mypkg/v2`

---

## 🧪 Module 10: Exercises

1. ✅ Create a module with 3 packages: `mathutil`, `dateutil`, `stringutil`.
2. ✅ Use an external module like `uuid`, `cobra`, or `gorilla/mux`.
3. ✅ Build a submodule (`/internal/logger`) and use it inside the module.
4. ✅ Create a monorepo with `shared/`, `api/`, and `cli/` modules.
5. ✅ Publish a versioned Go module on GitHub.

---

## 📘 Cheatsheet

| Task                     | Command / Concept                      |
| ------------------------ | -------------------------------------- |
| Init module              | `go mod init <module-path>`            |
| Add dependency           | `go get <module>`                      |
| Clean unused modules     | `go mod tidy`                          |
| Create internal package  | Place in `/internal/`                  |
| Exported function/struct | Start with uppercase                   |
| Private function/struct  | Start with lowercase                   |
| Publish                  | `git push && git tag v1.0.0`           |
| Use another module       | `import "github.com/username/pkgname"` |

---

## 🧱 Real-World Project Structure

```
myapp/
├── go.mod
├── main.go
├── mathutil/
│   └── math.go
├── stringutil/
│   └── string.go
├── internal/
│   └── config/
│       └── config.go
├── vendor/
├── README.md
```

---