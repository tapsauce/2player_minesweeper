package game

import "encoding/json"

// -----------------------------------------------------------------------------------

type Message struct {
	Id  string          `json:"id"`
	Msg json.RawMessage `json:"data"`
}

type MessageLobbyCreate struct {
	Width  byte `json:"width"`
	Height byte `json:"height"`
	Mines  int  `json:"mines"`
}

type MessageTileClicked struct {
	Tile int `json:"offset"`
}

type MessageLobbyJoin struct {
	Code string `json:"code"`
}

type MessageMapReset struct { //TODO
	Reset bool `json:"reset"`
}

// -----------------------------------------------------------------------------------

type EventLobbyCreated struct {
	Id     string `json:"id"`
	Width  byte   `json:"width"`
	Height byte   `json:"height"`
	Mines  int    `json:"mines"`
	Map    []byte `json:"map"`
	Code   string `json:"code"`
}

type EventMapUpdated struct {
	Id  string `json:"id"`
	Map []byte `json:"map"`
}

type EventTurnChanged struct {
	Id   string `json:"id"`
	Turn bool   `json:"turn"`
}

type EventGameEnd struct {
	Id    string `json:"id"`
	State bool   `json:"victory"`
}
