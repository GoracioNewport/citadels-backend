package character

import (
	"citadels-backend/logic/game/ability"
)

func NewCharacters(includeKing bool) []*Character {
	characters := []*Character{
		NewAssassin(),
		NewThief(),
		NewMagician(),
		NewBishop(),
		NewMerchant(),
		NewArchitect(),
		NewWarlord(),
	}

	if includeKing {
		characters = append(characters, NewKing())
	}

	return characters
}

func NewAssassin() *Character {
	return &Character{
		Class: Assassin,
		Order: 1,
		Abilities: append(
			ability.GenerateAbilities(ability.BaseAbilities),
			ability.GenerateAbilities(ability.AssassinAbilities)...,
		),
		Name: "Ассасин",
		Description: []string{
			"Можете назвать персожана, которого хотите убить. Когда его вызывают, его владелец молчит. Ход убитого персонажа пропускается.",
		},
		Image: "Assassin.png",
	}
}

func NewThief() *Character {
	return &Character{
		Class: Thief,
		Order: 2,
		Abilities: append(
			ability.GenerateAbilities(ability.BaseAbilities),
			ability.GenerateAbilities(ability.ThiefAbilities)...,
		),
		Name: "Вор",
		Description: []string{
			"Можете назвать персонажа, которого хотите обворовать (кроме персонажей 1-го ранга, а также убитых и заколдованных). Когда этого пресонажа вызывают, заберите у его владельца всё золото.",
		},
		Image: "Thief.png",
	}
}

func NewMagician() *Character {
	return &Character{
		Class: Magician,
		Order: 3,
		Abilities: append(
			ability.GenerateAbilities(ability.BaseAbilities),
			ability.GenerateAbilities(ability.MagicianAbilities)...,
		),
		Name: "Чародей",
		Description: []string{
			"Можете либо обменять все карты с руки на карты с руки другого игрока, либо вернуть сколько угодно карт с руки лицевой стороной вниз под колоду кварталов и взять из неё столько же новых.",
		},
		Image: "Magician.png",
	}
}

func NewKing() *Character {
	return &Character{
		Class: King,
		Order: 4,
		Abilities: append(
			ability.GenerateAbilities(ability.BaseAbilities),
			ability.GenerateAbilities(ability.KingAbilities)...,
		),
		Name: "Король",
		Description: []string{
			"В свой ход вы должны взять корону.",
			"Получите 1 <b>золотой</b> за каждый ваш дворянский квартал.",
		},
		Image: "King.png",
	}
}

func NewBishop() *Character {
	return &Character{
		Class: Bishop,
		Order: 5,
		Abilities: append(
			ability.GenerateAbilities(ability.BaseAbilities),
			ability.GenerateAbilities(ability.BishopAbilities)...,
		),
		Name: "Епископ",
		Description: []string{
			"Персонажи 8-го ранга не могут применять своё особое свойство против ваших кварталов.",
			"Получите 1 <b>золотой</b> за каждый ваш церковный квартал.",
		},
		Image: "Bishop.png",
	}
}

func NewMerchant() *Character {
	return &Character{
		Class: Merchant,
		Order: 6,
		Abilities: append(
			ability.GenerateAbilities(ability.BaseAbilities),
			ability.GenerateAbilities(ability.MerchantAbilities)...,
		),
		Name: "Купец",
		Description: []string{
			"Получите 1 дополнительный золотой.",
			"Получите 1 <b>золотой</b> за каждый ваш торговый квартал.",
		},
		Image: "Merchant.png",
	}
}

func NewArchitect() *Character {
	return &Character{
		Class: Architect,
		Order: 7,
		Abilities: append(
			ability.GenerateAbilities(ability.BaseAbilities),
			ability.GenerateAbilities(ability.ArchitectAbilities)...,
		),
		Name: "Зодчий",
		Description: []string{
			"Возьмите 2 дополнительные карты из колоды кварталов.",
			"Можете построить 3 квартала или меньше.",
		},
		Image: "Architect.png",
	}
}

func NewWarlord() *Character {
	return &Character{
		Class: Warlord,
		Order: 8,
		Abilities: append(
			ability.GenerateAbilities(ability.BaseAbilities),
			ability.GenerateAbilities(ability.WarlordAbilities)...,
		),
		Name: "Кондотьер",
		Description: []string{
			"Можете разрушить 1 квартал по вашему выбору, заплатив в банк на 1 золотой меньше стоимости квартала. Нельзя разрушать кварталы в достроенном городе.",
			"Получите 1 <b>золотой</b> за каждый ваш воинский квартал.",
		},
		Image: "Warlord.png",
	}
}
