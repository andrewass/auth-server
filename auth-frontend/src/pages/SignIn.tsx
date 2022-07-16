import {useNavigate} from "react-router-dom";
import {useState} from "react";

const SignIn = () => {
    const navigate = useNavigate();

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const submitSignIn = () => {

    }

    return (
        <>
            <form>
                <div>
                    <label htmlFor="username">Username</label>
                    <input type="text" name="firstname"
                           onChange={(e) => setUsername(e.target.value)}/>
                </div>
                <div>
                    <label htmlFor="password">Password</label>
                    <input type="password" name="password"
                           onChange={(e) => setPassword(e.target.value)}/>
                </div>
                <div>
                    <button type="submit" onClick={submitSignIn}>Sign In</button>
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