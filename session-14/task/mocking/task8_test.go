package mocking

import "testing"

type MockNotifier struct {
	mockSend func(message string) error
}

func (m MockNotifier) Send(message string) error {
	return m.mockSend(message)
}

func TestSendNotification(t *testing.T) {
	result := MockNotifier{
		mockSend: func(message string) error {
			return nil
		},
	}
	err := result.mockSend("hello world")
	if err != nil {
		t.Errorf("SendNotification returned unexpected error: %v", err)
	}
}

func TestSendNotificationEmptyMessage(t *testing.T) {
	notifier := MockNotifier{
		mockSend: func(message string) error {
			return nil
		},
	}
	err := SendNotification(notifier, "")
	if err == nil {
		t.Errorf("SendNotification expected error for empty message, got nil")
	}
}
