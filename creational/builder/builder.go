package builder

import "fmt"

// User is a complex object
type User struct {
	Name         string
	EmailAddress string
	PhoneNumber  string
	IsActive     bool
}

// UserBuilder provides a fluent API to build a User
type UserBuilder struct {
	user *User
}

// NewUserBuilder creates a new UserBuilder.
func NewUserBuilder(name, email string) *UserBuilder {
	return &UserBuilder{
		user: &User{
			Name:         name,
			EmailAddress: email,
		},
	}
}

// WithPhone sets a phone number for the user.
func (b *UserBuilder) WithPhone(phone string) *UserBuilder {
	b.user.PhoneNumber = phone
	return b
}

// Activate marks the user as active.
func (b *UserBuilder) Activate() *UserBuilder {
	b.user.IsActive = true
	return b
}

// Build finalizes and returns the constructed User object.
func (b *UserBuilder) Build() User {
	return *b.user
}

// Run demonstrates the Builder pattern
func Run() {
	user := NewUserBuilder("John Smith", "jsmith@email.com").
		WithPhone("123-456-7890").
		Activate().
		Build()

	fmt.Printf("User built: %+v\n", user)
}
