import { types } from './actions';

// returns an array of possible moves 
const getPositions = (moves) => {
    return moves.map((move) => {
    const n = move.length;
       const [position, castling] = checkCastling(move,moves)
        if (castling){
            return position
        };
            if (move.substring(n-1) === '+'){ // if a move results in check
                return move.substring(n-3, n-1)
            }; 
            return move.substring(n-2)
        });
};

const checkCastling = (move, moves) => {
    switch (move) {
        case "O-O": // kingside castle
            if (moves.includes("Kf1")){
                return ["g1", true]
            };
            if (moves.includes("Kf8")){
                return ["g8",true]
            };
            break;
        case "O-O-O": // queenside castle
            if (moves.includes("Kd1")){
                return ["c1",true]
            };
            if (moves.includes("Kd8")){
                return ["c8",true]
            };
            break;
        default:
            return [null,false]
    };

};

// takes in the previous game state and action to apply to the state, then returns a new updated state based on the action
const GameReducer = (state, action) =>{
    switch (action.type) {
        case types.SET_POSSIBLE_MOVES: // highlight possible cells to move to
            return {
                ...state,
                possibleMoves: getPositions(action.moves),
            };
            case types.CLEAR_POSSIBLE_MOVES: // unhighlight cells
                return {
                    ...state,
                    possibleMoves: [],
                };
            case types.SET_TURN:
                return {
                    ...state,
                    turn: action.player,
                    check: action.check,
                };
            case types.GAME_OVER:
                return {
                    ...state,
                    gameOver: true,
                    status: action.status,
                    turn: action.player,
                }

            case types.SET_PLAYER:
                return { ...state, playerName: action.name };
            case types.SET_PLAYER_COLOUR:
                return { ...state, playerColour: action.colour };
            case types.SET_OPPONENT:
                return { ...state, opponentName: action.name };
            case types.SET_MESSAGE:
                return { ...state, message: action.message };
            case types.CLEAR_MESSAGE:
                return { ...state, message: '' };
            case types.SET_OPPONENT_MOVES:
                return { ...state, opponentMoves: action.moves };
            case types.CLEAR_OPPONENT_MOVES:
                 return { ...state, opponentMoves: [] };
            case types.SET_OPPONENT_COLOUR:
                return { ...state, opponentColour: action.colour };
            default:
                return state;
    }
};

export default GameReducer;