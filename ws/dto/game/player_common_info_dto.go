package game

type PlayerCommonInfoDto struct {
	Name             string             `json:"name"`
	Coins            int                `json:"coins"`
	HandDeckSize     int                `json:"hand_deck_size"`
	Characters       []CharacterCardDto `json:"characters"`
	CharactersActive []bool             `json:"characters_active"`
	Crown            bool               `json:"crown"`
	Town             []CardDto          `json:"town"`
}
