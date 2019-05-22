package main

import (
	"fmt"
	"strconv"
)

type deck struct {
	suit     []string
	order    []string
	shuffled order.shuffle
}

// constructor function
func newDeck() *deck {
	var c deck
	c.suit = []string{"H", "S", "C", "D"}

	for i := 1; i < 13; i++ {
		for _, j := range c.suit {
			c.order = append(c.order, j+strconv.Itoa(i))
		}
	}
	return &c
}

func main() {
	cards := newDeck()
	for _, i := range cards.order {
		fmt.Println(i)
	}
}
