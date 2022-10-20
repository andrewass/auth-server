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
    const {post} = useAxiosWrapper()

    const postProcessAuthentication = (email: string) => {
        const url = new URL(String(redirectUri))
        const authorizationResponse = post(getAuthorizationResponse(email))
        url.searchParams.append("code", "dsfjasklrjelwkjrk3j24hjk3h24kjh423sdfsdfdsf")
        url.searchParams.append("state",String(state))
        window.location.replace(url)
    }

    if (displaySignUp) {
        return (
            <div>
                <SignUp setDisplaySignUp={setDisplaySignUp} postProcessSignUp={postProcessAuthentication}/>
            </div>
        )
    } else {
        return (
            <div>
                <SignIn setDisplaySignUp={setDisplaySignUp} postProcessSignIn={postProcessAuthentication}/>
            </div>
        )
    }
}

export default UserAuthentication