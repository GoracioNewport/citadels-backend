package game

import (
	"citadels-backend/ws/dto/game"
	"citadels-backend/ws/sender"
	"sort"
)

func (g *Game) endGame() {
	g.BroadcastServerMessage("Подсчитываем очки...")
	g.stage = EndGame

	g.BroadcastState()

	scores := make([]game.ScoreboardEntryDto, 0)

	for _, p := range g.players {
		score := calculateScore(g, p)

		scores = append(scores, game.ScoreboardEntryDto{
			Name:  p.Name,
			Score: score,
		})
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i].Score > scores[j].Score })

	for _, p := range g.players {
		sender.SendGameFinalScoresMessage(p.Connection, ManagerInstance.Server, scores)
	}

	g.BroadcastServerMessage("Игра окончена! Победил " + scores[0].Name + "!")
}

func calculateScore(g *Game, p *Player) int {
	score := 0

	for _, b := range p.Town {
		score += b.Price
	}

	if len(p.Town) >= townBuildingWinCondition {
		score += 2
	}

	if p == g.finishedPlayer {
		score += 2
	}

	return score
}
