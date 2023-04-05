import {useAuth} from "react-oidc-context";
import axios, {AxiosRequestConfig} from "axios";

export const useAxiosWrapper = () => {
    const auth = useAuth()

    function request() {
        return async (config: AxiosRequestConfig) => {
            config.headers = {Authorization: "Bearer " + auth.user?.access_token}
            const response = await axios(config);
            return response.data;
        }
    }

    return {
        axiosGet: request(),
        axiosPost: request(),
        axiosPut: request(),
        axiosDelete: request()
    }
}