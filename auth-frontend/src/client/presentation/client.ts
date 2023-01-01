export class Client {
    email: string
    name: string

    constructor(email: string, name: string) {
        this.email = email
        this.name = name
    }
}


export const mapToClient = (source: any) => {
    return new Client(
        source.user_email,
        source.client_name
    )
}
