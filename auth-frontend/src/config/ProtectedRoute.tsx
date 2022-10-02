import {useAuth} from "react-oidc-context";
import {CircularProgress} from "@chakra-ui/react";
import React from "react";

const ProtectedRoute = ({children} : any) => {
    const auth = useAuth();
    if (auth.isLoading) {
        return <CircularProgress/>
    }
    if (auth.isAuthenticated) {
        return children;
    }
    auth.signinRedirect().catch(error => console.log(error));

    return <></>
}

export default ProtectedRoute;
