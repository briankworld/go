package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Task struct {
  Description string
}

type TodoList struct {
  Tasks []Task
}

(t *TodoList)func Add(task string) {
  t.tasks = append(t.tasks, Task{ Description: task})
  fmt.Printf("Added task: %s\n", task)
}

func (t *TodoList) View() {
	if len(t.Tasks) == 0 {
		fmt.Println("Your ToDo list is empty.")
		return
	}
	fmt.Println("ToDo List:")
	for i, task := range t.Tasks {
		fmt.Printf("%d: %s\n", i+1, task.Description)
	}
}

(t *TodoList)func Remove(index int) {
  if index < 1 || index > len(t.Tasks) {
    fmt.Println("Invalid task number.")
		return
	}

  t.tasks = append(t.tasks[:index-1], t.tasks[index:]...)
  fmt.Printf("Removed task: %s\n", task)
}

func main() {
  todoList := &TodoList{}
  scanner := bufio.NewScanner(os.Stdin)

  for {

		fmt.Println("\nCommands:")
		fmt.Println("1. add [task] - Add a new task")
		fmt.Println("2. view - View all tasks")
		fmt.Println("3. remove [task number] - Remove a task")
		fmt.Println("4. exit - Exit the application")
		fmt.Print("Enter command: ")
    scanner.Scan()
    input := scanner.Text()
    commands := strings.Fields(input)

    switch commands[0] {
    case "add":
			task := strings.Join(commands[1:], " ")
			todoList.Add(task)

		case "view":
			todoList.View()

		case "remove":
			var index int
			fmt.Sscanf(parts[1], "%d", &index)
			todoList.Remove(index)

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Unknown command:", command)
    }
  }
}
