package prototype

import "fmt"

// Template represents a cloneable document template (Prototype Pattern)
type Template interface {
	Clone() Template
	Print()
}

// Document is a concrete type that implements Template
type Document struct {
	Title   string
	Content string
	Author  string
}

func (d *Document) Clone() Template {
	return &Document{
		Title:   d.Title,
		Content: d.Content,
		Author:  d.Author,
	}
}

func (d *Document) Print() {
	fmt.Printf("Document -> Title: %s, Content: %s, Author: %s\n", d.Title, d.Content, d.Author)
}

// Run demonstrates the Prototype pattern
func Run() {
	var t Template = &Document{
		Title:   "Monthly Report Template",
		Content: "This is a base content for reports.",
		Author:  "System",
	}

	report1 := t.Clone().(*Document)
	report1.Title = "April Financial Report"
	report1.Author = "John Smith"
	report1.Print()

	report2 := t.Clone().(*Document)
	report2.Title = "March Financial Report"
	report2.Author = "Peter Petersen"
	report2.Print()
}
