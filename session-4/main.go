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
	var completedTaskCount int
	var isProjectCompleted bool
	var isProjectCompletedString string
	var remainingTaskCount int
	var projectCompletionEstimate bool

	completedTaskCount = 99

	isProjectCompleted, completedTaskCount = taskCompletionChecks(completedTaskCount)
	remainingTaskCount, projectCompletionEstimate = remainingTaskCounter(completedTaskCount)

	isProjectCompletedString = strconv.FormatBool(isProjectCompleted)

	fmt.Println("Status check completed.")
	fmt.Println("Task list:")
	completedTaskListing("Task 1", "Task 2", "Task 3", "Task 4", "Task 5")

	fmt.Println("Recursive countdown of remaining tasks:")
	reverseRemainingTaskListing(remainingTaskCount)
	fmt.Println("Project near to completion?", projectCompletionEstimate)
	fmt.Println("Is the project completed?", isProjectCompletedString)

}

func taskCompletionChecks(completedTaskCount int) (bool, int) {

	var isProjectCompleted bool
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			completedTaskCount = totalTaskCount
		}
	}()

	if completedTaskCount > totalTaskCount {
		panic("Error: tasksCompleted exceeds total tasks")
	}

	if completedTaskCount == 0 {
		err := errors.New("You need to complete at least one task")
		fmt.Println(err)
		return false, completedTaskCount
	}

	if completedTaskCount == totalTaskCount {
		isProjectCompleted = true
	} else {
		isProjectCompleted = false
	}

	return isProjectCompleted, completedTaskCount

}

func remainingTaskCounter(completedTaskCount int) (int, bool) {
	var remainingTaskCount int
	var projectCompletionEstimate bool = false
	remainingTaskCount = totalTaskCount - completedTaskCount
	if completedTaskCount > 90 {
		projectCompletionEstimate = true
	}
	return remainingTaskCount, projectCompletionEstimate
}

func reverseRemainingTaskListing(remainingTaskCount int) {
	if remainingTaskCount <= 0 {
		return
	}
	if remainingTaskCount == 1 {
		fmt.Println("Tasks remaining:", remainingTaskCount)
		fmt.Println("All tasks completed")
		return
	}
	fmt.Println("Tasks remaining:", remainingTaskCount)
	reverseRemainingTaskListing(remainingTaskCount - 1)
}

func completedTaskListing(taskNames ...string) {
	for i, name := range taskNames {
		fmt.Printf("%d: %s\n", i+1, name)
	}
}
