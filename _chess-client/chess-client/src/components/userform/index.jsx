import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { DevTool } from '@hookform/devtools';
import axios from "axios";
import './userForm.css'



const UserForm = () => {
    const serverURL = "http://localhost:8080"

    const form = useForm();
    const { register, control, handleSubmit, formState } = form;
    const { errors } = formState;


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


    return (

        <div>
            <form className="userForm" onSubmit={handleSubmit(sendFormData)} noValidate> 
                <label>User name</label>
                <div className="form-control">
                    <input
                        type='text'
                        id = 'username'
                        placeholder='Enter username'
                        {...register("username", { 
                            required: 'Username is required',
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
                            required: 'Email is required',
                            pattern: {
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
            <DevTool control={control}/>
        </div>
    )

}

export default UserForm;