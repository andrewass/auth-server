import {AUTH_SERVER_URL} from "../../config/properties";
import {RegisterClientRequest} from "./clientDto";

export const GET_ALL_REGISTERED_CLIENTS = "getAllRegisteredClients";

export const registerClientConfig = (request: RegisterClientRequest) => {
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

export const deleteClientConfig = (clientId: string) => {
    return {
        method: "delete",
        url: AUTH_SERVER_URL + "/clients",
        params: {"client_id" : clientId}
    }
}

export const getRegisteredClientsConfig = (email?: string) => {
    return {
        method: "get",
        url: AUTH_SERVER_URL + "/clients",
        params: {"email": email}
    }
}

export const rotateSecretConfig = (clientId: string) => {
    return {
        method: "post",
        url: AUTH_SERVER_URL + "/client/rotate-secret",
        params: {"client_id": clientId}
    }
}