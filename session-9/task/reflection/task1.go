package reflection

import (
	"fmt"
	"reflect"
)

func IdentifyTypeAndKind(value interface{}) {
	fmt.Printf("Value: %v, Type: %v, Kind: %v\n", reflect.ValueOf(value), reflect.TypeOf(value), reflect.ValueOf(value).Kind())
}
