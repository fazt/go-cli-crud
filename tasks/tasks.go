package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

// list all tasks from the array with an icon
func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas.")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Complete {
			status = "✓"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Name)
	}
}

// add a new task to the array
func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       GetNextID(tasks),
		Name:     name,
		Complete: false,
	}
	return append(tasks, newTask)
}

// search task with id and mark as completed
func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			break
		}
	}
	return tasks
}

// Delete a task with an id
func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

// get the next available id for the next task
func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}

// save all tasks in a json file
func SaveTasks(file *os.File, tasks []Task) {
	// convert to a json file
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	// move the pointer at the start
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// clean the entier file or delete everything
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	// write the file
	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	// make sure the content was written
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
