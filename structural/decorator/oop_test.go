package decorator

import (
	"strings"
	"testing"
)

func Test_OOP(t *testing.T) {
	var notifier Notifier = &EmailNotifier{}
	notifier = &SMSNotifier{Notifier: notifier}
	notifier = &SlackNotifier{Notifier: notifier}

	output := captureOutput(func() {
		notifier.Send("Hello")
	})

	expectedKeywords := []string{
		"[Email]:",
		"[SMS]:",
		"[Slack]:",
	}

	for _, keyword := range expectedKeywords {
		if !strings.Contains(output, keyword) {
			t.Errorf("Expected output to contain %q", keyword)
		}
	}
}
