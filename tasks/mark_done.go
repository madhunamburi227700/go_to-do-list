package tasks

import (
	"fmt"
)


// Mark task as done
func MarkDone() {
	View()
	if len(TaskList) == 0 {
		return
	}

	fmt.Print("Enter task number to mark as done: ")
	var num int
	fmt.Scanln(&num)

	if num < 1 || num > len(TaskList) {
		fmt.Println("Invalid task number.")
		return
	}

	TaskList[num-1].Done = true
	fmt.Println("Task marked as done!")
}