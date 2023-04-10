import {Fragment, useState} from "react";
import {useSearchParams} from "react-router-dom";
import {getAuthorizationResponseConfig} from "./api/authenticationApi";
import {useAxiosWrapper} from "../config/axiosWrapper";
import {SignUpForm} from "./SignUpForm";
import {SignInForm} from "./SignInForm";

export const UserAuthentication = () => {
    const {axiosPost} = useAxiosWrapper();
    const [searchParams] = useSearchParams();

    const [displaySignUp, setDisplaySignUp] = useState(false);
    const redirectUri = searchParams.get("redirect_uri");
    const state = searchParams.get("state");
    const clientId = String(searchParams.get("client_id"));
    const codeChallenge = String(searchParams.get("code_challenge"));

    const codeChallengeMethod = String(searchParams.get("code_challenge_method"));

    const postProcessAuthenticationWithAuthCode = async (email: string) => {
        const url = new URL(String(redirectUri));
        const response = await axiosPost(getAuthorizationResponseConfig(
            email, clientId, codeChallenge, codeChallengeMethod
        ));
        url.searchParams.append("code", response.code);
        url.searchParams.append("state", String(state));
        window.location.replace(url);
    }

    if (displaySignUp) {
        return (
            <Fragment>
                <SignUpForm setDisplaySignUp={setDisplaySignUp}
                            postProcessSignUp={postProcessAuthenticationWithAuthCode}/>
            </Fragment>
        );
    } else {
        return (
            <Fragment>
                <SignInForm setDisplaySignUp={setDisplaySignUp}
                            postProcessSignIn={postProcessAuthenticationWithAuthCode}/>
            </Fragment>
        );
    }
}