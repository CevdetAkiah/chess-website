import axios from 'axios';

export const checkSession = () => {
    const serverURL = "http://localhost:8080/authUser";
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
                
            });
    });
};

export const checkGameID = async () => {
    const serverURL = "http://localhost:8080/gameID";
    const config = { withCredentials: true };

    try {
        return await new Promise((resolve) => {
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
    } catch (error) { }
}