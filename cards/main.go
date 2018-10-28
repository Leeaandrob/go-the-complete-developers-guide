package main

func main() {
	//var card string = "Ace of Spades"

	// creating a new slice
	cards := newDeck()

	// using func print of deck
	cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}
