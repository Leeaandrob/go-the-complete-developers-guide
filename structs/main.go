package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func newPerson(firstName string, lastName string) person {
	return person{firstName: firstName, lastName: lastName}
}

func updatePerson(p person, firstName string, lastName string) person {
	p.firstName = firstName
	p.lastName = lastName
	return p
}

func (p *person) printFirstName() {
	fmt.Println(p.firstName)
}

func (p *person) printLastName() {
	fmt.Println(p.lastName)
}

func (p *person) printFullName() {
	fmt.Println(p.firstName + " " + p.lastName)
}

func main() {
	person2 := newPerson("Bar", "Foo")
	person2.printFullName()
	person2 = updatePerson(person2, "Barr", "Fooo")
	person2.printFullName()
}
