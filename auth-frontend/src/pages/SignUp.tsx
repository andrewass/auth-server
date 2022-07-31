import {useNavigate} from "react-router-dom";
import React, {useState} from "react";
import {signUpUser} from "../api/authentication";

const SignUp = () => {
    const navigate = useNavigate()

    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [confirmedPassword, setConfirmedPassword] = useState("")

    const submitSignUp = (event: React.FormEvent<HTMLElement>) => {
        event.preventDefault()
        if (password !== confirmedPassword) {
            alert("Password mismatch")
        } else {
            signUpUser({email, password})
                .then(response => {
                    navigate("/confirmation", {state: })
                })
                .catch(error => console.log(error))
        }
    }

    return (
        <>
            <form onSubmit={submitSignUp}>
                <div>
                    <label>E-Mail</label>
                    <input type="email" onChange={(e) => setEmail(e.target.value)}/>
                </div>
                <div>
                    <label>Password</label>
                    <input type="password" onChange={(e) => setPassword(e.target.value)}/>
                </div>
                <div>
                    <label>Confirm Password</label>
                    <input type="password" onChange={(e) => setConfirmedPassword(e.target.value)}/>
                </div>
                <div>
                    <button type="submit">Sign Up</button>
                </div>
            </form>
            <div>
                <button onClick={() => navigate("/sign-in")}>
                    Sign in with existing account
                </button>
            </div>
        </>
    )
}

export default SignUp