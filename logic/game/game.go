package game

import (
	"citadels-backend/logic/game/card"
	"citadels-backend/logic/game/character"
	enums2 "citadels-backend/logic/game/enums"
	"math/rand"
	"strconv"
)

type Game struct {
	code                     string
	players                  []*Player
	stage                    enums2.Stage
	characters               []*character.Character
	mainDeck                 []card.Building
	currentCharacterIndex    int
	currentPlayerIndex       int
	playerRotationIndexStart int
	selectionRoundsLeft      int
	selectedCharacters       []*character.Character
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
	g.stage = enums2.TurnLoop

	for _, c := range g.characters {
		c.ResetCharacter()
	}

	g.currentCharacterIndex = 0

	g.StartTurn()
}

func (g *Game) StartTurn() {
	found := false
	for _, c := range g.selectedCharacters {
		if c.Class == g.GetCurrentCharacter().Class {
			if c.Robbed {
				thief := g.GetPlayerByCharacter(character.Thief)
				if thief == nil {
					continue
				}

				g.SendServerMessage(g.GetCurrentPlayer(), "Вас ограбили и вы теряете все свои деньги!")
				thief.Bank += g.GetCurrentPlayer().Bank
				g.GetCurrentPlayer().Bank = 0
			}

			if c.Dead {
				g.SendServerMessage(g.GetCurrentPlayer(), "Вас убили и вы пропускаете свой ход!")
				break
			}

			found = true
			break
		}
	}

	if !found {
		g.EndTurn()
		return
	}

	g.GetCurrentCharacter().ReloadCharacter()

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

func (g *Game) GetCharacter(class character.Class) *character.Character {
	for _, c := range g.characters {
		if c.Class == class {
			return c
		}
	}

	return nil
}

func (g *Game) GetPlayerByCharacter(class character.Class) *Player {
	for _, p := range g.players {
		for _, c := range p.Characters {
			if c.Class == class {
				return p
			}
		}
	}

	return nil
}
