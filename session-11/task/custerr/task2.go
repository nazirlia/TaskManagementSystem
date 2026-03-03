package custerr

import "fmt"

type DivideError struct {
	A float64
	B float64
}

func (e DivideError) Error() string {
	return fmt.Sprintf("Cannot divide %.2f by %.2f.", e.A, e.B)
}

func DivideBy(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, DivideError{
			A: a,
			B: b,
		}
	}
	return a / b, nil
}
