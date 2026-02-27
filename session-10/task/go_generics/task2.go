package go_generics

func MinMax[T int | float64](slice []T) (T, T) {
	var mi T
	var ma T
	mi, ma = slice[0], slice[0]
	for _, x := range slice {
		if x < mi {
			mi = x
		}
		if x > ma {
			ma = x
		}
	}
	return mi, ma
}

/// 100, 8, 20, 5, 10, 15
