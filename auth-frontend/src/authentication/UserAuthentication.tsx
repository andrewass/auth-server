import {Fragment, useState} from "react";
import SignUp from "./SignUp";
import SignIn from "./SignIn";
import {useSearchParams} from "react-router-dom";
import {getAuthorizationResponse} from "./api/authenticationApi";
import {useAxiosWrapper} from "../config/axiosWrapper";

const UserAuthentication = () => {
    const [searchParams] = useSearchParams()
    const [displaySignUp, setDisplaySignUp] = useState(false)
    const redirectUri = searchParams.get("redirect_uri")
    const state = searchParams.get("state")
    const clientId = String(searchParams.get("client_id"))
    const codeChallenge = String(searchParams.get("code_challenge"))
    const codeChallengeMethod = String(searchParams.get("code_challenge_method"))

    const {post} = useAxiosWrapper()

    const postProcessAuthenticationWithAuthCode = async (email: string) => {
        const url = new URL(String(redirectUri))
        const response = await post(getAuthorizationResponse(
            email, clientId, codeChallenge, codeChallengeMethod
        ))
        url.searchParams.append("code", response.code)
        url.searchParams.append("state", String(state))
        window.location.replace(url)
    }

    if (displaySignUp) {
        return (
            <Fragment>
                <SignUp setDisplaySignUp={setDisplaySignUp} postProcessSignUp={postProcessAuthenticationWithAuthCode}/>
            </Fragment>
        )
    } else {
        return (
            <Fragment>
                <SignIn setDisplaySignUp={setDisplaySignUp} postProcessSignIn={postProcessAuthenticationWithAuthCode}/>
            </Fragment>
        )
    }
}

export default UserAuthentication