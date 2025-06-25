package templatemethod

import (
	"strings"
	"testing"
)

type MockExporter struct {
	*BaseExporter
	log         []string
	fileContent string
}

func (m *MockExporter) Export() {
	m.BaseExporter.Export(m)
}

func (m *MockExporter) prepareData() []string {
	m.log = append(m.log, "prepareData")
	return []string{"a", "b", "c"}
}

func (m *MockExporter) formatData(data []string) string {
	m.log = append(m.log, "formatData")
	return strings.Join(data, "-")
}

func (m *MockExporter) writeToFile(content string) {
	m.log = append(m.log, "writeToFile")
	m.fileContent = content
}

func Test_TemplateMethod(t *testing.T) {
	m := &MockExporter{BaseExporter: &BaseExporter{}}
	m.Export()

	expectedLogs := []string{"prepareData", "formatData", "writeToFile"}
	for i, expectedLog := range expectedLogs {
		if m.log[i] != expectedLog {
			t.Errorf("Expected log %q at index %d, got %q", expectedLog, i, m.log[i])
		}
	}

	if m.fileContent != "a-b-c" {
		t.Errorf("Expected formatted data to be 'a-b-c', got %q", m.fileContent)
	}
}
