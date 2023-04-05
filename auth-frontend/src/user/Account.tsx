import {Box, VStack} from "@chakra-ui/react";
import {NavigationBar} from "../navigation/NavigationBar";
import {useQuery} from "react-query";
import {getUserInfoConfig} from "./api/userApi";
import {useAxiosWrapper} from "../config/axiosWrapper";

export const Account = () => {
    const {axiosGet} = useAxiosWrapper();

    const {isLoading, error, data} = useQuery('getUserInfo',  () => axiosGet(getUserInfoConfig()));

    if (isLoading) return <h5>'Loading...'</h5>

    if (error) return <h5>"An error has occurred"</h5>

    const {email} = data

    return (
        <VStack margin="auto">
            <NavigationBar/>
            <Box>
                Account details : email : {email}
            </Box>
        </VStack>
    );
}