package chesswebsocket

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

// return either white or black
func randColour() string {
	rand.NewSource(time.Now().UnixNano())
	colours := []string{white, black}
	return colours[rand.Intn(len(colours))]
}

// create a player from the user
func newPlayer(user *Player, wsc *websocket.Conn) *Player {
	player := &Player{}
	player.PlayerID = wsc
	player.Name = user.Name
	return player
}

// encode messages to json to disperse to websocket connections
func encodeMessage(sm interface{}) []byte {
	msg, err := json.Marshal(sm)
	if err != nil {
		fmt.Println("encoding error: ", err)
	}
	return msg
}

func newGameID() string {
	gameID := time.Now().String() + uuid.New().String()
	return gameID
}
