import {useAuth} from "react-oidc-context";
import {Grid, GridItem, Heading, IconButton, Menu, MenuButton} from "@chakra-ui/react";
import { GiHamburgerMenu } from "react-icons/gi";

export const NavigationBarMobile = () => {
    const auth = useAuth();

    return (
        <Grid templateColumns="repeat(3,1fr)" hideFrom="md"
              minWidth="100%" height="150px" bg="lightblue">
            <GridItem colSpan={1} margin="auto">
                <Heading as="h3" size="lg">
                    Auth Server
                </Heading>
            </GridItem>
            <GridItem colSpan={1} colStart={3} margin="auto">
                <Menu>
                    <MenuButton as={IconButton} icon={<GiHamburgerMenu/>} variant="ghost"/>
                </Menu>
            </GridItem>
        </Grid>
    );
}