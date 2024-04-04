package ability

import "citadels-backend/ws/enums"

type TargetContext struct {
	TargetArea  TargetArea
	TargetValue string
}

type TargetArea string

const (
	Player    TargetArea = "player"
	Character TargetArea = "character"
	Card      TargetArea = "card"
)

func TargetContextFromDto(dto enums.GameTargetAbilityMessage) TargetContext {
	return TargetContext{
		TargetArea:  TargetAreaFromDto(dto.TargetArea),
		TargetValue: dto.TargetValue,
	}
}

func TargetAreaFromDto(dto string) TargetArea {
	switch dto {
	case "player":
		return Player
	case "character":
		return Character
	case "card":
		return Card
	}

	return "unknown"
}
