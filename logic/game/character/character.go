package character

import (
	"citadels-backend/logic/game/ability"
	"citadels-backend/logic/game/ability/enums"
	gameDto "citadels-backend/ws/dto/game"
)

type Character struct {
	Class        Class
	Order        int
	Abilities    []*ability.Ability
	Name         string
	Description  []string
	Image        string
	Dead         bool
	Robbed       bool
	BannedHidden bool
	BannedOpen   bool
}

func (c *Character) GetUnusedAbilities() []ability.Ability {
	unusedList := make([]ability.Ability, 0)

	for _, a := range c.Abilities {
		if !a.Active {
			unusedList = append(unusedList, *a)
		}
	}

	return unusedList
}

func (c *Character) SetAbilityActive(key enums.Key, state bool) {
	for _, a := range c.Abilities {
		if a.Key == key {
			a.Active = state
		}
	}
}

func (c *Character) GetAbility(key enums.Key) *ability.Ability {
	for _, a := range c.Abilities {
		if a.Key == key {
			return a
		}
	}

	return nil
}

func (c *Character) CheckAbility(key enums.Key) bool {
	for _, a := range c.Abilities {
		if a.Key == key && a.Active {
			return true
		}
	}

	return false
}

func (c *Character) ResetAbilities() {
	for i := range c.Abilities {
		c.Abilities[i].Active = c.Abilities[i].DefaultState
	}
}

func (c *Character) ToCardDto(generateAbilities bool) gameDto.CharacterCardDto {
	abilities := make([]gameDto.CharacterAbilityDto, 0)

	if generateAbilities {
		for _, a := range c.Abilities {
			if !a.Hidden && a.Active {
				abilities = append(abilities, a.ToDto())
			}
		}
	}

	return gameDto.CharacterCardDto{
		Name:        c.Name,
		Description: c.Description,
		Image:       c.Image,
		Type:        c.Class.ToDto(),
		Abilities:   abilities,
	}
}

func (c *Character) ToInfoDto(currentClass Class, abilities bool) gameDto.CharacterInfoDto {
	return gameDto.CharacterInfoDto{
		Card:   c.ToCardDto(abilities),
		Dead:   c.Dead,
		Robbed: c.Robbed,
		Absent: c.BannedOpen,
		Turn:   currentClass == c.Class,
	}
}
