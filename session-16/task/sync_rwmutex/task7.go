package main

import (
	"fmt"
	"sync"
)

func Reader(name string, ages map[string]int, mu *sync.RWMutex) {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Printf("%s's age is %d\n", name, ages[name])
}

func Writer(name string, age int, ages map[string]int, mu *sync.RWMutex) {
	mu.Lock()
	defer mu.Unlock()
	ages[name] = age
	fmt.Printf("%s's age is set to %d\n", name, ages[name])
}

func main() {

	ages := map[string]int{
		"Ali":    34,
		"Madina": 28,
	}

	var wg sync.WaitGroup
	var mu sync.RWMutex

	wg.Add(4)

	go func() {
		defer wg.Done()
		Reader("Ali", ages, &mu)
	}()
	go func() {
		defer wg.Done()
		Reader("Madina", ages, &mu)
	}()
	go func() {
		defer wg.Done()
		Writer("Madina", 33, ages, &mu)
	}()
	go func() {
		defer wg.Done()
		Writer("Gerda", 32, ages, &mu)
	}()
	wg.Wait()

	fmt.Println(ages)

}
