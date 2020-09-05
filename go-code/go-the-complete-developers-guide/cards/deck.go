package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
// new types allow you to add functionality to
// existing types
type deck []string

// Add print method to the deck type
// this allows you to use the print method on
// a variable made from the deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
