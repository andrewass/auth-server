import {useEffect, useState} from "react";
import {useAxiosWrapper} from "../../config/axiosWrapper";
import {getRegisteredClientsConfig} from "../api/clientApi";


const ClientList = () => {
    const {get} = useAxiosWrapper()
    const [clientList, setClientList] = useState([])

    const getRegisteredClients = async () => {
        const response = await get(getRegisteredClientsConfig())
        setClientList(response)
    }

    useEffect(() => {
        getRegisteredClients()
            .catch(error => console.log(error))
    }, [])

    return (
        <p>Client List</p>
    )
}

export default ClientList