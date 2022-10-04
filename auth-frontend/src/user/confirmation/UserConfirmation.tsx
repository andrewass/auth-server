import {useLocation} from "react-router-dom";

interface CustomizedState {
    redirectUri: string
    state: string
}

const UserConfirmation = () => {
    const location = useLocation()
    const locationState = location.state as CustomizedState
    const {redirectUri, state} = locationState

    const createAuthorizationCode = () => {
        return "dsfdsfasdgrtwewe"
    }

    const navigateToRedirectUri = () => {
        const url = new URL(redirectUri)
        url.searchParams.append("code", createAuthorizationCode())
        url.searchParams.append("state",state)
        window.location.replace(url)
    }

    return (
        <>
            <button onClick={() => navigateToRedirectUri()}>Confirm</button>
        </>
    )
}

export default UserConfirmation