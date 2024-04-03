package game

import (
	"citadels-backend/logic/game/card"
	"citadels-backend/logic/game/character"
)

func canPlayCard(p *Player, c *character.Character, cd card.Building) bool {
	if p == nil || c == nil {
		return false
	}

	var cardInstance card.Building
	foundId := false
	foundName := false

	for _, c := range p.Hand {
		if c.Id == cd.Id {
			cardInstance = c
			foundId = true
		}

		if c.Name == cd.Name {
			foundName = true
		}
	}

	if foundName {
		return false
	}

	if !foundId {
		return false
	}

	if c.AllowedConstructions <= 0 {
		return false
	}

	price := calculateCardPrice(p, cardInstance)

	if price > p.Bank {
		return false
	}

	return true

}

func calculateCardPrice(p *Player, c card.Building) int {
	return c.Price
}
