package gameserver

import (
	"encoding/json"
	"fmt"
	custom_log "go-projects/chess/logger"
	"math/rand"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type PlayerSession struct {
	log       *custom_log.Logger
	sessionID int
	clientID  string
	username  string
	colour    colour
	inLobby   bool
	state     *PlayerState
	conn      *websocket.Conn
	ctx       *actor.Context
	serverPID *actor.PID
}

// create a new session actor
func newPlayerSession(serverPID *actor.PID, sid int, availableColours map[int]colour, conn *websocket.Conn) actor.Producer {
	// select colour at random
	rand.NewSource(time.Now().UnixNano())
	index := rand.Intn(len(availableColours))
	if len(availableColours) == 1 && availableColours[0] == "" {
		index = 1
	}
	colour := availableColours[index]
	// remove colour from the list
	delete(availableColours, index)

	return func() actor.Receiver {
		return &PlayerSession{
			log:       custom_log.NewLogger(),
			sessionID: sid,
			colour:    colour,
			conn:      conn,
			serverPID: serverPID,
		}
	}
}

// the session is started
func (session *PlayerSession) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		session.ctx = c
		go session.readLoop()
	case *PlayerState:
		session.sendState(msg)
	case StartGame:
		start, err := newOutGoingMessage(startGame, nil)
		if err != nil {
			panic(err)
		}
		session.handleMsg(start)
	case actor.Stopped:
		session.conn.Close()
	}
}

// send updated state to client. Received from opposing player
func (session *PlayerSession) sendState(state *PlayerState) {
	session.state = state
	msg, err := newOutGoingMessage(makeMove, session)
	if err != nil {
		session.log.Infof("from sendState: %b", err)
		panic(err)
	}
	session.conn.WriteJSON(msg)
}

// playersession will read any incoming messages from the websocket
func (session *PlayerSession) readLoop() {
	msg := newIncomingMessage()
	for {
		if err := session.conn.ReadJSON(&msg); err != nil {
			session.log.Infof("read websocket error: %b\n ", err)
			panic(err) // server will restart the session actor
		}
		go session.handleMsg(msg)
	}
}

// playersession will handle incoming message from the websocket client
func (session *PlayerSession) handleMsg(msg Message) {
	switch msg.MessageType() {
	case "join":
		var incomingJoin joinMsg
		if err := json.Unmarshal(msg.(incomingMessage).Data, &incomingJoin); err != nil {
			session.log.Infof("problem unmarshalling join message: %b", err)
			panic(err) // can panic here as the server will restart the playersession
		}
		session.username = incomingJoin.Username
		session.clientID = incomingJoin.ID
		session.handleJoin()
	case makeMove:
		var incomingMove moveMessage
		if err := json.Unmarshal(msg.(incomingMessage).Data, &incomingMove); err != nil {
			session.log.Infof("error while reading playerstate from websocket: %b", err)
			panic(err) // server will restart the session actor
		}
		state := newPlayerState(incomingMove)
		state.sessionID = session.sessionID
		session.state = state
		if session.ctx != nil {
			session.ctx.Send(session.serverPID, state) // send the state to the gameserver to send to the other player
		}
	case startGame:
		session.startGame(msg)
	}
}

func (session *PlayerSession) startGame(msg Message) {
	session.conn.WriteJSON(msg)
}

func (session *PlayerSession) handleJoin() {
	if session.clientID == "new-game" {
		session.inLobby = true
		msg, err := newOutGoingMessage(joinConfirmation, session)
		if err != nil {
			session.log.Infof("handleJoin error: %b", err)
			panic(err)
		}

		if err := session.conn.WriteJSON(msg); err != nil {
			session.log.Infof("error while writing JSON in handleJoin")
			panic(err)
		}
		fmt.Println("sending playerJoined message: ", msg)
	}

}

// generates a new game ID
func newGameID() string {
	gameID := time.Now().String() + uuid.New().String()
	return gameID
}
