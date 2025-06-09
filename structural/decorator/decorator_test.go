package decorator

import (
	"bytes"
	"os"
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
