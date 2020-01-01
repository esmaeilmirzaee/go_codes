package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type education struct {
	degree string
	field  string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	education
}

func main() {
	jim := person{
		firstName: "Jin",
		lastName:  "Party",
		contact: contactInfo{
			zipCode: 98001,
			email:   "e@g.co",
		},
		education: education{
			field:  "Computer Science",
			degree: "BSc",
		},
	}
	jim.print()
	pointerToJim := &jim
	pointerToJim.updateName("Jim")
	jim.print()
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Println(p.firstName, ", ", p.lastName, ", ", p.contact.email, ", ", p.contact.zipCode, ", ", p.education.degree, ", ", p.education.field)
}
