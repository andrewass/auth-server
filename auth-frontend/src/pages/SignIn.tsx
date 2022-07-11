const SignIn = () => {

    return (
        <>
            <form>
                <div>
                    <label htmlFor="username">Username</label>
                    <input type="text" name="firstname"/>
                </div>
                <div>
                    <label htmlFor="password">Password</label>
                    <input type="password" name="password"/>
                </div>
                <div>
                    <button type="submit">Sign In</button>
                </div>
            </form>
            <div>
                <button>Create new account</button>
            </div>
        </>
    )
}

export default SignIn