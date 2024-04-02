package game

import (
	"citadels-backend/logic/game/ability"
	"citadels-backend/logic/game/ability/enums"
	"citadels-backend/logic/game/card"
	"citadels-backend/logic/game/card/template"
	"citadels-backend/logic/game/character"
	"citadels-backend/ws/dto/game"
	"citadels-backend/ws/sender"
	"log"
)

type ActivateInstantAbility func(g *Game, c *character.Character, p *Player, a *ability.Ability)

func ActivateBaseLootBank(g *Game, c *character.Character, p *Player, a *ability.Ability) {
	p.Bank += 2
	a.Active = false

	c.SetAbilityActive(enums.BaseDrawCardsKey, false)
	c.SetAbilityActive(enums.BaseEndTurnKey, true)
}

func ActivateBaseDrawCards(g *Game, c *character.Character, p *Player, a *ability.Ability) {
	cardAmount := 2
	cardList := make([]card.Building, 0)

	for i := 0; i < cardAmount; i++ {
		cardInstance := g.DrawCardFromMainDeck()

		if cardInstance != nil {
			cardList = append(cardList, *cardInstance)
		}
	}

	p.SelectionBuffer = append(p.SelectionBuffer, cardList...)

	cardDtoList := make([]game.CardDto, 0)
	for _, cardInstance := range cardList {
		cardDto := cardInstance.ToDto()
		cardDto.Active = true

		cardDtoList = append(cardDtoList, cardDto)
	}

	sender.SendGameChooseCardsMessage(p.Connection, ManagerInstance.Server, cardDtoList)

	a.Active = false

	c.SetAbilityActive(enums.BaseLootBankKey, false)
	c.SetAbilityActive(enums.BaseEndTurnKey, true)
}

func ActivateBaseConstructBuilding(g *Game, c *character.Character, p *Player, a *ability.Ability) {
	a.Active = false
}

func ActivateBaseEndTurn(g *Game, c *character.Character, p *Player, a *ability.Ability) {
	a.Active = false
	g.EndTurn()
}

func activateColorLootCity(p *Player, a *ability.Ability, color template.BuildingColor) {
	loot := 0

	for _, b := range p.Town {
		if b.Color == color {
			loot++
		}
	}

	p.Bank += loot
	a.Active = false
}

func ActivateKingLootCity(g *Game, c *character.Character, p *Player, a *ability.Ability) {
	activateColorLootCity(p, a, template.Yellow)
}

func ActivateBishopLootCity(g *Game, c *character.Character, p *Player, a *ability.Ability) {
	activateColorLootCity(p, a, template.Blue)
}

func ActivateMerchantLootCity(g *Game, c *character.Character, p *Player, a *ability.Ability) {
	activateColorLootCity(p, a, template.Green)
}

func ActivateMerchantLootBank(g *Game, c *character.Character, p *Player, a *ability.Ability) {
	p.Bank++
	a.Active = false
}

func ActivateWarlordLootCity(g *Game, c *character.Character, p *Player, a *ability.Ability) {
	activateColorLootCity(p, a, template.Red)
}

var InstantAbilityFunctions = map[enums.Key]ActivateInstantAbility{
	enums.BaseLootBankKey:          ActivateBaseLootBank,
	enums.BaseDrawCardsKey:         ActivateBaseDrawCards,
	enums.BaseEndTurnKey:           ActivateBaseEndTurn,
	enums.BaseConstructBuildingKey: ActivateBaseConstructBuilding,
	enums.KingLootCityKey:          ActivateKingLootCity,
	enums.BishopLootCityKey:        ActivateBishopLootCity,
	enums.MerchantLootCityKey:      ActivateMerchantLootCity,
	enums.MerchantLootBankKey:      ActivateMerchantLootBank,
	enums.WarlordLootCityKey:       ActivateWarlordLootCity,
}

func GetInstantAbilityFunction(key enums.Key) ActivateInstantAbility {
	if _, ok := InstantAbilityFunctions[key]; !ok {
		log.Printf("Instant ability function for key %s not found", key)
		return nil
	}

	return InstantAbilityFunctions[key]
}
