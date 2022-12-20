import {AUTH_SERVER_URL} from "../../config/properties";
import {RegisterClientRequest} from "./clientDto";
import {getSessionEmail} from "../../config/oidcConfig";

const registerClientConfig = (request: RegisterClientRequest) => {
    return {
        method: "post",
        url: AUTH_SERVER_URL + "/clients",
        data: {
            "redirect_uris": request.redirectUris,
            "client_name": request.clientName,
            "user_email": request.userEmail
        }
    }
}

const getRegisteredClientsConfig = () => {
    return {
        method: "get",
        url: AUTH_SERVER_URL + "/clients",
        params: {"email": getSessionEmail()}
    }
}

export {
    registerClientConfig, getRegisteredClientsConfig
}