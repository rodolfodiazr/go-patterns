package mediator

import "fmt"

// Mediator defines the interface for communication
type Mediator interface {
	Send(sender, message string)
	Register(user *User)
}

// ChatRoom is the concrete mediator
type ChatRoom struct {
	users map[string]*User
}

// NewChatRoom creates a new chat room
func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users: make(map[string]*User),
	}
}

// Register adds a user to the chat room
func (c *ChatRoom) Register(user *User) {
	c.users[user.name] = user
	user.chat = c
}

// Send sends a message from one user to all others
func (c *ChatRoom) Send(sender string, message string) {
	for name, user := range c.users {
		if name != sender {
			user.Receive(sender, message)
		}
	}
}

// User is a participant in the chat room
type User struct {
	name string
	chat Mediator
}

// NewUser creates a new user
func NewUser(name string) *User {
	return &User{name: name}
}

// Send sends a message via the mediator
func (u *User) Send(message string) {
	fmt.Printf("[%s sends]: %s\n", u.name, message)
	u.chat.Send(u.name, message)
}

// Receive receives a message from the chat room
func (u *User) Receive(sender, message string) {
	fmt.Printf("[%s receives from %s]: %s\n", u.name, sender, message)
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
