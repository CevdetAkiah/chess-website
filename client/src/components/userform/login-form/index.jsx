import React, { useState, useContext, useEffect } from 'react';
import { useForm } from 'react-hook-form';
import axios from 'axios';
import './login-form.css'
import RegisterForm from '../register-form';
import { SiteContext } from '../../../context/website/ClientContext';
import { setClientUsername, setLoggedIn } from '../../../context/website/actions';


const LoginForm = () =>{
    
    const serverURL = "http://localhost:8080"

    const form = useForm();
    const { register, handleSubmit, formState } = form;
    const { errors } = formState;
    const [pop, setPop] = useState(false)
    const [close, setClose] = useState(false)
    const { state, dispatch } = useContext(SiteContext)
    const { username, loggedIn } = state;

        // send user data to the server
        const sendFormData = (data) => {
            const config = {
                headers: { 'Content-Type': 'application/json' },
                withCredentials: true,
            }
            if (data) {
                axios.post(serverURL + "/authenticate", JSON.stringify(data), config).then((response) => {
                    if (response.status === 200) {
                        // turn off register form
                        toggleLoginForm()
                        const name = response.data.username;
                        // set site state to logged in
                        console.log(name+ " you are logged in")
                        dispatch(setClientUsername(name))
                        dispatch(setLoggedIn(true))
                    }
                });     
            }
                        
        };

        useEffect(() =>{
            console.log("Username: ", username)
            console.log("Logged in: ", loggedIn)
        }, [username,loggedIn]);
    
        // controls register form popup
const toggleRegisterForm = () =>{
    setPop(!pop);
    setClose(!close)
};

// closes the login form
const toggleLoginForm = () => {
    setClose(!close);
};


        
    
return (
    <div >
        <form className="loginForm" style={ close ? { display: 'none'}: {}}  onSubmit={handleSubmit(sendFormData)} noValidate> 
        <header className="header">LOG IN</header>
            <div className="form-control">
            <label>Email</label>
            <input
                type='email'
                id = 'email'
                placeholder='Enter email'
                {...register("email", {
                    required: 'Email is required', // validation
                    pattern: { // validation
                        value: /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/,
                        message: 'Invalid email format',
                        validate: {
                            notAdmin: (fieldValue) => {
                                return (
                                    fieldValue !== "admin@chesswebsite.com" || "some other email address" // change depending on admin email
                                );
                            },
                            notBlackListed: (fieldValue) => {
                                return !fieldValue.endsWith("baddomain.com") || "This domain is not supported"
                            }
                        }
                    }
                })}
            />
                <p className="error">{errors.username?.message}</p>
            </div>

            <div  className="form-control">
                <label>Password</label>
                <input
                    type='password'
                    id = 'password'
                    placeholder='Enter password'
                    {...register("password", {
                        required: 'Password is required',
                    })}
                />
                <p className="error">{errors.password?.message}</p>
            </div>
            <button type="submit" className="submit">Submit</button>
            <button className="register" onClick={toggleRegisterForm} onSubmit="">REGISTER</button>      
        </form>     
        {pop ? <RegisterForm toggle={toggleRegisterForm}/>: null}
            
    </div>
)
}


export default LoginForm;