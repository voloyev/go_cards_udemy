package main

//import "fmt"

func main() {
	cards := newDeckFromFile("test")
	cards.shuffle()
	cards.print()
}
