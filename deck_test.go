package main

import (
	"os"
	"testing"
)

// write a test around newDeck()
// 52 items inside of it = length
// First card = Ace of Spades
// Last card = King of Clubs

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got %v", d[0])
	}

	if d[len(d) -1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, got %v", d[len(d) -1])
	}
}

// testing file IO including cleanup of files

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 card in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
