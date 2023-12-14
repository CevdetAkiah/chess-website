import React, { useContext, useEffect } from 'react';
import UserForm from '../userform/user-form';
import Profile from './profile';
import { SiteContext } from '../../context/website/ClientContext';
import { setClientUsername, setLoggedIn } from '../../context/website/actions';
import { checkSession } from '../../functions';
import  SignOut  from './logout'
import './navbar.css'


// swap components based on site state

const NavBar = () => {
    const { state, dispatch } = useContext(SiteContext);
    const { loggedIn } = state;

    useEffect(() => {
        checkSession()
            .then((userName) => {
                if (userName !== null) {
                    dispatch(setLoggedIn(true));
                    dispatch(setClientUsername(userName));
                }
            })
    }, [dispatch]);

    return (
        <nav className='nav'>
            {loggedIn ? <Profile /> : <UserForm />} {loggedIn && <SignOut />}
        </nav>
    );
};

export default NavBar;

