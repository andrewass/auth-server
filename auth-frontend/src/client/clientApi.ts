import {AUTH_SERVER_URL} from "../config/properties";

const registerClientConfig = (myData: any) => {
    return {
        method: "post",
        url: AUTH_SERVER_URL+"/",
        data: myData
    }
}

export {
    registerClientConfig
}