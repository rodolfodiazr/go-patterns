package mediator

import "testing"

type MockUser struct {
	name string
	chat Mediator

	receivedMessages []string
}

func NewMockUser(name string) *MockUser {
	return &MockUser{name: name}
}

func (u *MockUser) Send(message string) {
	u.chat.Send(u, message)
}

func (u *MockUser) Receive(sender, message string) {
	u.receivedMessages = append(u.receivedMessages, "["+sender+"]: "+message)
}

func (u *MockUser) Name() string {
	return u.name
}

func (u *MockUser) SetChat(chat Mediator) {
	u.chat = chat
}

func Test_ChatRoom_MessageBroadcast(t *testing.T) {
	// Create a new chat
	chat := NewChatRoom()

	// Create new users
	alice := NewMockUser("Alice")
	bob := NewMockUser("Bob")
	eve := NewMockUser("Eve")

	// Add the created users to the chat
	chat.Register(alice)
	chat.Register(bob)
	chat.Register(eve)

	// Send messages by chat
	alice.Send("Hello everyone!")

	expectedMessage := "[Alice]: Hello everyone!"
	if len(bob.receivedMessages) != 1 || bob.receivedMessages[0] != expectedMessage {
		t.Errorf("Bob should have received %q, got %v", expectedMessage, bob.receivedMessages)
	}

	if len(eve.receivedMessages) != 1 || eve.receivedMessages[0] != expectedMessage {
		t.Errorf("Eve should have received %q, got %v", expectedMessage, eve.receivedMessages)
	}

	if len(alice.receivedMessages) != 0 {
		t.Errorf("Alice should not receive her own message")
	}
}
