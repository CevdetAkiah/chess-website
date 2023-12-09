import axios from 'axios';

export const checkSession = () => {
    const serverURL = "http://localhost:8080/authUser";
    const config = { withCredentials: true };

    return new Promise((resolve) => {
        axios.get(serverURL, config)
            .then((response) => {
                if (response.status === 202) {
                    const userName = response.data.username;
                    resolve(userName);
                }else{
                    resolve(null)
                }
            })
            .catch((error) => {
            });
    });
};