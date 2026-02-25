package main

import (
	"fmt"
	"session-09/task/reflection"
	"session-09/task/using_reflection"
)

func main() {

	reflection.IdentifyTypeAndKind(3)

	p := reflection.Person{
		Name: "Seda Sayan",
		Age:  72,
	}
	reflection.InspectPerson(p)

	p1 := &using_reflection.Person{
		Name: "Bulent",
		Age:  76,
		City: "IST",
	}

	fmt.Printf("%+v\n", p1)
	using_reflection.SetFieldValue(p1, "City", "Ganja")
	fmt.Printf("%+v\n", p1)

	using_reflection.InvokeMethod(p1, "Greeting")
}
