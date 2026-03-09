package table_driven

import (
	"errors"
	"fmt"
)

func ConvertTemperature(value float64, scale string) (float64, error) {
	switch scale {
	case "C":
		return value*9/5 + 32, nil
	case "F":
		return (value - 32) * 5 / 9, nil
	default:
		return 0, errors.New(fmt.Sprintf("unsupported scale: %s", scale))
	}
}
