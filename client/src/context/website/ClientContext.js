import React, { createContext, useReducer } from 'react';
import ClientReducer from './ClientReducer';

const initialState = {
    username: "",
    loggedIn: false,
}

export const SiteContext = createContext();

export const SiteProvider = ({ children }) => {
    const [state, dispatch] = useReducer(ClientReducer, initialState)

    return (
        <SiteContext.Provider value = {{ state, dispatch }}>
            { children }
        </SiteContext.Provider>
    )
}