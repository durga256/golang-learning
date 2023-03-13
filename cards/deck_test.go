package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Error("Expected deck length of 52 but got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected first element to be Ace of Spades but got", d[0])
	}
	if d[len(d)-1] != "King of Diamonds" {
		t.Error("Expected Last element to be King of Diamonds but got", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckResult")

	deck := newDeck()
	if deck.saveToFile("_deckResult") == "OK" {
		t.Error("Save to File failed")
	}

	loadedDeck := newDeckFromFile("_deckResult")

	if len(loadedDeck) != 52 {
		t.Error("Expected 52 cards in deck, got ", len(loadedDeck))
	}

	os.Remove("_deckResult")
}
