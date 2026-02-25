package switching

import "fmt"

func PrintBasedOnType(typeOf any) {
	switch t := typeOf.(type) {
	case string:
		fmt.Println("Value is of type string:", t)
	case int:
		fmt.Println("Value is of type int:", t)
	case float64:
		fmt.Println("Value is of type float64:", t)
	case bool:
		fmt.Println("Value is of type bool:", t)
	default:
		fmt.Println("Value is of type unknown:", t)
	}

}
