package tasks

import (
	"fmt"
)


// Display all tasks
func View() {
	if len(TaskList) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("\nYour Tasks:")
	for i, t := range TaskList {
		status := "Not Done"
		if t.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, t.Title, status)
	}
}