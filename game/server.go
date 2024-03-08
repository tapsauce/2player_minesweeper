package game

import (
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns   map[*websocket.Conn]*Client
	lobbies map[string]*Lobby
}

func NewServer() *Server {
	return &Server{
		conns:   make(map[*websocket.Conn]*Client),
		lobbies: make(map[string]*Lobby),
	}
}

func (s *Server) Wshandler(ws *websocket.Conn) {
	s.ClientAdded(ws)
	s.read(ws)
}

func (s *Server) read(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("uh oh!")
			continue
		}
		client := s.conns[ws]
		s.HandleMsg(buf[:n], client)
	}
	s.ClientDisconnect(ws)
}

func (s *Server) ClientAdded(ws *websocket.Conn) {
	s.conns[ws] = NewClient(ws)
	fmt.Println("connection from: ", ws.RemoteAddr())
}

func (s *Server) ClientDisconnect(ws *websocket.Conn) {
	delete(s.conns, ws)
	fmt.Println("disconnected: ", ws.RemoteAddr())
}

func (s *Server) HandleMsg(msg []byte, client *Client) {
	message := Message{}
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return
	}
	switch message.Id {
	case "join_lobby":
		s.JoinLobby(message.Msg, client)
	case "create_lobby":
		s.CreateLobby(message.Msg, client)
	case "tile_clicked":
		s.TileClicked(message.Msg, client)
	}
}

func (s *Server) JoinLobby(data []byte, client *Client) {
	msg := MessageLobbyJoin{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return
	}
	lobby, ok := s.lobbies[msg.Code]
	if ok {
		lobby.AddPlayer(client)
		delete(s.lobbies, msg.Code)
		client.SendLobby()
		client.SendTurn()
	}
}

func (s *Server) CreateLobby(data []byte, client *Client) {
	if !client.Lobby.active {
		msg := MessageLobbyCreate{}
		err := json.Unmarshal(data, &msg)
		if err != nil {
			return
		}
		if msg.Mines > int(msg.Width)*int(msg.Height) {
			return
		}
		lobby := NewLobby(int(msg.Width), int(msg.Height), int(msg.Mines))
		lobby.AddPlayer(client)
		client.Lobby = lobby
		s.lobbies[lobby.code] = lobby
		client.SendLobby()
	}
}

func (s *Server) TileClicked(data []byte, client *Client) {
	msg := MessageTileClicked{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return
	}
	client.Lobby.TileClick(msg.Tile, client)
}
