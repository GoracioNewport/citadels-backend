package game

type StateDto struct {
	Players          []PlayerCommonInfoDto `json:"players"`
	Characters       []CharacterInfoDto    `json:"characters"`
	Player           PlayerInfoDto         `json:"player"`
	MainDeckSize     int                   `json:"main_deck_size"`
	MainDeckActive   bool                  `json:"main_deck_active"`
	GlobalBankActive bool                  `json:"global_bank_active"`
	EndTurnActive    bool                  `json:"end_turn_active"`
	Stage            Stage                 `json:"stage"`
}
