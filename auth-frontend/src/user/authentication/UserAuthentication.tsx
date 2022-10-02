import {useState} from "react";
import SignUp from "./SignUp";
import SignIn from "./SignIn";
import {useNavigate, useSearchParams} from "react-router-dom";

const UserAuthentication = () => {
    const navigate = useNavigate()
    const [displaySignUp, setDisplaySignUp] = useState(false)

    const redirectToAccount = () => {
        navigate("/account")
    }

    if (displaySignUp) {
        return (
            <div>
                <SignUp setDisplaySignUp={setDisplaySignUp} redirectToPage={redirectToAccount}/>
            </div>
        )
    } else {
        return (
            <div>
                <SignIn setDisplaySignUp={setDisplaySignUp} redirectToPage={redirectToAccount}/>
            </div>
        )
    }
}

export default UserAuthentication