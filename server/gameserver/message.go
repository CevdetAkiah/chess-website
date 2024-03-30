package gameserver

import (
	"encoding/json"
	"fmt"
)

// expected message types
const (
	joinConfirmation = "playerJoined" // confirmation that the player has joined the server
	makeMove         = "playerState"  // move set from the client
	startGame        = "startGame"    // send signal to client to start the game
)

type colour string

// accepted message types
var validMessages = map[string]bool{
	joinConfirmation: true,
	makeMove:         true,
	startGame:        true,
}

type Message interface {
	Encode() ([]byte, error)
	MessageType() string
}

type StartGame struct{}

// use for incoming websocket messages
type incomingMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func newIncomingMessage() incomingMessage {
	return incomingMessage{}
}

func (msg incomingMessage) MessageType() string {
	return msg.Type
}
func (msg incomingMessage) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

// use to send back a message confirming that a player has successfully joined the server
type playerMessage struct {
	Type     string `json:"type"`
	UserName string `json:"playerName"`
	Colour   string `json:"playerColour"`
	ID       string `json:"gameID"`
}

// return player colour
func (msg playerMessage) PlayerColour() string {
	return msg.Colour
}

// return the type of message
func (msg playerMessage) MessageType() string {
	return msg.Type
}

// encode the message in JSON to send over websocket
func (msg playerMessage) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

// use to update the opponent client with new moves
type moveMessage struct {
	Type string `json:"type"`
	ID   string `json:"gameID"`
	From string `json:"from"`
	To   string `json:"to"`
	Fen  string `json:"fen"`
}

// return message type
func (msg moveMessage) MessageType() string {
	return msg.Type
}

// encode the message in JSON to send over websocket
func (msg moveMessage) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

type startMessage struct {
	Type string `json:"type"`
}

// return message type
func (msg startMessage) MessageType() string {
	return msg.Type
}

// encode the message in JSON to send over websocket
func (msg startMessage) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

// construct a new outgoing websocket message based on the type. Session can be nil
func newOutGoingMessage(msgType string, session *PlayerSession) (Message, error) {

	if !validMessages[msgType] {
		return nil, fmt.Errorf("type not recognised")
	}

	switch msgType {
	case joinConfirmation:
		return playerMessage{
			Type:     joinConfirmation,
			UserName: session.username,
			Colour:   string(session.colour),
			ID:       newGameID(),
		}, nil
	case makeMove:
		return moveMessage{
			Type: makeMove,
			ID:   session.state.GameID,
			From: session.state.From,
			To:   session.state.To,
			Fen:  session.state.Fen,
		}, nil
	case startGame:
		return startMessage{
			Type: startGame,
		}, nil
	default:
		return nil, fmt.Errorf("message request not recognised")
	}

}

// use if incoming ws contains a join message
type joinMsg struct {
	Username string `json:"name"`
	ID       string `json:"uniqueID"`
}
