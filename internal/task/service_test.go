package task

import "testing"

type fakeRepository struct {
	tasks []Task
}

func (f *fakeRepository) Load() ([]Task, error) {
	return f.tasks, nil
}

func (f *fakeRepository) Save(tasks []Task) error {
	f.tasks = tasks
	return nil
}
func TestAddTask(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	createdTask, err := service.AddTask("Learn Go")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if createdTask.ID != 1 {
		t.Errorf("expected ID 1, got %d", createdTask.ID)
	}

	if createdTask.Description != "Learn Go" {
		t.Errorf("expected description 'Learn Go', got %q", createdTask.Description)
	}

	if createdTask.Status != StatusTodo {
		t.Errorf("expected status %q, got %q", StatusTodo, createdTask.Status)
	}

	if len(repository.tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(repository.tasks))
	}
}

func TestListTasks(t *testing.T) {
	repository := &fakeRepository{
		tasks: []Task{
			{ID: 1, Description: "Task 1", Status: StatusTodo},
			{ID: 2, Description: "Task 2", Status: StatusInProgress},
		},
	}

	service := NewService(repository)

	tasks, err := service.ListTasks()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].ID != 1 || tasks[1].ID != 2 {
		t.Fatalf("expected tasks in order [1,2], got [%d,%d]", tasks[0].ID, tasks[1].ID)
	}
}

func TestUpdateTask(t *testing.T) {
	repository := &fakeRepository{
		tasks: []Task{
			{
				ID:          1,
				Description: "Learn Go",
				Status:      StatusTodo,
			},
		},
	}

	service := NewService(repository)

	updatedTask, err := service.UpdateTask(1, "Learn Go deeply")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updatedTask.Description != "Learn Go deeply" {
		t.Errorf("expected description 'Learn Go deeply', got %q", updatedTask.Description)
	}

	if updatedTask.Status != StatusTodo {
		t.Errorf("expected status %q, got %q", StatusTodo, updatedTask.Status)
	}
}

func TestDeleteTask(t *testing.T) {
	repository := &fakeRepository{
		tasks: []Task{
			{ID: 1, Description: "Task 1"},
			{ID: 2, Description: "Task 2"},
		},
	}

	service := NewService(repository)

	err := service.DeleteTask(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repository.tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(repository.tasks))
	}

	if repository.tasks[0].ID != 2 {
		t.Errorf("expected remaining task ID 2, got %d", repository.tasks[0].ID)
	}
}

func TestDeleteTaskNotFound(t *testing.T) {
	repository := &fakeRepository{
		tasks: []Task{{ID: 1, Description: "Task 1"}},
	}

	service := NewService(repository)

	err := service.DeleteTask(99)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestMarkDone(t *testing.T) {
	repository := &fakeRepository{
		tasks: []Task{
			{
				ID:          1,
				Description: "Learn Go",
				Status:      StatusInProgress,
			},
		},
	}

	service := NewService(repository)

	updatedTask, err := service.MarkDone(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updatedTask.Status != StatusDone {
		t.Errorf("expected status %q, got %q", StatusDone, updatedTask.Status)
	}
}
func TestMarkInProgress(t *testing.T) {
	repository := &fakeRepository{
		tasks: []Task{
			{
				ID:          1,
				Description: "Learn Go",
				Status:      StatusTodo,
			},
		},
	}

	service := NewService(repository)

	updatedTask, err := service.MarkInProgress(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updatedTask.Status != StatusInProgress {
		t.Errorf("expected status %q, got %q", StatusInProgress, updatedTask.Status)
	}
}

func TestTaskNotFound(t *testing.T) {
	repository := &fakeRepository{
		tasks: []Task{
			{
				ID:          1,
				Description: "Learn Go",
				Status:      StatusTodo,
			},
		},
	}

	service := NewService(repository)

	_, err := service.UpdateTask(99, "Something")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestMarkDoneNotFound(t *testing.T) {
	repository := &fakeRepository{
		tasks: []Task{{ID: 1, Description: "Task 1", Status: StatusTodo}},
	}

	service := NewService(repository)

	_, err := service.MarkDone(99)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestMarkInProgressNotFound(t *testing.T) {
	repository := &fakeRepository{
		tasks: []Task{{ID: 1, Description: "Task 1", Status: StatusTodo}},
	}

	service := NewService(repository)

	_, err := service.MarkInProgress(99)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}