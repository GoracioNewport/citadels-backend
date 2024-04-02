package card

import (
	"citadels-backend/logic/game/card/template"
	"math/rand"
)

const (
	purpleCardLimit = 14
)

func GenerateMainDeck() []Building {
	var deck = make([]Building, 0)
	deck = append(deck, GenerateCards(template.Yellow)...)
	deck = append(deck, GenerateCards(template.Green)...)
	deck = append(deck, GenerateCards(template.Red)...)
	deck = append(deck, GenerateCards(template.Blue)...)

	var purpleCards = make([]Building, 0)
	purpleCards = append(purpleCards, GenerateCards(template.Purple)...)

	rand.Shuffle(len(purpleCards), func(i, j int) {
		purpleCards[i], purpleCards[j] = purpleCards[j], purpleCards[i]
	})

	for i := 0; i < min(purpleCardLimit, len(purpleCards)); i++ {
		deck = append(deck, purpleCards[i])
	}

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}
