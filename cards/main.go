package main

import "fmt"

func main() {

	cards := newDeck()

	hand, left := deal(cards, 5)

	fmt.Println("Hand:" , len(hand))
	fmt.Println("Left:" , len(left))

	fmt.Println(hand.toString())
	cards.saveToFile("my_cards.txt")


	fmt.Println("New Deck from file")
	newDeck := newDeckFromFile("my_cards.txt")

	hand2, _ := deal(newDeck, 5)
	hand2.shuffle()
	hand2.print()

}
