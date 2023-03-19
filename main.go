package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/faztweb/go-crud-cli/tasks"
)

func main() {
	// read or create the tasks.json file
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []task.Task

	// get file info
	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	// check the size
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	// check the user input and print options to choose
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Nombre de la tarea:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		task.SaveTasks(file, tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Debe proporcionar el ID de la tarea que se completará.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un número entero.")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveTasks(file, tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Debe proporcionar el ID de la tarea que se eliminará.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un número entero.")
			return
		}
		tasks = task.DeleteTask(tasks, id)
		task.SaveTasks(file, tasks)
	default:
		printUsage()
	}
}

// Imprimir el uso del programa
func printUsage() {
	fmt.Println("Uso: todo [list|add|complete|delete]")
}
