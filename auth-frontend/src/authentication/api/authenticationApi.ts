import {AUTH_SERVER_URL} from "../../config/properties";

export const signUpUser = (email: string, password: string): Promise<Response> => {
    const url = AUTH_SERVER_URL + "/user/sign-up"
    return fetch(url, {
        method: "POST",
        body: JSON.stringify({email, password})
    })
}

export const signInUserConfig = (email: string, password: string) => {
    return {
        url: AUTH_SERVER_URL + "/user/sign-in",
        method: "POST",
        data: JSON.stringify({email, password})
    }
}

export const getAuthorizationResponseConfig = (
    email: string, clientId: string, codeChallenge: string, codeChallengeMethod: string
) => {
    return {
        url: AUTH_SERVER_URL + "/authorization-response",
        method: "post",
        data: {
            "email": email,
            "client_id": clientId,
            "code_challenge": codeChallenge,
            "code_challenge_method": codeChallengeMethod
        }
    }
}
