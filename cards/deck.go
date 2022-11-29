package main

import "fmt"

//create a new type of 'deck'
//whic is a slice of strings
type deck []string // class declaration

func newDeck() deck { // como es la creación de un nuevo objeto, no necesita receiver
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}



func (d deck) print() { // agregamos el receiver solo cuando queremos hacer algo con el objeto ya creado
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
return d[:handSize],d[handSize:]
}
