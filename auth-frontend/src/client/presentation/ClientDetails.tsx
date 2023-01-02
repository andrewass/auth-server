import {Box, Card} from "@chakra-ui/react";
import NavigationBar from "../../navigation/NavigationBar";
import {Client} from "./client";
import {useLocation} from "react-router-dom";

const ClientDetails = () => {
    const {state} = useLocation()
    const client = state as Client

    return (
        <Box>
            <NavigationBar/>
            <Card  justify="center" marginTop="100px">
                <p>Client Details {client.name}</p>
            </Card>
        </Box>
    )
}

export default ClientDetails