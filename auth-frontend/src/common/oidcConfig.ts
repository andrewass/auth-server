import {User} from "oidc-client-ts";
import {SESSION_STORAGE_KEY} from "./properties";


interface SessionProfile {
    email: string
}

interface Session {
    profile: SessionProfile
}

export const oidcConfig = {
    authority: "http://localhost:8089",
    client_id: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQF",
    redirect_uri: "http://localhost:8090/account"
};

export const onSignInCallback = (_user: User | void): void => {
    window.history.replaceState(
        {},
        document.title,
        window.location.pathname
    )
}

export const getSessionEmail = () : string => {
    const sessionData =  String(sessionStorage.getItem(SESSION_STORAGE_KEY))
    const session : Session = JSON.parse(sessionData) as Session
    return session.profile.email
}