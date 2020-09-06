package main

import "fmt"

// Define a person struct with two properties.
type person struct {
	firstName string
	lastName  string
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
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}
	fmt.Println(alex)

	// Print properties and values with %+v
	fmt.Printf("%+v", alex)
}
