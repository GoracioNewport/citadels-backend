package game

import (
	"citadels-backend/logic/game/ability"
	"citadels-backend/logic/game/ability/enums"
	"citadels-backend/logic/game/character"
	"fmt"
	"log"
)

type TargetedAbilityActivator func(g *Game, c *character.Character, p *Player, a *ability.Ability, ctx ability.TargetContext)

func ActivateAssassinKillCharacter(g *Game, c *character.Character, p *Player, a *ability.Ability, ctx ability.TargetContext) {
	if ctx.TargetArea != ability.Character {
		log.Printf("Invalid target area for ability %s", a.Key)
		return
	}

	target, ok := character.ClassFromString(ctx.TargetValue)

	if !ok {
		log.Printf("Invalid target value for ability %s", a.Key)
		return
	}

	characterReference := g.GetCharacter(target)

	if characterReference == nil {
		log.Printf("Character %s not found", target)
		return
	}

	characterReference.Dead = true

	g.BroadcastServerMessage(fmt.Sprintf("Персонаж %s убит!", characterReference.Name))

	a.Active = false
}

func ActivateThiefLootCharacter(g *Game, c *character.Character, p *Player, a *ability.Ability, ctx ability.TargetContext) {
	if ctx.TargetArea != ability.Character {
		log.Printf("Invalid target area for ability %s", a.Key)
		return
	}

	target, ok := character.ClassFromString(ctx.TargetValue)

	if !ok {
		log.Printf("Invalid target value for ability %s", a.Key)
		return
	}

	characterReference := g.GetCharacter(target)

	if characterReference == nil {
		log.Printf("Character %s not found", target.ToDto())
		return
	}

	characterReference.Robbed = true

	g.BroadcastServerMessage(fmt.Sprintf("Персонаж %s ограблен!", characterReference.Name))

	a.Active = false
}

func ActivateMagicianSwapCards(g *Game, c *character.Character, p *Player, a *ability.Ability, ctx ability.TargetContext) {
	if ctx.TargetArea != ability.Player {
		log.Printf("Invalid target area for ability %s", a.Key)
		return
	}

	target := ctx.TargetValue

	playerReference := g.GetPlayer(target)

	if playerReference == nil {
		log.Printf("Player %s not found", target)
		return
	}

	playerReference.Hand, p.Hand = p.Hand, playerReference.Hand

	g.BroadcastServerMessage(fmt.Sprintf("Магическим образом карты игроков %s и %s меняются местами!", p.Name, playerReference.Name))

	a.Active = false
}

var TargetedAbilityFunctions = map[enums.Key]TargetedAbilityActivator{
	enums.AssassinKillCharacterKey: ActivateAssassinKillCharacter,
	enums.ThiefLootCharacterKey:    ActivateThiefLootCharacter,
	enums.MagicianSwapCardsKey:     ActivateMagicianSwapCards,
}

func GetTargetedAbilityFunction(key enums.Key) TargetedAbilityActivator {
	if _, ok := TargetedAbilityFunctions[key]; !ok {
		log.Printf("Targeted ability function for key %s not found", key)
		return nil
	}

	return TargetedAbilityFunctions[key]
}
