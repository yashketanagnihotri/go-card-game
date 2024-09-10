package main

import "fmt"

// create a new type of deck
// which is a slice of strings

// It is like a subclass of a slice of strings, has all behavior of a string with extra features
type deck []string

// This method whenever called will be returning a deck of cards and that has been mentioned here as the return type
// Since we are not taking any inputs so we are not having any recievers.
func newDeck() deck {
	cards := deck{}

	// instead of manually typing the cards we can generate it from 2 arrays

	cardsSuits := []string{"Diamonds", "Spades", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) deal(x int) deck {

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
