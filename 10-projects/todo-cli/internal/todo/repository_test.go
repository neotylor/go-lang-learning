package todo

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestRepo(t *testing.T) (*FileRepository, string) {
	t.Helper()

	// Use a temporary test file
	tmpFile := filepath.Join(os.TempDir(), "test_todos.json")
	filePath = tmpFile // override global variable

	// Ensure clean state
	_ = os.Remove(tmpFile)

	return NewRepository(), tmpFile
}

func TestAddAndList(t *testing.T) {
	repo, tmp := setupTestRepo(t)
	defer os.Remove(tmp)

	err := repo.Add("Write Go tests")
	if err != nil {
		t.Fatalf("Failed to add todo: %v", err)
	}

	todos, err := repo.List()
	if err != nil {
		t.Fatalf("Failed to list todos: %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("Expected 1 todo, got %d", len(todos))
	}

	if todos[0].Task != "Write Go tests" {
		t.Errorf("Expected task 'Write Go tests', got '%s'", todos[0].Task)
	}
}

func TestComplete(t *testing.T) {
	repo, tmp := setupTestRepo(t)
	defer os.Remove(tmp)

	_ = repo.Add("Complete this task")
	_ = repo.Complete(1)

	todos, _ := repo.List()
	if !todos[0].Completed {
		t.Errorf("Expected todo to be completed")
	}
}

func TestDelete(t *testing.T) {
	repo, tmp := setupTestRepo(t)
	defer os.Remove(tmp)

	_ = repo.Add("Delete me")
	_ = repo.Delete(1)

	todos, _ := repo.List()
	if len(todos) != 0 {
		t.Errorf("Expected todos to be empty after deletion")
	}
}

func TestCompleteInvalidID(t *testing.T) {
	repo, tmp := setupTestRepo(t)
	defer os.Remove(tmp)

	err := repo.Complete(999)
	if err == nil {
		t.Errorf("Expected error when completing invalid ID")
	}
}

func TestDeleteInvalidID(t *testing.T) {
	repo, tmp := setupTestRepo(t)
	defer os.Remove(tmp)

	err := repo.Delete(999)
	if err == nil {
		t.Errorf("Expected error when deleting invalid ID")
	}
}
