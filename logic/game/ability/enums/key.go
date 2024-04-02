package enums

type Key string

const (
	BaseEndTurnKey           Key = "base_end_turn"
	BaseLootBankKey          Key = "base_loot_bank"
	BaseDrawCardsKey         Key = "base_draw_cards"
	BaseConstructBuildingKey Key = "base_construct_building"
	KingLootCityKey          Key = "king_loot_city"
	BishopLootCityKey        Key = "bishop_loot_city"
	WarlordLootCityKey       Key = "warlord_loot_city"
	MerchantLootCityKey      Key = "merchant_loot_city"
	MerchantLootBankKey      Key = "merchant_loot_bank"
)

func (k Key) ToDto() string {
	return string(k)
}

func KeyFromDto(dto string) (key Key, ok bool) {
	abilityCollection := map[Key]struct{}{
		BaseEndTurnKey:           {},
		BaseLootBankKey:          {},
		BaseDrawCardsKey:         {},
		BaseConstructBuildingKey: {},
		KingLootCityKey:          {},
		BishopLootCityKey:        {},
		WarlordLootCityKey:       {},
		MerchantLootCityKey:      {},
		MerchantLootBankKey:      {},
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
	case BaseConstructBuildingKey:
		return "Построить здание"
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
	}

	return "Неизвестная способность"
}
