package gameserver

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"math"
	"math/rand"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/websocket"
)

const (
	port   = ":4000"
	origin = "http://localhost:3000"
)

type GameServer struct {
	log              *custom_log.Logger
	ctx              *actor.Context
	availableColours map[int]colour
	sessions         map[int]*actor.PID
}

func NewGameServer() actor.Receiver {
	l := custom_log.NewLogger()
	const white colour = "w"
	const black colour = "b"
	colours := make(map[int]colour)
	colours[0] = white
	colours[1] = black
	return &GameServer{
		log:              l,
		availableColours: colours,
		sessions:         make(map[int]*actor.PID),
	}
}

// actor receives messages to process
func (server *GameServer) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		server.startHTTP()
		server.ctx = c
	case *PlayerState:
		server.broadcast(c.Sender(), msg)
	default:
		fmt.Println("Received: ", msg)
	}

}

// broadcast the message to the other player
func (server *GameServer) broadcast(from *actor.PID, state *PlayerState) {
	for _, pid := range server.sessions { // send state to opposing player
		if !pid.Equals(from) {
			fmt.Println(pid)
			server.ctx.Send(pid, state) // this is received by the PlayerSession as *PlayerState
			// TODO: update DB with new fen
		}
	}
}

// listen for incoming websocket requests
func (server *GameServer) startHTTP() {
	fmt.Println("Starting game server on port", port)
	go func() {
		http.HandleFunc("/ws", server.handleWS)
		http.ListenAndServe(port, nil)
	}()
}

// upgrade the websocket connection
func (server *GameServer) handleWS(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{CheckOrigin: server.CheckOrigin}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		server.log.Error(err)
	}
	// create new player session for each incoming connection
	sid := rand.Intn(math.MaxInt)

	pid := server.ctx.SpawnChild(newPlayerSession(server.ctx.PID(), sid, server.availableColours, conn), fmt.Sprintf("session_%d", sid)) // spawn a new actor as a child of the player server
	server.sessions[sid] = pid

	// if there are two players in the server start the game
	if len(server.sessions) == 2 {
		server.startGame()
	}
	fmt.Printf("Client with sid %d and pid %s just connected \n", sid, pid)
}

// send the opponent info to each client so the game can start
func (server *GameServer) startGame() {
	for _, pid := range server.sessions {
		fmt.Println("server starting game")
		server.ctx.Send(pid, StartGame{})
	}
}

// checks the origin for the websocket upgrader
func (server *GameServer) CheckOrigin(r *http.Request) bool {
	requestOrigin := r.Header.Get("Origin")
	return requestOrigin == origin || requestOrigin == ""
}
