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
    const [click, setClick] = useState(false)

        // send user date to 
        const sendFormData = (data) => {
            console.log(data)
            const config = {
                headers: { 'Content-Type': 'multipart/form-data' }
            }
            if (data) {
                axios.post(serverURL + "/signupAccount", JSON.stringify(data), config).then((response) => {
                    console.log("Form response: ", response.status)
                });
            }
            
        }

const togglePop = () =>{
    setClick(!click);
};
        
    
return (
    <div >
        <form className="loginForm" style={click ? { display: 'none'}: {}}  onSubmit={handleSubmit(sendFormData)} noValidate> 
        <header className="header">LOG IN</header>
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
            <button className="register" onClick={togglePop} onSubmit="">REGISTER</button>      
        </form>     
        {click ? <RegisterForm toggle={togglePop}/>: null}
    </div>
)
}


export default LoginForm;