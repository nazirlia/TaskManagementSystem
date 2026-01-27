package main

import (
	"fmt"
	"strconv"
)

const totalTaskNumber = 100
const projectName = "Task Management System"
const (
	lowPriority = iota
	mediumPriority
	highPriority
)

func main() {
	var currentProjectStatus string
	var completedTaskNumber int
	var isProjectCompleted bool
	var isProjectCompletedString string

	currentProjectStatus = "IN PROGRESS"
	completedTaskNumber = 2
	isProjectCompleted = false
	isProjectCompletedString = strconv.FormatBool(isProjectCompleted)

	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is:", "2026-01-20 00:00:00")
	fmt.Println("Project:", projectName)

	fmt.Println("Current project status:", currentProjectStatus)
	fmt.Println("Tasks completed:", completedTaskNumber, "out of", totalTaskNumber)
	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", lowPriority, mediumPriority, highPriority)
	fmt.Println("Is Project Completed?:", isProjectCompletedString)
}
