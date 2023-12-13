import React, { useContext } from 'react';
import { SiteContext } from '../../../context/website/ClientContext';
import axios from 'axios';
import { setLoggedIn } from '../../../context/website/actions';

// logout
const SignOut = () => {
    const { dispatch } = useContext(SiteContext)
    const serverURL = "http://localhost:8080/logout"

    const logOut = () => {
        const config = {
            withCredentials: true,
        }   
        axios.post(serverURL,{},config).then((response) => {
            if (response.status === 204){
                dispatch(setLoggedIn(false))
            }else{
                console.log("error logging out failed")
            }
        })
    }
    return (
        <div><button onClick={logOut}>Sign out</button></div>
    )
}

export default SignOut