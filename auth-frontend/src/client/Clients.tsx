import {Box, SimpleGrid} from "@chakra-ui/react";
import {ClientRegistration} from "./registration/ClientRegistration";
import {NavigationBar} from "../navigation/NavigationBar";
import {ClientPresentation} from "./presentation/ClientPresentation";

export const Clients = () => {
    return (
        <Box>
            <NavigationBar/>
            <SimpleGrid columns={{base: 1, md: 2}} spacing="50px"
                        marginTop="80px" marginLeft="100px" marginRight="100px">
                <ClientRegistration/>
                <ClientPresentation/>
            </SimpleGrid>
        </Box>
    );
}