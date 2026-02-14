package embedding

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicle
	FuelType string
}
