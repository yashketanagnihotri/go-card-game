package main

import "fmt"

// create a new type of deck
// which is a slice of strings

// It is like a subclass of a slice of strings, has all behavior of a string with extra features
type deck []string

func (d deck) print(x int){
	for i, card := range d {
		fmt.Println(i, card,x)
	}
}

