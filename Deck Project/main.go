package main

func main() {
	cards := newDeck()

	fmt.Println(cards.toString())

	cards.saveToFile("my_deck")

	cards = cards.newDeckFromFile("my_deck")
	cards.print()

}
