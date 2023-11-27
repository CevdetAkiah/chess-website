import React, {useState, useEffect} from 'react';
import LoginForm from '../login-form'
import './userForm.css'

const UserForm = () => {
    const [click, setClick] = useState(false)

    // toggle the login form
    const togglePop = () =>{
        setClick(!click);
    };

    // collapse the forms
    useEffect(() => { 
        const handleEscape =(event) => {
            if (event.key === 'Escape') {
                setClick(false);
            }
        };
        
        document.addEventListener('keydown', handleEscape);
        return () => {
            document.removeEventListener('keydown', handleEscape)
        }
    }, []); 

    // button used to open the login form
    return (
        <div>
            <div className="userForm">
                <button onClick={togglePop}>SIGN IN</button>
            </div>
            {click ? <LoginForm toggle={togglePop}/> : null}
        </div>
    )
}

export default UserForm