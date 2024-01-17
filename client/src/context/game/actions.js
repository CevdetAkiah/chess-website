export const types = {
    SET_POSSIBLE_MOVES: 'SET_POSSIBLE_MOVES',
    CLEAR_POSSIBLE_MOVES: 'CLEAR_POSSIBLE_MOVES',
    SET_TURN: 'SET_TURN',


    SET_PLAYER: 'SET_PLAYER',
    SET_OPPONENT: 'SET_OPPONENT',
    SET_PLAYER_COLOUR: 'SET_PLAYER_COLOUR',
    SET_MESSAGE: 'SET_MESSAGE',
    CLEAR_MESSAGE: 'CLEAR_MESSAGE',
    SET_OPPONENT_MOVES: 'SET_OPPONENT_MOVES',
    CLEAR_OPPONENT_MOVES: 'CLEAR_OPPONENT_MOVES',
    SET_OPPONENT_COLOUR: 'SET_OPPONENT_COLOUR',
    
    SET_WEBSOCKET: 'SET_WEBSOCKET',
    SET_GAMEID: 'SET_GAMEID',
};

export const setPlayer = (name) => ({
    type: types.SET_PLAYER,
    name,
});

export const setOpponent = (opponent) => ({
    type: types.SET_OPPONENT,
    name: opponent?.name,
});

export const setPlayerColour = (colour) => ({
    type: types.SET_PLAYER_COLOUR,
    colour,
});

export const setMessage = (message) => ({
    type: types.SET_MESSAGE,
    message,
});

export const setOpponentMoves = (moves) => ({
    type: types.SET_OPPONENT_MOVES,
    moves,
});

export const clearOpponentMoves = () => ({
    type: types.CLEAR_OPPONENT_MOVES,
});

export const setOpponentColour = (colour) => ({
    type: types.SET_OPPONENT_COLOUR,
    colour,
});

export const setWebsocket = (websocket) => ({
    type: types.SET_WEBSOCKET,
    gameWebSocket: websocket,
});

export const setGameID = (id) => ({
    type: types.SET_GAMEID,
    gameID: id,
});