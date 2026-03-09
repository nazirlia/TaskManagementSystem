package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Employee struct {
	Name     string `xml:"name"`
	Position string `xml:"position"`
	Salary   int    `xml:"salary"`
}

type Employees struct {
	Employees []Employee `xml:"employee"`
}

func main() {
	var employees Employees
	file, err := os.ReadFile("employees.xml")
	if err != nil {
		fmt.Println(err)
	}
	xml.Unmarshal(file, &employees)

	fmt.Println("Employees with Salary above 50000:")
	for _, e := range employees.Employees {
		if e.Salary > 50000 {
			fmt.Printf("- %s %s\n", e.Name, e.Position)
		}
	}
}
