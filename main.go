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

// loadTasks loads the json file into the tasks[]
func loadTasks() {
	file, err := os.ReadFile(filename)

	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

// saveTasks saves the tasks json file in accordance to the tasks[]
func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(filename, data, 0644)
}

// addTask takes in a string and creates a task with the string as a title
// defaults to COMPLETE: false
func addTask(title string) {
	id := len(tasks) + 1
	tasks = append(tasks, Task{ID: id, TITLE: title, COMPLETE: false})
	saveTasks()
	fmt.Println("Task added successfully!")
}

// listTask lists all task that are save in the json file
func listTask() {
	if len(tasks) == 0 {
		fmt.Println("No current tasks.")
	}

	for _, task := range tasks {
		status := "❌"
		if task.COMPLETE {
			status = "✅"
		}
		fmt.Printf("[%d] %s %s\n", task.ID, status, task.TITLE)
	}
}

// completeTask takes in an id and changes its statur to complete
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
		addTask(os.Args[2])
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
