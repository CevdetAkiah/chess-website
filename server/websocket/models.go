package chesswebsocket

import (
	"go-projects/chess/service"
	"sync"

	"golang.org/x/net/websocket"
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

// TODO: experiment with getting rid of player and opponent and using a list  of players instead
type Game struct {
	gameID   int
	player   *Player
	opponent *Player
}

// send message to client
type sendMessage struct {
	Emit    string `json:"emit"`
	Message string `json:"message"`
}

// send chess move to client
type sendMove struct {
	Emit   string `json:"emit"`
	FromMV string `json:"from"`
	ToMV   string `json:"to"`
}

// send player info to client
type sendPlayerInfo struct {
	Emit         string `json:"emit"`
	PlayerName   string `json:"pname"`
	PlayerColour string `json:"playerColour"`
}

// send opponent info to client
type sendOpponentInfo struct {
	Emit           string `json:"emit"`
	OpponentName   string `json:"opponentName"`
	OpponentColour string `json:"opponentColour"`
}

// receive information from client
type receiveMessage struct {
	Emit    string `json:"emit"`
	User    Player `json:"user"`
	Message string `json:"message"`
	GameID  int    `json:"gameID"`
	FromMV  string `json:"from"`
	ToMV    string `json:"to"`
}