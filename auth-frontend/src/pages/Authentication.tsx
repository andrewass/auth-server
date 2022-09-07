import {useState} from "react";
import SignUp from "../components/authentication/SignUp";
import SignIn from "../components/authentication/SignIn";
import {useNavigate, useSearchParams} from "react-router-dom";

const Authentication = () => {
    const [searchParams] = useSearchParams()
    const navigate = useNavigate()
    const [displaySignUp, setDisplaySignUp] = useState(false)

    const[redirectUri] = useState(searchParams.get("redirect_uri"))

    const redirectToConfirmation = () => {
        navigate("/confirmation", {
            state: {
                redirectUri: redirectUri
            }
        })
    }

    if (displaySignUp) {
        return (
            <div>
                <SignUp setDisplaySignUp={setDisplaySignUp} redirectToConfirmation={redirectToConfirmation}/>
            </div>
        )
    } else {
        return (
            <div>
                <SignIn setDisplaySignUp={setDisplaySignUp} redirectToConfirmation={redirectToConfirmation} />
            </div>
        )
    }
}

export default Authentication