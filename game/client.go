package game

import (
	"encoding/json"

	"golang.org/x/net/websocket"
)

type Client struct {
	sock   *websocket.Conn
	Lobby  *Lobby
	turn   bool
	active bool
}

func NewClient(soket *websocket.Conn) *Client {
	return &Client{
		sock:   soket,
		Lobby:  &Lobby{},
		turn:   false,
		active: true,
	}
}

func (c *Client) SendMap(m *[]byte) {
	p, _ := json.Marshal(EventMapUpdated{
		Id:  "map_changed",
		Map: *m,
	})
	c.sock.Write(p)
}

func (c *Client) SendTurn() {
	p, _ := json.Marshal(EventTurnChanged{
		Id:   "turn_changed",
		Turn: c.turn,
	})
	c.sock.Write(p)
}

func (c *Client) SendLobby() {
	p, _ := json.Marshal(EventLobbyCreated{
		Id:     "lobby_created",
		Width:  byte(c.Lobby.game.Width),
		Height: byte(c.Lobby.game.height),
		Mines:  c.Lobby.game.mines,
		Map:    c.Lobby.game.MapClient(),
		Code:   c.Lobby.code,
	})
	c.sock.Write(p)
}

func (c *Client) SendGameEnd(win bool) {
	p, _ := json.Marshal(EventGameEnd{
		Id:    "game_end",
		State: win,
	})
	c.sock.Write(p)
}
