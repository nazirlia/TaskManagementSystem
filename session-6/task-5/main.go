package main

import "fmt"

func main() {
	var s []string
	s = []string{"A", "B", "A", "D", "E", "F", "E"}
	countFrequencyOfWords(s)
}

func countFrequencyOfWords(s []string) {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++

	}
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}
