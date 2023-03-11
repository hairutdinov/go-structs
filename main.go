package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// first way to declare:
	// alex := person{"Alex", "Anderson"}

	// second way
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// third way
	var alex person
	alex.firstName = "Alex"
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@mail.com",
			zipCode: "94000",
		},
	}
	fmt.Printf("%+v", jim)
}
