package json_saved

import (
	"encoding/json"
	"fmt"
	"os"
	"go_to-do-list/tasks"
)

const fileName = "tasks.json"

// Save tasks to JSON file
func SaveToFile() {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(tasks.TaskList)
}

// Load tasks from JSON file
func LoadFromFile() {
	file, err := os.Open(fileName)
	if err != nil {
		// If file doesn't exist, ignore
		return
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&tasks.TaskList)
}
