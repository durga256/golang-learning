package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.saveToFile("deckResult"))
	deckRead := newDeckFromFile("deckResult")
	fmt.Println((deckRead))

}
