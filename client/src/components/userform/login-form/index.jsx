import React, { useState, useContext } from 'react';
import { useForm } from 'react-hook-form';
import axios from 'axios';
import './login-form.css'
import RegisterForm from '../register-form';
import { SiteContext } from '../../../context/website/ClientContext';
import { setClientUsername, setLoggedIn } from '../../../context/website/actions';
import { INCORRECT_PASSWORD, USERNAME_NOT_FOUND } from '../error-types';

// POST user logs in
const LoginForm = () =>{
    
    
    const form = useForm();
    const { register, handleSubmit,setError, formState: { errors }, reset } = form;
    // const { errors, setError } = formState;
    const [pop, setPop] = useState(false)
    const [close, setClose] = useState(false)
    const { dispatch ,state } = useContext(SiteContext)
    const {endpoint, serverport} = state
    
    const serverURL = "http://" + endpoint + serverport + "/session"
        // send user data to the server
        const sendFormData = (data) => {
            const config = {
                headers: { 'Content-Type': 'application/json' },
                withCredentials: true,
            }
            if (data) {
                axios.post(serverURL, JSON.stringify(data), config).then((response) => {
                    if (response.status === 200) {
                        // turn off register form
                        toggleLoginForm()
                        const name = response.data.username;
                        // set site state to logged in
                        dispatch(setClientUsername(name))
                        dispatch(setLoggedIn(true))
                        alert("Loggin successful")
                    }else{
                        // user not found
                        console.log("unexpected server response ", response.status)
                    }
                })
                .catch(function (error) {
                    reset()
                    if (error.response.status === 401){
                        let errorName = "";
                        switch (error.response.data.trim()){
                            case INCORRECT_PASSWORD:
                                errorName = "password"
                                break;
                            case USERNAME_NOT_FOUND:
                                errorName = "email"
                                break;
                            default:
                                console.log("unexpected auth error")
                            }
                        setError(errorName,{
                            type: "server",
                            message: error.response.data
                        });
                    }else{
                        console.log("Authorization error: ", error.response)
                    }
                });
            }
    
                  
        };

    
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
                <p className="error">{errors.email?.message}</p>
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
            <button className="register" onClick={toggleRegisterForm} >REGISTER</button>      
        </form>     
        {pop ? <RegisterForm toggle={toggleRegisterForm}/>: null}
            
    </div>
)
}


export default LoginForm;