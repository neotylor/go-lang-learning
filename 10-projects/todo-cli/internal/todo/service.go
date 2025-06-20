package todo

import (
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// func (s *Service) AddTask(task string) error {
// 	if task == "" {
// 		return errors.New("task cannot be empty")
// 	}
// 	return s.repo.Add(task)
// }

func (s *Service) ListTasks() ([]Todo, error) {
	return s.repo.List()
}

func (s *Service) CompleteTask(id int) error {
	if id <= 0 {
		return errors.New("invalid task ID")
	}
	return s.repo.Complete(id)
}

func (s *Service) DeleteTask(id int) error {
	if id <= 0 {
		return errors.New("invalid task ID")
	}
	return s.repo.Delete(id)
}
