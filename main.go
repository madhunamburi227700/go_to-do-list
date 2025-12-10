package main

import (
	"fmt"
	"go_to-do-list/tasks"
	"go_to-do-list/json_saved"
)

func showMenu() {
	fmt.Println("\n=== TO-DO LIST APP ===")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Mark Task as Done")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option: ")
}

func main() {

	// Load tasks from file on startup
	json_saved.LoadFromFile()
	showMenu()
	
	for {
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			tasks.Add()
			json_saved.SaveToFile()

		case "2":
			tasks.View()

		case "3":
			tasks.MarkDone()
			json_saved.SaveToFile()

		case "4":
			tasks.Delete()
			json_saved.SaveToFile()

		case "5":
			fmt.Println("Saving tasks... Goodbye!")
			json_saved.SaveToFile()
			return

		default:
			fmt.Println("Invalid option. Try again.")
		}
		
		fmt.Print("Choose an option: ")
	}
}
