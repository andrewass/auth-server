import {AUTH_SERVER_URL, SESSION_STORAGE_KEY} from "../../config/properties";
import {RegisterClientRequest} from "./clientDto";

const registerClientConfig = (request: RegisterClientRequest) => {
    return {
        method: "post",
        url: AUTH_SERVER_URL + "/clients",
        data: {
            "redirect_uris": request.redirectUris,
            "client_name": request.clientName
        }
    }
}

const getRegisteredClientsConfig = () => {
    const response: any = sessionStorage.getItem(SESSION_STORAGE_KEY)
    const responseJson = JSON.parse(response)
    const {profile} = responseJson

    return {
        method: "get",
        url: AUTH_SERVER_URL + "/clients",
        params: {"email": profile.email}
    }
}

export {
    registerClientConfig, getRegisteredClientsConfig
}