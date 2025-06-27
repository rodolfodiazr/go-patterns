package state

import (
	"testing"
)

type MockState struct {
	edited    bool
	published bool
}

func (m *MockState) Edit(content string) {
	m.edited = true
}

func (m *MockState) Publish() {
	m.published = true
}

func Test_Document_DelegatesToStateMethods(t *testing.T) {
	state := &MockState{}

	if state.edited || state.published {
		t.Error("Neither Edit() nor Publish() should have been called yet")
	}

	doc := &Document{}
	doc.SetState(state)
	doc.Edit("Random content")
	doc.Publish()

	if !state.edited {
		t.Errorf("Expected Edit() to be called, but got edited = %v", state.edited)
	}

	if !state.published {
		t.Errorf("Expected Publish() to be called, but got published = %v", state.published)
	}
}
