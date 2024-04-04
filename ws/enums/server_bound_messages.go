package enums

import (
	"citadels-backend/ws/dto/lobby"
	"encoding/json"
)

type ServerBoundMessage string

const (
	RestoreConnection ServerBoundMessage = "restore_connection"

	LobbyCreate       ServerBoundMessage = "lobby_create"
	LobbyJoin         ServerBoundMessage = "lobby_join"
	LobbyMemberUpdate ServerBoundMessage = "lobby_member_update"
	LobbyKick         ServerBoundMessage = "lobby_kick"
	LobbyLeave        ServerBoundMessage = "lobby_leave"
	LobbyStartGame    ServerBoundMessage = "lobby_start_game"

	GameActivateAbility          ServerBoundMessage = "game_activate_ability"
	GameTargetAbility            ServerBoundMessage = "game_target_ability"
	GameChooseCard               ServerBoundMessage = "game_choose_card"
	GameConstructBuildingAbility ServerBoundMessage = "game_construct_building_ability"
	GamePickCharacter            ServerBoundMessage = "game_pick_character"
)

type GenericServerBoundMessage struct {
	Type    ServerBoundMessage `json:"type"`
	Payload json.RawMessage    `json:"payload"`
}

type RestoreConnectionMessage struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type LobbyCreateMessage struct {
	Host string `json:"host"`
}

type LobbyJoinMessage struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type LobbyMemberUpdateMessage struct {
	Code   string          `json:"code"`
	Member lobby.MemberDto `json:"member"`
}

type LobbyKickMessage struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type LobbyStartGameMessage struct {
	Code string `json:"code"`
}

type GameActivateAbilityMessage struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	AbilityKey string `json:"ability_key"`
}

type GameChooseCardMessage struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type GameConstructBuildingAbilityMessage struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type GamePickCharacterMessage struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Class string `json:"class"`
}

type GameTargetAbilityMessage struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	TargetArea  string `json:"target_area"`
	TargetValue string `json:"target_value"`
}
