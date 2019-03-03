package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := trucoDeck()
	expectedLength := 40

	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %v. Got %v", expectedLength, len(d))
	}
}
