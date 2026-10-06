package task

import "time"

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