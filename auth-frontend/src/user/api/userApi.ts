import {AUTH_SERVER_URL} from "../../config/properties";

export const GET_USER_INFO = "getUserInfo";

export const getUserInfoConfig = () => {
    return {
        method: "get",
        url: AUTH_SERVER_URL + "/user/info"
    }
}