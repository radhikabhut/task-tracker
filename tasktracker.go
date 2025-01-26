package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

type TaskManager struct {
	Tasks []Task `json:"tasks,omitempty"`
}

const taskFile = "task.json"

func loadTask() (*TaskManager, error) {
	file, err := os.Open(taskFile)
	if errors.Is(err, os.ErrNotExist) {
		return &TaskManager{}, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return &TaskManager{}, nil
	}
	taskManager := &TaskManager{}
	if err := json.Unmarshal(data, taskManager); err != nil {
		return nil, err
	}
	return taskManager, nil

}
func saveTasks(taskManager *TaskManager) error {
	data, err := json.MarshalIndent(taskManager, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}
func addTask(taskManager *TaskManager, description string, status string) {
	validStatuses := map[string]bool{"todo": true, "in-progress": true, "done": true}
	if !validStatuses[status] {
		fmt.Println("Invalid status.")
		return
	}
	id := len(taskManager.Tasks) + 1
	task := Task{
		ID:          id,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	taskManager.Tasks = append(taskManager.Tasks, task)
	fmt.Printf("Task added with ID: %d, Status: %s\n", id, status)
}
func updateTask(taskManager *TaskManager, id int, description string, status string) error {
	for i, task := range taskManager.Tasks {
		if task.ID == id {
			taskManager.Tasks[i].Description = description
			taskManager.Tasks[i].Status = status
			taskManager.Tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func deleteTask(taskManager *TaskManager, id int) error {
	for i, task := range taskManager.Tasks {
		if task.ID == id {
			taskManager.Tasks = append(taskManager.Tasks[:i], taskManager.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
func listTasks(taskManager *TaskManager, statusFilter string) {
	for _, task := range taskManager.Tasks {
		if statusFilter == "" || task.Status == statusFilter {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tasktracker <command> [arguments]")
		return
	}

	command := os.Args[1]
	taskManager, err := loadTask()
	if err != nil {
		fmt.Printf("Error loading tasks: %s\n", err)
		return
	}

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: tasktracker add <description> <status>")
			return
		}
		description := os.Args[2]
		status := os.Args[3]
		addTask(taskManager, description, status)
	case "update":
		if len(os.Args) < 5 {
			fmt.Println("Usage: tasktracker update <id> <description> <status>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		description := os.Args[3]
		status := os.Args[4]

		if err := updateTask(taskManager, id, description, status); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tasktracker delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := deleteTask(taskManager, id); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	case "list":
		statusFilter := ""
		if len(os.Args) >= 3 {
			statusFilter = os.Args[2]
		}
		listTasks(taskManager, statusFilter)

	default:
		fmt.Println("Unknown command")
	}
	if err := saveTasks(taskManager); err != nil {
		fmt.Printf("Error saving tasks: %s\n", err)
	}
}
