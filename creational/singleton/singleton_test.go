package singleton

import "testing"

func Test_GetConfig(t *testing.T) {
	cfg1 := GetConfig()
	cfg2 := GetConfig()

	if cfg1 != cfg2 {
		t.Errorf("Expected the same instance, got different ones.")
	}

	newAppName := "New Sample App"
	cfg1.AppName = newAppName

	if cfg2.AppName != newAppName {
		t.Errorf("Expected AppName to be '%s', got '%s'", newAppName, cfg2.AppName)
	}
}
