package main

func main() {
	cards := newDeck()


	hand, remainingCards := deal(cards, 5) // multiple return values from one function
	hand.print()
	remainingCards.print()
}

