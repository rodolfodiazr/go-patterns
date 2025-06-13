package adapter

import (
	"strings"
	"testing"
)

type MockExternalLogger struct {
	Message string
}

func (m *MockExternalLogger) Info(msg string, fields map[string]string) {
	var b strings.Builder

	b.WriteString("[ExternalLogger]: ")
	b.WriteString(msg)
	b.WriteString(". ")

	if len(fields) > 0 {
		b.WriteString("Data: ")
		for k, v := range fields {
			b.WriteString("'")
			b.WriteString(k)
			b.WriteString("': '")
			b.WriteString(v)
			b.WriteString("'")
		}
	}
	m.Message = b.String()
}

func Test_ExternalLoggerAdapter(t *testing.T) {
	mockLogger := &MockExternalLogger{}
	expectedMessage := "[ExternalLogger]: This is a new message. Data: 'source': 'app'"

	logger := NewThirdPartyLoggerAdapter(mockLogger)
	logger.Info("This is a new message")

	if mockLogger.Message != expectedMessage {
		t.Errorf("Expected %q but got %q", expectedMessage, mockLogger.Message)
	}
}
