package main

func main() {
	cards := newDeck(true)
	// // cards.print()

	// hand, remainingCards := cards.deal(4)
	// remainingCards.print()
	// hand.print()
	// hand.saveToFile("hand.txt")
	// remainingCards.saveToFile("remainingCards.txt")
	// cards := deckFromFile("hand.txt")
	// cards = deckFromFile("remainingCards2.txt")
	cards.print()
}
