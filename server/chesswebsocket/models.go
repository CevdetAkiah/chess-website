package chesswebsocket

import (
	"go-projects/chess/service"
	"sync"

	"golang.org/x/net/websocket"
)

var (
	emitMsg            = &emitMessage{Emit: "message"}
	emitPlayerJoined   = &emitMessage{Emit: "playerJoined"}
	emitOpponentJoined = &emitMessage{Emit: "opponentJoined"}
	emitOpponentMove   = &emitMessage{Emit: "opponentMove"}
)

const (
	white = "w"
	black = "b"
)

// handles websocket logic
type WsGame struct {
	lock        sync.Mutex
	conns       map[*websocket.Conn]bool
	gameSearch  []*Game
	gamesInPlay map[int]*Game
	DBAccess    service.DbService
}

type Player struct {
	Name     string `json:"name"`
	Colour   string
	PlayerID *websocket.Conn
	GameID   string
}

type Game struct {
	gameID    int
	playerOne *Player
	playerTwo *Player
}

type emitMessage struct {
	Emit string `json:"emit"`
}

// send message to client
type sendMessage struct {
	*emitMessage
	Message string `json:"message"`
}

// send chess move to client
type sendMove struct {
	*emitMessage
	FromMV string `json:"from"`
	ToMV   string `json:"to"`
}

// send player info to client
type sendPlayerInfo struct {
	*emitMessage
	PlayerName   string `json:"pname"`
	PlayerColour string `json:"playerColour"`
}

// send opponent info to client
type sendOpponentInfo struct {
	*emitMessage
	OpponentName   string `json:"opponentName"`
	OpponentColour string `json:"opponentColour"`
}

// receive message from client
type receiveMessage struct {
	Emit    string `json:"emit"`
	User    Player `json:"user"`
	Message string `json:"message"`
	GameID  int    `json:"gameID"`
	FromMV  string `json:"from"`
	ToMV    string `json:"to"`
}
