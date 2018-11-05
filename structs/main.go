package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	jim.print()
}

// will only be called when the pointer is pointing at a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// get the value from the pointer
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
