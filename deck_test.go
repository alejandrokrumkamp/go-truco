package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := trucoDeck()

	checkDeckLength(t, d, 40)
	checkExistenceOfSomeCardsInDeck(t, d, "2 de Copa")
	checkExistenceOfSomeCardsInDeck(t, d, "1 de Oro")
	checkExistenceOfSomeCardsInDeck(t, d, "7 de Espada")
	checkExistenceOfSomeCardsInDeck(t, d, "4 de Basto")
}

func TestSaveToFileAndNewReadDeckFromFile(t *testing.T) {
	d := trucoDeck()
	d.saveToFile("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")

	checkDeckLength(t, loadedDeck, 40)
	checkExistenceOfSomeCardsInDeck(t, d, "3 de Oro")
	checkExistenceOfSomeCardsInDeck(t, d, "2 de Basto")

	os.Remove("_decktesting")
}

/* Utils */
func checkExistenceOfSomeCardsInDeck(t *testing.T, d deck, c string) {
	if !d.isCardInDeck(c) {
		t.Errorf("Failed to find %v in deck", c)
	}
}

func checkDeckLength(t *testing.T, d deck, l int) {
	if len(d) != l {
		t.Errorf("Expected deck length of %v but got %v", l, len(d))
	}
}
