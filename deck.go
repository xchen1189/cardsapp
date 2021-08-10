package main

// Create a new type of `deck`
// which is a slice of strings

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i + 1, card)
	}
}

func (d deck) shuffle(){}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func NewDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}
	return cards
}
