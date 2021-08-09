package main

// Create a new type of `deck`
// which is a slice of strings

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
