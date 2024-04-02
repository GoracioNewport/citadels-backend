package game

type CharacterInfoDto struct {
	Card   CharacterCardDto `json:"card"`
	Dead   bool             `json:"dead"`
	Robbed bool             `json:"robbed"`
	Absent bool             `json:"absent"`
	Turn   bool             `json:"turn"`
}
