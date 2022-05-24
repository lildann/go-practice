package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string


func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}



func (d deck) print() {
	for i, card := range d {
		fmt.Println(i + 1, card)
	}
}
// d is the receiver, deck is the type, defined in deck.go
// by convention, refer to the receiver with a one-two
// letter abbreviation that matches the TYPE of the recevier
// 'd' is a value of type 'deck'
