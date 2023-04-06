import {Box, VStack} from "@chakra-ui/react";
import {NavigationBar} from "../navigation/NavigationBar";
import {useQuery} from "@tanstack/react-query";
import {GET_USER_INFO, getUserInfoConfig} from "./api/userApi";
import {useAxiosWrapper} from "../config/axiosWrapper";
import {UserDetails} from "./userTypes";

export const Account = () => {
    const {axiosGet} = useAxiosWrapper();

    const {isLoading, isError, data} = useQuery<UserDetails, Error>({
        queryKey: [GET_USER_INFO],
        queryFn: () => axiosGet(getUserInfoConfig())
    });

    if (isLoading) return <h5>'Loading...'</h5>

    if (isError) return <h5>"An error has occurred"</h5>

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