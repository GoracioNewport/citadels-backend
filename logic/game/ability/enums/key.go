package enums

type Key string

const (
	BaseEndTurnKey       Key = "base_end_turn"
	BaseLootBankKey      Key = "base_loot_bank"
	BaseDrawCardsKey     Key = "base_draw_cards"
	KingLootCityKey      Key = "king_loot_city"
	BishopLootCityKey    Key = "bishop_loot_city"
	WarlordLootCityKey   Key = "warlord_loot_city"
	MerchantLootCityKey  Key = "merchant_loot_city"
	MerchantLootBankKey  Key = "merchant_loot_bank"
	ArchitectLootDeckKey Key = "architect_loot_deck"
)

func (k Key) ToDto() string {
	return string(k)
}

func KeyFromDto(dto string) (key Key, ok bool) {
	abilityCollection := map[Key]struct{}{
		BaseEndTurnKey:       {},
		BaseLootBankKey:      {},
		BaseDrawCardsKey:     {},
		KingLootCityKey:      {},
		BishopLootCityKey:    {},
		WarlordLootCityKey:   {},
		MerchantLootCityKey:  {},
		MerchantLootBankKey:  {},
		ArchitectLootDeckKey: {},
	}

	ability := Key(dto)
	_, okMap := abilityCollection[ability]

	return ability, okMap
}

func (k Key) GetName() string {
	switch k {
	case BaseLootBankKey:
		return "Взять монеты из банка"
	case BaseDrawCardsKey:
		return "Взять карты из колоды"
	case KingLootCityKey:
		return "Собрать налоги"
	case BishopLootCityKey:
		return "Собрать налоги"
	case WarlordLootCityKey:
		return "Собрать налоги"
	case MerchantLootCityKey:
		return "Собрать налоги"
	case MerchantLootBankKey:
		return "+1 монета"
	case ArchitectLootDeckKey:
		return "+2 карты"
	}

	return "Неизвестная способность"
}
