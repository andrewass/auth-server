import {useEffect, useState} from "react";
import {useAxiosWrapper} from "../../config/axiosWrapper";
import {getRegisteredClientsConfig} from "../api/clientApi";
import {Client, mapToClient} from "./client";
import {ClientRow} from "./ClientRow";
import {Heading, List, Stack} from "@chakra-ui/react";
import {useNavigate} from "react-router-dom";
import {useAuth} from "react-oidc-context";


export function ClientList() {
    const {get} = useAxiosWrapper()
    const [name, setName] = useState(1)
    const [clientList, setClientList] = useState<Client[]>([])
    const navigate = useNavigate()
    const {user} = useAuth()

    const navigateToDetails = (client: Client) => {
        navigate("/client/details",)
    }

    const getRegisteredClients = async () => {
        const response = await get(getRegisteredClientsConfig(user!.profile.email))
        const clients: Client[] = []
        response.map((clientResponse: any) => clients.push(mapToClient(clientResponse)))
        setClientList(clients)
    }

    useEffect(() => {
        getRegisteredClients()
            .catch(error => console.log(error))
    }, [])

    return (
        <Stack>
            <Heading as="h6" size="s">Registered clients : </Heading>
            <List maxWidth={500} spacing={15}>
                {clientList.map(client => <ClientRow client={client} navigateToDetails={navigateToDetails}/>)}
            </List>
        </Stack>
    )
}