package game

type ScoreboardEntryDto struct {
	Name             string             `json:"name"`
	TotalScore       int                `json:"total_score"`
	ScoreDescription []ScoreDescription `json:"score_description"`
}

type ScoreDescription struct {
	Score int    `json:"score"`
	Text  string `json:"text"`
}
