export type ClientDetails = {
    email: string
    name: string
    clientId: string
    clientSecret: string
}

export const mapToClientDetails = (source: any): ClientDetails => {
    return {
        clientId: source.client_id,
        clientSecret: source.client_secret,
        email: source.user_email,
        name: source.client_name
    }
}
