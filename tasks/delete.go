package tasks

import (
	"fmt"
)



// Delete a task
func Delete() {
	View()
	if len(TaskList) == 0 {
		return
	}

	fmt.Print("Enter task number to delete: ")
	var num int
	fmt.Scanln(&num)

	if num < 1 || num > len(TaskList) {
		fmt.Println("Invalid task number.")
		return
	}

	TaskList = append(TaskList[:num-1], TaskList[num:]...)
	fmt.Println("Task deleted!")
}