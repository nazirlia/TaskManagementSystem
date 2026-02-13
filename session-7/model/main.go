package main

import (
	"fmt"
	"session-7/embedding"
	"session-7/methods"
	"session-7/structs"
)

func main() {
	var Book structs.Book
	Book = structs.Book{
		Title:  "Ere getmeyin 49 qaydasi",
		Author: "Ali & Nino Book Store",
		Pages:  429,
	}

	fmt.Printf("Book's Title: %s\nBook's Author: %s\nBook's page count: %d\n", Book.Title, Book.Author, Book.Pages)

	var Rectangle structs.Rectangle
	Rectangle = structs.Rectangle{
		Height: 12.5,
		Width:  10.0,
	}

	fmt.Printf(
		"Height: %.2f, Width: %.2f\nArea: %.2f\nPerimeter: %.2f\n",
		Rectangle.Height,
		Rectangle.Width,
		Rectangle.Area(),
		Rectangle.Perimeter(),
	)

	circle := methods.Circle{
		Radius: 10.0,
	}
	fmt.Printf("Circle's Radius: %.2f\nArea: %.2f\n", circle.Radius, circle.Area())

	student1 := methods.Student{
		Name:   "Ali",
		Grades: []int{57, 68, 79, 90, 99},
	}
	student2 := methods.Student{
		Name:   "Vali",
		Grades: []int{56, 67, 99, 90, 99},
	}

	fmt.Printf("Student 1: %s, Average Grade: %.2f\nStudent 2: %s, Average grade: %.2f\n", student1.Name, student1.Average(), student2.Name, student2.Average())

	if student1.Average() > student2.Average() {
		fmt.Printf("%s has a higher grade\n", student1.Name)
	} else {
		fmt.Printf("%s has a higher grade\n", student2.Name)
	}

	Employee := embedding.Employee{
		EmployeeId: 1122,
		Position:   "Senior Vibe Code Cleaner",
		Person: embedding.Person{
			FirstName: "Ilhame",
			LastName:  "Guliyeva",
		},
	}

	fmt.Printf("Name: %s %s\nEmployee ID: %d\nPosition: %s\n", Employee.FirstName, Employee.LastName, Employee.EmployeeId, Employee.Position)

	Car := embedding.Car{
		Vehicle: embedding.Vehicle{
			Make:  "BMW",
			Model: "X5 xDrive",
			Year:  2019,
		},
		FuelType: "Gasoline",
	}
	fmt.Printf("Make: %s\nModel: %s\nYear: %d\n FuelType: %s\n", Car.Make, Car.Model, Car.Year, Car.FuelType)

}
