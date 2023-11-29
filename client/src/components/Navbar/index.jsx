import React, { useContext } from 'react';
import UserForm from '../userform/user-form';
import Profile from '../profile';
import { SiteContext } from '../../context/website/ClientContext';


// swap components based on site state
const Navbar = () => {
    const { state } = useContext(SiteContext)
    const { loggedIn } = state;
    if (loggedIn) {
        return <Profile/>
    } else{
        return <UserForm/>
    }
}


export default Navbar;