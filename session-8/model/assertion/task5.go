package assertion

import "fmt"

func TypeAssertion(TypeOf any) {
	if value, ok := TypeOf.(string); ok {
		fmt.Printf("Value is of type string %s\n", value)
	} else if value, ok := TypeOf.(float64); ok {
		fmt.Printf("Value is of type float64 %.2f\n", value)
	} else if value, ok := TypeOf.(int); ok {
		fmt.Printf("Value is of type int %d\n", value)
	} else {
		fmt.Printf("Value is of type %T\n", TypeOf)
	}
}
