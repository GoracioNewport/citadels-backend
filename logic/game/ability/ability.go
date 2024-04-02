package ability

import (
	"citadels-backend/logic/game/ability/enums"
	"citadels-backend/ws/dto/game"
)

type Ability struct {
	Key            enums.Key
	ActivationType enums.ActivationType
	Active         bool
	DefaultState   bool
	Hidden         bool
}

func NewAbility(key enums.Key, activationType enums.ActivationType, defaultState bool, hidden bool) Ability {
	return Ability{Key: key, ActivationType: activationType, DefaultState: defaultState, Hidden: hidden}
}

func (a *Ability) ToDto() game.CharacterAbilityDto {
	return game.CharacterAbilityDto{
		Name: a.Key.GetName(),
		Key:  a.Key.ToDto(),
	}
}
