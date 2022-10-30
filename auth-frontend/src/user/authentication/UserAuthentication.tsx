import {useState} from "react";
import SignUp from "./SignUp";
import SignIn from "./SignIn";
import {useSearchParams} from "react-router-dom";
import {getAuthorizationResponse} from "../api/userApi";
import {useAxiosWrapper} from "../../config/axiosWrapper";

const UserAuthentication = () => {
    const [searchParams] = useSearchParams()
    const [displaySignUp, setDisplaySignUp] = useState(false)
    const redirectUri = searchParams.get("redirect_uri")
    const state = searchParams.get("state")
    const clientId = String(searchParams.get("client_id"))
    const {post} = useAxiosWrapper()

    const postProcessAuthenticationWithAuthCode = async (email: string) => {
        const url = new URL(String(redirectUri))
        const response = await post(getAuthorizationResponse(email, clientId))
        url.searchParams.append("code", response.code)
        url.searchParams.append("state", String(state))
        window.location.replace(url)
    }

    if (displaySignUp) {
        return (
            <div>
                <SignUp setDisplaySignUp={setDisplaySignUp} postProcessSignUp={postProcessAuthenticationWithAuthCode}/>
            </div>
        )
    } else {
        return (
            <div>
                <SignIn setDisplaySignUp={setDisplaySignUp} postProcessSignIn={postProcessAuthenticationWithAuthCode}/>
            </div>
        )
    }
}

export default UserAuthentication