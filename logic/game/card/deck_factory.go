package card

import (
	"citadels-backend/logic/game/card/template"
	"math/rand"
	"strconv"
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

var id = 0
var cardById = make(map[int]Building)

func BuildingFromTemplate(template template.BuildingTemplate) Building {
	cardId := id
	id++

	var image string
	if template.Image == "" {
		image = string(template.Color.ToDto()) + "_" + strconv.Itoa(template.Price) + ".png"
	} else {
		image = template.Image
	}

	building := Building{
		Name:        template.Name,
		Color:       template.Color,
		Description: template.Description,
		Price:       template.Price,
		Image:       image,
		Id:          cardId,
	}

	cardById[cardId] = building

	return building
}

func GetCardById(id int) *Building {
	card, ok := cardById[id]

	if !ok {
		return nil
	}

	return &card
}
