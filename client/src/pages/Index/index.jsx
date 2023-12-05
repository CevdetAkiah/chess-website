import React, { useContext } from 'react';

import NavBar from '../../components/Navbar';
import { SiteContext } from '../../context/website/ClientContext';

import { SiteProvider } from '../../context/website/ClientContext';
import  Game  from '../Game';
import { GameProvider } from '../../context/game/GameContext';
import { SplitScreen } from '../../components/layout/SplitScreen';
// import { BrowserRouter, Router, Route} from 'react-router-dom';

// const router = createBrowserRouter([

// ])

const LeftComponent =() =>{
    return(
            <SiteProvider><NavBar /></SiteProvider>
    )
}

const RightComponent =() =>{
    
    return(
        <h1>Hi</ h1>
    )
}


const IndexPage = () => {
    return (
        <SiteProvider>
                <SplitScreen leftWeight={1} rightWeight={3}>      
                        <LeftComponent/>                
                        <RightComponent/>
                </SplitScreen>
        </SiteProvider>
    );
};

export default IndexPage;