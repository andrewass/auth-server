import {AxiosRequestConfig} from "axios";
import {AUTH_SERVER_URL} from "../../config/properties";
import {AuthorizeUserRequest} from "./userDto";

export const signUpUser = (email: string, password: string): Promise<Response> => {
    const url = AUTH_SERVER_URL + "/user/sign-up"
    return fetch(url, {
        method: "POST",
        body: JSON.stringify({email, password})
    })
}

export const signInUser = (email: string, password: string): Promise<Response> => {
    const url = AUTH_SERVER_URL + "/user/sign-in"
    return fetch(url, {
        method: "POST",
        body: JSON.stringify({email, password})
    })
}

export const postConfirmation = (request: AuthorizeUserRequest) => {
    const url = AUTH_SERVER_URL + "/authorization-confirmation"
    return {
        method: "post"
    }
}

export const attachClientToUser = (requestData: AxiosRequestConfig) => {
    return {
        method: "post",
        url: AUTH_SERVER_URL + "/user/add-client",
        data: requestData
    }
}