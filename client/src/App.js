import React, { Component } from 'react';
import Game from './pages/Game';
import { GameProvider } from './context/game/GameContext';
import './App.css';
import UserForm from './components/userform/user-form';
import ComponentSwapper from './components/component-swapper';
import { SiteProvider } from './context/website/ClientContext';



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
        <ErrorBoundary>
                <SiteProvider>
                                 {/* <UserForm/> */}
                                 <ComponentSwapper/>
                </SiteProvider>
                <GameProvider>
                         <div className="game">
                                <Game/>
                        </div>
                </GameProvider>    
        </ErrorBoundary>

    );
}

export default App;