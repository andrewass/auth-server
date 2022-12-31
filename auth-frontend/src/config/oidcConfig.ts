import {User} from "oidc-client-ts";

export const oidcConfig = {
    authority: "http://authserver.io",
    client_id: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQF",
    client_secret: "DaFpLSjFbcXoEFfRsWxPLDnJObCsNV",
    redirect_uri: "http://authfrontend.io/account"
};

export const onSignInCallback = (_user: User | void): void => {
    window.history.replaceState(
        {},
        document.title,
        window.location.pathname
    )
}