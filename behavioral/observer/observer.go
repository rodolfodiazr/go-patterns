package observer

import "fmt"

// Observer defines a subscriber that responds to events.
type Observer interface {
	OnEvent(event string, payload any)
}

// EventManager handles event subscription and publishing.
type EventManager struct {
	subscribers map[string][]Observer
}

// NewEventManager initializes a new event manager.
func NewEventManager() *EventManager {
	return &EventManager{
		subscribers: make(map[string][]Observer),
	}
}

// Subscribe adds an observer to a specific event.
func (em *EventManager) Subscribe(e string, o Observer) {
	em.subscribers[e] = append(em.subscribers[e], o)
}

// Publish triggers all observers subscribed to the event.
func (em *EventManager) Publish(e string, p any) {
	if observers, ok := em.subscribers[e]; ok {
		for _, o := range observers {
			o.OnEvent(e, p)
		}
	}
}

// EmailNotifier is an observer that sends email notifications.
type EmailNotifier struct {
	EmailAddress string
}

func (e *EmailNotifier) OnEvent(event string, payload any) {
	switch event {
	case "product:in_stock":
		if product, ok := payload.(string); ok {
			fmt.Printf("[Email to %s] Event: %s -> %s is available!\n", e.EmailAddress, event, product)
		}
	}
}

// SlackNotifier is an observer that sends Slack notifications.
type SlackNotifier struct {
	Channel string
}

func (e *SlackNotifier) OnEvent(event string, payload any) {
	switch event {
	case "product:in_stock":
		if product, ok := payload.(string); ok {
			fmt.Printf("[Slack #%s] Event: %s -> %s is available!\n", e.Channel, event, product)
		}
	}
}

// Run demonstrates the Observer pattern
func Run() {
	eventManager := NewEventManager()

	email := &EmailNotifier{
		EmailAddress: "user@email.com",
	}

	slack := &SlackNotifier{
		Channel: "stock-updates",
	}

	eventManager.Subscribe("product:in_stock", email)
	eventManager.Subscribe("product:in_stock", slack)

	eventManager.Publish("product:in_stock", "Product")
}
