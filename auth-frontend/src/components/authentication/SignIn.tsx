import React, {useState} from "react";
import {signInUser} from "../../api/authentication";

interface Props {
    setDisplaySignUp: (value: boolean) => void
    redirectToConfirmation: () => void
}

const SignIn = ({setDisplaySignUp, redirectToConfirmation}: Props) => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    const submitSignIn = (event: React.FormEvent<HTMLElement>) => {
        event.preventDefault()
        signInUser({email, password})
            .then(() => redirectToConfirmation())
            .catch(error => console.log(error))
    }

    return (
        <>
            <form onSubmit={submitSignIn}>
                <div>
                    <label>E-mail</label>
                    <input type="text" onChange={(e) => setEmail(e.target.value)}/>
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
                <button onClick={() => setDisplaySignUp(true)}>
                    Create new account
                </button>
            </div>
        </>
    )
}

export default SignIn