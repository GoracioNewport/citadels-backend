package game

import (
	"citadels-backend/logic/game/ability/enums"
	"citadels-backend/logic/game/character"
	"citadels-backend/ws/dto/game"
)

type AbilityPopulator func(player *Player, currentCharacter *character.Character, game *Game, dto *game.StateDto)

func populateAssassinKillCharacterAbility(player *Player, currentCharacter *character.Character, game *Game, dto *game.StateDto) {
	for i, c := range dto.Characters {
		if c.Card.Order <= currentCharacter.Order {
			continue
		}

		dto.Characters[i].Card.Active = true
	}
}

func populateThiefLootCharacterAbility(player *Player, currentCharacter *character.Character, game *Game, dto *game.StateDto) {
	for i, c := range dto.Characters {
		if c.Card.Order <= currentCharacter.Order && !c.Dead {
			continue
		}

		dto.Characters[i].Card.Active = true
	}
}

func populateMagicianSwapCardsAbility(player *Player, currentCharacter *character.Character, game *Game, dto *game.StateDto) {
	for i, p := range dto.Players {
		if p.Name == player.Name {
			continue
		}

		dto.Players[i].Active = true
	}
}

var abilityPopulators = map[enums.Key]AbilityPopulator{
	enums.AssassinKillCharacterKey: populateAssassinKillCharacterAbility,
	enums.ThiefLootCharacterKey:    populateThiefLootCharacterAbility,
}

func GetAbilityPopulator(key enums.Key) AbilityPopulator {
	if _, ok := abilityPopulators[key]; !ok {
		return nil
	}

	return abilityPopulators[key]
}
