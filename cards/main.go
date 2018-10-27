package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"

	// creating a new slice
	cards := []string{newCard(), newCard()}

	// adding el to cards
	cards = append(cards, newCard())
	cards = append(cards, "Six of Spades")

	// print the el of the cards
	fmt.Println(cards)

	// lopping and print the el of the cards
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
