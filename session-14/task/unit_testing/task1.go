package unit_testing

import "fmt"

func Calculate(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero is not allowed")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", operation)
	}
}
