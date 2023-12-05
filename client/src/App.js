
import React, { Component, useContext, useState, useEffect} from 'react';
import './App.css';
import { BrowserRouter, Routes, Route, NavLink } from 'react-router-dom'
import { SiteContext, SiteProvider } from './context/website/ClientContext';
import { GameProvider } from './context/game/GameContext';
import  IndexPage  from './pages/Index';
import NavBar from './components/Navbar';
import Game from './pages/Game';




// class ErrorBoundary extends Component {
//         constructor(props) {
//                 super(props);
//                 this.state= { hasError: false };
//         }

//         componentDidCatch(error, errorInfo) {
//                 // Log error or perform error handling
//                 console.error(error);
//                 this.setState({ hasError: true });
//         }

//         render(){
//                 if (this.state.hasError) {
//                         return(<div>
//                                <h2> Something went wrong.</h2>
//                                <p>error</p>
//                                 </div>);
//                 }

//                 return this.props.children;
//         }
// }


const router = () => (
        <BrowserRouter>

        <nav>
                <NavLink to="game">Play</NavLink>
        </nav>        

        <main>
                <Routes>
                         <Route path="game/*" element={<GameProvider><Game /></GameProvider>}/>
                 </Routes>
         </main>
        </BrowserRouter> 
);


function App() {
    return (
        <SiteProvider>  
                 <NavBar />
                {router()}
          </SiteProvider>
    );
}

export default App;