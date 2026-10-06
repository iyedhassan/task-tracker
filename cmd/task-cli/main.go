package main

import (
	"fmt"
	"os"

	"task-tracker/internal/storage"
	"task-tracker/internal/task"
)

func main() {
	repository := storage.NewJSONRepository("tasks.json")
	service := task.NewService(repository)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := args[0]

	switch command {
case "add":
	if len(args) < 2 {
		fmt.Println("Usage: task-cli add <description>")
		return
	}

	description := args[1]

	createdTask, err := service.AddTask(description)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Task added successfully: %d\n", createdTask.ID)

case "list":
	tasks, err := service.ListTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {
		fmt.Printf(
			"%d | %s | %s\n",
			task.ID,
			task.Status,
			task.Description,
		)
	}

default:
	fmt.Println("Unknown command:", command)
}
}