export const BASE_URL = process.env.REACT_APP_BASE_URL
export const REDIRECT_URI = String(BASE_URL + "/account")
export const AUTH_SERVER_URL = process.env.REACT_APP_AUTH_SERVER_URL
export const CLIENT_ID  = String(process.env.REACT_APP_CLIENT_ID)
export const SESSION_STORAGE_KEY = "oidc.user:" + AUTH_SERVER_URL + ":" + CLIENT_ID