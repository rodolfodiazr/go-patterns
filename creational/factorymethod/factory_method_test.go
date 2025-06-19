package factorymethod

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	// Save the original stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	f()

	// Close the writer and restore stdout
	_ = w.Close()
	os.Stdout = old

	// Read the output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func Test_Notifier_CreateNotifier(t *testing.T) {
	tCases := []struct {
		name                 string
		notifierType         NotifierType
		expectedNotifierType string
		expectedError        error
	}{
		{
			name:          "Unknow notifier type.",
			notifierType:  "push",
			expectedError: errors.New("unknown notifier type: push"),
		},
		{
			name:                 "Valid notifier type: Email.",
			notifierType:         EmailNotifierType,
			expectedNotifierType: "*factorymethod.EmailNotifier",
			expectedError:        nil,
		},
		{
			name:                 "Valid notifier type: SMS.",
			notifierType:         SMSNotifierType,
			expectedNotifierType: "*factorymethod.SMSNotifier",
			expectedError:        nil,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			notifier, err := CreateNotifier(tc.notifierType)

			if tc.expectedError == nil && err != nil {
				t.Fatalf("Expected no error, got: %v", err)
			}

			if tc.expectedError != nil && err == nil {
				t.Fatalf("Expected error %q, got: nil", tc.expectedError.Error())
			}

			if tc.expectedError != nil && err != nil && tc.expectedError.Error() != err.Error() {
				t.Fatalf("Expected error %q, got: %q", tc.expectedError.Error(), err.Error())
			}

			if tc.expectedError == nil {
				notifierType := fmt.Sprintf("%T", notifier)
				if tc.expectedNotifierType != notifierType {
					t.Errorf("Expected type %q, got: %q", tc.expectedNotifierType, notifierType)
				}
			}
		})
	}
}

func Test_Notifier_Send(t *testing.T) {
	tCases := []struct {
		name            string
		notifier        Notifier
		expectedMessage string
	}{
		{
			name:            "Sending an email message",
			notifier:        &EmailNotifier{},
			expectedMessage: "[Email]: Sending a new message.",
		},
		{
			name:            "Sending an SMS message",
			notifier:        &SMSNotifier{},
			expectedMessage: "[SMS]: Sending a new message.",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			output := captureOutput(func() {
				tc.notifier.Send("Sending a new message.")
			})

			if output != tc.expectedMessage {
				t.Errorf("Expected message:\n%q\ngot:\n%q", tc.expectedMessage, output)
			}
		})
	}
}
