package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// func receiver
func (d deck) print() {
	// lopping and print the el of the cards
	for i, card := range d {
		fmt.Println(i, card)
	}
}
