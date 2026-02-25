package implementing

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct{}

type SMSNotifier struct{}

func (EmailNotifier) Notify(message string) {
	fmt.Printf("Sending email notification: %s\n", message)
}

func (SMSNotifier) Notify(message string) {
	fmt.Printf("Sending SMS notification: %s\n", message)
}

func Notify(notifier Notifier, message string) {
	notifier.Notify(message)
}
