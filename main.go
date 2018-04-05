package main

func main() {
	//	var card string = "Ace of Spades"
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
}
