import {useState} from "react";
import SignUp from "./SignUp";
import SignIn from "./SignIn";
import {useNavigate, useSearchParams} from "react-router-dom";

const UserAuthentication = () => {
    const [searchParams] = useSearchParams()
    const navigate = useNavigate()
    const [displaySignUp, setDisplaySignUp] = useState(false)
    const redirectUri = searchParams.get("redirect_uri")
    const state = searchParams.get("state")

    const redirectToClient = () => {
        const url = new URL(String(redirectUri))
        url.searchParams.append("code", "dsfjasklrjelwkjrk3j24hjk3h24kjh423sdfsdfdsf")
        url.searchParams.append("state",String(state))
        window.location.replace(url)
    }

    const redirectToAccount = () => {
        navigate("/account")
    }

    if (displaySignUp) {
        return (
            <div>
                <SignUp setDisplaySignUp={setDisplaySignUp} redirectToPage={redirectToClient}/>
            </div>
        )
    } else {
        return (
            <div>
                <SignIn setDisplaySignUp={setDisplaySignUp} redirectToPage={redirectToClient}/>
            </div>
        )
    }
}

export default UserAuthentication