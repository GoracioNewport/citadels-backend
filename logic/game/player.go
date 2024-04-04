package game

import (
	"citadels-backend/logic/game/card"
	"citadels-backend/logic/game/character"
	"citadels-backend/ws/dto/game"
	"github.com/gorilla/websocket"
)

type Player struct {
	Characters      []*character.Character
	Hand            []card.Building
	Town            []card.Building
	SelectionBuffer []card.Building
	Connection      *websocket.Conn
	Name            string
	Bank            int
	Crown           bool
	Game            *Game
}

func (p *Player) GiveCard(card card.Building) {
	p.Hand = append(p.Hand, card)
}

func (p *Player) HasCharacter(class character.Class) bool {
	for _, c := range p.Characters {
		if c.Class == class {
			return true
		}
	}

	return false
}

func (p *Player) GetCurrentCharacter() *character.Character {
	var characterReference *character.Character = nil

	if p.Game.GetCurrentCharacter() == nil {
		return nil
	}

	for _, c := range p.Characters {
		if c.Class == p.Game.GetCurrentCharacter().Class {
			characterReference = c
			break
		}
	}

	return characterReference
}

func (p *Player) ConstructBuilding(id int) bool {
	cardInstance := card.GetCardById(id)

	if cardInstance == nil {
		return false
	}

	currentCharacter := p.GetCurrentCharacter()

	if currentCharacter == nil {
		return false
	}

	if !canPlayCard(p, currentCharacter, *cardInstance) {
		return false
	}

	p.Town = append(p.Town, *cardInstance)

	for i, c := range p.Hand {
		if c.Id == id {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			break
		}
	}

	p.Bank -= calculateCardPrice(p, *cardInstance)
	currentCharacter.AllowedConstructions--

	return true
}

func (p *Player) TakeCardFromSelectionBuffer(id int) bool {
	for _, b := range p.SelectionBuffer {
		if b.Id == id {
			p.Hand = append(p.Hand, b)
			return true
		}
	}

	return false
}

func (p *Player) ToInfoDto() game.PlayerInfoDto {
	hand := make([]game.CardDto, 0)
	for _, b := range p.Hand {
		hand = append(hand, b.ToDto())
	}

	town := make([]game.CardDto, 0)
	for _, b := range p.Town {
		town = append(town, b.ToDto())
	}

	characters := make([]game.CharacterCardDto, 0)
	for _, c := range p.Characters {
		characters = append(characters, c.ToCardDto(true))
	}

	if len(p.Characters) == 0 {
		characters = append(characters, game.GetUnknownCharacterCardDto())
	}

	return game.PlayerInfoDto{
		Bank:       p.Bank,
		Crown:      p.Crown,
		Characters: characters,
		Hand:       hand,
		Town:       town,
	}
}

func (p *Player) ToCommonInfoDto(g *Game) game.PlayerCommonInfoDto {
	charactersCardDto := make([]game.CharacterCardDto, 0)

	for _, c := range p.Characters {
		if c.Order <= g.GetCurrentCharacter().Order && !c.Dead {
			charactersCardDto = append(charactersCardDto, c.ToCardDto(false))
		} else {
			charactersCardDto = append(charactersCardDto, game.GetUnknownCharacterCardDto())
		}
	}

	town := make([]game.CardDto, 0)
	for _, b := range p.Town {
		town = append(town, b.ToDto())
	}

	return game.PlayerCommonInfoDto{
		Name:         p.Name,
		Bank:         p.Bank,
		HandDeckSize: len(p.Hand),
		Characters:   charactersCardDto,
		Crown:        p.Crown,
		Town:         town,
	}
}
