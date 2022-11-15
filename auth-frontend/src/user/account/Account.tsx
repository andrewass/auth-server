import {Box, VStack} from "@chakra-ui/react";
import NavigationBar from "../../navigation/NavigationBar";

const Account = () => {

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