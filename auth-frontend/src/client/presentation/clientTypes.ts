export type ClientTypes = {
    email: string
    name: string
    clientId: string
    clientSecret: string
}

export const mapToClient = (source: any): ClientTypes => {
    return {
        clientId: source.client_id,
        clientSecret: source.client_secret,
        email: source.user_email,
        name: source.client_name
    }
}
