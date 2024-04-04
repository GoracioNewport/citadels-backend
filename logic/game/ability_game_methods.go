package game

import (
	"citadels-backend/logic/game/ability"
	"citadels-backend/logic/game/ability/enums"
)

func (g *Game) ActivateAbility(player *Player, key enums.Key) bool {
	if player == nil || player.Characters == nil {
		return false
	}

	for _, c := range player.Characters {
		if c.Class == g.GetCurrentCharacter().Class && c.CheckAbility(key) {
			ability := c.GetAbility(key)

			if !ability.Active {
				continue
			}

			if ability.ActivationType == enums.Instant {
				abilityFunction := GetInstantAbilityFunction(key)
				if abilityFunction == nil {
					return false
				}

				abilityFunction(g, c, player, ability)

				g.BroadcastState()
				return true
			}

			if ability.ActivationType == enums.Targeted {
				if c.PendingTargetedAbility == ability {
					c.PendingTargetedAbility = nil
				} else {
					c.PendingTargetedAbility = ability
				}

				g.BroadcastState()
				return true
			}
		}
	}

	return false
}

func (g *Game) TargetAbility(player *Player, ctx ability.TargetContext) bool {
	if player == nil || player.Characters == nil {
		return false
	}

	for _, c := range player.Characters {
		if c.PendingTargetedAbility != nil {
			if !c.PendingTargetedAbility.Active {
				return false
			}

			if c.PendingTargetedAbility.ActivationType != enums.Targeted {
				return false
			}

			abilityFunction := GetTargetedAbilityFunction(c.PendingTargetedAbility.Key)
			if abilityFunction == nil {
				return false
			}

			abilityFunction(g, c, player, c.PendingTargetedAbility, ctx)

			c.PendingTargetedAbility = nil
			return true
		}
	}

	return false
}
