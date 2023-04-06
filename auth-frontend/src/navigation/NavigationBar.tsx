import {Heading, HStack, IconButton} from "@chakra-ui/react";
import {NavLink} from "react-router-dom";
import {BiLogOut} from "react-icons/bi";
import {useAuth} from "react-oidc-context";

export const NavigationBar = () => {
    const auth = useAuth();

    return (
        <HStack width="100%" spacing="100px" backgroundColor="lightblue" justify="center" height="150px">
            <NavLink to="/account">
                <Heading as={"h4"} size="md">
                    Account
                </Heading>
            </NavLink>
            <NavLink to="/clients">
                <Heading as={"h4"} size="md">
                    Clients
                </Heading>
            </NavLink>
            <IconButton aria-label="Sign out" as={BiLogOut} variant="ghost"
                        onClick={() => auth.removeUser()}/>
        </HStack>
    );
}