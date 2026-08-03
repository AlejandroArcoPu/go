package main

import (
	"fmt"

	"example.com/addressbook"
)

func main() {

	a := new(addressbook.AddressBook)
	luke := addressbook.NewContact("Luke", "luke@gmail.com")
	andy := addressbook.NewContact("Andy", "andy@gmail.com")
	a.Add(luke)
	a.Add(andy)
	fmt.Println(a)
}
