package mocking

import "errors"

type Notifier interface {
	Send(message string) error
}

func SendNotification(notifier Notifier, message string) error {
	if message == "" {
		return errors.New("message cannot be empty")
	}
	return notifier.Send(message)
}
