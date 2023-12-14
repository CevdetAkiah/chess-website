import React, { useContext } from 'react';
import { SiteContext } from '../../../context/website/ClientContext'
import './profile.css'


// swap components based on site state
const Profile = () => {
    const { state } = useContext(SiteContext)
    const { username } = state;
        
    
    
    return (
        <div className='profile'>
             <button>{username}</button>
       </div>
)
}

export default Profile