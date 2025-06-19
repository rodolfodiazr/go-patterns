package factorymethod

import "fmt"

// Notifier is the product interface
type Notifier interface {
	Send(message string)
}

// EmailNotifier is one concrete product
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Print("[Email]: ", message)
}

// SMSNotifier is another concrete product
type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) {
	fmt.Print("[SMS]: ", message)
}

// NotifierType defines types of notifiers
type NotifierType string

const (
	EmailNotifierType NotifierType = "email"
	SMSNotifierType   NotifierType = "sms"
)

// CreateNotifier creates a new notifier. This is the factory method.
func CreateNotifier(t NotifierType) (Notifier, error) {
	switch t {
	case EmailNotifierType:
		return &EmailNotifier{}, nil
	case SMSNotifierType:
		return &SMSNotifier{}, nil
	default:
		return nil, fmt.Errorf("unknown notifier type: %s", t)
	}
}

// Run demonstrates the Factory Method Pattern
func Run() {
	emailNotifier, _ := CreateNotifier(EmailNotifierType)
	smsNotifier, _ := CreateNotifier(SMSNotifierType)

	emailNotifier.Send("Message sent via email.")
	smsNotifier.Send("Message sent via SMS.")
}
