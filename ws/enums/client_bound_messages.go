package enums

import (
	"citadels-backend/ws/dto/game"
	"citadels-backend/ws/dto/lobby"
	"encoding/json"
)

type ClientBoundMessage string

const (
	StatusCode       ClientBoundMessage = "status_code"
	LobbyCreated     ClientBoundMessage = "lobby_created"
	LobbyInfo        ClientBoundMessage = "lobby_info"
	LobbyKicked      ClientBoundMessage = "lobby_kicked"
	LobbyGameStarted ClientBoundMessage = "lobby_game_started"

	GameInfo        ClientBoundMessage = "game_info"
	GameChat        ClientBoundMessage = "game_chat"
	GameChooseCards ClientBoundMessage = "game_choose_cards"

	GameFinalScores ClientBoundMessage = "game_final_scores"
)

type GenericClientBoundMessage struct {
	Type    ClientBoundMessage `json:"type"`
	Payload json.RawMessage    `json:"payload"`
}

type StatusCodeMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type LobbyCreatedMessage struct {
	Code string `json:"code"`
}

type LobbyInfoMessage struct {
	State lobby.StateDto `json:"state"`
}

type LobbyKickedMessage struct{}

type LobbyStartedMessage struct{}

type GameInfoMessage struct {
	State game.StateDto `json:"state"`
}

type GameChatMessage struct {
	Message game.ChatMessageDto `json:"message"`
}

type GameChooseCardsMessage struct {
	Cards []game.CardDto `json:"cards"`
}

type GameFinalScoresMessage struct {
	Scores []game.ScoreboardEntryDto `json:"scores"`
}
