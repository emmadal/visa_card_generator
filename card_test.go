package main

import (
	"math/rand"
	"testing"
)

var cards = []int{4}
var cardLength = 16

func TestCardLength(t *testing.T) {
	for i := 1; i < cardLength; i++ {
		cards = append(cards, rand.Intn(10))
	}
	if len(cards) != cardLength {
		t.Errorf("Number of cards is incorrect, expected %d, got %d", cardLength, len(cards))
	}
}

func TestFirstNumber(t *testing.T) {
	if cards[0] != 4 {
		t.Errorf("Incorrect visa card, expected 4 for the first number, got %d",
			cards[0])
	}
}
