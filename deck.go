package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) deal() (deck, deck) {
	return d[:3], d[3:]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(path string) error {
	return ioutil.WriteFile(path, []byte(d.toString()), 0666)
}

func trucoDeck() deck {
	var aDeck deck
	cardsNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "10", "11", "12"}
	suits := []string{"Oro", "Copa", "Espada", "Basto"}

	for _, suit := range suits {
		for _, cardsNumber := range cardsNumbers {
			aDeck = append(aDeck, cardsNumber+" de "+suit)
		}
	}

	return aDeck
}
