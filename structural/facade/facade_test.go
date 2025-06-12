package facade

import (
	"testing"
)

const (
	mockCard         = "0000 0000 0000 0000"
	mockEmailAddress = "user1@fake.com"
	mockAmount       = 0.0
	mockEntry        = "Charged 0000 0000 0000 0000 for $0.00"
)

type mockValidator struct {
	validCard bool
}

func (m *mockValidator) Validate(card string) bool {
	return m.validCard
}

type mockGateway struct {
	cardCharged string
}

func (m *mockGateway) Charge(card string, amount float64) {
	m.cardCharged = card
}

type mockNotifier struct {
	emailNotified string
}

func (m *mockNotifier) SendReceipt(emailAddr string, amount float64) {
	m.emailNotified = emailAddr
}

type mockLogger struct {
	entryLogged string
}

func (m *mockLogger) Record(entry string) {
	m.entryLogged = entry
}

func Test_PaymentProcessor(t *testing.T) {
	tCases := []struct {
		name              string
		validCard         bool
		expectedCard      string
		expectedEmailAddr string
		expectedEntry     string
	}{
		{
			name:              "trigger all subsystems when card is valid",
			validCard:         true,
			expectedCard:      mockCard,
			expectedEmailAddr: mockEmailAddress,
			expectedEntry:     mockEntry,
		},
		{
			name:              "skip all subsystems when card is not valid",
			validCard:         false,
			expectedCard:      "",
			expectedEmailAddr: "",
			expectedEntry:     "",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := mockValidator{validCard: tc.validCard}
			gateway := mockGateway{}
			notifier := mockNotifier{}
			logger := mockLogger{}

			processor := NewPaymentProcessor(&validator, &gateway, &notifier, &logger)
			processor.Process(mockCard, mockEmailAddress, mockAmount)

			if tc.expectedCard != gateway.cardCharged {
				t.Errorf("expected card = %v, got %v", tc.expectedCard, gateway.cardCharged)
			}

			if tc.expectedEmailAddr != notifier.emailNotified {
				t.Errorf("expected email address = %v, got %v", tc.expectedEmailAddr, notifier.emailNotified)
			}

			if tc.expectedEntry != logger.entryLogged {
				t.Errorf("expected entry = %v, got %v", tc.expectedEntry, logger.entryLogged)
			}
		})
	}
}
