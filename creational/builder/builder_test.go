package builder

import "testing"

func Test_UserBuilder_New(t *testing.T) {
	name := "John"
	emailAddress := "jsmith@email.com"
	usr := NewUserBuilder(name, emailAddress).
		Build()

	if usr.Name != name {
		t.Errorf("expected name '%q', got %q", name, usr.Name)
	}

	if usr.EmailAddress != emailAddress {
		t.Errorf("expected email '%q', got %q", emailAddress, usr.EmailAddress)
	}

	if usr.PhoneNumber != "" {
		t.Errorf("expected empty phone number, got %q", usr.PhoneNumber)
	}

	if usr.IsActive {
		t.Error("expected user to be inactive by default")
	}
}

func Test_UserBuilder_WithPhone(t *testing.T) {
	name := "Bob"
	emailAddress := "Martinez"
	phoneNumber := "(123) 456-7890 "
	usr := NewUserBuilder(name, emailAddress).
		WithPhone(phoneNumber).
		Build()

	if usr.PhoneNumber != phoneNumber {
		t.Errorf("expected phone '%q', got %q", phoneNumber, usr.PhoneNumber)
	}
}

func TestUserBuilder_Activate(t *testing.T) {
	name := "Charlie"
	emailAddress := "ccharleston@email.com"
	usr := NewUserBuilder(name, emailAddress).
		Activate().
		Build()

	if !usr.IsActive {
		t.Error("expected user to be active")
	}
}
