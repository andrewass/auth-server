import {useLocation, useNavigate} from "react-router-dom";

interface CustomizedState {
    redirectUri: string
}

const Confirmation = () => {
    const location = useLocation()
    const navigate = useNavigate()

    const state = location.state as CustomizedState
    const {redirectUri} = state

    const navigateToRedirectUri = () => {
        window.location.replace(redirectUri)
    }

    return(
        <>
            <button onClick={() => navigateToRedirectUri()}>Confirm</button>
        </>
    )
}

export default Confirmation