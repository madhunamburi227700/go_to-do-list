package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Title string
	Done  bool
}

var TaskList []Task

// Add a new task
func Add() {
	fmt.Print("Enter new task: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	TaskList = append(TaskList, Task{Title: title, Done: false})
	fmt.Println("Task added!")
}