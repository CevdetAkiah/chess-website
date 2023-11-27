import React, { useContext } from 'react';
import { SiteContext } from '../../context/website/ClientContext';


// swap components based on site state
const Profile = () => {
    const { state } = useContext(SiteContext)
    const { username } = state;
        
    
    
    return (
        <div>
        <h1 style={{color: 'blue'}}>{username}</h1>
    </div>
)
}

export default Profile