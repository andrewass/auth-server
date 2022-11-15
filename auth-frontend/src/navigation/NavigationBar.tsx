import {HStack, IconButton} from "@chakra-ui/react";
import {Link} from "react-router-dom";
import {BiLogOut} from "react-icons/bi";
import {useAuth} from "react-oidc-context";

const NavigationBar = () => {

    const auth = useAuth()

    return (
        <HStack width="100%" spacing="100px" backgroundColor="lightblue" justify="center" height="150px">
            <Link to="/account">Account</Link>
            <Link to="/clients">Clients</Link>
            <IconButton aria-label="Sign out" as={BiLogOut} variant="ghost"
                        onClick={() => auth.removeUser()}/>
        </HStack>
    )
}

export default NavigationBar