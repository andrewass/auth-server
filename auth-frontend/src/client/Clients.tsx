import ClientRegistrationForm from "./registration/ClientRegistrationForm";
import {ClientList} from "./presentation/ClientList";
import {HStack} from "@chakra-ui/react";

const Clients = () => {

    return (
        <HStack spacing="250px">
            <ClientRegistrationForm/>
            <ClientList/>
        </HStack>
    )
}

export default Clients