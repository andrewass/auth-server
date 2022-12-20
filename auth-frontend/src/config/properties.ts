export const BASE_URL = "http://authfrontend.io"
export const REDIRECT_URI = String(BASE_URL + "/account")
export const AUTH_SERVER_URL = "http://authserver.io"
export const CLIENT_ID  = String(process.env.REACT_APP_CLIENT_ID)
export const SESSION_STORAGE_KEY = "oidc.user:" + AUTH_SERVER_URL + ":" + CLIENT_ID