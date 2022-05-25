package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5) // multiple return values from one function
	fmt.Println(hand.toString())
	remainingCards.print()

}

	// Type conversion to byte: fmt.Println([]byte("Hi There!")) -> [72 105 32 84 104 101 114 101 33]
