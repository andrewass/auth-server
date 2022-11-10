export class Client {
    email!: string

    constructor(email: string) {
        this.email = email
    }
}


export const mapToClient = (source: any) => {
    return new Client(
        source.user_email
    )
}
