package strategy

import "fmt"

// PaymentMethod defines the strategy interface
type PaymentMethod interface {
	Pay(amount float64)
}

// CreditCardPayment is a concrete strategy
type CreditCardPayment struct {
	Name       string
	CardNumber string
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card [%s]\n", amount, c.CardNumber)
}

// PayPalPayment is another concrete strategy
type PayPalPayment struct {
	EmailAddress string
}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal [%s]\n", amount, p.EmailAddress)
}

// BitcoinPayment is yet another concrete strategy
type BitcoinPayment struct {
	WalletAddress string
}

func (b *BitcoinPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Bitcoin Wallet [%s]\n", amount, b.WalletAddress)
}

// ShoppingCart is the context that uses a payment strategy
type ShoppingCart struct {
	Amount  float64
	Payment PaymentMethod
}

func (s *ShoppingCart) Checkout() {
	s.Payment.Pay(s.Amount)
}

// Run demonstrates the Strategy pattern
func Run() {
	cart := &ShoppingCart{
		Amount: 150.00,
	}

	cart.Payment = &CreditCardPayment{Name: "John Smith", CardNumber: "1234-5678-9876-5432"}
	cart.Checkout()

	cart.Payment = &PayPalPayment{EmailAddress: "jsmith@email.com"}
	cart.Checkout()

	cart.Payment = &BitcoinPayment{WalletAddress: "1A2B3C4D5E6F"}
	cart.Checkout()
}
