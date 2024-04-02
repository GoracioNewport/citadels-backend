package game

import (
	"citadels-backend/logic/game/ability/enums"
	"citadels-backend/logic/game/card"
	"citadels-backend/logic/game/character"
	"citadels-backend/ws/dto/game"
	"math/rand"
	"strconv"
)

type Game struct {
	code                     string
	players                  []*Player
	stage                    Stage
	characters               []*character.Character
	mainDeck                 []card.Building
	currentCharacterIndex    int
	currentPlayerIndex       int
	playerRotationIndexStart int
	selectionRoundsLeft      int
	selectedClasses          []character.Class
	bannedClasses            []character.Class
	finishedPlayer           *Player
}

const (
	initialHandSize          = 4
	initialBank              = 2
	townBuildingWinCondition = 7
)

func (g *Game) StartGame() {
	g.BroadcastServerMessage("Игра началась!")

	rand.Shuffle(len(g.players), func(i, j int) {
		g.players[i], g.players[j] = g.players[j], g.players[i]
	})

	g.mainDeck = card.GenerateMainDeck()

	for _, p := range g.players {
		for j := 0; j < initialHandSize; j++ {
			building := g.DrawCardFromMainDeck()

			if building != nil {
				p.Hand = append(p.Hand, *building)
			}
		}

		p.Bank = initialBank
	}

	g.startCharacterSelection()
}

func (g *Game) StartTurnLoop() {
	g.BroadcastServerMessage("Стадия: ход персонажей")
	g.stage = TurnLoop

	for _, c := range g.characters {
		c.ResetAbilities()
	}

	g.currentCharacterIndex = 0

	g.StartTurn()
}

func (g *Game) StartTurn() {
	found := false
	for _, c := range g.selectedClasses {
		if c == g.GetCurrentCharacter().Class {
			found = true
			break
		}
	}

	if !found {
		g.EndTurn()
		return
	}

	g.GetCurrentCharacter().ResetAbilities()

	player := g.GetCurrentPlayer()

	if player != nil {
		g.BroadcastServerMessage("Ход игрока " + player.Name + " (" + g.GetCurrentCharacter().Name + ")")

		if g.GetCurrentCharacter().Class == character.King {
			for _, p := range g.players {
				p.Crown = false
			}

			player.Crown = true

			g.BroadcastServerMessage("Корона перешла к игроку " + player.Name)
		}
	}

	g.BroadcastState()
}

func (g *Game) DrawCardFromMainDeck() *card.Building {
	if len(g.mainDeck) == 0 {
		return nil
	}

	building := g.mainDeck[0]
	g.mainDeck = g.mainDeck[1:]
	return &building
}

func (g *Game) ReturnCardToMainDeck(card card.Building) {
	g.mainDeck = append(g.mainDeck, card)
}

func (g *Game) ClearSelectionBuffer(player *Player, excludeId int) {
	if player == nil {
		return
	}

	for _, b := range player.SelectionBuffer {
		if b.Id == excludeId {
			continue
		}

		g.ReturnCardToMainDeck(b)
	}

	player.SelectionBuffer = nil
}

func (g *Game) EndTurn() {
	player := g.GetCurrentPlayer()

	if player != nil {
		finishedBuildings := len(player.Town)

		if finishedBuildings >= townBuildingWinCondition {
			g.finishedPlayer = player

			g.BroadcastServerMessage("Игрок " + player.Name + " построил " + strconv.Itoa(townBuildingWinCondition) + " кварталов. Игра закончится после хода всех персонажей.")
		}
	}

	g.currentCharacterIndex++

	if g.currentCharacterIndex >= len(g.characters) {
		g.currentCharacterIndex = 0

		if g.finishedPlayer != nil {
			g.currentCharacterIndex = len(g.characters) - 1
			g.endGame()
			return
		}

		g.startCharacterSelection()
		return
	}

	g.StartTurn()
}

func (g *Game) GetCurrentCharacter() *character.Character {
	if g.currentCharacterIndex < 0 || g.currentCharacterIndex >= len(g.characters) {
		return nil
	}

	return g.characters[g.currentCharacterIndex]
}

func (g *Game) GetCurrentPlayer() *Player {
	var index = -1

	if g.GetCurrentCharacter() == nil {
		return nil
	}

	for i, p := range g.players {
		for _, c := range p.Characters {
			if c.Class == g.GetCurrentCharacter().Class {
				index = i
			}
		}
	}

	if index == -1 {
		return nil
	}

	return g.players[index]
}

func (g *Game) GetPlayer(name string) *Player {
	for _, p := range g.players {
		if p.Name == name {
			return p
		}
	}

	return nil
}

func (g *Game) ActivateAbility(player *Player, key enums.Key) bool {
	if player == nil || player.Characters == nil {
		return false
	}

	for _, c := range player.Characters {
		if c.Class == g.GetCurrentCharacter().Class && c.CheckAbility(key) {
			ability := c.GetAbility(key)

			if !ability.Active {
				continue
			}

			if ability.ActivationType == enums.Instant {
				abilityFunction := GetInstantAbilityFunction(key)
				abilityFunction(g, c, player, ability)

				g.BroadcastState()
				return true
			}
		}
	}

	return false
}

func (g *Game) ToDto(playerInstance *Player) game.StateDto {
	if g.stage == CharacterSelection {
		return g.generateDtoSelectionStage(playerInstance)
	} else if g.stage == TurnLoop {
		return g.generateDtoTurnLoopStage(playerInstance)
	} else if g.stage == EndGame {
		return g.generateDtoEndGameStage(playerInstance)
	}

	return game.StateDto{}
}
