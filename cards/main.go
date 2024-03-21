package main

func main() {

	cards := deck{"Ace of Diamonds", createCard("Ace of Spades")}
	cards = append(cards, "Six of Spades")
	cards.print()

}
func createCard(name string) string {
	return name
}