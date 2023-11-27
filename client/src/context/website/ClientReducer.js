import { types } from './actions'

const ClientReducer = (state, action) => {
    switch (action.type) {
        case types.SET_CLIENT_USERNAME:
            return {
                ...state,
                username: action.username,
            };
        case types.SET_LOGGED_IN:
            return {
                ...state,
                loggedIn: action.value
            }
        default:
            return state;
        }      
}

export default ClientReducer;