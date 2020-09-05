package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedfirstCard := "Ace of Spades"
	expectedlastCard := "King of Clubs"
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 got: %v", len(d))
	}

	if d[0] != expectedfirstCard {
		t.Errorf("Expected %v got: %v", expectedfirstCard, d[0])
	}

	if d[len(d)-1] != expectedlastCard {
		t.Errorf("Expected %v got: %v", expectedlastCard, d[len(d)-1])
	}
}
