package reflection

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func InspectPerson(value interface{}) {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field Name: %s, Type: %v, Value: %v\n", field.Name, field.Type, v.FieldByName(field.Name))
	}
}
