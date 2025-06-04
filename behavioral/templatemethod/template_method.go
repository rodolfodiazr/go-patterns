package templatemethod

import "fmt"

// Exporter defines the template method
type Exporter interface {
	Export()
	prepareData() []string
	formatData(data []string) string
	writeToFile(formatted string)
}

// BaseExporter provides the template logic
type BaseExporter struct{}

func (b *BaseExporter) Export(e Exporter) {
	data := e.prepareData()
	formatted := e.formatData(data)
	e.writeToFile(formatted)
}

// CSVExporter is a concrete implementation
type CSVExporter struct {
	*BaseExporter
}

func (e *CSVExporter) Export() {
	e.BaseExporter.Export(e)
}

func (e *CSVExporter) prepareData() []string {
	return []string{"black", "white", "gray"}
}

func (e *CSVExporter) formatData(data []string) string {
	var output string
	for i, item := range data {
		output += item
		if i < len(data)-1 {
			output += ", "
		}
	}
	return output
}

func (e *CSVExporter) writeToFile(content string) {
	fmt.Println("Writing CSV:", content)
}

// JSONExporter is another concrete implementation
type JSONExporter struct {
	*BaseExporter
}

func (e *JSONExporter) Export() {
	e.BaseExporter.Export(e)
}

func (e *JSONExporter) prepareData() []string {
	return []string{"black", "white", "gray"}
}

func (e *JSONExporter) formatData(data []string) string {
	var output string
	for i, item := range data {
		output += item
		if i < len(data)-1 {
			output += `", "`
		}
	}
	return fmt.Sprintf(`["%s"]`, output)
}

func (e *JSONExporter) writeToFile(content string) {
	fmt.Println("Writing JSON:", content)
}

// Run demonstrates the Template Method pattern
func Run() {
	csv := &CSVExporter{&BaseExporter{}}
	csv.BaseExporter.Export(csv)

	json := &JSONExporter{&BaseExporter{}}
	json.BaseExporter.Export(json)
}
