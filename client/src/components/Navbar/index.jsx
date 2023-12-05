import React, { useContext } from 'react';
import UserForm from '../userform/user-form';
import Profile from '../profile';
import { SiteContext } from '../../context/website/ClientContext';


// swap components based on site state

const NavBar = () => {
    const { state } = useContext(SiteContext)
    const { loggedIn } = state;
 
    return (
        <nav>
            {loggedIn ? <Profile />  : <UserForm />}
        </nav>
    )
}


export default NavBar;
