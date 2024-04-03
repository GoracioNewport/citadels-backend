package game

import (
	"citadels-backend/logic/lobby"
	"citadels-backend/ws/api"
)

type Manager struct {
	games  map[string]*Game
	Server *api.WebSocketServer
}

var ManagerInstance Manager

func InitManager(server *api.WebSocketServer) {
	ManagerInstance = Manager{
		games:  make(map[string]*Game),
		Server: server,
	}
}

func (m *Manager) CreateGame(lobby lobby.Lobby) {
	game := &Game{
		code: lobby.Code,
	}

	var players = make([]*Player, 0)

	for _, mem := range lobby.Members {
		players = append(players, &Player{
			Name:       mem.Name,
			Connection: mem.Connection,
			Crown:      lobby.Host == mem.Name,
			Game:       game,
		})
	}

	game.players = players

	m.games[lobby.Code] = game
	m.games[lobby.Code].StartGame()
}

func (m *Manager) GetGame(code string) *Game {
	game, ok := m.games[code]
	if !ok {
		return nil
	}

	return game
}
