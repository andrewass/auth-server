import {ClientTypes} from "./clientTypes";
import {Button, IconButton, Td, Tr} from "@chakra-ui/react";
import {TiDelete} from "react-icons/ti";
import {useMutation} from "@tanstack/react-query";
import {queryClient} from "../../App";
import {deleteClientConfig} from "../api/clientApi";
import {useAxiosWrapper} from "../../config/axiosWrapper";
import {useNavigate} from "react-router-dom";


export function ClientRow({client}: { client: ClientTypes }) {
    const {axiosDelete} = useAxiosWrapper();
    const navigate = useNavigate();

    const deleteClient = async () => {
        await axiosDelete(deleteClientConfig(client.clientId));
    }

    const mutation = useMutation(deleteClient, {
        onSuccess: async () => {
            await queryClient.invalidateQueries("getUserClients");
        },
        onError: (error) => console.log("Unable to delete client : " + error)
    });

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
    );
}
