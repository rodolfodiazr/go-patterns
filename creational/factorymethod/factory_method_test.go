package factorymethod

import (
	"errors"
	"fmt"
	"testing"
)

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
