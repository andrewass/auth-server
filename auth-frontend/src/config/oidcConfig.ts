import {User} from "oidc-client-ts";

export const oidcConfig = {
    authority: "http://localhost:8089",
    client_id: "myClientId",
    redirect_uri: "http://localhost:8090/account",
    onSignInCallback: function (_user: User | void) {
        window.history.replaceState(
            {},
            document.title,
            window.location.pathname
        )
    }
};