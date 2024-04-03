package enums

import "citadels-backend/ws/dto/game"

type Stage int

const (
	CharacterSelection Stage = iota
	TurnLoop
	EndGame
)

func (s Stage) ToDto() game.Stage {
	switch s {
	case CharacterSelection:
		return game.CharacterSelection
	case TurnLoop:
		return game.TurnLoop
	case EndGame:
		return game.EndGame
	}

	return "unknown"
}
