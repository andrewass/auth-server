import {useState} from "react";
import SignUp from "./SignUp";
import SignIn from "./SignIn";
import {useNavigate, useSearchParams} from "react-router-dom";

const UserAuthentication = () => {
    const [searchParams] = useSearchParams()
    const navigate = useNavigate()
    const [displaySignUp, setDisplaySignUp] = useState(false)

    const[redirectUri] = useState(searchParams.get("redirect_uri"))
    const [state] = useState(searchParams.get("state"));

    const redirectToConfirmation = () => {
        navigate("/confirmation", {
            state: {
                redirectUri: redirectUri,
                state: state
            }
        })
    }

    if (displaySignUp) {
        return (
            <div>
                <SignUp setDisplaySignUp={setDisplaySignUp} redirectToPage={redirectToConfirmation}/>
            </div>
        )
    } else {
        return (
            <div>
                <SignIn setDisplaySignUp={setDisplaySignUp} redirectToPage={redirectToConfirmation} />
            </div>
        )
    }
}

export default UserAuthentication