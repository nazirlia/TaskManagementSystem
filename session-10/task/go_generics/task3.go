package go_generics

func Map[T any, U any](slice []T, fn func(T) U) []U {
	var result []U
	for i := 0; i < len(slice); i++ {
		result = append(result, fn(slice[i]))
	}
	return result
}
