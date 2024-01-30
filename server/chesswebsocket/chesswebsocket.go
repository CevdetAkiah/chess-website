package chesswebsocket

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

func NewWebsocket() *WsGame {
	return &WsGame{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (wsg *WsGame) HandleWS(wsc *websocket.Conn) {
	wsg.lock.Lock()

	// close the connection once readConn returns
	defer func() {
		wsg.lock.Lock()
		delete(wsg.conns, wsc)
		wsg.lock.Unlock()
		wsc.Close()
	}()

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
				err := wsg.handleJoin(message, wsc)
				fmt.Println(err)
			case "message":
				wsg.handleMessage(message)
			case "move":
				wsg.handleMove(message, wsc)
			case "reconnect":
				fmt.Println("emit: ", message.Emit)
				err := wsg.handleReconnect(message, wsc)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}
}

func (wsg *WsGame) handleReconnect(msg *receiveMessage, wsc *websocket.Conn) error {
	gameID := msg.GameID[:len(msg.GameID)-1] // remove the player colour
	colour := msg.GameID[len(msg.GameID)-1:] // retrieve the player colour

	if game, ok := wsg.gamesInPlay[gameID]; ok { // find game
		// update the player websocket and send game info back to player (colour, fen string)
		if game.playerOne.Colour == colour {
			game.playerOne.PlayerID.Close()
			game.playerOne.PlayerID = wsc

		} else {
			game.playerTwo.PlayerID.Close()
			game.playerTwo.PlayerID = wsc

		}

	} else {
		return fmt.Errorf("no game in play")
	}
	return nil
}

// send move to the player that is not the current websocket (THE OPPONENT RELATIVE TO THE WEBSOCKET)
func (wsg *WsGame) handleMove(msg *receiveMessage, wsc *websocket.Conn) error {
	message := &sendMove{emitMessage: emitOpponentMove, FromMV: msg.FromMV, ToMV: msg.ToMV}
	player := wsc

	if wsg.gamesInPlay[msg.GameID].playerOne.PlayerID == player {
		opponent := wsg.gamesInPlay[msg.GameID].playerTwo.PlayerID
		opponent.Write(encodeMessage(message))
	} else {
		opponent := wsg.gamesInPlay[msg.GameID].playerOne.PlayerID
		opponent.Write(encodeMessage(message))
	}
	return nil
}

// print message from client to console
func (wsg *WsGame) handleMessage(msg *receiveMessage) {
	fmt.Println(msg.Message)
}

// TODO: need concurrent safe functions to access shared data stores.
// handle the join event, join a game
func (wsg *WsGame) handleJoin(msg *receiveMessage, wsc *websocket.Conn) error {
	player := newPlayer(&msg.User, wsc)

	// join game or create new one

	if len(wsg.gameSearch) == 0 { //if no available games, create a new one
		// if game client is refreshed client has a gameID
		gameID := msg.GameID
		if gameID == "new-game" {
			gameID = newGameID()
		}
		game := &Game{
			ID: gameID,
		}
		player.Colour = randColour()
		player.GameID = gameID + player.Colour
		game.playerOne = player
		wsg.gameSearch = append(wsg.gameSearch, game) // add game to the search list

		playerInfo := &sendPlayerInfo{emitMessage: emitPlayerJoined, PlayerName: player.Name, PlayerColour: player.Colour, GameID: player.GameID}
		_, err := player.PlayerID.Write(encodeMessage(playerInfo))
		if err != nil {
			return fmt.Errorf("websocket emit error: %b", err)
		}

		playerMessage := &sendMessage{emitMessage: emitMsg, Message: "welcome " + player.Name + " you are playing as " + player.Colour}
		_, err = player.PlayerID.Write(encodeMessage(playerMessage))
		if err != nil {
			return fmt.Errorf("websocket emit error: %b", err)
		}

		return nil

	} else { // join someone's game if there is a game in the gameSearch list
		game := wsg.gameSearch[0]
		opponent := game.playerOne
		if opponent.Colour == white {
			player.Colour = black
			player.GameID = game.ID + black
		} else {
			player.Colour = white
			player.GameID = game.ID + white
		}
		game.playerTwo = opponent // add player to game
		game.playerOne = player
		if len(wsg.gamesInPlay) == 0 {
			wsg.gamesInPlay = make(map[string]*Game)
		}

		wsg.gamesInPlay[game.ID] = game // add game to started games map
		wsg.gameSearch = nil            // game is now in play

		// let the player know they have joined a game and their colour
		message := &sendMessage{emitMessage: emitMsg, Message: "welcome " + player.Name + " you are playing as " + player.Colour}
		player.PlayerID.Write(encodeMessage(message))

		playerInfo := &sendPlayerInfo{emitMessage: emitPlayerJoined, PlayerName: player.Name, PlayerColour: player.Colour, GameID: player.GameID}
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

	return nil

}
