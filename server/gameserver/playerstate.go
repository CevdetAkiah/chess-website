package gameserver

import "github.com/anthdm/hollywood/actor"

// hold chess game data
type PlayerState struct {
	sessionID int
	GameID    string `json:"gameID"`
	From      string `json:"from"`
	To        string `json:"to"`
	Fen       string `json:"fen"`
}

func newPlayerState(incomingMove Message) *PlayerState {
	return &PlayerState{
		GameID: incomingMove.MessageType(),
		From:   incomingMove.(moveMessage).From,
		To:     incomingMove.(moveMessage).To,
		Fen:    incomingMove.(moveMessage).Fen,
	}
}

func (state *PlayerState) Receive(c *actor.Context) {}
