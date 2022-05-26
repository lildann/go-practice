package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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

func deal(d deck, handSize int) (deck, deck) { // returning two values of type deck
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	biteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1: log the error and return a call to newDeck()
		// option 2: log the error and quit the program entirely - common pattern:
		fmt.Println("Error", err)
		os.Exit(1)
	}

	ss := strings.Split(string(biteSlice), ", ")
	return deck(ss)

}
