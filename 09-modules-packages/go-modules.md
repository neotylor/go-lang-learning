Here is a **complete tutorial/course on Go Modules**, organized in a **step-by-step format** with examples for **all possible use cases**, from beginner to advanced.

---

# 📘 Course: Mastering Go Modules – Dependency Management in Golang

> Learn to initialize, manage, version, and publish Go modules. Includes all possible examples and best practices.

---

## 📦 Module 1: Introduction to Go Modules

### 🔹 What is a Go Module?

* A Go module is a collection of Go packages stored in a file tree with a `go.mod` file at its root.
* Modules replace GOPATH-based development.

---

## 🏁 Module 2: Initialize and Create a Go Module

### 📄 `myproject/main.go`

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Modules")
}
```

### ✅ Steps

```bash
mkdir myproject && cd myproject
go mod init github.com/username/myproject
go run main.go
```

🔍 `go.mod` created:

```go
module github.com/username/myproject

go 1.22
```

---

## 📁 Module 3: Internal Packages in a Module

### 📁 Structure

```
myproject/
├── go.mod
├── main.go
└── greet/
    └── greet.go
```

### 📄 `greet/greet.go`

```go
package greet

func Message() string {
	return "Hello from greet package!"
}
```

### 📄 `main.go`

```go
package main

import (
	"fmt"
	"github.com/username/myproject/greet"
)

func main() {
	fmt.Println(greet.Message())
}
```

---

## 🌐 Module 4: Using External Modules

### ✅ Add dependency (e.g. UUID)

```bash
go get github.com/google/uuid
```

### 📄 `main.go`

```go
import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("UUID:", uuid.New())
}
```

---

## 📦 Module 5: Replace & Rename Modules

### 🔹 Use `replace` for local modules or forks

```go
replace github.com/original/package => ../localpackage
```

🔹 Example:

```go
replace github.com/google/uuid => ../uuid-fork
```

---

## 🔧 Module 6: Go Commands with Modules

| Command          | Description                   |
| ---------------- | ----------------------------- |
| `go mod init`    | Initialize a new module       |
| `go get`         | Install a dependency          |
| `go build`       | Build your project            |
| `go mod tidy`    | Clean up unused modules       |
| `go list -m all` | List all modules              |
| `go mod verify`  | Check dependencies are intact |
| `go mod graph`   | Print module dependency graph |

---

## 🔁 Module 7: Nested Modules (Monorepo)

### ✅ Setup:

```
company/
├── go.mod
├── service-a/
│   ├── go.mod
│   └── main.go
├── service-b/
│   ├── go.mod
│   └── main.go
```

Each sub-folder is its own module with independent `go.mod`.

Use `replace` in submodules to reference shared code from the root if needed.

---

## 🚀 Module 8: Publishing Modules to GitHub

### 📄 `go.mod`

```go
module github.com/yourname/mylib
```

### ✅ Steps:

1. Push code to GitHub:

   ```bash
   git push origin main
   ```
2. Use in other project:

   ```bash
   go get github.com/yourname/mylib@latest
   ```

📌 Supports semantic versioning:

```bash
go get github.com/yourname/mylib@v1.0.0
```

---

## 🔢 Module 9: Versioning & Tags

### 🔹 Semantic Versioning

```bash
git tag v1.0.0
git push origin v1.0.0
```

### 🔹 Major version upgrade

For `v2+`, you **must** include `/v2` in the module name and folder.

📄 `go.mod`

```go
module github.com/yourname/mylib/v2
```

---

## 🧪 Module 10: Real-World Exercises

1. ✅ Create a module with two internal packages: `math` and `utils`.
2. ✅ Add an external dependency (e.g., `github.com/joho/godotenv`) and use it.
3. ✅ Refactor one module into two using `replace`.
4. ✅ Tag and version your module, and use it from another project.
5. ✅ Create a monorepo with `shared/`, `serviceA/`, `serviceB/`.

---

## 📘 Cheatsheet

| Task                     | Command / File                       |
| ------------------------ | ------------------------------------ |
| Initialize module        | `go mod init <module-name>`          |
| Add a dependency         | `go get <module>`                    |
| Clean unused packages    | `go mod tidy`                        |
| Replace module           | `replace <old> => <new>`             |
| Semantic version tagging | `git tag v1.0.0 && git push origin`  |
| Use module from GitHub   | `go get github.com/user/repo@v1.0.0` |
| Nested modules           | Create separate go.mod per service   |

---

## ✅ Optional Advanced Topics

* Private modules with `GOPRIVATE`
* Proxy config (`GOPROXY`)
* Vendoring with `go mod vendor`
* Module authentication

---

## 🏁 Final Project Structure (Full Example)

```
go-modules-course/
├── go.mod
├── main.go
├── mathutil/
│   └── mathutil.go
├── utils/
│   └── utils.go
└── README.md
```

---