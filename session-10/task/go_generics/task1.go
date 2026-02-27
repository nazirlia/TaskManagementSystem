package go_generics

func Sum[T int | float64](slice []T) T {
	var sum T
	for _, x := range slice {
		sum += x
	}
	return sum
}
