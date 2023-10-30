import React, { useContext } from 'react';
import { isLightSquare, Cell as BoardCell } from '../../functions/';
import PropTypes from 'prop-types';
import Piece from'../piece';
import { GameContext } from '../../context/GameContext';
import './cell-styles.css';

const Cell = ( {cell, index, makeMove, setFromPos} ) => {
    const light = isLightSquare(cell.pos, index); /**returns true if cell should be light */
    const { possibleMoves, turn, check, playerColour  } = useContext(GameContext); 
    const isPossibleMove = possibleMoves.includes(cell.pos) && turn === playerColour; // check if this cell's position is a possible move
    const colour = cell.piece.toUpperCase() === cell.piece ? 'w' : 'b';
        // TODO: inCheck is returning false when it shouldn't
    const inCheck = () => {
        const king = cell.piece.toUpperCase === 'K';
        return turn === colour && king && check; // return true if the turn is the colour of current player and the piece is a king and in check (according to Chess.js)
    }

    const handleDrop = () => {
        if (turn === playerColour){
            makeMove(cell.pos);
        }
    };
        
    return ( 
            <div
                className={`cell ${light ? 'light' : 'dark'}`}
                onDrop={handleDrop}
                onDragOver={(e) => e.preventDefault()}
                >
                <div className={`overlay ${isPossibleMove && 'possible-move'} ${ inCheck() && 'check' }`}>
                        <Piece name={cell.piece} pos={cell.pos}  setFromPos={setFromPos} />
                </div>
            </div>
        );


        
};

Cell.prototype = {
    cell: PropTypes.instanceOf(BoardCell).isRequired,
    index: PropTypes.number.isRequired,
    makeMove: PropTypes.func,
    setFromPos: PropTypes.func,
}

export default Cell;