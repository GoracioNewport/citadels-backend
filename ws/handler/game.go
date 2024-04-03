package handler

import (
	"citadels-backend/logic/game"
	abilityEnums "citadels-backend/logic/game/ability/enums"
	"citadels-backend/logic/game/character"
	"citadels-backend/ws/api"
	"citadels-backend/ws/enums"
	"citadels-backend/ws/sender"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func HandleGamePickCharacter(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.GamePickCharacterMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	gameInstance := game.ManagerInstance.GetGame(payload.Code)
	if gameInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Game not found")
		return
	}

	class, ok := character.ClassFromString(payload.Class)

	if !ok {
		sender.SendStatusMessage(ws, server, 404, "No such class")
		return
	}

	playerInstance := gameInstance.GetPlayer(payload.Name)
	if playerInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Player not found")
		return
	}

	ok = gameInstance.PickCharacter(playerInstance, class)

	if !ok {
		sender.SendStatusMessage(ws, server, 500, "Cannot pick character")
		return
	}

	sender.SendStatusMessage(ws, server, 200, "Character picked")
}

func HandleGameActivateAbility(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.GameActivateAbilityMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	gameInstance := game.ManagerInstance.GetGame(payload.Code)
	if gameInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Game not found")
		return
	}

	abilityKey, ok := abilityEnums.KeyFromDto(payload.AbilityKey)

	if !ok {
		sender.SendStatusMessage(ws, server, 404, "No such ability")
		return
	}

	playerInstance := gameInstance.GetPlayer(payload.Name)
	if playerInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Player not found")
		return
	}

	ok = gameInstance.ActivateAbility(playerInstance, abilityKey)

	if !ok {
		sender.SendStatusMessage(ws, server, 404, "Cannot activate ability")
		return
	}

	sender.SendStatusMessage(ws, server, 200, "Ability activated")
}

func HandleGameChooseCard(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.GameChooseCardMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	gameInstance := game.ManagerInstance.GetGame(payload.Code)
	if gameInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Game not found")
		return
	}

	playerInstance := gameInstance.GetPlayer(payload.Name)
	if playerInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Player not found")
		return
	}

	ok := playerInstance.TakeCardFromSelectionBuffer(payload.Id)

	if !ok {
		sender.SendStatusMessage(ws, server, 404, "Card not found in selection buffer")
		return
	}

	sender.SendStatusMessage(ws, server, 200, "Card chosen")

	gameInstance.ClearSelectionBuffer(playerInstance, payload.Id)
	gameInstance.BroadcastState()
}

func HandleGameConstructBuildingAbility(ws *websocket.Conn, server *api.WebSocketServer, message enums.GenericServerBoundMessage) {
	var payload enums.GameConstructBuildingAbilityMessage
	var err = json.Unmarshal(message.Payload, &payload)

	if err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	gameInstance := game.ManagerInstance.GetGame(payload.Code)
	if gameInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Game not found")
		return
	}

	playerInstance := gameInstance.GetPlayer(payload.Name)
	if playerInstance == nil {
		sender.SendStatusMessage(ws, server, 404, "Player not found")
		return
	}

	ok := playerInstance.ConstructBuilding(payload.Id)

	if !ok {
		sender.SendStatusMessage(ws, server, 404, "Cannot construct building")
		return
	}

	sender.SendStatusMessage(ws, server, 200, "Building constructed")
	gameInstance.BroadcastState()
}
