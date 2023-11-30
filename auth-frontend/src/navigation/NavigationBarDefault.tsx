import {useAuth} from "react-oidc-context";
import {Grid, GridItem, Heading, IconButton} from "@chakra-ui/react";
import {NavLink} from "react-router-dom";
import { BiLogOut } from "react-icons/bi";

export const NavigationBarDefault = () => {
    const auth = useAuth();

    return (
        <Grid templateColumns="repeat(8,1fr)" hideBelow="md"
              minWidth="100%" height="150px" bg="lightblue">
            <GridItem colSpan={2} margin="auto">
                <Heading as="h3" size="lg">
                    Auth Server
                </Heading>
            </GridItem>
            <GridItem colSpan={2} margin="auto">
                <NavLink to="/account">
                    <Heading as="h4" size="md">
                        Account
                    </Heading>
                </NavLink>
            </GridItem>
            <GridItem colSpan={2} margin="auto">
                <NavLink to="/clients">
                    <Heading as="h4" size="md">
                        Clients
                    </Heading>
                </NavLink>
            </GridItem>
            <GridItem colSpan={2} margin="auto">
                <IconButton aria-label="Sign out" as={BiLogOut} variant="ghost"
                            onClick={() => auth.removeUser()}/>
            </GridItem>
        </Grid>
    );
}