package decorator

import "fmt"

// Notifier is a Component Interface
type Notifier interface {
	Send(message string)
}

// EmailNotifier is a Concrete Component (always enabled)
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("[Email]: ", message)
}

// SMSNotifier is a Concrete Decorator
type SMSNotifier struct {
	Notifier
}

func (s *SMSNotifier) Send(message string) {
	s.Notifier.Send(message)
	fmt.Println("[SMS]: ", message)
}

// SMSNotifier is another Concrete Decorator
type SlackNotifier struct {
	Notifier
}

func (s *SlackNotifier) Send(message string) {
	s.Notifier.Send(message)
	fmt.Println("[Slack]: ", message)
}
