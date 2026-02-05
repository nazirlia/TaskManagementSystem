package main

import (
	"errors"
	"fmt"
	"strconv"
)

const totalTaskCount = 100
const projectName = "Task Management System"
const (
	lowPriority = iota
	mediumPriority
	highPriority
)

func main() {
	var currentProjectStatus string
	var completedTaskCount int
	var isProjectCompleted bool
	var isProjectCompletedString string
	var remainingTaskCount int
	var phase string

	completedTaskCount = 0

	if completedTaskCount > totalTaskCount {
		err := errors.New("ERROR: tasks completed exceeds total tasks, resetting to maximum")
		fmt.Println(err)
		completedTaskCount = totalTaskCount
	} else if completedTaskCount == 0 {
		panic("You need to complete at least one task")
	}

	if completedTaskCount == totalTaskCount {
		isProjectCompleted = true
		currentProjectStatus = "COMPLETED"
	} else {
		isProjectCompleted = false
		currentProjectStatus = "IN PROGRESS"
	}

	isProjectCompletedString = strconv.FormatBool(isProjectCompleted)
	remainingTaskCount = totalTaskCount - completedTaskCount

	switch {
	case completedTaskCount < 30:
		phase = "Starting phase"
	case completedTaskCount < 60 && completedTaskCount >= 30:
		phase = "Midway"
	case completedTaskCount >= 60 && completedTaskCount <= 100:
		phase = "Final phase"
	}

	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is:", "2026-01-20 00:00:00")
	fmt.Println("Project:", projectName)

	if completedTaskCount > 90 && completedTaskCount < totalTaskCount {
		fmt.Println("Project is near completion!")
	}

	fmt.Println("Current project status:", currentProjectStatus)
	fmt.Println("Project is in the", phase+".")
	fmt.Println("Tasks remaining:", remainingTaskCount, "out of", totalTaskCount)
	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", lowPriority, mediumPriority, highPriority)
	fmt.Println("Is Project Completed?:", isProjectCompletedString)

	for i := 1; i <= remainingTaskCount; i++ {
		fmt.Println("- Task", i)
	}
}
