package decorator

import (
	"strings"
	"testing"
)

func Test_Func(t *testing.T) {
	base := func(price float64) float64 {
		return price * 1.5
	}

	pc := Loggin(base)

	output := captureOutput(func() {
		calculatedPrice := pc(100)
		if calculatedPrice != 150 {
			t.Errorf("Expected 150, got %v", calculatedPrice)
		}
	})

	if !strings.Contains(output, "Starting price calculation") ||
		!strings.Contains(output, "Ending price calculation") {
		t.Error("Expected logging output not found")
	}
}
