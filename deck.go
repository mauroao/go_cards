package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// with is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	bytes := []byte(d.toString())
	return os.WriteFile(fileName, bytes, 0666)
}

func newDeckFromFile(fileName string) deck {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	sliceString := strings.Split(text, ",")
	return deck(sliceString)
}

func (d deck) shuffle() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	customRand := rand.New(source)

	for i := range d {
		newPosition := customRand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
