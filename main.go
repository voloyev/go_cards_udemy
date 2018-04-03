package main

func main() {
	//	var card string = "Ace of Spades"
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Five of Diamonds")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
