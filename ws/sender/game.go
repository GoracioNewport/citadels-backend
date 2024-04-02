package sender

import (
	"citadels-backend/ws/api"
	"citadels-backend/ws/dto/game"
	"citadels-backend/ws/enums"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func SendChatMessage(ws *websocket.Conn, server *api.WebSocketServer, message game.ChatMessageDto) {
	var payload, err = json.Marshal(enums.GameChatMessage{Message: message})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.GameChat,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}

func SendGameStateMessage(ws *websocket.Conn, server *api.WebSocketServer, state game.StateDto) {
	var payload, err = json.Marshal(enums.GameInfoMessage{State: state})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.GameInfo,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}

func SendGameChooseCardsMessage(ws *websocket.Conn, server *api.WebSocketServer, cards []game.CardDto) {
	var payload, err = json.Marshal(enums.GameChooseCardsMessage{Cards: cards})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.GameChooseCards,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}

func SendGameFinalScoresMessage(ws *websocket.Conn, server *api.WebSocketServer, scores []game.ScoreboardEntryDto) {
	var payload, err = json.Marshal(enums.GameFinalScoresMessage{Scores: scores})

	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	var data = enums.GenericClientBoundMessage{
		Type:    enums.GameFinalScores,
		Payload: json.RawMessage(payload),
	}

	server.MessageSender(ws, data)
}
