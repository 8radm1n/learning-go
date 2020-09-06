package main

import "fmt"

// Define a person struct with two properties.
type person struct {
	firstName string
	lastName  string
}

func main() {
	// Shorthand syntax
	// alex := person{"Alex", "Anderson"}

	// Preferred method to create struct instance
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}
	fmt.Println(alex)
}
