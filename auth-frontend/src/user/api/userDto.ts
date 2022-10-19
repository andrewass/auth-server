
export interface AuthorizeUserRequest{
    clientId: string
    redirectUri: string
    state: string
}