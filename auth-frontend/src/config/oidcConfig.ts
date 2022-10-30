import {User} from "oidc-client-ts";

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