package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) deal() (deck, deck) {
	return d[:3], d[3:]
}

func (d deck) shuffle() deck {
	now := time.Now().UnixNano()
	source := rand.NewSource(now)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func bsToDeck(bs []byte) deck {
	return strings.Split(string(bs), ",")
}

func (d deck) saveToFile(path string) error {
	return ioutil.WriteFile(path, []byte(d.toString()), 0666)
}

func (d deck) isCardInDeck(c string) bool {
	for _, deckCard := range d {
		if deckCard == c {
			return true
		}
	}
	return false
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

func readDeckFromFile(path string) deck {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bsToDeck(bs)
}
