package main

func main() {
	cards := newDeck(true)

	hand, remainingCards := deal(cards, 4)
	remainingCards.print()
	hand.print()
	hand.saveToFile("hand.txt")
	remainingCards.saveToFile("remainingCards.txt")
	hand = deckFromFile("hand.txt")
	cards = deckFromFile("remainingCards2.txt")
	hand.print()
	cards.print()

}
