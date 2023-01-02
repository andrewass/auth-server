import {Client} from "./client";
import {Button, IconButton, Td, Tr} from "@chakra-ui/react";
import {TiDelete} from "react-icons/ti";
import {useMutation} from "react-query";
import {queryClient} from "../../App";
import {deleteClientConfig} from "../api/clientApi";
import {useAxiosWrapper} from "../../config/axiosWrapper";
import {useNavigate} from "react-router-dom";


interface Props {
    client: Client
}

export function ClientRow({client}: Props) {
    const {axiosDelete} = useAxiosWrapper()
    const navigate = useNavigate()

    const naviagateToDetails = () => {
        navigate("/client/details", {state: client})
    }

    const deleteClient = async () => {
        await axiosDelete(deleteClientConfig(client.clientId))
    }

    const mutation = useMutation(deleteClient, {
        onSuccess: async () => {
            await queryClient.invalidateQueries("getUserClients");
        },
        onError: (error) => console.log("Unable to delete client : " + error)
    })

    return (
        <Tr>
            <Td>
                <Button onClick={() => navigate("/client/details", {state: client})}>
                    {client.name}
                </Button>
            </Td>
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
