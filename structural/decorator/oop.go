package decorator

import "fmt"

// Notifier is a component interface
type Notifier interface {
	Send(message string)
}

// EmailNotifier is a concrete component (always enabled)
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("[Email]: ", message)
}

// SMSNotifier is a concrete decorator
type SMSNotifier struct {
	Notifier
}

func (s *SMSNotifier) Send(message string) {
	s.Notifier.Send(message)
	fmt.Println("[SMS]: ", message)
}

// SMSNotifier is another concrete decorator
type SlackNotifier struct {
	Notifier
}

func (s *SlackNotifier) Send(message string) {
	s.Notifier.Send(message)
	fmt.Println("[Slack]: ", message)
}
