import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import axios from 'axios';
import './register-form.css'
import { EMAIL_DUPLICATE, USERNAME_DUPLICATE } from '../error-types';


// POST user registration
const RegisterForm = () => {
    const apiURL = process.env.REACT_APP_BACKEND_URL;
    const apiPORT = process.env.REACT_APP_BACKEND_PORT;
    const serverURL = `${apiURL}:${apiPORT}/user`
    const form = useForm();
    const { register, handleSubmit, formState, reset, setError } = form;
    const { errors } = formState;
    const [click, setClick] = useState(false)

    // control register form popup
    const togglePop = () => {
        setClick(!click)
    }

    // send user date to 
    const sendFormData = (data) => {
        console.log(data)
        const config = {
            headers: { 'Content-Type': 'multipart/form-data' }
        }
        if (data) {
            axios.post(serverURL, JSON.stringify(data), config).then((response) => {
                if (response.status === 201) {
                    // TODO: register confirmation
                    // turn off register form
                    togglePop()
                }
            })
            .catch(function (error) {
                let errorName ="";
                reset()
                if (error.response.status === 409){
                    switch (error.response.data.trim()){
                        case EMAIL_DUPLICATE:
                            errorName = "email"
                            break;
                        case USERNAME_DUPLICATE:
                            errorName = "username"
                            break;
                        default:
                            console.log("unexpected signup error")
                    }
                    setError(errorName,{
                        type: "server",
                        message: error.response.data
                    });
                } else{
                    setError(errorName,{
                        type: "server",
                        message: error.response.data
                    });
                }
            })
        }
    }


 
    return (
               
            <div>
                <form style = {click ? {display: 'none'}: {}}autoComplete="off" className="registerForm" onSubmit={handleSubmit(sendFormData)} noValidate> 
                <header className="header">REGISTER</header>
                    <div className="form-control">
                    <label>User name</label>
                        <input
                            type='text'
                            id = 'username'
                            placeholder='Enter username'
                            {...register("username", { 
                                required: 'Username is required', // validation
                            })}
                        />
                        <p className="error">{errors.username?.message}</p>
                    </div>

                    <div  className="form-control">
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
                    <button type="submit">Submit</button>

                </form>
            </div>
   

    )
}

export default RegisterForm;