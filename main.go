package main

func main() {
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")

	//cards := NewDeck()
	//cards.print()
	//hand, remainning := deal(cards, 5)
	//hand.print()
	//remainning.print()
	//cards.saveToFile("my_cards")
	cards, _ := newDeckFromFile("my_cards")
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
