package game

type PlayerInfoDto struct {
	Bank       int                `json:"bank"`
	Crown      bool               `json:"crown"`
	Town       []CardDto          `json:"town"`
	Hand       []CardDto          `json:"hand"`
	Characters []CharacterCardDto `json:"characters"`
}
