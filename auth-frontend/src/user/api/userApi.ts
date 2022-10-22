import {AUTH_SERVER_URL} from "../../config/properties";

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

export const getAuthorizationResponse = (email: string, clientId: string) => {
    return {
        url: AUTH_SERVER_URL + "/authorization-response",
        method: "post",
        data: {
            "email": email,
            "client_id": clientId
        }
    }
}
