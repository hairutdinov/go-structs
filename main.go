package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// first way to declare:
	// alex := person{"Alex", "Anderson"}

	// second way
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// third way
	var alex person
	alex.firstName = "Alex"
	fmt.Printf("%+v", alex)
}
