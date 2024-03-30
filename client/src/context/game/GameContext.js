import React, { createContext, useReducer } from 'react';
import GameReducer from './GameReducer';

const initialState = {
    possibleMoves: [],
    turn: '',
    check: false, // true if the side to move (current turn) is in check.
    gameOver: false,
    status: '', // game over status eg checkmate or stalemate
    playerName: '',
    playerColour: '',
    opponentName: '',
    opponentColour: '',
    gameStart: false,
    message: '',
    opponentMoves: [],
    webSocket: null,
    gameID: "new-game",
};

export const GameContext = createContext(initialState);

// GameProvider wraps the Game in App.js which will expose the game reducer functions and state to it's children components
export const GameProvider = ({ children }) => {
    const [state, dispatch] = useReducer(GameReducer, initialState);

    if (!GameContext){
        throw new Error("GameContext must be used within a GameProvider")
    }

    return (
        <GameContext.Provider value ={{ ...state, dispatch }}>
            {children}
        </GameContext.Provider>
    )
}