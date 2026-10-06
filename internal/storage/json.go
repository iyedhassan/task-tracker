package storage

import (
	"encoding/json"
	"os"

	"task-tracker/internal/task"
)

type JSONRepository struct {
	filePath string
}

func NewJSONRepository(filePath string) *JSONRepository {
	return &JSONRepository{
		filePath: filePath,
	}
}

func (r *JSONRepository) Load() ([]task.Task, error) {
	data, err := os.ReadFile(r.filePath)

	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}

		tasks := []task.Task{}
		if err := r.Save(tasks); err != nil {
			return nil, err
		}

		return tasks, nil
	}

	var tasks []task.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []task.Task{}
	}

	return tasks, nil
}

func (r *JSONRepository) Save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	data = append(data, '\n')

	return os.WriteFile(r.filePath, data, 0644)
}