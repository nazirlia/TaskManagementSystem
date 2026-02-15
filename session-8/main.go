package main

import (
	"fmt"
	"session-8/model/assertion"
	"session-8/model/defining"
	"session-8/model/embedding"
	"session-8/model/implementing"
	"session-8/model/switching"
)

func main() {
	
	//Task 1
	r := defining.Rectangle{
		Width:  10,
		Height: 20,
	}
	
	c := defining.Circle{
		Radius: 10,
	}
	
	fmt.Printf("Circle Area: %.2f\n", defining.AreaOfShapes(c))
	fmt.Printf("Rectangle Area: %.2f\n", defining.AreaOfShapes(r))
	// End of Task 1
	
	// Task2
	p := defining.Person{
		Audio: "Hello!",
	}
	
	d := defining.Dog{
		Audio: "Woof!",
	}
	
	fmt.Printf("Dog says: %s\n", defining.Speak(d))
	fmt.Printf("Person says: %s\n", defining.Speak(p))
	// End of Task 2
	
	// Task 3
	cc := implementing.CreditCard{}
	
	pp := implementing.PayPal{}
	
	implementing.ProcessPayment(cc, 100)
	implementing.ProcessPayment(pp, 22.9)
	// End of Task 3
	
	// Task 4
	
	sn := implementing.SMSNotifier{}
	en := implementing.EmailNotifier{}
	
	implementing.Notify(sn, "Your item has shipped")
	implementing.Notify(en, "Your item has been shipped")
	
	// End of Task 4
	
	// Task 5
	var anyTypeTask5 any
	anyTypeTask5 = 45
	assertion.TypeAssertion(anyTypeTask5)
	anyTypeTask5 = "Adam and Eve"
	assertion.TypeAssertion(anyTypeTask5)
	// End of Task 5
	
	// Task 6
	
	var anyType any
	anyType = 5
	switching.PrintBasedOnType(anyType)
	anyType = "Game of Thrones"
	switching.PrintBasedOnType(anyType)
	anyType = implementing.SMSNotifier{}
	switching.PrintBasedOnType(anyType)
	
	// End of Task 6
	
	// Task 7
	ab := embedding.AutoBot{}
	
	ab.Say()
	ab.Move()
	// End of Task 7
	
}
