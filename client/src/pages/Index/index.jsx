import React, { useContext } from 'react';
import Navbar from '../../components/Navbar';
import { SiteContext } from '../../context/website/ClientContext';
import  Game  from '../Game';
import { GameProvider } from '../../context/game/GameContext';


const IndexPage = () => {
    const { state } = useContext(SiteContext)
    const { loggedIn } = state;

    return (
        <div>
            <div><Navbar/></div>
            <GameProvider>
            <div>{loggedIn && <Game/>}</div>
            </GameProvider>
        </div>
    ) 
};

export default IndexPage;