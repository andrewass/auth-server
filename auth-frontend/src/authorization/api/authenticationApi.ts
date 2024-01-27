import {AUTH_SERVER_URL} from "../../config/properties";

export const signUpUserConfig = (email: string, password: string) => {
    return {
        url: AUTH_SERVER_URL + "/api/user/sign-up",
        method: "POST",
        data: JSON.stringify({email, password})
    }
}

export const signInUserConfig = (email: string, password: string) => {
    return {
        url: AUTH_SERVER_URL + "/api/user/sign-in",
        method: "POST",
        data: JSON.stringify({email, password})
    }
}

export const getAuthorizationResponseConfig = (
    email: string, clientId: string, codeChallenge: string, codeChallengeMethod: string
) => {
    return {
        url: AUTH_SERVER_URL + "/api/auth/response",
        method: "post",
        data: {
            "email": email,
            "client_id": clientId,
            "code_challenge": codeChallenge,
            "code_challenge_method": codeChallengeMethod
        }
    }
}
