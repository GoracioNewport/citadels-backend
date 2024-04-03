package game

import (
	"citadels-backend/logic/game/ability/enums"
	"citadels-backend/logic/game/character"
	enums2 "citadels-backend/logic/game/enums"
	"citadels-backend/ws/dto/game"
)

func (g *Game) generateDtoTurnLoopStage(playerInstance *Player) game.StateDto {
	currentTurn := false
	var currentCharacter *character.Character = nil

	mainDeckActive := false
	globalBankActive := false

	endTurnActive := true

	for _, c := range playerInstance.Characters {
		if c.Class == g.GetCurrentCharacter().Class {
			currentTurn = true
			currentCharacter = c
			break
		}
	}

	if currentTurn {
		globalBankActive = currentCharacter.CheckAbility(enums.BaseLootBankKey)
		mainDeckActive = currentCharacter.CheckAbility(enums.BaseDrawCardsKey)
		endTurnActive = currentCharacter.CheckAbility(enums.BaseEndTurnKey)
	}

	if !currentTurn {
		endTurnActive = false
	}

	characters := make([]game.CharacterInfoDto, 0)
	for _, c := range g.characters {
		characters = append(characters, c.ToInfoDto(g.GetCurrentCharacter().Class, false))
	}

	return game.StateDto{
		Players:          generatePlayersDto(g),
		Characters:       characters,
		Player:           generatePlayerDto(playerInstance, currentCharacter),
		MainDeckSize:     len(g.mainDeck),
		MainDeckActive:   mainDeckActive,
		GlobalBankActive: globalBankActive,
		EndTurnActive:    endTurnActive,
		Stage:            g.stage.ToDto(),
	}
}

func (g *Game) generateDtoSelectionStage(playerInstance *Player) game.StateDto {
	currentTurn := false

	if g.players[g.currentPlayerIndex].Name == playerInstance.Name {
		currentTurn = true
	}

	characters := make([]game.CharacterInfoDto, 0)
	for _, c := range g.characters {
		characterDto := c.ToInfoDto(g.GetCurrentCharacter().Class, false)
		characterDto.Turn = false

		found := false
		for _, sc := range g.selectedClasses {
			if sc == c.Class {
				found = true
				break
			}
		}

		if !found && currentTurn && !c.BannedOpen && !c.BannedHidden {
			characterDto.Card.Active = true
		}

		characters = append(characters, characterDto)
	}

	return game.StateDto{
		Players:          generatePlayersDto(g),
		Characters:       characters,
		Player:           generatePlayerDto(playerInstance, nil),
		MainDeckSize:     len(g.mainDeck),
		MainDeckActive:   false,
		GlobalBankActive: false,
		EndTurnActive:    false,
		Stage:            g.stage.ToDto(),
	}
}

func (g *Game) generateDtoEndGameStage(playerInstance *Player) game.StateDto {
	characters := make([]game.CharacterInfoDto, 0)
	for _, c := range g.characters {
		characterDto := c.ToInfoDto(g.GetCurrentCharacter().Class, false)

		characterDto.Turn = false

		characters = append(characters, characterDto)
	}

	return game.StateDto{
		Players:          generatePlayersDto(g),
		Characters:       characters,
		Player:           generatePlayerDto(playerInstance, nil),
		MainDeckSize:     len(g.mainDeck),
		MainDeckActive:   false,
		GlobalBankActive: false,
		EndTurnActive:    false,
		Stage:            g.stage.ToDto(),
	}
}

func generatePlayersDto(g *Game) []game.PlayerCommonInfoDto {
	players := make([]game.PlayerCommonInfoDto, 0)
	for _, p := range g.players {
		players = append(players, p.ToCommonInfoDto(g))
	}

	return players
}

func generatePlayerDto(playerInstance *Player, character *character.Character) game.PlayerInfoDto {
	playerDto := playerInstance.ToInfoDto()

	// Hand active population
	if character != nil {
		for i, c := range playerInstance.Hand {
			if canPlayCard(playerInstance, character, c) {
				playerDto.Hand[i].Active = true
			}
		}
	}

	return playerDto
}

func (g *Game) ToDto(playerInstance *Player) game.StateDto {
	if g.stage == enums2.CharacterSelection {
		return g.generateDtoSelectionStage(playerInstance)
	} else if g.stage == enums2.TurnLoop {
		return g.generateDtoTurnLoopStage(playerInstance)
	} else if g.stage == enums2.EndGame {
		return g.generateDtoEndGameStage(playerInstance)
	}

	return game.StateDto{}
}
