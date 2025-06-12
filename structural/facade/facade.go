package facade

import "fmt"

type Validator interface {
	Validate(cardNumber string) bool
}

// CardValidator is a subsystem
type CardValidator struct{}

func (v *CardValidator) Validate(cardNumber string) bool {
	fmt.Println("Validating card:", cardNumber)
	return cardNumber != ""
}

type Gateway interface {
	Charge(cardNumber string, amount float64)
}

// PaymentGateway is a subsystem
type PaymentGateway struct{}

func (pg *PaymentGateway) Charge(cardNumber string, amount float64) {
	fmt.Printf("Charging %.2f to card %s\n", amount, cardNumber)
}

type Notifier interface {
	SendReceipt(emailAddr string, amount float64)
}

// NotificationService is a subsystem
type NotificationService struct{}

func (n *NotificationService) SendReceipt(emailAddress string, amount float64) {
	fmt.Printf("Sending receipt to %s for amount $%.2f\n", emailAddress, amount)
}

type Logger interface {
	Record(transaction string)
}

// AuditLog is a subsystem
type AuditLog struct{}

func (a *AuditLog) Record(transaction string) {
	fmt.Println("Audit log:", transaction)
}

// PaymentProcessor is a facade
type PaymentProcessor struct {
	validator Validator
	gateway   Gateway
	notifier  Notifier
	logger    Logger
}

// NewPaymentProcessor creates a new PaymentProcessor
func NewPaymentProcessor(v Validator, g Gateway, n Notifier, l Logger) *PaymentProcessor {
	return &PaymentProcessor{
		validator: v,
		gateway:   g,
		notifier:  n,
		logger:    l,
	}
}

func (p *PaymentProcessor) Process(card string, emailAddress string, amount float64) {
	if !p.validator.Validate(card) {
		fmt.Println("Invalid card")
		return
	}

	p.gateway.Charge(card, amount)
	p.notifier.SendReceipt(emailAddress, amount)
	p.logger.Record(fmt.Sprintf("Charged %s for $%.2f", card, amount))
}

func Run() {
	processor := NewPaymentProcessor(
		&CardValidator{},
		&PaymentGateway{},
		&NotificationService{},
		&AuditLog{},
	)
	processor.Process("1234-5678-9876-5432", "jsmith@email.com", 150.00)
}
