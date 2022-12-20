import {AUTH_SERVER_URL} from "../../config/properties";

type UserInfo = {
    username: string
}

export const fetchUserInfo = async (): Promise<UserInfo> => {
    const url = AUTH_SERVER_URL + "/user/info"
    const response = await fetch(url)
    return response.json()
}
