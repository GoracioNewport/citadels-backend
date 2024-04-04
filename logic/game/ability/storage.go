package ability

import (
	"citadels-backend/logic/game/ability/enums"
)

// Base

var BaseAbilities = []Generator{
	GenerateBaseLootBank,
	GenerateBaseDrawCards,
	GenerateEndTurn,
}

func GenerateBaseLootBank() Ability {
	return NewAbility(enums.BaseLootBankKey, enums.Instant, true, true)
}
func GenerateBaseDrawCards() Ability {
	return NewAbility(enums.BaseDrawCardsKey, enums.Instant, true, true)
}
func GenerateEndTurn() Ability { return NewAbility(enums.BaseEndTurnKey, enums.Instant, false, true) }

// Assassin

var AssassinAbilities = []Generator{
	GenerateAssassinKillCharacter,
}

func GenerateAssassinKillCharacter() Ability {
	return NewAbility(enums.AssassinKillCharacterKey, enums.Targeted, true, false)
}

// Thief

var ThiefAbilities = []Generator{
	GenerateThiefLootCharacter,
}

func GenerateThiefLootCharacter() Ability {
	return NewAbility(enums.ThiefLootCharacterKey, enums.Targeted, true, false)
}

// Magician

var MagicianAbilities = []Generator{
	GenerateMagicianSwapCards,
}

func GenerateMagicianSwapCards() Ability {
	return NewAbility(enums.MagicianSwapCardsKey, enums.Targeted, true, false)
}

// King

var KingAbilities = []Generator{
	GenerateKingLootCity,
}

func GenerateKingLootCity() Ability {
	return NewAbility(enums.KingLootCityKey, enums.Instant, true, false)
}

// Bishop

var BishopAbilities = []Generator{
	GenerateBishopLootCity,
}

func GenerateBishopLootCity() Ability {
	return NewAbility(enums.BishopLootCityKey, enums.Instant, true, false)
}

// Merchant

var MerchantAbilities = []Generator{
	GenerateMerchantLootCity,
	GenerateMerchantLootBank,
}

func GenerateMerchantLootCity() Ability {
	return NewAbility(enums.MerchantLootCityKey, enums.Instant, true, false)
}

func GenerateMerchantLootBank() Ability {
	return NewAbility(enums.MerchantLootBankKey, enums.Instant, true, false)
}

// Architect

var ArchitectAbilities = []Generator{
	GenerateArchitectLootDeck,
}

func GenerateArchitectLootDeck() Ability {
	return NewAbility(enums.ArchitectLootDeckKey, enums.Instant, true, false)
}

// Warlord

var WarlordAbilities = []Generator{
	GenerateWarlordLootCity,
}

func GenerateWarlordLootCity() Ability {
	return NewAbility(enums.WarlordLootCityKey, enums.Instant, true, false)
}
