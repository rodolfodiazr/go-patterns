package strategy

import (
	"testing"
)

type MockPayment struct {
	called bool
	amount float64
}

func (m *MockPayment) Pay(amount float64) {
	m.amount = amount
	m.called = true
}

func Test_PaymentMethod_Checkout(t *testing.T) {
	m := &MockPayment{}
	cart := &ShoppingCart{
		Amount:  55.00,
		Payment: m,
	}

	cart.Checkout()

	if !m.called {
		t.Errorf("Expected Pay Method to be called")
	}

	if m.amount != 50.00 {
		t.Errorf("Expected amount = 50.00, got %v", m.amount)
	}
}
