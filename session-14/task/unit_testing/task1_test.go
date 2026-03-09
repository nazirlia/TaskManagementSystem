package unit_testing

import (
	"testing"
)

func TestCalculateAddition(t *testing.T) {
	result, err := Calculate(10, 5, "+")
	if err != nil {
		t.Errorf("Calculate error: %s", err)
	}
	if result != 15 {
		t.Errorf("Calculate: want %d, got %v", 15, result)
	}
}

func TestCalculateSubtraction(t *testing.T) {
	result, err := Calculate(10, 5, "-")
	if err != nil {
		t.Errorf("Calculate error: %s", err)
	}
	if result != 5 {
		t.Errorf("Calculate: want %d, got %v", 5, result)
	}
}

func TestCalculateMultiplication(t *testing.T) {
	result, err := Calculate(10, 5, "*")
	if err != nil {
		t.Errorf("Calculate error: %s", err)
	}
	if result != 50 {
		t.Errorf("Calculate: want %d, got %v", 50, result)
	}
}

func TestCalculateDivision(t *testing.T) {
	result, err := Calculate(10, 5, "/")
	if err != nil {
		t.Errorf("Calculate error: %s", err)
	}
	if result != 2 {
		t.Errorf("Calculate: want %d, got %v", 2, result)
	}
}

func TestCalculateDivisionByZero(t *testing.T) {
	result, err := Calculate(10, 0, "/")
	if err == nil {
		t.Errorf("Expected error but succeed")
	}
	if err.Error() != "division by zero is not allowed" {
		t.Fatalf("unexpected error message: %v", err)
	}
	if result != 0 {
		t.Errorf("Calculate: want %d, got %v", 0, result)
	}
}

func TestCalculateUnknownOperation(t *testing.T) {
	_, err := Calculate(10, 5, "%")
	if err == nil {
		t.Errorf("Expected error but succeed")
	}
	if err.Error() != "unsupported operation: %" {
		t.Errorf("unexpected error message: %v", err)
	}
}
