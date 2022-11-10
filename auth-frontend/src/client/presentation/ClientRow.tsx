import {Client} from "./client";


type Props = {
    client : Client
}

export function ClientRow({client}: Props){

    return(
        <p>A Row : {client.email}</p>
    )
}
