import {useAxiosWrapper} from "../../config/axiosWrapper";
import {getRegisteredClientsConfig} from "../api/clientApi";
import {Client, mapToClient} from "./client";
import {ClientRow} from "./ClientRow";
import {Heading, List, Stack, Table, TableCaption, TableContainer, Tbody, Th, Thead, Tr} from "@chakra-ui/react";
import {useNavigate} from "react-router-dom";
import {useAuth} from "react-oidc-context";
import {useQuery} from "react-query";


export function ClientList() {
    const {axiosGet} = useAxiosWrapper()
    const navigate = useNavigate()
    const {user} = useAuth()

    const navigateToDetails = (client: Client) => {
        navigate("/client/details")
    }

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
                        <ClientRow client={mapToClient(clientData)} navigateToDetails={navigateToDetails}/>)}
                </Tbody>
            </Table>
        </TableContainer>
    )
}