package game

type CharacterCardDto struct {
	Name        string                `json:"name"`
	Description []string              `json:"description"`
	Image       string                `json:"image"`
	Type        CharacterClassDto     `json:"class"`
	Active      bool                  `json:"active"`
	Abilities   []CharacterAbilityDto `json:"abilities"`
}

func GetUnknownCharacterCardDto() CharacterCardDto {
	return CharacterCardDto{
		Name:        "Неизвестный персонаж",
		Description: []string{"Персонаж этого игрока пока не раскрыт!"},
		Image:       "",
		Type:        Unknown,
		Abilities:   nil,
	}
}
