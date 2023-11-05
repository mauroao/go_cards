package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	dLength := len(d)
	if dLength != 16 {
		t.Errorf("Expected deck length of 16, bug got %v", dLength)
	}

	spected := "Ace of Spades"
	if d[0] != spected {
		t.Errorf("Expected first deck of %v but got %v", spected, d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	loadedDesk := newDeckFromFile(fileName)

	dLength := len(loadedDesk)
	if dLength != 16 {
		t.Errorf("Expected deck length of 16, bug got %v", dLength)
	}

	os.Remove(fileName)
}
