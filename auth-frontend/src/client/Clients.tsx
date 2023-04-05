import {ClientList} from "./presentation/ClientList";
import {Box, HStack} from "@chakra-ui/react";
import ClientRegistration from "./registration/ClientRegistration";
import {NavigationBar} from "../navigation/NavigationBar";

export const Clients = () => {
    return (
        <Box>
            <NavigationBar/>
            <HStack spacing="250px" justify="center" marginTop="100px">
                <ClientRegistration/>
                <ClientList/>
            </HStack>
        </Box>
    );
}