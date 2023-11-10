import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { DevTool } from '@hookform/devtools';
import axios from 'axios';
import './loginForm.css'
import RegisterForm from '../regiserForm';


const LoginForm = () =>{

    const serverURL = "http://localhost:8080"

    const form = useForm();
    const { register, control, handleSubmit, formState } = form;
    const { errors } = formState;
    const [pop, setPop] = useState(false)
    const [close, setClose] = useState(false)

        // send user date to 
        const sendFormData = (data) => {
            const config = {
                headers: { 'Content-Type': 'multipart/form-data' }
            }
            if (data) {
                axios.post(serverURL + "/authenticate", JSON.stringify(data), config).then((response) => {
                    if (response.status === 200) {
                        // turn off register form
                        toggleLoginForm()

                        console.log(response.statusText)
                    }
                });
            }
                        
        }

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