package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

//create a new type of deck

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func newDeck() deck {
	cards := deck{}
	cardTypes := []string{"Spades", "Hearts", "Clovers", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, value := range cardValues {
		for _, suite := range cardTypes {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	sliceDeck := []string(d)
	deckString := strings.Join(sliceDeck, ",")
	return deckString
}

func (d deck) saveToFile(filename string) string {
	err := os.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		return ("Error: " + err.Error())
	} else {
		return ("OK")
	}

}
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
func (d deck) shuffle() {
	for i := range d {
		newIndex := rand.Int() % len(d)
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}
