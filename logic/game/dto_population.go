package game

import (
	"citadels-backend/logic/game/ability/enums"
	"citadels-backend/ws/dto/game"
)

func (g *Game) populateTurnLoopDto(playerInstance *Player, dto *game.StateDto) {
	currentCharacter := g.GetCurrentCharacter()
	currentTurn := playerInstance.HasCharacter(currentCharacter.Class)

	if !currentTurn {
		return
	}

	if currentCharacter.PendingTargetedAbility != nil {
		abilityPopulator := GetAbilityPopulator(currentCharacter.PendingTargetedAbility.Key)
		if abilityPopulator != nil {
			dto.PendingAbility = true

			abilityPopulator(playerInstance, currentCharacter, g, dto)

			for i, c := range dto.Player.Characters {
				if c.Type != currentCharacter.Class.ToDto() {
					continue
				}

				for j, a := range c.Abilities {
					abilityKey, ok := enums.KeyFromDto(a.Key)
					if !ok {
						continue
					}

					if abilityKey != currentCharacter.PendingTargetedAbility.Key {
						continue
					}

					dto.Player.Characters[i].Abilities[j].Name = "Отменить"
					dto.Player.Characters[i].Abilities[j].Active = true
				}
			}

			return
		}

	}

	dto.GlobalBankActive = currentCharacter.CheckAbility(enums.BaseLootBankKey)
	dto.MainDeckActive = currentCharacter.CheckAbility(enums.BaseDrawCardsKey)
	dto.EndTurnActive = currentCharacter.CheckAbility(enums.BaseEndTurnKey)

	for i, c := range dto.Player.Characters {
		alternativeIndex := findAlternativeIndex(playerInstance, c)
		if alternativeIndex == -1 {
			continue
		}

		for j, a := range c.Abilities {
			abilityKey, ok := enums.KeyFromDto(a.Key)
			if !ok {
				continue
			}

			dto.Player.Characters[i].Abilities[j].Active = playerInstance.Characters[alternativeIndex].CheckAbility(abilityKey)
		}
	}
}

func findAlternativeIndex(playerInstance *Player, c game.CharacterCardDto) int {
	alternativeIndex := -1

	for i, cc := range playerInstance.Characters {
		if c.Type == cc.Class.ToDto() {
			alternativeIndex = i
			break
		}
	}

	return alternativeIndex
}
