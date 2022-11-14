package main

import "testing"

//the functions are automatically called on go test command

func TestNewDeck(t *testing.T) { // t *testing.T is to be the default parameter in tests
	d := newDeck()
	//check if deck length is 16 , if not show error
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got: %v", len(d))
	}

	//check if the first card is Ace of Spaces

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected First card: Ace of Spades, but got: %v", d[0])
	}

	//check id last card is Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Last card: Four of Clubs, but got: %v", d[len(d)-1])
	}
}
