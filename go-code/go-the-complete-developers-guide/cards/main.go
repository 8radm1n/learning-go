package main

func main() {
	cards := newDeck()
	cards.saveToFile("my-cards.txt")
	cards1 := newDeckFromFile("my-cards.txt")
	cards1.print()
}
