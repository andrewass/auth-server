import {useNavigate, useSearchParams} from "react-router-dom";
import {useState} from "react";

const SignIn = () => {
    const navigate = useNavigate()
    const [searchParams] = useSearchParams()

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const redirectUri = searchParams.get("redirect_uri")

    const submitSignIn = () => {
    }

    return (
        <>
            <form onSubmit={submitSignIn}>
                <div>
                    <label>Username</label>
                    <input type="text" onChange={(e) => setUsername(e.target.value)}/>
                </div>
                <div>
                    <label>Password</label>
                    <input type="password" onChange={(e) => setPassword(e.target.value)}/>
                </div>
                <div>
                    <button type="submit">Sign In</button>
                </div>
            </form>
            <div>
                <button onClick={() => navigate("/sign-up")}>
                    Create new account
                </button>
            </div>
        </>
    )
}

export default SignIn