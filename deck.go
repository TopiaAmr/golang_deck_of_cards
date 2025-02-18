// Package main implements a simple deck of cards functionality.
// It provides operations for creating, shuffling, and managing a deck of playing cards.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// deck represents a collection of playing cards as a slice of strings.
type deck []string

// newDeck creates and returns a new deck of cards.
// If shuffle is true, the deck will be randomly shuffled before being returned.
// The deck consists of combinations of suits (Spades, Hearts, Diamonds, Clubs)
// and values (Ace, Two, Three, Four).
func newDeck(shuffle bool) deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, s := range cardSuits {
		for _, v := range cardValues {
			cards = append(cards, v+" of "+s)
		}
	}
	if shuffle {
		cards.shuffleCards()
	}
	return cards
}

// shuffleCards randomly reorders the cards in the deck using the math/rand package.
// It returns a new shuffled deck while leaving the original deck unchanged.
func (d deck) shuffleCards() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

// print displays all cards in the deck to standard output,
// with each card preceded by its index in the deck.
func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

// deal splits the deck into two parts: a hand and the remaining deck.
// It returns two decks: the first handSize cards (the hand),
// and the remaining cards (the rest of the deck).
// If handSize is greater than the deck size, it will panic.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString converts the deck into a single string with cards
// separated by commas. This format is suitable for saving to a file.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saveToFile writes the deck to a file in CSV format.
// The filename parameter specifies where to save the deck.
// Returns an error if the write operation fails.
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// deckFromFile reads a deck from a file that was previously saved using saveToFile.
// The filename parameter specifies the file to read from.
// If there's an error reading the file, it prints the error and exits the program.
func deckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		println("Error: ", err.Error())
		os.Exit(1)
	}
	d := strings.Split(string(data), ",")
	return deck(d)
}
