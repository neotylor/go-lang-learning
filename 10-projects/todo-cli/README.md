Great! Let's scaffold a scalable and best-practice-friendly Go project for a **CLI-based To-Do app**, named `todo-cli`.

We'll follow idiomatic Go project structure with **modular organization**, clean architecture principles, and CLI tooling (like `cobra` or native `flag` package support). Here's the recommended file/folder structure:

---

### 🗂️ Project Structure

```
todo-cli/
│
├── cmd/                     # CLI commands (main entry point)
│   └── root.go              # Root command for the CLI
│   └── add.go               # 'add' command
│   └── list.go              # 'list' command
│   └── complete.go          # 'complete' command
│   └── delete.go            # 'delete' command
│
├── internal/                # Internal business logic (not exposed outside)
│   ├── todo/                # Package for todo business logic
│   │   ├── service.go       # Core service logic
│   │   ├── model.go         # Todo data models
│   │   └── repository.go    # Data persistence logic (e.g., file or DB)
│
├── pkg/                     # Public reusable packages (optional)
│   └── logger/              # Logging utilities
│       └── logger.go
│
├── config/                  # Configuration related logic
│   └── config.go
│
├── data/                    # Data storage (e.g., JSON or DB files)
│   └── todos.json
│
├── go.mod                   # Go module definition
├── go.sum                   # Module checksum file
├── main.go                  # Main entry point
└── README.md
```

---

### 🛠 Steps to Create This Structure

#### 1. **Initialize the Module**

```bash
mkdir todo-cli && cd todo-cli
go mod init github.com/yourname/todo-cli
```

---

#### 2. **Create Directories**

```bash
mkdir -p cmd internal/todo pkg/logger config data
```

---

#### 3. **Create Entry File: `main.go`**

```go
// main.go
package main

import (
	"github.com/yourname/todo-cli/cmd"
)

func main() {
	cmd.Execute()
}
```

---

#### 4. **Setup CLI with Cobra (recommended)**

```bash
go install github.com/spf13/cobra-cli@latest
cobra-cli init
cobra-cli add add
cobra-cli add list
cobra-cli add complete
cobra-cli add delete
```

> This auto-generates the `cmd/` folder and basic command templates.

---

#### 5. **Create a Basic Todo Model**

```go
// internal/todo/model.go
package todo

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}
```

---

#### 6. **Create Repository Interface**

```go
// internal/todo/repository.go
package todo

type Repository interface {
	Add(todo Todo) error
	List() ([]Todo, error)
	Delete(id int) error
	Complete(id int) error
}
```

---

#### 7. **Create a File-based Repository Implementation (optional)**

We can use JSON in `data/todos.json` or use BoltDB/SQLite later.

---

#### 8. **Logger Utility**

```go
// pkg/logger/logger.go
package logger

import (
	"log"
)

func Info(msg string) {
	log.Println("[INFO]", msg)
}

func Error(msg string) {
	log.Println("[ERROR]", msg)
}
```

---

#### 9. **Add Configuration Loader**

```go
// config/config.go
package config

// You can define environment-based config here if needed
```

---
Perfect. We'll now implement the functionality of the CLI commands:

* `add`: Add a new todo.
* `list`: Show all todos.
* `complete`: Mark a todo as completed.
* `delete`: Delete a todo.

We'll use a **file-based JSON store** (`data/todos.json`) for simplicity.

---

## ✅ Prerequisites

Before adding command files, let’s define the reusable logic.

---

### `internal/todo/model.go`

```go
package todo

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}
```

---

### `internal/todo/repository.go`

```go
package todo

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

var filePath = "data/todos.json"
var mu sync.Mutex

type FileRepository struct{}

func NewRepository() *FileRepository {
	return &FileRepository{}
}

func (r *FileRepository) readTodos() ([]Todo, error) {
	mu.Lock()
	defer mu.Unlock()

	var todos []Todo
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &todos)
	return todos, err
}

func (r *FileRepository) writeTodos(todos []Todo) error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func (r *FileRepository) Add(task string) error {
	todos, err := r.readTodos()
	if err != nil {
		return err
	}
	id := 1
	if len(todos) > 0 {
		id = todos[len(todos)-1].ID + 1
	}
	newTodo := Todo{ID: id, Task: task, Completed: false}
	todos = append(todos, newTodo)
	return r.writeTodos(todos)
}

func (r *FileRepository) List() ([]Todo, error) {
	return r.readTodos()
}

func (r *FileRepository) Complete(id int) error {
	todos, err := r.readTodos()
	if err != nil {
		return err
	}
	found := false
	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = true
			found = true
			break
		}
	}
	if !found {
		return errors.New("todo not found")
	}
	return r.writeTodos(todos)
}

func (r *FileRepository) Delete(id int) error {
	todos, err := r.readTodos()
	if err != nil {
		return err
	}
	newTodos := make([]Todo, 0)
	found := false
	for _, t := range todos {
		if t.ID == id {
			found = true
			continue
		}
		newTodos = append(newTodos, t)
	}
	if !found {
		return errors.New("todo not found")
	}
	return r.writeTodos(newTodos)
}
```

---

## 🔧 CLI Command Files (`cmd/`)

Ensure you import the repository like:

```go
import "github.com/yourname/todo-cli/internal/todo"
```

Replace `yourname` accordingly.

---

### `cmd/add.go`

```go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yourname/todo-cli/internal/todo"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new todo task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := todo.NewRepository()
		task := args[0]
		err := repo.Add(task)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Todo added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
```

---

### `cmd/list.go`

```go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yourname/todo-cli/internal/todo"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		repo := todo.NewRepository()
		todos, err := repo.List()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(todos) == 0 {
			fmt.Println("No todos found.")
			return
		}
		for _, t := range todos {
			status := "[ ]"
			if t.Completed {
				status = "[x]"
			}
			fmt.Printf("%s %d: %s\n", status, t.ID, t.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
```

---

### `cmd/complete.go`

```go
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yourname/todo-cli/internal/todo"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark a todo as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := todo.NewRepository()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID:", args[0])
			return
		}
		err = repo.Complete(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Todo marked as completed!")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
```

---

### `cmd/delete.go`

```go
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yourname/todo-cli/internal/todo"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a todo by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := todo.NewRepository()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID:", args[0])
			return
		}
		err = repo.Delete(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Todo deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
```

---

### `cmd/root.go`

```go
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo is a CLI-based todo manager",
	Long:  `A simple CLI app to manage your todo tasks using Go.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
```

---

## ✅ Usage

Make sure `data/todos.json` exists, or it will be auto-created.

```bash
go run main.go add "Learn Go"
go run main.go list
go run main.go complete 1
go run main.go delete 1
```

---

### ✅ What's Next?

You can develop some more things:

* Fill in the contents of each CLI command (add, list, complete, delete)
* Add file-based JSON storage
* Add test cases
* Add logging or config parsing
* Add Cobra command auto-completion

Would you like me to generate and fill in the contents of each of the major files?


## 📦 Next Steps

* Add unit tests?
* Add color to the CLI output?
* Support for due dates or priorities?
* Convert to use BoltDB or SQLite for persistent storage?

Let me know! if you enhance!

Happy Coding!