package game

type Stage string

const (
	CharacterSelection Stage = "character_selection"
	TurnLoop           Stage = "turn_loop"
	EndGame            Stage = "end_game"
)
