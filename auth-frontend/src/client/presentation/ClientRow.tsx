import {Client} from "./client";
import {IconButton, Td, Tr} from "@chakra-ui/react";
import {TiDelete} from "react-icons/ti";
import {useMutation} from "react-query";
import {queryClient} from "../../App";
import {deleteClientConfig} from "../api/clientApi";
import {useAxiosWrapper} from "../../config/axiosWrapper";


interface Props {
    client: Client
    navigateToDetails: (client: Client) => void
}

export function ClientRow({client, navigateToDetails}: Props) {

    const {axiosDelete} = useAxiosWrapper()

    const deleteClient = async () => {
        await axiosDelete(deleteClientConfig("fdfsdf"))
    }

    const mutation = useMutation(deleteClient, {
        onSuccess: async () => {
            await queryClient.invalidateQueries("getUserClients");
        },
        onError: (error) => console.log("Unable to delete client : " + error)
    })

    return (
        <Tr>
            <Td>{client.name}</Td>
            <Td>'
                <IconButton
                    onClick={() => mutation.mutate()}
                    aria-label="Delete client"
                    size="sm"
                    icon={<TiDelete size="25px"/>}/>
            </Td>
        </Tr>
    )
}
