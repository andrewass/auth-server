import UserAuthentication from "./UserAuthentication";
import {Stack} from "@chakra-ui/react";

const ExternalAuthentication = () => {

    return (
        <Stack maxWidth={500} margin="auto" spacing={5} marginTop={90}>
            <p>External app</p>
            <UserAuthentication/>
            <p>Scope of app</p>
        </Stack>
    )
}

export default ExternalAuthentication