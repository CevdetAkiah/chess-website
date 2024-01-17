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
    setGameID,
    types, } from '../../context/game/actions';
import { checkGameID, checkSession, getGameOverState } from '../../functions';
import { SiteContext } from '../../context/website/ClientContext';


const serverURL = 'ws://localhost:8080/ws'

const FEN = 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1';
// checkmate for testing game over screen
// const FEN = 'rnb1kbnr/pppp1ppp/8/4p3/5PPq/8/PPPPP2P/RNBQKBNR w KQkq - 1 3';

const Game = ()=> {
    const[fen, setFen] = useState(FEN);
    const {current: chess} = useRef(new Chess(fen));
    const [board,setBoard] = useState(createBoard(fen));
    const { state } = useContext(SiteContext);
    const { possibleMoves, dispatch, opponentName, gameID } = useContext(GameContext);
    const { username, loggedIn } = state; // get username from log in details
    
    const gameIDRef = useRef(gameID);
    const wsRef = useRef(null)

    useEffect(() => {
        checkSession()

        checkGameID()
        .then((gameID) => {
            if (gameID !== null){
                gameIDRef.current = gameID
                dispatch(setGameID(gameID))
            }
        })
    },[dispatch])


    
    //TODO: on initial startup the user joins, gets the "welcome user you are b", then the websocket closes immediately?
    useEffect(() => {
    
        if (!wsRef.current) {
            wsRef.current = new WebSocket(serverURL)
            console.log("NEW WEBSOCKET")
        }
        
        wsRef.current.onopen = (event) =>{
            console.log("connection established: ", event)
            
            const joinName = loggedIn ? username : 'Anonymous';
            
            console.log("gameID: ",gameIDRef.current)
                const apiRequest = {emit: "join", user : {name : joinName}, uniqueID: gameIDRef.current}
                wsRef.current.send(JSON.stringify(apiRequest)) 
            }            
            wsRef.current.onerror = (err) => {
                console.log("Websocket error: ",err)
            }
    
            wsRef.current.onmessage = (event) => { 
                const msgReceived = JSON.parse(event.data)
                const emit = msgReceived.emit
                switch (emit) {
                    case 'message':
                        console.log(msgReceived.message) 
                        break;
                    case 'playerJoined':
                        dispatch(setPlayer(msgReceived.playerName))
                        dispatch(setPlayerColour(msgReceived.playerColour))
                        dispatch(setGameID(msgReceived.gameID))
                        document.cookie = "gameID=" + msgReceived.gameID +"; SameSite=None";
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
    
            wsRef.current.onclose = (event) => {
                console.log("connection closed: ", event)
                wsRef.current.close();
            }

            return () => {
                if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN){
                    wsRef.current.close();
                }
            };

        },[dispatch, chess, username, loggedIn,gameID]);
    



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
        if (opponentName === ''){
            return
        };
        var validMove = possibleMoves.includes(to)
         if (validMove){
            chess.move({ from, to });
            const moveRequest = {emit: "move", gameID: gameID, from: from, to: to }
            
            wsRef.current.send(moveRequest)
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