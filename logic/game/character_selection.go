package game

import (
	"citadels-backend/logic/game/character"
	"math/rand"
	"sort"
)

func (g *Game) startCharacterSelection() {
	g.stage = CharacterSelection
	g.BroadcastServerMessage("Стадия: выбор персонажей")

	g.selectedClasses = make([]character.Class, 0)
	g.bannedClasses = make([]character.Class, 0)

	g.currentPlayerIndex = 0
	for i, p := range g.players {
		if p.Crown {
			g.currentPlayerIndex = i
		}

		p.Characters = nil
	}

	g.selectionRoundsLeft = 1
	hiddenCharacterBanRounds := 1
	openCharacterBanRounds := 1

	if len(g.players) < 4 {
		//g.selectionRoundsLeft = 2
		openCharacterBanRounds = 0
	} else {
		if len(g.players) == 4 {
			openCharacterBanRounds = 2
		} else if len(g.players) == 5 {
			openCharacterBanRounds = 1
		} else {
			openCharacterBanRounds = 0
		}
	}

	g.characters = character.NewCharacters(false)

	rand.Shuffle(len(g.characters), func(i, j int) {
		g.characters[i], g.characters[j] = g.characters[j], g.characters[i]
	})

	for i := 0; i < hiddenCharacterBanRounds; i++ {
		g.characters[i].BannedHidden = true
		g.bannedClasses = append(g.bannedClasses, g.characters[i].Class)
	}

	for i := 0; i < openCharacterBanRounds; i++ {
		g.characters[i+hiddenCharacterBanRounds].BannedOpen = true
		g.bannedClasses = append(g.bannedClasses, g.characters[i+hiddenCharacterBanRounds].Class)
	}

	g.characters = append(g.characters, character.NewKing())

	sort.Slice(g.characters, func(i, j int) bool { return g.characters[i].Order < g.characters[j].Order })

	g.playerRotationIndexStart = g.currentPlayerIndex

	g.startSelectionRound()
}

func (g *Game) startSelectionRound() {

	if g.currentPlayerIndex == g.playerRotationIndexStart {
		g.selectionRoundsLeft--
	}

	if g.selectionRoundsLeft < 0 {
		g.StartTurnLoop()
		return
	}

	g.BroadcastState()
}

func (g *Game) availableForSelection(class character.Class) bool {
	for _, c := range g.selectedClasses {
		if c == class {
			return false
		}
	}

	for _, c := range g.characters {
		if c.Class == class {
			if c.BannedHidden || c.BannedOpen {
				return false
			}
		}
	}

	return true
}

func (g *Game) PickCharacter(player *Player, class character.Class) bool {
	if g.stage != CharacterSelection {
		return false
	}

	if player.Name != g.players[g.currentPlayerIndex].Name {
		return false
	}

	if !g.availableForSelection(class) {
		return false
	}

	var selectedCharacter *character.Character
	for _, c := range g.characters {
		if c.Class == class {
			selectedCharacter = c
			break
		}
	}

	g.selectedClasses = append(g.selectedClasses, class)
	player.Characters = append(player.Characters, selectedCharacter)

	g.currentPlayerIndex++
	if g.currentPlayerIndex >= len(g.players) {
		g.currentPlayerIndex = 0
	}

	g.startSelectionRound()

	return true
}
