import React, { useContext } from 'react';
import UserForm from '../userform/user-form';
import Profile from '../profile';
import { SiteContext } from '../../context/website/ClientContext';
import Game from '../../pages/Game';


// swap components based on site state
const ComponentSwapper = () => {
    const { state } = useContext(SiteContext)
    const { loggedIn } = state;
    console.log(loggedIn)
    if (loggedIn) {
        return <Profile/>
    } else{
        return <UserForm/>
    }
}


export default ComponentSwapper