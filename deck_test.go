package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected := 16

	if len(d) != expected {
		t.Errorf("Expected deck len %v, but got %v", expected, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last Four of Clubs but got %v", d[len(d)-1])
	}
}
