package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

func main() {
	s1 := Student{
		Name:  "Ali",
		Age:   20,
		Grade: 42,
	}
	s2 := Student{
		Name:  "Jack",
		Age:   20,
		Grade: 52,
	}
	s3 := Student{
		Name:  "David",
		Age:   21,
		Grade: 72,
	}

	var students = []Student{s1, s2, s3}
	studentsJson, _ := json.Marshal(students)
	err := os.WriteFile("students.json", studentsJson, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// Part 2
	file, err := os.ReadFile("students.json")
	if err != nil {
		fmt.Println(err)
	}
	var data []Student
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Students with passing grades:")
	for _, stu := range data {
		if stu.Grade >= 60 {
			fmt.Printf("- %s", stu.Name)
		}
	}

}
