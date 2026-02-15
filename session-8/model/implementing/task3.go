package implementing

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type CreditCard struct{}

type PayPal struct{}

func (p PayPal) ProcessPayment(amount float64) {
	fmt.Printf("Processing PayPal payment of %.2f\n", amount)
}

func (c CreditCard) ProcessPayment(amount float64) {
	fmt.Printf("Processing credit card payment of %.2f\n", amount)
}

func ProcessPayment(p PaymentProcessor, amount float64) {
	p.ProcessPayment(amount)
}
