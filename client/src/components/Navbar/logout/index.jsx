import React, { useContext } from 'react';
import { SiteContext } from '../../../context/website/ClientContext';
import axios from 'axios';
import { setLoggedIn } from '../../../context/website/actions';
import './logout.css'

// DELETE a session to log a user out
const SignOut = () => {
    const { dispatch } = useContext(SiteContext)
    const serverURL = "http://localhost:8080/session" 

    const logOut = () => {
        const config = {
            withCredentials: true,
        }   
        axios.delete(serverURL,config).then((response) => {
            if (response.status === 204){
                dispatch(setLoggedIn(false))
            }else{
                console.log("error logging out failed")
            }
        })
    }
    return (
        <div className='logout'><button onClick={logOut}>SIGN OUT</button></div>
    )
}

export default SignOut