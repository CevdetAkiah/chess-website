import React, { Component, useContext, useState, useEffect } from 'react';
import Game from './pages/Game';
import { GameProvider } from './context/game/GameContext';
import './App.css';
import { SiteContext, SiteProvider } from './context/website/ClientContext';
import  IndexPage  from './pages/Index';
import { SplitScreen } from './components/layout/SplitScreen';
import Navbar from './components/Navbar';
import { SiteProvider } from './context/website/ClientContext';
import  IndexPage  from './pages/Index';



class ErrorBoundary extends Component {
        constructor(props) {
                super(props);
                this.state= { hasError: false };
        }

        componentDidCatch(error, errorInfo) {
                // Log error or perform error handling
                console.error(error);
                this.setState({ hasError: true });
        }

        render(){
                if (this.state.hasError) {
                        return(<div>
                               <h2> Something went wrong.</h2>
                               <p>error</p>
                                </div>);
                }

                return this.props.children;
        }
}



function App() {
    return (
        <SiteProvider>
                <IndexPage />
        </SiteProvider>
    );
}

export default App;