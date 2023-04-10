import {Box, Heading} from "@chakra-ui/react";
import {ClientTable} from "./ClientTable";

export const ClientPresentation = () => {
    return (
        <Box maxWidth="500px">
            <Heading as={"h6"} size="md" marginLeft="25px">
                Registered Clients :
            </Heading>
            <ClientTable/>
        </Box>
    );
}