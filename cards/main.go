package main

<<<<<<< HEAD
func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
=======
import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
	printState()
}

func newCard() string {
	return "Five of Diamonds"
>>>>>>> leaning golang
}
