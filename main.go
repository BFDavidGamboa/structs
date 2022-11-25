package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "jimsmith@example.com",
			zipCode: 12345,
		},
	}

	// jim.print()
	jim.updateName("Jimmy")
	jim.print()``
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
