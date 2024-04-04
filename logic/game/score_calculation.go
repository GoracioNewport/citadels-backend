package game

import (
	"citadels-backend/logic/game/card/template"
	"citadels-backend/ws/dto/game"
	"sort"
)

func (g *Game) calculateFinalScores() []game.ScoreboardEntryDto {
	scores := make([]game.ScoreboardEntryDto, 0)

	for _, p := range g.players {
		scores = append(scores, calculateScore(g, p))
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i].TotalScore > scores[j].TotalScore })

	return scores
}

func calculateScore(g *Game, p *Player) game.ScoreboardEntryDto {
	scores := make([]game.ScoreDescription, 0)

	townScore := 0
	for _, b := range p.Town {
		townScore += b.Price
	}
	scores = append(scores, game.ScoreDescription{
		Score: townScore,
		Text:  "Суммарная стоимость построенных кварталов",
	})

	if len(p.Town) >= townBuildingWinCondition {
		if p == g.finishedPlayer {
			scores = append(scores, game.ScoreDescription{
				Score: 4,
				Text:  "Первый завершённый город",
			})
		} else {
			scores = append(scores, game.ScoreDescription{
				Score: 2,
				Text:  "Завершённый город",
			})
		}
	}

	colors := make(map[template.BuildingColor]int)

	for _, b := range p.Town {
		colors[b.Color]++
	}

	if len(colors) == template.ColorAmount {
		scores = append(scores, game.ScoreDescription{
			Score: 3,
			Text:  "Построены кварталы всех видов",
		})
	}

	totalScore := 0
	for _, s := range scores {
		totalScore += s.Score
	}

	return game.ScoreboardEntryDto{
		Name:             p.Name,
		TotalScore:       totalScore,
		ScoreDescription: scores,
	}
}
