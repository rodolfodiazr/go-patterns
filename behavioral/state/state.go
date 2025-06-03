package state

import "fmt"

// State defines the interface for document states
type State interface {
	Edit(content string)
	Publish()
}

// Document is the context that changes behavior based on state
type Document struct {
	content string
	state   State
}

// NewDocument creates a new document
func NewDocument() *Document {
	document := &Document{}
	document.state = &DraftState{document}
	return document
}

func (d *Document) SetState(state State) {
	d.state = state
}

func (d *Document) Edit(content string) {
	d.state.Edit(content)
}

func (d *Document) Publish() {
	d.state.Publish()
}

// DraftState allows editing and moving to moderation
type DraftState struct {
	doc *Document
}

func (s *DraftState) Edit(content string) {
	s.doc.content = content
	fmt.Println("Document updated.")
}

func (s *DraftState) Publish() {
	s.doc.SetState(&ModerationState{s.doc})
	fmt.Println("Document sent for review.")
}

// ModerationState allows publishing, not editing
type ModerationState struct {
	doc *Document
}

func (s *ModerationState) Edit(content string) {
	fmt.Println("Cannot edit. Document is under moderation.")
}

func (s *ModerationState) Publish() {
	s.doc.SetState(&PublishedState{s.doc})
	fmt.Println("Document approved and published.")
}

// PublishedState is read-only
type PublishedState struct {
	doc *Document
}

func (s *PublishedState) Edit(content string) {
	fmt.Println("Cannot edit. Document is already published.")
}

func (s *PublishedState) Publish() {
	fmt.Println("Document is already published.")
}

func Run() {
	doc := NewDocument()

	fmt.Println("Document (Draft):")
	doc.Edit("Initial draft content.")
	doc.Publish() // Move to Moderation

	fmt.Println("Document (In Review):")
	doc.Edit("Trying to edit under moderation.")
	doc.Publish() // Move to Published

	fmt.Println("Document (Published):")
	doc.Edit("Trying to edit after publication.")
	doc.Publish() // It is already published
}
