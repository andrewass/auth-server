import {AUTH_SERVER_URL} from "../../config/properties";

const registerClientConfig = (requestData: any) => {
    return {
        method: "post",
        url: AUTH_SERVER_URL+"/clients",
        data: requestData
    }
}

export {
    registerClientConfig
}