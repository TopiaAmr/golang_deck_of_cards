package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck(false)

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndTestNewDeckFromFile(t *testing.T) {
	fileName := "_testdeckfile"
	os.Remove(fileName)

	d := newDeck(false)
	d.saveToFile(fileName)

	loadedFromDisk := deckFromFile(fileName)

	if len(loadedFromDisk) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedFromDisk))
	}

	if loadedFromDisk[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", loadedFromDisk[0])
	}

	if loadedFromDisk[len(loadedFromDisk)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", loadedFromDisk[len(loadedFromDisk)-1])
	}

	os.Remove(fileName)
}

func TestDeal(t *testing.T) {
	d := newDeck(false)

	h, r := deal(d, 5)

	if len(h) != 5 {
		t.Errorf("Expected to have 5 cards in hand, but got %v", len(h))
	}

	if len(r) != len(d)-5 {
		t.Errorf("Expected to have %v remaining cards, but got %v", len(d)-5, len(r))
	}
}
