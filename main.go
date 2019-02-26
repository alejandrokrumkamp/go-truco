package main

func main() {
	aDeck := trucoDeck()
	aDeck.saveToFile("newDeck")
	anotherDeck := readDeckFromFile("newDeck")
	anotherDeck.print()
}
