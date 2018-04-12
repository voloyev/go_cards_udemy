package main
import "fmt"

func main() {
	//	var card string = "Ace of Spades"
	cards := newDeck()
	fmt.Println(cards.toString())
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	cards.saveToFile("test")
}
