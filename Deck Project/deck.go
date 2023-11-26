package main

import (
	"fmt"
	"os"
	"strings"
)

// create a new type of 'deck'
// which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// reciever of type deck (d deck)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// reciever of type deck (d deck)
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// reciever of type deck (d deck)
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return to call newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}
