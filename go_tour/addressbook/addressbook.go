package addressbook

type Contact struct {
	Name  string
	Email string
}

type AddressBook struct {
	Contacts []Contact
}

func NewContact(name, email string) Contact {
	return Contact{name, email}
}

func (d *AddressBook) Add(contact Contact) {
	d.Contacts = append(d.Contacts, contact)
}
