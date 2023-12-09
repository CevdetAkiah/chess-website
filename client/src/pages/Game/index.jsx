import React, { useState, useRef, useEffect, useContext} from 'react';
import './game.css';
import {Chess} from 'chess.js';
import { createBoard } from '../../functions/create-board';
import Board from  '../../components/board';
import { GameContext } from '../../context/game/GameContext';
import {     
    setOpponent,
    setOpponentColour,
    setOpponentMoves,
    setPlayer,
    setPlayerColour,
    types, } from '../../context/game/actions';
import { getGameOverState } from '../../functions';
import { SiteContext } from '../../context/website/ClientContext';
import  { randGameID }  from '../../functions/game-ID';

const serverURL = 'ws://localhost:8080/ws'

const FEN = 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1';
// checkmate for testing game over screen
// const FEN = 'rnb1kbnr/pppp1ppp/8/4p3/5PPq/8/PPPPP2P/RNBQKBNR w KQkq - 1 3';

// const ws = new WebSocket(serverURL)
const Game = ()=> {
    const[fen, setFen] = useState(FEN);
    const {current: chess} = useRef(new Chess(fen));
    const [board,setBoard] = useState(createBoard(fen));
    const { possibleMoves, dispatch, opponentName } = useContext(GameContext);
    const { state } = useContext(SiteContext);
    const { username, loggedIn } = state; // get username from log in details
    
    const ws = useRef(null); // useRef allows a persistent wesbsocket across re renders; ensuring the connection is only created once.

    useEffect(() => {
        ws.current = new WebSocket(serverURL)
            ws.current.onopen = (event) =>{
                console.log("connection established: ", event)
                const gameID  = randGameID
                
                const joinName = loggedIn ? username : 'Anonymous';
                const apiRequest = {emit: "join", user : {name : joinName, uniqueID: gameID}}
                ws.current.send(JSON.stringify(apiRequest)) 
            }            
            ws.current.onerror = (err) => {
                console.log("Websocket error: ",err)
            }
    
            ws.current.onmessage = (event) => { 
                const msgReceived = JSON.parse(event.data)
                const emit = msgReceived.emit
                switch (emit) {
                    case 'message':
                        console.log(msgReceived.message) 
                        break;
                    case 'playerJoined':
                        console.log("player: ", msgReceived.playerName)
                        dispatch(setPlayer(msgReceived.playerName))
                        dispatch(setPlayerColour(msgReceived.playerColour))
                        break;
                    case 'opponentJoined':
                        console.log("opponent: ", msgReceived.opponentName)  
                        dispatch(setOpponent(msgReceived.opponentName))
                        dispatch(setOpponentColour(msgReceived.opponentColour))
                        break;
                    case 'opponentMove':
                        const from = msgReceived.from
                        const to = msgReceived.to
                        chess.move({ from, to })
                        setFen(chess.fen()); // update the fen with the new move/piece positions
                        dispatch(setOpponentMoves([from,to]))
                        break;
                    default:
                    };
                }
    
            ws.current.onclose = (event) => {
                console.log("connection closed: ", event)
                ws.current.close();
            }

            return () => {
                if (ws.current){
                    ws.current.close();
                }
            };

        },[dispatch, chess, username, loggedIn]);
    



    // every time a change is made to the state of the game, the board is updated with the new fen
    useEffect(() =>{
        setBoard(createBoard(fen));
    }, [fen]);



    // will detect if a player is in check
    useEffect(() =>{
        const [gameOver, status] = getGameOverState(chess);
        if (gameOver) {
            dispatch({ type: types.GAME_OVER, status, player: chess.turn() });
            return
        }
        dispatch({
            type: types.SET_TURN,
            player: chess.turn(),
            check: chess.isCheck(),
        });
    }, [fen, dispatch, chess]);

    /** move handling */ 
    const fromPos = useRef(); /** follow setFromPos to the piece component */

    // share move set with components

    const makeMove = (pos) =>{
        const from = fromPos.current;
        const to = pos;
        if (opponentName == ''){
            return
        };
        var validMove = possibleMoves.includes(to)
         if (validMove){
            chess.move({ from, to });
            const apiRequest = {emit: "move", gameID: 1, from: from, to: to }
            ws.current.send(JSON.stringify(apiRequest))
            dispatch({ type: types.CLEAR_POSSIBLE_MOVES}) // unhighlight possible moves
            setFen(chess.fen()); // update the fen with the new move/piece positions
        }
    };

    // this is called once a piece is dragged
    const setFromPos = (pos) => {
        fromPos.current = pos
        dispatch({
            type:types.SET_POSSIBLE_MOVES,
            moves: chess.moves({ square: pos }) // send the possible moves from the currently selected position to highlight
        })
    };

    // if (gameOver) {
    //     return <GameOver />
    // };

    return (
            <div className="game">
            <Board cells = {board} makeMove={makeMove} setFromPos={setFromPos}/>
            </div>
    );

};

export default Game;