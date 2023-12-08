import axios from 'axios';

export const checkSession = () => {
    const serverURL = "http://localhost:8080/authUser"
    const config = { withCredentials: true};

    return new Promise((resolve) => {
        axios.get(serverURL, config)
        .then((response) => {
            if (response.status === 202) {
                const name = response.data.username;
                resolve(name)    
            }else{
                resolve(null);
            }
         })
    });  
};