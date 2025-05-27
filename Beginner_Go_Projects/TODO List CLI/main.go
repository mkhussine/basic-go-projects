// TODO List CLI
// Add/remove/display tasks using basic data structures.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Task struct {
	ID   int
	Name string
	Done bool
}

const filename = "tasks.json"

func printTasks(tasks []Task) {
	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, i, task.Name)
	}
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func loadTasks() ([]Task, error) {
	var tasks []Task
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return tasks, nil // No file yet
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return tasks, err
	}
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func main() {
	var tasks []Task
	tasks, _ = loadTasks()

	for {
		fmt.Println("\nWelcome to our TODO List CLI")
		fmt.Println("1- Add a new task")
		fmt.Println("2- List all tasks")
		fmt.Println("3- Mark task as done")
		fmt.Println("4- Save tasks to file")
		fmt.Println("5- Load tasks from file")
		fmt.Println("0- Exit")

		var input int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var taskName string
			fmt.Print("Enter task name: ")
			fmt.Scanln(&taskName)
			newTask := Task{ID: len(tasks), Name: taskName, Done: false}
			tasks = append(tasks, newTask)
			fmt.Println("Task added.")

		case 2:
			fmt.Println("Your Tasks:")
			printTasks(tasks)

		case 3:
			printTasks(tasks)
			var taskID int
			fmt.Print("Enter task ID to mark as done: ")
			fmt.Scanln(&taskID)
			if taskID >= 0 && taskID < len(tasks) {
				if tasks[taskID].Done {
					fmt.Println("Task is already done.")
				} else {
					tasks[taskID].Done = true
					fmt.Println("Task marked as done.")
				}
			} else {
				fmt.Println("Invalid task ID.")
			}

		case 4:
			err := saveTasks(tasks)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			} else {
				fmt.Println("Tasks saved successfully.")
			}

		case 5:
			tasks, _ = loadTasks()
			fmt.Println("Tasks loaded.")

		case 0:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}
