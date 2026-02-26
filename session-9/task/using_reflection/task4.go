package using_reflection

import (
	"fmt"
	"reflect"
)

func (p *Person) Greeting() string {
	return fmt.Sprintf("Hi, I'm %s", p.Name)
}

func InvokeMethod(value interface{}, methodName string) {
	fmt.Printf("Invoke Method: %s, Result: %s\n", methodName, reflect.ValueOf(value).MethodByName("Greeting").Call(nil))
}
