package addressbook

import (
	"testing"
)

func TestCreateContact(t *testing.T) {
	c := NewContact("Alejandro", "alejandroarpu@gmail.com")

	if c.Name != "Alejandro" {
		t.Error("The contact name doesn't match")
	}

	if c.Email != "alejandroarpu@gmail.com" {
		t.Error("The email doesn't match")
	}
}

func TestAddContact(t *testing.T) {
	d := new(AddressBook)
	c := NewContact("Alejandro", "alejandroarpu@gmail.com")
	d.Add(c)

	if d.Contacts == nil {
		t.Error("The addressbook hasn't been initialized")
	}

	if len(d.Contacts) <= 0 {
		t.Error("The addressbook should have a contact")
	}

	if name := d.Contacts[0].Name; name != "Alejandro" {
		t.Error("The contact name doesn't match")
	}
}
