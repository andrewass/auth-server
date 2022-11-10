import {useEffect, useState} from "react";
import {useAxiosWrapper} from "../../common/axiosWrapper";
import {getRegisteredClientsConfig} from "../api/clientApi";
import {Client, mapToClient} from "./client";
import {ClientRow} from "./ClientRow";
import {Stack} from "@chakra-ui/react";


export function ClientList(){
    const {get} = useAxiosWrapper()
    const [clientList, setClientList] = useState<Client[]>([])

    const getRegisteredClients = async () => {
        const response = await get(getRegisteredClientsConfig())
        response.map((entry: any) => clientList.push(mapToClient(entry)))
    }

    useEffect(() => {
        getRegisteredClients()
            .catch(error => console.log(error))
    }, [])

    return (
        <Stack maxWidth={500} margin="auto" spacing={15}>
            {clientList.map(client => <ClientRow client={client} />)}
        </Stack>
    )
}