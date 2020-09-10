package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	ebot := englishBot{}
	sbot := spanishBot{}

	printGreeting(ebot)
	printGreeting(sbot)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

// If you are not using the reference that you
// attach to the method you can omit it
// from the receiver.
// func (spanishBot) getGreeting() string {
// 	return "Hola"
// }
