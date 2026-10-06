package task

import (
	"fmt"
	"time"
)
type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) AddTask(description string) (Task, error) {
	tasks, err := s.repository.Load()
	if err != nil {
		return Task{}, err
	}

	now := time.Now()

	task := Task{
		ID:          nextID(tasks),
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)

	if err := s.repository.Save(tasks); err != nil {
		return Task{}, err
	}

	return task, nil
}

func nextID(tasks []Task) int {
	maxID := 0

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}
func (s *Service) ListTasks() ([]Task, error) {
	return s.repository.Load()
}

func (s *Service) UpdateTask(id int, description string) (Task, error) {
	tasks, err := s.repository.Load()
	if err != nil {
		return Task{}, err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()

			if err := s.repository.Save(tasks); err != nil {
				return Task{}, err
			}

			return tasks[i], nil
		}
	}

	return Task{}, fmt.Errorf("task with id %d not found", id)
}

func (s *Service) DeleteTask(id int) error {
	tasks, err := s.repository.Load()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			return s.repository.Save(tasks)
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}

func (s *Service) MarkDone(id int) (Task, error) {
	return s.updateStatus(id, StatusDone)
}

func (s *Service) MarkInProgress(id int) (Task, error) {
	return s.updateStatus(id, StatusInProgress)
}

func (s *Service) updateStatus(id int, status Status) (Task, error) {
	tasks, err := s.repository.Load()
	if err != nil {
		return Task{}, err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()

			if err := s.repository.Save(tasks); err != nil {
				return Task{}, err
			}

			return tasks[i], nil
		}
	}

	return Task{}, fmt.Errorf("task with id %d not found", id)
}

