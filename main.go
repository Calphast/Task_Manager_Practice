package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	TITLE    string `json:"title"`
	COMPLETE bool   `json:"complete"`
}

var tasks []Task

const filename = "tasks.json"

func loadTasks() {
	file, err := os.ReadFile(filename)

	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(filename, data, 0644)
}

func addtask(title string) {
	id := len(tasks) + 1
	tasks = append(tasks, Task{ID: id, TITLE: title, COMPLETE: false})
	saveTasks()
	fmt.Println("Task added successfully!")
}

func listTask() {
	for _, task := range tasks {
		status := "❌"
		if task.COMPLETE {
			status = "✅"
		}
		fmt.Printf("[%d] %s %s\n", task.ID, status, task.TITLE)
	}
}

func completeTask(id int) {
	for i, task := range tasks {
		if task.ID == int(id) {
			fmt.Println("ID Found")
			tasks[i].COMPLETE = true
			saveTasks()
			fmt.Println("Task Complete")
			return
		}
	}
	fmt.Println("No task found")
}

func main() {
	loadTasks()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go-task-manager <command> [arguments]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go-task-manager add <task>")
			return
		}
		addtask(os.Args[2])
	case "list":
		listTask()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go-task-manager done <task_id>")
			return
		}

		var id int
		fmt.Sscanf(os.Args[2], "%d", &id)
		completeTask(id)
	default:
		fmt.Println("Unknown command. Use 'add', 'list', or 'done'.")
	}
}
