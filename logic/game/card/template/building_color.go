package template

import "citadels-backend/ws/dto/game"

type BuildingColor int

const (
	Yellow BuildingColor = iota
	Green
	Red
	Blue
	Purple
)

const (
	ColorAmount = 5
)

func (c BuildingColor) ToDto() game.CardColorDto {
	switch c {
	case Yellow:
		return game.Yellow
	case Green:
		return game.Green
	case Red:
		return game.Red
	case Blue:
		return game.Blue
	case Purple:
		return game.Purple
	}

	return ""
}

func BuildingColorFromDto(dto game.CardColorDto) BuildingColor {
	switch dto {
	case game.Yellow:
		return Yellow
	case game.Green:
		return Green
	case game.Red:
		return Red
	case game.Blue:
		return Blue
	case game.Purple:
		return Purple
	}

	return -1
}
