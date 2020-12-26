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

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// when we do not need index, it can be replaced by underscore as shown below
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error - ", err)
		/*
			0 indicates program terminated successfully
			whereas a non-zero values indicates something incorrect happened
		*/
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	seed := time.Now().UnixNano()

	// creating a new source that shall take the generated seed for random no. generation
	source := rand.NewSource(seed)
	randomNumGen := rand.New(source)

	for i := range d {
		newPosition := randomNumGen.Intn(len(d) - 1)

		// swap in GoLang
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

type evenOddSlice []int

func evenOdd() {
	eo := evenOddSlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range eo {
		if eo[i]%2 == 0 {
			fmt.Printf("%v is even\n", eo[i])
		} else {
			fmt.Printf("%v is odd\n", eo[i])
		}
	}
}
