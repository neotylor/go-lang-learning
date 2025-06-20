package todo

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type Repository interface {
	Add(todo Todo) error
	List() ([]Todo, error)
	Delete(id int) error
	Complete(id int) error
}

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
			return []Todo{}, nil // file doesn't exist → return empty list
		}
		return nil, err
	}

	if len(file) == 0 { // ✅ fix: skip unmarshalling empty files
		return []Todo{}, nil
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
	// func (r *FileRepository) Add(todo Todo) error {
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
