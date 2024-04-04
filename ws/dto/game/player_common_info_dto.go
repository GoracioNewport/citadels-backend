package game

type PlayerCommonInfoDto struct {
	Name         string             `json:"name"`
	Bank         int                `json:"bank"`
	HandDeckSize int                `json:"hand_deck_size"`
	Characters   []CharacterCardDto `json:"characters"`
	Crown        bool               `json:"crown"`
	Town         []CardDto          `json:"town"`
	Active       bool               `json:"active"`
}
