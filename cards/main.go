package main

import "fmt"

func main() {

	cards := newDeckFromFile("my_cards")
	fmt.Println(cards.toString())
	cards.print()

	w_cards := newDeckFromFile("my_cars")
	w_cards.print()
}
