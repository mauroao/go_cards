package main

import "fmt"

func main() {
	cards := newDeckFromFile("abc.txt")

	cards.shuffle()

	fmt.Println(cards)

}
