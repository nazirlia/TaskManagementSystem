package embedding

type Person struct {
	FirstName string
	LastName  string
}

type Employee struct {
	EmployeeId int
	Position   string
	Person
}
