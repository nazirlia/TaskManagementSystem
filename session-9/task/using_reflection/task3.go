package using_reflection

import "reflect"

type Person struct {
	Name string
	Age  int
	City string
}

func SetFieldValue(value interface{}, fieldName string, newValue interface{}) {
	reflect.ValueOf(value).Elem().FieldByName(fieldName).Set(reflect.ValueOf(newValue))
}
