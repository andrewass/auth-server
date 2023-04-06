import {useAxiosWrapper} from "../../config/axiosWrapper";
import {GET_ALL_REGISTERED_CLIENTS, getRegisteredClientsConfig} from "../api/clientApi";
import {ClientRow} from "./ClientRow";
import {Table, TableContainer, Tbody, Th, Thead, Tr} from "@chakra-ui/react";
import {useAuth} from "react-oidc-context";
import {useQuery} from "@tanstack/react-query";
import {mapToClientDetails} from "../clientTypes";


export function ClientList() {
    const {axiosGet} = useAxiosWrapper();
    const {user} = useAuth();

    const {isLoading, isError, data} = useQuery({
        queryKey: [GET_ALL_REGISTERED_CLIENTS],
        queryFn: () => axiosGet(getRegisteredClientsConfig(user!.profile.email))
    });

    if (isLoading) return <h5>'Loading...'</h5>

    if (isError) return <h5>"An error has occurred"</h5>

    return (
        <TableContainer>
            <Table variant="striped" colorScheme="teal">
                <Thead>
                    <Tr>
                        <Th>Client name</Th>
                        <Th>Delete</Th>
                    </Tr>
                </Thead>
                <Tbody>
                    {data.map((clientData: any) =>
                        <ClientRow client={mapToClientDetails(clientData)}/>)}
                </Tbody>
            </Table>
        </TableContainer>
    );
}