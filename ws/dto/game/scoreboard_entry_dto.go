package game

type ScoreboardEntryDto struct {
	Name       string `json:"name"`
	TotalScore int    `json:"score"`
	Source     string `json:"source"`
}
