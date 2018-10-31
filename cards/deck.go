package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// creating a instance of deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// loop to append suit and value into cards(deck)
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// func receiver
func (d deck) print() {
	// lopping and print the el of the cards
	for i, card := range d {
		fmt.Println(i, card)
		// print []byte of card
		// fmt.Println([]byte(card))
	}
}

// doing a slice on deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// converting a deck into string separating using comma
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saving deck into file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// reading file and create a new deck from file opened
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("error opened the file " + filename)
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// will random a number and get the position of deck using the random number
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)

		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
