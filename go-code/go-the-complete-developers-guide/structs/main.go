package main

import "fmt"

// Define a person struct with two properties.
type person struct {
	firstName string
	lastName  string
	contact   contactInfo // embedded struct
	// Alternately you can ommit the property and use 'contactInfo' by itself
	// contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// Create person struct with zero values assigned.
	// Zero value for a string is ""
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// Shorthand syntax to create and assign properties
	// alex := person{"Alex", "Anderson"}

	// Preferred method to create struct instance
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }
	// fmt.Println(alex)

	// // Print properties and values with %+v
	// fmt.Printf("%+v", alex)

	// Creating an embeded struct
	jim := person{
		firstName: "Jim",
		lastName:  "Parsons",
		// contactInfo: contactInfo{} if immiting the 'contact' property in the struct
		contact: contactInfo{
			email:   "jp@bigbang.theory",
			zipCode: 90210,
		},
	}
	// This is the long way to do it.
	// jimPointer := &jim
	// jimPointer.updateName("Leonard")

	// Go will automatically grab the memory address if you have a function
	// with a receiver type of *person so you dont have to create a transient
	// variable 'jimPointer' provinding a nice shortcut.
	jim.updateName("Leonard")
	jim.print()
}

// You can add reciever functions to structs
func (p person) print() {
	fmt.Printf("%+v", p)
}

// This does not work because go passes a copy of
// the variable to the function. Go is a pass by
// value language.
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
