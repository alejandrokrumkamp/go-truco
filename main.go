package main

func main() {
	aDeck := trucoDeck()
	aDeck.print()
	aDeck.saveToFile("newDeck")
}
