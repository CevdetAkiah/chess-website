package gameserver

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"nhooyr.io/websocket"
)

type GameServer struct {
	log *custom_log.Logger
}

func NewGameServer() actor.Receiver {
	l := custom_log.NewLogger()
	return &GameServer{log: l}
}

// actor receives messages to process
func (server *GameServer) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		server.startHTTP()
		_ = msg
	}
}

// upgrade the websocket connection
func (server *GameServer) handleWS(w http.ResponseWriter, r *http.Request) {
	originPatterns := []string{"localhost:3000"}
	acceptOptions := websocket.AcceptOptions{
		OriginPatterns: originPatterns,
	}
	conn, err := websocket.Accept(w, r, &acceptOptions)
	if err != nil {
		server.log.Error(err)
	}
	defer conn.Close(websocket.StatusInternalError, "websocket closing")

	fmt.Println("new connection: ", conn)
}

func (server *GameServer) startHTTP() {
	fmt.Println("Starting game server on port: 4000")
	go func() {
		http.HandleFunc("/ws", server.handleWS)
		http.ListenAndServe(":4000", nil)
	}()

}
