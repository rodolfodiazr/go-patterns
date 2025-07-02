package mediator

import "fmt"

// Mediator defines the interface for communication
type Mediator interface {
	Send(user User, message string)
	Register(user User)
}

// User defines the interface for users
type User interface {
	Name() string
	Send(message string)
	Receive(sender string, message string)
	SetChat(chat Mediator)
}

// ChatRoom is the concrete mediator
type ChatRoom struct {
	users map[string]User
}

// NewChatRoom creates a new chat room
func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users: make(map[string]User),
	}
}

// Register adds a user to the chat room
func (c *ChatRoom) Register(user User) {
	fmt.Printf("[System] %s has joined the chat.\n", user.Name())
	c.users[user.Name()] = user
	user.SetChat(c)
}

// Send sends a message from one user to all others
func (c *ChatRoom) Send(sender User, message string) {
	for name, user := range c.users {
		if name != sender.Name() {
			user.Receive(sender.Name(), message)
		}
	}
}

// ChatUser is a participant in the chat room
type ChatUser struct {
	name string
	chat Mediator
}

// NewUser creates a new user
func NewUser(name string) *ChatUser {
	return &ChatUser{name: name}
}

// Send sends a message via the mediator
func (u *ChatUser) Send(message string) {
	u.chat.Send(u, message)
}

// Receive receives a message from the chat room
func (u *ChatUser) Receive(sender, message string) {
	fmt.Printf("[%s receives from %s]: %s\n", u.name, sender, message)
}

func (u *ChatUser) Name() string {
	return u.name
}

func (u *ChatUser) SetChat(chat Mediator) {
	u.chat = chat
}

// Run demonstrates the Mediator pattern
func Run() {
	chat := NewChatRoom()

	// Create new users
	alice := NewUser("Alice")
	becky := NewUser("Becky")
	charles := NewUser("Charles")

	// Add the created users to the chat
	chat.Register(alice)
	chat.Register(becky)
	chat.Register(charles)

	// Send messages by chat
	alice.Send("Hello everyone!")
	becky.Send("Hi Alice!")
}
