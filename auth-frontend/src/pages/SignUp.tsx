import {useNavigate} from "react-router-dom";

const SignUp = () => {

    const navigate = useNavigate();

    return (
        <>
            <form>
                <div>
                    <label htmlFor="username">Username</label>
                    <input type="text" name="firstname"/>
                </div>
                <div>
                    <label htmlFor="email">E-Mail</label>
                    <input type="email" name="email"/>
                </div>
                <div>
                    <label htmlFor="password">Password</label>
                    <input type="password" name="password"/>
                </div>
                <div>
                    <label htmlFor="confirmPassword">Confirm Password</label>
                    <input type="password" name="confirmPassword"/>
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