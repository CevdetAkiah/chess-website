
import React  from 'react';
import './App.css';
import { SiteProvider } from './context/website/ClientContext';
import NavBar from './components/Navbar';
import router from './components/router';




function App() {
    return (
        <SiteProvider>  
                 <NavBar />
                {router()}
        </SiteProvider>
    );
}

export default App;