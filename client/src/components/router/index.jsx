import { BrowserRouter, Routes, Route, NavLink } from 'react-router-dom';
import Game from '../../pages/Game';
import { GameProvider } from '../../context/game/GameContext';
import React, {  useState, useEffect } from 'react';
import './router.css'
import Healthz from '../healthcheck';





const Router = () => {
    const [display, toggleDisplay] = useState(true)
    const [gameActive, toggleGameActive] = useState(false)

    useEffect(() => { // don't show play button if we're on the game screen
        toggleGameActive(window.location.pathname.startsWith('/game'))
        if (gameActive){
            toggleDisplay(!display)
        }
    }, [gameActive]);
    const handleSubmit = () => {
        toggleGameActive(true)
    };

      return (<BrowserRouter>
                <nav className="nav-game">
                        <NavLink to="game" onClick={handleSubmit} style={{display: display ? 'block': 'none'}}>Play</NavLink>
                </nav>        

                <main>
                        <Routes>
                                <Route path="healthz" element={<Healthz />}/>
                                <Route path="game/*" element={<GameProvider><Game /></GameProvider>}/>
                        </Routes>
                </main>
        </BrowserRouter> 
        )
};


export default Router;