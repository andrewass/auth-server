import {useAxiosWrapper} from "../../config/axiosWrapper";
import {getRegisteredClientsConfig} from "../api/clientApi";
import {mapToClient} from "./clientTypes";
import {ClientRow} from "./ClientRow";
import {Table, TableContainer, Tbody, Th, Thead, Tr} from "@chakra-ui/react";
import {useAuth} from "react-oidc-context";
import {useQuery} from "react-query";


export function ClientList() {
    const {axiosGet} = useAxiosWrapper()
    const {user} = useAuth()

    const {isLoading, error, data} = useQuery(['getUserClients', user!.profile.email],
        () => axiosGet(getRegisteredClientsConfig(user!.profile.email)))

    if (isLoading) return <h5>'Loading...'</h5>

    if (error) return <h5>"An error has occurred"</h5>

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
                        <ClientRow client={mapToClient(clientData)}/>)}
                </Tbody>
            </Table>
        </TableContainer>
    );
}