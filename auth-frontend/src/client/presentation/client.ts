export interface Client {
    email: string
    name: string
    clientId: string
}

export const mapToClient = (source: any): Client => {
    return {
        clientId: source.client_id,
        email: source.user_email,
        name: source.client_name
    }
}
