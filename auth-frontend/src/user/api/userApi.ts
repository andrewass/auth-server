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
    return {
        url: AUTH_SERVER_URL + "/authorization-confirmation",
        method: "post",
        data: {
            "client_id": request.clientId,
            "redirect_uri": request.redirectUri,
            "state": request.state
        }
    }
}
