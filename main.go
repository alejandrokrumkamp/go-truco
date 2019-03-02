package main

func main() {
	aDeck := trucoDeck()
	aDeck.saveToFile("newDeck")
	anotherDeck := readDeckFromFile("newDeck")
	anotherDeck = anotherDeck.shuffle()
	anotherDeck.print()
}
