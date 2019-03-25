package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := trucoDeck()

	expectedDeckLength := 40
	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length of %v but got %v", expectedDeckLength, len(d))
	}

	expectedFirstCard := "1 de Oro"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %v but got %v", expectedFirstCard, d[0])
	}

	expectedLastCard := "12 de Basto"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card to be %v but got %v", expectedLastCard, d[len(d)-1])
	}
}

func TestSaveToFileAndNewReadDeckFromFile(t *testing.T) {
	d := trucoDeck()
	d.saveToFile("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")

	expectedDeckLength := 40
	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected loaded deck to have %v. Got %v", expectedDeckLength, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
