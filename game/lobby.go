package game

import (
	"crypto/rand"
	"encoding/hex"
)

type Lobby struct {
	player1 *Client
	player2 *Client
	game    Minesweeper
	active  bool
	code    string
}

func NewLobby(width int, height int, minec int) *Lobby {
	return &Lobby{
		game: *NewSweeper(width, height, minec),
		code: CreateLobbyCode(),
	}
}

func (l *Lobby) AddPlayer(client *Client) {
	if l.player1 == nil {
		l.player1 = client
	} else if l.player2 == nil {
		l.player2 = client
	}
	client.Lobby = l

	if l.player1 != nil && l.player2 != nil {
		l.player1.turn = true
		l.player1.SendTurn()
	}
}

func (l *Lobby) RemovePlayer(client *Client) {
	if l.player1 == client {
		l.player1 = nil
	} else if l.player2 == client {
		l.player2 = nil
	}
}

func CreateLobbyCode() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (l *Lobby) TileClick(pos int, client *Client) {
	if (!client.turn) && l.active {
		return
	}

	if l.player1 == client {
		l.player1.turn = false
		l.player2.turn = true
	} else {
		l.player1.turn = true
		l.player2.turn = false
	}

	mine_cliecked := l.game.Reveal(pos)
	board_completed := l.game.CheckWin()

	cmap := l.game.MapClient()
	l.player1.SendMap(&cmap)
	l.player2.SendMap(&cmap)

	if !mine_cliecked && !board_completed {
		l.player1.SendTurn()
		l.player2.SendTurn()
	} else if mine_cliecked {
		if l.player1 == client {
			l.player1.SendGameEnd(false)
			l.player2.SendGameEnd(true)
		} else {
			l.player1.SendGameEnd(true)
			l.player2.SendGameEnd(false)
		}
	} else if board_completed {
		if l.player1 == client {
			l.player1.SendGameEnd(true)
			l.player2.SendGameEnd(false)
		} else {
			l.player1.SendGameEnd(false)
			l.player2.SendGameEnd(true)
		}
	}
}
