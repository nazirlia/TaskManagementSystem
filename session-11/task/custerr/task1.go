package custerr

type DivisionError struct{}

func (e DivisionError) Error() string {
	return "Division by zero is not allowed."
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivisionError{}
	}
	return a / b, nil
}
