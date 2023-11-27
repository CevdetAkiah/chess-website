export const types = {
    SET_CLIENT_USERNAME: 'SET_CLIENT_USERNAME',
    SET_LOGGED_IN: 'SET_LOGGED_IN',
};

export const setClientUsername = (username) => ({
    type: types.SET_CLIENT_USERNAME,
    username,
});

export const setLoggedIn = (value) => ({
    type: types.SET_LOGGED_IN,
    value
});