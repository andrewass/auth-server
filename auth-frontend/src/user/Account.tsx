import {Box, VStack} from "@chakra-ui/react";
import NavigationBar from "../navigation/NavigationBar";
import {useQuery} from "react-query";
import {fetchUserInfo} from "./api/userApi";

const Account = () => {

    const {isLoading, error, data} = useQuery('getUserInfo', async () => {
            const response = await fetchUserInfo()
        }
    )

    if (isLoading) return <h5>'Loading...'</h5>

    if (error) return <h5>"An error has occurred"</h5>

    return (
        <VStack margin="auto">
            <NavigationBar/>
            <Box>
                Account details
            </Box>
        </VStack>
    )
}

export default Account