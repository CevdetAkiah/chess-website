package chesswebsocket

import (
	"fmt"
	"go-projects/chess/service"
	"io"

	"golang.org/x/net/websocket"
)

func NewChessWebsocket(DBA service.DbService) *WsGame {
	return &WsGame{
		conns:    make(map[*websocket.Conn]bool),
		DBAccess: DBA,
	}
}

func (wsg *WsGame) HandleWS(wsc *websocket.Conn) {
	wsg.lock.Lock()
	wsg.conns[wsc] = true
	wsg.lock.Unlock()

	wsg.readConn(wsc)

}

func (wsg *WsGame) readConn(wsc *websocket.Conn) {
	message := &receiveMessage{}
	for {
		err := websocket.JSON.Receive(wsc, &message)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("msg read error: ", err)
		}

		if message != nil {
			switch message.Emit {
			case "join":
				wsg.handleJoin(message, wsc)
			case "message":
				wsg.handleMessage(message)
			case "move":
				wsg.handleMove(message, wsc)
			}
		}
	}
}

// send move to the player that is not the current websocket (THE OPPONENT RELATIVE TO THE WEBSOCKET)
func (wsg *WsGame) handleMove(msg *receiveMessage, wsc *websocket.Conn) {
	message := &sendMove{emitMessage: emitOpponentMove, FromMV: msg.FromMV, ToMV: msg.ToMV}
	player := wsc

	if wsg.gamesInPlay[msg.GameID].playerOne.PlayerID == player {
		opponent := wsg.gamesInPlay[msg.GameID].playerTwo.PlayerID
		opponent.Write(encodeMessage(message))
	} else {
		opponent := wsg.gamesInPlay[msg.GameID].playerOne.PlayerID
		opponent.Write(encodeMessage(message))
	}
	// wsg.Broadcast(msg.GameID, message)
}

func (wsg *WsGame) handleMessage(msg *receiveMessage) {
	fmt.Println(msg.Message)
}

// TODO: the logic for writing out to players is clumsy and needs re organising

// handle the join event, join a game
func (wsg *WsGame) handleJoin(msg *receiveMessage, wsc *websocket.Conn) {
	player := newPlayer(&msg.User, wsc)
	// join game or create new one
	if len(wsg.gameSearch) == 0 { //if no available games, create a new one
		game := &Game{
			gameID: 1,
		}
		player.Colour = randColour()
		game.playerOne = player
		wsg.gameSearch = append(wsg.gameSearch, game) // add game to the search list

		playerInfo := &sendPlayerInfo{emitMessage: emitPlayerJoined, PlayerName: player.Name, PlayerColour: player.Colour}
		player.PlayerID.Write(encodeMessage(playerInfo))

		playerMessage := &sendMessage{emitMessage: emitMsg, Message: "welcome " + player.Name + " you are playing as " + player.Colour}
		player.PlayerID.Write(encodeMessage(playerMessage))

	} else { // join someone's game
		game := wsg.gameSearch[0]
		opponent := game.playerOne
		if opponent.Colour == white {
			player.Colour = black
		} else {
			player.Colour = white
		}
		game.playerTwo = opponent // add player to game
		game.playerOne = player
		if len(wsg.gamesInPlay) == 0 {
			wsg.gamesInPlay = make(map[int]*Game)
		}

		wsg.gamesInPlay[game.gameID] = game // add game to started games map
		wsg.gameSearch = nil                // game is now in play

		// let the player know they have joined a game and their colour
		message := &sendMessage{emitMessage: emitMsg, Message: "welcome " + player.Name + " you are playing as " + player.Colour}
		player.PlayerID.Write(encodeMessage(message))

		playerInfo := &sendPlayerInfo{emitMessage: emitPlayerJoined, PlayerName: player.Name, PlayerColour: player.Colour}
		player.PlayerID.Write(encodeMessage(playerInfo))

		// let the player know their opponent info and set the opponent in their client
		opponentInfo := &sendOpponentInfo{emitMessage: emitOpponentJoined, OpponentName: opponent.Name, OpponentColour: opponent.Colour}
		player.PlayerID.Write(encodeMessage(opponentInfo))
		message = &sendMessage{emitMessage: emitMsg, Message: "you are playing against " + opponent.Name + " who is playing as " + opponent.Colour + "\nStart the game"}
		player.PlayerID.Write(encodeMessage(message))

		// let the opponent know a player has joined and set the opponent's opponent as the player
		message = &sendMessage{emitMessage: emitMsg, Message: opponent.Name + " has joined and is playing " + opponent.Colour + "\nStart the game"}
		opponent.PlayerID.Write(encodeMessage(message))
		opponentInfo = &sendOpponentInfo{emitMessage: emitOpponentJoined, OpponentName: opponent.Name, OpponentColour: opponent.Colour}
		opponent.PlayerID.Write(encodeMessage(opponentInfo))
	}

}
