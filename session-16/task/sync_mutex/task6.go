// package sync_mutex
package main

import (
	"fmt"
	"sync"
)

func GradeUpdaterOrAdder(name string, grade int, grades map[string]int, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	grades[name] = grade
}

func main() {

	var mu sync.Mutex
	var wg sync.WaitGroup
	grades := make(map[string]int)

	wg.Add(6)
	go func() {
		defer wg.Done()
		GradeUpdaterOrAdder("Garay", 97, grades, &mu)

	}()
	go func() {
		defer wg.Done()
		GradeUpdaterOrAdder("Garay", 15, grades, &mu)

	}()
	go func() {
		defer wg.Done()
		GradeUpdaterOrAdder("Garay", 90, grades, &mu)

	}()
	go func() {
		defer wg.Done()
		GradeUpdaterOrAdder("Ali", 47, grades, &mu)

	}()
	go func() {
		defer wg.Done()
		GradeUpdaterOrAdder("Ali", 85, grades, &mu)

	}()
	go func() {
		defer wg.Done()
		GradeUpdaterOrAdder("Medina", 78, grades, &mu)

	}()

	wg.Wait()

	fmt.Println(grades)

}
