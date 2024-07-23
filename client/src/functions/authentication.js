import axios from 'axios';

// TODO: incorporate these functions into the structure of the app so context can be used

// GET checkSession checks if a session is active
export const checkSession = () => {
    const apiURL = process.env.REACT_APP_BACKEND_URL;
    const serverURL = `${apiURL}/session`; 
    const config = { withCredentials: true };

    return new Promise((resolve) => {
        axios.get(serverURL, config)
            .then((response) => {
                if (response.status === 202 || response.status === 200) {
                    const userName = response.data.username;
                    resolve(userName);
                }else{
                    resolve(null);
                }
            })
            .catch((error) => {
                if (error.status === 408){
                    console.log("request timed out")
                }
            });
    });
};

// GET checkGameID checks if a game is occurring for this client
export const checkGameID = async () => {
    const apiURL = process.env.REACT_APP_BACKEND_URL;
    const serverURL = `${apiURL}/game`;
    const config = { withCredentials: true };

    try {
        return new Promise((resolve) => {
            axios.get(serverURL, config)
                .then((response) => {
                    if (response.status === 202 || response.status === 200) {
                        const gameID = response.data.gameID;
                        resolve(gameID);
                    } else {
                        resolve(null);
                    }
                });
        });
    } catch (error) {
        if (error.status === 408){
            console.log("request timed out")
        } else{
            console.error("checkGameID error: ", error)
        }
     };
}