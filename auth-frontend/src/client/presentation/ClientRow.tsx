import {Client} from "./client";
import {Box, ListItem} from "@chakra-ui/react";


interface Props {
    client: Client
    navigateToDetails: (client: Client) => void
}

export function ClientRow({client, navigateToDetails}: Props) {


    return (
        <ListItem onClick={() => navigateToDetails(client)}>
            <Box>
                Client : {client.email}
            </Box>
        </ListItem>
    )
}
