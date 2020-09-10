package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	ebot := englishBot{}
	sbot := spanishBot{}

	ebot.printGreeting()
	sbot.printGreeting()
}

func (eb englishBot) printGreeting() {
	fmt.Println(eb.getGreeting())
}

func (sb spanishBot) printGreeting() {
	fmt.Println(sb.getGreeting())
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
