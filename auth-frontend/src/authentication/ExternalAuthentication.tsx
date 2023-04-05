import {Stack} from "@chakra-ui/react";
import { UserAuthentication } from "./UserAuthentication";

export const ExternalAuthentication = () => {

    return (
        <Stack maxWidth={500} margin="auto" spacing={5} marginTop={90}>
            <p>External app</p>
            <UserAuthentication/>
            <p>Scope of app</p>
        </Stack>
    );
}