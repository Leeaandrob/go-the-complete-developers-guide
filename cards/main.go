package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"

	// creating a new slice
	cards := deck{newCard(), newCard()}

	// adding el to cards
	cards = append(cards, newCard())
	cards = append(cards, "Six of Spades")

	// print the el of the cards
	fmt.Println(cards)

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
