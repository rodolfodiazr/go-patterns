package prototype

import (
	"testing"
)

func Test_Prototype_Clone(t *testing.T) {
	document := &Document{
		Title:   "Base Title",
		Content: "Base Content",
		Author:  "Base Author",
	}

	clone1 := document.Clone().(*Document)

	if clone1 == document {
		t.Errorf("Expected clone1 to be a new instance, but it's the same.")
	}

	if clone1.Title != document.Title ||
		clone1.Content != document.Content ||
		clone1.Author != document.Author {
		t.Errorf("Expected clone1 fields to match original document.")
	}

	clone2 := document.Clone().(*Document)
	clone2.Title = "[Clone #2] Base Title"

	if clone1.Title == clone2.Title {
		t.Errorf("Expected clone1 and clone2 titles to be different.")
	}
}
