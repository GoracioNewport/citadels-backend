package game

import (
	"citadels-backend/logic/game/enums"
	"citadels-backend/ws/sender"
)

func (g *Game) endGame() {
	g.BroadcastServerMessage("Подсчитываем очки...")
	g.stage = enums.EndGame

	g.BroadcastState()

	scores := g.calculateFinalScores()

	for _, p := range g.players {
		sender.SendGameFinalScoresMessage(p.Connection, ManagerInstance.Server, scores)
	}

	g.BroadcastServerMessage("Игра окончена! Победил " + scores[0].Name + "!")
}
