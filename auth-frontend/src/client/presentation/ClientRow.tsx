import {Button, IconButton, Td, Tr} from "@chakra-ui/react";
import {TiDelete} from "react-icons/ti";
import {useMutation} from "@tanstack/react-query";
import {queryClient} from "../../App";
import {deleteClientConfig, GET_ALL_REGISTERED_CLIENTS} from "../api/clientApi";
import {useAxiosWrapper} from "../../config/axiosWrapper";
import {useNavigate} from "react-router-dom";
import {ClientDetails} from "../clientTypes";


export function ClientRow({client}: { client: ClientDetails }) {
    const {axiosDelete} = useAxiosWrapper();
    const navigate = useNavigate();

    const mutation = useMutation({
        mutationFn: async () => await axiosDelete(deleteClientConfig(client.clientId)),
        onSuccess: async () => {
            await queryClient.invalidateQueries([GET_ALL_REGISTERED_CLIENTS]);
        },
        onError: (error) => console.log("Unable to delete client : " + error)
    });

    return (
        <Tr>
            <Td>
                <Button onClick={() => navigate("/client/details",
                    {state: client})} variant="ghost" _hover={{bg: "white"}}>
                    {client.name}
                </Button>
            </Td>
            <Td>'
                <IconButton
                    onClick={() => mutation.mutate()} aria-label="Delete client" size="sm"
                    icon={<TiDelete size="25px"/>} variant="ghost" _hover={{bg: "white"}}/>
            </Td>
        </Tr>
    );
}
