package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func newPerson(firstName string, lastName string, email string, zipCode int) person {
	return person{firstName: firstName, lastName: lastName, contact: contactInfo{email: email, zipCode: zipCode}}
}

func (p person) updateFirstName(newFirstName string) {
	// give the value this memory address is pointing at
	p.firstName = newFirstName
}

func (p *person) printFirstName() {
	fmt.Println((*p).firstName)
}

func (p *person) printLastName() {
	fmt.Println((*p).lastName)
}

func (p *person) printFullName() {
	fmt.Println((*p).firstName + " " + p.lastName)
}

func (p person) printFullInfo() {
	fmt.Printf("%+v\n", p)
}

func main() {
	person2 := newPerson("Bar", "Foo", "bar@foo.com", 2275300)
	person2.updateFirstName("Baz")
	person2.printFullName()
	person2.printFullInfo()
}
